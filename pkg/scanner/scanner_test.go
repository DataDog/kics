package scanner

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/Checkmarx/kics/internal/storage"
	"github.com/Checkmarx/kics/internal/tracker"
	"github.com/Checkmarx/kics/pkg/engine"
	"github.com/Checkmarx/kics/pkg/engine/provider"
	"github.com/Checkmarx/kics/pkg/engine/source"
	"github.com/Checkmarx/kics/pkg/kics"
	"github.com/Checkmarx/kics/pkg/parser"
	"github.com/Checkmarx/kics/pkg/progress"
	"github.com/Checkmarx/kics/pkg/resolver"
	"github.com/stretchr/testify/require"

	jsonParser "github.com/Checkmarx/kics/pkg/parser/json"
	terraformParser "github.com/Checkmarx/kics/pkg/parser/terraform"
	yamlParser "github.com/Checkmarx/kics/pkg/parser/yaml"
)

var (
	sourcePath = []string{filepath.FromSlash("../../assets/queries")}
)

type testContext struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func TestScanner_StartScan(t *testing.T) {
	type args struct {
		scanID     string
		noProgress bool
	}
	type fields struct {
		types          []string
		cloudProviders []string
	}

	ctx := context.Background()
	tests := []struct {
		name   string
		args   args
		fields fields
		ctx    testContext
	}{
		{
			name: "testing_start_scan_no_ctx_timeout",
			args: args{
				scanID:     "console",
				noProgress: true,
			},
			fields: fields{
				types:          []string{""},
				cloudProviders: []string{""},
			},
			ctx: createContext(ctx, time.Second*99999),
		},
		{
			name: "testing_start_scan_with_ctx_timeout",
			args: args{
				scanID:     "console",
				noProgress: true,
			},
			fields: fields{
				types:          []string{""},
				cloudProviders: []string{""},
			},
			ctx: createContext(ctx, time.Second*1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.ctx.cancel()
			services, store, err := createServices(tt.fields.types, tt.fields.cloudProviders)
			require.NoError(t, err)
			err = StartScan(tt.ctx.ctx, tt.args.scanID, progress.PbBuilder{}, services)
			require.NoError(t, err)
			require.NotEmpty(t, &store)
		})
	}
}

func createServices(types, cloudProviders []string) (serviceSlice, *storage.MemoryStorage, error) {
	ctx := context.Background()
	filesSource, err := provider.NewFileSystemSourceProvider(ctx, []string{filepath.FromSlash("../../test")}, []string{})
	if err != nil {
		return nil, nil, err
	}

	t, err := tracker.NewTracker(1)
	if err != nil {
		return nil, nil, err
	}
	querySource := source.NewFilesystemSource(ctx, sourcePath, types, cloudProviders, filepath.FromSlash("../../assets/libraries"), true)

	inspector, err := engine.NewInspector(context.Background(),
		querySource, engine.DefaultVulnerabilityBuilder,
		t, &source.QueryInspectorParameters{}, map[string]bool{}, 60, true, true, 1, false)
	if err != nil {
		return nil, nil, err
	}

	// secretsInspector, err := secrets.NewInspector(
	// 	context.Background(),
	// 	map[string]bool{},
	// 	t,
	// 	&source.QueryInspectorParameters{},
	// 	false,
	// 	60,
	// 	assets.SecretsQueryRegexRulesJSON,
	// 	false,
	// )
	// if err != nil {
	// 	return nil, nil, err
	// }

	combinedParser, err := parser.NewBuilder(ctx).
		Add(&jsonParser.Parser{}).
		Add(&yamlParser.Parser{}).
		Add(terraformParser.NewDefault()).
		Build(types, cloudProviders)
	if err != nil {
		return nil, nil, err
	}

	combinedResolver, err := resolver.NewBuilder().
		Build(ctx)
	if err != nil {
		return nil, nil, err
	}

	store := storage.NewMemoryStorage()

	services := make([]*kics.Service, 0, len(combinedParser))

	for _, parser := range combinedParser {
		services = append(services, &kics.Service{
			SourceProvider: filesSource,
			Storage:        store,
			Parser:         parser,
			Inspector:      inspector,
			Tracker:        t,
			Resolver:       combinedResolver,
		})
	}
	return services, store, nil
}

func createContext(ctx context.Context, timeout time.Duration) testContext {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	return testContext{
		ctx,
		cancel,
	}
}

func TestScanner_ConcurrentScans(t *testing.T) {
	ctx := context.Background()

	t.Run("concurrent_scans_same_services", func(t *testing.T) {
		services, store, err := createServices([]string{""}, []string{""})
		require.NoError(t, err)
		require.NotEmpty(t, store)

		numGoroutines := 5
		var wg sync.WaitGroup
		errChan := make(chan error, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				scanID := fmt.Sprintf("concurrent-scan-%d", id)
				err := StartScan(ctx, scanID, progress.PbBuilder{}, services)
				if err != nil {
					errChan <- fmt.Errorf("goroutine %d failed: %w", id, err)
				}
			}(i)
		}

		wg.Wait()
		close(errChan)

		for err := range errChan {
			require.NoError(t, err)
		}
	})

	t.Run("concurrent_scans_different_services", func(t *testing.T) {
		numGoroutines := 3
		var wg sync.WaitGroup
		errChan := make(chan error, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				services, store, err := createServices([]string{""}, []string{""})
				if err != nil {
					errChan <- fmt.Errorf("goroutine %d service creation failed: %w", id, err)
					return
				}
				require.NotEmpty(t, store)

				scanID := fmt.Sprintf("concurrent-diff-services-%d", id)
				err = StartScan(ctx, scanID, progress.PbBuilder{}, services)
				if err != nil {
					errChan <- fmt.Errorf("goroutine %d scan failed: %w", id, err)
				}
			}(i)
		}

		wg.Wait()
		close(errChan)

		for err := range errChan {
			require.NoError(t, err)
		}
	})

	t.Run("concurrent_prepare_and_scan", func(t *testing.T) {
		numGoroutines := 4
		var wg sync.WaitGroup
		errChan := make(chan error, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				services, store, err := createServices([]string{""}, []string{""})
				if err != nil {
					errChan <- fmt.Errorf("goroutine %d service creation failed: %w", id, err)
					return
				}
				require.NotEmpty(t, store)

				scanID := fmt.Sprintf("concurrent-prepare-scan-%d", id)
				err = PrepareAndScan(ctx, scanID, false, 5, progress.PbBuilder{}, services)
				if err != nil {
					errChan <- fmt.Errorf("goroutine %d prepare and scan failed: %w", id, err)
				}
			}(i)
		}

		wg.Wait()
		close(errChan)

		for err := range errChan {
			require.NoError(t, err)
		}
	})
}

func TestScanner_ConcurrentScansWithTimeout(t *testing.T) {
	t.Run("concurrent_scans_with_context_timeout", func(t *testing.T) {
		numGoroutines := 3
		var wg sync.WaitGroup
		errChan := make(chan error, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()

				testCtx := createContext(context.Background(), time.Second*30)
				defer testCtx.cancel()

				services, store, err := createServices([]string{""}, []string{""})
				if err != nil {
					errChan <- fmt.Errorf("goroutine %d service creation failed: %w", id, err)
					return
				}
				require.NotEmpty(t, store)

				scanID := fmt.Sprintf("concurrent-timeout-%d", id)
				err = StartScan(testCtx.ctx, scanID, progress.PbBuilder{}, services)
				if err != nil {
					errChan <- fmt.Errorf("goroutine %d scan failed: %w", id, err)
				}
			}(i)
		}

		wg.Wait()
		close(errChan)

		for err := range errChan {
			require.NoError(t, err)
		}
	})
}

func TestScanner_StressTestConcurrency(t *testing.T) {
	t.Run("stress_test_many_concurrent_scans", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping stress test in short mode")
		}

		ctx := context.Background()
		numGoroutines := 10
		var wg sync.WaitGroup
		errChan := make(chan error, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				services, store, err := createServices([]string{""}, []string{""})
				if err != nil {
					errChan <- fmt.Errorf("goroutine %d service creation failed: %w", id, err)
					return
				}
				require.NotEmpty(t, store)

				scanID := fmt.Sprintf("stress-test-%d", id)
				err = PrepareAndScan(ctx, scanID, false, 5, progress.PbBuilder{}, services)
				if err != nil {
					errChan <- fmt.Errorf("goroutine %d stress test failed: %w", id, err)
				}
			}(i)
		}

		wg.Wait()
		close(errChan)

		for err := range errChan {
			require.NoError(t, err)
		}
	})
}

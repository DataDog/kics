package test

import (
	"path/filepath"
	"testing"

	"github.com/Checkmarx/kics"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/Checkmarx/kics/pkg/scan"
	"github.com/stretchr/testify/require"
)

func Test_E2EExclusions(t *testing.T) {
	tests := []struct {
		name           string
		testFile       string
		expectedOutput scan.ScanStats
	}{
		{
			name:     "no exclusions",
			testFile: filepath.Join("fixtures", "no-exclusions.tf"),
			expectedOutput: scan.ScanStats{
				Violations: 6,
				Files:      1,
				Rules:      1126,
				ViolationBreakdowns: map[string][]string{
					"LOW":    {"c5b31ab9-0f26-4a49-b8aa-4cc064392f4d", "e592a0c5-5bdb-414c-9066-5dba7cdea370", "c5b31ab9-0f26-4a49-b8aa-4cc064392f4d"},
					"MEDIUM": {"f861041c-8c9f-4156-acfc-5e6e524f5884", "568a4d22-3517-44a6-a7ad-6a7eed88722c"},
				},
			},
		},
		{
			name:     "disabled rule inline",
			testFile: filepath.Join("fixtures", "inline-disabled-rule.tf"),
			expectedOutput: scan.ScanStats{
				Violations: 5,
				Files:      1,
				Rules:      1126,
				ViolationBreakdowns: map[string][]string{
					"LOW":    {"c5b31ab9-0f26-4a49-b8aa-4cc064392f4d", "c5b31ab9-0f26-4a49-b8aa-4cc064392f4d"},
					"MEDIUM": {"f861041c-8c9f-4156-acfc-5e6e524f5884", "568a4d22-3517-44a6-a7ad-6a7eed88722c"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metadata, _, err := kics.ExecuteKICSScan(
				[]string{tt.testFile},
				"",
				model.SCIInfo{
					DiffAware: model.DiffAware{
						Enabled: false,
					},
					RunType: "code_update",
					RepositoryCommitInfo: model.RepositoryCommitInfo{
						RepositoryUrl: "test/url",
						CommitSHA:     "test/hash",
						Branch:        "test/branch",
					},
				},
			)
			require.NoError(t, err)
			require.Equal(t, tt.expectedOutput.Violations, metadata.Stats.Violations)
			require.Equal(t, tt.expectedOutput.Files, metadata.Stats.Files)
			require.Equal(t, tt.expectedOutput.Rules, metadata.Stats.Rules)
			require.ElementsMatch(t, tt.expectedOutput.ViolationBreakdowns["LOW"], metadata.Stats.ViolationBreakdowns["LOW"])
			require.ElementsMatch(t, tt.expectedOutput.ViolationBreakdowns["MEDIUM"], metadata.Stats.ViolationBreakdowns["MEDIUM"])
		})
	}

}

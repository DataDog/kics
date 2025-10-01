package test

import (
	"path/filepath"
	"testing"

	"github.com/Checkmarx/kics"
	"github.com/Checkmarx/kics/pkg/featureflags"
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
				Violations: 5,
				Files:      1,
				Rules:      1124,
				ViolationBreakdowns: map[string]map[string]int{
					"INFO":   {
						"a2b3c4d5-e6f7-8901-gh23-ijkl456m7890": 1,
					},
					"LOW":    {
						"c5b31ab9-0f26-4a49-b8aa-4cc064392f4d": 2,
					},
					"MEDIUM": {
						"f861041c-8c9f-4156-acfc-5e6e524f5884": 1,
						"568a4d22-3517-44a6-a7ad-6a7eed88722c": 1,
					},
				},
			},
		},
		{
			name:     "disabled rule inline",
			testFile: filepath.Join("fixtures", "inline-disabled-rule.tf"),
			expectedOutput: scan.ScanStats{
				Violations: 4,
				Files:      1,
				Rules:      1124,
				ViolationBreakdowns: map[string]map[string]int{
					"LOW":    {
						"c5b31ab9-0f26-4a49-b8aa-4cc064392f4d": 2,
					},
					"MEDIUM": {
						"f861041c-8c9f-4156-acfc-5e6e524f5884": 1,
						"568a4d22-3517-44a6-a7ad-6a7eed88722c": 1,
					},
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
				featureflags.NewLocalEvaluator(),
			)
			require.NoError(t, err)
			require.Equal(t, tt.expectedOutput.Violations, metadata.Stats.Violations)
			require.Equal(t, tt.expectedOutput.Files, metadata.Stats.Files)
			require.Equal(t, tt.expectedOutput.Rules, metadata.Stats.Rules)
			require.Equal(t, tt.expectedOutput.ViolationBreakdowns, metadata.Stats.ViolationBreakdowns)
		})
	}

}

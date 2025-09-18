/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */

package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/Checkmarx/kics/internal/console"
	"github.com/Checkmarx/kics/pkg/featureflags"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/Checkmarx/kics/pkg/scan"
	"github.com/rs/zerolog/log"
)

/*
// example local usage
func main() {
	inputPaths := []string{"assets/queries/terraform/gcp/sql_database_instance_does_not_have_skip_show_database/test/"}
	outputPath := "/path/to/dir"
	sci := model.SCIInfo{
		DiffAware: model.DiffAware{
			Enabled: false,
		},
		RunType: "code_update",
		RepositoryCommitInfo: model.RepositoryCommitInfo{
			RepositoryUrl: "github.com/blah",
			Branch:        "main",
			CommitSHA:     "1234567890",
		},
	}
	ExecuteKICSScan(inputPaths, outputPath, sci, featureflags.LocalEvaluator{},true)
}*/

func ExecuteKICSScan(inputPaths []string, outputPath string, sciInfo model.SCIInfo, FlagEvaluator featureflags.FlagEvaluator, consolePrint ...bool) (scan.ScanMetadata, string, error) {
	ctx := context.Background()
	extraInfos := map[string]string{
		"org":        fmt.Sprintf("%d", sciInfo.OrgId),
		"branch":     sciInfo.RepositoryCommitInfo.Branch,
		"sha":        sciInfo.RepositoryCommitInfo.CommitSHA,
		"repository": sciInfo.RepositoryCommitInfo.RepositoryUrl,
	}

	params, logCtx := scan.GetDefaultParameters(ctx, outputPath, extraInfos, consolePrint...)
	params.Path = inputPaths
	params.OutputPath = outputPath
	params.SCIInfo = sciInfo
	params.FlagEvaluator = FlagEvaluator
	metadata, err := console.ExecuteScan(logCtx, params)

	if err != nil {
		log.Error().Int64(
			"org", sciInfo.OrgId,
		).Str(
			"branch", sciInfo.RepositoryCommitInfo.Branch,
		).Str(
			"sha", sciInfo.RepositoryCommitInfo.CommitSHA,
		).Str(
			"repository", sciInfo.RepositoryCommitInfo.RepositoryUrl,
		).Msgf("failed to execute scan: %v", err)
		return scan.ScanMetadata{}, "", err
	}

	log.Info().Int64(
		"org", sciInfo.OrgId,
	).Str(
		"branch", sciInfo.RepositoryCommitInfo.Branch,
	).Str(
		"sha", sciInfo.RepositoryCommitInfo.CommitSHA,
	).Str(
		"repository", sciInfo.RepositoryCommitInfo.RepositoryUrl,
	).Msgf(
		"Scan completed successfully with metadata: %v", metadata,
	)

	resultsFile := filepath.Join(outputPath, params.OutputName)
	return metadata, resultsFile, nil
}

/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */

package main

import (
	"path/filepath"

	"github.com/Checkmarx/kics/internal/console"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/Checkmarx/kics/pkg/scan"
	"github.com/rs/zerolog/log"
)

func main() {
	inputPaths := []string{"assets/queries/terraform/gcp/sql_database_instance_does_not_have_skip_show_database/test/positive.tf"}
	outputPath := "/Users/bahar.shah/dev/kics"
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
	ExecuteKICSScan(inputPaths, outputPath, sci, true, true)
}

func ExecuteKICSScan(inputPaths []string, outputPath string, sciInfo model.SCIInfo, includeRemediations bool, useDetailedDescriptionText bool) (scan.ScanMetadata, string) {
	params := scan.GetDefaultParameters(outputPath)
	params.Path = inputPaths
	params.OutputPath = outputPath
	params.SCIInfo = sciInfo
	params.IncludeRemediations = includeRemediations
	params.UseDetailedDescriptionText = useDetailedDescriptionText
	metadata, err := console.ExecuteScan(params)
	if err != nil {
		log.Fatal().Int64(
			"org", sciInfo.OrgId,
		).Str(
			"branch", sciInfo.RepositoryCommitInfo.Branch,
		).Str(
			"sha", sciInfo.RepositoryCommitInfo.CommitSHA,
		).Str(
			"repository", sciInfo.RepositoryCommitInfo.RepositoryUrl,
		).Msgf("failed to execute scan: %v", err)
		return scan.ScanMetadata{}, ""
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
	return metadata, resultsFile

}

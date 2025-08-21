/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */

package kics

import (
	"fmt"
	"path/filepath"

	"github.com/Checkmarx/kics/internal/console"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/Checkmarx/kics/pkg/scan"
	"github.com/rs/zerolog/log"
)

func ExecuteKICSScan(inputPaths []string, outputPath string, sciInfo model.SCIInfo, consolePrint ...bool) (scan.ScanMetadata, string, error) {
	extraInfos := map[string]string{
		"org":        fmt.Sprintf("%d", sciInfo.OrgId),
		"branch":     sciInfo.RepositoryCommitInfo.Branch,
		"sha":        sciInfo.RepositoryCommitInfo.CommitSHA,
		"repository": sciInfo.RepositoryCommitInfo.RepositoryUrl,
	}

	params := scan.GetDefaultParameters(outputPath, extraInfos, consolePrint...)
	params.Path = inputPaths
	params.OutputPath = outputPath
	params.SCIInfo = sciInfo
	metadata, err := console.ExecuteScan(params)

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

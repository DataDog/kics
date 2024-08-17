/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */

package kics

import (
	"embed"

	"github.com/Checkmarx/kics/internal/console"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/Checkmarx/kics/pkg/scan"
)

//go:embed assets/queries
var queryDir embed.FS

func ExecuteKICSScan(inputPaths []string, outputPath string, sciInfo model.SCIInfo) {
	params := scan.GetDefaultParameters()
	params.Path = inputPaths
	params.OutputPath = outputPath
	params.SCIInfo = sciInfo
	params.QueryDir = queryDir
	console.ExecuteScan(params)
}

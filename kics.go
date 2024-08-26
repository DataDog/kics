/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */

package main

import (
	"log"
	"path/filepath"

	"github.com/Checkmarx/kics/internal/console"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/Checkmarx/kics/pkg/scan"
)

func main() {
	inputPaths := []string{
		"/Users/bahar.shah/go/src/github.com/DataDog/innovation-week-cloud-to-tf/terraform/ami.tf",
	}
	outputPath := "/Users/bahar.shah/go/src/github.com/DataDog/innovation-week-cloud-to-tf"
	sci := model.SCIInfo{
		DiffAware: model.DiffAware{
			Enabled: false,
		},
	}
	ExecuteKICSScan(inputPaths, outputPath, sci)
}

func ExecuteKICSScan(inputPaths []string, outputPath string, sciInfo model.SCIInfo) (scan.ScanMetadata, string) {
	params := scan.GetDefaultParameters()
	params.Path = inputPaths
	params.OutputPath = outputPath
	params.SCIInfo = sciInfo
	metadata, err := console.ExecuteScan(params)
	if err != nil {
		log.Fatalf("failed to execute scan: %v", err)
		return scan.ScanMetadata{}, ""
	}
	log.Printf("Scan completed successfully with metadata: %v", metadata)
	resultsFile := filepath.Join(outputPath, params.OutputName)
	return metadata, resultsFile
}

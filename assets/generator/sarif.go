package main

import (
	"github.com/Checkmarx/kics"
	"github.com/Checkmarx/kics/pkg/model"
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
		RunType: "code_update",
		RepositoryCommitInfo: model.RepositoryCommitInfo{
			RepositoryUrl: "github.com/blah",
			Branch:        "main",
			CommitSHA:     "1234567890",
		},
	}
	kics.ExecuteKICSScan(inputPaths, outputPath, sci)
}

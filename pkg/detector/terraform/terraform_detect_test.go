/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package terraform

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/Checkmarx/kics/pkg/model"
	"github.com/Checkmarx/kics/pkg/utils"
	"github.com/stretchr/testify/require"
)

var OriginalData1 = `resource "aws_s3_bucket" "bucket" {
  bucket = "innovationweek-fancy2023-bucket-${random_id.bucket_id.hex}"
  acl    = "authenticated-read"

  versioning {
    enabled = true
  }

  tags = {
    Demo = "true"
  }
}

resource "aws_s3_bucket" "bucket2" {
  bucket = "innovationweek2-2023-bucket-${random_id.bucket_id.hex}"

  tags = {
    Demo = "true"
    Team = "infrastructure-as-code"
  }
}

resource "aws_s3_bucket" "test3" {
  bucket = "iac-remediation-demo-bucket-notags"
  tags = {
    Demo = "true"
  }
}
`

// TestDetectTerraformLine tests the functions [DetectTerraformLine()] and all the methods called by them
func TestDetectTerraformLine(t *testing.T) { //nolint
	testCases := []struct {
		expected  model.VulnerabilityLines
		searchKey string
		file      *model.FileMetadata
	}{
		{
			expected: model.VulnerabilityLines{
				Line: 3,
				VulnLines: &[]model.CodeLine{
					{
						Position: 2,
						Line:     "  bucket = \"innovationweek-fancy2023-bucket-${random_id.bucket_id.hex}\"",
					},
					{
						Position: 3,
						Line:     "  acl    = \"authenticated-read\"",
					},
					{
						Position: 4,
						Line:     "",
					},
				},
				VulnerablilityLocation: model.ResourceLocation{
					Start: model.ResourceLine{
						Line: 1,
						Col:  1,
					},
					End: model.ResourceLine{
						Line: 12,
						Col:  2,
					},
				},
				LineWithVulnerability: "  acl    = \"authenticated-read\"",
				ResolvedFile:          "",
				ResourceSource:        "resource \"aws_s3_bucket\" \"bucket\" {\n  bucket = \"innovationweek-fancy2023-bucket-${random_id.bucket_id.hex}\"\n  acl    = \"authenticated-read\"\n\n  versioning {\n    enabled = true\n  }\n\n  tags = {\n    Demo = \"true\"\n  }\n}\n",
				FileSource:            strings.Split(OriginalData1, "\n"),
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{
						Line: 1,
						Col:  1,
					},
					End: model.ResourceLine{
						Line: 12,
						Col:  2,
					},
				},
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{
						Line: 3,
						Col:  3,
					},
					End: model.ResourceLine{
						Line: 3,
						Col:  3,
					},
				},
			},
			searchKey: "aws_s3_bucket[bucket].acl",
			file: &model.FileMetadata{
				ScanID:            "Test2",
				ID:                "Test2",
				Kind:              model.KindTerraform,
				OriginalData:      OriginalData1,
				LinesOriginalData: utils.SplitLines(OriginalData1),
			},
		},
		{
			expected: model.VulnerabilityLines{
				Line: 15,
				VulnLines: &[]model.CodeLine{
					{
						Position: 14,
						Line:     "resource \"aws_s3_bucket\" \"bucket2\" {",
					},
					{
						Position: 15,
						Line:     "  bucket = \"innovationweek2-2023-bucket-${random_id.bucket_id.hex}\"",
					},
					{
						Position: 16,
						Line:     "",
					},
				},
				LineWithVulnerability: "  bucket = \"innovationweek2-2023-bucket-${random_id.bucket_id.hex}\"",
				VulnerablilityLocation: model.ResourceLocation{
					Start: model.ResourceLine{
						Line: 14,
						Col:  1,
					},
					End: model.ResourceLine{
						Line: 21,
						Col:  2,
					},
				},
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{
						Line: 14,
						Col:  1,
					},
					End: model.ResourceLine{
						Line: 21,
						Col:  2,
					},
				},
				ResourceSource: "resource \"aws_s3_bucket\" \"bucket2\" {\n  bucket = \"innovationweek2-2023-bucket-${random_id.bucket_id.hex}\"\n\n  tags = {\n    Demo = \"true\"\n    Team = \"infrastructure-as-code\"\n  }\n}\n",
				FileSource:     strings.Split(OriginalData1, "\n"),
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{
						Line: 15,
						Col:  3,
					},
					End: model.ResourceLine{
						Line: 15,
						Col:  3,
					},
				},
			},
			searchKey: "aws_s3_bucket[bucket2].bucket",
			file: &model.FileMetadata{
				ScanID:            "Test2",
				ID:                "Test2",
				Kind:              model.KindTerraform,
				OriginalData:      OriginalData1,
				LinesOriginalData: utils.SplitLines(OriginalData1),
			},
		},
		{
			expected: model.VulnerabilityLines{
				Line: 25,
				VulnLines: &[]model.CodeLine{
					{
						Position: 24,
						Line:     "  bucket = \"iac-remediation-demo-bucket-notags\"",
					},
					{
						Position: 25,
						Line:     "  tags = {",
					},
					{
						Position: 26,
						Line:     "    Demo = \"true\"",
					},
				},
				VulnerablilityLocation: model.ResourceLocation{
					Start: model.ResourceLine{
						Line: 23,
						Col:  1,
					},
					End: model.ResourceLine{
						Line: 28,
						Col:  2,
					},
				},
				LineWithVulnerability: "  tags = {",
				ResolvedFile:          "",
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{
						Line: 26,
						Col:  5,
					},
					End: model.ResourceLine{
						Line: 26,
						Col:  5,
					},
				},
				ResourceSource: "resource \"aws_s3_bucket\" \"test3\" {\n  bucket = \"iac-remediation-demo-bucket-notags\"\n  tags = {\n    Demo = \"true\"\n  }\n}\n",
				FileSource:     strings.Split(OriginalData1, "\n"),
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{
						Line: 23,
						Col:  1,
					},
					End: model.ResourceLine{
						Line: 28,
						Col:  2,
					},
				},
			},
			searchKey: "aws_s3_bucket[test3].tags",
			file: &model.FileMetadata{
				ScanID:            "Test3",
				ID:                "Test3",
				Kind:              model.KindTerraform,
				OriginalData:      OriginalData1,
				LinesOriginalData: utils.SplitLines(OriginalData1),
			},
		},
	}

	ctx := context.Background()
	for i, testCase := range testCases {
		detector := DetectKindLine{}
		t.Run(fmt.Sprintf("detectTerraformLine-%d", i), func(t *testing.T) {
			v := detector.DetectLine(ctx, testCase.file, testCase.searchKey, 3)
			require.Equal(t, testCase.expected, v)
		})
	}
}

func TestDetectTerraformLineRemediations(t *testing.T) {
	testCases := []struct {
		expected  model.VulnerabilityLines
		searchKey string
		file      *model.FileMetadata
	}{
		{
			expected: model.VulnerabilityLines{
				Line: 24,
				VulnLines: &[]model.CodeLine{
					{Position: 23, Line: "\tversioning {"},
					{Position: 24, Line: "\t  enabled = false"},
					{Position: 25, Line: "\t}"},
				},
				VulnerablilityLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 14, Col: 3},
					End:   model.ResourceLine{Line: 26, Col: 4},
				},
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 24, Col: 4},
					End:   model.ResourceLine{Line: 24, Col: 4},
				},
				LineWithVulnerability: "\t  enabled = false",
				ResolvedFile:          "",
				ResourceSource:        "  resource \"aws_s3_bucket\" \"positive1\" {\n\tbucket = \"my-tf-test-bucket\"\n\tacl    = \"private\"\n\n\ttags = {\n\t  Name        = \"My bucket\"\n\t  Environment = \"Dev\"\n\t}\n\n\tversioning {\n\t  enabled = false\n\t}\n  }\n",
				FileSource:            strings.Split(OriginalDataPositive1, "\n"),
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 14, Col: 3},
					End:   model.ResourceLine{Line: 26, Col: 4},
				},
			},
			searchKey: "aws_s3_bucket[positive1].versioning.enabled",
			file: &model.FileMetadata{
				ScanID:            "console",
				ID:                "positive1.versioning.enabled",
				Kind:              model.KindTerraform,
				OriginalData:      OriginalDataPositive1,
				LinesOriginalData: utils.SplitLines(OriginalDataPositive1),
			},
		},
		{
			expected: model.VulnerabilityLines{
				Line: 14,
				VulnLines: &[]model.CodeLine{
					{Position: 13, Line: ""},
					{Position: 14, Line: "  resource \"aws_s3_bucket\" \"positive2\" {"},
					{Position: 15, Line: "\tbucket = \"my-tf-test-bucket\""},
				},
				VulnerablilityLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 14, Col: 3},
					End:   model.ResourceLine{Line: 22, Col: 4},
				},
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 17, Col: 2},
					End:   model.ResourceLine{Line: 17, Col: 2},
				},
				LineWithVulnerability: "  resource \"aws_s3_bucket\" \"positive2\" {",
				ResolvedFile:          "",
				ResourceSource:        "  resource \"aws_s3_bucket\" \"positive2\" {\n\tbucket = \"my-tf-test-bucket\"\n\tacl    = \"private\"\n\n\ttags = {\n\t  Name        = \"My bucket\"\n\t  Environment = \"Dev\"\n\t}\n  }\n",
				FileSource:            strings.Split(OriginalDataPositive2, "\n"),
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 14, Col: 3},
					End:   model.ResourceLine{Line: 22, Col: 4},
				},
			},
			searchKey: "aws_s3_bucket[positive2]",
			file: &model.FileMetadata{
				ScanID:            "console",
				ID:                "positive2.missing_field",
				Kind:              model.KindTerraform,
				OriginalData:      OriginalDataPositive2,
				LinesOriginalData: utils.SplitLines(OriginalDataPositive2),
			},
		},
		{
			expected: model.VulnerabilityLines{
				Line: 23,
				VulnLines: &[]model.CodeLine{
					{Position: 22, Line: ""},
					{Position: 23, Line: "\tversioning {"},
					{Position: 24, Line: "\t  mfa_delete = true"},
				},
				VulnerablilityLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 14, Col: 3},
					End:   model.ResourceLine{Line: 26, Col: 4},
				},
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 24, Col: 4},
					End:   model.ResourceLine{Line: 24, Col: 4},
				},
				LineWithVulnerability: "\tversioning {",
				ResolvedFile:          "",
				ResourceSource:        "  resource \"aws_s3_bucket\" \"positive3\" {\n\tbucket = \"my-tf-test-bucket\"\n\tacl    = \"private\"\n\n\ttags = {\n\t  Name        = \"My bucket\"\n\t  Environment = \"Dev\"\n\t}\n\n\tversioning {\n\t  mfa_delete = true\n\t}\n  }\n",
				FileSource:            strings.Split(OriginalDataPositive3, "\n"),
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 14, Col: 3},
					End:   model.ResourceLine{Line: 26, Col: 4},
				},
			},
			searchKey: "aws_s3_bucket[positive3].versioning",
			file: &model.FileMetadata{
				ScanID:            "console",
				ID:                "positive3.versioning",
				Kind:              model.KindTerraform,
				OriginalData:      OriginalDataPositive3,
				LinesOriginalData: utils.SplitLines(OriginalDataPositive3),
			},
		},
		{
			expected: model.VulnerabilityLines{
				Line: 27,
				VulnLines: &[]model.CodeLine{
					{Position: 26, Line: "\tversioning_configuration {"},
					{Position: 27, Line: "\t  status = \"Suspended\""},
					{Position: 28, Line: "\t}"},
				},
				VulnerablilityLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 23, Col: 3},
					End:   model.ResourceLine{Line: 29, Col: 4},
				},
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 27, Col: 4},
					End:   model.ResourceLine{Line: 27, Col: 4},
				},
				LineWithVulnerability: "\t  status = \"Suspended\"",
				ResolvedFile:          "",
				ResourceSource:        "  resource \"aws_s3_bucket_versioning\" \"example\" {\n\tbucket = aws_s3_bucket.b0.id\n\n\tversioning_configuration {\n\t  status = \"Suspended\"\n\t}\n  }\n",
				FileSource:            strings.Split(OriginalDataPositive7, "\n"),
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 23, Col: 3},
					End:   model.ResourceLine{Line: 29, Col: 4},
				},
			},
			searchKey: "aws_s3_bucket_versioning[example].versioning_configuration.status",
			file: &model.FileMetadata{
				ScanID:            "console",
				ID:                "positive7.versioning_configuration.status",
				Kind:              model.KindTerraform,
				OriginalData:      OriginalDataPositive7,
				LinesOriginalData: utils.SplitLines(OriginalDataPositive7),
			},
		},
		{
			expected: model.VulnerabilityLines{
				Line: 1,
				VulnLines: &[]model.CodeLine{
					{Position: 1, Line: "resource \"aws_s3_bucket\" \"no_versioning\" {"},
					{Position: 2, Line: "\tbucket = \"my-tf-test-bucket\""},
					{Position: 3, Line: ""},
				},
				VulnerablilityLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 1, Col: 1},
					End:   model.ResourceLine{Line: 7, Col: 4},
				},
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 3, Col: 2},
					End:   model.ResourceLine{Line: 3, Col: 2},
				},
				LineWithVulnerability: "resource \"aws_s3_bucket\" \"no_versioning\" {",
				ResolvedFile:          "",
				ResourceSource:        "resource \"aws_s3_bucket\" \"no_versioning\" {\n\tbucket = \"my-tf-test-bucket\"\n\n\ttags = {\n\t  Name = \"My bucket\"\n\t}\n  }\n",
				FileSource:            strings.Split(OriginalDataMissingVersioning, "\n"),
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 1, Col: 1},
					End:   model.ResourceLine{Line: 7, Col: 4},
				},
			},
			searchKey: "aws_s3_bucket[no_versioning].versioning.enabled",
			file: &model.FileMetadata{
				ScanID:            "console",
				ID:                "missing.versioning.enabled",
				Kind:              model.KindTerraform,
				OriginalData:      OriginalDataMissingVersioning,
				LinesOriginalData: utils.SplitLines(OriginalDataMissingVersioning),
			},
		},
		{
			expected: model.VulnerabilityLines{
				Line: 4,
				VulnLines: &[]model.CodeLine{
					{Position: 3, Line: "    configuration {"},
					{Position: 4, Line: "      status = \"Enabled\""},
					{Position: 5, Line: "    }"},
				},
				VulnerablilityLocation: model.ResourceLocation{
					Start: model.ResourceLine{Col: 1, Line: 1},
					End:   model.ResourceLine{Col: 2, Line: 7},
				},
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{Col: 7, Line: 4},
					End:   model.ResourceLine{Col: 7, Line: 4},
				},
				LineWithVulnerability: "      status = \"Enabled\"",
				ResolvedFile:          "",
				ResourceSource:        "resource \"aws_s3_bucket\" \"example\" {\n  versioning {\n    configuration {\n      status = \"Enabled\"\n    }\n  }\n}\n",
				FileSource:            strings.Split("resource \"aws_s3_bucket\" \"example\" {\n  versioning {\n    configuration {\n      status = \"Enabled\"\n    }\n  }\n}\n", "\n"),
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{Col: 1, Line: 1},
					End:   model.ResourceLine{Col: 2, Line: 7},
				},
			},
			searchKey: "aws_s3_bucket[example].versioning.configuration.status",
			file: &model.FileMetadata{
				ScanID:            "deep",
				ID:                "deep-nested",
				Kind:              model.KindTerraform,
				OriginalData:      "resource \"aws_s3_bucket\" \"example\" {\n  versioning {\n    configuration {\n      status = \"Enabled\"\n    }\n  }\n}\n",
				LinesOriginalData: utils.SplitLines("resource \"aws_s3_bucket\" \"example\" {\n  versioning {\n    configuration {\n      status = \"Enabled\"\n    }\n  }\n}\n"),
			},
		},
		{
			expected: model.VulnerabilityLines{
				Line: 3,
				VulnLines: &[]model.CodeLine{
					{Position: 2, Line: "  statement = [{"},
					{Position: 3, Line: "    actions = [\"s3:GetObject\"]"},
					{Position: 4, Line: "  }]"},
				},
				VulnerablilityLocation: model.ResourceLocation{
					Start: model.ResourceLine{Col: 1, Line: 1},
					End:   model.ResourceLine{Col: 2, Line: 5},
				},
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{Col: 5, Line: 3},
					End:   model.ResourceLine{Col: 5, Line: 3},
				},
				LineWithVulnerability: "    actions = [\"s3:GetObject\"]",
				ResolvedFile:          "",
				ResourceSource:        "resource \"aws_iam_policy_document\" \"example\" {\n  statement = [{\n    actions = [\"s3:GetObject\"]\n  }]\n}\n",
				FileSource:            strings.Split("resource \"aws_iam_policy_document\" \"example\" {\n  statement = [{\n    actions = [\"s3:GetObject\"]\n  }]\n}\n", "\n"),
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{Col: 1, Line: 1},
					End:   model.ResourceLine{Col: 2, Line: 5},
				},
			},
			searchKey: "aws_iam_policy_document[example].statement[0].actions[0]",
			file: &model.FileMetadata{
				ScanID:            "indexing",
				ID:                "policy.actions",
				Kind:              model.KindTerraform,
				OriginalData:      "resource \"aws_iam_policy_document\" \"example\" {\n  statement = [{\n    actions = [\"s3:GetObject\"]\n  }]\n}\n",
				LinesOriginalData: utils.SplitLines("resource \"aws_iam_policy_document\" \"example\" {\n  statement = [{\n    actions = [\"s3:GetObject\"]\n  }]\n}\n"),
			},
		},
		{
			expected: model.VulnerabilityLines{
				Line: 1,
				VulnLines: &[]model.CodeLine{
					{Position: 1, Line: "resource \"aws_instance\" \"example\" {"},
					{Position: 2, Line: "  tags = {"},
					{Position: 3, Line: "    Name = \"web-server\""},
				},
				VulnerablilityLocation: model.ResourceLocation{
					Start: model.ResourceLine{Col: 1, Line: 1},
					End:   model.ResourceLine{Col: 2, Line: 5},
				},
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{Col: 1, Line: 1},
					End:   model.ResourceLine{Col: 1, Line: 1},
				},
				LineWithVulnerability: "resource \"aws_instance\" \"example\" {",
				ResolvedFile:          "",
				ResourceSource:        "resource \"aws_instance\" \"example\" {\n  tags = {\n    Name = \"web-server\"\n  }\n}\n",
				FileSource:            strings.Split("resource \"aws_instance\" \"example\" {\n  tags = {\n    Name = \"web-server\"\n  }\n}\n", "\n"),
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{Col: 1, Line: 1},
					End:   model.ResourceLine{Col: 2, Line: 5},
				},
			},
			searchKey: "aws_instance[example].tags[\"Name\"]",
			file: &model.FileMetadata{
				ScanID:            "mapkey",
				ID:                "tags.name",
				Kind:              model.KindTerraform,
				OriginalData:      "resource \"aws_instance\" \"example\" {\n  tags = {\n    Name = \"web-server\"\n  }\n}\n",
				LinesOriginalData: utils.SplitLines("resource \"aws_instance\" \"example\" {\n  tags = {\n    Name = \"web-server\"\n  }\n}\n"),
			},
		},
	}

	ctx := context.Background()
	for i, testCase := range testCases {
		detector := DetectKindLine{}
		t.Run(fmt.Sprintf("detectTerraformLine-%d", i), func(t *testing.T) {
			v := detector.DetectLine(ctx, testCase.file, testCase.searchKey, 3)
			require.Equal(t, testCase.expected, v)
		})
	}
}

const OriginalDataPositive1 = `provider "aws" {
	region = "us-east-1"
  }

  terraform {
	required_providers {
	  aws = {
		source  = "hashicorp/aws"
		version = "~> 3.0"
	  }
	}
  }

  resource "aws_s3_bucket" "positive1" {
	bucket = "my-tf-test-bucket"
	acl    = "private"

	tags = {
	  Name        = "My bucket"
	  Environment = "Dev"
	}

	versioning {
	  enabled = false
	}
  }`

const OriginalDataPositive2 = `provider "aws" {
	region = "us-east-1"
  }

  terraform {
	required_providers {
	  aws = {
		source  = "hashicorp/aws"
		version = "~> 3.0"
	  }
	}
  }

  resource "aws_s3_bucket" "positive2" {
	bucket = "my-tf-test-bucket"
	acl    = "private"

	tags = {
	  Name        = "My bucket"
	  Environment = "Dev"
	}
  }`

const OriginalDataPositive3 = `provider "aws" {
	region = "us-east-1"
  }

  terraform {
	required_providers {
	  aws = {
		source  = "hashicorp/aws"
		version = "~> 3.0"
	  }
	}
  }

  resource "aws_s3_bucket" "positive3" {
	bucket = "my-tf-test-bucket"
	acl    = "private"

	tags = {
	  Name        = "My bucket"
	  Environment = "Dev"
	}

	versioning {
	  mfa_delete = true
	}
  }`

const OriginalDataPositive7 = `terraform {
	required_providers {
	  aws = {
		source = "hashicorp/aws"
		version = "4.2.0"
	  }
	}
  }

  provider "aws" {
	# Configuration options
  }

  resource "aws_s3_bucket" "b0" {
	bucket = "my-tf-test-bucket"

	tags = {
	  Name        = "My bucket"
	  Environment = "Dev"
	}
  }

  resource "aws_s3_bucket_versioning" "example" {
	bucket = aws_s3_bucket.b0.id

	versioning_configuration {
	  status = "Suspended"
	}
  }`

const OriginalDataMissingVersioning = `resource "aws_s3_bucket" "no_versioning" {
	bucket = "my-tf-test-bucket"

	tags = {
	  Name = "My bucket"
	}
  }
  `

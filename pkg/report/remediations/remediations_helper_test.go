package model

import (
	"fmt"
	"testing"

	"github.com/Checkmarx/kics/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDetectTerraformLine(t *testing.T) { //nolint
	testCases := []struct {
		expected      model.SarifFix
		vuln          model.VulnerableFile
		startLocation model.SarifResourceLocation
		endLocation   model.SarifResourceLocation
	}{
		{
			vuln: model.VulnerableFile{
				FileName:        "assets/queries/terraform/aws/s3_bucket_without_versioning/test/positive6.tf",
				SimilarityID:    "3f57a6f32db9ebf0d1da3a73122381ae2b7605a6ba3dc432219fe9fd5a44d7d5",
				OldSimilarityID: "",
				Line:            1,
				ResourceLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 1, Col: 1},
					End:   model.ResourceLine{Line: 9, Col: 2},
				},
				VulnLines:        nil,
				ResourceType:     "n/a",
				ResourceName:     "n/a",
				IssueType:        "MissingAttribute",
				SearchKey:        "module[s3_bucket]",
				SearchLine:       1,
				SearchValue:      "",
				KeyExpectedValue: "'block_public_acls' should equal 'true'",
				KeyActualValue:   "'block_public_acls' is missing",
				Value:            nil,
				Remediation:      "block_public_acls = true",
				RemediationType:  "addition",
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 8, Col: 2},
					End:   model.ResourceLine{Line: 8, Col: 2},
				},
				LineWithVulnerability: "module \"s3_bucket\" {",
				ResourceSource:        "module \"s3_bucket\" {\n  source = \"terraform-aws-modules/s3-bucket/aws\"\n\n  version = \"3.7.0\"\n\n  bucket = \"my-s3-bucket\"\n  acl    = \"private\"\n\n}\n",
				FileSource:            []string{"module \"s3_bucket\" {", "  source = \"terraform-aws-modules/s3-bucket/aws\"", "", "  version = \"3.7.0\"", "", "  bucket = \"my-s3-bucket\"", "  acl    = \"private\"", "", "}", ""},
				BlockLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: 1, Col: 1},
					End:   model.ResourceLine{Line: 9, Col: 2},
				},
			},
			startLocation: model.SarifResourceLocation{Line: 8, Col: 2},
			endLocation:   model.SarifResourceLocation{Line: 8, Col: 2},
			expected: model.SarifFix{
				ArtifactChanges: []model.ArtifactChange{
					{
						ArtifactLocation: model.ArtifactLocation{
							URI: "assets/queries/terraform/aws/s3_bucket_without_versioning/test/positive6.tf",
						},
						Replacements: []model.FixReplacement{
							{
								DeletedRegion: model.SarifRegion{
									StartLine:   8,
									EndLine:     8,
									StartColumn: 2,
									EndColumn:   2,
								},
								InsertedContent: model.FixContent{
									Text: "\n  block_public_acls = true\n",
								},
							},
						},
					},
				},
				Description: model.FixMessage{
					Text: "Apply remediation: block_public_acls = true",
				},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TransformToSarifFix-%d", i), func(t *testing.T) {
			v, err := TransformToSarifFix(testCase.vuln, testCase.startLocation, testCase.endLocation)
			assert.Nil(t, err)
			require.Equal(t, testCase.expected, v)
		})
	}
}

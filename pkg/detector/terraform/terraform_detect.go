/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package terraform

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/Checkmarx/kics/pkg/detector"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/rs/zerolog"
)

// DetectKindLine defines a kindDetectLine type
type DetectKindLine struct {
}

const (
	undetectedVulnerabilityLine = -1
)

// DetectLine searches vulnerability line in terraform files
func (d DetectKindLine) DetectLine(file *model.FileMetadata, searchKey string,
	outputLines int, logwithfields *zerolog.Logger) model.VulnerabilityLines {
	det := &detector.DefaultDetectLineResponse{
		CurrentLine:     0,
		IsBreak:         false,
		FoundAtLeastOne: false,
		ResolvedFile:    file.FilePath,
		ResolvedFiles:   make(map[string]model.ResolvedFileSplit),
	}

	var extractedString [][]string
	extractedString = detector.GetBracketValues(searchKey, extractedString, "")
	sKey := searchKey
	for idx, str := range extractedString {
		sKey = strings.Replace(sKey, str[0], `{{`+strconv.Itoa(idx)+`}}`, -1)
	}

	lines := *file.LinesOriginalData
	splitSanitized := strings.Split(sKey, ".")
	for index, split := range splitSanitized {
		if strings.Contains(split, "$ref") {
			splitSanitized[index] = strings.Join(splitSanitized[index:], ".")
			splitSanitized = splitSanitized[:index+1]
			break
		}
	}

	for _, key := range splitSanitized {
		substr1, substr2 := detector.GenerateSubstrings(key, extractedString)
		det, _ = det.DetectCurrentLine(substr1, substr2, 0, lines)

		if det.IsBreak {
			break
		}
	}

	if det.FoundAtLeastOne {
		line := det.CurrentLine + 1

		resourceStart, resourceEnd, lineContent, err := parseAndFindTerraformBlock([]byte(file.OriginalData), line)
		if err != nil {
			fmt.Printf("Failed to parse and find Terraform block for line %d in file %s: %s\n", line, file.FilePath, err)
			return model.VulnerabilityLines{
				Line:         undetectedVulnerabilityLine,
				VulnLines:    &[]model.CodeLine{},
				ResolvedFile: file.FilePath,
				ResourceLocation: model.ResourceLocation{
					ResourceStart: resourceStart,
					ResourceEnd:   resourceEnd,
				},
			}
		}

		return model.VulnerabilityLines{
			Line:         line,
			VulnLines:    detector.GetAdjacentVulnLines(det.CurrentLine, outputLines, lines),
			ResolvedFile: file.FilePath,
			ResourceLocation: model.ResourceLocation{
				ResourceStart: resourceStart,
				ResourceEnd:   resourceEnd,
			},
			LineWithVulnerability: lineContent,
		}
	}

	logwithfields.Warn().Msgf("Failed to detect Terraform line, query response %s", sKey)

	return model.VulnerabilityLines{
		Line:         undetectedVulnerabilityLine,
		VulnLines:    &[]model.CodeLine{},
		ResolvedFile: file.FilePath,
	}
}

func parseAndFindTerraformBlock(src []byte, identifyingLine int) (model.ResourceLine, model.ResourceLine, string, error) {
	filePath := "temp.tf"
	lineContent := ""

	resourceStart := model.ResourceLine{Line: -1, Col: -1}
	resourceEnd := model.ResourceLine{Line: -1, Col: -1}

	hclSyntaxFile, diagnostics := hclsyntax.ParseConfig(src, filePath, hcl.InitialPos)
	if diagnostics != nil && diagnostics.HasErrors() {
		return resourceStart, resourceEnd, lineContent, fmt.Errorf("failed to parse hcl file %s because of errors %s", filePath, diagnostics.Errs())
	}

	lines := bytes.Split(src, []byte("\n"))

	var findDeepestBlock func(blocks hclsyntax.Blocks) *hclsyntax.Block
	findDeepestBlock = func(blocks hclsyntax.Blocks) *hclsyntax.Block {
		for _, block := range blocks {
			blockStart := block.TypeRange.Start
			blockEnd := block.Body.EndRange.End

			if blockStart.Line <= identifyingLine && identifyingLine <= blockEnd.Line {
				// Check nested blocks first
				nested := findDeepestBlock(block.Body.Blocks)
				if nested != nil {
					return nested
				}
				return block // No deeper match, return this one
			}
		}
		return nil
	}

	rootBlocks := hclSyntaxFile.Body.(*hclsyntax.Body).Blocks
	targetBlock := findDeepestBlock(rootBlocks)

	if targetBlock == nil {
		return resourceStart, resourceEnd, lineContent, fmt.Errorf("failed to find block for line %d in file %s", identifyingLine, filePath)
	}

	blockStart := targetBlock.TypeRange.Start
	blockEnd := targetBlock.Body.EndRange.End

	resourceStart = model.ResourceLine{Line: blockStart.Line, Col: blockStart.Column}
	resourceEnd = model.ResourceLine{Line: blockEnd.Line, Col: blockEnd.Column}

	// Defensive: make sure we have the actual line
	lineIndex := identifyingLine - 1
	if lineIndex >= 0 && lineIndex < len(lines) {
		lineContent = string(lines[lineIndex])
	}

	return resourceStart, resourceEnd, lineContent, nil
}

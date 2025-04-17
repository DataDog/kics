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

type BlockInfo struct {
	Block   *hclsyntax.Block
	Depth   int
	Parent  *hclsyntax.Block
	IsMatch bool
}

func parseAndFindTerraformBlock(src []byte, identifyingLine int) (model.ResourceLine, model.ResourceLine, string, error) {
	filePath := "temp.tf"
	lines := bytes.Split(src, []byte("\n"))
	if identifyingLine <= 0 || identifyingLine > len(lines) {
		return model.ResourceLine{}, model.ResourceLine{}, "", fmt.Errorf("line %d is out of range", identifyingLine)
	}
	lineContent := string(lines[identifyingLine-1])

	resourceStart := model.ResourceLine{Line: -1, Col: -1}
	resourceEnd := model.ResourceLine{Line: -1, Col: -1}

	hclSyntaxFile, diagnostics := hclsyntax.ParseConfig(src, filePath, hcl.InitialPos)
	if diagnostics != nil && diagnostics.HasErrors() {
		return resourceStart, resourceEnd, lineContent, fmt.Errorf("failed to parse hcl file %s: %v", filePath, diagnostics.Errs())
	}

	for _, block := range hclSyntaxFile.Body.(*hclsyntax.Body).Blocks {
		start := block.TypeRange.Start
		end := block.Body.SrcRange.End

		if identifyingLine >= start.Line && identifyingLine <= end.Line {
			// Check if inside nested block
			nestedBlock, _, nestedEnd := findInnermostBlock(block, identifyingLine)
			if nestedBlock != nil {
				// Identifying line is in a nested block → insert before nested block's closing brace
				resourceStart = model.ResourceLine{
					Line: nestedEnd.Line,
					Col:  1,
				}
				resourceEnd = model.ResourceLine{
					Line: nestedEnd.Line,
					Col:  nestedEnd.Column,
				}
			} else {
				// Top-level block → insert before outer block's closing brace
				resourceStart = model.ResourceLine{
					Line: end.Line,
					Col:  1,
				}
				resourceEnd = model.ResourceLine{
					Line: end.Line,
					Col:  end.Column,
				}
			}
			break
		}
	}

	if resourceStart.Line == -1 || resourceEnd.Line == -1 {
		return resourceStart, resourceEnd, lineContent, fmt.Errorf("failed to locate block for line %d", identifyingLine)
	}

	return resourceStart, resourceEnd, lineContent, nil
}

func findInnermostBlock(block *hclsyntax.Block, line int) (*hclsyntax.Block, hcl.Pos, hcl.Pos) {
	for _, nested := range block.Body.Blocks {
		start := nested.TypeRange.Start
		end := nested.Body.SrcRange.End
		if line >= start.Line && line <= end.Line {
			if deeper, s, e := findInnermostBlock(nested, line); deeper != nil {
				return deeper, s, e
			}
			return nested, start, end
		}
	}
	return nil, hcl.Pos{}, hcl.Pos{}
}

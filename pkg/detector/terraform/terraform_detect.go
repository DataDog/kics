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
		blockStart := block.TypeRange.Start
		blockEnd := block.Body.SrcRange.End

		if identifyingLine >= blockStart.Line && identifyingLine <= blockEnd.Line {
			// Check for nested block or inline object
			structureName, nestedStart, nestedEnd, isAttribute := findContainingStructure(block, identifyingLine)

			var insertionLine int
			var insertionCol int

			if structureName != "" {
				// Line is inside a nested block or object attribute
				insertionLine = nestedEnd.Line - 1 // one line before the closing brace

				if isAttribute {
					// For object-style attributes, infer indent from inner attribute lines
					for i := insertionLine - 1; i > nestedStart.Line && i <= len(lines); i-- {
						trimmed := strings.TrimSpace(string(lines[i-1]))
						if trimmed != "" && !strings.HasPrefix(trimmed, "#") && strings.Contains(trimmed, "=") {
							insertionCol = countLeadingSpacesOrTabs(lines[i-1]) + 1
							break
						}
					}
				} else {
					// For real nested blocks, use closing brace indent
					if insertionLine-1 >= 0 && insertionLine-1 < len(lines) {
						insertionCol = countLeadingSpacesOrTabs(lines[insertionLine-1]) + 1
					}
				}
			} else {
				// Top-level block — insert before outer block's closing brace
				insertionLine = blockEnd.Line - 1 // one line before the closing brace

				// Try to infer indentation from the last meaningful line inside the block
				for i := insertionLine; i > blockStart.Line && i <= len(lines); i-- {
					trimmed := strings.TrimSpace(string(lines[i-1]))
					if trimmed != "" && !strings.HasPrefix(trimmed, "#") {
						insertionCol = countLeadingSpacesOrTabs(lines[i-1]) + 1
						break
					}
				}
			}

			// Fallback if nothing was found
			if insertionCol == 0 && insertionLine-1 < len(lines) {
				insertionCol = countLeadingSpacesOrTabs(lines[insertionLine-1]) + 1
			}

			resourceStart = model.ResourceLine{Line: insertionLine, Col: insertionCol}
			resourceEnd = model.ResourceLine{Line: insertionLine, Col: insertionCol}

			return resourceStart, resourceEnd, lineContent, nil
		}
	}

	return resourceStart, resourceEnd, lineContent, fmt.Errorf("failed to locate block for line %d", identifyingLine)
}

func countLeadingSpacesOrTabs(line []byte) int {
	count := 0
	for _, b := range line {
		if b == ' ' || b == '\t' {
			count++
		} else {
			break
		}
	}
	return count
}

// findContainingStructure returns the name, start, and end of a nested block or attribute containing the line.
// If the line is not part of a nested block or object-style attribute, it returns empty string and zero positions.
func findContainingStructure(block *hclsyntax.Block, line int) (string, hcl.Pos, hcl.Pos, bool) {
	for _, nested := range block.Body.Blocks {
		start := nested.TypeRange.Start
		end := nested.Body.SrcRange.End
		if line >= start.Line && line <= end.Line {
			if deeperType, deeperStart, deeperEnd, isAttr := findContainingStructure(nested, line); deeperType != "" {
				return deeperType, deeperStart, deeperEnd, isAttr
			}
			return nested.Type, start, end, false
		}
	}

	for name, attr := range block.Body.Attributes {
		if _, ok := attr.Expr.(*hclsyntax.ObjectConsExpr); ok {
			start := attr.SrcRange.Start
			end := attr.SrcRange.End
			if line >= start.Line && line <= end.Line {
				return name, start, end, true
			}
		}
	}

	return "", hcl.Pos{}, hcl.Pos{}, false
}

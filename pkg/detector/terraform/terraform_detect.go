/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package terraform

import (
	"bytes"
	"fmt"
	"regexp"
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
	// Sanitize malformed Go formatting artifacts
	searchKey = sanitizeSearchKey(searchKey)

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
		// Only replace raw bracketed values (e.g., [abc]), not placeholders (e.g., [{{var}}])
		if !strings.Contains(str[0], "{{") {
			sKey = strings.Replace(sKey, str[0], `{{`+strconv.Itoa(idx)+`}}`, -1)
		}
	}

	lines := *file.LinesOriginalData
	splitSanitized := strings.FieldsFunc(sKey, func(r rune) bool {
		return r == '.' || r == '/'
	})
	for index, split := range splitSanitized {
		if strings.Contains(split, "$ref") {
			splitSanitized[index] = strings.Join(splitSanitized[index:], ".")
			splitSanitized = splitSanitized[:index+1]
			break
		}
	}

	for _, key := range splitSanitized {
		substr1, substr2 := detector.GenerateSubstrings(key, extractedString, lines, det.CurrentLine)
		det, _ = det.DetectCurrentLine(substr1, substr2, 0, lines)

		if det.IsBreak {
			break
		}
	}

	if det.FoundAtLeastOne {
		line := det.CurrentLine + 1

		resourceStart, resourceEnd, remediationStart, remediationEnd, lineContent, resourceSource, blockStart, blockEnd, err := parseAndFindTerraformBlock([]byte(file.OriginalData), line)
		if err != nil {
			fmt.Printf("Failed to parse and find Terraform block for line %d in file %s: %s\n", line, file.FilePath, err)
			return model.VulnerabilityLines{
				Line:         undetectedVulnerabilityLine,
				VulnLines:    &[]model.CodeLine{},
				ResolvedFile: file.FilePath,
				VulnerablilityLocation: model.ResourceLocation{
					Start: resourceStart,
					End:   resourceEnd,
				},
				RemediationLocation: model.ResourceLocation{
					Start: remediationStart,
					End:   remediationEnd,
				},
				ResourceSource: resourceSource,
				FileSource:     *file.LinesOriginalData,
				BlockLocation: model.ResourceLocation{
					Start: blockStart,
					End:   blockEnd,
				},
			}
		}

		return model.VulnerabilityLines{
			Line:         line,
			VulnLines:    detector.GetAdjacentVulnLines(det.CurrentLine, outputLines, lines),
			ResolvedFile: file.FilePath,
			VulnerablilityLocation: model.ResourceLocation{
				Start: resourceStart,
				End:   resourceEnd,
			},
			RemediationLocation: model.ResourceLocation{
				Start: remediationStart,
				End:   remediationEnd,
			},
			LineWithVulnerability: lineContent,
			ResourceSource:        resourceSource,
			FileSource:            *file.LinesOriginalData,
			BlockLocation: model.ResourceLocation{
				Start: blockStart,
				End:   blockEnd,
			},
		}
	}

	logwithfields.Warn().Msgf("Failed to detect Terraform line, query response %s", sKey)

	return model.VulnerabilityLines{
		Line:           undetectedVulnerabilityLine,
		VulnLines:      &[]model.CodeLine{},
		ResolvedFile:   file.FilePath,
		ResourceSource: "",
		FileSource:     *file.LinesOriginalData,
	}
}

type BlockInfo struct {
	Block   *hclsyntax.Block
	Depth   int
	Parent  *hclsyntax.Block
	IsMatch bool
}

func sanitizeSearchKey(key string) string {
	// Replace any instance of [%!s(int=N)] with [N]
	re := regexp.MustCompile(`\[%!s\(int=(\d+)\)\]`)
	return re.ReplaceAllString(key, "[$1]")
}

func parseAndFindTerraformBlock(src []byte, identifyingLine int) (model.ResourceLine, model.ResourceLine, model.ResourceLine, model.ResourceLine, string, string, model.ResourceLine, model.ResourceLine, error) {
	filePath := "temp.tf"
	lines := bytes.Split(src, []byte("\n"))

	if identifyingLine <= 0 || identifyingLine > len(lines) {
		return model.ResourceLine{}, model.ResourceLine{}, model.ResourceLine{}, model.ResourceLine{}, "", "", model.ResourceLine{}, model.ResourceLine{}, fmt.Errorf("line %d is out of range", identifyingLine)
	}

	lineContent := string(lines[identifyingLine-1])
	var vulnerabilitySource string

	vulnerabilityStart := model.ResourceLine{Line: -1, Col: -1}
	vulnerabilityEnd := model.ResourceLine{Line: -1, Col: -1}
	remediationStart := model.ResourceLine{Line: -1, Col: -1}
	remediationEnd := model.ResourceLine{Line: -1, Col: -1}
	blockLocationStart := model.ResourceLine{Line: -1, Col: -1}
	blockLocationEnd := model.ResourceLine{Line: -1, Col: -1}

	hclSyntaxFile, diagnostics := hclsyntax.ParseConfig(src, filePath, hcl.InitialPos)
	if diagnostics != nil && diagnostics.HasErrors() {
		return vulnerabilityStart, vulnerabilityEnd, remediationStart, remediationEnd, lineContent, "", blockLocationStart, blockLocationEnd, fmt.Errorf("failed to parse HCL file %s: %v", filePath, diagnostics.Errs())
	}

	blocks := hclSyntaxFile.Body.(*hclsyntax.Body).Blocks

	for _, block := range blocks {
		blockStart := block.TypeRange.Start
		blockEnd := block.Body.SrcRange.End

		blockLines := lines[blockStart.Line-1 : blockEnd.Line]
		var sb strings.Builder
		for _, l := range blockLines {
			sb.Write(l)
			sb.WriteByte('\n')
		}
		vulnerabilitySource = sb.String()

		if identifyingLine >= blockStart.Line && identifyingLine <= blockEnd.Line {
			var insertionLine int
			var insertionCol int
			var caseType string

			structureName, nestedStart, nestedEnd, _ := findContainingStructure(block, identifyingLine)

			if structureName != "" {
				if identifyingLine == nestedEnd.Line {
					insertionLine = nestedEnd.Line - 1
					caseType = "nested-end"
				} else if identifyingLine == nestedStart.Line {
					insertionLine = nestedStart.Line + 1
					caseType = "nested-start"
				} else {
					insertionLine = identifyingLine
					caseType = "nested-body"
				}
			} else {
				if identifyingLine == blockStart.Line {
					insertionLine = blockEnd.Line - 1
					for i := insertionLine; i >= blockStart.Line; i-- {
						_, nestedStart, nestedEnd, isAttr := findContainingStructure(block, i)
						if isAttr && nestedEnd.Line >= insertionLine {
							insertionLine = nestedStart.Line - 1
							continue
						}
						break
					}
					caseType = "block-start"
				} else {
					insertionLine = identifyingLine
					caseType = "block-body"
				}
			}

			insertionCol = determineInsertionIndent(
				toStringLines(lines),
				insertionLine,
				caseType,
				nestedStart.Line,
				nestedEnd.Line,
				blockStart.Line,
				blockEnd.Line,
			) + 1

			// if this is block start and the insertion line contains } we want to insert at the end and not at the start
			if caseType == "block-start" && strings.TrimSpace(string(lines[insertionLine-1])) == "}" {
				insertionCol = len(lines[insertionLine-1]) + 1
			}

			remediationStart = model.ResourceLine{Line: insertionLine, Col: insertionCol}
			remediationEnd = model.ResourceLine{Line: insertionLine, Col: insertionCol}
			vulnerabilityStart = model.ResourceLine{Line: blockStart.Line, Col: blockStart.Column}
			vulnerabilityEnd = model.ResourceLine{Line: blockEnd.Line, Col: blockEnd.Column}
			blockLocationStart = model.ResourceLine{Line: blockStart.Line, Col: blockStart.Column}
			blockLocationEnd = model.ResourceLine{Line: blockEnd.Line, Col: blockEnd.Column}

			return vulnerabilityStart, vulnerabilityEnd, remediationStart, remediationEnd, lineContent, vulnerabilitySource, blockLocationStart, blockLocationEnd, nil
		}
	}

	return model.ResourceLine{}, model.ResourceLine{}, model.ResourceLine{}, model.ResourceLine{}, "", "", model.ResourceLine{}, model.ResourceLine{}, fmt.Errorf("failed to locate block for line %d", identifyingLine)
}

func toStringLines(byteLines [][]byte) []string {
	result := make([]string, len(byteLines))
	for i, b := range byteLines {
		result[i] = string(b)
	}
	return result
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

// determineInsertionIndent determines correct indentation based on insertion case.
func determineInsertionIndent(lines []string, insertionLine int, caseType string, nestedStart int, nestedEnd int, blockStart int, blockEnd int) int {
	switch caseType {
	case "nested-end":
		// Look upwards inside nested block
		for i := nestedEnd - 2; i >= nestedStart-1; i-- {
			trimmed := strings.TrimSpace(lines[i])
			if trimmed != "" && !strings.HasPrefix(trimmed, "#") {
				return countLeadingSpacesOrTabs([]byte(lines[i]))
			}
		}
	case "nested-start":
		// 2 spaces deeper than nested block header
		return countLeadingSpacesOrTabs([]byte(lines[nestedStart-1])) + 2
	case "nested-body":
		// Match current line
		return countLeadingSpacesOrTabs([]byte(lines[insertionLine-1]))
	case "block-start":
		// if line being inserted at is not } then we want to insert at the start of the content
		if strings.TrimSpace(lines[insertionLine-1]) != "}" {
			nonWhitespaceIndex := firstNonWhitespaceIndex(lines[insertionLine-1])
			if nonWhitespaceIndex != -1 {
				return nonWhitespaceIndex
			}
		}

		// 2 spaces deeper than block header
		return 1
	case "block-body":
		// Match current line
		return countLeadingSpacesOrTabs([]byte(lines[insertionLine-1]))
	}
	return 0
}

func firstNonWhitespaceIndex(line string) int {
	for i, r := range line {
		if r != ' ' && r != '\t' {
			return i
		}
	}
	return -1
}

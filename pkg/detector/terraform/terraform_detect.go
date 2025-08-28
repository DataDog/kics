/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package terraform

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Checkmarx/kics/pkg/detector"
	"github.com/Checkmarx/kics/pkg/logger"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type DetectKindLine struct{}

const undetectedVulnerabilityLine = -1

// DetectLine searches vulnerability line in terraform files
func (d DetectKindLine) DetectLine(ctx context.Context, file *model.FileMetadata, searchKey string, outputLines int) model.VulnerabilityLines {
	logger := logger.FromContext(ctx)
	searchKey = sanitizeSearchKey(searchKey)

	detection := &detector.DefaultDetectLineResponse{
		CurrentLine:     0,
		IsBreak:         false,
		FoundAtLeastOne: false,
		ResolvedFile:    file.FilePath,
		ResolvedFiles:   make(map[string]model.ResolvedFileSplit),
	}

	extracted := detector.GetBracketValues(searchKey, [][]string{}, "")
	normalizedKey := searchKey
	for i, match := range extracted {
		if !strings.Contains(match[0], "{{") {
			normalizedKey = strings.Replace(normalizedKey, match[0], `{{`+strconv.Itoa(i)+`}}`, -1)
		}
	}

	keyParts := strings.FieldsFunc(normalizedKey, func(r rune) bool {
		return r == '.' || r == '/'
	})
	for i, part := range keyParts {
		if strings.Contains(part, "$ref") {
			keyParts = append(keyParts[:i+1], keyParts[i+1:]...)
			break
		}
	}

	lines := *file.LinesOriginalData

	for _, part := range keyParts {
		s1, s2 := GenerateSubstrings(ctx, part, extracted, lines, detection.CurrentLine)
		detection, _, _, _ = detection.DetectCurrentLine(s1, s2, 0, lines)
		if detection.IsBreak {
			break
		}
	}

	if detection.FoundAtLeastOne {
		line := detection.CurrentLine + 1
		vulnLines, err := locateTerraformBlock(ctx, []byte(file.OriginalData), line, lines)
		if err != nil {
			logger.Error().Err(err).Msgf("Failed to parse block at line %d in file %s", line, file.FilePath)
			return buildEmptyVulnerabilityLines(file)
		}
		vulnLines.Line = line
		vulnLines.VulnLines = detector.GetAdjacentVulnLines(detection.CurrentLine, outputLines, lines)
		vulnLines.ResolvedFile = file.FilePath
		vulnLines.FileSource = lines
		return vulnLines
	}

	logger.Warn().Msgf("Failed to detect Terraform line, query response %s", normalizedKey)
	return buildEmptyVulnerabilityLines(file)
}

func buildEmptyVulnerabilityLines(file *model.FileMetadata) model.VulnerabilityLines {
	return model.VulnerabilityLines{
		Line:           undetectedVulnerabilityLine,
		VulnLines:      &[]model.CodeLine{},
		ResolvedFile:   file.FilePath,
		ResourceSource: "",
		FileSource:     *file.LinesOriginalData,
	}
}

func sanitizeSearchKey(key string) string {
	re := regexp.MustCompile(`\[%!s\(int=(\d+)\)\]`)
	return re.ReplaceAllString(key, "[$1]")
}

func locateTerraformBlock(ctx context.Context, src []byte, identifyingLine int, strLines []string) (model.VulnerabilityLines, error) {
	logger := logger.FromContext(ctx)
	filePath := "temp.tf"

	if identifyingLine <= 0 || identifyingLine > len(strLines) {
		err := fmt.Errorf("line %d is out of range", identifyingLine)
		logger.Error().Msg(err.Error())
		return model.VulnerabilityLines{}, err
	}

	hclFile, diagnostics := hclsyntax.ParseConfig(src, filePath, hcl.InitialPos)
	if diagnostics.HasErrors() {
		err := fmt.Errorf("failed to parse HCL: %v", diagnostics.Errs())
		logger.Error().Msg(err.Error())
		return model.VulnerabilityLines{}, err
	}

	body, ok := hclFile.Body.(*hclsyntax.Body)
	if !ok {
		err := fmt.Errorf("unexpected HCL body type")
		logger.Error().Msg(err.Error())
		return model.VulnerabilityLines{}, err
	}

	for _, block := range body.Blocks {
		start := block.TypeRange.Start
		end := block.Body.SrcRange.End
		if identifyingLine >= start.Line && identifyingLine <= end.Line {
			blockSrc := extractBlockSource(strLines, start.Line, end.Line)

			insertionLine, insertionCol := calculateInsertionPoint(block, identifyingLine, strLines)
			return model.VulnerabilityLines{
				VulnerablilityLocation: model.ResourceLocation{
					Start: toResourceLine(start),
					End:   toResourceLine(end),
				},
				RemediationLocation: model.ResourceLocation{
					Start: model.ResourceLine{Line: insertionLine, Col: insertionCol},
					End:   model.ResourceLine{Line: insertionLine, Col: insertionCol},
				},
				BlockLocation: model.ResourceLocation{
					Start: toResourceLine(start),
					End:   toResourceLine(end),
				},
				LineWithVulnerability: string(strLines[identifyingLine-1]),
				ResourceSource:        blockSrc,
			}, nil
		}
	}

	err := fmt.Errorf("failed to locate block for line %d", identifyingLine)
	logger.Error().Msg(err.Error())
	return model.VulnerabilityLines{}, err
}

func toResourceLine(pos hcl.Pos) model.ResourceLine {
	return model.ResourceLine{Line: pos.Line, Col: pos.Column}
}

func extractBlockSource(lines []string, start, end int) string {
	return string(strings.Join(lines[start-1:end], "\n")) + "\n"
}

func calculateInsertionPoint(block *hclsyntax.Block, line int, lines []string) (int, int) {
	name, nestedStart, nestedEnd, isAttr := findContainingStructure(block, line)

	var insertionLine int
	var caseType string

	if name != "" {
		// Detect if this is a function call like merge(...) and avoid inserting inside it
		lineText := strings.TrimSpace(lines[nestedStart.Line-1])
		if isAttr && strings.Contains(lineText, "(") && strings.Contains(lineText, "{") {
			// Likely a function call wrapping a block (e.g., merge(..., { ... }))
			// We do not want to insert inside the nested structure
			insertionLine = nestedEnd.Line + 1
			caseType = "block-body"
		} else {
			switch {
			case line == nestedEnd.Line:
				insertionLine = nestedEnd.Line - 1
				caseType = "nested-end"
			case line == nestedStart.Line:
				insertionLine = nestedStart.Line + 1
				caseType = "nested-start"
			default:
				insertionLine = line
				caseType = "nested-body"
			}
		}
	} else {
		if line == block.TypeRange.Start.Line {
			insertionLine = block.Body.SrcRange.End.Line - 1
			for i := insertionLine; i >= block.TypeRange.Start.Line; i-- {
				_, s, e, attr := findContainingStructure(block, i)
				if attr && e.Line >= insertionLine {
					insertionLine = s.Line - 1
				} else {
					break
				}
			}
			caseType = "block-start"
		} else {
			insertionLine = line
			caseType = "block-body"
		}
	}

	col := determineInsertionIndent(lines, insertionLine, caseType, nestedStart.Line, nestedEnd.Line, block.TypeRange.Start.Line, block.Body.SrcRange.End.Line) + 1
	trimmed := strings.TrimSpace(lines[insertionLine-1])
	if caseType == "block-start" && (strings.Contains(trimmed, "}") || isHeredocTerminator(trimmed, lines, insertionLine-1)) {
		col = len(lines[insertionLine-1]) + 1
	}
	return insertionLine, col
}

func findContainingStructure(block *hclsyntax.Block, line int) (string, hcl.Pos, hcl.Pos, bool) {
	for _, nested := range block.Body.Blocks {
		if line >= nested.TypeRange.Start.Line && line <= nested.Body.SrcRange.End.Line {
			if name, s, e, isAttr := findContainingStructure(nested, line); name != "" {
				return name, s, e, isAttr
			}
			return nested.Type, nested.TypeRange.Start, nested.Body.SrcRange.End, false
		}
	}

	for name, attr := range block.Body.Attributes {
		start := attr.SrcRange.Start
		end := attr.SrcRange.End
		if line >= start.Line && line <= end.Line {
			switch attr.Expr.(type) {
			case *hclsyntax.ObjectConsExpr:
				return name, start, end, true
			case *hclsyntax.FunctionCallExpr:
				return name, start, end, true // allow insertion logic to handle function call detection
			}
		}
	}

	return "", hcl.Pos{}, hcl.Pos{}, false
}

func determineInsertionIndent(lines []string, insertionLine int, caseType string, nestedStart, nestedEnd, blockStart, blockEnd int) int {
	switch caseType {
	case "nested-end":
		for i := nestedEnd - 2; i >= nestedStart-1; i-- {
			if trimmed := strings.TrimSpace(lines[i]); trimmed != "" && !strings.HasPrefix(trimmed, "#") {
				return countLeadingSpacesOrTabs([]byte(lines[i]))
			}
		}
	case "nested-start":
		return countLeadingSpacesOrTabs([]byte(lines[nestedStart-1])) + 2
	case "nested-body", "block-body":
		return countLeadingSpacesOrTabs([]byte(lines[insertionLine-1]))
	case "block-start":
		if strings.TrimSpace(lines[insertionLine-1]) != "}" {
			if idx := firstNonWhitespaceIndex(lines[insertionLine-1]); idx != -1 {
				return idx
			}
		}
		return 1
	}
	return 0
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

func firstNonWhitespaceIndex(line string) int {
	for i, r := range line {
		if r != ' ' && r != '\t' {
			return i
		}
	}
	return -1
}

func isHeredocTerminator(line string, lines []string, idx int) bool {
	for i := idx - 1; i >= 0; i-- {
		text := strings.TrimSpace(lines[i])
		if strings.Contains(text, "<<") {
			parts := strings.Split(text, "<<")
			if len(parts) == 2 {
				marker := strings.TrimSpace(parts[1])
				return line == marker
			}
		}
	}
	return false
}

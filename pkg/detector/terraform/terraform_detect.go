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
	"github.com/hashicorp/hcl/v2/hclwrite"
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
	lineContent := ""
	resourceStart := model.ResourceLine{Line: -1, Col: -1}
	resourceEnd := model.ResourceLine{Line: -1, Col: -1}

	hclFile, diagnostics := hclwrite.ParseConfig(src, filePath, hcl.InitialPos)
	if diagnostics != nil && diagnostics.HasErrors() {
		return resourceStart, resourceEnd, lineContent, fmt.Errorf("failed to parse hcl file %s because of errors %s", filePath, diagnostics.Errs())
	}

	hclSyntaxFile, diagnostics := hclsyntax.ParseConfig(src, filePath, hcl.InitialPos)
	if diagnostics != nil && diagnostics.HasErrors() {
		return resourceStart, resourceEnd, lineContent, fmt.Errorf("failed to parse hcl file %s because of errors %s", filePath, diagnostics.Errs())
	}

	if hclFile == nil || hclSyntaxFile == nil {
		return resourceStart, resourceEnd, lineContent, fmt.Errorf("failed to parse hcl file %s", filePath)
	}

	syntaxBlocks := hclSyntaxFile.Body.(*hclsyntax.Body).Blocks
	lines := bytes.Split(src, []byte("\n"))

	if identifyingLine <= 0 || identifyingLine > len(lines) {
		return resourceStart, resourceEnd, lineContent, fmt.Errorf("line %d is out of range", identifyingLine)
	}
	lineContent = string(lines[identifyingLine-1])

	var matchedBlock *hclsyntax.Block
	var nestedBlockMatched bool

	for _, block := range syntaxBlocks {
		blockStart := block.TypeRange.Start
		blockEnd := block.Body.SrcRange.End

		if blockStart.Line <= identifyingLine && identifyingLine <= blockEnd.Line {
			// Search nested blocks
			for _, nestedBlock := range block.Body.Blocks {
				nStart := nestedBlock.TypeRange.Start
				nEnd := nestedBlock.Body.SrcRange.End
				if nStart.Line <= identifyingLine && identifyingLine <= nEnd.Line {
					matchedBlock = nestedBlock
					nestedBlockMatched = true
					break
				}
			}
			if !nestedBlockMatched {
				matchedBlock = block
			}
			break
		}
	}

	if matchedBlock == nil {
		return resourceStart, resourceEnd, lineContent, fmt.Errorf("failed to find block for line %d in file %s", identifyingLine, filePath)
	}

	if nestedBlockMatched {
		// After opening brace of nested block
		openBracePos := matchedBlock.OpenBraceRange.Start
		resourceStart = model.ResourceLine{
			Line: openBracePos.Line,
			Col:  openBracePos.Column + 1, // Right after opening brace
		}
		resourceEnd = model.ResourceLine{
			Line: openBracePos.Line,
			Col:  openBracePos.Column + 1,
		}
	} else {
		// Before closing brace of outer block
		closeBracePos := matchedBlock.Body.SrcRange.End
		resourceStart = model.ResourceLine{
			Line: closeBracePos.Line,
			Col:  1,
		}
		resourceEnd = model.ResourceLine{
			Line: closeBracePos.Line,
			Col:  closeBracePos.Column,
		}
	}

	return resourceStart, resourceEnd, lineContent, nil
}

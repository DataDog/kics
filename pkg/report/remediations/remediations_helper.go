package model

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/Checkmarx/kics/pkg/model"
)

func TransformToSarifFix(vuln model.VulnerableFile, startLocation model.SarifResourceLocation, endLocation model.SarifResourceLocation) (model.SarifFix, error) {
	switch vuln.RemediationType {
	case "replacement":
		return buildReplacementFix(vuln, startLocation, endLocation)
	case "addition":
		return buildAdditionFix(vuln, startLocation)
	case "removal":
		return buildRemovalFix(vuln, startLocation, endLocation)
	default:
		return model.SarifFix{}, fmt.Errorf("unsupported remediation type: %s", vuln.RemediationType)
	}
}

func buildReplacementFix(vuln model.VulnerableFile, startLocation, endLocation model.SarifResourceLocation) (model.SarifFix, error) {
	var patch map[string]string
	if err := json.Unmarshal([]byte(vuln.Remediation), &patch); err != nil {
		return model.SarifFix{}, fmt.Errorf("invalid remediation format for replacement: %v", err)
	}

	before := strings.TrimSpace(patch["before"])
	after := strings.TrimSpace(patch["after"])

	// Regex for key and value (list/quoted/unquoted)
	keyValRegex := regexp.MustCompile(`(?m)["']?(\w+)["']?\s*[:=]\s*(\[.*?\]|".*?"|[^#]+)`)
	matches := keyValRegex.FindStringSubmatch(vuln.LineWithVulnerability)
	if len(matches) < 3 {
		return model.SarifFix{}, fmt.Errorf("could not parse key-value from line: %s", vuln.LineWithVulnerability)
	}

	key := strings.TrimSpace(matches[1])
	rawValue := strings.TrimSpace(matches[2])
	normalizedFullLine := normalize(fmt.Sprintf("%s = %s", key, rawValue))
	normalizedRawValue := normalize(rawValue)

	leadingWhitespace := determineActualBaseIndent(vuln.FileSource, startLocation.Line, vuln.BlockLocation.Start.Line)

	// Extract alternative values (just the value parts)
	altValues := parseAlternatives(before)

	matched := false
	insertedText := ""

	// First try direct match (for flat key=value lines)
	for _, alt := range altValues {
		altKV := normalize(fmt.Sprintf("%s = %s", key, alt))
		if altKV == normalizedFullLine || strings.Contains(normalizedRawValue, alt) {
			matched = true
			break
		}
	}

	// List fallback matching
	if !matched && strings.HasPrefix(rawValue, "[") && strings.HasSuffix(rawValue, "]") {
		// Parse the list from the raw value (preserves quotes)
		listRaw := strings.Trim(rawValue, "[] ")
		items := strings.Split(listRaw, ",")

		replaced := false
		newList := make([]string, 0, len(items))

		for _, item := range items {
			trimmed := strings.TrimSpace(strings.Trim(item, `"`))
			norm := normalize(trimmed)
			replacedItem := false

			for _, alt := range altValues {
				if norm == alt {
					replaced = true
					replacedItem = true
					if strings.HasPrefix(item, `"`) && strings.HasSuffix(item, `"`) {
						newList = append(newList, `"`+after+`"`)
					} else {
						newList = append(newList, after)
					}
					break
				}
			}

			if !replacedItem {
				newList = append(newList, item)
			}
		}

		if replaced {
			rebuilt := "[" + strings.Join(newList, ", ") + "]"
			insertedText = fmt.Sprintf("%s = %s", key, rebuilt)
			matched = true
		}
	}

	if !matched {
		return model.SarifFix{}, fmt.Errorf("line value '%s' does not match any expected values '%s'", normalizedFullLine, strings.Join(altValues, " | "))
	}

	// Preserve quoting if the original value was quoted
	wasQuoted := strings.HasPrefix(rawValue, `"`) && strings.HasSuffix(rawValue, `"`)
	if wasQuoted && !(strings.HasPrefix(after, `"`) && strings.HasSuffix(after, `"`)) {
		after = `"` + after + `"`
	}

	// Fallback: replace entire value
	if insertedText == "" {
		idxs := keyValRegex.FindStringSubmatchIndex(vuln.LineWithVulnerability)
		if len(idxs) < 6 {
			return model.SarifFix{}, fmt.Errorf("could not determine exact value location")
		}

		isFullLine := strings.Contains(after, "=")
		if isFullLine {
			insertedText = leadingWhitespace + after
		} else {
			prefix := vuln.LineWithVulnerability[:idxs[4]]
			suffix := vuln.LineWithVulnerability[idxs[5]:]
			insertedText = prefix + after + suffix
		}
	}

	return buildSarifFix(vuln, insertedText, model.SarifResourceLocation{
		Line: vuln.RemediationLocation.Start.Line,
		Col:  1,
	}, model.SarifResourceLocation{
		Line: vuln.Line,
		Col:  len(vuln.LineWithVulnerability) + 1,
	}, vuln.LineWithVulnerability), nil
}

func buildAdditionFix(vuln model.VulnerableFile, startLocation model.SarifResourceLocation) (model.SarifFix, error) {
	normalized := normalizeIndentation(vuln.Remediation, 2)
	lines := strings.Split(normalized, "\n")
	baseIndent := determineActualBaseIndent(vuln.FileSource, startLocation.Line, vuln.BlockLocation.Start.Line)
	nested := isInsertingInsideNestedBlock(vuln.FileSource, startLocation, vuln.BlockLocation.Start.Line, vuln.BlockLocation.End.Line)

	var result []string
	nestingLevel := 0
	postIndent := baseIndent

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		opening := strings.HasSuffix(trimmed, "{")
		closing := trimmed == "}"

		indent := baseIndent
		if nestingLevel > 0 && !closing {
			indent += strings.Repeat("  ", nestingLevel)
		}
		if closing {
			if len(strings.Repeat("  ", nestingLevel)) > len(baseIndent) {
				postIndent = strings.Repeat("  ", nestingLevel)
			} else {
				postIndent = baseIndent
			}
		}
		if nested && !closing {
			postIndent = baseIndent
		}

		result = append(result, indent+trimmed)

		if opening {
			nestingLevel++
		}
		if closing && nestingLevel > 0 {
			nestingLevel--
		}
	}

	insertedText := "\n" + strings.Join(result, "\n") + "\n" + postIndent
	return buildSarifFix(vuln, insertedText, startLocation, startLocation, ""), nil
}

func buildRemovalFix(vuln model.VulnerableFile, startLocation, endLocation model.SarifResourceLocation) (model.SarifFix, error) {
	return buildSarifFix(vuln, "", startLocation, endLocation, ""), nil
}

func buildSarifFix(vuln model.VulnerableFile, insertedText string, startLocation, endLocation model.SarifResourceLocation, originalLine string) model.SarifFix {
	replacement := model.FixReplacement{
		DeletedRegion: model.SarifRegion{
			StartLine:   startLocation.Line,
			StartColumn: startLocation.Col,
			EndLine:     endLocation.Line,
			EndColumn:   endLocation.Col,
		},
	}
	if insertedText != "" {
		replacement.InsertedContent = model.FixContent{Text: insertedText}
	}

	return model.SarifFix{
		ArtifactChanges: []model.ArtifactChange{{
			ArtifactLocation: model.ArtifactLocation{URI: vuln.FileName},
			Replacements:     []model.FixReplacement{replacement},
		}},
		Description: model.FixMessage{
			Text: fmt.Sprintf("Apply remediation: %s", vuln.Remediation),
		},
	}
}

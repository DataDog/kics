package model

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/Checkmarx/kics/pkg/model"
)

func TransformToSarifFix(vuln model.VulnerableFile, startLocation model.SarifResourceLocation, endLocation model.SarifResourceLocation) (model.SarifFix, error) {
	var insertedText string
	fixStart := startLocation
	fixEnd := endLocation

	switch vuln.RemediationType {
	case "replacement":
		var patch map[string]string
		err := json.Unmarshal([]byte(vuln.Remediation), &patch)
		if err != nil {
			return model.SarifFix{}, fmt.Errorf("invalid remediation format for replacement: %v", err)
		}

		before := strings.TrimSpace(patch["before"])
		after := strings.TrimSpace(patch["after"])

		re := regexp.MustCompile(`(?m)["']?(?P<key>\w+)["']?\s*[:=]\s*(?P<value>\[.*?\]|".*?"|true|false|\d+)[,]?\s*`)
		matches := re.FindStringSubmatch(vuln.LineWithVulnerability)
		if len(matches) < 3 {
			return model.SarifFix{}, fmt.Errorf("could not parse key-value from line: %s", vuln.LineWithVulnerability)
		}

		key := strings.TrimSpace(matches[1])
		value := strings.TrimSpace(matches[2])

		if idx := strings.IndexAny(value, "#/"); idx != -1 {
			value = strings.TrimSpace(value[:idx])
		}

		fullLine := key + " = " + value

		normalizedFullLine := normalize(fullLine)
		normalizedValue := normalize(value)

		// Support multiple 'before' values separated by 'or'
		normalizedAlternatives := []string{normalize(before)}
		if strings.Contains(before, " or ") {
			rawParts := strings.Split(before, " or ")
			normalizedAlternatives = make([]string, 0, len(rawParts))
			for _, part := range rawParts {
				normalizedAlternatives = append(normalizedAlternatives, normalize(strings.TrimSpace(part)))
			}
		}

		insertedText = ""
		matched := false

		for _, alt := range normalizedAlternatives {
			altKey := ""
			altValue := alt

			if parts := strings.SplitN(alt, "=", 2); len(parts) == 2 {
				altKey = normalize(strings.TrimSpace(parts[0]))
				altValue = normalize(strings.TrimSpace(parts[1]))
			} else {
				altValue = normalize(alt)
			}

			if (alt == normalizedFullLine) ||
				(altKey == normalize(key) && altValue == normalizedValue) ||
				strings.Contains(normalizedValue, altValue) ||
				strings.HasPrefix(normalizedValue, altValue) ||
				strings.HasPrefix(altValue, normalizedValue) {
				matched = true
				break
			}
		}

		// Block-style fallback (e.g., nested blocks like metadata_options)
		if !matched && strings.Contains(normalize(before), "{") && strings.Contains(normalize(before), "=") {
			reInner := regexp.MustCompile(`(?m)(\w+)\s*=\s*(".*?"|\[.*?\]|\S+)`)
			for _, inner := range reInner.FindAllString(before, -1) {
				n := normalize(inner)
				for _, alt := range normalizedAlternatives {
					if n == normalizedFullLine || n == normalizedValue || strings.Contains(normalizedFullLine, n) || strings.Contains(n, alt) {
						matched = true
						break
					}
				}
				if matched {
					break
				}
			}
		}

		// List-style fallback
		if !matched && strings.HasPrefix(normalizedValue, "[") && strings.HasSuffix(normalizedValue, "]") {
			listText := strings.Trim(normalizedValue, "[] ")
			listItems := strings.Split(listText, ",")

			newList := make([]string, 0, len(listItems))
			replaced := false

			// Pre-extract values from all alternatives (e.g., extract just "DISABLED" from "key = DISABLED")
			normalizedAltValues := make([]string, 0, len(normalizedAlternatives))
			for _, alt := range normalizedAlternatives {
				parts := strings.SplitN(alt, "=", 2)
				if len(parts) == 2 {
					normalizedAltValues = append(normalizedAltValues, strings.TrimSpace(parts[1]))
				} else {
					normalizedAltValues = append(normalizedAltValues, strings.TrimSpace(alt))
				}
			}

			for _, item := range listItems {
				itemStripped := strings.TrimSpace(strings.Trim(item, `"`))
				normOriginal := normalize(itemStripped)
				replacedItem := false

				for _, altVal := range normalizedAltValues {
					normAltVal := normalize(altVal)

					if normOriginal == normAltVal ||
						strings.Contains(normOriginal, normAltVal) ||
						strings.Contains(normAltVal, normOriginal) {
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
				joined := "[" + strings.Join(newList, ", ") + "]"
				insertedText = strings.Replace(vuln.LineWithVulnerability, value, joined, 1)
				matched = true
			} else {
				// Check if the normalized list as a whole matches one of the alt values
				for _, altVal := range normalizedAltValues {
					if normalizedValue == normalize(altVal) {
						matched = true
						insertedText = vuln.LineWithVulnerability
						break
					}
				}
			}
		}

		if !matched {
			return model.SarifFix{}, fmt.Errorf("line value '%s' does not match any of expected values '%s'", normalizedFullLine, strings.Join(normalizedAlternatives, " | "))
		}

		// Preserve quotes if necessary
		wasQuoted := strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`)
		if wasQuoted && !(strings.HasPrefix(after, `"`) && strings.HasSuffix(after, `"`)) {
			after = `"` + after + `"`
		}

		// Determine replacement range
		idxs := re.FindStringSubmatchIndex(vuln.LineWithVulnerability)
		if len(idxs) < 6 {
			return model.SarifFix{}, fmt.Errorf("could not determine exact value location")
		}

		isFullLine := strings.Contains(after, "=")

		if insertedText == "" {
			if isFullLine {
				insertedText = after
			} else {
				prefix := vuln.LineWithVulnerability[:idxs[4]]
				suffix := vuln.LineWithVulnerability[idxs[5]:]
				insertedText = prefix + after + suffix
			}
		}

		fixStart = model.SarifResourceLocation{
			Line: vuln.RemediationLocation.Start.Line,
			Col:  1,
		}
		fixEnd = model.SarifResourceLocation{
			Line: vuln.Line,
			Col:  len(vuln.LineWithVulnerability) + 1,
		}

	case "addition":
		normalizedRemediation := normalizeIndentation(vuln.Remediation, 2)

		insertedLines := strings.Split(normalizedRemediation, "\n")

		var baseIndent string

		nestedInsert := isInsertingInsideNestedBlock(vuln.FileSource, startLocation, vuln.BlockLocation.Start.Line, vuln.BlockLocation.End.Line)

		baseIndent = determineActualBaseIndent(vuln.FileSource, startLocation.Line, vuln.BlockLocation.Start.Line)

		var result []string
		nestingLevel := 0
		postIndent := baseIndent

		for _, line := range insertedLines {
			trimmed := strings.TrimSpace(line)

			if trimmed == "" {
				// No extra blank line
				continue
			}

			isOpeningBrace := strings.HasSuffix(trimmed, "{")
			isClosingBrace := trimmed == "}"

			// Safely calculate current indentation
			currentIndent := baseIndent

			if nestingLevel > 0 && !isClosingBrace {
				currentIndent += strings.Repeat("  ", nestingLevel)
			}

			// append the base indent after if closing brace
			if isClosingBrace {
				postIndent = strings.Repeat("  ", nestingLevel)
			}

			// if nested insert and not closing brace then add the base indent to the following line
			if nestedInsert && !isClosingBrace {
				postIndent = baseIndent
			}

			// Append the properly indented line
			result = append(result, currentIndent+trimmed)

			// Adjust nesting AFTER appending
			if isOpeningBrace {
				nestingLevel++
			}
			if isClosingBrace && nestingLevel > 0 {
				nestingLevel--
			}
		}

		// Build insertedText cleanly
		insertedText = "\n" + strings.Join(result, "\n")

		// sourceLines := strings.Split(vuln.ResourceSource, "\n")
		// if determineIfShouldAppendIndent(sourceLines, startLocation, vuln.ResourceLocation, vuln.FileSource) {
		// 	if nestedInsert {
		// 		insertedText += "\n" + strings.Repeat(" ", startLocation.Col-1)
		// 	} else {
		// 		insertedText += "\n" + strings.Repeat(" ", startLocation.Col)
		// 	}
		// } else {
		insertedText = insertedText + "\n" + postIndent
		// }

		fixStart = startLocation
		fixEnd = startLocation
	case "removal":
		// Just remove the whole region
		insertedText = ""
		fixStart = startLocation
		fixEnd = endLocation
	}

	replacement := model.FixReplacement{
		DeletedRegion: model.SarifRegion{
			StartLine:   fixStart.Line,
			StartColumn: fixStart.Col,
			EndLine:     fixEnd.Line,
			EndColumn:   fixEnd.Col,
		},
	}

	if insertedText != "" {
		replacement.InsertedContent = model.FixContent{Text: insertedText}
	}

	a := model.SarifFix{
		ArtifactChanges: []model.ArtifactChange{{
			ArtifactLocation: model.ArtifactLocation{URI: vuln.FileName},
			Replacements:     []model.FixReplacement{replacement},
		}},
		Description: model.FixMessage{
			Text: fmt.Sprintf("Apply remediation: %s", vuln.Remediation),
		},
	}

	return a, nil
}

func determineActualBaseIndent(fileLines []string, startLine int, blockStartLine int) string {
	// Try to find the first non-empty, non-comment line above startLine within block
	for i := startLine - 1; i >= blockStartLine-1; i-- {
		line := strings.TrimSpace(fileLines[i])
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}
		return fileLines[i][:len(fileLines[i])-len(strings.TrimLeft(fileLines[i], " \t"))]
	}
	return "" // default fallback
}

// findResourceStartLine searches for where resourceSource starts inside fileSource.
// Returns the starting line number (1-based), or -1 if not found.
func findResourceStartLine(fileSource []string, resourceSource []string) int {
	resourceLines := trimTrailingEmptyLines(resourceSource)

	if len(resourceLines) == 0 {
		return -1
	}

	for i := 0; i <= len(fileSource)-len(resourceLines); i++ {
		match := true
		for j := 0; j < len(resourceLines); j++ {
			if strings.TrimSpace(fileSource[i+j]) != strings.TrimSpace(resourceLines[j]) {
				match = false
				break
			}
		}
		if match {
			return i + 1
		}
	}

	return -1 // Not found
}

// trimTrailingEmptyLines removes trailing empty lines from a slice of lines.
func trimTrailingEmptyLines(lines []string) []string {
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func normalizeIndentation(input string, spacesPerTab int) string {
	lines := strings.Split(input, "\n")
	tabSpaces := strings.Repeat(" ", spacesPerTab)

	for i, line := range lines {
		// Replace tabs with spaces and trim trailing whitespace
		line = strings.ReplaceAll(line, "\t", tabSpaces)
		lines[i] = strings.TrimRight(line, " \t")
	}

	return strings.Join(lines, "\n")
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, `\"`, `"`)     // Handle escaped quotes
	s = strings.ReplaceAll(s, `"`, `"`)      // Double safety
	s = strings.ReplaceAll(s, `“`, `"`)      // Smart quotes (optional)
	s = strings.ReplaceAll(s, `”`, `"`)      // Smart quotes (optional)
	s = strings.Join(strings.Fields(s), " ") // normalize extra whitespace
	s = strings.Trim(s, `"`)
	return s
}

func isInsertingInsideNestedBlock(fileLines []string, startLocation model.SarifResourceLocation, blockStartLine int, blockEndLine int) bool {
	if startLocation.Line <= blockStartLine || startLocation.Line >= blockEndLine {
		// Inserting outside or exactly at block start/end
		return false
	}

	// Read the line we are inserting at
	if startLocation.Line-1 < 0 || startLocation.Line-1 >= len(fileLines) {
		return false
	}

	lineContent := strings.TrimSpace(fileLines[startLocation.Line-1])

	// If inserting into a field or structure inside the block body
	if lineContent == "" {
		// blank line — could be inside, assume not nested by blankness
		return false
	}
	if lineContent == "}" {
		// if right at closing brace, not inside nested body
		return false
	}

	// Now, if there is a `{` before or after nearby, you're likely inside a nested structure
	// (this is heuristic: it covers 99% terraform style)

	// look backwards for a nearby opening `{`
	for i := startLocation.Line - 2; i >= blockStartLine-1; i-- {
		if strings.Contains(fileLines[i], "{") {
			return true // nested structure found
		}
		trimmed := strings.TrimSpace(fileLines[i])
		if trimmed == "" {
			continue
		}
		if trimmed != "}" {
			break // if it's a field or something else, stop
		}
	}
	return false
}

package terraform

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

func GenerateSubstrings(ctx context.Context, key string, extracted [][]string, lines []string, currentLine int) (string, string) {
	var substr1, substr2 string

	// Replace placeholders back to bracketed values
	for idx, str := range extracted {
		placeholder := fmt.Sprintf("{{%d}}", idx)
		key = strings.Replace(key, placeholder, str[0], 1)
	}

	// Handle [something] or ["key"]
	if strings.Contains(key, "[") && strings.HasSuffix(key, "]") {
		start := strings.Index(key, "[")
		end := strings.LastIndex(key, "]")
		if start > 0 && end > start {
			base := key[:start]
			bracketValue := key[start+1 : end]

			// Strip quotes from map-style keys like ["Name"]
			bracketValue = strings.Trim(bracketValue, `"'`)

			// Handle placeholders like [{{label}}]
			if strings.HasPrefix(bracketValue, "{{") && strings.HasSuffix(bracketValue, "}}") {
				bracketValue = strings.TrimPrefix(bracketValue, "{{")
				bracketValue = strings.TrimSuffix(bracketValue, "}}")
			}

			// Handle numeric index
			if index, err := strconv.Atoi(bracketValue); err == nil {
				// Use list logic only if this looks like an actual list
				if looksLikeListAttribute(base, lines) {
					substr1 = base
					substr2 = resolveListIndex(base, index, lines)
					return substr1, substr2
				}
				// Otherwise treat it as a block navigation
				substr1 = base
				substr2 = "" // no value needed
				return substr1, substr2
			}

			// Resource label or map key
			substr1 = base
			substr2 = bracketValue
			return substr1, substr2
		}
	}

	// Fallback: key = value
	parts := strings.SplitN(key, "=", 2)
	substr1 = strings.TrimSpace(parts[0])
	if len(parts) > 1 {
		substr2 = strings.TrimSpace(parts[1])
	}

	// Final fallback: try to resolve value from lines
	if substr2 == "" && substr1 != "" {
		for i := currentLine; i < len(lines); i++ {
			line := lines[i]
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, substr1+" =") {
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					substr2 = strings.TrimSpace(parts[1])
				}
				break
			}
		}
	}

	return substr1, substr2
}

func looksLikeListAttribute(attrName string, lines []string) bool {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, attrName+" = [") {
			return true
		}
	}
	return false
}

func resolveListIndex(attrName string, index int, lines []string) string {
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, attrName+" = [") {
			start := strings.Index(trimmed, "[")
			end := strings.Index(trimmed, "]")
			if start >= 0 && end > start {
				items := strings.Split(trimmed[start+1:end], ",")
				if index < len(items) {
					return strings.Trim(strings.TrimSpace(items[index]), `"`)
				}
			}
		}
	}
	return ""
}

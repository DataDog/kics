package model

import (
	"strings"

	"github.com/Checkmarx/kics/pkg/model"
)

func parseAlternatives(before string) []string {
	if strings.Contains(before, " or ") {
		parts := strings.Split(before, " or ")
		result := make([]string, 0, len(parts))
		for _, part := range parts {
			_, val := splitKeyValue(part)
			result = append(result, normalizeListItem(val))
		}
		return result
	}
	_, val := splitKeyValue(before)
	return []string{normalizeListItem(val)}
}

func normalizeListItem(val string) string {
	val = strings.TrimSpace(val)
	if strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]") {
		val = val[1 : len(val)-1]
		val = strings.TrimSpace(val)
	}
	if strings.HasPrefix(val, `"`) && strings.HasSuffix(val, `"`) {
		val = val[1 : len(val)-1]
	}
	return normalize(val)
}
func splitKeyValue(expr string) (string, string) {
	parts := strings.SplitN(expr, "=", 2)
	if len(parts) == 2 {
		return normalize(strings.TrimSpace(parts[0])), normalize(strings.TrimSpace(parts[1]))
	}
	return "", normalize(expr)
}

func determineActualBaseIndent(fileLines []string, startLine int, blockStartLine int) string {
	if startLine == 0 {
		return ""
	}
	for i := startLine - 1; i >= blockStartLine-1; i-- {
		line := strings.TrimSpace(fileLines[i])
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}
		return fileLines[i][:len(fileLines[i])-len(strings.TrimLeft(fileLines[i], " \t"))]
	}
	return ""
}

func normalizeIndentation(input string, spacesPerTab int) string {
	lines := strings.Split(input, "\n")
	tabSpaces := strings.Repeat(" ", spacesPerTab)

	for i, line := range lines {
		line = strings.ReplaceAll(line, "\t", tabSpaces)
		lines[i] = strings.TrimRight(line, " \t")
	}
	return strings.Join(lines, "\n")
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, `\"`, `"`)
	s = strings.ReplaceAll(s, `“`, `"`)
	s = strings.ReplaceAll(s, `”`, `"`)
	s = strings.Join(strings.Fields(s), " ")
	s = strings.Trim(s, `"`)
	return s
}

func isInsertingInsideNestedBlock(fileLines []string, startLocation model.SarifResourceLocation, blockStartLine, blockEndLine int) bool {
	if startLocation.Line <= blockStartLine || startLocation.Line >= blockEndLine {
		return false
	}

	if startLocation.Line-1 < 0 || startLocation.Line-1 >= len(fileLines) {
		return false
	}

	lineContent := strings.TrimSpace(fileLines[startLocation.Line-1])
	if lineContent == "" || lineContent == "}" {
		return false
	}

	for i := startLocation.Line - 2; i >= blockStartLine-1; i-- {
		if strings.Contains(fileLines[i], "{") {
			return true
		}
		trimmed := strings.TrimSpace(fileLines[i])
		if trimmed == "" {
			continue
		}
		if trimmed != "}" {
			break
		}
	}
	return false
}

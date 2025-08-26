/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package remediation

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Checkmarx/kics/pkg/logger"
)

// Report includes all query results
type Report struct {
	Queries []Query `json:"queries"`
}

// Query includes all the files that presents a result related to the queryID
type Query struct {
	Files   []File `json:"files"`
	QueryID string `json:"query_id"`
}

// File presents the result information related to the file
type File struct {
	FilePath        string `json:"file_name"`
	Line            int    `json:"line"`
	Remediation     string `json:"remediation"`
	RemediationType string `json:"remediation_type"`
	SimilarityID    string `json:"similarity_id"`
	SearchKey       string `json:"search_key"`
	ExpectedValue   string `json:"expected_value"`
	ActualValue     string `json:"actual_value"`
}

// Remediation presents all the relevant information for the fix
type Remediation struct {
	Line          int
	Remediation   string
	SimilarityID  string
	QueryID       string
	SearchKey     string
	ExpectedValue string
	ActualValue   string
}

// Set includes all the replacements and additions related to a file
type Set struct {
	Replacement []Remediation
	Addition    []Remediation
}

// RemediateFile remediationSets the replacements first and secondly, the additions sorted down
func (s *Summary) RemediateFile(ctx context.Context, filePath string, remediationSet Set, openAPIResolveReferences bool, maxResolverDepth int) error {
	logger := logger.FromContext(ctx)
	filepath.Clean(filePath)
	content, err := os.ReadFile(filePath)

	if err != nil {
		logger.Error().Msgf("failed to read file: %s", err)
		return err
	}

	lines := strings.Split(string(content), "\n")

	// do replacements first
	if len(remediationSet.Replacement) > 0 {
		for i := range remediationSet.Replacement {
			r := remediationSet.Replacement[i]
			remediatedLines := replacement(ctx, &r, lines)
			if len(remediatedLines) > 0 && willRemediate(ctx, remediatedLines, filePath, &r, openAPIResolveReferences, maxResolverDepth) {
				lines = s.writeRemediation(ctx, remediatedLines, lines, filePath, r.SimilarityID)
			}
		}
	}

	// do additions after
	if len(remediationSet.Addition) > 0 {
		// descending order
		sort.Slice(remediationSet.Addition, func(i, j int) bool {
			return remediationSet.Addition[i].Line > remediationSet.Addition[j].Line
		})

		for i := range remediationSet.Addition {
			a := remediationSet.Addition[i]
			remediatedLines := addition(ctx, &a, &lines)
			if len(remediatedLines) > 0 && willRemediate(ctx, remediatedLines, filePath, &a, openAPIResolveReferences, maxResolverDepth) {
				lines = s.writeRemediation(ctx, remediatedLines, lines, filePath, a.SimilarityID)
			}
		}
	}

	return nil
}

// ReplacementInfo presents the relevant information to do the replacement
type ReplacementInfo struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

func replacement(ctx context.Context, r *Remediation, lines []string) []string {
	logger := logger.FromContext(ctx)
	originalLine := lines[r.Line-1]

	var replacement ReplacementInfo
	err := json.Unmarshal([]byte(r.Remediation), &replacement)

	if err != nil || replacement == (ReplacementInfo{}) {
		return []string{}
	}

	remediated := strings.Replace(lines[r.Line-1], replacement.Before, replacement.After, 1)

	if originalLine == remediated {
		logger.Info().Msgf("remediation '%s' is already done", r.SimilarityID)
		return []string{}
	}

	// replace the original line with remediation
	lines[r.Line-1] = remediated

	return lines
}

func addition(ctx context.Context, r *Remediation, lines *[]string) []string {
	logger := logger.FromContext(ctx)
	fatherNumberLine := r.Line - 1

	if len(*lines) <= fatherNumberLine+1 {
		return []string{}
	}

	firstLine := strings.Split(r.Remediation, "\n")[0]

	if strings.TrimSpace((*lines)[fatherNumberLine+1]) == strings.TrimSpace(firstLine) {
		logger.Info().Msgf("remediation '%s' is already done", r.SimilarityID)
		return []string{}
	}

	begin := make([]string, len(*lines))
	end := make([]string, len(*lines))

	copy(begin, *lines)
	copy(end, *lines)

	begin = begin[:fatherNumberLine+1]
	end = end[fatherNumberLine+1:]

	before := getBefore((*lines)[fatherNumberLine+1])

	remediation := begin
	remediation = append(remediation, before+r.Remediation)
	remediation = append(remediation, end...)

	return remediation
}

func (s *Summary) writeRemediation(ctx context.Context, remediatedLines, lines []string, filePath, similarityID string) []string {
	logger := logger.FromContext(ctx)
	remediated := []byte(strings.Join(remediatedLines, "\n"))

	if err := os.WriteFile(filePath, remediated, os.ModePerm); err != nil {
		logger.Error().Msgf("failed to write file: %s", err)
		return lines
	}

	logger.Info().Msgf("file '%s' was remediated with '%s'", filePath, similarityID)
	s.ActualRemediationDoneNumber++

	return remediatedLines
}

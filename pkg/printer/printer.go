/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package printer

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/Checkmarx/kics/pkg/logger"
	"github.com/Checkmarx/kics/pkg/utils"

	"github.com/Checkmarx/kics/pkg/model"
	"github.com/gookit/color"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
)

const (
	charsLimitPerLine = 255
)

var (
	consoleLogger = zerolog.ConsoleWriter{Out: io.Discard}
	fileLogger    = zerolog.ConsoleWriter{Out: io.Discard}

	outFileLogger    interface{}
	outConsoleLogger = io.Discard

	loggerFile  interface{}
	initialized bool
)

// Printer wil print console output with colors
// Medium is for medium sevevity results
// High is for high sevevity results
// Low is for low sevevity results
// Info is for info sevevity results
// Success is for successful prints
// Line is the color to print the line with the vulnerability
// minVersion is a bool that if true will print the results output in a minimum version
type Printer struct {
	Critical            color.RGBColor
	Medium              color.RGBColor
	High                color.RGBColor
	Low                 color.RGBColor
	Info                color.RGBColor
	Success             color.RGBColor
	Line                color.RGBColor
	VersionMessage      color.RGBColor
	ContributionMessage color.RGBColor
	minimal             bool
}

// WordWrap Wraps text at the specified number of words
func WordWrap(s, indentation string, limit int) string {
	if strings.TrimSpace(s) == "" {
		return s
	}

	wordSlice := strings.Fields(s)
	var result string

	for len(wordSlice) >= 1 {
		result = result + indentation + strings.Join(wordSlice[:limit], " ") + "\r\n"

		wordSlice = wordSlice[limit:]
		if len(wordSlice) < limit {
			limit = len(wordSlice)
		}
	}
	return result
}

// PrintResult prints on output the summary results
func PrintResult(ctx context.Context, summary *model.Summary, printer *Printer, usingCustomQueries bool, sciInfo model.SCIInfo) error {
	logger := logger.FromContext(ctx)
	logger.Debug().Msg("helpers.PrintResult()")
	fmt.Printf("\n\n")

	for index := range summary.Queries {
		idx := len(summary.Queries) - index - 1
		if summary.Queries[idx].Severity == model.SeverityTrace {
			continue
		}

		logger.Info().Msgf(
			"%s, Severity: %s, Results: %d\n",
			printer.PrintBySev(summary.Queries[idx].QueryName, string(summary.Queries[idx].Severity)),
			printer.PrintBySev(string(summary.Queries[idx].Severity), string(summary.Queries[idx].Severity)),
			len(summary.Queries[idx].Files),
		)

		if summary.Queries[idx].Experimental {
			logger.Info().Msgf("Note: this is an experimental query")
		}

		if !printer.minimal {
			if summary.Queries[idx].CISDescriptionID != "" {
				logger.Info().Msgf("%s %s\n", printer.Bold("Description ID:"), summary.Queries[idx].CISDescriptionIDFormatted)
				logger.Info().Msgf("%s %s\n", printer.Bold("Title:"), summary.Queries[idx].CISDescriptionTitle)
				logger.Info().Msgf("%s %s\n", printer.Bold("Description:"), summary.Queries[idx].CISDescriptionTextFormatted)
			} else {
				logger.Info().Msgf("%s %s\n", printer.Bold("Description:"), summary.Queries[idx].Description)
			}
			logger.Info().Msgf("%s %s\n", printer.Bold("Platform:"), summary.Queries[idx].Platform)

			if summary.Queries[idx].CWE != "" {
				logger.Info().Msgf("%s %s\n", printer.Bold("CWE:"), summary.Queries[idx].CWE)
			}

			// checks if should print queries URL DOCS based on the use of custom queries and invalid ids
			if !usingCustomQueries && validQueryID(summary.Queries[idx].QueryID) {
				queryURLId := summary.Queries[idx].QueryID
				queryURLPlatform := strings.ToLower(summary.Queries[idx].Platform)

				if queryURLPlatform == "common" && strings.Contains(strings.ToLower(summary.Queries[idx].QueryName), "passwords and secrets") {
					queryURLId = "a88baa34-e2ad-44ea-ad6f-8cac87bc7c71"
				}

				logger.Info().Msgf("%s %s\n\n",
					printer.Bold("Learn more about this vulnerability:"),
					fmt.Sprintf("https://docs.kics.io/latest/queries/%s-queries/%s%s",
						queryURLPlatform,
						normalizeURLCloudProvider(summary.Queries[idx].CloudProvider),
						queryURLId))
			}
		}
		printFiles(ctx, &summary.Queries[idx], printer)
	}
	logger.Info().Msgf("\nResults Summary:\n")
	printSeverityCounter(ctx, model.SeverityCritical, summary.SeveritySummary.SeverityCounters[model.SeverityCritical])
	printSeverityCounter(ctx, model.SeverityHigh, summary.SeveritySummary.SeverityCounters[model.SeverityHigh])
	printSeverityCounter(ctx, model.SeverityMedium, summary.SeveritySummary.SeverityCounters[model.SeverityMedium])
	printSeverityCounter(ctx, model.SeverityLow, summary.SeveritySummary.SeverityCounters[model.SeverityLow])
	printSeverityCounter(ctx, model.SeverityInfo, summary.SeveritySummary.SeverityCounters[model.SeverityInfo])
	logger.Info().Msgf("TOTAL: %d\n\n", summary.SeveritySummary.TotalCounter)

	logger.Info().Int64(
		"org", sciInfo.OrgId,
	).Str(
		"branch", sciInfo.RepositoryCommitInfo.Branch,
	).Str(
		"sha", sciInfo.RepositoryCommitInfo.CommitSHA,
	).Str(
		"repository", sciInfo.RepositoryCommitInfo.RepositoryUrl,
	).Msgf("Scanned Files: %d", summary.ScannedFiles)

	logger.Info().Int64(
		"org", sciInfo.OrgId,
	).Str(
		"branch", sciInfo.RepositoryCommitInfo.Branch,
	).Str(
		"sha", sciInfo.RepositoryCommitInfo.CommitSHA,
	).Str(
		"repository", sciInfo.RepositoryCommitInfo.RepositoryUrl,
	).Msgf("Parsed Files: %d", summary.ParsedFiles)

	logger.Info().Msgf("Scanned Lines: %d", summary.ScannedFilesLines)
	logger.Info().Msgf("Parsed Lines: %d", summary.ParsedFilesLines)
	logger.Info().Msgf("Ignored Lines: %d", summary.IgnoredFilesLines)
	logger.Info().Msgf("Queries loaded: %d", summary.TotalQueries)
	logger.Info().Msgf("Queries failed to execute: %d", summary.FailedToExecuteQueries)
	logger.Info().Msg("Inspector stopped")

	return nil
}

func printSeverityCounter(ctx context.Context, severity string, counter int) {
	logger := logger.FromContext(ctx)
	logger.Info().Msgf("%s: %d\n", severity, counter)
}

func printFiles(ctx context.Context, query *model.QueryResult, printer *Printer) {
	logger := logger.FromContext(ctx)
	for fileIdx := range query.Files {
		logger.Info().Msgf("\t%s %s:%s\n", printer.PrintBySev(fmt.Sprintf("[%d]:", fileIdx+1), string(query.Severity)),
			query.Files[fileIdx].FileName, printer.Success.Sprint(query.Files[fileIdx].Line))
		if !printer.minimal {
			fmt.Println()
			for _, line := range *query.Files[fileIdx].VulnLines {
				if len(line.Line) > charsLimitPerLine {
					line.Line = line.Line[:charsLimitPerLine]
				}
				if line.Position == query.Files[fileIdx].Line {
					printer.Line.Printf("\t\t%03d: %s\n", line.Position, line.Line)
				} else {
					logger.Info().Msgf("\t\t%03d: %s\n", line.Position, line.Line)
				}
			}
			fmt.Print("\n\n")
		}
	}
}

// SetupPrinter - configures stdout and log options with given FlagSet
func SetupPrinter(flags *pflag.FlagSet) error {

	initialized = true
	return nil
}

// IsInitialized returns true if printer is ready, false otherwise
func IsInitialized() bool {
	return initialized
}

// NewPrinter initializes a new Printer
func NewPrinter(minimal bool) *Printer {
	return &Printer{
		Critical:            color.HEX("#ff0000"),
		Medium:              color.HEX("#ff7213"),
		High:                color.HEX("#bb2124"),
		Low:                 color.HEX("#edd57e"),
		Success:             color.HEX("#22bb33"),
		Info:                color.HEX("#5bc0de"),
		Line:                color.HEX("#f0ad4e"),
		VersionMessage:      color.HEX("#ff9913"),
		ContributionMessage: color.HEX("ffe313"),
		minimal:             minimal,
	}
}

// PrintBySev will print the output with the specific severity color given the severity of the result
func (p *Printer) PrintBySev(content, sev string) string {
	switch strings.ToUpper(sev) {
	case model.SeverityCritical:
		return p.Critical.Sprintf(content)
	case model.SeverityHigh:
		return p.High.Sprintf(content)
	case model.SeverityMedium:
		return p.Medium.Sprintf(content)
	case model.SeverityLow:
		return p.Low.Sprintf(content)
	case model.SeverityInfo:
		return p.Info.Sprintf(content)
	}
	return content
}

// Bold returns the output in a bold format
func (p *Printer) Bold(content string) string {
	return color.Bold.Sprintf(content)
}

func validQueryID(queryID string) bool {
	if queryID == "" {
		return false
	} else if queryID != "" {
		return utils.ValidateUUID(queryID)
	}
	return true
}

func normalizeURLCloudProvider(cloudProvider string) string {
	cloudProvider = strings.ToLower(cloudProvider)
	if cloudProvider == "common" {
		cloudProvider = ""
	} else if cloudProvider != "" {
		cloudProvider += "/"
	}
	return cloudProvider
}

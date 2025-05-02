/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package scan

import (
	_ "embed" // Embed kics CLI img and scan-flags
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	consoleHelpers "github.com/Checkmarx/kics/internal/console/helpers"
	"github.com/Checkmarx/kics/pkg/engine/provider"
	"github.com/Checkmarx/kics/pkg/model"
	consolePrinter "github.com/Checkmarx/kics/pkg/printer"
	"github.com/Checkmarx/kics/pkg/progress"
	"github.com/Checkmarx/kics/pkg/report"
	"github.com/rs/zerolog/log"
)

func (c *Client) getSummary(results []model.Vulnerability, end time.Time, pathParameters model.PathParameters) model.Summary {
	counters := model.Counters{
		ScannedFiles:           c.Tracker.FoundFiles,
		ScannedFilesLines:      c.Tracker.FoundCountLines,
		ParsedFilesLines:       c.Tracker.ParsedCountLines,
		ParsedFiles:            c.Tracker.ParsedFiles,
		IgnoredFilesLines:      c.Tracker.IgnoreCountLines,
		TotalQueries:           c.Tracker.LoadedQueries,
		FailedToExecuteQueries: c.Tracker.ExecutingQueries - c.Tracker.ExecutedQueries,
		FailedSimilarityID:     c.Tracker.FailedSimilarityID,
		FoundResources:         c.Tracker.FoundResources,
	}

	summary := model.CreateSummary(
		counters,
		results,
		c.ScanParams.ScanID,
		pathParameters.PathExtractionMap,
		c.Tracker.Version,
		c.ScanParams.OutputPath,
	)
	summary.Times = model.Times{
		Start: c.ScanStartTime,
		End:   end,
	}

	// if c.ScanParams.DisableFullDesc {
	// 	log.Warn().Msg("Skipping descriptions because provided disable flag is set")
	// } else {
	// 	err := descriptions.RequestAndOverrideDescriptions(&summary)
	// 	if err != nil {
	// 		log.Warn().Msgf("Unable to get descriptions: %s", err)
	// 		log.Warn().Msgf("Using default descriptions")
	// 	}
	// }

	return summary
}

func (c *Client) resolveOutputs(
	summary *model.Summary,
	documents model.Documents,
	printer *consolePrinter.Printer,
	proBarBuilder progress.PbBuilder,
) error {
	log.Debug().Msg("console.resolveOutputs()")

	usingCustomQueries := usingCustomQueries(c.ScanParams.QueriesPath)
	if err := consolePrinter.PrintResult(summary, printer, usingCustomQueries, c.ScanParams.SCIInfo); err != nil {
		return err
	}
	if c.ScanParams.PayloadPath != "" {
		if err := report.ExportJSONReport(
			filepath.Dir(c.ScanParams.PayloadPath),
			filepath.Base(c.ScanParams.PayloadPath),
			documents,
		); err != nil {
			return err
		}
	}

	return printOutput(
		c.ScanParams.OutputPath,
		c.ScanParams.OutputName,
		summary, c.ScanParams.ReportFormats,
		proBarBuilder,
		c.ScanParams.SCIInfo,
		c.ScanParams.IncludeRemediations,
	)
}

func printOutput(outputPath, filename string, body interface{}, formats []string, proBarBuilder progress.PbBuilder, sciInfo model.SCIInfo, includeRemediations bool) error {
	log.Debug().Msg("console.printOutput()")
	if outputPath == "" {
		return nil
	}
	if len(formats) == 0 {
		formats = []string{"json"}
	}

	log.Debug().Msgf("Output formats provided [%v]", strings.Join(formats, ","))
	log.Debug().Msgf("SCIInfo: %v", sciInfo)
	err := consoleHelpers.GenerateReport(outputPath, filename, body, formats, proBarBuilder, sciInfo, includeRemediations)

	return err
}

// postScan is responsible for the output results
func (c *Client) postScan(scanResults *Results) (ScanMetadata, error) {
	metadata := ScanMetadata{}
	if scanResults == nil {
		log.Info().Msg("No files were scanned")
		scanResults = &Results{
			Results:        []model.Vulnerability{},
			ExtractedPaths: provider.ExtractedPath{},
			Files:          model.FileMetadatas{},
			FailedQueries:  map[string]error{},
		}
	}

	// // mask results preview if Secrets Scan is disabled
	// if c.ScanParams.DisableSecrets {
	// 	err := maskPreviewLines(c.ScanParams.SecretsRegexesPath, scanResults)
	// 	if err != nil {
	// 		log.Err(err)
	// 		return err
	// 	}
	// }
	sort.Strings(c.ScanParams.Path)
	summary := c.getSummary(scanResults.Results, time.Now(), model.PathParameters{
		ScannedPaths:      c.ScanParams.Path,
		PathExtractionMap: scanResults.ExtractedPaths.ExtractionMap,
	})

	if err := c.resolveOutputs(
		&summary,
		scanResults.Files.Combine(c.ScanParams.LineInfoPayload),
		c.Printer,
		*c.ProBarBuilder); err != nil {
		log.Err(err).Msgf("failed to resolve outputs %v", err)
		return metadata, err
	}

	logger := consolePrinter.NewLogger(nil)
	endTime := time.Now()
	scanDuration := endTime.Sub(c.ScanStartTime)
	consolePrinter.PrintScanDuration(&logger, scanDuration)

	log.Info().Int64(
		"org", c.ScanParams.SCIInfo.OrgId,
	).Str(
		"branch", c.ScanParams.SCIInfo.RepositoryCommitInfo.Branch,
	).Str(
		"sha", c.ScanParams.SCIInfo.RepositoryCommitInfo.CommitSHA,
	).Str(
		"repository", c.ScanParams.SCIInfo.RepositoryCommitInfo.RepositoryUrl,
	).Str(
		"exclusion_source", "config_file",
	).Int(
		"excluded_paths", len(c.ScanParams.PreAnalysisExcludePaths),
	).Int(
		"excluded_categories", len(c.ScanParams.ExcludeCategories),
	).Int(
		"excluded_severities", len(c.ScanParams.ExcludeSeverities),
	).Int(
		"excluded_queries", len(c.ScanParams.ExcludeQueries),
	).Int(
		"excluded_results", len(c.ScanParams.ExcludeResults),
	).Msg("Exclusions Info")

	exitCode := consoleHelpers.ResultsExitCode(&summary)
	if consoleHelpers.ShowError("results") && exitCode != 0 {
		os.Exit(exitCode)
	}

	// generate metadata payload
	metadata = c.generateMetadata(scanResults, c.ScanStartTime, endTime)

	return metadata, nil
}

func (c *Client) generateMetadata(scanResults *Results, startTime time.Time, endTime time.Time) ScanMetadata {
	stats := c.generateStats(scanResults, endTime.Sub(startTime))
	ruleStats := c.generateRuleStats(scanResults)

	metadata := ScanMetadata{
		StartTime:      startTime,
		EndTime:        endTime,
		CoresAvailable: int(consoleHelpers.GetNumCPU()),
		DiffAware:      c.ScanParams.SCIInfo.DiffAware.Enabled,
		Stats:          stats,
		RuleStats:      ruleStats,
	}

	return metadata
}

func (c *Client) generateStats(scanResults *Results, scanDuration time.Duration) ScanStats {
	// iterate through scanResults and create a map of severity to count
	violationBreakdowns := make(map[string][]string)
	severitySet := make(map[model.Severity]bool)
	for _, sev := range model.AllSeverities {
		severitySet[sev] = true
	}
	for _, vuln := range scanResults.Results {
		if _, exists := severitySet[vuln.Severity]; exists {
			temp := []string{}
			if _, exists := violationBreakdowns[string(vuln.Severity)]; exists {
				temp = violationBreakdowns[string(vuln.Severity)]
			}
			violationBreakdowns[string(vuln.Severity)] = append(temp, vuln.QueryID)
		}
	}

	return ScanStats{
		Violations:          len(scanResults.Results),
		Files:               c.Tracker.FoundFiles,
		Rules:               c.Tracker.ExecutedQueries,
		Duration:            scanDuration,
		ViolationBreakdowns: violationBreakdowns,
		ResourcesScanned:    c.Tracker.FoundResources,
	}
}

func (c *Client) generateRuleStats(scanResults *Results) RuleStats {
	failedQueries := make([]string, 0, len(scanResults.FailedQueries))
	for _, q := range failedQueries {
		failedQueries = append(failedQueries, q)
	}
	return RuleStats{
		TimedOut:          failedQueries,
		MostExpensiveRule: getLongestRunningQuery(scanResults.Results),
		SlowestRule:       getLongestRunningQuery(scanResults.Results),
	}
}

func getLongestRunningQuery(vulns []model.Vulnerability) RuleTiming {
	longestQuery := ""
	longestQueryTime := time.Duration(0)
	for _, vuln := range vulns {
		if vuln.QueryDuration > longestQueryTime {
			longestQueryTime = vuln.QueryDuration
			longestQuery = vuln.QueryName
		}
	}
	return RuleTiming{
		Name: longestQuery,
		Time: longestQueryTime,
	}
}

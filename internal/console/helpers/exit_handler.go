/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package helpers

import (
	"fmt"
	"strings"

	"github.com/Checkmarx/kics/pkg/model"
)

type ExitHandler struct {
	ShouldIgnore string
	ShouldFail   map[string]struct{}
}

func NewExitHandler() *ExitHandler {
	return &ExitHandler{
		ShouldIgnore: "",
		ShouldFail:   map[string]struct{}{},
	}
}

// ResultsExitCode calculate exit code base on severity of results, returns 0 if no results was reported
func ResultsExitCode(summary *model.Summary) (*ExitHandler, int) {
	exitHandler := NewExitHandler()
	return exitHandler, ResultsExitCodeWithConfig(exitHandler, summary)
}

func ResultsExitCodeWithConfig(config *ExitHandler, summary *model.Summary) int {
	// severityArr is needed to make sure 'for' cycle is made in an ordered fashion
	severityArr := []model.Severity{"CRITICAL", "HIGH", "MEDIUM", "LOW", "INFO", "TRACE"}
	codeMap := map[model.Severity]int{"CRITICAL": 60, "HIGH": 50, "MEDIUM": 40, "LOW": 30, "INFO": 20, "TRACE": 0}
	exitMap := summary.SeveritySummary.SeverityCounters

	for _, severity := range severityArr {
		if _, reportSeverity := config.ShouldFail[strings.ToLower(string(severity))]; !reportSeverity {
			continue
		}
		if exitMap[severity] > 0 {
			return codeMap[severity]
		}
	}
	return 0
}

// InitShouldIgnoreArg initializes what kind of errors should be used on exit codes
func (config *ExitHandler) InitShouldIgnoreArg(arg string) error {
	validArgs := []string{"none", "all", "results", "errors"}
	for _, validArg := range validArgs {
		if strings.EqualFold(validArg, arg) {
			config.ShouldIgnore = strings.ToLower(arg)
			return nil
		}
	}
	return fmt.Errorf("unknown argument for --ignore-on-exit: %s\nvalid arguments:\n  %s", arg, strings.Join(validArgs, "\n  "))
}

// InitShouldFailArg initializes which kind of vulnerability severity should changes exit code
func (config *ExitHandler) InitShouldFailArg(args []string) error {
	possibleArgs := map[string]struct{}{
		"critical": {},
		"high":     {},
		"medium":   {},
		"low":      {},
		"info":     {},
	}

	if len(args) == 0 {
		config.ShouldFail = possibleArgs
		return nil
	}

	argsConverted := make(map[string]struct{})
	for _, arg := range args {
		if _, ok := possibleArgs[strings.ToLower(arg)]; !ok {
			validArgs := []string{"critical", "high", "medium", "low", "info"}
			return fmt.Errorf("unknown argument for --fail-on: %s\nvalid arguments:\n  %s", arg, strings.Join(validArgs, "\n  "))
		}
		argsConverted[strings.ToLower(arg)] = struct{}{}
	}

	config.ShouldFail = argsConverted
	return nil
}

// ShowError returns true if should show error, otherwise returns false
func (config *ExitHandler) ShowError(kind string) bool {
	return strings.EqualFold(config.ShouldIgnore, "none") || (!strings.EqualFold(config.ShouldIgnore, "all") && !strings.EqualFold(config.ShouldIgnore, kind))
}

// RemediateExitCode calculate exit code base on the difference between remediation selected and done
func RemediateExitCode(selectedRemediationNumber, actualRemediationDoneNumber int) int {
	statusCode := 70
	if selectedRemediationNumber != actualRemediationDoneNumber {
		// KICS AR was not able to remediate all the selected remediation
		return statusCode
	}

	return 0
}

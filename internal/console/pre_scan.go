/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package console

import (
	_ "embed" // Embed kics CLI img and scan-flags
	"fmt"
	"runtime"
	"strings"

	consoleHelpers "github.com/Checkmarx/kics/internal/console/helpers"
	"github.com/Checkmarx/kics/internal/constants"
	internalPrinter "github.com/Checkmarx/kics/pkg/printer"
	"github.com/Checkmarx/kics/pkg/progress"
	"github.com/Checkmarx/kics/pkg/scan"

	"github.com/rs/zerolog/log"
)

var (
	//go:embed assets/kics-console
	banner string
)

type console struct {
	Printer       *internalPrinter.Printer
	ProBarBuilder *progress.PbBuilder
}

func newConsole() *console {
	return &console{}
}

// preScan is responsible for scan preparation
func (console *console) preScan(params *scan.Parameters) {
	log.Debug().Msg("console.scan()")

	printer := internalPrinter.NewPrinter(true)
	printer.Success.Printf("\n%s\n", banner)

	versionMsg := fmt.Sprintf("\nScanning with %s\n\n", constants.GetVersion())
	fmt.Println(versionMsg)
	log.Info().Msgf(strings.ReplaceAll(versionMsg, "\n", ""))

	log.Info().Msgf("Operating system: %s", runtime.GOOS)

	cpu := consoleHelpers.GetNumCPU()
	log.Info().Msgf("CPU: %.1f", cpu)

	log.Info().Msgf("Max file size permitted for scanning: %d MB", params.MaxFileSizeFlag)
	log.Info().Msgf("Max resolver depth permitted for resolving files: %d", params.MaxResolverDepth)

	noProgress := false

	proBarBuilder := progress.InitializePbBuilder(
		noProgress,
		false,
		true)

	console.Printer = printer
	console.ProBarBuilder = proBarBuilder
}

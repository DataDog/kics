/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package printer

import (
	"context"
	"io"
	"os"
	"path/filepath"

	consoleHelpers "github.com/Checkmarx/kics/internal/console/helpers"
	"github.com/Checkmarx/kics/internal/constants"
	"github.com/Checkmarx/kics/pkg/logger"
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// CI - enable only log messages to CLI output
func CI(ctx context.Context, opt interface{}) error {
	logger := logger.FromContext(ctx)
	ci := opt.(bool)
	if ci {
		color.SetOutput(io.Discard)
		logger.Output(zerolog.MultiLevelWriter(outConsoleLogger, outFileLogger.(io.Writer)))
		os.Stdout = nil
	}
	return nil
}

// LogFormat - configures the logs format (JSON,pretty).
func LogFormat(ctx context.Context, logFormat string) error {
	logger := logger.FromContext(ctx)
	if logFormat == constants.LogFormatJSON {
		logger.Output(zerolog.MultiLevelWriter(outConsoleLogger, loggerFile.(io.Writer)))
		outFileLogger = loggerFile
		outConsoleLogger = os.Stdout
	} else if logFormat == constants.LogFormatPretty {
		fileLogger = consoleHelpers.CustomConsoleWriter(&zerolog.ConsoleWriter{Out: loggerFile.(io.Writer), NoColor: true})
		logger.Output(zerolog.MultiLevelWriter(consoleLogger, fileLogger))
		outFileLogger = fileLogger
		outConsoleLogger = zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true}
	} else {
		return errors.New("invalid log format")
	}
	return nil
}

// LogPath - sets the log files location
func LogPath(opt interface{}, changed bool) error {
	logPath := opt.(string)
	var err error
	if !changed {
		if loggerFile == nil {
			loggerFile = io.Discard
			return nil
		}
		return nil
	}
	if logPath == "" {
		logPath, err = constants.GetDefaultLogPath()
		if err != nil {
			return err
		}
	} else if filepath.Dir(logPath) != "." {
		if createErr := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); createErr != nil {
			return createErr
		}
	}

	loggerFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

type LogSink struct {
	logs []string
}

func (l *LogSink) Write(p []byte) (n int, err error) {
	l.logs = append(l.logs, string(p))
	return len(p), nil
}

func (l *LogSink) Index(i int) string {
	return l.logs[i]
}

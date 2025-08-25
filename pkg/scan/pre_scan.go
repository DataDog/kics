package scan

import (
	"context"
	"os"
	"path/filepath"
	"time"

	consoleHelpers "github.com/Checkmarx/kics/internal/console/helpers"
	"github.com/Checkmarx/kics/internal/constants"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type ConfigParameters struct {
	ExcludeCategories []string
	ExcludePaths      []string
	ExcludeQueries    []string
	ExcludeResults    []string
	ExcludeSeverities []string
	IncludeQueries    []string
}

func setupConfigFile(ctx context.Context, rootPath string) (bool, error) {
	logger := log.Ctx(ctx)
	configPath := filepath.Join(rootPath, constants.DefaultConfigFilename)
	_, err := os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Info().Msgf("Config file not found at %s", configPath)
			return true, nil
		}
		logger.Info().Msgf("Error reading config file at %s", configPath)
		return true, err
	}

	logger.Info().Msgf("Config file found at %s", configPath)
	return false, nil
}

func initializeConfig(ctx context.Context, rootPath string, extraInfos map[string]string, consolePrint ...bool) (ConfigParameters, context.Context, error) {
	var logCtx context.Context
	baseLogger := log.Logger
	if len(consolePrint) > 0 && consolePrint[0] {
		baseLogger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			With().
			Timestamp().Logger()
		logCtx = baseLogger.WithContext(ctx)
	} else {
		infos := baseLogger.With()
		for k, v := range extraInfos {
			infos = infos.Str(k, v)
		}

		logCtx = infos.Logger().WithContext(ctx)
		baseLogger = infos.Logger()
	}

	baseLogger.Debug().Msg("console.initializeConfig()")

	configParams := ConfigParameters{}

	v := viper.New()
	v.SetEnvPrefix("KICS")
	v.AutomaticEnv()

	exit, err := setupConfigFile(ctx, rootPath)
	if err != nil {
		return configParams, logCtx, err
	}
	if exit {
		return configParams, logCtx, nil
	}
	configPath := filepath.Join(rootPath, constants.DefaultConfigFilename)

	base := filepath.Base(constants.DefaultConfigFilename)
	v.SetConfigName(base)
	v.AddConfigPath(rootPath)
	ext, err := consoleHelpers.FileAnalyzer(configPath)
	if err != nil {
		baseLogger.Debug().Msgf("Error analyzing config file base %s at %s", base, configPath)
		return configParams, logCtx, err
	}
	v.SetConfigType(ext)
	if err := v.ReadInConfig(); err != nil {
		baseLogger.Debug().Msgf("Error reading config file base %s at %s", base, configPath)
		return configParams, logCtx, err
	}

	if v.Get("exclude-categories") != nil {
		configParams.ExcludeCategories = v.GetStringSlice("exclude-categories")
	}
	if v.Get("exclude-paths") != nil {
		configParams.ExcludePaths = v.GetStringSlice("exclude-paths")
	}
	if v.Get("exclude-queries") != nil {
		configParams.ExcludeQueries = v.GetStringSlice("exclude-queries")
	}
	if v.Get("exclude-results") != nil {
		configParams.ExcludeResults = v.GetStringSlice("exclude-results")
	}
	if v.Get("exclude-severities") != nil {
		configParams.ExcludeSeverities = v.GetStringSlice("exclude-severities")
	}
	if v.Get("include-queries") != nil {
		configParams.IncludeQueries = v.GetStringSlice("include-queries")
	}

	return configParams, logCtx, nil
}

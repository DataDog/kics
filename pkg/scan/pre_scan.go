package scan

import (
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

func setupConfigFile(rootPath string) (bool, error) {
	configPath := filepath.Join(rootPath, constants.DefaultConfigFilename)
	_, err := os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info().Msgf("Config file not found at %s", configPath)
			return true, nil
		}
		log.Info().Msgf("Error reading config file at %s", configPath)
		return true, err
	}

	log.Info().Msgf("Config file found at %s", configPath)
	return false, nil
}

func initializeConfig(rootPath string, extraInfos map[string]string, consolePrint ...bool) (ConfigParameters, error) {
	baseLogger := log.Logger
	if len(consolePrint) > 0 && consolePrint[0] {
		baseLogger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			With().
			Timestamp().Logger()
	}

	infos := (&baseLogger).With()
	for k, v := range extraInfos {
		infos = infos.Str(k, v)
	}
	log.Logger = infos.Logger()

	log.Debug().Msg("console.initializeConfig()")

	configParams := ConfigParameters{}

	v := viper.New()
	v.SetEnvPrefix("KICS")
	v.AutomaticEnv()

	exit, err := setupConfigFile(rootPath)
	if err != nil {
		return configParams, err
	}
	if exit {
		return configParams, nil
	}
	configPath := filepath.Join(rootPath, constants.DefaultConfigFilename)

	base := filepath.Base(constants.DefaultConfigFilename)
	v.SetConfigName(base)
	v.AddConfigPath(rootPath)
	ext, err := consoleHelpers.FileAnalyzer(configPath)
	if err != nil {
		log.Debug().Msgf("Error analyzing config file base %s at %s", base, configPath)
		return configParams, err
	}
	v.SetConfigType(ext)
	if err := v.ReadInConfig(); err != nil {
		log.Debug().Msgf("Error reading config file base %s at %s", base, configPath)
		return configParams, err
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

	return configParams, nil
}

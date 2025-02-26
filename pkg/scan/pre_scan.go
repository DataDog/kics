package scan

import (
	"os"
	"path/filepath"

	consoleHelpers "github.com/Checkmarx/kics/internal/console/helpers"
	"github.com/Checkmarx/kics/internal/constants"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type ConfigParameters struct {
	ExcludeCategories []string
	ExcludePaths      []string
	ExcludeQueries    []string
	ExcludeResults    []string
	ExcludeSeverities []string
}

func setupConfigFile() (bool, error) {
	_, err := os.Stat(constants.DefaultConfigFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return true, nil
		}
		return true, err
	}

	return false, nil
}

func initializeConfig(rootPath string) (ConfigParameters, error) {
	log.Debug().Msg("console.initializeConfig()")

	configParams := ConfigParameters{}

	v := viper.New()
	v.SetEnvPrefix("KICS")
	v.AutomaticEnv()

	exit, err := setupConfigFile()
	if err != nil {
		return configParams, err
	}
	if exit {
		return configParams, nil
	}
	configPath := filepath.Join(rootPath, constants.DefaultConfigFilename)

	base := filepath.Base(configPath)
	v.SetConfigName(base)
	v.AddConfigPath(filepath.Dir(base))
	ext, err := consoleHelpers.FileAnalyzer(base)
	if err != nil {
		return configParams, err
	}
	v.SetConfigType(ext)
	if err := v.ReadInConfig(); err != nil {
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

	return configParams, nil
}

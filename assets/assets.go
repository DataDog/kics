package assets

import (
	"embed" // used for embedding KICS libraries
	"io/fs"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

//go:embed libraries/*.rego
var embeddedLibraries embed.FS

//go:embed libraries/*.json
var embeddedLibraryData embed.FS

// GetEmbeddedLibrary returns the embedded library.rego for the platform passed in the argument
func GetEmbeddedLibrary(platform string) (string, error) {
	content, err := embeddedLibraries.ReadFile("libraries/" + platform + ".rego")

	return string(content), err
}

// GetEmbeddedLibrary returns the embedded library.rego for the platform passed in the argument
func GetEmbeddedLibraryData(platform string) (string, error) {
	content, err := embeddedLibraryData.ReadFile("libraries/" + platform + ".json")

	return string(content), err
}

//go:embed queries/cicd
var cicdEmbeddedQueries embed.FS

//go:embed queries/terraform
var terraformEmbeddedQueries embed.FS

var embeddedQueries []embed.FS = []embed.FS{cicdEmbeddedQueries, terraformEmbeddedQueries}

func GetEmbeddedQueryDirs() ([]string, error) {
	var out []string
	for _, embeddedQueriesList := range embeddedQueries {
		err := fs.WalkDir(embeddedQueriesList, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				log.Info().Msgf("Failed to walk directory: %s", path)
				return err
			}
			baseDir := filepath.Base(path)
			if baseDir == "test" {
				return nil
			}
			if d.IsDir() {
				out = append(out, path)
			}
			return nil
		})
		if err != nil {
			log.Error().Msgf("Failed to walk embedded directory for queries: %v", err)
			return nil, err
		}
	}
	return out, nil
}

func GetEmbeddedQueryFile(path string) (string, error) {
	var content []byte
	var err error
	for _, embeddedQueriesList := range embeddedQueries {
		content, err = embeddedQueriesList.ReadFile(path)
		if err == nil {
			return string(content), nil
		}
	}
	log.Debug().Msgf("Failed to read file for path %s: %v", path, err)
	return "", err
}

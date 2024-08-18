package assets

import (
	"embed" // used for embedding KICS libraries
	"fmt"
	"io/fs"

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

//go:embed queries
var embeddedQueries embed.FS

func GetEmbeddedQueryDirs() ([]string, error) {
	var out []string
	err := fs.WalkDir(embeddedQueries, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Info().Msgf("Failed to walk directory: %s", path)
			return err
		}
		fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())
		if d.IsDir() {
			out = append(out, path)
		}
		return nil
	})
	if err != nil {
		log.Error().Msgf("Failed to walk embedded directory for queries: %v", err)
		return nil, err
	}
	return out, nil
}

func GetEmbeddedQueryFile(path string) (string, error) {
	content, err := embeddedQueries.ReadFile(path)
	if err != nil {
		log.Error().Msgf("Failed to read file for path %s: %v", path, err)
		return "", err
	}
	return string(content), nil
}

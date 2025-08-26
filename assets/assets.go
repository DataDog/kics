package assets

import (
	"context"
	"embed" // used for embedding KICS libraries
	"io/fs"
	"path/filepath"

	"github.com/Checkmarx/kics/pkg/logger"
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

//go:embed queries/terraform
var embeddedQueries embed.FS

func GetEmbeddedQueryDirs(ctx context.Context) ([]string, error) {
	logger := logger.FromContext(ctx)
	var out []string
	err := fs.WalkDir(embeddedQueries, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			logger.Info().Msgf("Failed to walk directory: %s", path)
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
		logger.Error().Msgf("Failed to walk embedded directory for queries: %v", err)
		return nil, err
	}
	return out, nil
}

func GetEmbeddedQueryFile(ctx context.Context, path string) (string, error) {
	logger := logger.FromContext(ctx)
	content, err := embeddedQueries.ReadFile(path)
	if err != nil {
		logger.Debug().Msgf("Failed to read file for path %s: %v", path, err)
		return "", err
	}
	return string(content), nil
}

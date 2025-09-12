/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package source

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Checkmarx/kics/assets"
	"github.com/Checkmarx/kics/internal/constants"
	"github.com/Checkmarx/kics/pkg/logger"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/pkg/errors"
)

// FilesystemSource this type defines a struct with a path to a filesystem source of queries
// Source is the path to the queries
// Types are the types given by the flag --type for query selection mechanism
type FilesystemSource struct {
	Source              []string
	Types               []string
	CloudProviders      []string
	Library             string
	ExperimentalQueries bool
}

const (
	// QueryFileName The default query file name
	QueryFileName = "query.rego"
	// MetadataFileName The default metadata file name
	MetadataFileName = "metadata.json"
	// LibrariesDefaultBasePath the path to rego libraries
	LibrariesDefaultBasePath = "./assets/libraries"

	emptyInputData = "{}"

	common = "Common"

	kicsDefault = "default"
)

// NewFilesystemSource initializes a NewFilesystemSource with source to queries and types of queries to load
func NewFilesystemSource(ctx context.Context, source, types, cloudProviders []string, libraryPath string, experimentalQueries bool) *FilesystemSource {
	logger := logger.FromContext(ctx)
	logger.Debug().Msg("source.NewFilesystemSource()")

	if len(types) == 0 {
		types = []string{""}
	}

	if len(cloudProviders) == 0 {
		cloudProviders = []string{""}
	}

	for s := range source {
		source[s] = filepath.FromSlash(source[s])
	}

	return &FilesystemSource{
		Source:              source,
		Types:               types,
		CloudProviders:      cloudProviders,
		Library:             filepath.FromSlash(libraryPath),
		ExperimentalQueries: experimentalQueries,
	}
}

// ListSupportedPlatforms returns a list of supported platforms
func ListSupportedPlatforms() []string {
	keys := make([]string, len(constants.AvailablePlatforms))
	i := 0
	for k := range constants.AvailablePlatforms {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

// ListSupportedCloudProviders returns a list of supported cloud providers
func ListSupportedCloudProviders() []string {
	return []string{"alicloud", "aws", "azure", "gcp", "nifcloud", "tencentcloud"}
}

func getLibraryInDir(ctx context.Context, platform, libraryDirPath string) string {
	logger := logger.FromContext(ctx)
	var libraryFilePath string
	err := filepath.Walk(libraryDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.EqualFold(filepath.Base(path), platform+".rego") { // try to find the library file <platform>.rego
			libraryFilePath = path
		}
		return nil
	})
	if err != nil {
		logger.Error().Msgf("Failed to analyze path %s: %s", libraryDirPath, err)
	}
	return libraryFilePath
}

func isDefaultLibrary(libraryPath string) bool {
	return filepath.FromSlash(libraryPath) == filepath.FromSlash(LibrariesDefaultBasePath)
}

// GetPathToCustomLibrary - returns the libraries path for a given platform
func GetPathToCustomLibrary(ctx context.Context, platform, libraryPathFlag string) string {
	logger := logger.FromContext(ctx)
	libraryFilePath := kicsDefault

	if !isDefaultLibrary(libraryPathFlag) {
		logger.Debug().Msgf("Trying to load custom libraries from %s", libraryPathFlag)

		library := getLibraryInDir(ctx, platform, libraryPathFlag)
		// found a library named according to the platform
		if library != "" {
			libraryFilePath = library
		}
	}

	return libraryFilePath
}

// GetQueryLibrary returns the library.rego for the platform passed in the argument
func (s *FilesystemSource) GetQueryLibrary(ctx context.Context, platform string) (RegoLibraries, error) {
	logger := logger.FromContext(ctx)
	library := GetPathToCustomLibrary(ctx, platform, s.Library)

	if library == "" {
		return RegoLibraries{}, errors.New("unable to get libraries path")
	}

	// getting embedded library
	embeddedLibraryCode, errGettingEmbeddedLibrary := assets.GetEmbeddedLibrary(strings.ToLower(platform))
	if errGettingEmbeddedLibrary != nil {
		return RegoLibraries{}, errGettingEmbeddedLibrary
	}

	embeddedLibraryData, errGettingEmbeddedLibraryCode := assets.GetEmbeddedLibraryData(strings.ToLower(platform))
	if errGettingEmbeddedLibraryCode != nil {
		logger.Debug().Msgf("Could not open embedded library data for %s platform", platform)
		embeddedLibraryData = emptyInputData
	}

	regoLibrary := RegoLibraries{
		LibraryCode:      embeddedLibraryCode,
		LibraryInputData: embeddedLibraryData,
	}
	return regoLibrary, nil
}

// CheckType checks if the queries have the type passed as an argument in '--type' flag to be loaded
func (s *FilesystemSource) CheckType(queryPlatform interface{}) bool {
	if queryPlatform.(string) == common {
		return true
	}
	if s.Types[0] != "" {
		for _, t := range s.Types {
			if strings.EqualFold(t, queryPlatform.(string)) {
				return true
			}
		}
		return false
	}
	return true
}

// CheckCloudProvider checks if the queries have the cloud provider passed as an argument in '--cloud-provider' flag to be loaded
func (s *FilesystemSource) CheckCloudProvider(cloudProvider interface{}) bool {
	if cloudProvider != nil {
		if strings.EqualFold(cloudProvider.(string), common) {
			return true
		}
		if s.CloudProviders[0] != "" {
			return strings.Contains(strings.ToUpper(strings.Join(s.CloudProviders, ",")), strings.ToUpper(cloudProvider.(string)))
		}
	}

	if s.CloudProviders[0] == "" {
		return true
	}

	return false
}

func checkQueryInclude(ctx context.Context, id interface{}, includedQueries []string) bool {
	logger := logger.FromContext(ctx)
	queryMetadataKey, ok := id.(string)
	if !ok {
		logger.Warn().
			Msgf("Can't cast query metadata key = %v", id)
		return false
	}
	for _, includedQuery := range includedQueries {
		if queryMetadataKey == includedQuery {
			return true
		}
	}
	return false
}

func checkQueryExcludeField(ctx context.Context, id interface{}, excludeQueries []string) bool {
	logger := logger.FromContext(ctx)
	queryMetadataKey, ok := id.(string)
	if !ok {
		logger.Warn().
			Msgf("Can't cast query metadata key = %v", id)
		return false
	}
	for _, excludedQuery := range excludeQueries {
		if strings.EqualFold(queryMetadataKey, excludedQuery) {
			return true
		}
	}
	return false
}

func checkQueryExclude(ctx context.Context, metadata map[string]interface{}, queryParameters *QueryInspectorParameters) bool {
	return checkQueryExcludeField(ctx, metadata["id"], queryParameters.ExcludeQueries.ByIDs) ||
		checkQueryExcludeField(ctx, metadata["category"], queryParameters.ExcludeQueries.ByCategories) ||
		checkQueryExcludeField(ctx, metadata["severity"], queryParameters.ExcludeQueries.BySeverities) ||
		(!queryParameters.BomQueries && metadata["severity"] == model.SeverityTrace) ||
		checkQueryFeatureFlagDisabled(ctx, metadata, queryParameters)
}

func checkQueryFeatureFlagDisabled(ctx context.Context, metadata map[string]interface{}, queryParameters *QueryInspectorParameters) bool {
	if queryParameters.FlagEvaluator == nil {
		return false
	}

	// Extract KICS ID from query metadata
	kicsID, exists := metadata["id"]
	if !exists {
		return false
	}

	kicsIDStr, ok := kicsID.(string)
	if !ok {
		return false
	}

	// Create custom variables with the KICS ID
	customVariables := map[string]interface{}{
		"KICS_RULE_ID": kicsIDStr,
	}

	logger := logger.FromContext(ctx)
	// Check if the rule is disabled via feature flag
	disabled, err := queryParameters.FlagEvaluator.EvaluateWithOrgAndCustomVariables("k9-iac-disable-kics-rule", customVariables)
	if err != nil {
		// If feature flag evaluation fails, log and continue (fail open)
		logger.Warn().Err(err).Str("kics_id", kicsIDStr).Msg("Failed to evaluate feature flag for KICS rule")
		return false
	}

	if disabled {
		logger.Info().Str("kics_id", kicsIDStr).Msg("KICS rule disabled by feature flag")
	}

	return disabled
}

// GetQueries walks a given filesource path returns all queries found in an array of
// QueryMetadata struct
func (s *FilesystemSource) GetQueries(ctx context.Context, queryParameters *QueryInspectorParameters) ([]model.QueryMetadata, error) {
	logger := logger.FromContext(ctx)
	// queryDirs, err := s.iterateSources()
	logger.Info().Msg("iterateEmbeddedQuerySources()")
	dirs, err := s.iterateEmbeddedQuerySources(ctx)
	if err != nil {
		return nil, err
	}

	queries := s.iterateQueryDirs(ctx, dirs, queryParameters)

	return queries, nil
}

func (s *FilesystemSource) iterateSources() ([]string, error) {
	queryDirs := make([]string, 0)

	for _, source := range s.Source {
		err := filepath.Walk(source,
			func(p string, f os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if f.IsDir() || f.Name() != QueryFileName {
					return nil
				}

				querypathDir := filepath.Dir(p)

				if err == nil {
					queryDirs = append(queryDirs, querypathDir)
				} else if err != nil {
					return errors.Wrap(err, "Failed to get query relative path")
				}

				return nil
			})
		if err != nil {
			return nil, errors.Wrap(err, "failed to get query Source")
		}
	}

	return queryDirs, nil
}

func getAllDirs(ctx context.Context, embedfs *embed.FS, path string) ([]string, error) {
	logger := logger.FromContext(ctx)
	var out []string
	err := fs.WalkDir(embedfs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			logger.Error().Msgf("Failed to walk directory: %s", path)
			return err
		}
		logger.Info().Msgf("path=%q, isDir=%v\n", path, d.IsDir())
		if d.IsDir() {
			out = append(out, path)
		}
		return nil
	})
	if err != nil {
		logger.Error().Msgf("Failed to walk directory: %s: %v", path, err)
		return nil, err
	}
	return out, nil
}

// iterate over the embedded query directory and read the respective queries
func (s *FilesystemSource) iterateEmbeddedQuerySources(ctx context.Context) ([]string, error) {
	logger := logger.FromContext(ctx)
	// dirEntries, err := queryDir.ReadDir(".")
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to read embedded query directory")
	// }
	logger.Info().Msg("getAllDirs()")

	queryDirs, err := assets.GetEmbeddedQueryDirs(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get query sources")
	}
	// err := fs.WalkDir(queryDir, ".", func(p string, d fs.DirEntry, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if d.IsDir() || d.Name() != QueryFileName {
	// 		return nil
	// 	}

	// 	querypathDir := filepath.Dir(p)

	// 	logger.Info().Msgf("Query path: %s", querypathDir)
	// 	queryDirs = append(queryDirs, querypathDir)

	// 	return nil
	// })

	if err != nil {
		logger.Error().Msgf("failed to get query Source: %v", err)
	}

	return queryDirs, nil
}

// iterateQueryDirs iterates all query directories and reads the respective queries
func (s *FilesystemSource) iterateQueryDirs(ctx context.Context, queryDirs []string, queryParameters *QueryInspectorParameters) []model.QueryMetadata {
	// logger := logger.FromContext(ctx)
	queries := make([]model.QueryMetadata, 0, len(queryDirs))

	for _, queryDir := range queryDirs {
		query, errRQ := ReadQuery(ctx, queryDir)
		if errRQ != nil {
			// sentryReport.ReportSentry(&sentryReport.Report{
			// 	Message:  fmt.Sprintf("Query provider failed to read query, query=%s", path.Base(queryDir)),
			// 	Err:      errRQ,
			// 	Location: "func GetQueries()",
			// 	FileName: path.Base(queryDir),
			// }, true)
			continue
		}

		if query.Experimental && !queryParameters.ExperimentalQueries {
			continue
		}

		if !s.CheckType(query.Metadata["platform"]) {
			continue
		}

		if !s.CheckCloudProvider(query.Metadata["cloudProvider"]) {
			continue
		}

		// customInputData, readInputErr := readInputData(baseDir, filepath.Join(queryParameters.InputDataPath, query.Metadata["id"].(string)+".json"))
		// if readInputErr != nil {
		// 	logger.Err(errRQ).
		// 		Msgf("failed to read input data, query=%s", path.Base(queryDir))
		// 	continue
		// }

		// inputData, mergeError := MergeInputData(query.InputData, customInputData)
		// if mergeError != nil {
		// 	logger.Err(mergeError).
		// 		Msgf("failed to merge input data, query=%s", path.Base(queryDir))
		// 	continue
		// }
		// query.InputData = inputData

		if len(queryParameters.IncludeQueries.ByIDs) > 0 {
			if checkQueryInclude(ctx, query.Metadata["id"], queryParameters.IncludeQueries.ByIDs) {
				queries = append(queries, query)
			}
		} else {
			if checkQueryExclude(ctx, query.Metadata, queryParameters) {
				continue
			}

			queries = append(queries, query)
		}
	}
	return queries
}

// validateMetadata prevents panics when KICS queries metadata fields are missing
func validateMetadata(metadata map[string]interface{}) (exist bool, field string) {
	fields := []string{
		"id",
		"platform",
	}
	for _, field = range fields {
		if _, exist = metadata[field]; !exist {
			return
		}
	}
	return
}

// ReadQuery reads query's files for a given path and returns a QueryMetadata struct with it's
// content
func ReadQuery(ctx context.Context, queryDir string) (model.QueryMetadata, error) {
	logger := logger.FromContext(ctx)
	queryContent, err := assets.GetEmbeddedQueryFile(ctx, path.Join(queryDir, QueryFileName))
	// queryContent, err := os.ReadFile(filepath.Clean(path.Join(queryDir, QueryFileName)))
	if err != nil {
		return model.QueryMetadata{}, errors.Wrapf(err, "failed to read query %s", path.Base(queryDir))
	}

	metadata, err := ReadMetadata(ctx, queryDir)
	if err != nil {
		return model.QueryMetadata{}, errors.Wrapf(err, "failed to read query %s", path.Base(queryDir))
	}

	if valid, missingField := validateMetadata(metadata); !valid {
		err := fmt.Errorf("failed to read metadata field: %s", missingField)
		logger.Error().Msg(err.Error())
		return model.QueryMetadata{}, err
	}

	platform := getPlatform(metadata["platform"].(string))

	aggregation := 1
	if agg, ok := metadata["aggregation"]; ok {
		aggregation = int(agg.(float64))
	}

	experimental := getExperimental(metadata["experimental"])

	return model.QueryMetadata{
		Query:        path.Base(filepath.ToSlash(queryDir)),
		Content:      string(queryContent),
		Metadata:     metadata,
		Platform:     platform,
		InputData:    "{}",
		Aggregation:  aggregation,
		Experimental: experimental,
	}, nil
}

// ReadMetadata read query's metadata file inside the query directory
func ReadMetadata(ctx context.Context, queryDir string) (map[string]interface{}, error) {
	logger := logger.FromContext(ctx)
	f, err := assets.GetEmbeddedQueryFile(ctx, filepath.Clean(path.Join(queryDir, MetadataFileName)))
	// f, err := os.Open(filepath.Clean(path.Join(queryDir, MetadataFileName)))
	if err != nil {
		logger.Error().Msgf("Queries provider can't read metadata, query=%s: %v", path.Base(queryDir), err)
		return nil, err
	}

	var metadata map[string]interface{}
	if err := json.Unmarshal([]byte(f), &metadata); err != nil {
		return nil, err
	}

	return metadata, nil
}

type supportedPlatforms map[string]string

var supPlatforms = &supportedPlatforms{
	"Ansible":                 "ansible",
	"CloudFormation":          "cloudFormation",
	"Common":                  "common",
	"Crossplane":              "crossplane",
	"Knative":                 "knative",
	"Kubernetes":              "k8s",
	"OpenAPI":                 "openAPI",
	"Terraform":               "terraform",
	"AzureResourceManager":    "azureResourceManager",
	"GRPC":                    "grpc",
	"GoogleDeploymentManager": "googleDeploymentManager",
	"Buildah":                 "buildah",
	"Pulumi":                  "pulumi",
	"ServerlessFW":            "serverlessFW",
	"CICD":                    "cicd",
}

func getPlatform(metadataPlatform string) string {
	if p, ok := (*supPlatforms)[metadataPlatform]; ok {
		return p
	}
	return "unknown"
}

func getExperimental(experimental interface{}) bool {
	readExperimental, _ := experimental.(string)
	if readExperimental == "true" {
		return true
	} else {
		return false
	}
}

func readInputData(ctx context.Context, inputDataPath string) (string, error) {
	logger := logger.FromContext(ctx)
	inputData, err := assets.GetEmbeddedQueryFile(ctx, filepath.Clean(inputDataPath))
	if err != nil {
		if os.IsNotExist(err) {
			return emptyInputData, nil
		}
		return emptyInputData, errors.Wrapf(err, "failed to read query input data %s", path.Base(inputDataPath))
	}
	logger.Info().Msgf("Input data found in file: %s", string(inputData))
	return string(inputData), nil
}

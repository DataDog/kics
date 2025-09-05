/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package remediation

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Checkmarx/kics/pkg/engine"
	"github.com/Checkmarx/kics/pkg/kics"
	"github.com/Checkmarx/kics/pkg/logger"
	"github.com/Checkmarx/kics/pkg/minified"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/Checkmarx/kics/pkg/scan"
	"github.com/Checkmarx/kics/pkg/utils"
	"github.com/open-policy-agent/opa/topdown"

	"github.com/Checkmarx/kics/internal/tracker"
	"github.com/Checkmarx/kics/pkg/engine/source"
	"github.com/Checkmarx/kics/pkg/parser"
	buildahParser "github.com/Checkmarx/kics/pkg/parser/buildah"
	protoParser "github.com/Checkmarx/kics/pkg/parser/grpc"
	jsonParser "github.com/Checkmarx/kics/pkg/parser/json"
	terraformParser "github.com/Checkmarx/kics/pkg/parser/terraform"
	yamlParser "github.com/Checkmarx/kics/pkg/parser/yaml"
	"github.com/open-policy-agent/opa/rego"
)

type runQueryInfo struct {
	payload   model.Documents
	query     *engine.PreparedQuery
	inspector *engine.Inspector
	tmpFile   string
	files     model.FileMetadatas
}

// scanTmpFile scans a temporary file against a specific query
func scanTmpFile(
	ctx context.Context,
	tmpFile, queryID string,
	remediated []byte,
	openAPIResolveReferences bool,
	maxResolverDepth int) ([]model.Vulnerability, error) {
	logger := logger.FromContext(ctx)
	// get payload
	files, err := getPayload(ctx, tmpFile, remediated, openAPIResolveReferences, maxResolverDepth)

	if err != nil {
		logger.Err(err).Msg("didn't get payload")
		return []model.Vulnerability{}, err
	}

	if len(files) == 0 {
		logger.Error().Msg("failed to get payload")
		return []model.Vulnerability{}, errors.New("failed to get payload")
	}

	payload := files.Combine(ctx, false)

	// init scan
	inspector, err := initScan(ctx, queryID)

	if err != nil {
		logger.Err(err).Msg("")
		return []model.Vulnerability{}, err
	}
	logger.Info().Msg("Scan initialized")

	// load query
	logger.Info().Msg("Loading query")
	query, err := loadQuery(ctx, inspector, queryID)

	if err != nil {
		logger.Err(err)
		return []model.Vulnerability{}, err
	}

	// run query
	info := &runQueryInfo{
		payload:   payload,
		query:     query,
		inspector: inspector,
		tmpFile:   tmpFile,
		files:     files,
	}

	return runQuery(ctx, info), nil
}

// getPayload gets the payload of a file
func getPayload(ctx context.Context, filePath string, content []byte, openAPIResolveReferences bool, maxResolverDepth int) (model.FileMetadatas, error) {
	logger := logger.FromContext(ctx)
	ext, _ := utils.GetExtension(ctx, filePath)
	var p []*parser.Parser
	var err error

	switch ext {
	case ".tf":
		p, err = parser.NewBuilder(ctx).Add(terraformParser.NewDefault()).Build([]string{""}, []string{""})

	case ".proto":
		p, err = parser.NewBuilder(ctx).Add(&protoParser.Parser{}).Build([]string{""}, []string{""})

	case ".yaml", ".yml":
		p, err = parser.NewBuilder(ctx).Add(&yamlParser.Parser{}).Build([]string{""}, []string{""})

	case ".json":
		p, err = parser.NewBuilder(ctx).Add(&jsonParser.Parser{}).Build([]string{""}, []string{""})

	case ".sh":
		p, err = parser.NewBuilder(ctx).Add(&buildahParser.Parser{}).Build([]string{""}, []string{""})
	}

	if err != nil {
		logger.Error().Msgf("failed to get parser: %s", err)
		return model.FileMetadatas{}, err
	}

	if len(p) == 0 {
		logger.Info().Msg("failed to get parser")
		return model.FileMetadatas{}, errors.New("failed to get parser")
	}

	isMinified := minified.IsMinified(filePath, content)
	documents, er := p[0].Parse(ctx, filePath, content, openAPIResolveReferences, isMinified, maxResolverDepth)

	if er != nil {
		logger.Error().Msgf("failed to parse file '%s': %s", filePath, er)
		return model.FileMetadatas{}, er
	}

	var files model.FileMetadatas

	for _, document := range documents.Docs {
		_, err = json.Marshal(document)
		if err != nil {
			continue
		}

		file := model.FileMetadata{
			FilePath:          filePath,
			Document:          kics.PrepareScanDocument(ctx, document, documents.Kind),
			LineInfoDocument:  document,
			Commands:          p[0].CommentsCommands(ctx, filePath, content),
			OriginalData:      string(content),
			LinesOriginalData: utils.SplitLines(string(content)),
			IsMinified:        documents.IsMinified,
		}

		files = append(files, file)
	}

	return files, nil
}

// runQuery runs a query and returns its results
func runQuery(ctx context.Context, r *runQueryInfo) []model.Vulnerability {
	logger := logger.FromContext(ctx)
	queryStart := time.Now()
	queryExecTimeout := time.Duration(60) * time.Second

	timeoutCtx, cancel := context.WithTimeout(context.Background(), queryExecTimeout)
	defer cancel()

	options := []rego.EvalOption{rego.EvalInput(r.payload)}

	results, err := r.query.OpaQuery.Eval(timeoutCtx, options...)

	if err != nil {
		if topdown.IsCancel(err) {
			logger.Err(err)
		}

		logger.Err(err)
	}

	bcgCtx := context.Background()

	queryCtx := &engine.QueryContext{
		Ctx:           bcgCtx,
		Query:         r.query,
		BaseScanPaths: []string{r.tmpFile},
		Files:         r.files.ToMap(),
	}

	timeoutCtxToDecode, cancelDecode := context.WithTimeout(context.Background(), queryExecTimeout)
	defer cancelDecode()
	decoded, err := r.inspector.DecodeQueryResults(ctx, queryCtx, timeoutCtxToDecode, results, time.Since(queryStart))

	if err != nil {
		logger.Err(err)
	}

	return decoded
}

func initScan(ctx context.Context, queryID string) (*engine.Inspector, error) {
	logger := logger.FromContext(ctx)
	scanParams := &scan.Parameters{
		CloudProvider:               []string{""},
		DisableFullDesc:             false,
		ExcludeCategories:           []string{},
		ExcludeQueries:              []string{},
		ExcludeResults:              []string{},
		ExcludeSeverities:           []string{},
		ExcludePaths:                []string{},
		ExperimentalQueries:         false,
		IncludeQueries:              []string{},
		InputData:                   "",
		OutputName:                  "kics-result",
		PayloadPath:                 "",
		PreviewLines:                3,
		QueriesPath:                 []string{"./assets/queries"},
		LibrariesPath:               "./assets/libraries",
		ReportFormats:               []string{"sarif"},
		Platform:                    []string{""},
		TerraformVarsPath:           "",
		QueryExecTimeout:            60,
		LineInfoPayload:             false,
		DisableSecrets:              true,
		SecretsRegexesPath:          "",
		ChangedDefaultQueryPath:     false,
		ChangedDefaultLibrariesPath: false,
		ScanID:                      "console",
		BillOfMaterials:             false,
		ExcludeGitIgnore:            false,
		OpenAPIResolveReferences:    false,
		ParallelScanFlag:            0,
		MaxFileSizeFlag:             5,
		UseOldSeverities:            false,
		MaxResolverDepth:            15,
		ExcludePlatform:             []string{""},
	}

	c := &scan.Client{
		ScanParams: scanParams,
	}

	_, err := c.GetQueryPath(ctx)
	if err != nil {
		logger.Err(err)
		return &engine.Inspector{}, err
	}

	queriesSource := source.NewFilesystemSource(
		ctx,
		c.ScanParams.QueriesPath,
		c.ScanParams.Platform,
		c.ScanParams.CloudProvider,
		c.ScanParams.LibrariesPath,
		c.ScanParams.ExperimentalQueries)

	includeQueries := source.IncludeQueries{
		ByIDs: []string{queryID},
	}

	queryFilter := source.QueryInspectorParameters{
		IncludeQueries: includeQueries,
	}

	t, err := tracker.NewTracker(c.ScanParams.PreviewLines)
	if err != nil {
		logger.Err(err)
		return &engine.Inspector{}, err
	}

	// bcgCtx := context.Background()

	logger.Info().Msgf("Preparing to inspect query source %v", queriesSource)

	inspector, err := engine.NewInspector(ctx,
		queriesSource,
		engine.DefaultVulnerabilityBuilder,
		t,
		&queryFilter,
		make(map[string]bool),
		c.ScanParams.QueryExecTimeout,
		c.ScanParams.UseOldSeverities,
		false,
		c.ScanParams.ParallelScanFlag,
		c.ScanParams.KicsComputeNewSimID,
		utils.FeatureFlags{},
	)

	return inspector, err
}

func loadQuery(ctx context.Context, inspector *engine.Inspector, queryID string) (*engine.PreparedQuery, error) {
	logger := logger.FromContext(ctx)
	if len(inspector.QueryLoader.QueriesMetadata) == 1 {
		queryOpa, err := inspector.QueryLoader.LoadQuery(context.Background(), &inspector.QueryLoader.QueriesMetadata[0], nil)

		if err != nil {
			logger.Err(err)
			return &engine.PreparedQuery{}, err
		}

		query := &engine.PreparedQuery{
			OpaQuery: *queryOpa,
			Metadata: inspector.QueryLoader.QueriesMetadata[0],
		}

		return query, nil
	}

	return &engine.PreparedQuery{}, errors.New("unable to load query" + queryID)
}

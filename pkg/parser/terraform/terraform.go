/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package terraform

import (
	"context"
	"os"
	"path/filepath"
	"regexp"

	"github.com/Checkmarx/kics/pkg/model"
	"github.com/Checkmarx/kics/pkg/parser/terraform/comment"
	"github.com/Checkmarx/kics/pkg/parser/terraform/converter"
	"github.com/Checkmarx/kics/pkg/parser/utils"
	masterUtils "github.com/Checkmarx/kics/pkg/utils"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

// RetriesDefaultValue is default number of times a parser will retry to execute
const RetriesDefaultValue = 50

// Converter returns content json, error line, error
type Converter func(ctx context.Context, file *hcl.File, inputVariables converter.VariableMap) (model.Document, error)

// Parser struct that contains the function to parse file and the number of retries if something goes wrong
type Parser struct {
	convertFunc       Converter
	numOfRetries      int
	terraformVarsPath string
	sciInfo           model.SCIInfo
	inputVariables    converter.VariableMap
}

// NewDefault initializes a parser with Parser default values
func NewDefault() *Parser {
	return &Parser{
		numOfRetries: RetriesDefaultValue,
		convertFunc:  converter.DefaultConverted,
	}
}

// NewDefaultWithVarsPath initializes a parser with the default values using a variables path
func NewDefaultWithVarsPath(terraformVarsPath string) *Parser {
	parser := NewDefault()
	parser.terraformVarsPath = terraformVarsPath
	return parser
}

func NewDefaultWithParams(terraformVarsPath string, sciInfo model.SCIInfo) *Parser {
	parser := NewDefault()
	parser.terraformVarsPath = terraformVarsPath
	parser.sciInfo = sciInfo
	return parser
}

// Resolve - replace or modifies in-memory content before parsing
func (p *Parser) Resolve(ctx context.Context, fileContent []byte, filename string, _ bool, _ int) ([]byte, error) {
	// handle panic during resolve process
	defer func() {
		if r := recover(); r != nil {
			errMessage := "Recovered from panic during resolve of file " + filename
			masterUtils.HandlePanic(ctx, r, errMessage)
		}
	}()
	inputVars := getInputVariables(ctx, filepath.Dir(filename), string(fileContent), p.terraformVarsPath)
	p.inputVariables = getDataSourcePolicy(ctx, filepath.Dir(filename), inputVars)
	return fileContent, nil
}

func processContent(ctx context.Context, elements model.Document, content, path string) {
	var certInfo map[string]interface{}
	if content != "" {
		certInfo = utils.AddCertificateInfo(ctx, path, content)
		if certInfo != nil {
			elements["certificate_body"] = certInfo
		}
	}
}

func processElements(ctx context.Context, elements model.Document, path string) {
	for k, v3 := range elements { // resource elements
		if k != "certificate_body" {
			continue
		}
		switch value := v3.(type) {
		case string:
			content := utils.CheckCertificate(value)
			processContent(ctx, elements, content, path)
		case ctyjson.SimpleJSONValue:
			content := utils.CheckCertificate(value.Value.AsString())
			processContent(ctx, elements, content, path)
		}
	}
}

func processResourcesElements(ctx context.Context, resourcesElements model.Document, path string) error {
	for _, v2 := range resourcesElements {
		switch t := v2.(type) {
		case []interface{}:
			return errors.New("failed to process resources")
		case interface{}:
			if elements, ok := t.(model.Document); ok {
				processElements(ctx, elements, path)
			}
		}
	}
	return nil
}

func processResources(ctx context.Context, doc model.Document, path string) error {
	var resourcesElements model.Document

	defer func() {
		if r := recover(); r != nil {
			errMessage := "Recovered from panic during process of resources in file " + path
			masterUtils.HandlePanic(ctx, r, errMessage)
		}
	}()

	for _, resources := range doc {
		switch t := resources.(type) {
		case []interface{}: // support the case of nameless resources - where we get a list of resources
			for _, value := range t {
				resourcesElements = value.(model.Document)
				err := processResourcesElements(ctx, resourcesElements, path)
				if err != nil {
					return err
				}
			}

		case interface{}:
			resourcesElements = t.(model.Document)
			err := processResourcesElements(ctx, resourcesElements, path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func addExtraInfo(ctx context.Context, json []model.Document, path string) ([]model.Document, error) {
	// handle panic during resource processing
	defer func() {
		if r := recover(); r != nil {
			errMessage := "Recovered from panic during resource processing for file " + path
			masterUtils.HandlePanic(ctx, r, errMessage)
		}
	}()
	for _, documents := range json { // iterate over documents
		if resources, ok := documents["resource"].(model.Document); ok {
			err := processResources(ctx, resources, path)
			if err != nil {
				return []model.Document{}, err
			}
		}
	}

	return json, nil
}

func parseFile(filename string, shouldReplaceDataSource bool) (*hcl.File, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if shouldReplaceDataSource {
		replaceDataIdentifiers := regexp.MustCompile(`(data\.[A-Za-z0-9._-]+)`)
		file = []byte(replaceDataIdentifiers.ReplaceAllString(string(file), "\"$1\""))
	}
	parsedFile, _ := hclsyntax.ParseConfig(file, filename, hcl.Pos{Line: 1, Column: 1})

	return parsedFile, nil
}

// Parse execute parser for the content in a file
func (p *Parser) Parse(ctx context.Context, path string, content []byte) ([]model.Document, []int, error) {
	logger := log.Ctx(ctx)
	file, diagnostics := hclsyntax.ParseConfig(content, filepath.Base(path), hcl.Pos{Byte: 0, Line: 1, Column: 1})
	defer func() {
		if r := recover(); r != nil {
			errMessage := "Recovered from panic during parsing of file " + path
			masterUtils.HandlePanic(ctx, r, errMessage)
		}
	}()
	if diagnostics != nil && diagnostics.HasErrors() && len(diagnostics.Errs()) > 0 {
		err := diagnostics.Errs()[0]
		return nil, []int{}, err
	}

	ignore, err := comment.ParseComments(content, path)
	if err != nil {
		logger.Err(err).Msg("failed to parse comments")
	}

	logger.Info().Int64(
		"org", p.sciInfo.OrgId,
	).Str(
		"branch", p.sciInfo.RepositoryCommitInfo.Branch,
	).Str(
		"sha", p.sciInfo.RepositoryCommitInfo.CommitSHA,
	).Str(
		"repository", p.sciInfo.RepositoryCommitInfo.RepositoryUrl,
	).Str(
		"exclusion_source", "inline",
	).Int(
		"lines_excluded", len(ignore[model.IgnoreLine]),
	).Int("blocks_excluded", len(ignore[model.IgnoreBlock])).Msg("Exclusions Info")

	linesToIgnore := comment.GetIgnoreLines(ignore, file.Body.(*hclsyntax.Body))

	fc, parseErr := p.convertFunc(ctx, file, p.inputVariables)
	json, err := addExtraInfo(ctx, []model.Document{fc}, path)
	if err != nil {
		return json, []int{}, errors.Wrap(err, "failed terraform parse")
	}

	return json, linesToIgnore, errors.Wrap(parseErr, "failed terraform parse")
}

// SupportedExtensions returns Terraform extensions
func (p *Parser) SupportedExtensions() []string {
	return []string{".tf", ".tfvars"}
}

// SupportedTypes returns types supported by this parser, which are terraform
func (p *Parser) SupportedTypes() map[string]bool {
	return map[string]bool{"terraform": true}
}

// GetKind returns Terraform kind parser
func (p *Parser) GetKind() model.FileKind {
	return model.KindTerraform
}

// GetCommentToken return the comment token of Terraform - #
func (p *Parser) GetCommentToken() string {
	return "#"
}

// StringifyContent converts original content into string formatted version
func (p *Parser) StringifyContent(content []byte) (string, error) {
	return string(content), nil
}

// GetResolvedFiles returns the files that are resolved
func (p *Parser) GetResolvedFiles() map[string]model.ResolvedFile {
	return make(map[string]model.ResolvedFile)
}

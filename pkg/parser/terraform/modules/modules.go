package tfmodules

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/Checkmarx/kics/pkg/model"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/rs/zerolog/log"
	"github.com/zclconf/go-cty/cty"
)

type ParsedModule struct {
	Name           string
	AbsSource      string
	Source         string
	Version        string
	IsLocal        bool
	SourceType     string // local, git, registry, etc.
	RegistryScope  string // public, private, or "" (non-registry)
	AttributesData map[string]ModuleAttributesInfo
}

type ModuleParseResult struct {
	Module ParsedModule
	Error  error
}

type ModuleAttributesInfo struct {
	Resources []string          `json:"resources"`
	Inputs    map[string]string `json:"inputs"`
}

var registryPattern = regexp.MustCompile(`^[a-z0-9\-]+/[a-z0-9\-]+/[a-z0-9\-]+$`)

func isValidRegistryFormat(s string) bool {
	return registryPattern.MatchString(s)
}

func resolveModulePath(source string, rootDir string) string {
	clean := strings.TrimPrefix(source, "file://")
	clean = strings.TrimPrefix(clean, "git::")
	return filepath.Clean(filepath.Join(rootDir, clean))
}

// ParseTerraformModules parses HCL content and extracts module source/version, resolving locals/variables if possible.
func ParseTerraformModules(files model.FileMetadatas) (map[string]ParsedModule, error) {
	modules := make(map[string]ParsedModule)
	localsMap := make(map[string]string)
	varsMap := make(map[string]string)

	for _, file := range files {
		filePath := file.FilePath
		baseDir := filepath.Dir(filePath)

		file.Content = getFileContent(file)

		hclFile, diags := hclsyntax.ParseConfig([]byte(file.Content), filePath, hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			log.Warn().Msgf("Skipping file %s due to HCL parse errors: %s", filePath, diags.Error())
			continue
		}

		body, ok := hclFile.Body.(*hclsyntax.Body)
		if !ok {
			log.Error().Msgf("Unexpected body type in %s", filePath)
			continue
		}

		// Collect locals and variable defaults
		for _, block := range body.Blocks {
			switch block.Type {
			case "locals":
				for name, attr := range block.Body.Attributes {
					val, diag := attr.Expr.Value(nil)
					if !diag.HasErrors() && val.Type().Equals(cty.String) {
						localsMap[name] = val.AsString()
					}
				}
			case "variable":
				if len(block.Labels) != 1 {
					continue
				}
				varName := block.Labels[0]
				if defAttr, ok := block.Body.Attributes["default"]; ok {
					val, diag := defAttr.Expr.Value(nil)
					if !diag.HasErrors() && val.Type().Equals(cty.String) {
						varsMap[varName] = val.AsString()
					}
				}
			}
		}

		// Extract module blocks
		for _, block := range body.Blocks {
			if block.Type != "module" || len(block.Labels) == 0 {
				continue
			}

			mod := ParsedModule{Name: block.Labels[0]}

			for key, attr := range block.Body.Attributes {
				resolved := resolveExpr(attr.Expr, localsMap, varsMap)

				switch key {
				case "source":
					mod.Source = resolved
					mod.SourceType, mod.RegistryScope = DetectModuleSourceType(resolved)
					mod.IsLocal = LooksLikeLocalModuleSource(strings.TrimPrefix(resolved, "git::"))

					if mod.IsLocal {
						// Normalize relative path to absolute
						absPath := filepath.Join(baseDir, strings.TrimPrefix(resolved, "file://"))
						mod.AbsSource = filepath.Clean(absPath)
						err := validateModuleSource(mod.AbsSource)
						if err != nil {
							log.Warn().Msgf("Invalid local module source %q: %v", mod.Source, err)
							continue
						}
					}

				case "version":
					mod.Version = resolved
				}
			}

			if _, exists := modules[mod.Source]; !exists {
				modules[mod.Source] = mod
			}
		}
	}

	return modules, nil
}

func validateModuleSource(absPath string) error {
	// Attempt to read the directory contents
	entries, err := os.ReadDir(absPath)
	if err != nil {
		log.Error().Msgf("Module source path %q is not accessible: %v", absPath, err)
		return fmt.Errorf("module source path %q is not accessible: %w", absPath, err)
	}

	// Check for at least one .tf file
	valid := false
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".tf") {
			valid = true
			break
		}
	}

	if !valid {
		log.Warn().Msgf("Module at %s does not contain any .tf files", absPath)
		return fmt.Errorf("module at %s does not contain any .tf files", absPath)
	}
	return nil
}

func getFileContent(file model.FileMetadata) string {
	var builder strings.Builder
	for _, line := range *file.LinesOriginalData {
		builder.WriteString(line)
		builder.WriteString("\n")
	}
	return builder.String()
}

// resolveExpr evaluates HCL expressions using known locals and vars
func resolveExpr(expr hclsyntax.Expression, locals map[string]string, vars map[string]string) string {
	switch e := expr.(type) {
	case *hclsyntax.LiteralValueExpr:
		if e.Val.Type().Equals(cty.String) {
			return e.Val.AsString()
		}
		return "__NON_STRING_LITERAL__"

	case *hclsyntax.TemplateExpr:
		var result strings.Builder
		for _, part := range e.Parts {
			switch p := part.(type) {
			case *hclsyntax.LiteralValueExpr:
				if p.Val.Type().Equals(cty.String) {
					result.WriteString(p.Val.AsString())
				}
			case *hclsyntax.ScopeTraversalExpr:
				result.WriteString(resolveScopeTraversal(p, locals, vars))
			default:
				result.WriteString("${UNSUPPORTED_TEMPLATE_EXPR}")
			}
		}
		return result.String()

	case *hclsyntax.ScopeTraversalExpr:
		return resolveScopeTraversal(e, locals, vars)

	case *hclsyntax.FunctionCallExpr:
		return resolveFunctionCall(e, locals, vars)

	default:
		val, diag := expr.Value(nil)
		if !diag.HasErrors() && val.Type().Equals(cty.String) {
			return val.AsString()
		}
		return "__UNRESOLVED__"
	}
}

func resolveScopeTraversal(expr *hclsyntax.ScopeTraversalExpr, locals map[string]string, vars map[string]string) string {
	traversal := expr.Traversal
	if len(traversal) < 2 {
		return "__INVALID_TRAVERSAL__"
	}

	root := traversal[0].(hcl.TraverseRoot).Name

	switch root {
	case "local":
		if attr, ok := traversal[1].(hcl.TraverseAttr); ok {
			if val, ok := locals[attr.Name]; ok {
				return val
			}
		}
	case "var":
		if attr, ok := traversal[1].(hcl.TraverseAttr); ok {
			if val, ok := vars[attr.Name]; ok {
				return val
			}
		}
	case "data":
		// Convert traversal to something like: data_ref:aws_s3_bucket.logs.bucket_domain_name
		parts := []string{}
		for _, step := range traversal[1:] {
			switch s := step.(type) {
			case hcl.TraverseAttr:
				parts = append(parts, s.Name)
			default:
				parts = append(parts, "__UNKNOWN__")
			}
		}
		return "data_ref:" + strings.Join(parts, ".")
	}

	return "__UNKNOWN_REF__"
}

func resolveFunctionCall(expr *hclsyntax.FunctionCallExpr, locals map[string]string, vars map[string]string) string {
	switch expr.Name {
	case "format":
		if len(expr.Args) < 1 {
			return "__INVALID_FORMAT__"
		}
		formatStr := resolveExpr(expr.Args[0], locals, vars)
		args := make([]interface{}, 0, len(expr.Args)-1)
		for _, arg := range expr.Args[1:] {
			args = append(args, resolveExpr(arg, locals, vars))
		}
		return fmt.Sprintf(formatStr, args...)

	case "join":
		if len(expr.Args) != 2 {
			return "__INVALID_JOIN__"
		}
		sep := resolveExpr(expr.Args[0], locals, vars)
		listExpr, ok := expr.Args[1].(*hclsyntax.TupleConsExpr)
		if !ok {
			return "__INVALID_JOIN_LIST__"
		}
		items := []string{}
		for _, item := range listExpr.Exprs {
			items = append(items, resolveExpr(item, locals, vars))
		}
		return strings.Join(items, sep)

	default:
		return fmt.Sprintf("__UNSUPPORTED_FUNC_%s__", expr.Name)
	}
}

// LooksLikeLocalModuleSource uses heuristics to determine if the resolved source string is likely local
func LooksLikeLocalModuleSource(source string) bool {
	source = strings.TrimSpace(source)

	if source == "" {
		return false
	}

	// Unwrap common go-getter schemes like git:: or hg::
	schemes := []string{"git::", "hg::", "file::", "http::", "https::"}
	for _, scheme := range schemes {
		if strings.HasPrefix(source, scheme) {
			source = strings.TrimPrefix(source, scheme)
			break
		}
	}

	// Lowercase for uniformity when checking prefixes
	lower := strings.ToLower(source)

	return strings.HasPrefix(lower, "./") ||
		strings.HasPrefix(lower, "../") ||
		filepath.IsAbs(source)
}

func DetectModuleSourceType(source string) (string, string) {
	source = strings.TrimSpace(source)

	if source == "" {
		return "unknown", ""
	}

	if strings.HasPrefix(source, "data_ref:") {
		return "data_ref", ""
	}

	// Recognize git-based sources
	if strings.HasPrefix(source, "git::") {
		return "git", ""
	}

	// Recognize public registry hostname
	if strings.HasPrefix(source, "registry.terraform.io/") {
		return "registry", "public"
	}

	// Recognize private registries by fully qualified domain with 3 parts
	if strings.Count(source, "/") == 3 && strings.Contains(source, ".") {
		return "registry", "private"
	}

	// Recognize implicit public registry format (namespace/name/provider)
	if isValidRegistryFormat(source) {
		return "registry", "public"
	}

	if LooksLikeLocalModuleSource(source) {
		return "local", ""
	}

	return "unknown", ""
}

func ParseAllModuleVariables(modules map[string]ParsedModule, rootDir string) []ParsedModule {
	numWorkers := 4

	input := make(chan ParsedModule)
	output := make(chan ModuleParseResult)

	var wg sync.WaitGroup

	// Fan-out: Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for mod := range input {
				if !mod.IsLocal {
					output <- ModuleParseResult{Module: mod}
					continue
				}
				modulePath := resolveModulePath(mod.AbsSource, rootDir)

				attributesData, err := generateEquivalentMap(modulePath)
				if err != nil {
					log.Warn().Msg("Failed to generate equivalent map")
				} else {
					mod.AttributesData = attributesData
				}
				output <- ModuleParseResult{Module: mod, Error: err}
			}
		}()
	}

	// Fan-in: Close output when all workers are done
	go func() {
		wg.Wait()
		close(output)
	}()

	// Feed input channel
	go func() {
		for _, mod := range modules {
			input <- mod
		}
		close(input)
	}()

	// Collect results
	finalModules := make([]ParsedModule, 0, len(modules))
	for res := range output {
		if res.Error != nil {
			log.Warn().Msgf("Failed to parse module %s: %v", res.Module.Name, res.Error)
		}
		finalModules = append(finalModules, res.Module)
	}

	return finalModules
}

func generateEquivalentMap(modulePath string) (map[string]ModuleAttributesInfo, error) {
	equivalentMap := make(map[string]ModuleAttributesInfo)
	resourceTypesMap := make(map[string]map[string]bool)

	entries, err := os.ReadDir(modulePath)
	if err != nil {
		log.Error().Msgf("Failed to read module source directory: %s", modulePath)
		return nil, err
	}

	for _, entry := range entries {
		path := filepath.Join(modulePath, entry.Name())

		if entry.IsDir() {
			log.Debug().Msgf("Skipping directory: %s", path)
			continue
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			log.Error().Msgf("Failed to read file: %s", path)
			return nil, err
		}

		hclFile, diag := hclwrite.ParseConfig(contents, "", hcl.InitialPos)
		if diag.HasErrors() {
			return nil, fmt.Errorf("error parsing input Terraform block in file %s: %s", path, diag.Error())
		}

		for _, block := range hclFile.Body().Blocks() {
			if block.Type() != "resource" {
				continue
			}

			if len(block.Labels()) < 1 {
				log.Warn().Msgf("Skipping malformed resource block with no labels in file %s", path)
				continue
			}

			resourceType := block.Labels()[0]
			provider, err := GetProviderFromResourceType(resourceType)
			if err != nil {
				log.Warn().Msgf("Failed to get provider from resource type '%s' in file %s: %v", resourceType, path, err)
				continue
			}

			// Store resource type to the set
			if _, ok := resourceTypesMap[provider]; !ok {
				resourceTypesMap[provider] = make(map[string]bool)
			}
			resourceTypesMap[provider][resourceType] = true

			// Create or update the module info object for current provider
			modInfo, ok := equivalentMap[provider]
			if !ok {
				modInfo = ModuleAttributesInfo{
					Resources: []string{},
					Inputs:    make(map[string]string),
				}
			}

			// Update inputs mapping with all attributes referencing a variable
			maps.Copy(modInfo.Inputs, getVariableAttributes(block))

			// Assign the updated modInfo back to the map
			equivalentMap[provider] = modInfo
		}
	}

	// After iterating through all files and blocks, populate the unique resources slice
	for provider, typesSet := range resourceTypesMap {
		modInfo := equivalentMap[provider]
		for rt := range typesSet {
			modInfo.Resources = append(modInfo.Resources, rt)
		}
		equivalentMap[provider] = modInfo
	}

	return equivalentMap, nil
}

func getVariableAttributes(block *hclwrite.Block) map[string]string {
	attributeToVariableMap := make(map[string]string)
	for name, attr := range block.Body().Attributes() {
		value := string(attr.Expr().BuildTokens(nil).Bytes())
		if !isVariableReference(value) {
			continue
		}

		if varName := parseVariableReference(value); varName != "" {
			attributeToVariableMap[name] = varName
		}
	}

	// Handle nested blocks too
	for _, nestedBlock := range block.Body().Blocks() {
		maps.Copy(attributeToVariableMap, getVariableAttributes(nestedBlock))
	}
	return attributeToVariableMap
}

func isVariableReference(s string) bool {
	return strings.Contains(s, "var.")
}

var reVarRef = regexp.MustCompile(`^var\.(\w+)$`)

func parseVariableReference(s string) string {
	match := reVarRef.FindStringSubmatch(strings.TrimSpace(s))
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

// GetProviderFromResourceType extracts the provider name from a Terraform resource type.
// For example: "aws_s3_bucket" → "aws", "azurerm_network_interface" → "azurerm"
func GetProviderFromResourceType(resourceType string) (string, error) {
	if resourceType == "" {
		return "", errors.New("resource type cannot be empty")
	}
	parts := strings.SplitN(resourceType, "_", 2)
	if len(parts) < 2 {
		return "", errors.New("invalid Terraform resource type format")
	}
	return parts[0], nil
}

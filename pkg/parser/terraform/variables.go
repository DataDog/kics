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
	"strings"

	"github.com/Checkmarx/kics/pkg/logger"
	"github.com/Checkmarx/kics/pkg/parser/terraform/converter"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

func mergeMaps(baseMap, newItems converter.VariableMap) {
	for key, value := range newItems {
		baseMap[key] = value
	}
}

func setInputVariablesDefaultValues(filename string) (converter.VariableMap, error) {
	parsedFile, err := parseFile(filename, false)
	if err != nil || parsedFile == nil {
		return nil, err
	}
	content, _, _ := parsedFile.Body.PartialContent(&hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "variable",
				LabelNames: []string{"name"},
			},
		},
	})

	defaultValuesMap := make(converter.VariableMap)
	for _, block := range content.Blocks {
		if len(block.Labels) == 0 || block.Labels[0] == "" {
			continue
		}
		attr, _ := block.Body.JustAttributes()
		if len(attr) == 0 {
			continue
		}
		if defaultValue, exists := attr["default"]; exists {
			defaultVar, _ := defaultValue.Expr.Value(nil)
			defaultValuesMap[block.Labels[0]] = defaultVar
		}
	}
	return defaultValuesMap, nil
}

func checkTfvarsValid(ctx context.Context, f *hcl.File, filename string) error {
	logger := logger.FromContext(ctx)
	content, _, _ := f.Body.PartialContent(&hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "variable",
				LabelNames: []string{"name"},
			},
		},
	})
	if len(content.Blocks) > 0 {
		logger.Debug().Msgf("failed to get variables from %s, .tfvars file is used to assing values not to declare new variables", filename)
	}
	return nil
}

func getInputVariablesFromFile(filename string) (converter.VariableMap, error) {
	src, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	parsedFile, diags := hclsyntax.ParseConfig(src, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return nil, diags
	}

	variables := make(converter.VariableMap)

	body, ok := parsedFile.Body.(*hclsyntax.Body)
	if !ok {
		return variables, nil
	}

	// Case 1: .tfvars-style assignments
	if filepath.Ext(filename) == ".tfvars" || strings.HasSuffix(filename, ".auto.tfvars") {
		for name, attr := range body.Attributes {
			val, _ := attr.Expr.Value(&hcl.EvalContext{})
			variables[name] = val
		}
		return variables, nil
	}

	// Case 2: .tf variable "x" { default = ... }
	for _, block := range body.Blocks {
		if block.Type == "variable" && len(block.Labels) == 1 {
			varName := block.Labels[0]
			if defaultAttr, exists := block.Body.Attributes["default"]; exists {
				val, diags := defaultAttr.Expr.Value(&hcl.EvalContext{})
				if !diags.HasErrors() && val.IsKnown() {
					variables[varName] = val
				}
			}
		}
	}

	return variables, nil
}

func getInputLocalsFromFile(filename string) (converter.VariableMap, error) {
	src, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	parsedFile, diags := hclsyntax.ParseConfig(src, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return nil, diags
	}

	locals := make(converter.VariableMap)
	type unresolved struct {
		name string
		expr hclsyntax.Expression
	}
	var unevaluated []unresolved

	// Collect all locals
	for _, block := range parsedFile.Body.(*hclsyntax.Body).Blocks {
		if block.Type != "locals" {
			continue
		}
		for name, attr := range block.Body.Attributes {
			unevaluated = append(unevaluated, unresolved{name, attr.Expr})
		}
	}

	// Resolve locals in multiple passes
	for pass := 0; pass < 3; pass++ {
		for _, item := range unevaluated {
			if _, ok := locals[item.name]; ok {
				continue // already resolved
			}
			val, diags := item.expr.Value(&hcl.EvalContext{
				Variables: map[string]cty.Value{
					"local": cty.ObjectVal(locals),
				},
			})
			if !diags.HasErrors() && val.IsKnown() {
				locals[item.name] = val
			}
		}
	}

	return locals, nil
}

func sanitizeCtyMap(in map[string]cty.Value) map[string]cty.Value {
	out := make(map[string]cty.Value)
	for k, v := range in {
		if !v.IsKnown() || v.IsNull() {
			// default to empty string or another valid fallback
			out[k] = cty.NullVal(v.Type())
			continue
		}
		out[k] = v
	}
	return out
}

func getInputVariables(ctx context.Context, currentPath, fileContent, terraformVarsPath string) converter.VariableMap {
	logger := logger.FromContext(ctx)
	variablesMap := make(converter.VariableMap)
	localsMap := make(converter.VariableMap)

	tfFiles, err := filepath.Glob(filepath.Join(currentPath, "*.tf"))
	if err != nil {
		logger.Error().Msg("Error getting .tf files")
	}

	// Parse all .tf files for variables and locals
	for _, tfFile := range tfFiles {
		// Get variables
		vars, errVars := getInputVariablesFromFile(tfFile)
		if errVars != nil {
			logger.Error().Msgf("Error getting default values from %s", tfFile)
			logger.Err(errVars)
		} else {
			mergeMaps(variablesMap, vars)
		}

		// Get locals
		locals, errLocals := getInputLocalsFromFile(tfFile)
		if errLocals != nil {
			logger.Error().Msgf("Error getting locals from %s", tfFile)
			logger.Err(errLocals)
		} else {
			mergeMaps(localsMap, locals)
		}
	}

	// Parse *.auto.tfvars files
	tfVarsFiles, err := filepath.Glob(filepath.Join(currentPath, "*.auto.tfvars"))
	if err != nil {
		logger.Error().Msg("Error getting .auto.tfvars files")
	}

	// Add terraform.tfvars if it exists
	if _, err := os.Stat(filepath.Join(currentPath, "terraform.tfvars")); err == nil {
		tfVarsFiles = append(tfVarsFiles, filepath.Join(currentPath, "terraform.tfvars"))
	}

	for _, tfVarsFile := range tfVarsFiles {
		vars, errInput := getInputVariablesFromFile(tfVarsFile)
		if errInput != nil {
			logger.Error().Msgf("Error getting values from %s", tfVarsFile)
			logger.Err(errInput)
			continue
		}
		mergeMaps(variablesMap, vars)
	}

	if terraformVarsPath == "" {
		terraformVarsPathRegex := regexp.MustCompile(`(?m)^\s*// kics_terraform_vars: ([\w/\\.:-]+)\r?\n`)
		match := terraformVarsPathRegex.FindStringSubmatch(fileContent)
		if match != nil {
			terraformVarsPath = match[1]
			if !strings.Contains(terraformVarsPath, ":") {
				terraformVarsPath = filepath.Join(currentPath, terraformVarsPath)
			}
		}
	}

	if terraformVarsPath != "" {
		if _, err := os.Stat(terraformVarsPath); err == nil {
			vars, errInput := getInputVariablesFromFile(terraformVarsPath)
			if errInput != nil {
				logger.Error().Msgf("Error getting values from %s", terraformVarsPath)
				logger.Err(errInput)
			} else {
				mergeMaps(variablesMap, vars)
			}
		} else {
			logger.Trace().Msgf("%s file not found", terraformVarsPath)
		}
	}

	// Final sanitized maps
	cleanVars := sanitizeCtyMap(variablesMap)
	cleanLocals := sanitizeCtyMap(localsMap)

	result := make(converter.VariableMap)
	result["var"] = cty.ObjectVal(cleanVars)
	result["local"] = cty.ObjectVal(cleanLocals)
	return result
}

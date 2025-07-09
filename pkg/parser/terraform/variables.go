/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package terraform

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Checkmarx/kics/pkg/parser/terraform/converter"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/rs/zerolog/log"
	"github.com/zclconf/go-cty/cty"
)

var inputVariableMap = make(converter.VariableMap)

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

func checkTfvarsValid(f *hcl.File, filename string) error {
	content, _, _ := f.Body.PartialContent(&hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "variable",
				LabelNames: []string{"name"},
			},
		},
	})
	if len(content.Blocks) > 0 {
		log.Debug().Msgf("failed to get variables from %s, .tfvars file is used to assing values not to declare new variables", filename)
	}
	return nil
}

func getInputVariablesFromFile(filename string) (converter.VariableMap, error) {
	parsedFile, err := parseFile(filename, false)
	if err != nil || parsedFile == nil {
		return nil, err
	}
	// err = checkTfvarsValid(parsedFile, filename)
	// if err != nil {
	// 	return nil, err
	// }
	variables := make(converter.VariableMap)

	attrs := parsedFile.Body.(*hclsyntax.Body).Attributes
	for name, attr := range attrs {
		value, _ := attr.Expr.Value(&hcl.EvalContext{})
		variables[name] = value
	}

	for _, block := range parsedFile.Body.(*hclsyntax.Body).Blocks {
		if block.Type == "locals" {
			// If the block is a locals block, we need to get the attributes inside it
			block_attrs := block.Body.Attributes
			for name, block_attr := range block_attrs {
				value, _ := block_attr.Expr.Value(&hcl.EvalContext{})
				variables[name] = value
			}
		}
	}
	return variables, nil
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

func getInputVariables(currentPath, fileContent, terraformVarsPath string) {
	variablesMap := make(converter.VariableMap)
	localsMap := make(converter.VariableMap)
	tfFiles, err := filepath.Glob(filepath.Join(currentPath, "*.tf"))
	if err != nil {
		log.Error().Msg("Error getting .tf files")
	}
	for _, tfFile := range tfFiles {
		variables, errDefaultValues := setInputVariablesDefaultValues(tfFile)
		if errDefaultValues != nil {
			log.Error().Msgf("Error getting default values from %s", tfFile)
			log.Err(errDefaultValues)
			continue
		}
		mergeMaps(variablesMap, variables)
	}
	tfVarsFiles, err := filepath.Glob(filepath.Join(currentPath, "*.auto.tfvars"))
	if err != nil {
		log.Error().Msg("Error getting .auto.tfvars files")
	}

	_, err = os.Stat(filepath.Join(currentPath, "terraform.tfvars"))
	if err == nil {
		tfVarsFiles = append(tfVarsFiles, filepath.Join(currentPath, "terraform.tfvars"))
	}

	for _, tfVarsFile := range tfVarsFiles {
		variables, errInputVariables := getInputVariablesFromFile(tfVarsFile)
		if errInputVariables != nil {
			log.Error().Msgf("Error getting values from %s", tfVarsFile)
			log.Err(errInputVariables)
			continue
		}
		mergeMaps(variablesMap, variables)
	}

	// If the flag is empty let's look for the value in the first written line of the file
	if terraformVarsPath == "" {
		terraformVarsPathRegex := regexp.MustCompile(`(?m)^\s*// kics_terraform_vars: ([\w/\\.:-]+)\r?\n`)
		terraformVarsPathMatch := terraformVarsPathRegex.FindStringSubmatch(fileContent)
		if terraformVarsPathMatch != nil {
			// There is a path tp the variables file in the file so that will be the path to the variables tf file
			terraformVarsPath = terraformVarsPathMatch[1]
			// If the path contains ":" assume its a global path
			if !strings.Contains(terraformVarsPath, ":") {
				// If not then add the current folder path before so that the comment path can be relative
				terraformVarsPath = filepath.Join(currentPath, terraformVarsPath)
			}
		}
	}

	// If the terraformVarsPath is empty, this means that it is not in the flag
	// and it is not in the first written line of the file
	if terraformVarsPath != "" {
		_, err = os.Stat(terraformVarsPath)
		if err != nil {
			log.Trace().Msgf("%s file not found", terraformVarsPath)
		} else {
			variables, errInputVariables := getInputVariablesFromFile(terraformVarsPath)
			if errInputVariables != nil {
				log.Error().Msgf("Error getting values from %s", terraformVarsPath)
				log.Err(errInputVariables)
			} else {
				mergeMaps(variablesMap, variables)
			}
		}
	}

	// check if variables.tf or variable.tf file exists in the current path
	var foundVarFile string
	for _, fname := range []string{"variables.tf", "variable.tf"} {
		path := filepath.Join(currentPath, fname)
		if _, err := os.Stat(path); err == nil {
			foundVarFile = path
			break
		}
	}

	if foundVarFile != "" {
		variables, errInputVariables := getInputVariablesFromFile(foundVarFile)
		if errInputVariables != nil {
			log.Error().Msgf("Error getting values from %s: %v", foundVarFile, errInputVariables)
			log.Err(errInputVariables)
		} else {
			mergeMaps(variablesMap, variables)
		}
	}
	cleanVars := sanitizeCtyMap(variablesMap)
	inputVariableMap["var"] = cty.ObjectVal(cleanVars)

	// check if @locals.tf or locals.tf file exists in the current path
	for _, fname := range []string{"@locals.tf", "locals.tf"} {
		path := filepath.Join(currentPath, fname)
		if _, err := os.Stat(path); err == nil {
			foundVarFile = path
			break
		}
	}

	if foundVarFile != "" {
		variables, errInputVariables := getInputVariablesFromFile(foundVarFile)
		if errInputVariables != nil {
			log.Error().Msgf("Error getting values from %s: %v", foundVarFile, errInputVariables)
			log.Err(errInputVariables)
		} else {
			mergeMaps(localsMap, variables)
		}
	}
	cleanLocals := sanitizeCtyMap(localsMap)
	inputVariableMap["local"] = cty.ObjectVal(cleanLocals)

}

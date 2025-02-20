package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Metadata struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Provider string `json:"provider"`
}

// Function to get user input with a default value
func getInput(prompt, defaultValue string) string {
	fmt.Printf("%s (default: %s): ", prompt, defaultValue)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}
	return input
}

func main() {
	// Set default values
	defaultRepoPath := "/Users/bahar.shah/dev/kics"
	defaultPlatform := "terraform"
	defaultProvider := "aws"

	// Get user input with default values
	repoPath := getInput("Enter the path to the repo", defaultRepoPath)
	platform := getInput("Enter the platform", defaultPlatform)
	provider := getInput("Enter the provider", defaultProvider)
	ruleName := getInput("Enter the rule name", "")
	if ruleName == "" {
		fmt.Println("Rule name is required. Exiting.")
		return
	}

	var numNegatives int
	fmt.Print("Enter the number of negative examples (default: 0): ")
	fmt.Scanln(&numNegatives)

	// Construct the base directory path
	baseDir := filepath.Join(repoPath, "assets", "queries", platform, provider, ruleName)
	testDir := filepath.Join(baseDir, "test")

	// Create directories
	dirs := []string{baseDir, testDir}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Failed to create directory %s: %v\n", dir, err)
			return
		}
	}

	// Create query.rego file
	queryFilePath := filepath.Join(baseDir, "query.rego")
	if err := os.WriteFile(queryFilePath, []byte("package "+ruleName+"\n\n# Define your Rego policy here"), 0644); err != nil {
		fmt.Printf("Failed to create file %s: %v\n", queryFilePath, err)
		return
	}

	// Create metadata.json file
	metadata := Metadata{
		Name:     ruleName,
		Platform: platform,
		Provider: provider,
	}
	metadataFilePath := filepath.Join(baseDir, "metadata.json")
	metadataJSON, _ := json.MarshalIndent(metadata, "", "  ")
	if err := os.WriteFile(metadataFilePath, metadataJSON, 0644); err != nil {
		fmt.Printf("Failed to create file %s: %v\n", metadataFilePath, err)
		return
	}

	// Create test files
	positiveTFPath := filepath.Join(testDir, "positive.tf")
	positiveExpectedPath := filepath.Join(testDir, "positive_expected_result.json")
	os.WriteFile(positiveTFPath, []byte("# Positive test case"), 0644)
	os.WriteFile(positiveExpectedPath, []byte("{ \"result\": \"pass\" }"), 0644)

	// Create negative test cases
	if numNegatives == 1 {
		negativeTFPath := filepath.Join(testDir, "negative.tf")
		os.WriteFile(negativeTFPath, []byte("# Single negative test case"), 0644)
	} else {
		for i := 0; i < numNegatives; i++ {
			negativeTFPath := filepath.Join(testDir, "negative"+strconv.Itoa(i)+".tf")
			os.WriteFile(negativeTFPath, []byte("# Negative test case "+strconv.Itoa(i)), 0644)
		}
	}

	fmt.Println("Directory structure and files created successfully!")
}

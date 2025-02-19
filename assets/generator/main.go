package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Metadata struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Provider string `json:"provider"`
}

func main() {
	var repoPath, platform, provider, ruleName string
	var numNegatives int

	// Get user input
	fmt.Print("Enter the path to the repo: (should be the root of the local repository) ")
	fmt.Scanln(&repoPath)
	fmt.Print("Enter the platform: (ex: terraform)")
	fmt.Scanln(&platform)
	fmt.Print("Enter the provider: (ex: aws)")
	fmt.Scanln(&provider)
	fmt.Print("Enter the rule name: (ex: ec2_instance_no_public_ip)")
	fmt.Scanln(&ruleName)
	fmt.Print("Enter the number of negative examples: (minimum 1) ")
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
	os.WriteFile(positiveExpectedPath, []byte("[]"), 0644)

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

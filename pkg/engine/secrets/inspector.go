/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package secrets

import (
	"math"
	"regexp"
	"strings"
)

const (
	Base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
	HexChars    = "1234567890abcdefABCDEF"
)

// SecretTracker is Struct created to keep track of the secrets found in the inspector
// it used for masking all the secrets in the vulnerability preview in the different report formats
type SecretTracker struct {
	ResolvedFilePath string
	Line             int
	OriginalContent  string
	MaskedContent    string
}

type Entropy struct {
	Group int     `json:"group"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
}

type MultilineResult struct {
	DetectLineGroup int `json:"detectLineGroup"`
}

type AllowRule struct {
	Description string `json:"description"`
	RegexStr    string `json:"regex"`
	Regex       *regexp.Regexp
}

type RegexQuery struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Multiline   MultilineResult `json:"multiline"`
	RegexStr    string          `json:"regex"`
	SpecialMask string          `json:"specialMask"`
	Entropies   []Entropy       `json:"entropies"`
	AllowRules  []AllowRule     `json:"allowRules"`
	Regex       *regexp.Regexp
}

// IsAllowRule check if string matches any of the allow rules for the secret queries
func IsAllowRule(s string, query *RegexQuery, allowRules []AllowRule) bool {
	regexMatch := query.Regex.FindStringIndex(s)
	if regexMatch != nil {
		allowRuleMatches := AllowRuleMatches(s, append(query.AllowRules, allowRules...))

		for _, allowMatch := range allowRuleMatches {
			allowStart, allowEnd := allowMatch[0], allowMatch[1]
			regexStart, regexEnd := regexMatch[0], regexMatch[1]

			if (allowStart <= regexEnd && allowStart >= regexStart) || (regexStart <= allowEnd && regexStart >= allowStart) {
				return true
			}
		}
	}

	return false
}

// AllowRuleMatches return all the allow rules matches for the secret queries
func AllowRuleMatches(s string, allowRules []AllowRule) [][]int {
	allowRuleMatches := [][]int{}
	for i := range allowRules {
		allowRuleMatches = append(allowRuleMatches, allowRules[i].Regex.FindAllStringIndex(s, -1)...)
	}
	return allowRuleMatches
}

// CheckEntropyInterval - verifies if a given token's entropy is within expected bounds
func CheckEntropyInterval(entropy Entropy, token string) (isEntropyInInterval bool, entropyLevel float64) {
	base64Entropy := calculateEntropy(token, Base64Chars)
	hexEntropy := calculateEntropy(token, HexChars)
	highestEntropy := math.Max(base64Entropy, hexEntropy)
	if insideInterval(entropy, base64Entropy) || insideInterval(entropy, hexEntropy) {
		return true, highestEntropy
	}
	return false, highestEntropy
}

func insideInterval(entropy Entropy, floatEntropy float64) bool {
	return floatEntropy >= entropy.Min && floatEntropy <= entropy.Max
}

// calculateEntropy - calculates the entropy of a string based on the Shannon formula
func calculateEntropy(token, charSet string) float64 {
	if token == "" {
		return 0
	}
	charMap := map[rune]float64{}
	for _, char := range token {
		if strings.Contains(charSet, string(char)) {
			charMap[char]++
		}
	}

	var freq float64
	length := float64(len(token))
	for _, count := range charMap {
		freq += count * math.Log2(count)
	}

	return math.Log2(length) - freq/length
}

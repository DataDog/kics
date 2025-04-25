/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package model

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Checkmarx/kics/internal/constants"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var categoriesNotFound = make(map[string]bool)

var severityLevelEquivalence = map[model.Severity]string{
	"INFO":     "none",
	"LOW":      "note",
	"MEDIUM":   "warning",
	"HIGH":     "error",
	"CRITICAL": "error",
}

var targetTemplate = sarifDescriptorReference{
	ToolComponent: sarifComponentReference{
		ComponentReferenceGUID:  "58cdcc6f-fe41-4724-bfb3-131a93df4c3f",
		ComponentReferenceName:  "Categories",
		ComponentReferenceIndex: 0,
	},
}

var cweTemplate = cweDescriptorReference{
	ToolComponent: cweComponentReference{
		ComponentReferenceGUID: "1489b0c4-d7ce-4d31-af66-6382a01202e3",
		ComponentReferenceName: "CWE",
	},
}

type sarifProperties map[string]interface{}

type ruleMetadata struct {
	queryID          string
	queryName        string
	queryDescription string
	queryURI         string
	queryCategory    string
	queryCwe         string
	severity         model.Severity
}

type ruleCISMetadata struct {
	descriptionText string
	id              string
	title           string
}

type sarifMessage struct {
	Text              string          `json:"text"`
	MessageProperties sarifProperties `json:"properties,omitempty"`
}

type sarifComponentReference struct {
	ComponentReferenceName  string `json:"name,omitempty"`
	ComponentReferenceGUID  string `json:"guid,omitempty"`
	ComponentReferenceIndex int    `json:"index,omitempty"`
}

type cweComponentReference struct {
	ComponentReferenceGUID string `json:"guid"`
	ComponentReferenceName string `json:"name"`
}

type sarifDescriptorReference struct {
	ReferenceID    string                  `json:"id,omitempty"`
	ReferenceGUID  string                  `json:"guid,omitempty"`
	ReferenceIndex int                     `json:"index,omitempty"`
	ToolComponent  sarifComponentReference `json:"toolComponent,omitempty"`
}

type cweMessage struct {
	Text string `json:"text"`
}

type cweCsv struct {
	CweID            string     `json:"id"`
	FullDescription  cweMessage `json:"fullDescription"`
	ShortDescription cweMessage `json:"shortDescription"`
	GUID             string     `json:"guid"`
	HelpURI          string     `json:"helpUri"`
}

type cweDescriptorReference struct {
	ReferenceID   string                `json:"id"`
	ReferenceGUID string                `json:"guid"`
	ToolComponent cweComponentReference `json:"toolComponent"`
}

type sarifConfiguration struct {
	Level string `json:"level"`
}

type sarifRelationship struct {
	Relationship sarifDescriptorReference `json:"target,omitempty"`
}

type sarifRule struct {
	RuleID               string              `json:"id"`
	RuleName             string              `json:"name"`
	RuleShortDescription sarifMessage        `json:"shortDescription"`
	RuleFullDescription  sarifMessage        `json:"fullDescription"`
	DefaultConfiguration sarifConfiguration  `json:"defaultConfiguration"`
	HelpURI              string              `json:"helpUri"`
	Relationships        []sarifRelationship `json:"relationships,omitempty"`
	RuleProperties       sarifProperties     `json:"properties,omitempty"`
}

type sarifDriver struct {
	ToolName     string              `json:"name"`
	ToolVersion  string              `json:"version"`
	ToolFullName string              `json:"fullName"`
	ToolURI      string              `json:"informationUri"`
	Properties   sarifToolProperties `json:"properties"`
	Rules        []sarifRule         `json:"rules"`
}

type sarifToolProperties struct {
	Tags []string `json:"tags"`
}

type sarifTool struct {
	Driver sarifDriver `json:"driver"`
}

type sarifResourceLocation struct {
	Line int `json:"line"`
	Col  int `json:"col"`
}

type sarifRegion struct {
	StartLine   int `json:"startLine"`
	EndLine     int `json:"endLine"`
	StartColumn int `json:"startColumn"`
	EndColumn   int `json:"endColumn"`
	// StartResource sarifResourceLocation `json:"startResource"`
	// EndResource   sarifResourceLocation `json:"endResource"`
}

type sarifArtifactLocation struct {
	ArtifactURI string `json:"uri"`
}

type sarifPhysicalLocation struct {
	ArtifactLocation sarifArtifactLocation `json:"artifactLocation"`
	Region           sarifRegion           `json:"region"`
}

type sarifLocation struct {
	PhysicalLocation sarifPhysicalLocation `json:"physicalLocation"`
}

type sarifResult struct {
	ResultRuleID        string                   `json:"ruleId"`
	ResultRuleIndex     int                      `json:"ruleIndex"`
	ResultKind          string                   `json:"kind,omitempty"`
	ResultMessage       sarifMessage             `json:"message"`
	ResultLocations     []sarifLocation          `json:"locations"`
	PartialFingerprints SarifPartialFingerprints `json:"partialFingerprints,omitempty"`
	ResultLevel         string                   `json:"level"`
	ResultProperties    sarifProperties          `json:"properties,omitempty"`
	ResultFixes         []sarifFix               `json:"fixes,omitempty"`
}

type SarifPartialFingerprints struct {
	Sha                string `json:"SHA,omitempty"`
	DatadogFingerprint string `json:"DATADOG_FINGERPRINT,omitempty"`
	CommitSha          string `json:"commitSha,omitempty"`
	Email              string `json:"email,omitempty"`
	Author             string `json:"author,omitempty"`
	Date               string `json:"date,omitempty"`
	CommitMessage      string `json:"commitMessage,omitempty"`
}

type taxonomyDefinitions struct {
	DefinitionGUID             string     `json:"guid,omitempty"`
	DefinitionName             string     `json:"name,omitempty"`
	DefinitionID               string     `json:"id"`
	DefinitionShortDescription cweMessage `json:"shortDescription"`
	DefinitionFullDescription  cweMessage `json:"fullDescription"`
	HelpURI                    string     `json:"helpUri,omitempty"`
}

type sarifFix struct {
	ArtifactChanges []artifactChange `json:"artifactChanges"`
	Description     fixMessage       `json:"description"`
}

type artifactChange struct {
	ArtifactLocation artifactLocation `json:"artifactLocation"`
	Replacements     []fixReplacement `json:"replacements"`
}

type artifactLocation struct {
	URI string `json:"uri"`
}

type fixReplacement struct {
	DeletedRegion   sarifRegion `json:"deletedRegion"`
	InsertedContent fixContent  `json:"insertedContent,omitempty"`
}

type fixContent struct {
	Text string `json:"text"`
}

type fixMessage struct {
	Text string `json:"text"`
}

type cweTaxonomiesWrapper struct {
	Taxonomies sarifTaxonomy `json:"taxonomies"`
}

type sarifTaxonomy struct {
	TaxonomyGUID                              string                `json:"guid"`
	TaxonomyName                              string                `json:"name"`
	TaxonomyFullDescription                   sarifMessage          `json:"fullDescription,omitempty"`
	TaxonomyShortDescription                  sarifMessage          `json:"shortDescription"`
	TaxonomyDownloadURI                       string                `json:"downloadUri,omitempty"`
	TaxonomyInformationURI                    string                `json:"informationUri,omitempty"`
	TaxonomyIsComprehensive                   bool                  `json:"isComprehensive,omitempty"`
	TaxonomyLanguage                          string                `json:"language,omitempty"`
	TaxonomyMinRequiredLocDataSemanticVersion string                `json:"minimumRequiredLocalizedDataSemanticVersion,omitempty"`
	TaxonomyOrganization                      string                `json:"organization,omitempty"`
	TaxonomyRealeaseDateUtc                   string                `json:"releaseDateUtc,omitempty"`
	TaxonomyDefinitions                       []taxonomyDefinitions `json:"taxa"`
}

// SarifRun - sarifRun is a component of the SARIF report
type SarifRun struct {
	Tool       sarifTool       `json:"tool"`
	Results    []sarifResult   `json:"results"`
	Taxonomies []sarifTaxonomy `json:"taxonomies,omitempty"`
}

// SarifReport represents a usable sarif report reference
type SarifReport interface {
	BuildSarifIssue(issue *model.QueryResult, sciInfo model.SCIInfo) string
	RebuildTaxonomies(cwes []string, guids map[string]string)
	GetGUIDFromRelationships(idx int, cweID string) string
	AddTags(summary *model.Summary, diffAware *model.DiffAware) error
	ResolveFilepaths(basePath string) error
	SetToolVersionType(runType string)
}

type sarifReport struct {
	Schema       string     `json:"$schema"`
	SarifVersion string     `json:"version"`
	Runs         []SarifRun `json:"runs"`
}

const (
	diffAwareConfigDigestTag = "DATADOG_DIFF_AWARE_CONFIG_DIGEST:%s"
	diffAwareEnabledTag      = "DATADOG_DIFF_AWARE_ENABLED:%v"
	diffAwareBaseShaTag      = "DATADOG_DIFF_AWARE_BASE_SHA:%s"
	diffAwareFileTag         = "DATADOG_DIFF_AWARE_FILE:%s"
	executionTimeTag         = "DATADOG_EXECUTION_TIME_SECS:%v"
	ruleTypeProperty         = "DATADOG_RULE_TYPE:IAC_SCANNING"
	categoryTag              = "DATADOG_CATEGORY:%s"
)

func initSarifTool() sarifTool {
	return sarifTool{
		Driver: sarifDriver{
			ToolName:     "Datadog IaC Scanning", // DO NOT CHANGE THIS VALUE. This value is needed for downstream mappings that rely on this value being "Datadog IaC Scanning"
			ToolVersion:  constants.Version,
			ToolFullName: constants.Fullname,
			ToolURI:      constants.URL,
			Rules:        make([]sarifRule, 0),
		},
	}
}

func initSarifCategories() []taxonomyDefinitions {
	allCategories := []taxonomyDefinitions{noCategory}
	for _, category := range categories {
		allCategories = append(allCategories, category)
	}
	return allCategories
}

// initCweCategories is responsible for building the CWE taxa field, inside taxonomies
func initCweCategories(cweIDs []string, guids map[string]string) []taxonomyDefinitions {
	absPath, err := filepath.Abs(".")
	if err != nil {
		return []taxonomyDefinitions{}
	}

	cweCSVPath := filepath.Join(absPath, "assets", "cwe_csv", "Software-Development-CWE.csv")
	cweCsvList, err := readCWECsvInfo(cweCSVPath)
	if err != nil {
		return []taxonomyDefinitions{}
	}

	var taxonomyList []taxonomyDefinitions
	for _, cweID := range cweIDs {
		var matchingCweEntry cweCsv

		for _, cweEntry := range cweCsvList {
			if cweEntry.CweID == cweID {
				matchingCweEntry = cweEntry
				break
			}
		}

		if matchingCweEntry.CweID == "" {
			continue
		}

		guid, exists := guids[cweID]
		if !exists {
			continue
		}

		taxonomy := taxonomyDefinitions{
			DefinitionID:               matchingCweEntry.CweID,
			DefinitionGUID:             guid,
			DefinitionFullDescription:  matchingCweEntry.FullDescription,
			DefinitionShortDescription: matchingCweEntry.ShortDescription,
			HelpURI:                    matchingCweEntry.HelpURI,
		}
		taxonomyList = append(taxonomyList, taxonomy)
	}

	return taxonomyList
}

func initSarifTaxonomies() []sarifTaxonomy {
	var taxonomies []sarifTaxonomy
	// Categories

	if targetTemplate.ToolComponent.ComponentReferenceName == "Categories" {
		categories := sarifTaxonomy{
			TaxonomyGUID: targetTemplate.ToolComponent.ComponentReferenceGUID,
			TaxonomyName: targetTemplate.ToolComponent.ComponentReferenceName,
			TaxonomyShortDescription: sarifMessage{
				Text: "Vulnerabilities categories",
			},
			TaxonomyFullDescription: sarifMessage{
				Text: "This taxonomy contains the types an issue can assume",
			},
			TaxonomyDefinitions: initSarifCategories(),
		}
		taxonomies = append(taxonomies, categories)
	}

	cweInfoPath := filepath.Join(".", "assets", "cwe_csv", "cwe_taxonomies_latest.json")

	if cweTemplate.ToolComponent.ComponentReferenceName == "CWE" {
		cweInfo, err := readCWEInfo(cweInfoPath)
		if err != nil {
			return taxonomies
		}
		cweTaxonomy := sarifTaxonomy{
			TaxonomyGUID:                              cweInfo.TaxonomyGUID,
			TaxonomyName:                              cweInfo.TaxonomyName,
			TaxonomyInformationURI:                    cweInfo.TaxonomyInformationURI,
			TaxonomyIsComprehensive:                   cweInfo.TaxonomyIsComprehensive,
			TaxonomyLanguage:                          cweInfo.TaxonomyLanguage,
			TaxonomyOrganization:                      cweInfo.TaxonomyOrganization,
			TaxonomyRealeaseDateUtc:                   cweInfo.TaxonomyRealeaseDateUtc,
			TaxonomyMinRequiredLocDataSemanticVersion: cweInfo.TaxonomyMinRequiredLocDataSemanticVersion,
			TaxonomyDownloadURI:                       cweInfo.TaxonomyDownloadURI,
			TaxonomyFullDescription:                   sarifMessage{Text: cweInfo.TaxonomyFullDescription.Text},
			TaxonomyShortDescription:                  sarifMessage{Text: cweInfo.TaxonomyShortDescription.Text},
			TaxonomyDefinitions:                       []taxonomyDefinitions{},
		}
		taxonomies = append(taxonomies, cweTaxonomy)
	}

	return taxonomies
}

func initSarifRun() []SarifRun {
	return []SarifRun{
		{
			Tool:    initSarifTool(),
			Results: make([]sarifResult, 0),
			// Taxonomies: initSarifTaxonomies(),
		},
	}
}

// NewSarifReport creates and start a new sarif report with default values respecting SARIF schema 2.1.0
func NewSarifReport() SarifReport {
	return &sarifReport{
		Schema:       "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json",
		SarifVersion: "2.1.0",
		Runs:         initSarifRun(),
	}
}

func (sr *sarifReport) findSarifCategory(category string) int {
	for idx, taxonomy := range sr.Runs[0].Taxonomies[0].TaxonomyDefinitions {
		if taxonomy.DefinitionName == category {
			return idx
		}
	}
	return 0
}

// readCWEInfo is responsible for reading the CWE taxonomy info from the json file
func readCWEInfo(filePath string) (sarifTaxonomy, error) {
	var wrapper cweTaxonomiesWrapper
	fileContent, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return sarifTaxonomy{}, err
	}

	err = json.Unmarshal(fileContent, &wrapper)
	if err != nil {
		return sarifTaxonomy{}, err
	}

	return wrapper.Taxonomies, nil
}

func generateGUID() string {
	id := uuid.New()
	return id.String()
}

// readCWECsvInfo is responsible for reading the CWE taxonomy info from the corresponding csv file
func readCWECsvInfo(filePath string) ([]cweCsv, error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Note: -1 means records may have a variable number of fields in the csv file
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var cweEntries []cweCsv
	numRecords := 23

	for _, record := range records {
		if len(record) >= numRecords {
			cweEntry := cweCsv{
				CweID: record[0],
				FullDescription: cweMessage{
					Text: record[5],
				},
				ShortDescription: cweMessage{
					Text: record[4],
				},
				GUID:    generateGUID(),
				HelpURI: "https://cwe.mitre.org/data/definitions/" + record[0] + ".html",
			}

			// Check if Extended Description is empty, fill it with Description if so
			if cweEntry.FullDescription.Text == "" {
				cweEntry.FullDescription.Text = record[4]
			}

			cweEntries = append(cweEntries, cweEntry)
		}
	}

	return cweEntries, nil
}

// buildCweCategory builds the CWE category in taxonomies, with info from CWE and CSV
func (sr *sarifReport) buildCweCategory(cweID string) sarifDescriptorReference {
	absPath, err := filepath.Abs(".")
	if err != nil {
		return sarifDescriptorReference{}
	}

	cweInfoPath := filepath.Join(absPath, "assets", "cwe_csv", "cwe_taxonomies_latest.json")
	cweCSVPath := filepath.Join(absPath, "assets", "cwe_csv", "Software-Development-CWE.csv")

	cweInfo, err := readCWEInfo(cweInfoPath)
	if err != nil {
		return sarifDescriptorReference{}
	}

	_ = cweInfo

	cweCsvList, err := readCWECsvInfo(cweCSVPath)
	if err != nil {
		return sarifDescriptorReference{}
	}

	var matchingCweEntry cweCsv
	for _, cweEntry := range cweCsvList {
		if cweEntry.CweID == cweID {
			matchingCweEntry = cweEntry
			break
		}
	}

	if matchingCweEntry.CweID == "" {
		return sarifDescriptorReference{}
	}

	newGUID := generateGUID()

	cwe := sarifDescriptorReference{
		ReferenceID:   matchingCweEntry.CweID,
		ReferenceGUID: newGUID,
		ToolComponent: sarifComponentReference{
			ComponentReferenceGUID: "1489b0c4-d7ce-4d31-af66-6382a01202e3",
			ComponentReferenceName: "CWE",
		},
	}

	return cwe
}

func (sr *sarifReport) buildSarifCategory(category string) sarifDescriptorReference {
	target := targetTemplate
	categoryIndex := sr.findSarifCategory(category)

	if categoryIndex >= 0 {
		taxonomy := sr.Runs[0].Taxonomies[0].TaxonomyDefinitions[categoryIndex]

		target.ReferenceID = taxonomy.DefinitionID
	}
	target.ReferenceIndex = categoryIndex

	if categoryIndex == -1 {
		if _, exists := categoriesNotFound[category]; !exists {
			log.Warn().Msgf("Category %s not found.", category)
			categoriesNotFound[category] = true
		}
	}

	return target
}

func (sr *sarifReport) findSarifRuleIndex(ruleID string) int {
	for idx := range sr.Runs[0].Tool.Driver.Rules {
		if sr.Runs[0].Tool.Driver.Rules[idx].RuleID == ruleID {
			return idx
		}
	}
	return -1
}

func (sr *sarifReport) buildSarifRule(queryMetadata *ruleMetadata, cisMetadata ruleCISMetadata) int {
	index := sr.findSarifRuleIndex(queryMetadata.queryID)

	if index < 0 {
		helpURI := "https://docs.kics.io/"
		if queryMetadata.queryURI != "" {
			helpURI = queryMetadata.queryURI
		}

		tags := []string{ruleTypeProperty}
		cwe := queryMetadata.queryCwe
		if cwe != "" {
			cweTag := GetCWETag(cwe)
			tags = append(tags, cweTag)
		}

		categoryTag := GetCategoryTag(queryMetadata.queryCategory)
		kicsRuleIDTag := GetKICSRuleIDTag(queryMetadata.queryID)
		tags = append(tags, categoryTag, kicsRuleIDTag)

		rule := sarifRule{
			RuleID:               queryMetadata.queryName,
			RuleName:             queryMetadata.queryName,
			RuleShortDescription: sarifMessage{Text: queryMetadata.queryName},
			RuleFullDescription:  sarifMessage{Text: fmt.Sprintf("[%s] %s", queryMetadata.queryID, queryMetadata.queryDescription)},
			DefaultConfiguration: sarifConfiguration{Level: severityLevelEquivalence[queryMetadata.severity]},
			// Relationships:        relationships,
			HelpURI: helpURI,
			RuleProperties: sarifProperties{
				"tags": tags,
			},
		}
		if cisMetadata.id != "" {
			rule.RuleFullDescription.Text = cisMetadata.descriptionText
			rule.RuleProperties = sarifProperties{
				"cisId":    cisMetadata.id,
				"cisTitle": cisMetadata.title,
			}
		}

		sr.Runs[0].Tool.Driver.Rules = append(sr.Runs[0].Tool.Driver.Rules, rule)
		index = len(sr.Runs[0].Tool.Driver.Rules) - 1
	}
	return index
}

// GetGUIDFromRelationships gets the GUID from the relationship for each CWE item
func (sr *sarifReport) GetGUIDFromRelationships(idx int, cweID string) string {
	if len(sr.Runs) > 0 {
		if len(sr.Runs[0].Tool.Driver.Rules) > 0 {
			relationships := sr.Runs[0].Tool.Driver.Rules[idx].Relationships
			for _, relationship := range relationships {
				target := relationship.Relationship

				if target.ReferenceID == cweID {
					return target.ReferenceGUID
				}
			}
		}
	}

	return ""
}

// RebuildTaxonomies builds the taxonomies with the CWEs and the GUIDs coming from each relationships field
func (sr *sarifReport) RebuildTaxonomies(cwes []string, guids map[string]string) {
	if len(cwes) > 0 {
		result := initCweCategories(cwes, guids)
		if len(sr.Runs) > 0 {
			if len(sr.Runs[0].Taxonomies) == 2 {
				sr.Runs[0].Taxonomies[1].TaxonomyDefinitions = result
			}
		}
	}
}

// BuildSarifIssue creates a new entries in Results (one for each file) and new entry in Rules and Taxonomy if necessary
func (sr *sarifReport) BuildSarifIssue(issue *model.QueryResult, sciInfo model.SCIInfo) string {
	if len(issue.Files) > 0 {
		metadata := ruleMetadata{
			queryID:          issue.QueryID,
			queryName:        issue.QueryName,
			queryDescription: issue.Description,
			queryURI:         issue.QueryURI,
			queryCategory:    issue.Category,
			queryCwe:         issue.CWE,
			severity:         issue.Severity,
		}
		cisDescriptions := ruleCISMetadata{
			id:              issue.CISDescriptionIDFormatted,
			title:           issue.CISDescriptionTitle,
			descriptionText: issue.CISDescriptionTextFormatted,
		}
		ruleIndex := sr.buildSarifRule(&metadata, cisDescriptions)

		categoryTag := GetCategoryTag(issue.Category)
		tags := []string{categoryTag}
		cwe := issue.CWE
		if cwe != "" {
			cweTag := GetCWETag(cwe)
			tags = append(tags, cweTag)
		}

		for idx := range issue.Files {
			line := issue.Files[idx].Line
			if line < 1 {
				line = 1
			}
			vulnerability := issue.Files[idx]

			resourceType := vulnerability.ResourceType
			resourceName := vulnerability.ResourceName
			resourceTypeTag := GetResourceTypeTag(resourceType)
			resourceNameTag := GetResourceNameTag(resourceName)

			resultTags := append(tags, resourceTypeTag, resourceNameTag)

			resourceLocation := vulnerability.ResourceLocation

			if resourceLocation.Start.Line < 1 {
				resourceLocation.Start.Line = 1
			}
			if resourceLocation.End.Line < 1 {
				resourceLocation.End.Line = resourceLocation.Start.Line
			}
			if resourceLocation.Start.Col < 1 {
				resourceLocation.Start.Col = 1
			}
			if resourceLocation.End.Col < 1 {
				resourceLocation.End.Col = resourceLocation.Start.Col
			}

			resourceStartLocation := sarifResourceLocation{
				Line: resourceLocation.Start.Line,
				Col:  resourceLocation.Start.Col,
			}
			resourceEndLocation := sarifResourceLocation{
				Line: resourceLocation.End.Line,
				Col:  resourceLocation.End.Col,
			}

			remediationLocation := vulnerability.RemediationLocation
			remediationStartLocation := sarifResourceLocation{
				Line: remediationLocation.Start.Line,
				Col:  remediationLocation.Start.Col,
			}
			remediationEndLocation := sarifResourceLocation{
				Line: remediationLocation.End.Line,
				Col:  remediationLocation.End.Col,
			}

			if resourceStartLocation.Col < 1 {
				resourceStartLocation.Col = 1
			}

			absoluteFilePath := strings.ReplaceAll(issue.Files[idx].FileName, "../", "")
			result := sarifResult{
				ResultRuleID:    issue.QueryName,
				ResultRuleIndex: ruleIndex,
				ResultLevel:     severityLevelEquivalence[issue.Severity],
				ResultMessage: sarifMessage{
					Text: issue.Files[idx].KeyActualValue,
				},
				ResultLocations: []sarifLocation{
					{
						PhysicalLocation: sarifPhysicalLocation{
							ArtifactLocation: sarifArtifactLocation{ArtifactURI: absoluteFilePath},
							Region: sarifRegion{
								StartLine:   resourceStartLocation.Line,
								EndLine:     resourceEndLocation.Line,
								StartColumn: resourceStartLocation.Col,
								EndColumn:   resourceEndLocation.Col,
							},
						},
					},
				},
				ResultProperties: sarifProperties{
					"tags": resultTags,
				},
				PartialFingerprints: SarifPartialFingerprints{
					DatadogFingerprint: GetDatadogFingerprintHash(sciInfo, absoluteFilePath, line, issue.QueryID),
				},
			}
			if vulnerability.Remediation != "" && vulnerability.RemediationType != "" {
				sarifFix, err := TransformToSarifFix(
					vulnerability,
					remediationStartLocation,
					remediationEndLocation,
				)
				if err != nil {
					// we do not want to fail the whole process if we cannot transform the fix
					// so we just log the error and continue
					log.Err(err).Msgf("failed to transform to sarif fix: %v", err)
				} else {
					// if vulnerability.RemediationType == "addition" {
					// 	result.ResultLocations[0].PhysicalLocation.Region.StartLine = resourceEndLocation.Line
					// 	result.ResultLocations[0].PhysicalLocation.Region.EndLine = resourceEndLocation.Line + 4
					// }
					result.ResultFixes = append(result.ResultFixes, sarifFix)
				}
			}
			sr.Runs[0].Results = append(sr.Runs[0].Results, result)
		}
		return issue.CWE
	}
	return ""
}

func (sr *sarifReport) SetToolVersionType(runType string) {
	if len(runType) > 0 {
		for idx := range sr.Runs {
			sr.Runs[idx].Tool.Driver.ToolVersion = runType
		}
		log.Info().Msgf("Tool version set to %s", runType)
	}
}

func (sr *sarifReport) AddTags(summary *model.Summary, diffAware *model.DiffAware) error {
	if len(sr.Runs) != 1 {
		return errors.New("sarifReport must have exactly one run")
	}
	tagsToAppend := []string{}
	executionTimeTag := GetScanDurationTag(*summary)
	diffAwareEnabledTag := GetDiffAwareEnabledTag(*diffAware)

	tagsToAppend = append(tagsToAppend, executionTimeTag, diffAwareEnabledTag)

	if diffAware.Enabled {
		if diffAware.BaseSha == "" || diffAware.Files == "" || diffAware.ConfigDigest == "" {
			return fmt.Errorf(
				"diffAware enabled but base sha %s, files %s, config digest %s provided",
				diffAware.BaseSha,
				diffAware.Files,
				diffAware.ConfigDigest,
			)
		}
		diffAwareConfigDigestTag := GetDiffAwareConfigDigestTag(*diffAware)
		diffAwareBaseShaTag := GetDiffAwareBaseShaTag(*diffAware)
		diffAwareFileTag := GetDiffAwareFilesTag(*diffAware)

		tagsToAppend = append(tagsToAppend, diffAwareConfigDigestTag, diffAwareBaseShaTag, diffAwareFileTag)
	}

	sarifToolProperties := &sr.Runs[0].Tool.Driver.Properties
	sarifToolProperties.Tags = append(
		sarifToolProperties.Tags,
		tagsToAppend...,
	)

	return nil
}

func (sr *sarifReport) ResolveFilepaths(basePath string) error {
	for idx := range sr.Runs[0].Results {
		artifactLocation := &sr.Runs[0].Results[idx].ResultLocations[0].PhysicalLocation.ArtifactLocation.ArtifactURI
		resolvedLocation := strings.TrimPrefix(*artifactLocation, basePath+"/")
		sr.Runs[0].Results[idx].ResultLocations[0].PhysicalLocation.ArtifactLocation.ArtifactURI = resolvedLocation
	}
	return nil
}

func GetDatadogFingerprintHash(sciInfo model.SCIInfo, filePath string, startLine int, ruleId string) string {
	runTypeRelatedInfo := sciInfo.RepositoryCommitInfo.CommitSHA
	if sciInfo.RunType == "full_scan" {
		runTypeRelatedInfo = time.Now().Format("2006/01/02")
	}
	return StringToHash(fmt.Sprintf("%s|%s|%s|%s|%d|%s", sciInfo.RunType, runTypeRelatedInfo, sciInfo.RepositoryCommitInfo.RepositoryUrl, filePath, startLine, ruleId))
}

func TransformToSarifFix(vuln model.VulnerableFile, startLocation sarifResourceLocation, endLocation sarifResourceLocation) (sarifFix, error) {
	var insertedText string
	fixStart := startLocation
	fixEnd := endLocation
	sourceLines := vuln.ResourceSource
	fileLines := vuln.FileSource

	// Helper to detect indentation
	detectIndent := func(line string) string {
		for i, r := range line {
			if r != ' ' && r != '\t' {
				return line[:i]
			}
		}
		return ""
	}

	// Detect base indentation from first non-empty source line
	baseIndent := ""
	lines := strings.Split(sourceLines, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			baseIndent = detectIndent(line)
			break
		}
	}

	switch vuln.RemediationType {

	// case "replacement":
	// 	var patch map[string]string
	// 	err := json.Unmarshal([]byte(vuln.Remediation), &patch)
	// 	if err != nil {
	// 		return sarifFix{}, fmt.Errorf("invalid remediation format for replacement: %v", err)
	// 	}
	// 	after := strings.TrimSpace(patch["after"])

	// 	// Re-indent `after` content to match block
	// 	afterLines := strings.Split(after, "\n")
	// 	for i, line := range afterLines {
	// 		afterLines[i] = baseIndent + strings.TrimRight(line, " ")
	// 	}
	// 	insertedText = strings.Join(afterLines, "\n")
	// 	if len(insertedText) > 0 && insertedText[len(insertedText)-1] != '\n' {
	// 		insertedText += "\n"
	// 	}

	// 	fixStart = startLocation
	// 	fixEnd = endLocation

	case "replacement":
		var patch map[string]string
		err := json.Unmarshal([]byte(vuln.Remediation), &patch)
		if err != nil {
			return sarifFix{}, fmt.Errorf("invalid remediation format for replacement: %v", err)
		}

		before := strings.TrimSpace(patch["before"])
		after := strings.TrimSpace(patch["after"])

		re := regexp.MustCompile(`(?m)["']?(?P<key>\w+)["']?\s*[:=]\s*(?P<value>\[.*?\]|".*?"|true|false|\d+)[,]?\s*`)
		matches := re.FindStringSubmatch(vuln.LineWithVulnerability)
		if len(matches) < 3 {
			return sarifFix{}, fmt.Errorf("could not parse key-value from line: %s", vuln.LineWithVulnerability)
		}

		key := strings.TrimSpace(matches[1])
		value := strings.TrimSpace(matches[2])

		if idx := strings.IndexAny(value, "#/"); idx != -1 {
			value = strings.TrimSpace(value[:idx])
		}

		fullLine := key + " = " + value

		normalizedFullLine := normalize(fullLine)
		normalizedValue := normalize(value)

		// Support multiple 'before' values separated by 'or'
		normalizedAlternatives := []string{normalize(before)}
		if strings.Contains(before, " or ") {
			rawParts := strings.Split(before, " or ")
			normalizedAlternatives = make([]string, 0, len(rawParts))
			for _, part := range rawParts {
				normalizedAlternatives = append(normalizedAlternatives, normalize(strings.TrimSpace(part)))
			}
		}

		insertedText = ""
		matched := false

		for _, alt := range normalizedAlternatives {
			altKey := ""
			altValue := alt

			if parts := strings.SplitN(alt, "=", 2); len(parts) == 2 {
				altKey = normalize(strings.TrimSpace(parts[0]))
				altValue = normalize(strings.TrimSpace(parts[1]))
			} else {
				altValue = normalize(alt)
			}

			if (alt == normalizedFullLine) ||
				(altKey == normalize(key) && altValue == normalizedValue) ||
				strings.Contains(normalizedValue, altValue) ||
				strings.HasPrefix(normalizedValue, altValue) ||
				strings.HasPrefix(altValue, normalizedValue) {
				matched = true
				break
			}
		}

		// Block-style fallback (e.g., nested blocks like metadata_options)
		if !matched && strings.Contains(normalize(before), "{") && strings.Contains(normalize(before), "=") {
			reInner := regexp.MustCompile(`(?m)(\w+)\s*=\s*(".*?"|\[.*?\]|\S+)`)
			for _, inner := range reInner.FindAllString(before, -1) {
				n := normalize(inner)
				for _, alt := range normalizedAlternatives {
					if n == normalizedFullLine || n == normalizedValue || strings.Contains(normalizedFullLine, n) || strings.Contains(n, alt) {
						matched = true
						break
					}
				}
				if matched {
					break
				}
			}
		}

		// List-style fallback
		if !matched && strings.HasPrefix(normalizedValue, "[") && strings.HasSuffix(normalizedValue, "]") {
			listText := strings.Trim(normalizedValue, "[] ")
			listItems := strings.Split(listText, ",")

			newList := make([]string, 0, len(listItems))
			replaced := false

			// Pre-extract values from all alternatives (e.g., extract just "DISABLED" from "key = DISABLED")
			normalizedAltValues := make([]string, 0, len(normalizedAlternatives))
			for _, alt := range normalizedAlternatives {
				parts := strings.SplitN(alt, "=", 2)
				if len(parts) == 2 {
					normalizedAltValues = append(normalizedAltValues, strings.TrimSpace(parts[1]))
				} else {
					normalizedAltValues = append(normalizedAltValues, strings.TrimSpace(alt))
				}
			}

			for _, item := range listItems {
				itemStripped := strings.TrimSpace(strings.Trim(item, `"`))
				normOriginal := normalize(itemStripped)
				replacedItem := false

				for _, altVal := range normalizedAltValues {
					normAltVal := normalize(altVal)

					if normOriginal == normAltVal ||
						strings.Contains(normOriginal, normAltVal) ||
						strings.Contains(normAltVal, normOriginal) {
						replaced = true
						replacedItem = true
						if strings.HasPrefix(item, `"`) && strings.HasSuffix(item, `"`) {
							newList = append(newList, `"`+after+`"`)
						} else {
							newList = append(newList, after)
						}
						break
					}
				}

				if !replacedItem {
					newList = append(newList, item)
				}
			}

			if replaced {
				joined := "[" + strings.Join(newList, ", ") + "]"
				insertedText = strings.Replace(vuln.LineWithVulnerability, value, joined, 1)
				matched = true
			} else {
				// Check if the normalized list as a whole matches one of the alt values
				for _, altVal := range normalizedAltValues {
					if normalizedValue == normalize(altVal) {
						matched = true
						insertedText = vuln.LineWithVulnerability
						break
					}
				}
			}
		}

		if !matched {
			return sarifFix{}, fmt.Errorf("line value '%s' does not match any of expected values '%s'", normalizedFullLine, strings.Join(normalizedAlternatives, " | "))
		}

		// Preserve quotes if necessary
		wasQuoted := strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`)
		if wasQuoted && !(strings.HasPrefix(after, `"`) && strings.HasSuffix(after, `"`)) {
			after = `"` + after + `"`
		}

		// Determine replacement range
		idxs := re.FindStringSubmatchIndex(vuln.LineWithVulnerability)
		if len(idxs) < 6 {
			return sarifFix{}, fmt.Errorf("could not determine exact value location")
		}

		isFullLine := strings.Contains(after, "=")

		if insertedText == "" {
			if isFullLine {
				insertedText = after
			} else {
				prefix := vuln.LineWithVulnerability[:idxs[4]]
				suffix := vuln.LineWithVulnerability[idxs[5]:]
				insertedText = prefix + after + suffix
			}
		}

		fixStart = sarifResourceLocation{
			Line: vuln.Line,
			Col:  startLocation.Col,
		}
		fixEnd = sarifResourceLocation{
			Line: vuln.Line,
			Col:  len(vuln.LineWithVulnerability) + 1,
		}
	case "addition":
		// Add content with proper indentation
		// Use baseIndent (block-level indent) + one level deeper
		// innerIndent := baseIndent + "  " // assuming 2-space indentation; adjust if needed
		normalizedRemediation := normalizeIndentation(vuln.Remediation, 2)

		followingIndent := detectIndent(fileLines[startLocation.Line-1])
		// // Try to get indentation from the first non-empty line after insertion
		// for i := startLocation.Line - 1; i < len(lines)-1; i++ {
		// 	nextLine := strings.TrimSpace(lines[i+1])
		// 	if nextLine != "" {
		// 		followingIndent = detectIndent(lines[i+1])
		// 		break
		// 	}
		// }

		insertedLines := strings.Split(normalizedRemediation, "\n")
		for i, line := range insertedLines {
			if strings.TrimSpace(line) != "" {
				if strings.TrimSpace(line) == "{" || strings.TrimSpace(line) == "}" {
					insertedLines[i] = baseIndent + strings.TrimRight(line, " ")
				} else {
					insertedLines[i] = followingIndent
				}
			}
		}
		insertedText = "\n" + strings.Join(insertedLines, "\n") + "\n" + followingIndent

		fixStart = startLocation
		fixEnd = startLocation

	case "removal":
		// Just remove the whole region
		insertedText = ""
		fixStart = startLocation
		fixEnd = endLocation
	}

	replacement := fixReplacement{
		DeletedRegion: sarifRegion{
			StartLine:   fixStart.Line,
			StartColumn: fixStart.Col,
			EndLine:     fixEnd.Line,
			EndColumn:   fixEnd.Col,
		},
	}

	if insertedText != "" {
		replacement.InsertedContent = fixContent{Text: insertedText}
	}

	return sarifFix{
		ArtifactChanges: []artifactChange{{
			ArtifactLocation: artifactLocation{URI: vuln.FileName},
			Replacements:     []fixReplacement{replacement},
		}},
		Description: fixMessage{
			Text: fmt.Sprintf("Apply remediation: %s", vuln.Remediation),
		},
	}, nil
}

func normalizeIndentation(input string, spacesPerTab int) string {
	lines := strings.Split(input, "\n")
	tabSpaces := strings.Repeat(" ", spacesPerTab)

	for i, line := range lines {
		// Replace tabs with spaces and trim trailing whitespace
		line = strings.ReplaceAll(line, "\t", tabSpaces)
		lines[i] = strings.TrimRight(line, " \t")
	}

	return strings.Join(lines, "\n")
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, `\"`, `"`)     // Handle escaped quotes
	s = strings.ReplaceAll(s, `"`, `"`)      // Double safety
	s = strings.ReplaceAll(s, `“`, `"`)      // Smart quotes (optional)
	s = strings.ReplaceAll(s, `”`, `"`)      // Smart quotes (optional)
	s = strings.Join(strings.Fields(s), " ") // normalize extra whitespace
	s = strings.Trim(s, `"`)
	return s
}

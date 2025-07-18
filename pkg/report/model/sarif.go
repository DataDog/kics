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
	"strings"

	"github.com/Checkmarx/kics/internal/constants"
	"github.com/Checkmarx/kics/pkg/model"
	remediationsHelper "github.com/Checkmarx/kics/pkg/report/remediations"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var categoriesNotFound = make(map[string]bool)

var severityLevelEquivalence = map[model.Severity]string{
	"INFO":     "none",
	"LOW":      "none",
	"MEDIUM":   "note",
	"HIGH":     "warning",
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

type sarifArtifactLocation struct {
	ArtifactURI string `json:"uri"`
}

type sarifPhysicalLocation struct {
	ArtifactLocation sarifArtifactLocation `json:"artifactLocation"`
	Region           model.SarifRegion     `json:"region"`
}

type SarifLocation struct {
	PhysicalLocation sarifPhysicalLocation `json:"physicalLocation"`
}

type sarifResult struct {
	ResultRuleID        string                   `json:"ruleId"`
	ResultRuleIndex     int                      `json:"ruleIndex"`
	ResultKind          string                   `json:"kind,omitempty"`
	ResultMessage       sarifMessage             `json:"message"`
	ResultLocations     []SarifLocation          `json:"locations"`
	PartialFingerprints SarifPartialFingerprints `json:"partialFingerprints,omitempty"`
	ResultLevel         string                   `json:"level"`
	ResultProperties    sarifProperties          `json:"properties,omitempty"`
	ResultFixes         []model.SarifFix         `json:"fixes,omitempty"`
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
	scannedFileCountTag      = "DATADOG_SCANNED_FILE_COUNT:%d"
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
			RuleFullDescription:  sarifMessage{Text: fmt.Sprintf("%s\nRule ID: [%s]", queryMetadata.queryDescription, queryMetadata.queryID)},
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

			if resourceLocation.Start.Line < 1 || resourceLocation.End.Line < 1 {
				log.Warn().Msgf("Invalid resource location for file %s", issue.Files[idx].FileName)
				continue
			}

			if resourceLocation.Start.Col < 1 {
				resourceLocation.Start.Col = 1
			}
			if resourceLocation.End.Col < 1 {
				resourceLocation.End.Col = resourceLocation.Start.Col
			}

			resourceStartLocation := model.SarifResourceLocation{
				Line: resourceLocation.Start.Line,
				Col:  resourceLocation.Start.Col,
			}
			resourceEndLocation := model.SarifResourceLocation{
				Line: resourceLocation.End.Line,
				Col:  resourceLocation.End.Col,
			}

			remediationLocation := vulnerability.RemediationLocation
			remediationStartLocation := model.SarifResourceLocation{
				Line: remediationLocation.Start.Line,
				Col:  remediationLocation.Start.Col,
			}
			remediationEndLocation := model.SarifResourceLocation{
				Line: remediationLocation.End.Line,
				Col:  remediationLocation.End.Col,
			}

			if resourceStartLocation.Col < 1 {
				resourceStartLocation.Col = 1
			}

			if remediationStartLocation.Line >= resourceStartLocation.Line && remediationEndLocation.Line <= resourceEndLocation.Line {
				resourceStartLocation = remediationStartLocation
				resourceEndLocation = remediationEndLocation
			}

			absoluteFilePath := strings.ReplaceAll(issue.Files[idx].FileName, "../", "")
			result := sarifResult{
				ResultRuleID:    issue.QueryName,
				ResultRuleIndex: ruleIndex,
				ResultLevel:     severityLevelEquivalence[issue.Severity],
				ResultMessage: sarifMessage{
					Text: issue.Files[idx].KeyActualValue,
				},
				ResultLocations: []SarifLocation{
					{
						PhysicalLocation: sarifPhysicalLocation{
							ArtifactLocation: sarifArtifactLocation{ArtifactURI: absoluteFilePath},
							Region: model.SarifRegion{
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
					DatadogFingerprint: GetDatadogFingerprintHash(sciInfo, absoluteFilePath, resourceType, resourceName, issue.QueryID),
				},
			}
			if vulnerability.Remediation != "" && vulnerability.RemediationType != "" {
				sarifFix, err := remediationsHelper.TransformToSarifFix(
					vulnerability,
					remediationStartLocation,
					remediationEndLocation,
				)
				if err != nil {
					// we do not want to fail the whole process if we cannot transform the fix
					// so we just log the error and continue
					log.Err(err).Msgf("failed to transform to sarif fix: %v", err)
				} else {
					// we want the location displayed in the UI to properly highlight the remediation
					result.ResultLocations[0].PhysicalLocation.Region.StartLine = remediationStartLocation.Line
					result.ResultLocations[0].PhysicalLocation.Region.EndLine = remediationEndLocation.Line

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
	scannedFileCountTag := GetScannedFilesCountTag(summary.ScannedFiles)
	diffAwareEnabledTag := GetDiffAwareEnabledTag(*diffAware)

	tagsToAppend = append(tagsToAppend, executionTimeTag, scannedFileCountTag, diffAwareEnabledTag)

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

func GetDatadogFingerprintHash(sciInfo model.SCIInfo, filePath, resourceType, resourceName, ruleId string) string {
	return StringToHash(fmt.Sprintf("%s|%s|%s|%s|%s", sciInfo.RepositoryCommitInfo.RepositoryUrl, filePath, resourceType, resourceName, ruleId))
}

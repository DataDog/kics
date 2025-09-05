/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package model

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/Checkmarx/kics/pkg/model"
)

const (
	kicsRuleIDTag           = "KICS_RuleID:%s"
	cweTag                  = "CWE:%s"
	resourceTypeTag         = "IAC_RESOURCE_TYPE:%s"
	resourceNameTag         = "IAC_RESOURCE_NAME:%s"
	frameworkTag            = "framework:%s"
	frameworkVersionTag     = "framework_version:%s"
	frameworkRequirementTag = "requirement:%s"
	frameworkControlTag     = "control:%s"
)

func GetScanDurationTag(summary model.Summary) string {
	scanDuration := int(summary.Times.End.Sub(summary.Times.Start).Seconds())
	executionTimeTag := fmt.Sprintf(executionTimeTag, scanDuration)
	return executionTimeTag
}

func GetDiffAwareEnabledTag(diffAware model.DiffAware) string {
	return fmt.Sprintf(diffAwareEnabledTag, diffAware.Enabled)
}

func GetDiffAwareConfigDigestTag(diffAware model.DiffAware) string {
	return fmt.Sprintf(diffAwareConfigDigestTag, diffAware.ConfigDigest)
}

func GetDiffAwareBaseShaTag(diffAware model.DiffAware) string {
	return fmt.Sprintf(diffAwareBaseShaTag, diffAware.BaseSha)
}

func GetDiffAwareFilesTag(diffAware model.DiffAware) string {
	return fmt.Sprintf(diffAwareFileTag, diffAware.Files)
}

func GetCategoryTag(category string) string {
	return fmt.Sprintf(categoryTag, category)
}

func GetKICSRuleIDTag(ruleID string) string {
	return fmt.Sprintf(kicsRuleIDTag, ruleID)
}

func GetCWETag(cwe string) string {
	return fmt.Sprintf(cweTag, cwe)
}

func GetResourceTypeTag(resourceType string) string {
	return fmt.Sprintf(resourceTypeTag, resourceType)
}

func GetResourceNameTag(resourceName string) string {
	return fmt.Sprintf(resourceNameTag, resourceName)
}

func GetScannedFilesCountTag(scannedFiles int) string {
	return fmt.Sprintf(scannedFileCountTag, scannedFiles)
}

func GetFrameworkTags(framework model.CustomFramework) []string {
	return []string{
		fmt.Sprintf(frameworkTag, framework.Framework),
		fmt.Sprintf(frameworkVersionTag, framework.FrameworkVersion),
		fmt.Sprintf(frameworkRequirementTag, framework.Requirement),
		fmt.Sprintf(frameworkControlTag, framework.Control),
	}
}

// stringToHash returns a SHA256 hash of the input string.
func StringToHash(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

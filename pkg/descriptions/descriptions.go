/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package descriptions

import (
	"context"
	"fmt"

	"github.com/Checkmarx/kics/pkg/model"
)

var (
	descClient HTTPDescription = &Client{}
)

// RequestAndOverrideDescriptions - Requests descriptions and override default descriptions
func RequestAndOverrideDescriptions(ctx context.Context, summary *model.Summary) error {
	descriptionIDs := make([]string, 0)
	for idx := range summary.Queries {
		descriptionIDs = append(descriptionIDs, summary.Queries[idx].DescriptionID)
	}

	if err := descClient.CheckConnection(ctx); err != nil {
		return err
	}

	descriptionMap, err := descClient.RequestDescriptions(ctx, descriptionIDs)
	if err != nil {
		return err
	}

	for idx := range summary.Queries {
		if descriptionMap[summary.Queries[idx].DescriptionID].DescriptionID == "" &&
			descriptionMap[summary.Queries[idx].DescriptionID].RationaleText == "" {
			continue
		}
		descriptionID := summary.Queries[idx].DescriptionID

		summary.Queries[idx].CISDescriptionID = descriptionMap[descriptionID].DescriptionID
		summary.Queries[idx].CISDescriptionTitle = descriptionMap[descriptionID].DescriptionTitle
		summary.Queries[idx].CISDescriptionText = descriptionMap[descriptionID].DescriptionText
		summary.Queries[idx].CISRationaleText = descriptionMap[descriptionID].RationaleText
		summary.Queries[idx].CISBenchmarkName = descriptionMap[descriptionID].BenchmarkName
		summary.Queries[idx].CISBenchmarkVersion = descriptionMap[descriptionID].BenchmarkVersion

		summary.Queries[idx].CISDescriptionIDFormatted = fmt.Sprintf(
			"Security - %s v%s - Rule %s",
			descriptionMap[descriptionID].BenchmarkName,
			descriptionMap[descriptionID].BenchmarkVersion,
			descriptionMap[descriptionID].DescriptionID,
		)
		summary.Queries[idx].CISDescriptionTextFormatted = fmt.Sprintf(
			"%s\n%s",
			descriptionMap[descriptionID].DescriptionText,
			descriptionMap[descriptionID].RationaleText,
		)
	}
	return nil
}

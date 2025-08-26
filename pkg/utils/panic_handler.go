/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package utils

import (
	"context"
	"fmt"

	"github.com/Checkmarx/kics/pkg/logger"
)

func HandlePanic(ctx context.Context, r any, errMessage string) {
	logger := logger.FromContext(ctx)
	err := fmt.Errorf("panic: %v", r)
	logger.Err(err).Msg(errMessage)
}

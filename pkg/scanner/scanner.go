/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package scanner

import (
	"context"
	"sync"

	"github.com/Checkmarx/kics/internal/metrics"
	"github.com/Checkmarx/kics/pkg/kics"
	"github.com/Checkmarx/kics/pkg/logger"
)

type serviceSlice []*kics.Service

func PrepareAndScan(
	ctx context.Context,
	scanID string,
	openAPIResolveReferences bool,
	maxResolverDepth int,
	services serviceSlice,
) error {
	metrics.Metric.Start("prepare_sources")
	var wg sync.WaitGroup
	wgDone := make(chan bool)
	errCh := make(chan error)
	defer close(errCh)

	for _, service := range services {
		wg.Add(1)
		go service.PrepareSources(ctx, scanID, openAPIResolveReferences, maxResolverDepth, &wg, errCh)
	}

	go func() {
		defer close(wgDone)
		wg.Wait()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-wgDone:
		metrics.Metric.Stop()
		err := StartScan(ctx, scanID, services)
		if err != nil {
			return err
		}
		break
	case err := <-errCh:
		return err
	}
	return nil
}

// StartScan will run concurrent scans by parser
func StartScan(ctx context.Context, scanID string, services serviceSlice) error {
	logger := logger.FromContext(ctx)
	defer metrics.Metric.Stop()
	metrics.Metric.Start("start_scan")
	var wg sync.WaitGroup
	wgDone := make(chan bool)
	errCh := make(chan error)

	logger.Info().Msgf("Starting scan with id: %s", scanID)

	total := services.GetQueriesLength()
	logger.Info().Msgf("Got %d queries", total)

	for _, service := range services {
		wg.Add(1)
		go service.StartScan(ctx, scanID, errCh, &wg)
	}

	go func() {
		defer close(wgDone)
		wg.Wait()
	}()

	select {
	case <-ctx.Done():
		close(errCh)
		return ctx.Err()
	case <-wgDone:
		break
	case err := <-errCh:
		close(errCh)
		return err
	}
	return nil
}

// GetQueriesLength returns the Total of queries for all Services
func (s serviceSlice) GetQueriesLength() int {
	count := 0
	for _, service := range s {
		count += service.Inspector.LenQueriesByPlat(service.Parser.Platform)
	}
	return count
}

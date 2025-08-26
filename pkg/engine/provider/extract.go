/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package provider

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"sync"

	"github.com/Checkmarx/kics/pkg/kuberneter"
	"github.com/Checkmarx/kics/pkg/logger"
	"github.com/Checkmarx/kics/pkg/model"
	"github.com/Checkmarx/kics/pkg/utils"

	"github.com/hashicorp/go-getter"
)

const (
	channelLength = 2
)

// ExtractedPath is a struct that contains the paths, temporary paths to remove
// and extraction map path of the sources
// Path is the slice of paths to scan
// ExtractionMap is a map that correlates the temporary path to the given path
// RemoveTmp is the slice containing temporary paths to be removed
type ExtractedPath struct {
	Path          []string
	ExtractionMap map[string]model.ExtractedPathObject
}

type getterStruct struct {
	ctx         context.Context
	cancel      context.CancelFunc
	mode        getter.ClientMode
	pwd         string
	opts        []getter.ClientOption
	destination string
	source      string
}

// GetKuberneterSources uses Kubernetes API to download runtime resources
// After Downloaded files kics scan the files as normal local files
func GetKuberneterSources(ctx context.Context, source []string, destinationPath string) (ExtractedPath, error) {
	logger := logger.FromContext(ctx)
	extrStruct := ExtractedPath{
		Path:          []string{},
		ExtractionMap: make(map[string]model.ExtractedPathObject),
	}

	for _, path := range source {
		exportedPath, err := kuberneter.Import(ctx, path, destinationPath)
		if err != nil {
			logger.Error().Msgf("failed to import %s: %s", path, err)
		}

		extrStruct.ExtractionMap[exportedPath] = model.ExtractedPathObject{
			Path:      exportedPath,
			LocalPath: true,
		}

		extrStruct.Path = append(extrStruct.Path, exportedPath)
	}

	return extrStruct, nil
}

// GetSources goes through the source slice, and determines the of source type (ex: zip, git, local).
// It than extracts the files to be scanned. If the source given is not local, a temp dir
// will be created where the files will be stored.
func GetSources(ctx context.Context, source []string, downloadDir string) (ExtractedPath, error) {
	// logger := logger.FromContext(ctx)
	extrStruct := ExtractedPath{
		Path:          []string{},
		ExtractionMap: make(map[string]model.ExtractedPathObject),
	}
	for _, path := range source {
		destination := filepath.Join(downloadDir, utils.NextRandom())
		// logger.Info().Msgf("Extracting %s to %s", path, destination)
		// mode := getter.ClientModeAny

		// pwd := "external/com_github_checkmarx_kics/"

		// opts := []getter.ClientOption{}

		// opts = append(opts, getter.WithInsecure())

		// ctx, cancel := context.WithCancel(context.Background())

		// goGetter := getterStruct{
		// 	ctx:         ctx,
		// 	cancel:      cancel,
		// 	mode:        mode,
		// 	pwd:         pwd,
		// 	opts:        opts,
		// 	destination: destination,
		// 	source:      path,
		// }

		// getterDst, err := getPaths(&goGetter)
		// if err != nil {
		// 	if ignoreDamagedFiles(path) {
		// 		continue
		// 	}
		// 	logger.Error().Msgf("%s", err)
		// 	return ExtractedPath{}, err
		// }
		// tempDst, local := checkSymLink(getterDst, path)

		extrStruct.ExtractionMap[destination] = model.ExtractedPathObject{
			Path:      path,
			LocalPath: true,
		}

		extrStruct.Path = append(extrStruct.Path, path)
	}
	return extrStruct, nil
}

func getPaths(ctx context.Context, g *getterStruct) (string, error) {
	logger := logger.FromContext(ctx)
	if isEncrypted(ctx, g.source) {
		err := errors.New("zip encrypted files are not supported")
		logger.Err(err)
		return "", err
	}

	// Build the client
	client := &getter.Client{
		Ctx:     g.ctx,
		Src:     g.source,
		Dst:     g.destination,
		Pwd:     g.pwd,
		Mode:    g.mode,
		Options: g.opts,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	errChan := make(chan error, channelLength)
	go func() {
		defer wg.Done()
		defer g.cancel()
		if err := client.Get(); err != nil {
			errChan <- err
		}
	}()

	c := make(chan os.Signal, channelLength)
	signal.Notify(c, os.Interrupt)

	select {
	case <-c:
		signal.Reset(os.Interrupt)
		g.cancel()
		wg.Wait()
	case <-g.ctx.Done():
		wg.Wait()
	case err := <-errChan:
		wg.Wait()
		return "", err
	}

	return g.destination, nil
}

// check if the dst is a symbolic link
func checkSymLink(ctx context.Context, getterDst, pathFile string) (string, bool) {
	logger := logger.FromContext(ctx)
	var local bool
	_, err := os.Stat(pathFile)
	if err == nil { // check if file exist locally
		local = true
	}

	info, err := os.Lstat(getterDst)
	if err != nil {
		logger.Error().Msgf("failed lstat for %s: %v", getterDst, err)
	}

	fileInfo := getFileInfo(info, getterDst, pathFile)

	if info.Mode()&os.ModeSymlink != 0 { // if it's a symbolic Link
		path, err := os.Readlink(getterDst) // get location of symbolic Link
		if err != nil {
			logger.Error().Msgf("failed Readlink for %s: %v", getterDst, err)
		}
		getterDst = path // change path to local path
	} else if !fileInfo.IsDir() { // symbolic links are not created for single files
		if local { // check if file exist locally
			getterDst = pathFile
		}
	}
	return getterDst, local
}

func getFileInfo(info fs.FileInfo, dst, pathFile string) fs.FileInfo {
	var extension = filepath.Ext(pathFile)
	var path string
	if extension == "" {
		path = filepath.Join(dst, filepath.Base(pathFile[0:len(pathFile)-len(extension)])) // for single file
	} else {
		path = filepath.Join(dst, filepath.Base(pathFile)) // for directories
	}
	fileInfo, err := os.Lstat(path)
	if err != nil {
		fileInfo = info
	}
	return fileInfo
}

func isEncrypted(ctx context.Context, sourceFile string) bool {
	logger := logger.FromContext(ctx)
	if filepath.Ext(sourceFile) != ".zip" {
		return false
	}
	logger.Error().Msgf("file %s is zipped", sourceFile)
	return false
}

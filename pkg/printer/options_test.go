/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package printer

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestOptions_LogPath(t *testing.T) {
	type args struct {
		opt string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_log_path",
			args: args{
				"",
			},
			wantErr: false,
		},
		{
			name: "test_log_path_error",
			args: args{
				"kics/results.json",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := LogPath(tt.args.opt, true)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogPath() = %v, wantErr = %v", err, tt.wantErr)
			} else if tt.args.opt != "" {
				require.FileExists(t, filepath.FromSlash(tt.args.opt))
				os.RemoveAll(filepath.Dir(filepath.FromSlash(tt.args.opt)))
			}
		})
	}
}

func TestOptions_CI(t *testing.T) {
	res := os.Stdout
	defer func() {
		os.Stdout = res
	}()
	type args struct {
		opt bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    zerolog.ConsoleWriter
	}{
		{
			name: "test_ci",
			args: args{
				opt: true,
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outConsoleLogger = io.Discard
			outFileLogger = io.Discard
			err := CI(ctx, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("CI() = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestOptions_LogFormat(t *testing.T) {
	type args struct {
		opt string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_log_format_json",
			args: args{
				opt: "json",
			},
			wantErr: false,
		},
		{
			name: "test_log_format_pretty",
			args: args{
				opt: "pretty",
			},
			wantErr: false,
		},
		{
			name: "test_log_format_error",
			args: args{
				opt: "error",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := LogFormat(ctx, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogFormat() = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

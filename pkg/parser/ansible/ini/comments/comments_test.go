/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 *
 * This product includes software developed at Datadog (https://www.datadoghq.com)  Copyright 2024 Datadog, Inc.
 */
package comments

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_getKicsIgnore(t *testing.T) {
	tests := []struct {
		comment string
		want    string
	}{
		{
			comment: "    # NEW comment",
			want:    "    # new comment",
		},
		{
			comment: "# KICS:ignore-line\n",
			want:    "# kics:ignore-line",
		},
	}
	for _, tt := range tests {
		t.Run(tt.comment, func(t *testing.T) {
			if got := getKicsIgnore(tt.comment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_getIgnoreLinesFromBlock(t *testing.T) {
	tests := []struct {
		lines           []string
		ignoreBlockLine int
		want            int
	}{
		{
			lines: []string{
				"; dd-iac-scan ignore-block",
				"ansible_host=0.0.0.0",
			},
			ignoreBlockLine: 0,
			want:            0,
		},
		{
			lines: []string{
				"; dd-iac-scan ignore-block",
				"[group_name]",
			},
			ignoreBlockLine: 0,
			want:            1,
		},
		{
			lines: []string{
				"; dd-iac-scan ignore-block",
				"[group_name]",
				"ansible_host=0.0.0.0",
				"",
				"[another_group]",
			},
			ignoreBlockLine: 0,
			want:            3,
		},
	}

	for i, tt := range tests {
		t.Run("Test number "+strconv.Itoa(i), func(t *testing.T) {
			if got := getIgnoreLinesFromBlock(tt.lines, tt.ignoreBlockLine); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_getIgnoreLines(t *testing.T) {
	tests := []struct {
		lines []string
		want  []int
	}{
		{
			lines: []string{
				"# dd-iac-scan ignore-block",
			},
			want: []int{0},
		},
		{
			lines: []string{
				"# dd-iac-scan ignore-line",
			},
			want: []int{0},
		},
		{
			lines: []string{
				"# dd-iac-scan ignore-block",
				"ansible_host=0.0.0.0",
			},
			want: []int{0},
		},
		{
			lines: []string{
				"; dd-iac-scan ignore-block",
				"[group_name]",
			},
			want: []int{0, 1},
		},
		{
			lines: []string{
				"; dd-iac-scan ignore-block",
				"[group_name]",
				"ansible_host=0.0.0.0",
				"",
				"[another_group]",
			},
			want: []int{0, 1, 2, 3},
		},
		{
			lines: []string{
				"; dd-iac-scan ignore-line",
				"ansible_host=0.0.0.0",
			},
			want: []int{0, 1},
		},
	}

	for i, tt := range tests {
		t.Run("Test number "+strconv.Itoa(i), func(t *testing.T) {
			if got := GetIgnoreLines(tt.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %d, want %d", got, tt.want)
			}
		})
	}
}

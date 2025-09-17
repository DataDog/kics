package terraform

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateSubstrings(t *testing.T) {
	lines := []string{
		`tags = ["a", "b", "c"]`,
		`description = "hello"`,
		`name = "web"`,
		`labels = { Name = "db" }`,
	}

	tests := []struct {
		name        string
		key         string
		extracted   [][]string
		currentLine int
		wantSubstr1 string
		wantSubstr2 string
	}{
		{
			name:        "list index resolves to value",
			key:         `tags[1]`,
			currentLine: 0,
			wantSubstr1: "tags",
			wantSubstr2: "b",
		},
		{
			name:        "numeric bracket but not a list -> block navigation (no value)",
			key:         `ingress[0]`,
			currentLine: 0,
			wantSubstr1: "ingress",
			wantSubstr2: "",
		},
		{
			name:        "map-style quoted key",
			key:         `labels["Name"]`,
			currentLine: 0,
			wantSubstr1: "labels",
			wantSubstr2: "Name",
		},
		{
			name:        "unquoted label key",
			key:         `resource[web]`,
			currentLine: 0,
			wantSubstr1: "resource",
			wantSubstr2: "web",
		},
		{
			name:        "placeholder inside brackets gets restored",
			key:         `labels[{{0}}]`,
			extracted:   [][]string{{"Environment"}},
			currentLine: 0,
			wantSubstr1: "labels",
			wantSubstr2: "Environment",
		},
		{
			name:        "plain key=value line",
			key:         `description = "hello"`,
			currentLine: 0,
			wantSubstr1: "description",
			// Note: function preserves quotes from the right-hand side
			wantSubstr2: `"hello"`,
		},
		{
			name:        "fallback scan resolves value later in lines",
			key:         `name`,
			currentLine: 0,
			wantSubstr1: "name",
			wantSubstr2: `"web"`,
		},
		{
			name:        "numeric placeholder but not list -> treated as block navigation (no value)",
			key:         `block[{{0}}]`,
			extracted:   [][]string{{"0"}},
			currentLine: 0,
			wantSubstr1: "block",
			wantSubstr2: "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := GenerateSubstrings(context.TODO(), tt.key, tt.extracted, lines, tt.currentLine)
			require.Equal(t, got1, tt.wantSubstr1)
			require.Equal(t, got2, tt.wantSubstr2)
		})
	}
}

func TestLooksLikeListAttribute(t *testing.T) {
	lines := []string{
		`tags = ["a", "b", "c"]`,
		`ingress {}`,
	}

	require.Equal(t, looksLikeListAttribute("tags", lines), true)

	require.Equal(t, looksLikeListAttribute("ingress", lines), false)
}

func TestResolveListIndex(t *testing.T) {
	lines := []string{
		`tags = ["a", "b", "c"]`,
		`names = ["web"]`,
	}

	require.Equal(t, resolveListIndex("tags", 2, lines), "c")

	require.Equal(t, resolveListIndex("tags", 5, lines), "")

	require.Equal(t, resolveListIndex("names", 0, lines), "web")
}

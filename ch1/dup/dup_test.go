package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReportDuplicates(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{
			name: "one duplicate pair",
			in:   "a\nb\na\n",
			want: []string{"2\ta"},
		},
		{
			name: "multiple duplicates",
			in:   "x\ny\nx\ny\nx\n",
			want: []string{"3\tx", "2\ty"},
		},
		{
			name: "unique only",
			in:   "only\nonce\n",
			want: nil,
		},
		{
			name: "empty input",
			in:   "",
			want: nil,
		},
		{
			name: "empty duplicated lines",
			in:   "\n\n",
			want: []string{"2\t"},
		},
		{
			name: "whitespace sensitive identity",
			in:   " hi\nhi\n hi\n",
			want: []string{"2\t hi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			require.NoError(t, reportDuplicates(strings.NewReader(tt.in), &out))
			assert.ElementsMatch(t, tt.want, nonEmptyLines(out.String()))
		})
	}
}

func nonEmptyLines(s string) []string {
	if s == "" {
		return nil
	}
	var lines []string
	for _, line := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
		lines = append(lines, line)
	}
	return lines
}

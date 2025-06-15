package utils

import (
	"slices"
	"testing"
)

func TestJoinStringSeq(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		separator string
		expected  string
	}{
		{
			name:      "empty sequence",
			input:     []string{},
			separator: ",",
			expected:  "",
		},
		{
			name:      "single element",
			input:     []string{"a"},
			separator: ",",
			expected:  "a",
		},
		{
			name:      "multiple elements",
			input:     []string{"a", "b", "c"},
			separator: ",",
			expected:  "a,b,c",
		},
		{
			name:      "different separator",
			input:     []string{"a", "b", "c"},
			separator: "-",
			expected:  "a-b-c",
		},
		{
			name:      "elements with empty string",
			input:     []string{"a", "", "c"},
			separator: ",",
			expected:  "a,,c",
		},
		{
			name:      "empty separator",
			input:     []string{"a", "b", "c"},
			separator: "",
			expected:  "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := slices.Values(tt.input)
			result := JoinStringSeq(seq, tt.separator)
			if result != tt.expected {
				t.Errorf("JoinStringSeq(%v, %q) = %q; want %q", tt.input, tt.separator, result, tt.expected)
			}
		})
	}
}
func TestJoinStringSlice(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		separator string
		expected  string
	}{
		{
			name:      "empty slice",
			input:     []string{},
			separator: ",",
			expected:  "",
		},
		{
			name:      "single element",
			input:     []string{"a"},
			separator: ",",
			expected:  "a",
		},
		{
			name:      "multiple elements",
			input:     []string{"a", "b", "c"},
			separator: ",",
			expected:  "a,b,c",
		},
		{
			name:      "different separator",
			input:     []string{"a", "b", "c"},
			separator: "-",
			expected:  "a-b-c",
		},
		{
			name:      "elements with empty string",
			input:     []string{"a", "", "c"},
			separator: ",",
			expected:  "a,,c",
		},
		{
			name:      "empty separator",
			input:     []string{"a", "b", "c"},
			separator: "",
			expected:  "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := JoinStringSlice(tt.input, tt.separator)
			if result != tt.expected {
				t.Errorf("JoinStringSlice(%v, %q) = %q; want %q", tt.input, tt.separator, result, tt.expected)
			}
		})
	}
}

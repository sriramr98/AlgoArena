package evaluators

import (
	"testing"

	"github.com/sriramr98/dsa_server/problems"
)

func TestStringEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		name           string
		output         string
		expected       interface{}
		comparisonMode problems.ComparisonMode
		wantPassed     bool
		wantErr        bool
	}{
		{
			name:           "exact string match",
			output:         "hello world",
			expected:       "hello world",
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "string mismatch",
			output:         "hello world",
			expected:       "hello",
			comparisonMode: problems.ExactMode,
			wantPassed:     false,
			wantErr:        false,
		},
		{
			name:           "string with trailing newline",
			output:         "hello world\n",
			expected:       "hello world",
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "string with trailing whitespace",
			output:         "hello world   ",
			expected:       "hello world",
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "empty string match",
			output:         "",
			expected:       "",
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "expected type mismatch",
			output:         "hello",
			expected:       123, // int instead of string
			comparisonMode: problems.ExactMode,
			wantPassed:     false,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCase := problems.TestCase{
				Expected: tt.expected,
			}

			result, err := EvaluateString(tt.output, testCase, tt.comparisonMode)

			if (err != nil) != tt.wantErr {
				t.Errorf("StringEvaluator.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && result.Passed != tt.wantPassed {
				t.Errorf("StringEvaluator.Evaluate() passed = %v, want %v", result.Passed, tt.wantPassed)
			}
		})
	}
}

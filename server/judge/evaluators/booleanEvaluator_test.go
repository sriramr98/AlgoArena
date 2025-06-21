package evaluators

import (
	"testing"

	"github.com/sriramr98/dsa_server/problems"
)

func TestBooleanEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		name           string
		output         string
		expected       interface{}
		comparisonMode problems.ComparisonMode
		wantPassed     bool
		wantErr        bool
	}{
		{
			name:           "true matches true",
			output:         "true",
			expected:       true,
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "false matches false",
			output:         "false",
			expected:       false,
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "true does not match false",
			output:         "true",
			expected:       false,
			comparisonMode: problems.ExactMode,
			wantPassed:     false,
			wantErr:        false,
		},
		{
			name:           "FALSE in caps should match false",
			output:         "FALSE",
			expected:       false,
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "invalid boolean causes error",
			output:         "not_a_boolean",
			expected:       true,
			comparisonMode: problems.ExactMode,
			wantPassed:     false,
			wantErr:        true,
		},
		{
			name:           "expected type mismatch",
			output:         "true",
			expected:       "true", // string instead of bool
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

			result, err := EvaluateBool(tt.output, testCase, tt.comparisonMode)

			if (err != nil) != tt.wantErr {
				t.Errorf("BooleanEvaluator.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && result.Passed != tt.wantPassed {
				t.Errorf("BooleanEvaluator.Evaluate() passed = %v, want %v", result.Passed, tt.wantPassed)
			}
		})
	}
}

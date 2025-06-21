package evaluators

import (
	"testing"

	"github.com/sriramr98/dsa_server/problems"
)

func TestFloatEvaluator_Evaluate(t *testing.T) {
	evaluator := FloatEvaluator{}

	tests := []struct {
		name           string
		output         string
		expected       interface{}
		comparisonMode problems.ComparisonMode
		wantPassed     bool
		wantErr        bool
	}{
		{
			name:           "exact float match",
			output:         "3.14",
			expected:       3.14,
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "float mismatch",
			output:         "3.14",
			expected:       3.15,
			comparisonMode: problems.ExactMode,
			wantPassed:     false,
			wantErr:        false,
		},
		{
			name:           "integer as float",
			output:         "5",
			expected:       5.0,
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "negative float",
			output:         "-2.5",
			expected:       -2.5,
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "scientific notation",
			output:         "1.23e-4",
			expected:       0.000123,
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "almost equal floats (within epsilon)",
			output:         "0.10000000000000001",
			expected:       0.1,
			comparisonMode: problems.ExactMode,
			wantPassed:     true,
			wantErr:        false,
		},
		{
			name:           "invalid float",
			output:         "not_a_float",
			expected:       1.0,
			comparisonMode: problems.ExactMode,
			wantPassed:     false,
			wantErr:        true,
		},
		{
			name:           "expected type mismatch",
			output:         "3.14",
			expected:       "3.14", // string instead of float
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

			result, err := evaluator.Evaluate(tt.output, testCase, tt.comparisonMode)

			if (err != nil) != tt.wantErr {
				t.Errorf("FloatEvaluator.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && result.Passed != tt.wantPassed {
				t.Errorf("FloatEvaluator.Evaluate() passed = %v, want %v", result.Passed, tt.wantPassed)
			}
		})
	}
}

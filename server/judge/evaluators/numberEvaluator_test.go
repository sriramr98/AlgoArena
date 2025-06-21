package evaluators

import (
	"errors"
	"testing"

	"github.com/sriramr98/dsa_server/problems"
	"github.com/stretchr/testify/assert"
)

func TestIntegerEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		name           string
		executorOutput string
		testCase       problems.TestCase
		comparisonMode problems.ComparisonMode
		want           EvaluatorResult
		wantErr        bool
		expectedErr    error
	}{
		{
			name:           "successful evaluation with matching values",
			executorOutput: "42",
			testCase: problems.TestCase{
				Expected: 42,
			},
			want: EvaluatorResult{
				Passed:         true,
				ActualOutput:   42,
				ExpectedOutput: 42,
				Error:          nil,
			},
			wantErr: false,
		},
		{
			name:           "successful evaluation with non-matching values",
			executorOutput: "43",
			testCase: problems.TestCase{
				Expected: 42,
			},
			want: EvaluatorResult{
				Passed:         false,
				ActualOutput:   43,
				ExpectedOutput: 42,
				Error:          nil,
			},
			wantErr: false,
		},
		{
			name:           "successful evaluation with trailing newline",
			executorOutput: "42\n",
			testCase: problems.TestCase{
				Expected: 42,
			},
			want: EvaluatorResult{
				Passed:         true,
				ActualOutput:   42,
				ExpectedOutput: 42,
				Error:          nil,
			},
			wantErr: false,
		},
		{
			name:           "invalid stdout - not a number",
			executorOutput: "not a number",
			testCase: problems.TestCase{
				Expected: 42,
			},
			want:    EvaluatorResult{},
			wantErr: true,
		},
		{
			name:           "expected type is not int",
			executorOutput: "42",
			testCase: problems.TestCase{
				Expected: "42", // string instead of int
			},
			want:        EvaluatorResult{},
			wantErr:     true,
			expectedErr: errors.New("expected data type doesn't match input data type"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := tt.executorOutput
			got, err := EvaluateInteger(stdout, tt.testCase, tt.comparisonMode)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedErr != nil {
					assert.Equal(t, tt.expectedErr.Error(), err.Error())
				}
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

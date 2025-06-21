package evaluators

import (
	"errors"
	"strings"

	"github.com/sriramr98/dsa_server/problems"
)

func EvaluateString(output string, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error) {
	// Trim any trailing newlines or whitespace that might be added by the execution environment
	expected, ok := testCase.Expected.(string)
	if !ok {
		return EvaluatorResult{}, errors.New("expected data type doesn't match input data type")
	}

	areEquals := expected == strings.TrimSpace(output)

	return EvaluatorResult{
		Passed:         areEquals,
		ActualOutput:   output,
		ExpectedOutput: expected,
		Error:          nil,
	}, nil
}

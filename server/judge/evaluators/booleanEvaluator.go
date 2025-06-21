package evaluators

import (
	"errors"
	"strconv"
	"strings"

	"github.com/sriramr98/dsa_server/problems"
)

func EvaluateBool(output string, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error) {
	// Trim any whitespace and convert to lowercase for comparison
	actualResultStr := strings.ToLower(strings.TrimSpace(output))

	// Parse the boolean value from the string output
	actualResult, err := strconv.ParseBool(actualResultStr)
	if err != nil {
		return EvaluatorResult{}, err
	}

	expected, ok := testCase.Expected.(bool)
	if !ok {
		return EvaluatorResult{}, errors.New("expected data type doesn't match input data type")
	}

	areEquals := expected == actualResult

	return EvaluatorResult{
		Passed:         areEquals,
		ActualOutput:   actualResult,
		ExpectedOutput: expected,
		Error:          nil,
	}, nil
}

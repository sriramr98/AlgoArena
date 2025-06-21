package evaluators

import (
	"errors"
	"strconv"
	"strings"

	"github.com/sriramr98/dsa_server/problems"
)

func EvaluateFloat(output string, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error) {
	// Convert the string to a float64
	actualResult, err := strconv.ParseFloat(strings.TrimSpace(output), 64)
	if err != nil {
		return EvaluatorResult{}, err
	}

	expected, ok := testCase.Expected.(float64)
	if !ok {
		return EvaluatorResult{}, errors.New("expected value must be a float64")
	}

	// For floating point comparison, using exact equality might be problematic
	// due to floating point precision issues. Consider using an epsilon-based comparison.
	// Ref: https://stackoverflow.com/a/47969546
	const epsilon = 1e-9
	diff := expected - actualResult
	areEquals := diff < epsilon && diff > -epsilon // This checks if the difference is within a small range

	return EvaluatorResult{
		Passed:         areEquals,
		ActualOutput:   actualResult,
		ExpectedOutput: expected,
		Error:          nil,
	}, nil
}

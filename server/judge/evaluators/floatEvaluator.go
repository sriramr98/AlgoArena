package evaluators

import (
	"errors"
	"strconv"
	"strings"

	"github.com/sriramr98/dsa_server/judge/executors"
	"github.com/sriramr98/dsa_server/problems"
)

type FloatEvaluator struct{}

func (fe FloatEvaluator) Evaluate(executionResult executors.ExecutorOutput, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error) {
	// Trim any whitespace from the output
	trimmedOutput := strings.TrimSpace(executionResult.Run.Stdout)

	// Convert the string to a float64
	actualResult, err := strconv.ParseFloat(trimmedOutput, 64)
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

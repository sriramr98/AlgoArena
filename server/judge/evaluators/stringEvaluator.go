package evaluators

import (
	"errors"
	"strings"

	"github.com/sriramr98/dsa_server/judge/executors"
	"github.com/sriramr98/dsa_server/problems"
)

type StringEvaluator struct{}

func (se StringEvaluator) Evaluate(executionResult executors.ExecutorOutput, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error) {
	// Trim any trailing newlines or whitespace that might be added by the execution environment
	actualResult := strings.TrimSpace(executionResult.Run.Stdout)

	expected, ok := testCase.Expected.(string)
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

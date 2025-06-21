package evaluators

import (
	"errors"
	"strconv"

	"github.com/sriramr98/dsa_server/judge/executors"
	"github.com/sriramr98/dsa_server/problems"
)

type IntegerEvaluator struct{}

func (ne IntegerEvaluator) Evaluate(executionResult executors.ExecutorOutput, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error) {
	actualResult, err := strconv.Atoi(executionResult.Run.Stdout)
	if err != nil {
		return EvaluatorResult{}, err
	}

	expected, ok := testCase.Expected.(int)
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

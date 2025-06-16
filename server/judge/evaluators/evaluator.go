package evaluators

import (
	"fmt"

	"github.com/sriramr98/dsa_server/judge/executors"
	"github.com/sriramr98/dsa_server/problems"
)

type EvaluatorResult struct {
	Passed         bool
	ActualOutput   any
	ExpectedOutput any
	Error          error
}

type Evaluator interface {
	Evaluate(executionResult executors.ExecutorOutput, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error)
}

func GetEvaluator(problem problems.Problem) (Evaluator, error) {
	var defaultEval Evaluator
	switch problem.Output.Type {
	case problems.ArrayType:
		return ArrayEvaluator{}, nil
	default:
		return defaultEval, fmt.Errorf("no evaluator found for output type %s", problem.Output.Type)
	}
}

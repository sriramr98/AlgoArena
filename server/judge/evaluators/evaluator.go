package evaluators

import (
	"fmt"

	"github.com/sriramr98/dsa_server/problems"
)

type EvaluatorResult struct {
	Passed         bool
	ActualOutput   any
	ExpectedOutput any
	Error          error
}

type Evaluator func(output string, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error)

func GetEvaluator(problem problems.Problem) (Evaluator, error) {
	var defaultEval Evaluator
	switch problem.Output.Type {
	case problems.ArrayType:
		return EvaluateArray, nil
	case problems.NumberType:
		return EvaluateInteger, nil
	case problems.FloatType:
		return EvaluateFloat, nil
	case problems.StringType:
		return EvaluateString, nil
	case problems.BooleanType:
		return EvaluateBool, nil
	default:
		return defaultEval, fmt.Errorf("no evaluator found for output type %s", problem.Output.Type)
	}
}

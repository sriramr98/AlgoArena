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

type Evaluator interface {
	Evaluate(output string, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error)
}

func GetEvaluator(problem problems.Problem) (Evaluator, error) {
	var defaultEval Evaluator
	switch problem.Output.Type {
	case problems.ArrayType:
		return ArrayEvaluator{}, nil
	case problems.NumberType:
		return IntegerEvaluator{}, nil
	case problems.FloatType:
		return FloatEvaluator{}, nil
	case problems.StringType:
		return StringEvaluator{}, nil
	case problems.BooleanType:
		return BooleanEvaluator{}, nil
	default:
		return defaultEval, fmt.Errorf("no evaluator found for output type %s", problem.Output.Type)
	}
}

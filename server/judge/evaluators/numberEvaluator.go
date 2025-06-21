package evaluators

import (
	"errors"
	"strconv"
	"strings"

	"github.com/sriramr98/dsa_server/problems"
)

func EvaluateInteger(output string, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error) {
	actualResult, err := strconv.Atoi(strings.TrimSpace(output))
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

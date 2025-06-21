package evaluators

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sriramr98/dsa_server/problems"
)

func EvaluateInteger(output string, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error) {
	actualResult, err := strconv.Atoi(strings.TrimSpace(output))
	if err != nil {
		return EvaluatorResult{}, err
	}

	// We first cast to float and then to it because this comes from a json input which always maps numbers to float64
	expected, ok := testCase.Expected.(float64)

	if !ok {
		return EvaluatorResult{}, fmt.Errorf("expected data type doesn't match input data type")
	}

	areEquals := int(expected) == actualResult

	return EvaluatorResult{
		Passed:         areEquals,
		ActualOutput:   actualResult,
		ExpectedOutput: int(expected),
		Error:          nil,
	}, nil
}

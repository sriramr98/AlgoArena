package evaluators

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/sriramr98/dsa_server/problems"
)

func EvaluateArray(output string, testCase problems.TestCase, comparisonMode problems.ComparisonMode) (EvaluatorResult, error) {
	actualResult, err := format(strings.TrimSpace(output))
	if err != nil {
		return EvaluatorResult{}, err
	}

	expected, ok := testCase.Expected.([]any)
	if !ok {
		return EvaluatorResult{}, errors.New("expected data type doesn't match input data type")
	}

	areEquals := false
	if comparisonMode == problems.OrderedMode {
		areEquals = compareOrdered(expected, actualResult)
	} else {
		areEquals = compareUnordered(expected, actualResult)
	}

	return EvaluatorResult{
		Passed:         areEquals,
		ActualOutput:   actualResult,
		ExpectedOutput: expected,
		Error:          nil,
	}, nil
}

func format(data string) ([]any, error) {
	var result []any
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return result, err
	}

	return result, nil
}

func compareOrdered(expected []any, actual []any) bool {
	if len(expected) != len(actual) {
		return false
	}

	for idx, val := range expected {
		if actual[idx] != val {
			return false
		}
	}

	return true
}

func compareUnordered(expected []any, actual []any) bool {
	if len(expected) != len(actual) {
		return false
	}

	occurenceMap := make(map[any]int)
	for _, val := range expected {
		occurenceMap[val] += 1
	}

	for _, val := range actual {
		occurenceMap[val] -= 1
		if occurenceMap[val] < 0 {
			return false
		}
	}

	return true
}

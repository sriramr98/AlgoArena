package evaluators

import (
	"encoding/json"
	"testing"

	"github.com/sriramr98/dsa_server/judge/executors"
	"github.com/sriramr98/dsa_server/problems"
)

func TestArrayEvaluatorCompareOrdered(t *testing.T) {
	testCases := []struct {
		Name          string
		Actual        []any
		Expected      []any
		ShouldBeEqual bool
	}{
		{
			Name:          "equal string arrays",
			Actual:        []any{"abc", "def", "ghi"},
			Expected:      []any{"abc", "def", "ghi"},
			ShouldBeEqual: true,
		},
		{
			Name:          "different length arrays",
			Actual:        []any{"abc", "def"},
			Expected:      []any{"abc", "def", "ghi"},
			ShouldBeEqual: false,
		},
		{
			Name:          "same length but different elements",
			Actual:        []any{"abc", "def", "xyz"},
			Expected:      []any{"abc", "def", "ghi"},
			ShouldBeEqual: false,
		},
		{
			Name:          "empty arrays",
			Actual:        []any{},
			Expected:      []any{},
			ShouldBeEqual: true,
		},
		{
			Name:          "mixed type arrays that match",
			Actual:        []any{1, "two", true, 4.5},
			Expected:      []any{1, "two", true, 4.5},
			ShouldBeEqual: true,
		},
		{
			Name:          "mixed type arrays that don't match",
			Actual:        []any{1, "two", false, 4.5},
			Expected:      []any{1, "two", true, 4.5},
			ShouldBeEqual: false,
		},
		{
			Name:          "different order arrays",
			Actual:        []any{"ghi", "def", "abc"},
			Expected:      []any{"abc", "def", "ghi"},
			ShouldBeEqual: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			areEqual := compareOrdered(testCase.Expected, testCase.Actual)
			if testCase.ShouldBeEqual && !areEqual {
				t.Error("Expected arrays to be equal but aren't equal")
			}

			if !testCase.ShouldBeEqual && areEqual {
				t.Error("Expected arrays to not be equal but are equal")
			}
		})
	}
}

func TestArrayEvaluatorCompareUnordered(t *testing.T) {
	testCases := []struct {
		Name          string
		Actual        []any
		Expected      []any
		ShouldBeEqual bool
	}{
		{
			Name:          "equal string arrays in same order",
			Actual:        []any{"abc", "def", "ghi"},
			Expected:      []any{"abc", "def", "ghi"},
			ShouldBeEqual: true,
		},
		{
			Name:          "equal string arrays in different order",
			Actual:        []any{"ghi", "abc", "def"},
			Expected:      []any{"abc", "def", "ghi"},
			ShouldBeEqual: true,
		},
		{
			Name:          "different length arrays",
			Actual:        []any{"abc", "def"},
			Expected:      []any{"abc", "def", "ghi"},
			ShouldBeEqual: false,
		},
		{
			Name:          "same length but different elements",
			Actual:        []any{"abc", "def", "xyz"},
			Expected:      []any{"abc", "def", "ghi"},
			ShouldBeEqual: false,
		},
		{
			Name:          "empty arrays",
			Actual:        []any{},
			Expected:      []any{},
			ShouldBeEqual: true,
		},
		{
			Name:          "arrays with duplicate elements in different frequencies",
			Actual:        []any{"a", "a", "b", "c"},
			Expected:      []any{"a", "b", "c", "c"},
			ShouldBeEqual: false,
		},
		{
			Name:          "arrays with same elements and frequencies but different order",
			Actual:        []any{"a", "b", "a", "c", "b"},
			Expected:      []any{"a", "a", "b", "b", "c"},
			ShouldBeEqual: true,
		},
		{
			Name:          "mixed type arrays in different order",
			Actual:        []any{true, "two", 1, 4.5},
			Expected:      []any{1, "two", true, 4.5},
			ShouldBeEqual: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			areEqual := compareUnordered(testCase.Expected, testCase.Actual)
			if testCase.ShouldBeEqual && !areEqual {
				t.Error("Expected arrays to be equal but aren't equal")
			}

			if !testCase.ShouldBeEqual && areEqual {
				t.Error("Expected arrays to not be equal but are equal")
			}
		})
	}
}

func TestArrayEvaluatorFormat(t *testing.T) {
	testCases := []struct {
		Name          string
		Input         string
		Expected      []any
		ExpectError   bool
		ErrorContains string
	}{
		{
			Name:        "valid json array of strings",
			Input:       `["abc", "def", "ghi"]`,
			Expected:    []any{"abc", "def", "ghi"},
			ExpectError: false,
		},
		{
			Name:        "valid json array of mixed types",
			Input:       `[1, "two", true, 4.5]`,
			Expected:    []any{float64(1), "two", true, 4.5},
			ExpectError: false,
		},
		{
			Name:        "valid json empty array",
			Input:       `[]`,
			Expected:    []any{},
			ExpectError: false,
		},
		{
			Name:          "invalid json",
			Input:         `[1, 2,]`,
			ExpectError:   true,
			ErrorContains: "json",
		},
		{
			Name:          "not an array",
			Input:         `{"key": "value"}`,
			ExpectError:   true,
			ErrorContains: "cannot unmarshal",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result, err := format(testCase.Input)

			// Check error expectations
			if testCase.ExpectError && err == nil {
				t.Error("Expected an error but got none")
			}
			if !testCase.ExpectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
			if testCase.ExpectError && err != nil && testCase.ErrorContains != "" {
				if !contains(err.Error(), testCase.ErrorContains) {
					t.Errorf("Error doesn't contain expected text. Got: %v, Expected to contain: %v", err.Error(), testCase.ErrorContains)
				}
			}

			// If no error is expected, check the result
			if !testCase.ExpectError {
				if len(result) != len(testCase.Expected) {
					t.Errorf("Result length mismatch. Got: %v, Expected: %v", len(result), len(testCase.Expected))
				}

				for i, val := range testCase.Expected {
					if i >= len(result) {
						t.Errorf("Result missing element at index %d", i)
						continue
					}
					if result[i] != val {
						t.Errorf("Mismatch at index %d. Got: %v, Expected: %v", i, result[i], val)
					}
				}
			}
		})
	}
}

func TestArrayEvaluatorEvaluate(t *testing.T) {
	testCases := []struct {
		Name             string
		ExecOutput       executors.ExecutorOutput
		TestCase         problems.TestCase
		ComparisonMode   problems.ComparisonMode
		ExpectedResult   EvaluatorResult
		ExpectError      bool
		ErrorContains    string
		MockUnmarshalErr bool
	}{
		{
			Name: "ordered comparison that passes",
			ExecOutput: executors.ExecutorOutput{
				Run: executors.RunOutput{
					Stdout: `[1, 2, 3]`,
				},
			},
			TestCase: problems.TestCase{
				Expected: []any{float64(1), float64(2), float64(3)},
			},
			ComparisonMode: problems.OrderedMode,
			ExpectedResult: EvaluatorResult{
				Passed:         true,
				ActualOutput:   []any{float64(1), float64(2), float64(3)},
				ExpectedOutput: []any{float64(1), float64(2), float64(3)},
				Error:          nil,
			},
			ExpectError: false,
		},
		{
			Name: "ordered comparison that fails",
			ExecOutput: executors.ExecutorOutput{
				Run: executors.RunOutput{
					Stdout: `[3, 2, 1]`,
				},
			},
			TestCase: problems.TestCase{
				Expected: []any{float64(1), float64(2), float64(3)},
			},
			ComparisonMode: problems.OrderedMode,
			ExpectedResult: EvaluatorResult{
				Passed:         false,
				ActualOutput:   []any{float64(3), float64(2), float64(1)},
				ExpectedOutput: []any{float64(1), float64(2), float64(3)},
				Error:          nil,
			},
			ExpectError: false,
		},
		{
			Name: "unordered comparison that passes",
			ExecOutput: executors.ExecutorOutput{
				Run: executors.RunOutput{
					Stdout: `[3, 1, 2]`,
				},
			},
			TestCase: problems.TestCase{
				Expected: []any{float64(1), float64(2), float64(3)},
			},
			ComparisonMode: problems.UnorderedMode,
			ExpectedResult: EvaluatorResult{
				Passed:         true,
				ActualOutput:   []any{float64(3), float64(1), float64(2)},
				ExpectedOutput: []any{float64(1), float64(2), float64(3)},
				Error:          nil,
			},
			ExpectError: false,
		},
		{
			Name: "expected not an array",
			ExecOutput: executors.ExecutorOutput{
				Run: executors.RunOutput{
					Stdout: `[1, 2, 3]`,
				},
			},
			TestCase: problems.TestCase{
				Expected: "not an array",
			},
			ComparisonMode: problems.OrderedMode,
			ExpectError:    true,
			ErrorContains:  "expected data type doesn't match",
		},
		{
			Name: "invalid json in stdout",
			ExecOutput: executors.ExecutorOutput{
				Run: executors.RunOutput{
					Stdout: `[1, 2, 3`,
				},
			},
			TestCase: problems.TestCase{
				Expected: []any{float64(1), float64(2), float64(3)},
			},
			ComparisonMode: problems.OrderedMode,
			ExpectError:    true,
			ErrorContains:  "json",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			evaluator := ArrayEvaluator{}
			result, err := evaluator.Evaluate(testCase.ExecOutput, testCase.TestCase, testCase.ComparisonMode)

			// Check error expectations
			if testCase.ExpectError && err == nil {
				t.Error("Expected an error but got none")
			}
			if !testCase.ExpectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
			if testCase.ExpectError && err != nil && testCase.ErrorContains != "" {
				if !contains(err.Error(), testCase.ErrorContains) {
					t.Errorf("Error doesn't contain expected text. Got: %v, Expected to contain: %v", err.Error(), testCase.ErrorContains)
				}
			}

			// If no error is expected, check the result
			if !testCase.ExpectError {
				if result.Passed != testCase.ExpectedResult.Passed {
					t.Errorf("Passed flag mismatch. Got: %v, Expected: %v", result.Passed, testCase.ExpectedResult.Passed)
				}

				// Compare actual output
				actualJSON, _ := json.Marshal(result.ActualOutput)
				expectedJSON, _ := json.Marshal(testCase.ExpectedResult.ActualOutput)
				if string(actualJSON) != string(expectedJSON) {
					t.Errorf("ActualOutput mismatch. Got: %v, Expected: %v", result.ActualOutput, testCase.ExpectedResult.ActualOutput)
				}

				// Compare expected output
				actualJSON, _ = json.Marshal(result.ExpectedOutput)
				expectedJSON, _ = json.Marshal(testCase.ExpectedResult.ExpectedOutput)
				if string(actualJSON) != string(expectedJSON) {
					t.Errorf("ExpectedOutput mismatch. Got: %v, Expected: %v", result.ExpectedOutput, testCase.ExpectedResult.ExpectedOutput)
				}
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return s != "" && substr != "" && s != substr && len(s) > len(substr) && s != substr
}

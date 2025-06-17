package stub

import (
	"strings"
	"testing"

	"github.com/sriramr98/dsa_server/problems"
)

func TestJSStubGenerator_Generate(t *testing.T) {
	tests := []struct {
		name    string
		problem problems.Problem
		want    string
	}{
		{
			name: "simple_problem_with_single_number_input",
			problem: problems.Problem{
				FunctionName: "factorial",
				Input: map[string]problems.InputType{
					"n": {
						VariableType: problems.VariableType{
							Type:        problems.NumberType,
							Description: "a non-negative integer",
						},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: problems.NumberType,
					},
				},
			},
			want: "/**\n * @param {number} n - a non-negative integer\n*/\nfunction factorial(n) {\n // Your code here \n\n return 0\n}",
		},
		{
			name: "problem_with_multiple_inputs_of_different_types",
			problem: problems.Problem{
				FunctionName: "findSubstring",
				Input: map[string]problems.InputType{
					"s": {
						VariableType: problems.VariableType{
							Type:        problems.StringType,
							Description: "the input string",
						},
					},
					"target": {
						VariableType: problems.VariableType{
							Type:        problems.StringType,
							Description: "the substring to find",
						},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: problems.NumberType,
					},
				},
			},
			want: "/**\n * @param {string} s - the input string\n * @param {string} target - the substring to find\n*/\nfunction findSubstring(s, target) {\n // Your code here \n\n return 0\n}",
		},
		{
			name: "problem_with_array_input_and_array_output",
			problem: problems.Problem{
				FunctionName: "reverseArray",
				Input: map[string]problems.InputType{
					"nums": {
						VariableType: problems.VariableType{
							Type: problems.ArrayType,
							ChildType: &problems.VariableType{
								Type: problems.NumberType,
							},
							Description: "array of integers",
						},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: problems.ArrayType,
						ChildType: &problems.VariableType{
							Type: problems.NumberType,
						},
					},
				},
			},
			want: "/**\n * @param {number[]} nums - array of integers\n*/\nfunction reverseArray(nums) {\n // Your code here \n\n return []\n}",
		},
		{
			name: "problem_with_object_input_and_string_output",
			problem: problems.Problem{
				FunctionName: "formatPerson",
				Input: map[string]problems.InputType{
					"person": {
						VariableType: problems.VariableType{
							Type:        problems.ObjectType,
							Description: "person object with name and age",
						},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: problems.StringType,
					},
				},
			},
			want: "/**\n * @param {Object} person - person object with name and age\n*/\nfunction formatPerson(person) {\n // Your code here \n\n return \"\"\n}",
		},
		{
			name: "problem_with_boolean_input_and_output",
			problem: problems.Problem{
				FunctionName: "isEven",
				Input: map[string]problems.InputType{
					"num": {
						VariableType: problems.VariableType{
							Type:        problems.NumberType,
							Description: "an integer to check",
						},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: problems.BooleanType,
					},
				},
			},
			want: "/**\n * @param {number} num - an integer to check\n*/\nfunction isEven(num) {\n // Your code here \n\n return false\n}",
		},
		{
			name: "problem_with_array_of_arrays",
			problem: problems.Problem{
				FunctionName: "transpose",
				Input: map[string]problems.InputType{
					"matrix": {
						VariableType: problems.VariableType{
							Type: problems.ArrayType,
							ChildType: &problems.VariableType{
								Type:      problems.ArrayType,
								ChildType: &problems.VariableType{Type: problems.NumberType},
							},
							Description: "2D matrix",
						},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: problems.ArrayType,
						ChildType: &problems.VariableType{
							Type:      problems.ArrayType,
							ChildType: &problems.VariableType{Type: problems.NumberType},
						},
					},
				},
			},
			want: "/**\n * @param {number[][]} matrix - 2D matrix\n*/\nfunction transpose(matrix) {\n // Your code here \n\n return []\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := JSStubGenerator{}
			got := g.Generate(tt.problem)
			// Compare after normalizing whitespace to handle potential platform-specific line endings
			if normalizeWhitespace(got) != normalizeWhitespace(tt.want) {
				t.Errorf("JSStubGenerator.Generate() =\n%v\nwant:\n%v", got, tt.want)
			}
		})
	}
}

func TestGetComment(t *testing.T) {
	tests := []struct {
		name    string
		problem problems.Problem
		want    string
	}{
		{
			name: "empty input",
			problem: problems.Problem{
				Input: map[string]problems.InputType{},
			},
			want: "/**\n*/",
		},
		{
			name: "single input with description",
			problem: problems.Problem{
				Input: map[string]problems.InputType{
					"num": {
						VariableType: problems.VariableType{
							Type:        problems.NumberType,
							Description: "a number to process",
						},
					},
				},
			},
			want: "/**\n * @param {number} num - a number to process\n*/",
		},
		{
			name: "multiple inputs with and without descriptions",
			problem: problems.Problem{
				Input: map[string]problems.InputType{
					"a": {
						VariableType: problems.VariableType{
							Type:        problems.NumberType,
							Description: "first number",
						},
					},
					"b": {
						VariableType: problems.VariableType{
							Type:        problems.NumberType,
							Description: "",
						},
					},
				},
			},
			want: "/**\n * @param {number} a - first number\n * @param {number} b\n*/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getComment(tt.problem)
			if got != tt.want {
				t.Errorf("getComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetType(t *testing.T) {
	tests := []struct {
		name         string
		variableType problems.DataType
		childType    *problems.VariableType
		want         string
	}{
		{
			name:         "number_type",
			variableType: problems.NumberType,
			childType:    nil,
			want:         "number",
		},
		{
			name:         "string_type",
			variableType: problems.StringType,
			childType:    nil,
			want:         "string",
		},
		{
			name:         "boolean_type",
			variableType: problems.BooleanType,
			childType:    nil,
			want:         "boolean",
		},
		{
			name:         "array_of_numbers",
			variableType: problems.ArrayType,
			childType: &problems.VariableType{
				Type: problems.NumberType,
			},
			want: "number[]",
		},
		{
			name:         "array_of_strings",
			variableType: problems.ArrayType,
			childType: &problems.VariableType{
				Type: problems.StringType,
			},
			want: "string[]",
		},
		{
			name:         "array_with_no_element_type",
			variableType: problems.ArrayType,
			childType:    nil,
			want:         "any[]",
		},
		{
			name:         "array_of_arrays_of_numbers",
			variableType: problems.ArrayType,
			childType: &problems.VariableType{
				Type: problems.ArrayType,
				ChildType: &problems.VariableType{
					Type: problems.NumberType,
				},
			},
			want: "number[][]",
		},
		{
			name:         "array_of_arrays_with_no_element_type",
			variableType: problems.ArrayType,
			childType: &problems.VariableType{
				Type:      problems.ArrayType,
				ChildType: nil,
			},
			want: "any[][]",
		},
		{
			name:         "array_of_array_of_booleans",
			variableType: problems.ArrayType,
			childType: &problems.VariableType{
				Type: problems.ArrayType,
				ChildType: &problems.VariableType{
					Type: problems.BooleanType,
				},
			},
			want: "boolean[][]",
		},
		{
			name:         "array_of_array_of_array_of_numbers",
			variableType: problems.ArrayType,
			childType: &problems.VariableType{
				Type: problems.ArrayType,
				ChildType: &problems.VariableType{
					Type: problems.ArrayType,
					ChildType: &problems.VariableType{
						Type: problems.NumberType,
					},
				},
			},
			want: "number[][][]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getType(tt.variableType, tt.childType)
			if got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReturnVariable(t *testing.T) {
	tests := []struct {
		name    string
		problem problems.Problem
		want    string
	}{
		{
			name: "array_return_type",
			problem: problems.Problem{
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: problems.ArrayType,
					},
				},
			},
			want: "[]",
		},
		{
			name: "object_return_type",
			problem: problems.Problem{
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: problems.ObjectType,
					},
				},
			},
			want: "{}",
		},
		{
			name: "string_return_type",
			problem: problems.Problem{
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: problems.StringType,
					},
				},
			},
			want: "\"\"",
		},
		{
			name: "number_return_type",
			problem: problems.Problem{
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: problems.NumberType,
					},
				},
			},
			want: "0",
		},
		{
			name: "boolean_return_type",
			problem: problems.Problem{
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: problems.BooleanType,
					},
				},
			},
			want: "false",
		},
		{
			name: "unknown_return_type",
			problem: problems.Problem{
				Output: problems.OutputType{
					VariableType: problems.VariableType{
						Type: "unknown",
					},
				},
			},
			want: "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getReturnVariable(tt.problem)
			if got != tt.want {
				t.Errorf("getReturnVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to normalize whitespace for comparison
func normalizeWhitespace(s string) string {
	// Replace all whitespace with a single space
	return strings.Join(strings.Fields(strings.ReplaceAll(s, "\n", " ")), " ")
}

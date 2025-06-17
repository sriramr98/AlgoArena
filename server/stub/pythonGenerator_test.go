package stub

import (
	"fmt"
	"testing"

	"github.com/sriramr98/dsa_server/problems"
)

func TestPythonStubGenerator_Generate(t *testing.T) {
	tests := []struct {
		name    string
		problem problems.Problem
		want    string
	}{
		{
			name: "ReturnMode with Number Input and Array Output",
			problem: problems.Problem{
				ID:           "two-sum",
				FunctionName: "two_sum",
				Input: map[string]problems.InputType{
					"nums": {
						VariableType: problems.VariableType{Type: problems.ArrayType},
					},
					"target": {
						VariableType: problems.VariableType{Type: problems.NumberType},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{Type: problems.ArrayType},
				},
				ExecutionMode: problems.ReturnMode,
			},
			want: "def two_sum(nums: List[Any], target: int) -> List[Any]:\n    # Your code here\n    pass\n\n",
		},
		{
			name: "ReturnMode with String Input and String Output",
			problem: problems.Problem{
				ID:           "reverse-words",
				FunctionName: "reverse_words",
				Input: map[string]problems.InputType{
					"s": {
						VariableType: problems.VariableType{Type: problems.StringType},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{Type: problems.StringType},
				},
				ExecutionMode: problems.ReturnMode,
			},
			want: "def reverse_words(s: str) -> str:\n    # Your code here\n    pass\n\n",
		},
		{
			name: "InPlaceMode with Array Input",
			problem: problems.Problem{
				ID:           "reverse-string",
				FunctionName: "reverse_string",
				Input: map[string]problems.InputType{
					"s": {
						VariableType: problems.VariableType{Type: problems.ArrayType},
						Output:       true,
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{Type: problems.ArrayType},
				},
				ExecutionMode: problems.InPlaceMode,
			},
			want: "def reverse_string(s: List[Any]):\n    # Your code here\n    pass\n\n",
		},
		{
			name: "ReturnMode with Boolean Input and Output",
			problem: problems.Problem{
				ID:           "check-condition",
				FunctionName: "check_condition",
				Input: map[string]problems.InputType{
					"flag": {
						VariableType: problems.VariableType{Type: problems.BooleanType},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{Type: problems.BooleanType},
				},
				ExecutionMode: problems.ReturnMode,
			},
			want: "def check_condition(flag: bool) -> bool:\n    # Your code here\n    pass\n\n",
		},
		{
			name: "Multiple Inputs with Different Types",
			problem: problems.Problem{
				ID:           "process-data",
				FunctionName: "process_data",
				Input: map[string]problems.InputType{
					"str": {
						VariableType: problems.VariableType{Type: problems.StringType},
					},
					"num": {
						VariableType: problems.VariableType{Type: problems.NumberType},
					},
					"arr": {
						VariableType: problems.VariableType{Type: problems.ArrayType},
					},
					"obj": {
						VariableType: problems.VariableType{Type: problems.ObjectType},
					},
					"boolean": {
						VariableType: problems.VariableType{Type: problems.BooleanType},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{Type: problems.BooleanType},
				},
				ExecutionMode: problems.ReturnMode,
			},
			want: "def process_data(arr: List[Any], boolean: bool, num: int, obj: dict, str: str) -> bool:\n    # Your code here\n    pass\n\n",
		},
		{
			name: "Array with Child Type",
			problem: problems.Problem{
				ID:           "nested-arrays",
				FunctionName: "process_matrix",
				Input: map[string]problems.InputType{
					"matrix": {
						VariableType: problems.VariableType{
							Type: problems.ArrayType,
							ChildType: &problems.VariableType{
								Type: problems.ArrayType,
								ChildType: &problems.VariableType{
									Type: problems.NumberType,
								},
							},
						},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{Type: problems.NumberType},
				},
				ExecutionMode: problems.ReturnMode,
			},
			want: "def process_matrix(matrix: List[List[int]]) -> int:\n    # Your code here\n    pass\n\n",
		},
		{
			name: "Inputs in unsorted order with different types should be sorted in function signature",
			problem: problems.Problem{
				ID:           "unsorted-inputs",
				FunctionName: "handle_unsorted_inputs",
				Input: map[string]problems.InputType{
					"b": {
						VariableType: problems.VariableType{Type: problems.BooleanType},
					},
					"a": {
						VariableType: problems.VariableType{Type: problems.StringType},
					},
					"c": {
						VariableType: problems.VariableType{Type: problems.NumberType},
					},
				},
				Output: problems.OutputType{
					VariableType: problems.VariableType{Type: problems.StringType},
				},
				ExecutionMode: problems.ReturnMode,
			},
			want: "def handle_unsorted_inputs(a: str, b: bool, c: int) -> str:\n    # Your code here\n    pass\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := PythonStubGenerator{}
			got := g.Generate(tt.problem)

			wanted := fmt.Sprintf("from typing import *\n\n%s", tt.want)
			if got != wanted {
				t.Errorf("PythonStubGenerator.Generate() = %v, want %v", got, wanted)
			}
		})
	}
}

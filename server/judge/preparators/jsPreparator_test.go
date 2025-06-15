package preparators

import (
	"strings"
	"testing"

	"github.com/sriramr98/dsa_server/problems"
)

type args struct {
	userCode string
	problem  problems.Problem
	testCase problems.TestCase
}

type expected struct {
	varDeclarations []string
	functionCall    string
	printResult     string
}

func TestJsPreparator_Prepare(t *testing.T) {
	tests := []struct {
		name    string
		p       JsPreparator
		args    args
		want    expected
		wantErr bool
	}{
		{
			name: "ReturnMode with Number Input and Array Output",
			p:    JsPreparator{},
			args: args{
				userCode: "function twoSum(nums, target) {\n  // Implementation\n  return [0, 1];\n}",
				problem: problems.Problem{
					ID:           "two-sum",
					FunctionName: "twoSum",
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
				testCase: problems.TestCase{
					Input: map[string]any{
						"nums":   []any{2, 7, 11, 15},
						"target": 9,
					},
					Expected: []any{0, 1},
				},
			},
			want: expected{
				varDeclarations: []string{
					"let nums = [2,7,11,15];",
					"let target = 9;",
				},
				functionCall: "let result = twoSum(nums, target)",
				printResult:  "console.log(JSON.stringify(result));",
			},
			wantErr: false,
		},
		{
			name: "InPlaceMode with Array Input and Output",
			p:    JsPreparator{},
			args: args{
				userCode: "function reverseString(s) {\n  // Implementation\n  return s.reverse();\n}",
				problem: problems.Problem{
					ID:           "reverse-string",
					FunctionName: "reverseString",
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
				testCase: problems.TestCase{
					Input: map[string]any{
						"s": []any{"h", "e", "l", "l", "o"},
					},
					Expected: []any{"o", "l", "l", "e", "h"},
				},
			},
			wantErr: false,
			want: expected{
				varDeclarations: []string{
					"let s = [\"h\",\"e\",\"l\",\"l\",\"o\"];",
				},
				functionCall: "reverseString(s)",
				printResult:  "console.log(JSON.stringify(s));",
			},
		},
		{
			name: "ReturnMode with String Input and String Output",
			p:    JsPreparator{},
			args: args{
				userCode: "function reverseWords(str) {\n  // Implementation\n  return str.split(' ').reverse().join(' ');\n}",
				problem: problems.Problem{
					ID:           "reverse-words",
					FunctionName: "reverseWords",
					Input: map[string]problems.InputType{
						"str": {
							VariableType: problems.VariableType{Type: problems.StringType},
						},
					},
					Output: problems.OutputType{
						VariableType: problems.VariableType{Type: problems.StringType},
					},
					ExecutionMode: problems.ReturnMode,
				},
				testCase: problems.TestCase{
					Input: map[string]any{
						"str": "hello world",
					},
					Expected: "world hello",
				},
			},
			wantErr: false,
			want: expected{
				varDeclarations: []string{
					"let str = \"hello world\";",
				},
				functionCall: "let result = reverseWords(str)",
				printResult:  "console.log(result);",
			},
		},
		{
			name: "ReturnMode with Object Input and Object Output",
			p:    JsPreparator{},
			args: args{
				userCode: "function transformObject(obj) {\n  // Implementation\n  return { ...obj, transformed: true };\n}",
				problem: problems.Problem{
					ID:           "transform-object",
					FunctionName: "transformObject",
					Input: map[string]problems.InputType{
						"obj": {
							VariableType: problems.VariableType{Type: problems.ObjectType},
						},
					},
					Output: problems.OutputType{
						VariableType: problems.VariableType{Type: problems.ObjectType},
					},
					ExecutionMode: problems.ReturnMode,
				},
				testCase: problems.TestCase{
					Input: map[string]any{
						"obj": map[string]any{"name": "test", "value": 42},
					},
					Expected: map[string]any{"name": "test", "value": 42, "transformed": true},
				},
			},
			wantErr: false,
			want: expected{
				varDeclarations: []string{
					"let obj = {\"name\":\"test\",\"value\":42};",
				},
				functionCall: "let result = transformObject(obj)",
				printResult:  "console.log(JSON.stringify(result));",
			},
		},
		{
			name: "InPlaceMode with Object Input and Output",
			p:    JsPreparator{},
			args: args{
				userCode: "function addProperty(obj) {\n  // Implementation\n  obj.modified = true;\n}",
				problem: problems.Problem{
					ID:           "add-property",
					FunctionName: "addProperty",
					Input: map[string]problems.InputType{
						"obj": {
							VariableType: problems.VariableType{Type: problems.ObjectType},
							Output:       true,
						},
					},
					Output: problems.OutputType{
						VariableType: problems.VariableType{Type: problems.ObjectType},
					},
					ExecutionMode: problems.InPlaceMode,
				},
				testCase: problems.TestCase{
					Input: map[string]any{
						"obj": map[string]any{"initial": true},
					},
					Expected: map[string]any{"initial": true, "modified": true},
				},
			},
			wantErr: false,
			want: expected{
				varDeclarations: []string{
					"let obj = {\"initial\":true};",
				},
				functionCall: "addProperty(obj)",
				printResult:  "console.log(JSON.stringify(obj));",
			},
		},
		{
			name: "Multiple Inputs with Different Types",
			p:    JsPreparator{},
			args: args{
				userCode: "function processData(str, num, arr, obj) {\n  // Implementation\n  return { processed: true };\n}",
				problem: problems.Problem{
					ID:           "process-data",
					FunctionName: "processData",
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
						"floats": {
							VariableType: problems.VariableType{Type: problems.FloatType},
						},
					},
					Output: problems.OutputType{
						VariableType: problems.VariableType{Type: problems.ObjectType},
					},
					ExecutionMode: problems.ReturnMode,
				},
				testCase: problems.TestCase{
					Input: map[string]any{
						"str":     "test",
						"num":     42,
						"arr":     []any{1, 2, 3},
						"obj":     map[string]any{"key": "value"},
						"boolean": true,
						"floats":  3.14,
					},
					Expected: map[string]any{"processed": true},
				},
			},
			wantErr: false,
			want: expected{
				varDeclarations: []string{
					"let str = \"test\";",
					"let num = 42;",
					"let arr = [1,2,3];",
					"let obj = {\"key\":\"value\"};",
					"let boolean = true;",
					"let floats = 3.14;",
				},
				functionCall: "let result = processData(str, num, arr, obj, boolean floats)",
				printResult:  "console.log(JSON.stringify(result));",
			},
		},
		{
			name: "Boolean Input and Output",
			p:    JsPreparator{},
			args: args{
				userCode: "function checkCondition(flag) {\n  // Implementation\n  return !flag;\n}",
				problem: problems.Problem{
					ID:           "check-condition",
					FunctionName: "checkCondition",
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
				testCase: problems.TestCase{
					Input: map[string]any{
						"flag": true,
					},
					Expected: false,
				},
			},
			wantErr: false,
			want: expected{
				varDeclarations: []string{
					"let flag = true;",
				},
				functionCall: "let result = checkCondition(flag)",
				printResult:  "console.log(result);",
			},
		},
		{
			name: "Error: InPlaceMode with No Output Variable",
			p:    JsPreparator{},
			args: args{
				userCode: "function noOutputFunc(data) {\n  // Implementation\n}",
				problem: problems.Problem{
					ID:           "no-output",
					FunctionName: "noOutputFunc",
					Input: map[string]problems.InputType{
						"data": {
							VariableType: problems.VariableType{Type: problems.ObjectType},
							Output:       false, // No output variable specified
						},
					},
					Output: problems.OutputType{
						VariableType: problems.VariableType{Type: problems.ObjectType},
					},
					ExecutionMode: problems.InPlaceMode,
				},
				testCase: problems.TestCase{
					Input: map[string]any{
						"data": map[string]any{},
					},
					Expected: map[string]any{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := JsPreparator{}
			got, err := p.Prepare(tt.args.userCode, tt.args.problem, tt.args.testCase)
			if tt.wantErr {
				if err == nil {
					t.Errorf("JsPreparator.Prepare() expected error, got nil")
				}
				return
			}

			// Verify the code structure
			verifyCodeStructure(t, got, tt.args.userCode, tt.want)
		})
	}
}

// verifyCodeStructure validates that the prepared code follows the expected format
func verifyCodeStructure(t *testing.T, code string, usercode string, want expected) {
	t.Logf("Prepared code:\n%s", code)
	// Check for variable declarations for all inputs
	for _, declaration := range want.varDeclarations {
		if !strings.Contains(code, declaration) {
			t.Errorf("Missing variable declaration: %s", declaration)
			return
		}
	}

	// Check for function call with proper parameters
	if !strings.Contains(code, want.functionCall) {
		t.Errorf("Missing function call: %s", want.functionCall)
		return
	}

	// Check for print statement
	if !strings.Contains(code, want.printResult) {
		t.Errorf("Missing print result statement: %s", want.printResult)
		return
	}

	// Check for user code
	if !strings.Contains(code, usercode) {
		t.Errorf("User code not included in prepared code")
		return
	}
}

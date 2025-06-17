package preparators

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/sriramr98/dsa_server/problems"
	"github.com/sriramr98/dsa_server/utils"
)

type PythonPreparator struct{}

func (p PythonPreparator) Prepare(userCode string, problem problems.Problem, testCase problems.TestCase) (string, error) {
	log.Println("User Code: ", userCode)
	userCode = strings.ReplaceAll(userCode, "from typing import *", "") // Remove import from userCode since it'll be injected later

	code := strings.ReplaceAll(CODE_TEMPLATE, "{imports}", "from typing import *\nimport json\n\n")

	sortedInputKeys := utils.MapKeysSorted(problem.Input)

	var variableDeclarations []string
	for _, inputKey := range sortedInputKeys {
		inputInfo := problem.Input[inputKey]
		var declaration string
		switch inputInfo.Type {
		case problems.ArrayType, problems.ObjectType:
			testCaseJSON, err := json.Marshal(testCase.Input[inputKey])
			if err != nil {
				return "", err
			}
			declaration = inputKey + " = " + string(testCaseJSON)
		case problems.StringType:
			declaration = fmt.Sprintf("%s = \"%s\"", inputKey, testCase.Input[inputKey])
		default:
			declaration = fmt.Sprintf("%s = %v", inputKey, testCase.Input[inputKey])
		}
		variableDeclarations = append(variableDeclarations, declaration)
	}

	varDeclarations := strings.Join(variableDeclarations, "\n")
	code = strings.ReplaceAll(code, "{var_declarations}", varDeclarations)

	code = strings.ReplaceAll(code, "{user_code}", userCode)

	funcArgs := utils.JoinStringSlice(sortedInputKeys, ", ")
	functionCall := fmt.Sprintf("%s(%s)", problem.FunctionName, funcArgs)
	if problem.ExecutionMode == problems.ReturnMode {
		functionCall = fmt.Sprintf("result = %s", functionCall)
	}
	code = strings.ReplaceAll(code, "{function_call}", functionCall)
	var printResult string
	if problem.ExecutionMode == problems.ReturnMode {
		if problem.Output.Type == problems.ArrayType || problem.Output.Type == problems.ObjectType {
			printResult = "print(json.dumps(result))"
		} else {
			printResult = "print(result)"
		}
	} else {
		outputVariable, _, found := utils.FindInMap(problem.Input, func(key string, value problems.InputType) bool {
			return value.Output
		})

		if !found {
			return "", fmt.Errorf("no output variable found in problem input")
		}

		printResult = fmt.Sprintf("print(%s)", outputVariable)
	}
	code = strings.ReplaceAll(code, "{print_result}", printResult)

	return code, nil
}

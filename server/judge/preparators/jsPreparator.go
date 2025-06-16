package preparators

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/sriramr98/dsa_server/problems"
	"github.com/sriramr98/dsa_server/utils"
)

type JsPreparator struct{}

func (p JsPreparator) Prepare(userCode string, problem problems.Problem, testCase problems.TestCase) (string, error) {
	codeToExecute := strings.ReplaceAll(CODE_TEMPLATE, "{imports}", "")

	var variableDeclarations []string
	for inputKey, inputInfo := range problem.Input {
		var declaration string
		switch inputInfo.Type {
		case problems.ArrayType, problems.ObjectType:
			testCaseJSON, err := json.Marshal(testCase.Input[inputKey])
			if err != nil {
				return "", err
			}
			declaration = fmt.Sprintf("let %s = %s;", inputKey, string(testCaseJSON))
		case problems.StringType:
			declaration = fmt.Sprintf("let %s = \"%s\";", inputKey, testCase.Input[inputKey])
		default:
			declaration = fmt.Sprintf("let %s = %v;", inputKey, testCase.Input[inputKey])
		}

		variableDeclarations = append(variableDeclarations, declaration)
	}

	varDeclarations := strings.Join(variableDeclarations, "\n")
	codeToExecute = strings.ReplaceAll(codeToExecute, "{var_declarations}", varDeclarations)

	codeToExecute = strings.ReplaceAll(codeToExecute, "{user_code}", userCode)

	funcArgs := utils.JoinStringSlice(utils.MapKeysSorted(problem.Input), ", ")
	functionCall := fmt.Sprintf("%s(%s)", problem.FunctionName, funcArgs)

	if problem.ExecutionMode == problems.ReturnMode {
		functionCall = fmt.Sprintf("let result = %s", functionCall)
	}

	codeToExecute = strings.ReplaceAll(codeToExecute, "{function_call}", functionCall)

	var printResult string
	if problem.ExecutionMode == problems.ReturnMode {
		if problem.Output.Type == problems.ArrayType || problem.Output.Type == problems.ObjectType {
			printResult = "console.log(JSON.stringify(result));"
		} else {
			printResult = "console.log(result);"
		}
	} else {
		var outputVariable string
		for inputKey, inputConf := range problem.Input {
			if inputConf.Output {
				outputVariable = inputKey
				break
			}
		}

		if outputVariable == "" {
			return "", errors.New("no output variable found for InPlace return type")
		}

		if problem.Input[outputVariable].Type == problems.ArrayType || problem.Input[outputVariable].Type == problems.ObjectType {
			printResult = fmt.Sprintf("console.log(JSON.stringify(%s));", outputVariable)
		} else {
			printResult = fmt.Sprintf("console.log(%s);", outputVariable)
		}
	}

	codeToExecute = strings.ReplaceAll(codeToExecute, "{print_result}", printResult)

	return codeToExecute, nil
}

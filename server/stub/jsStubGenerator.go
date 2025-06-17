package stub

import (
	"fmt"

	"github.com/sriramr98/dsa_server/problems"
	"github.com/sriramr98/dsa_server/utils"
)

type JSStubGenerator struct{}

func (g JSStubGenerator) Generate(problem problems.Problem) string {
	comment := getComment(problem)

	functionArgs := utils.JoinStringSlice(utils.MapKeysSorted(problem.Input), ", ")
	functionSignature := fmt.Sprintf("function %s(%s) {\n    // Your code here \n\n    return %s\n}", problem.FunctionName, functionArgs, getReturnVariable(problem))

	return fmt.Sprintf("%s\n%s", comment, functionSignature)
}

func getComment(problem problems.Problem) string {
	comment := "/**"
	inputKeysSorted := utils.MapKeysSorted(problem.Input)
	for _, variable := range inputKeysSorted {
		variableType := problem.Input[variable].VariableType
		comment += fmt.Sprintf("\n * @param {%s} %s", getType(variableType.Type, variableType.ChildType), variable)
		if variableType.Description != "" {
			comment += fmt.Sprintf(" - %s", variableType.Description)
		}
	}
	comment += "\n*/"
	return comment
}

func getType(variableType problems.DataType, childType *problems.VariableType) string {
	switch variableType {
	case problems.NumberType:
		return "number"
	case problems.StringType:
		return "string"
	case problems.BooleanType:
		return "boolean"
	case problems.ArrayType:
		if childType == nil {
			// If no child type is specified, we assume it can be any type
			return "any[]"
		}
		if childType.Type == problems.ArrayType {
			// If the child type is also an array, we need to recursively get the type
			if childType.ChildType != nil {
				return fmt.Sprintf("%s[]", getType(childType.Type, childType.ChildType))
			}
			return "any[][]" // Array of arrays with no element type
		} else if childType.Type == problems.NullType {
			return "any[]"
		} else if childType.Type == problems.NumberType || childType.Type == problems.StringType || childType.Type == problems.BooleanType {
			return fmt.Sprintf("%s[]", getType(childType.Type, nil))
		}
		return "any[]"
	case problems.ObjectType:
		return "Object" //TODO: Handle objects
	default:
		return "any"
	}
}

func getReturnVariable(problem problems.Problem) string {
	switch problem.Output.Type {
	case problems.ArrayType:
		return "[]"
	case problems.ObjectType:
		return "{}"
	case problems.StringType:
		return "\"\""
	case problems.NumberType:
		return "0"
	case problems.BooleanType:
		return "false"
	default:
		return "null"
	}
}

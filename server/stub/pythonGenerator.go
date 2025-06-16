package stub

import (
	"fmt"

	"github.com/sriramr98/dsa_server/problems"
	"github.com/sriramr98/dsa_server/utils"
)

type PythonStubGenerator struct {
}

func (g PythonStubGenerator) Generate(problem problems.Problem) string {
	functionArgs := getFunctionArgs(problem)
	stub := fmt.Sprintf("def %s(%s)", problem.FunctionName, functionArgs)

	if problem.ExecutionMode == problems.ReturnMode {
		stub += " -> " + getPyType(problem.Output.VariableType.Type, problem.Output.VariableType.ChildType)
	}

	stub += ":\n"
	stub += "    # Your code here\n    pass\n\n"

	return stub
}

func getFunctionArgs(problem problems.Problem) string {
	sortedFuncArgs := utils.MapKeysSorted(problem.Input)

	args := make([]string, len(sortedFuncArgs))
	for idx, arg := range sortedFuncArgs {
		variableType := problem.Input[arg].VariableType
		typeStr := getPyType(variableType.Type, variableType.ChildType)

		args[idx] = fmt.Sprintf("%s: %s", arg, typeStr)
	}

	return utils.JoinStringSlice(args, ", ")
}

func getPyType(variableType problems.DataType, childType *problems.VariableType) string {
	switch variableType {
	case problems.NumberType:
		return "int"
	case problems.StringType:
		return "str"
	case problems.BooleanType:
		return "bool"
	case problems.ArrayType:
		if childType == nil || childType.Type == problems.NullType {
			return "List[Any]"
		}
		if childType.Type == problems.ArrayType {
			if childType.ChildType != nil {
				return fmt.Sprintf("List[%s]", getPyType(childType.Type, childType.ChildType))
			}
			return "List[List[Any]]" // Array of arrays with no element type
		} else {
			return fmt.Sprintf("List[%s]", getPyType(childType.Type, nil))
		}
	case problems.ObjectType:
		return "dict" //TODO: Handle objects
	default:
		return "Any"
	}
}

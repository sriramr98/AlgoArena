import Preparator from ".";
import { Problem } from "../../problems";
import { TEMPLATE } from "../codeTemplate";
import { TestCase } from "../../problems";

const jsPreparator: Preparator = {
  // This function creates the final code that'll be executed for Javascript
  prepare(userCode: string, problem: Problem, testCase: TestCase): string {
    let codeToExecute = TEMPLATE.replace("{imports}", "");

    let variableDeclarations: string[] = [];
    Object.keys(problem.input).forEach((key) => {
      const inputType = problem.input[key].type;
      if (inputType === "array" || inputType === "object") {
        variableDeclarations.push(`let ${key} = ${JSON.stringify(testCase.input[key])};`);
      } else if (inputType === "string") {
        variableDeclarations.push(`let ${key} = "${testCase.input[key]}";`);
      } else {
        variableDeclarations.push(`let ${key} = ${testCase.input[key]};`);
      }
    });
    codeToExecute = codeToExecute.replace(
      "{var_declarations}",
      variableDeclarations.join("\n")
    );

    codeToExecute = codeToExecute.replace(
      "{user_code}",
      userCode
    );

    let functionCall = `${problem.functionName}(${Object.keys(problem.input).join(", ")});`;
    if (problem.executionMode === "return") {
      functionCall = `let result = ${functionCall}`;
    }
    codeToExecute = codeToExecute.replace(
      "{function_call}",
      functionCall
    );

    let printResult = "";
    if (problem.executionMode === "return") {
      printResult = `console.log(result);`;
    } else {
      const outputVariable = Object.keys(problem.input).filter(key => problem.input[key].output)[0];
      printResult = `console.log(${outputVariable});`;
    }

    codeToExecute = codeToExecute.replace(
      "{print_result}",
      printResult
    );

    return codeToExecute;
  }
};

export = jsPreparator;
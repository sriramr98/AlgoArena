import { EvaluationResult, Evaluator } from ".";
import { TestCase, ComparisonMode } from "../../problems";
import {
  areArraysEqualOrdered,
  areArraysEqualUnordered,
} from "../comparators/array";
import { ExecutionResult } from "../executor";

const SUPPORTED_COMPARISON_TYPES = [ComparisonMode.ORDERED, ComparisonMode.UNORDERED];

class ArrayEvaluator implements Evaluator {
  evaluate(executionResult: ExecutionResult, testCase: TestCase, comparsionMode: ComparisonMode): EvaluationResult {
    const { stdout = "", stderr = "", code = 1 } = executionResult;

    if (!SUPPORTED_COMPARISON_TYPES.includes(comparsionMode)) {
      return {
        passed: false,
        error: `Unsupported comparison type: ${comparsionMode}. Supported types are ${SUPPORTED_COMPARISON_TYPES.join(", ")}.`,
        actualOutput: stdout,
        expectedOutput: testCase.expected,
      };
    }

    try {
      if (code !== 0) {
        return {
          passed: false,
          //TODO: get error from executionResult
          error: `Execution failed with code ${code}`,
          actualOutput: stdout,
          expectedOutput: testCase.expected,
        };
      }

      if (stderr) {
        return {
          passed: false,
          error: stderr,
          actualOutput: stdout,
          expectedOutput: testCase.expected,
        };
      }

      const expectedOutput = testCase.expected;
      const actualOutput = JSON.parse(stdout);

      if (!Array.isArray(actualOutput)) {
        return {
          passed: false,
          error: `Expected an array but got ${typeof actualOutput}`,
          actualOutput,
          expectedOutput,
        };
      }

      const compareArrays =
        comparsionMode === ComparisonMode.ORDERED
          ? areArraysEqualOrdered
          : areArraysEqualUnordered;
      if (compareArrays(expectedOutput, actualOutput)) {
        return {
          passed: true,
          actualOutput,
          expectedOutput,
        };
      }

      return {
        passed: false,
        actualOutput,
        expectedOutput,
      };
    } catch (error: any) {
      return {
        passed: false,
        error: `Error in evaluation: ${error.message || error}`,
      };
    }

  }
}

export default ArrayEvaluator;

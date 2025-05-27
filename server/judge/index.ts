import { Problem } from "../problems";
import codeExecutor from "./executor";
import ArrayEvaluator from "./evaluators/arrayEvaluator";
import { Evaluator } from "./evaluators";
import Preparator from "./preparators";
import jsPreparator from "./preparators/jsPreparator";

interface TestResult {
  input: any;
  expectedOutput: any;
  actualOutput: any;
  passed: boolean;
  error: string;
}

interface JudgeResult {
  language: string;
  passed: number;
  total: number;
  testResults: TestResult[];
  success?: boolean;
  successRate?: number;
  message?: string;
}

const getPreparator = (language: string): Preparator | undefined => {
  const preparators: Record<string, Preparator> = {
    javascript: jsPreparator,
  };
  return preparators[language];
};

const getEvaluator = (problem: Problem): Evaluator => {
  const evaluators: Record<string, Evaluator> = {
    array: new ArrayEvaluator(),
    // "string": require("./evaluators/stringEvaluator"),
    // "number": require("./evaluators/numberEvaluator"),
    // "boolean": require("./evaluators/booleanEvaluator"),
    // "object": require("./evaluators/objectEvaluator"),
  };
  return evaluators[problem.output.type];
};

const judge = async (
  userCode: string,
  problem: Problem,
  language: string,
  noOfTestCases = 2,
): Promise<JudgeResult> => {
  const preparator = getPreparator(language);
  if (!preparator) {
    return {
      language,
      passed: 0,
      total: 0,
      testResults: [],
      success: false,
      message: `Language ${language} is not supported`,
    };
  }

  const testCases = problem.testCases.slice(0, noOfTestCases);
  const results: JudgeResult = {
    language,
    passed: 0,
    total: noOfTestCases,
    testResults: [],
  };

  for (const testCase of testCases) {
    try {
      const codeToExecute = preparator.prepare(userCode, problem, testCase);
      console.log(codeToExecute);
      const executionResult = await codeExecutor(codeToExecute, language);

      const evaluator = getEvaluator(problem);
      if (!evaluator) {
        throw new Error(
          `Evaluator not found for problem ${problem.id} output type ${problem.output.type}`,
        );
      }
      const evalResult = evaluator.evaluate(
        executionResult,
        testCase,
        problem.comparisonMode,
      );

      if (evalResult.passed) {
        results.passed++;
      }
      results.testResults.push({
        input: testCase.input,
        expectedOutput: testCase.expected,
        actualOutput: evalResult.actualOutput,
        passed: evalResult.passed,
        error: evalResult.error || "",
      });
    } catch (error: any) {
      results.testResults.push({
        input: testCase.input,
        expectedOutput: testCase.expected,
        actualOutput: null,
        passed: false,
        error: error.message || "An error occurred during execution",
      });
    }
  }

  results.success = results.passed === results.total;
  results.successRate = (results.passed / results.total) * 100;

  return results;
};

export default judge;

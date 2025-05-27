import { ComparisonMode, TestCase } from "../../problems";
import { ExecutionResult } from "../executor";

export interface EvaluationResult {
    passed: boolean;
    error?: string;
    actualOutput?: any;
    expectedOutput?: any;
}

export interface Evaluator {
  evaluate(
    executionResult: ExecutionResult,
    testCase: TestCase,
    comparsionMode: ComparisonMode,
  ): EvaluationResult;
}
import { Problem, TestCase } from "../../problems";

export default interface Preparator {
  prepare(userCode: string, problem: Problem, testCase: TestCase): string;
}
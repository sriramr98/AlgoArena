// Types for problem structures

export enum DataType {
  STRING = "string",
  NUMBER = "number",
  BOOLEAN = "boolean",
  ARRAY = "array",
  OBJECT = "object",
}

export enum ExecutionMode {
  RETURN = "return",
  IN_PLACE = "in-place",
}

export enum ComparisonMode {
  EXACT = "exact",
  ORDERED = "ordered",
  UNORDERED = "unordered",
}

export interface TestCase {
  input: any;
  expected: any;
}

export interface ProblemExample {
  input: string;
  output: string;
  explanation?: string;
}

export interface InputType {
  type: DataType;
  output?: boolean;
}

export interface Problem {
  id: string;
  title: string;
  difficulty: string;
  description: string;
  examples: ProblemExample[];
  constraints: string[];
  input: Record<string, InputType>;
  output: {
    type: DataType;
  };
  functionName?: string;
  executionMode: ExecutionMode;
  comparisonMode: ComparisonMode;
  testCases: TestCase[];
}

// Sample hardcoded problems
export const problems: Problem[] = [
  {
    id: "two-sum",
    title: "Two Sum",
    difficulty: "Easy",
    description:
      "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.",
    examples: [
      {
        input: "nums = [2,7,11,15], target = 9",
        output: "[0,1]",
        explanation: "Because nums[0] + nums[1] == 9, we return [0, 1].",
      },
      {
        input: "nums = [3,2,4], target = 6",
        output: "[1,2]",
      },
    ],
    constraints: [
      "2 <= nums.length <= 10^4",
      "-10^9 <= nums[i] <= 10^9",
      "-10^9 <= target <= 10^9",
      "Only one valid answer exists.",
    ],
    input: {
      nums: {
        type: DataType.ARRAY,
      },
      target: {
        type: DataType.NUMBER,
      },
    },
    output: {
      type: DataType.ARRAY,
    },
    functionName: "twoSum",
    executionMode: ExecutionMode.RETURN,
    comparisonMode: ComparisonMode.ORDERED,
    testCases: [
      {
        input: {
          nums: [2, 7, 11, 15],
          target: 9,
        },
        expected: [0, 1],
      },
      {
        input: {
          nums: [3, 2, 4],
          target: 6,
        },
        expected: [1, 2],
      },
      {
        input: {
          nums: [3, 3],
          target: 6,
        },
        expected: [0, 1],
      },
      {
        input: {
          nums: [1, 2, 3, 4, 5],
          target: 9,
        },
        expected: [3, 4],
      },
      {
        input: {
          nums: [0, 4, 3, 0],
          target: 0,
        },
        expected: [0, 3],
      },
    ],
  },
  {
    id: "reverse-string",
    title: "Reverse String",
    difficulty: "Easy",
    description:
      "Write a function that reverses a string. The input string is given as an array of characters s.\n\nYou must do this by modifying the input array in-place with O(1) extra memory.",
    examples: [
      {
        input: 'Input: s = "hello"',
        output: "olleh",
      },
      {
        input: 'Input: s = "Hannah"',
        output: "hannaH",
      },
    ],
    constraints: [
      "1 <= s.length <= 10^5",
      "s[i] is a printable ascii character.",
    ],
    input: {
      s: {
        type: DataType.ARRAY,
        output: true,
      },
    },
    output: {
      type: DataType.ARRAY,
    },
    testCases: [
      {
        input: {
          s: ["h", "e", "l", "l", "o"],
        },
        expected: ["o", "l", "l", "e", "h"],
      },
      {
        input: {
          s: ["H", "a", "n", "n", "a", "h"],
        },
        expected: ["h", "a", "n", "n", "a", "H"],
      },
      {
        input: {
          s: ["a"],
        },
        expected: ["a"],
      },
      {
        input: {
          s: ["a", "b", "c", "d"],
        },
        expected: ["d", "c", "b", "a"],
      },
    ],
    executionMode: ExecutionMode.IN_PLACE,
    comparisonMode: ComparisonMode.ORDERED,
  },
  {
    id: "longest-substring-without-repeating-characters",
    title: "Longest Substring Without Repeating Characters",
    difficulty: "Medium",
    description:
      "Given a string s, find the length of the longest substring without repeating characters.",
    examples: [
      {
        input: 'Input: s = "abcabcbb"',
        output: "3",
        explanation: "The answer is 'abc', with the length of 3.",
      },
      {
        input: 'Input: s = "bbbbb"',
        output: "1",
        explanation: "The answer is 'b', with the length of 1.",
      },
      {
        input: 'Input: s = "pwwkew"',
        output: "3",
        explanation:
          "The answer is 'wke', with the length of 3. Notice that the answer must be a substring, 'pwke' is a subsequence and not a substring.",
      },
    ],
    constraints: [
      "0 <= s.length <= 5 * 10^4",
      "s consists of English letters, digits, symbols and spaces.",
    ],
    input: {
      s: {
        type: DataType.STRING,
      },
    },
    output: {
      type: DataType.NUMBER,
    },
    testCases: [
      {
        input: { s: "abcabcbb" },
        expected: 3,
      },
      {
        input: { s: "bbbbb" },
        expected: 1,
      },
      {
        input: { s: "pwwkew" },
        expected: 3,
      },
      {
        input: { s: "" },
        expected: 0,
      },
      {
        input: { s: "aab" },
        expected: 2,
      },
      {
        input: { s: "dvdf" },
        expected: 3,
      },
    ],
    executionMode: ExecutionMode.RETURN,
    comparisonMode: ComparisonMode.EXACT,
  },
];

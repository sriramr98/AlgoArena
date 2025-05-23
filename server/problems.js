// Sample hardcoded problems
const problems = [
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
    executionMode: "return",
    comparisonMode: "unordered-array",
    testCases: [
      {
        input: {
          nums: {
            value: [2, 7, 11, 15],
            type: "array",
          },
          target: {
            value: 9,
            type: "number",
          },
        },
        expected: {
          value: [0, 1],
          type: "array",
        },
      },
      {
        input: {
          nums: {
            value: [3, 2, 4],
            type: "array",
          },
          target: {
            value: 6,
            type: "number",
          },
        },
        expected: {
          value: [1, 2],
          type: "array",
        },
      },
      {
        input: {
          nums: {
            value: [3, 3],
            type: "array",
          },
          target: {
            value: 6,
            type: "number",
          },
        },
        expected: {
          value: [0, 1],
          type: "array",
        },
      },
      {
        input: {
          nums: {
            value: [1, 2, 3, 4, 5],
            type: "array",
          },
          target: {
            value: 9,
            type: "number",
          },
        },
        expected: {
          value: [3, 4],
          type: "array",
        },
      },
      {
        input: {
          nums: {
            value: [0, 4, 3, 0],
            type: "array",
          },
          target: {
            value: 0,
            type: "number",
          },
        },
        expected: {
          value: [0, 3],
          type: "array",
        },
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
    testCases: [
      {
        input: {
          s: {
            value: ["h", "e", "l", "l", "o"],
            type: "array",
          },
        },
        expected: {
          value: ["o", "l", "l", "e", "h"],
          type: "array",
        },
      },
      {
        input: {
          s: {
            value: ["H", "a", "n", "n", "a", "h"],
            type: "array",
          },
        },
        expected: {
          value: ["h", "a", "n", "n", "a", "H"],
          type: "array",
        },
      },
      {
        input: {
          s: {
            value: ["a"],
            type: "array",
          },
        },
        expected: {
          value: ["a"],
          type: "array",
        },
      },
      {
        input: {
          s: {
            value: ["a", "b", "c", "d"],
            type: "array",
          },
        },
        expected: {
          value: ["d", "c", "b", "a"],
          type: "array",
        },
      },
    ],
    executionMode: "in-place",
    comparisonMode: "array",
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
    executionMode: "return",
    comparisonMode: "exact",
  },
];

module.exports = problems;

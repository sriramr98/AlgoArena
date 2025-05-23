// Code execution manager
const fs = require("fs");
const executeJavaScript = require("./javascriptExecutor");
const executePython = require("./pythonExecutor");
const executeCpp = require("./cppExecutor");

/**
 * Executes code in the specified language against test cases
 * @param {string} code - The user's submitted code
 * @param {string} language - The programming language (javascript, python, cpp)
 * @param {Object} problem - The problem details
 * @returns {Promise<Object>} - Results of the test case executions
 */
async function executeCode(code, language, problem, preview = true) {
  try {
    // Choose the right executor based on the language
    const executor = getExecutor(language);
    if (!executor) {
      throw new Error(`Unsupported language: ${language}`);
    }

    const testCasesToRun = preview
      ? problem.testCases.slice(0, 2)
      : problem.testCases;
    const results = {
      language,
      passed: 0,
      total: testCasesToRun.length,
      testResults: [],
    };

    // Run each test case
    for (const testCase of testCasesToRun) {
      try {
        const codeToExecute = prepareCodeToExecute(problem.id, code, testCase);

        const result = await executor(codeToExecute, testCase);
        console.log("Test case result:", result);

        // Add to results
        results.testResults.push({
          input: JSON.stringify(testCase.input),
          expectedOutput: result.expected,
          actualOutput: result.output,
          passed: result.passed,
          error: null,
        });

        if (result.passed) {
          results.passed++;
        }
      } catch (error) {
        console.error("Test case execution error:", error);
        // Handle test case execution errors
        results.testResults.push({
          input: JSON.stringify(testCase.input),
          expectedOutput: JSON.stringify(testCase.expected),
          actualOutput: null,
          passed: false,
          error: "Unable to run code..",
        });
      }
    }

    // Calculate overall result
    results.success = results.passed === results.total;
    results.successRate = (results.passed / results.total) * 100;

    return results;
  } catch (error) {
    console.error("Code execution error:", error);
    return {
      language,
      error: error.message || "Failed to execute code",
      success: false,
      passed: 0,
      total: problem.testCases.length,
      testResults: [],
    };
  }
}

/**
 * Returns the appropriate executor function for a given language
 * @param {string} language
 * @returns {Function} Executor function
 */
function getExecutor(language) {
  switch (language.toLowerCase()) {
    case "javascript":
      return executeJavaScript;
    case "python":
      return executePython;
    case "cpp":
      return executeCpp;
    default:
      return null;
  }
}

const prepareCodeToExecute = (problemId, userCode, testCase) => {
  let codeToExecute = fs.readFileSync(
    `${process.cwd()}/code_templates/${problemId}/exec/javascript`,
    "utf-8",
  );
  codeToExecute = codeToExecute.replace("{user_code}", userCode);

  Object.entries(testCase.input).forEach(([key, value]) => {
    if (value.type === "array" || value.type === "object") {
      value = JSON.stringify(value.value);
    } else if (value.type != "string") {
      value = `${value.value}`;
    } else {
      value = `"${value.value}"`;
    }
    codeToExecute = codeToExecute.replace(`{${key}}`, value);
  });

  console.log(codeToExecute);

  return codeToExecute;
};

module.exports = { executeCode };

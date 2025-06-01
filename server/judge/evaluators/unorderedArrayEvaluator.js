const areArraysEqual = (arr1 = [], arr2 = []) => {
  const occurences = new Map();
  arr1.forEach((item) => {
    occurences.set(item, (occurences.get(item) || 0) + 1);
  });

  arr2.forEach((item) => {
    if (!occurences.has(item)) {
      return false;
    }
    occurences.set(item, occurences.get(item) - 1);
  });

  for (const [key, value] of occurences) {
    if (value !== 0) {
      return false;
    }
  }

  return true;
};

module.exports = {
  evaluate(executionResult, testCase) {
    try {
      const { stdout, stderr, code } = executionResult;
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

      if (areArraysEqual(expectedOutput, actualOutput)) {
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
    } catch (error) {
      return {
        passed: false,
        error: `Error in evaluation: ${error.error}`,
      };
    }
  },
};

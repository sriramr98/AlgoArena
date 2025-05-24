const codeExecutor = require('./executor');

const getPreparator = (language) => {
    const preparators = {
        javascript: require("./preparators/jsPreparator"),
    };
    return preparators[language];
}

const getEvaluator = (problem) => {
    const evaluators = {
        "unordered-array": require("./evaluators/unorderedArrayEvaluator"),
        // "ordered-array": require("./evaluators/orderedArrayEvaluator"),
        // "string": require("./evaluators/stringEvaluator"),
        // "number": require("./evaluators/numberEvaluator"),
        // "boolean": require("./evaluators/booleanEvaluator"),
        // "object": require("./evaluators/objectEvaluator"),
    }
    return evaluators[problem.comparisonMode];
}

const judge = async (userCode, problem, language, noOfTestCases = 2) => {
    const preparator = getPreparator(language);
    if (!preparator) {
        return {
            success: false,
            message: `Language ${language} is not supported`,
        };
    }

    const testCases = problem.testCases.slice(0, noOfTestCases);
    const results = {
        language,
        passed: 0,
        total: noOfTestCases,
        testResults: []
    }

    for (const testCase of testCases) {
        try {
            const codeToExecute = preparator.prepare(userCode, problem, testCase);
            console.log(codeToExecute);
            const executionResult = await codeExecutor(codeToExecute, language);

            const evaluator = getEvaluator(problem)
            const evalResult = evaluator.evaluate(executionResult, testCase);

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

        } catch (error) {
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
}

module.exports = judge
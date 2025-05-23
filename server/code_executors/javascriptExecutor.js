const { areArraysEqual } = require('../utils/arrayUtils');
const callPiston = require('./piston');

/**
 * Execute JavaScript code with provided test cases
 * @param {string} code - User submitted code
 * @param {Object} problem - Problem metadata
 * @param {Object} testCase - Test case to run
 * @returns {Promise<Object>} - Result of execution
 */
const executeJavaScript= async (code, testCase) => {
    const execResult = await callPiston("javascript", code) 
    const data = execResult.data;
    console.log("Data from Piston: ", data);
    if (data.run.code === 0) {
        // Process exited successfully, no errors thrown

        const output = data.run.stdout;
        const expected = testCase.expected;

        if (expected.type === "array") {
            let userOutput;
            try {
                userOutput = JSON.parse(output); 
            } catch (e) {
                return {
                    passed: false,
                    expected: JSON.stringify(expected.value),
                    output: output,
                }
            }

            const baseResult = {
                expected: JSON.stringify(expected.value),
                output: JSON.stringify(userOutput),
            }

            if (userOutput.length !== expected.value.length) {
                return {
                    passed: false,
                    ...baseResult
                }
            }

            if (areArraysEqual(userOutput, expected.value)) {
                return {
                    passed: true,
                    ...baseResult
                }
            } else {
                return {
                    passed: false,
                    ...baseResult
                }
            }
        }


        return {
            passed: output === expected.value,
            expected: expected.value,
            output: output,
        }

    }
    return {
        passed: false
    }
}

module.exports = executeJavaScript;
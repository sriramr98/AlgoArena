const fs = require('fs');
const PROBLEMS = require('./../server/problems')
const SUPPORTED_LANGUAGES = require('./../server/languages');

const getExpectedFunctionDeclaration = (functionName, inputSchema, language) => {
    if (language === 'python') {
        const args = Object.keys(inputSchema).join(', ');
        return `def ${functionName}(${args}):`;
    } else if (language === 'javascript') {
        const args = Object.keys(inputSchema).join(', ');
        return `function ${functionName}(${args}) {`;
    }
    // Add more languages as needed
    throw new Error(`Unsupported language: ${language}`);
}

for (const problem of PROBLEMS) {
    const issues = [];

    if(!problem.id || typeof problem.id !== 'string') {
        issues.push(`Problem ID is missing or not a string: ${problem.id}`);
    }

    if (!problem.title || typeof problem.title !== 'string') {
        issues.push(`Problem title is missing or not a string: ${problem.title}`);
    }

    if (!problem.difficulty || typeof problem.difficulty !== 'string') {
        issues.push(`Problem difficulty is missing or not a string: ${(problem.difficulty)}`);
    }
    if (["Easy", "Medium", "Hard"].indexOf(problem.difficulty) === -1) {
        issues.push(`Problem difficulty is not one of Easy, Medium, Hard: ${(problem.difficulty)}`);
    }
    
    if (!problem.description || typeof problem.description !== 'string') {
        issues.push(`Problem description is missing or not a string: ${(problem.description)}`);
    }
    
    if (!Array.isArray(problem.examples)) {
        issues.push(`Problem examples are missing or not an array: ${(problem.examples)}`);
    }
    
    for (const example of problem.examples) {
        if (!example.input || !example.output || !example.explanation) {
            issues.push(`Example is missing input, output or explanation: ${JSON.stringify(example)}`);
        }
    }
    
    if (!problem.constraints || !Array.isArray(problem.constraints)) {
        issues.push(`Problem constraints are missing or not an array: ${JSON.stringify(problem.constraints)}`);
    }
    
    if (!problem.input || typeof problem.input !== 'object') {
        issues.push(`Problem input schema is missing or not an objectL ${JSON.stringify(problem.input)}`);
    }
    
    if (!problem.output || typeof problem.output !== 'object') {
        issues.push(`Problem output schema is missing or not an object: ${JSON.stringify(problem.output)}`);
    }

    if (!problem.functionName || typeof problem.functionName !== 'string') {
        issues.push(`Problem function name is missing or not a string: ${problem.functionName}`);
    }
    
    if (!Array.isArray(problem.testCases)) {
        issues.push(`Problem test cases are missing or not an array: ${JSON.stringify(problem.testCases)}`);
    }
    
    for (const testCase of problem.testCases) {
        if (!testCase.input || !testCase.expected) {
            issues.push(`Test case is missing input or expected output: ${JSON.stringify(testCase)}`);
        }
    }
    
    if (problem.executionMode && typeof problem.executionMode !== 'string') {
        issues.push(`Problem execution mode is not a string: ${(problem.executionMode)}`);
    }
    
    if (problem.comparisonMode && typeof problem.comparisonMode !== 'string') {
        issues.push(`Problem comparison mode is not a string: ${(problem.comparisonMode)}`);
    }

    for (const language of SUPPORTED_LANGUAGES) {
        const stub = fs.readFileSync(`./../server/code_templates/${problem.id}/stub/${language}`, 'utf8');
        const expectedFuncDeclaration = getExpectedFunctionDeclaration(problem.functionName, problem.input, language);
        if (!stub.includes(expectedFuncDeclaration)) {
            issues.push(`Stub for ${language} does not contain expected function declaration: ${expectedFuncDeclaration}`);
        }
    }

    if (issues.length > 0) {
        console.log(`Issues found in problem ${problem.id} (${problem.title}):`);
        for (const issue of issues) {
            console.log(`- ${issue}`);
        }
    } else {
        console.log(`Problem ${problem.id} (${problem.title}) is valid.`);
    }
    console.log('-----------------------------------');
}
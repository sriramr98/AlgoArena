package judge

import (
	"errors"
	"log"

	"github.com/sriramr98/dsa_server/judge/evaluators"
	"github.com/sriramr98/dsa_server/judge/executors"
	"github.com/sriramr98/dsa_server/judge/preparators"
	"github.com/sriramr98/dsa_server/problems"
	"github.com/sriramr98/dsa_server/utils"
)

type TestResult struct {
	Input          map[string]any
	ExpectedOutput any
	ActualOutput   any
	Passed         bool
	Error          error
}

type JudgeResult struct {
	Language    string
	Passed      bool
	TotalCases  int
	TotalPassed int
	TotalFailed int
	TestResults []TestResult
	SuccessRate int
}

func JudgeProblem(problem problems.Problem, userCode string, language string, noOfTestCases int) (JudgeResult, error) {
	preparator, err := preparators.GetPreparator(language)
	if err != nil {
		return JudgeResult{}, err
	}

	evaluator, err := evaluators.GetEvaluator(problem)
	if err != nil {
		return JudgeResult{}, err
	}

	codeExecutor := executors.GetExecutor()

	testCases := problem.TestCases
	if len(testCases) > noOfTestCases {
		testCases = testCases[0:noOfTestCases]
	}

	judgeResult := JudgeResult{
		Language:    language,
		Passed:      false,
		TotalCases:  len(testCases),
		TestResults: make([]TestResult, 0),
	}

	for _, testCase := range testCases {
		codeToExecute, err := preparator.Prepare(userCode, problem, testCase)
		if err != nil {
			utils.LogError(err)
			return judgeResult, err
		}

		executionResult, err := codeExecutor.Execute(codeToExecute, executors.ExecutorConfig{Language: language})
		log.Printf("Execution result: %+v", executionResult)
		if err != nil {
			utils.LogError(err)
			return judgeResult, err
		}

		if executionErr := validateExecution(executionResult); executionErr != nil {
			utils.LogError(executionErr)
			return judgeResult, err
		}

		evaluationResult, err := evaluator.Evaluate(executionResult, testCase, problem.ComparisonMode)
		if err != nil {
			utils.LogError(err)
			return judgeResult, err
		}

		if evaluationResult.Passed {
			judgeResult.TotalPassed += 1
		} else {
			judgeResult.TotalFailed += 1
		}

		judgeResult.TestResults = append(judgeResult.TestResults, TestResult{
			Input:          testCase.Input,
			ExpectedOutput: testCase.Expected,
			ActualOutput:   evaluationResult.ActualOutput,
			Passed:         evaluationResult.Passed,
			Error:          evaluationResult.Error,
		})

	}

	judgeResult.Passed = judgeResult.TotalPassed == judgeResult.TotalCases
	judgeResult.SuccessRate = (judgeResult.TotalPassed / judgeResult.TotalCases) * 100

	return judgeResult, nil
}

func validateExecution(result executors.ExecutorOutput) error {
	if result.Run.Code != 0 {
		return errors.New("error executing code")
	}

	if result.Run.Stderr != "" {
		return errors.New(result.Run.Stderr)
	}

	return nil
}

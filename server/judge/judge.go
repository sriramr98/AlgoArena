package judge

import (
	"errors"
	"log"
	"strings"

	"github.com/sriramr98/dsa_server/judge/evaluators"
	"github.com/sriramr98/dsa_server/judge/executors"
	"github.com/sriramr98/dsa_server/judge/preparators"
	"github.com/sriramr98/dsa_server/problems"
	"github.com/sriramr98/dsa_server/utils"
)

type TestResult struct {
	Input          map[string]any `json:"input"`
	ExpectedOutput any            `json:"expectedOutput"`
	ActualOutput   any            `json:"actualOutput"`
	Passed         bool           `json:"passed"`
	Error          error          `json:"error,omitempty"`
}

type JudgeResult struct {
	Language    string       `json:"language"`
	Passed      bool         `json:"passed"`
	TotalCases  int          `json:"totalCases"`
	TotalPassed int          `json:"totalPassed"`
	TotalFailed int          `json:"totalFailed"`
	TestResults []TestResult `json:"testResults"`
	SuccessRate int          `json:"successRate"`
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

		output := strings.TrimSpace(executionResult.Run.Stdout)

		evaluationResult, err := evaluator(output, testCase, problem.ComparisonMode)
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

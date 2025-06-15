package controllers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/sriramr98/dsa_server/judge"
	"github.com/sriramr98/dsa_server/problems"
	"github.com/sriramr98/dsa_server/utils"
)

type SubmitProblemRequest struct {
	Code     string `json:"code"`
	Language string `json:"language"`
	Run      bool   `json:"run"`
}

type ProblemsResponse struct {
	Id         string              `json:"id"`
	Title      string              `json:"title"`
	Difficulty problems.Difficulty `json:"difficulty"`
}

type ProblemController struct{}

func (pc ProblemController) GetProblems(ctx *gin.Context) {
	problems, err := problems.Problems()
	if err != nil {
		utils.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.FailureResponse("Something went wrong"))
	}

	problemsResp := []ProblemsResponse{}
	for _, problem := range problems {
		problemsResp = append(problemsResp, ProblemsResponse{
			Title:      problem.Title,
			Id:         problem.ID,
			Difficulty: problem.Difficulty,
		})
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse(problemsResp))
}

func (pc ProblemController) GetProblemDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	problem, err := problems.ProblemForID(id)
	if err != nil {
		if errors.Is(err, problems.ErrProblemNotFound) {
			ctx.JSON(http.StatusNotFound, utils.FailureResponse("Problem not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.FailureResponse("Unable to fetch problem"))
	}

	// Send only first two test cases
	problem.TestCases = problem.TestCases[0:2]
	ctx.JSON(http.StatusOK, utils.SuccessResponse(problem))
}

func (pc ProblemController) GetProblemStub(ctx *gin.Context) {
	id := ctx.Param("id")
	language := ctx.Param("language")
	fmt.Printf("Fetching stub for problem %s language %s\n", id, language)

	if !slices.Contains(problems.SUPPORTED_LANGUAGES, language) {
		ctx.JSON(http.StatusBadRequest, utils.FailureResponse("Unsupported language "+language))
		return
	}

	//TODO: Create stub dynamically from the problem definition with sorted args
	root, err := os.Getwd()
	if err != nil {
		utils.LogError(err)
		ctx.JSON(http.StatusInternalServerError, utils.FailureResponse("Something went wrong"))
	}

	stubFolderPath := path.Join(root, "code_templates", id, "stub", language)
	fmt.Println("Stub Path: ", stubFolderPath)
	data, err := os.ReadFile(stubFolderPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FailureResponse("Unable to find stub for problem"))
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse(string(data)))
}

func (pc ProblemController) SubmitProblem(ctx *gin.Context) {
	id := ctx.Param("id")
	base64Encoded := ctx.Query("base64Encoded")

	var submitRequest SubmitProblemRequest

	if err := ctx.BindJSON(&submitRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FailureResponse(err.Error()))
		return
	}

	problem, err := problems.ProblemForID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FailureResponse(err.Error()))
		return
	}

	noOfTestCases := len(problem.TestCases)
	if submitRequest.Run {
		noOfTestCases = 2
	}

	var userCode string
	if base64Encoded == "true" {
		userCodeBytes, err := base64.StdEncoding.DecodeString(submitRequest.Code)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.FailureResponse("Invalid base64 code format"))
			return
		}
		userCode = string(userCodeBytes)
	} else {
		userCode = submitRequest.Code
	}

	res, err := judge.JudgeProblem(problem, string(userCode), submitRequest.Language, noOfTestCases)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FailureResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse(res))
}

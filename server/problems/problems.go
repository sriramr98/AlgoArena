package problems

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

// Language represents a list of all supported languages

var SUPPORTED_LANGUAGES = []string{"javascript", "python"}

// Difficulty represents the difficulty level of a problem
type Difficulty string

const (
	Easy   Difficulty = "EASY"
	Medium Difficulty = "MEDIUM"
	Hard   Difficulty = "HARD"
)

// DataType represents the data type for inputs and outputs
type DataType string

const (
	StringType  DataType = "string"
	NumberType  DataType = "number"
	BooleanType DataType = "boolean"
	ArrayType   DataType = "array"
	ObjectType  DataType = "object"
	FloatType   DataType = "float"
	NullType    DataType = "null"
)

// ExecutionMode defines how the solution should be executed
type ExecutionMode string

const (
	ReturnMode  ExecutionMode = "return"
	InPlaceMode ExecutionMode = "in-place"
)

// ComparisonMode defines how outputs should be compared
type ComparisonMode string

const (
	ExactMode     ComparisonMode = "exact"
	OrderedMode   ComparisonMode = "ordered"
	UnorderedMode ComparisonMode = "unordered"
)

// TestCase represents a single test case with input and expected output
type TestCase struct {
	Input    map[string]any `json:"input"`
	Expected any            `json:"expected"`
}

// ProblemExample provides an example of the problem with input, output and explanation
type ProblemExample struct {
	Input       string `json:"input"`
	Output      string `json:"output"`
	Explanation string `json:"explanation,omitempty"`
}

type VariableType struct {
	Type    DataType `json:"type"`
	SubType DataType `json:"subType"` // For example, if the input is to be an array of strings, Type will be ArrayType and SubType will be StringType
}

// InputType defines the type of an input parameter
type InputType struct {
	VariableType
	Output bool `json:"output,omitempty"`
}

// OutputType defines the type of the output of the function
type OutputType struct {
	VariableType
}

// Problem represents a complete problem definition
type Problem struct {
	ID             string               `json:"id"`
	Title          string               `json:"title"`
	Difficulty     Difficulty           `json:"difficulty"`
	Description    string               `json:"description"`
	Examples       []ProblemExample     `json:"examples"`
	Constraints    []string             `json:"constraints"`
	Input          map[string]InputType `json:"input"`
	Output         OutputType           `json:"output"`
	FunctionName   string               `json:"functionName,omitempty"`
	ExecutionMode  ExecutionMode        `json:"executionMode"`
	ComparisonMode ComparisonMode       `json:"comparisonMode"`
	TestCases      []TestCase           `json:"testCases"`
}

var ErrProblemNotFound = errors.New("problem not found")

var problems []Problem

func Problems() ([]Problem, error) {
	if len(problems) > 0 {
		return problems, nil
	}

	cwd, err := os.Getwd()
	if err != nil {
		return problems, err
	}

	problemsFolder := path.Join(cwd, "problems")

	problemFiles, err := os.ReadDir(problemsFolder)
	if err != nil {
		return []Problem{}, err
	}

	for _, file := range problemFiles {
		// Ignore folders
		if file.IsDir() {
			continue
		}

		// Ignore not json files
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		filePath := path.Join(problemsFolder, file.Name())
		problemData, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			// Ignore file with error
			fmt.Println("Unable to read file: ", filePath)
			continue
		}

		var problem Problem
		if err = json.Unmarshal(problemData, &problem); err != nil {
			fmt.Println(err)
			fmt.Println("Invalid problem at :", filePath)
			continue
		}

		problems = append(problems, problem)
	}

	return problems, nil
}

func ProblemForID(id string) (Problem, error) {
	problems, err := Problems()
	if err != nil {
		return Problem{}, err
	}

	for _, problem := range problems {
		if problem.ID == id {
			return problem, nil
		}
	}

	return Problem{}, ErrProblemNotFound
}

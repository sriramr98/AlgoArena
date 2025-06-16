package executors

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/sriramr98/dsa_server/utils"
)

const (
	PISTON_API_HOST_KEY = "PISTON_API_URL"
	PISTON_LOCAL_HOST   = "http://localhost:2000"
)

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type PistonPayload struct {
	Language           string `json:"language"`
	Version            string `json:"version"`
	Files              []File `json:"files"`
	CompileTimeout     int    `json:"compile_timeout"`
	RunTimeout         int    `json:"run_timeout"`
	CompileCpuTime     int    `json:"compile_cpu_time"`
	RunCpuTime         int    `json:"run_cpu_time"`
	CompileMemoryLimit int    `json:"compile_memory_limit"`
	RunMemoryLimit     int    `json:"run_memory_limit"`
}

type PistonExecutor struct{}

func (pe PistonExecutor) Execute(code string, config ExecutorConfig) (ExecutorOutput, error) {
	apiBaseUrl := utils.GetEnv(PISTON_API_HOST_KEY, PISTON_LOCAL_HOST)
	executeUrl := fmt.Sprintf("%s/api/v2/execute", apiBaseUrl)

	payload := PistonPayload{
		Language: config.Language,
		Version:  getLanguageVersion(config.Language),
		Files: []File{
			{
				Name:    fmt.Sprintf("code.%s", getLanguageExtension(config.Language)),
				Content: code,
			},
		},
		CompileTimeout:     10000,
		RunTimeout:         3000,
		CompileCpuTime:     10000,
		RunCpuTime:         3000,
		CompileMemoryLimit: -1,
		RunMemoryLimit:     -1,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		utils.LogError(err)
		return ExecutorOutput{}, err
	}

	log.Println("Payload:", string(jsonPayload))
	request, err := http.NewRequest(http.MethodPost, executeUrl, bytes.NewReader(jsonPayload))
	if err != nil {
		return ExecutorOutput{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		utils.LogError(err)
		return ExecutorOutput{}, err
	}

	defer resp.Body.Close()

	// respbody, err := io.ReadAll(resp.Body)
	// fmt.Println("response: ", respbody, err)

	var respData ExecutorOutput
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return ExecutorOutput{}, err
	}

	if resp.StatusCode != 200 {
		return respData, errors.New(fmt.Sprintf("error executing code with status %d", resp.StatusCode))
	}

	return respData, nil
}

func getLanguageVersion(language string) string {
	switch language {
	case "python":
		return "3.12.0"
	case "javascript":
		return "20.11.1"
	default:
		return ""
	}
}

func getLanguageExtension(language string) string {
	switch language {
	case "python":
		return "py"
	case "javascript":
		return "js"
	default:
		return ""
	}
}

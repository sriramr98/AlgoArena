package executors

type ExecutorConfig struct {
	Language string
}

type RunOutput struct {
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	Output   string `json:"output"`
	Code     int    `json:"code"`
	CpuTime  int    `json:"cpu_time"`
	WallTime int    `json:"wall_time"`
	Memory   int    `json:"memory"`
}

type ExecutorOutput struct {
	Language string    `json:"language"`
	Version  string    `json:"version"`
	Run      RunOutput `json:"run"`
}

type CodeExecutor interface {
	Execute(code string, config ExecutorConfig) (ExecutorOutput, error)
}

func GetExecutor() CodeExecutor {
	return PistonExecutor{}
}

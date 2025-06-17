package stub

import (
	"github.com/sriramr98/dsa_server/problems"
)

type StubGenerator interface {
	Generate(problem problems.Problem) string
}

func GetStubGenerator(language string) StubGenerator {
	switch language {
	case "javascript":
		return JSStubGenerator{}
	case "python":
		return PythonStubGenerator{}
	default:
		return nil
	}
}

package stub

import (
	"log"

	"github.com/sriramr98/dsa_server/problems"
)

type StubGenerator interface {
	Generate(problem problems.Problem) string
}

func GetStubGenerator(language string) StubGenerator {
	switch language {
	case "javascript":
		log.Println("Using JavaScript stub generator")
		return JSStubGenerator{}
	default:
		return nil
	}
}

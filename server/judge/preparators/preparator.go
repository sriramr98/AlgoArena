package preparators

import (
	"errors"

	"github.com/sriramr98/dsa_server/problems"
)

var ErrUnsupportedLanguage error = errors.New("unsupported language")

type Preparator interface {
	Prepare(userCode string, problem problems.Problem, testCase problems.TestCase) (string, error)
}

var preparators = map[string]Preparator{
	"javascript": JsPreparator{},
}

func GetPreparator(language string) (Preparator, error) {
	preparator, ok := preparators[language]
	if !ok {
		return nil, ErrUnsupportedLanguage
	}

	return preparator, nil
}

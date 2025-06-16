// filepath: /Users/sriram.ramasamy/Desktop/learning/dsa/server/judge/formatters/jsonFormatter.go
package formatters

import "encoding/json"

type ValueFormatter[T any] struct{}

func (f ValueFormatter[T]) Format(input string) (T, error) {
	var output T

	// This works even if the input is not a valid JSON, it will format it into the type of T which can be used to convert strings into numbers, booleans etc..
	if err := json.Unmarshal([]byte(input), &output); err != nil {
		return output, err
	}

	return output, nil
}

package main

import (
	"encoding/json"
)

type Formatter[T any] interface {
	Format(input string) T
}

type ArrayFormatter[T any] struct{}

func (f ArrayFormatter[T]) Format(input string) ([]T, error) {
	var output []T

	if err := json.Unmarshal([]byte(input), &output); err != nil {
		return output, err
	}

	return output, nil
}

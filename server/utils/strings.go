package utils

import (
	"iter"
	"strings"
)

func JoinStringSeq(seq iter.Seq[string], separator string) string {
	var builder strings.Builder
	first := true

	for value := range seq {
		if !first {
			builder.WriteString(separator)
		}
		builder.WriteString(value)
		first = false
	}

	return builder.String()
}

func JoinStringSlice(slice []string, separator string) string {
	var builder strings.Builder
	first := true

	for _, value := range slice {
		if !first {
			builder.WriteString(separator)
		}
		builder.WriteString(value)
		first = false
	}

	return builder.String()
}

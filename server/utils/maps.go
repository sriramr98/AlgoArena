package utils

import "slices"

func MapKeysSorted[T any](input map[string]T) []string {
	keys := make([]string, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}

	// Sort the keys
	slices.Sort(keys)
	return keys
}

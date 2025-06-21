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

func FindInMap[T any](input map[string]T, predicate func(string, T) bool) (string, T, bool) {
	for key, value := range input {
		if predicate(key, value) {
			return key, value, true
		}
	}
	var zeroValue T
	return "", zeroValue, false
}

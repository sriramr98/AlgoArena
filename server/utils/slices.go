package utils

func FindInSlice[T any](input []T, predicate func(T) bool) (T, bool) {
	for _, item := range input {
		if predicate(item) {
			return item, true
		}
	}
	var zeroValue T
	return zeroValue, false
}

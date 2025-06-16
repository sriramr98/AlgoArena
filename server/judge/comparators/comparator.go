package comparators

type Comparator[T any] interface {
	Compare(a T, b T) bool
}

package is

// Not returns a function that inverts the result of the given predicate.
func Not[T any](predicate func(T) bool) func(T) bool {
	return func(x T) bool {
		return !predicate(x)
	}
}

package is

// Not returns a function that inverts the result of the given predicate.
func Not[T any](predicate func(T) bool) func(T) bool {
	return func(x T) bool {
		return !predicate(x)
	}
}

// NotOrError returns a function that inverts the result of the given predicate.
func NotOrError[T any](predicate func(T) (bool, error)) func(T) (bool, error) {
	return func(x T) (bool, error) {
		ok, err := predicate(x)
		return !ok, err
	}
}

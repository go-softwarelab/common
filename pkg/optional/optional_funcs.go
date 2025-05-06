package optional

// Map is a function that maps the value of optional if it is present.
func Map[E, R any](o Value[E], f func(E) R) Value[R] {
	if o.IsEmpty() {
		return Empty[R]()
	}

	return Of(f(*o.value))
}

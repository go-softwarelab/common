package optional

// Map is a function that maps the value of optional if it is present.
func Map[E, R any](o Elem[E], f func(E) R) Elem[R] {
	if o.IsEmpty() {
		return Empty[R]()
	}

	return Of(f(*o.value))
}

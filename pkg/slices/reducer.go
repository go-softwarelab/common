package slices

// Reduce applies a function against an accumulator and each element in the slice (from left to right) to reduce it to a single value.
func Reduce[E any, R any](collection []E, accumulator func(agg R, item E) R, initial R) R {
	for _, v := range collection {
		initial = accumulator(initial, v)
	}

	return initial
}

// ReduceRight applies a function against an accumulator and each element in the slice (from right to left) to reduce it to a single value.
func ReduceRight[E any, R any](collection []E, accumulator func(agg R, item E) R, initial R) R {
	for i := len(collection) - 1; i >= 0; i-- {
		initial = accumulator(initial, collection[i])
	}

	return initial
}

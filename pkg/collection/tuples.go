package collection

// Pair creates a pair tuple from given values.
func Pair[A, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{A: a, B: b}
}

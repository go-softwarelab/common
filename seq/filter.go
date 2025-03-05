package seq

import "iter"

// Predicate is a function that takes an element and returns a boolean.
type Predicate[E any] = Mapper[E, bool]

// Filter returns a new sequence that contains only the elements that satisfy the predicate.
func Filter[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range seq {
			if predicate(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

// Where returns a new sequence that contains only the elements that satisfy the predicate.
// SQL-like alias for Filter
func Where[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E] {
	return Filter(seq, predicate)
}

// Skip returns a new sequence that skips the first n elements of the given sequence.
func Skip[E any](seq iter.Seq[E], n int) iter.Seq[E] {
	return func(yield func(E) bool) {
		i := 1
		for v := range seq {
			if i > n {
				if !yield(v) {
					break
				}
			}
			i++
		}
	}
}

// SkipWhile returns a new sequence that skips elements while the predicate is true.
func SkipWhile[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		skip := true
		for v := range seq {
			if skip {
				skip = predicate(v)
			}
			if !skip {
				if !yield(v) {
					break
				}
			}
		}
	}
}

// SkipUntil returns a new sequence that skips elements until the predicate is true.
func SkipUntil[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		skip := true
		for v := range seq {
			if skip {
				skip = !predicate(v)
			}
			if !skip {
				if !yield(v) {
					break
				}
			}
		}
	}
}

// Offset returns a new sequence that skips the first n elements of the given sequence.
// SQL-like alias for Skip
func Offset[E any](seq iter.Seq[E], n int) iter.Seq[E] {
	return Skip(seq, n)
}

// Take returns a new sequence that contains only the first n elements of the given sequence.
func Take[E any](seq iter.Seq[E], n int) iter.Seq[E] {
	return func(yield func(E) bool) {
		i := 1
		for v := range seq {
			if i > n {
				break
			}
			if !yield(v) {
				break
			}
			i++
		}
	}
}

// TakeWhile returns a new sequence that contains elements while the predicate is true.
func TakeWhile[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range seq {
			if !predicate(v) {
				break
			}
			if !yield(v) {
				break
			}
		}
	}
}

// TakeUntil returns a new sequence that contains elements until the predicate is true.
func TakeUntil[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range seq {
			if predicate(v) {
				break
			}
			if !yield(v) {
				break
			}
		}
	}
}

// Limit returns a new sequence that contains only the first n elements of the given sequence.
// SQL-like alias for Take
func Limit[E any](seq iter.Seq[E], n int) iter.Seq[E] {
	return Take(seq, n)
}

// Uniq returns a sequence with only unique elements.
func Uniq[E comparable](seq iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		seen := make(map[E]struct{})
		for v := range seq {
			if _, ok := seen[v]; !ok {
				seen[v] = struct{}{}
				if !yield(v) {
					break
				}
			}
		}
	}
}

// Distinct returns a sequence with only unique elements.
// SQL-like alias for Uniq
func Distinct[E comparable](seq iter.Seq[E]) iter.Seq[E] {
	return Uniq(seq)
}

// UniqBy returns a sequence with only unique elements based on a key.
func UniqBy[E any, K comparable](seq iter.Seq[E], mapper Mapper[E, K]) iter.Seq[E] {
	return func(yield func(E) bool) {
		seen := make(map[K]struct{})
		for v := range seq {
			key := mapper(v)
			if _, ok := seen[key]; !ok {
				seen[key] = struct{}{}
				if !yield(v) {
					break
				}
			}
		}
	}
}

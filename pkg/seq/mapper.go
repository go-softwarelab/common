package seq

import "iter"

// Mapper is a function that maps a value of type T to a value of type R.
type Mapper[T any, R any] = func(T) R

// Map applies a mapper function to each element of the sequence.
func Map[E any, R any](seq iter.Seq[E], mapper Mapper[E, R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			result := mapper(v)
			if !yield(result) {
				break
			}
		}
	}
}

// MapOrErr applies a mapper function which can return error to each element of the sequence.
func MapOrErr[E any, R any](seq iter.Seq[E], mapper func(E) (R, error)) iter.Seq2[R, error] {
	return func(yield func(R, error) bool) {
		for v := range seq {
			result, err := mapper(v)
			if !yield(result, err) {
				break
			}
		}
	}
}

// Select applies a mapper function to each element of the sequence.
// SQL-like alias for Map
func Select[E any, R any](seq iter.Seq[E], mapper Mapper[E, R]) iter.Seq[R] {
	return Map[E, R](seq, mapper)
}

// FlatMap applies a mapper function to each element of the sequence and flattens the result.
func FlatMap[E any, R any](seq iter.Seq[E], mapper Mapper[E, iter.Seq[R]]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			for r := range mapper(v) {
				if !yield(r) {
					return
				}
			}
		}
	}
}

// Flatten flattens a sequence of sequences.
func Flatten[Seq iter.Seq[iter.Seq[E]], E any](seq Seq) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range seq {
			for elem := range v {
				if !yield(elem) {
					return
				}
			}
		}
	}
}

// FlattenSlices flattens a sequence of slices.
func FlattenSlices[Seq iter.Seq[[]E], E any](seq Seq) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range seq {
			for _, elem := range v {
				if !yield(elem) {
					return
				}
			}
		}
	}
}

// Cycle repeats the sequence indefinitely.
func Cycle[E any](seq iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// CycleTimes repeats the sequence specific number of times.
func CycleTimes[E any](seq iter.Seq[E], count int) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := 0; i < count; i++ {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

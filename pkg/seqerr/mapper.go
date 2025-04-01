package seqerr

import (
	"iter"

	"github.com/go-softwarelab/common/pkg/to"
)

// MapSeq applies a mapper function to each element of the sequence.
// The mapper function can return an error.
func MapSeq[E any, R any](seq iter.Seq[E], mapper MapperWithError[E, R]) iter.Seq2[R, error] {
	return func(yield func(R, error) bool) {
		for v := range seq {
			result, err := mapper(v)
			if !yield(result, err) {
				break
			}
		}
	}
}

// Map applies a mapper function to each element of the sequence.
func Map[E any, R any](seq iter.Seq2[E, error], mapper MapperWithoutError[E, R]) iter.Seq2[R, error] {
	return func(yield func(R, error) bool) {
		for v, err := range seq {
			if err != nil {
				yield(to.ZeroValue[R](), err)
				break
			}
			result := mapper(v)
			if !yield(result, nil) {
				break
			}
		}
	}
}

// MapOrErr applies a mapper function that can return error to each element of the sequence.
func MapOrErr[E any, R any](seq iter.Seq2[E, error], mapper MapperWithError[E, R]) iter.Seq2[R, error] {
	return func(yield func(R, error) bool) {
		for v, err := range seq {
			if err != nil {
				yield(to.ZeroValue[R](), err)
				break
			}
			result, err := mapper(v)
			if !yield(result, err) {
				break
			}
		}
	}
}

// FlatMap applies a mapper function to each element of the sequence and flattens the result.
func FlatMap[E any, R any](seq iter.Seq2[E, error], mapper MapperWithoutError[E, iter.Seq[R]]) iter.Seq2[R, error] {
	return func(yield func(R, error) bool) {
		for v, err := range seq {
			if err != nil {
				yield(to.ZeroValue[R](), err)
				break
			}

			collection := mapper(v)
			for element := range collection {
				if !yield(element, nil) {
					return
				}
			}
		}
	}
}

// FlatMapOrErr applies a mapper function that can return error to each element of the sequence and flattens the result.
func FlatMapOrErr[E any, R any](seq iter.Seq2[E, error], mapper MapperWithError[E, iter.Seq[R]]) iter.Seq2[R, error] {
	return func(yield func(R, error) bool) {
		for v, err := range seq {
			if err != nil {
				yield(to.ZeroValue[R](), err)
				break
			}

			collection, err := mapper(v)
			if err != nil {
				yield(to.ZeroValue[R](), err)
				break
			}

			for element := range collection {
				if !yield(element, nil) {
					return
				}
			}
		}
	}
}

// Flatten flattens a sequence of sequences.
func Flatten[Seq iter.Seq2[iter.Seq[E], error], E any](seq Seq) iter.Seq2[E, error] {
	return func(yield func(E, error) bool) {
		for v, err := range seq {
			if err != nil {
				yield(to.ZeroValue[E](), err)
				break
			}

			for elem := range v {
				if !yield(elem, nil) {
					return
				}
			}
		}
	}
}

// FlattenSlices flattens a sequence of slices.
func FlattenSlices[Seq iter.Seq2[[]E, error], E any](seq Seq) iter.Seq2[E, error] {
	return func(yield func(E, error) bool) {
		for v, err := range seq {
			if err != nil {
				yield(to.ZeroValue[E](), err)
				break
			}

			for _, elem := range v {
				if !yield(elem, nil) {
					return
				}
			}
		}
	}
}

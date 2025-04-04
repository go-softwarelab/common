package seqerr

import (
	"iter"
)

// Of creates a new sequence of elements and nil errors, from the provided elements.
func Of[E any](elements ...E) iter.Seq2[E, error] {
	return FromSlice(elements)
}

// FromSlice converts a slice of elements into a sequence of elements and nil errors.
func FromSlice[E any](slice []E) iter.Seq2[E, error] {
	return func(yield func(E, error) bool) {
		for _, elem := range slice {
			if !yield(elem, nil) {
				break
			}
		}
	}
}

// FromSeq converts a sequence of elements into a sequence of elements and nil errors.
func FromSeq[E any](sequence iter.Seq[E]) iter.Seq2[E, error] {
	return func(yield func(E, error) bool) {
		for elem := range sequence {
			if !yield(elem, nil) {
				break
			}
		}
	}
}

// Produce returns a new sequence that is filled by the results of calling the next function.
func Produce[E, A any](next func(A) ([]E, A, error)) iter.Seq2[[]E, error] {
	iterator := &statefulIterator[E, A]{
		next: next,
	}

	return iterator.iterate()
}

// ProduceWithArg returns a new sequence that is filled by the results of calling the next function with the provided argument.
func ProduceWithArg[E, A any](next func(A) ([]E, A, error), arg A) iter.Seq2[[]E, error] {
	iterator := &statefulIterator[E, A]{
		next: next,
		arg:  arg,
	}

	return iterator.iterate()
}

type statefulIterator[E, A any] struct {
	arg  A
	next func(A) ([]E, A, error)
}

func (i *statefulIterator[E, A]) iterate() iter.Seq2[[]E, error] {
	return func(yield func([]E, error) bool) {
		for {
			elems, arg, err := i.next(i.arg)
			if err != nil {
				yield(nil, err)
				break
			}
			if len(elems) == 0 {
				break
			}
			if !yield(elems, nil) {
				break
			}
			i.arg = arg
		}
	}
}

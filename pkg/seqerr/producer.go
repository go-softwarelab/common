package seqerr

import (
	"iter"
)

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

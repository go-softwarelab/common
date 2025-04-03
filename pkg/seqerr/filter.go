package seqerr

import (
	"iter"

	"github.com/go-softwarelab/common/pkg/is"
)

// Filter returns a new sequence that contains only the elements that satisfy the predicate.
func Filter[E any, P Predicate[E]](seq iter.Seq2[E, error], predicate P) iter.Seq2[E, error] {
	filter := toPredicateWithError[E](predicate)

	return func(yield func(E, error) bool) {
		for v, err := range seq {
			if err != nil {
				yield(v, err)
				break
			}

			ok, err := filter(v)
			if err != nil {
				yield(v, err)
				break
			}

			if !ok {
				continue
			}

			if !yield(v, nil) {
				break
			}
		}
	}
}

// Take returns a new sequence that contains only the first n elements of the given sequence.
func Take[E any](seq iter.Seq2[E, error], n int) iter.Seq2[E, error] {
	return func(yield func(E, error) bool) {
		i := 1
		for v, err := range seq {
			if err != nil {
				yield(v, err)
				break
			}
			if i > n {
				break
			}
			if !yield(v, nil) {
				break
			}
			i++
		}
	}
}

// TakeWhile returns a new sequence that takes elements from the given sequence while the predicate is satisfied.
func TakeWhile[E any, P Predicate[E]](seq iter.Seq2[E, error], predicate P) iter.Seq2[E, error] {
	filter := toPredicateWithError[E](predicate)
	return TakeUntil(seq, is.NotOrError(filter))
}

// TakeUntil returns a new sequence that takes elements from the given sequence until the predicate is satisfied.
func TakeUntil[E any, P Predicate[E]](seq iter.Seq2[E, error], predicate P) iter.Seq2[E, error] {
	filter := toPredicateWithError[E](predicate)
	return func(yield func(E, error) bool) {
		for v, err := range seq {
			if err != nil {
				yield(v, err)
				break
			}

			ok, err := filter(v)
			if err != nil {
				yield(v, err)
				break
			}

			if ok || !yield(v, nil) {
				break
			}
		}
	}
}

// TakeWhileTrue returns a new sequence that takes elements from the given sequence while the stop condition is satisfied.
func TakeWhileTrue[E any](seq iter.Seq2[E, error], stopCondition func() bool) iter.Seq2[E, error] {
	return TakeUntilTrue(seq, func() bool {
		return !stopCondition()
	})
}

// TakeUntilTrue returns a new sequence that takes elements from the given sequence until the stop condition is satisfied.
func TakeUntilTrue[E any](seq iter.Seq2[E, error], stopCondition func() bool) iter.Seq2[E, error] {
	return func(yield func(E, error) bool) {
		for v, err := range seq {
			if err != nil {
				yield(v, err)
				break
			}

			if stopCondition() || !yield(v, nil) {
				break
			}
		}
	}
}

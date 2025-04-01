package seqerr

import (
	"iter"
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

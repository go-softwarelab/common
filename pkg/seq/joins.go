package seq

import (
	"iter"

	"github.com/go-softwarelab/common/pkg/types"
)

// Concat concatenates multiple sequences into a single sequence.
func Concat[E any](sequences ...iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, seq := range sequences {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Union returns a sequence that contains all distinct elements from both input sequences.
func Union[E types.Comparable](seq1 iter.Seq[E], seq2 iter.Seq[E]) iter.Seq[E] {
	return Distinct(UnionAll(seq1, seq2))
}

// UnionAll returns a sequence that contains all elements from both input sequences.
func UnionAll[E any](seq1 iter.Seq[E], seq2 iter.Seq[E]) iter.Seq[E] {
	return Concat(seq1, seq2)
}

// Append appends elements to the end of a sequence.
func Append[E any](seq iter.Seq[E], elems ...E) iter.Seq[E] {
	return Concat(seq, Of(elems...))
}

// Prepend prepends elements to the beginning of a sequence.
func Prepend[E any](seq iter.Seq[E], elems ...E) iter.Seq[E] {
	return Concat(Of(elems...), seq)
}

// Zip combines two sequences into a iter.Seq2.
func Zip[E any, R any](seq1 iter.Seq[E], seq2 iter.Seq[R]) iter.Seq2[E, R] {
	return func(yield func(E, R) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()
		next2, stop2 := iter.Pull(seq2)
		defer stop2()

		for {
			v1, exist1 := next1()
			v2, exist2 := next2()
			if !exist1 && !exist2 {
				break
			}
			if !yield(v1, v2) {
				break
			}
		}
	}
}

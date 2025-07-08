package seq2

import (
	"iter"

	"github.com/go-softwarelab/common/pkg/seq"
)

// Concat concatenates multiple sequences into a single sequence.
// It also safely handles nil iterators treating them as an empty iterator.
func Concat[K any, V any](sequences ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range sequences {
			if seq == nil {
				continue
			}
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Union returns a sequence that contains all distinct elements from both input sequences.
func Union[K comparable, V comparable](seq1 iter.Seq2[K, V], seq2 iter.Seq2[K, V]) iter.Seq2[K, V] {
	return Distinct(UnionAll(seq1, seq2))
}

// UnionAll returns a sequence that contains all elements from both input sequences.
func UnionAll[K any, V any](seq1 iter.Seq2[K, V], seq2 iter.Seq2[K, V]) iter.Seq2[K, V] {
	return Concat(seq1, seq2)
}

// UnZip splits a sequence of pairs into two sequences.
func UnZip[K any, V any](seq iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V]) {
	return Keys(seq), Values(seq)
}

// Split splits a sequence of pairs into two sequences.
func Split[K any, V any](sequence iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V]) {
	if sequence == nil {
		return seq.Empty[K](), seq.Empty[V]()
	}

	return Keys(sequence), Values(sequence)
}

// Append appends element to the end of a sequence.
func Append[K any, V any](seq iter.Seq2[K, V], key K, value V) iter.Seq2[K, V] {
	return Concat(seq, Single(key, value))
}

// Prepend prepends element to the beginning of a sequence.
func Prepend[K any, V any](seq iter.Seq2[K, V], key K, value V) iter.Seq2[K, V] {
	return Concat(Single(key, value), seq)
}

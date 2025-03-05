package seq2

import "iter"

// Concat concatenates multiple sequences into a single sequence.
func Concat[K any, V any](sequences ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range sequences {
			for k, v := range seq {
				if !yield(k, v) {
					break
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
	return func(yield func(K, V) bool) {
		for k, v := range seq1 {
			if !yield(k, v) {
				break
			}
		}
		for k, v := range seq2 {
			if !yield(k, v) {
				break
			}
		}
	}
}

// UnZip splits a sequence of pairs into two sequences.
func UnZip[K any, V any](seq iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V]) {
	return Keys(seq), Values(seq)
}

// Split splits a sequence of pairs into two sequences.
func Split[K any, V any](seq iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V]) {
	return Keys(seq), Values(seq)
}

// Append appends element to the end of a sequence.
func Append[K any, V any](seq iter.Seq2[K, V], key K, value V) iter.Seq2[K, V] {
	return Concat(seq, Pair(key, value))
}

// Prepend prepends element to the beginning of a sequence.
func Prepend[K any, V any](seq iter.Seq2[K, V], key K, value V) iter.Seq2[K, V] {
	return Concat(Pair(key, value), seq)
}

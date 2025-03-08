package seq2

import "iter"

// Keys returns a sequence of keys from a sequence of key-value pairs.
func Keys[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range seq {
			if !yield(k) {
				break
			}
		}
	}
}

// Values returns a sequence of values from a sequence of key-value pairs.
func Values[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				break
			}
		}
	}
}

// IsEmpty returns true if the sequence is empty.
func IsEmpty[K, V any](seq iter.Seq2[K, V]) bool {
	for range seq {
		return false
	}
	return true
}

// IsNotEmpty returns true if the sequence is not empty.
func IsNotEmpty[K, V any](seq iter.Seq2[K, V]) bool {
	return !IsEmpty(seq)
}

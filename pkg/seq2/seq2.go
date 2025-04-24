package seq2

import "iter"

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

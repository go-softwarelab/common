package seq

import (
	"iter"
)

// IsNotEmpty returns true if the sequence is not empty.
func IsNotEmpty[E any](seq iter.Seq[E]) bool {
	for range seq {
		return true
	}
	return false
}

// IsEmpty returns true if the sequence is empty.
func IsEmpty[E any](seq iter.Seq[E]) bool {
	return !IsNotEmpty(seq)
}

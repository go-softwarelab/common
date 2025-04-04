package seq

import (
	"iter"
	"slices"
)

// Of creates a new sequence from the given elements.
func Of[E any](elems ...E) iter.Seq[E] {
	return FromSlice(elems)
}

// FromSlice creates a new sequence from the given slice.
func FromSlice[Slice ~[]E, E any](slice Slice) iter.Seq[E] {
	return slices.Values(slice)
}

// PointersFromSlice creates a new sequence of pointers for the given slice of value elements.
func PointersFromSlice[Slice ~[]E, E any](slice Slice) iter.Seq[*E] {
	return func(yield func(*E) bool) {
		for i := range slice {
			if !yield(&slice[i]) {
				break
			}
		}
	}
}

// Reverse creates a new sequence that iterates over the elements of the given sequence in reverse order.
func Reverse[E any](seq iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		slice := slices.Collect(seq)
		for i := len(slice) - 1; i >= 0; i-- {
			if !yield(slice[i]) {
				break
			}
		}
	}
}

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

package seqerr

import (
	"iter"
)

// ToSlice collects the elements of the given sequence into a slice.
func ToSlice[Slice ~[]E, E any](seq iter.Seq2[E, error], slice Slice) (Slice, error) {
	for v, err := range seq {
		if err != nil {
			return slice, err
		}
		slice = append(slice, v)
	}
	return slice, nil
}

// Collect collects the elements of the given sequence into a slice.
func Collect[E any](seq iter.Seq2[E, error]) ([]E, error) {
	return ToSlice(seq, []E(nil))
}

// Count returns the number of elements in the sequence.
func Count[E any](seq iter.Seq2[E, error]) (int, error) {
	i := 0
	for _, err := range seq {
		if err != nil {
			return 0, err
		}
		i++
	}
	return i, nil
}

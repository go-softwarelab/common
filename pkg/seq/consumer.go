package seq

import (
	"iter"
	"slices"
)

// Consumer is a function that consumes an element of a sequence.
type Consumer[E any] = func(E)

// Tap returns a sequence that applies the given consumer to each element of the input sequence and pass it further.
func Tap[E any](seq iter.Seq[E], consumer func(E)) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range seq {
			consumer(v)
			if !yield(v) {
				break
			}
		}
	}
}

// Each is an alias for Tap.
func Each[E any](seq iter.Seq[E], consumer Consumer[E]) iter.Seq[E] {
	return Tap(seq, consumer)
}

// ForEach applies consumer to each element of the input sequence.
func ForEach[E any](seq iter.Seq[E], consumer Consumer[E]) {
	for v := range seq {
		consumer(v)
	}
}

// Flush consumes all elements of the input sequence.
func Flush[E any](seq iter.Seq[E]) {
	for range seq {
	}
}

// ToSlice collects the elements of the given sequence into a slice.
func ToSlice[Slice ~[]E, E any](seq iter.Seq[E], slice Slice) Slice {
	return slices.AppendSeq(slice, seq)
}

// Collect collects the elements of the given sequence into a slice.
func Collect[E any](seq iter.Seq[E]) []E {
	return slices.Collect(seq)
}

// Count returns the number of elements in the sequence.
func Count[E any](seq iter.Seq[E]) int {
	i := 0
	for range seq {
		i++
	}
	return i
}

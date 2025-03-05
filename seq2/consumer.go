package seq2

import (
	"iter"
	"maps"
)

// Consumer is a function that consumes an element of an iter.Seq2.
type Consumer[K any, V any] = func(K, V)

// Tap returns a sequence that applies the given consumer to each element of the input sequence and pass it further.
func Tap[K any, V any](seq iter.Seq2[K, V], consumer Consumer[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for v, r := range seq {
			consumer(v, r)
			if !yield(v, r) {
				break
			}
		}
	}
}

// Each is an alias for Tap.
func Each[K any, V any](seq iter.Seq2[K, V], consumer Consumer[K, V]) iter.Seq2[K, V] {
	return Tap(seq, consumer)
}

// ForEach applies consumer to each element of the input sequence.
func ForEach[K any, V any](seq iter.Seq2[K, V], consumer Consumer[K, V]) {
	for v, r := range seq {
		consumer(v, r)
	}
}

// Flush consumes all elements of the input sequence.
func Flush[K any, V any](seq iter.Seq2[K, V]) {
	for range seq {
	}
}

// ToMap collects the elements of the given sequence into a map.
func ToMap[Map ~map[K]V, K comparable, V any](seq iter.Seq2[K, V], m Map) {
	maps.Insert(m, seq)
}

// Collect collects the elements of the given sequence into a map.
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	return maps.Collect(seq)
}

// Count returns the number of elements in the sequence.
func Count[K any, V any](seq iter.Seq2[K, V]) int {
	i := 0
	for range seq {
		i++
	}
	return i
}

package seq2

import (
	"iter"
	"maps"
	"slices"

	"github.com/go-softwarelab/common/pkg/types"
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

// Each returns a sequence that applies the given consumer to each element of the input sequence and pass it further.
// Each is an alias for Tap.
// Comparing to ForEach, this is a lazy function and doesn't consume the input sequence.
func Each[K any, V any](seq iter.Seq2[K, V], consumer Consumer[K, V]) iter.Seq2[K, V] {
	return Tap(seq, consumer)
}

// ForEach applies consumer to each element of the input sequence.
// Comparing to Each, this is not a lazy function and consumes the input sequence.
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

// CollectToMap collects the elements of the given sequence into a map.
func CollectToMap[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	return maps.Collect(seq)
}

// Collect collects the elements of the given sequence into a slice of types.Pair of K and V.
func Collect[K comparable, V any](seq iter.Seq2[K, V]) []types.Pair[K, V] {
	pairs := MapTo(seq, func(k K, v V) types.Pair[K, V] {
		return *types.NewPair(k, v)
	})
	return slices.Collect(pairs)
}

// Count returns the number of elements in the sequence.
func Count[K any, V any](seq iter.Seq2[K, V]) int {
	i := 0
	for range seq {
		i++
	}
	return i
}

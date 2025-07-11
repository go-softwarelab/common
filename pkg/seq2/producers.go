package seq2

import (
	"iter"
	"slices"
	"time"

	"github.com/go-softwarelab/common/pkg/types"
)

// Empty returns an empty sequence.
func Empty[K, V any]() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {}
}

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

// Single returns a sequence with given key value.
func Single[K, V any](k K, v V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		yield(k, v)
	}
}

// OfIndexed creates a new indexed sequence from the given elements.
func OfIndexed[E any](elems ...E) iter.Seq2[int, E] {
	return slices.All(elems)
}

// FromSlice creates a new sequence from the given slice with index as keys.
func FromSlice[Slice ~[]E, E any](slice Slice) iter.Seq2[int, E] {
	return slices.All(slice)
}

// FromMap creates a new iter.Seq2 from the given map.
func FromMap[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				break
			}
		}
	}
}

// WithIndex creates a new indexed sequence from the given sequence.
func WithIndex[E any](seq iter.Seq[E]) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		i := 0
		for v := range seq {
			if !yield(i, v) {
				break
			}
			i++
		}
	}
}

// WithoutIndex creates a new sequence from the given indexed sequence.
func WithoutIndex[E any](indexed iter.Seq2[int, E]) iter.Seq[E] {
	return Values(indexed)
}

// Repeat repeats the given pair `count` times.
func Repeat[K any, V any, N types.Integer](key K, value V, count N) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := N(0); i < count; i++ {
			if !yield(key, value) {
				break
			}
		}
	}
}

// Tick returns a sequence that yields the tick number and the current time every duration.
func Tick(d time.Duration) iter.Seq2[int, time.Time] {
	return func(yield func(int, time.Time) bool) {
		idx := 0

		ticker := time.NewTicker(d)
		defer ticker.Stop()

		//nolint:gosimple
		for {
			select {
			case t := <-ticker.C:
				// thanks to that, it's starting on 1
				idx++
				if !yield(idx, t) {
					return
				}
			}
		}
	}
}

// Reverse reverses the given sequence.
func Reverse[K, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	type pair struct {
		k K
		v V
	}
	seqOfPairs := MapTo[K, V, *pair](seq, func(k K, v V) *pair {
		return &pair{k, v}
	})

	return func(yield func(K, V) bool) {
		pairs := slices.Collect(seqOfPairs)
		for i := len(pairs) - 1; i >= 0; i-- {
			if !yield(pairs[i].k, pairs[i].v) {
				break
			}
		}
	}
}

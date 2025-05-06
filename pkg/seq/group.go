package seq

import (
	"iter"
	"slices"
)

// Partition splits the sequence into chunks of the given size.
func Partition[E any](seq iter.Seq[E], size int) iter.Seq[iter.Seq[E]] {
	if size <= 0 {
		panic("size must be greater than 0")
	}
	return func(yield func(iter.Seq[E]) bool) {
		chunk := make([]E, 0, size)
		for v := range seq {
			chunk = append(chunk, v)
			if len(chunk) == size {
				if !yield(slices.Values(chunk)) {
					break
				}
				chunk = make([]E, 0, size)
			}
		}
		if len(chunk) > 0 {
			yield(slices.Values(chunk))
		}
	}
}

// Chunk splits the sequence into chunks of the given size.
func Chunk[E any](seq iter.Seq[E], size int) iter.Seq[iter.Seq[E]] {
	return Partition(seq, size)
}

// PartitionBy splits the sequence into chunks based on the given key.
// It splits the sequence when ever the key changes, the order matters here.
func PartitionBy[E any, K comparable](seq iter.Seq[E], by Mapper[E, K]) iter.Seq[iter.Seq[E]] {
	return func(yield func(iter.Seq[E]) bool) {
		var lastKey K
		chunk := make([]E, 0)
		for v := range seq {
			key := by(v)
			if key == lastKey {
				chunk = append(chunk, v)
			} else {
				if !yield(FromSlice(chunk)) {
					return
				}
				chunk = make([]E, 0)
				chunk = append(chunk, v)
				lastKey = key
			}
		}
		yield(FromSlice(chunk))
	}
}

// GroupBy groups the sequence by the given key.
func GroupBy[E any, K comparable](seq iter.Seq[E], by Mapper[E, K]) iter.Seq2[K, iter.Seq[E]] {
	return func(yield func(K, iter.Seq[E]) bool) {
		result := make(map[K][]E)
		for v := range seq {
			key := by(v)
			result[key] = append(result[key], v)
		}
		for k, v := range result {
			if !yield(k, slices.Values(v)) {
				break
			}
		}
	}
}

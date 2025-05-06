package seq2

import (
	"iter"
)

// Reduce applies a function against an accumulator and each element in the sequence (from left to right) to reduce it to a single value.
func Reduce[K any, V any, R any](seq2 iter.Seq2[K, V], accumulator func(agg R, key K, value V) R, initial R) R {
	result := initial
	for k, v := range seq2 {
		result = accumulator(result, k, v)
	}
	return result
}

// ReduceRight applies a function against an accumulator and each element in the sequence (from right to left) to reduce it to a single value.
func ReduceRight[K any, V any, R any](seq2 iter.Seq2[K, V], accumulator func(agg R, key K, value V) R, initial R) R {
	return Reduce(Reverse(seq2), accumulator, initial)
}

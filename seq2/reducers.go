package seq2

import (
	"iter"

	"github.com/go-softwarelab/common/optional"
	"github.com/go-softwarelab/common/seq"
	"github.com/go-softwarelab/common/types"
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

// FoldValues applies a function against an accumulator and each element in the sequence (from left to right) to reduce it to a single value.
// Notice: first key will be omitted and the first value will be used as initial value.
func FoldValues[K any, V any](seq2 iter.Seq2[K, V], accumulator func(agg V, key K, value V) V) optional.Elem[V] {
	next, stop := iter.Pull2(seq2)
	defer stop()

	k, v, ok := next()
	if !ok {
		return optional.Empty[V]()
	}

	result := v
	for {
		k, v, ok = next()
		if !ok {
			break
		}
		result = accumulator(result, k, v)
	}

	return optional.Of(result)
}

// FoldValuesRight applies a function against an accumulator and each element in the sequence (from right to left) to reduce it to a single value.
// Notice: last key will be omitted and the last value will be used as initial value.
func FoldValuesRight[K any, V any](seq2 iter.Seq2[K, V], accumulator func(agg V, key K, value V) V) optional.Elem[V] {
	return FoldValues(Reverse(seq2), accumulator)
}

// FoldKeys applies a function against an accumulator and each element in the sequence (from left to right) to reduce it to a single value.
// Notice: first value will be omitted and the first key will be used as initial value.
func FoldKeys[K any, V any](seq2 iter.Seq2[K, V], accumulator func(agg K, key K, value V) K) optional.Elem[K] {
	next, stop := iter.Pull2(seq2)
	defer stop()

	k, v, ok := next()
	if !ok {
		return optional.Empty[K]()
	}

	result := k
	for {
		k, v, ok = next()
		if !ok {
			break
		}
		result = accumulator(result, k, v)
	}

	return optional.Of(result)
}

// FoldKeysRight applies a function against an accumulator and each element in the sequence (from right to left) to reduce it to a single value.
// Notice: last value will be omitted and the last key will be used as initial value.
func FoldKeysRight[K any, V any](seq2 iter.Seq2[K, V], accumulator func(agg K, key K, value V) K) optional.Elem[K] {
	return FoldKeys(Reverse(seq2), accumulator)
}

// MaxValue returns the maximum value in the sequence.
func MaxValue[K any, V types.Ordered](seq2 iter.Seq2[K, V]) optional.Elem[V] {
	vals := Values(seq2)
	return seq.Max(vals)
}

// MinValue returns the minimum value in the .
func MinValue[K any, V types.Ordered](seq2 iter.Seq2[K, V]) optional.Elem[V] {
	vals := Values(seq2)
	return seq.Min(vals)
}

// MaxKey returns the maximum key in the sequence.
func MaxKey[K types.Ordered, V any](seq2 iter.Seq2[K, V]) optional.Elem[K] {
	keys := Keys(seq2)
	return seq.Max(keys)
}

// MinKey returns the minimum key in the sequence.
func MinKey[K types.Ordered, V any](seq2 iter.Seq2[K, V]) optional.Elem[K] {
	keys := Keys(seq2)
	return seq.Min(keys)
}

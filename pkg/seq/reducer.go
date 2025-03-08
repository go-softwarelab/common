package seq

import (
	"iter"

	"github.com/go-softwarelab/common/pkg/optional"
	"github.com/go-softwarelab/common/types"
)

// Reduce applies a function against an accumulator and each element in the sequence (from left to right) to reduce it to a single value.
func Reduce[E any, R any](seq iter.Seq[E], accumulator func(agg R, item E) R, initial R) R {
	result := initial
	for v := range seq {
		result = accumulator(result, v)
	}
	return result
}

// ReduceRight applies a function against an accumulator and each element in the sequence (from right to left) to reduce it to a single value.
func ReduceRight[E any, R any](seq iter.Seq[E], accumulator func(agg R, item E) R, initial R) R {
	return Reduce(Reverse(seq), accumulator, initial)
}

// Fold applies a function against an accumulator and each element in the sequence (from left to right) to reduce it to a single value.
func Fold[E any](seq iter.Seq[E], accumulator func(agg E, item E) E) optional.Elem[E] {
	next, stop := iter.Pull(seq)
	defer stop()

	result, ok := next()
	if !ok {
		return optional.Empty[E]()
	}

	for {
		v, ok := next()
		if !ok {
			break
		}
		result = accumulator(result, v)
	}

	return optional.Of(result)
}

// FoldRight applies a function against an accumulator and each element in the sequence (from right to left) to reduce it to a single value.
func FoldRight[E any](seq iter.Seq[E], accumulator func(agg E, item E) E) optional.Elem[E] {
	return Fold(Reverse(seq), accumulator)
}

// Max returns the maximum element in the sequence.
func Max[E types.Ordered](seq iter.Seq[E]) optional.Elem[E] {
	var result E
	next, stop := iter.Pull(seq)
	defer stop()

	result, ok := next()
	if !ok {
		return optional.Empty[E]()
	}

	for {
		v, ok := next()
		if !ok {
			break
		}
		if v > result {
			result = v
		}
	}
	return optional.Of(result)
}

// Min returns the minimum element in the sequence.
func Min[E types.Ordered](seq iter.Seq[E]) optional.Elem[E] {
	var result E
	next, stop := iter.Pull(seq)
	defer stop()

	result, ok := next()
	if !ok {
		return optional.Empty[E]()
	}

	for {
		v, ok := next()
		if !ok {
			break
		}
		if v < result {
			result = v
		}
	}
	return optional.Of(result)
}

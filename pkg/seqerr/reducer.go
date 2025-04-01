package seqerr

import (
	"iter"

	"github.com/go-softwarelab/common/pkg/optional"
	"github.com/go-softwarelab/common/pkg/seq2"
	"github.com/go-softwarelab/common/pkg/to"
)

// Reduce applies a function against an accumulator and each element in the sequence (from left to right) to reduce it to a single value.
func Reduce[E any, R any](seq iter.Seq2[E, error], accumulator func(agg R, item E) R, initial R) (R, error) {
	result := initial
	for v, err := range seq {
		if err != nil {
			return to.ZeroValue[R](), err
		}
		result = accumulator(result, v)
	}
	return result, nil
}

// ReduceRight applies a function against an accumulator and each element in the sequence (from right to left) to reduce it to a single value.
func ReduceRight[E any, R any](seq iter.Seq2[E, error], accumulator func(agg R, item E) R, initial R) (R, error) {
	return Reduce(seq2.Reverse(seq), accumulator, initial)
}

// Fold applies a function against an accumulator and each element in the sequence (from left to right) to reduce it to a single value.
func Fold[E any](seq iter.Seq2[E, error], accumulator func(agg E, item E) E) (optional.Elem[E], error) {
	next, stop := iter.Pull2(seq)
	defer stop()

	result, err, ok := next()
	if !ok {
		return optional.Empty[E](), nil
	}
	if err != nil {
		return optional.Empty[E](), err
	}

	for {
		v, err, ok := next()
		if !ok {
			break
		}
		if err != nil {
			return optional.Empty[E](), err
		}

		result = accumulator(result, v)
	}

	return optional.Of(result), nil
}

// FoldRight applies a function against an accumulator and each element in the sequence (from right to left) to reduce it to a single value.
func FoldRight[E any](seq iter.Seq2[E, error], accumulator func(agg E, item E) E) (optional.Elem[E], error) {
	return Fold(seq2.Reverse(seq), accumulator)
}

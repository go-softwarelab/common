package seqerr

import (
	"iter"

	"github.com/go-softwarelab/common/pkg/seq2"
)

// Concat concatenates multiple sequences into a single sequence.
func Concat[E any](sequences ...iter.Seq2[E, error]) iter.Seq2[E, error] {
	return seq2.Concat(sequences...)
}

// Append appends element to the end of a sequence.
func Append[E any](seq iter.Seq2[E, error], elem E) iter.Seq2[E, error] {
	return seq2.Append(seq, elem, nil)
}

// Prepend prepends element to the beginning of a sequence.
func Prepend[E any](seq iter.Seq2[E, error], elem E) iter.Seq2[E, error] {
	return seq2.Prepend(seq, elem, nil)
}

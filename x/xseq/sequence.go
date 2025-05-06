package xseq

import (
	"iter"

	"github.com/go-softwarelab/common/pkg/optional"
	"github.com/go-softwarelab/common/pkg/seq"
)

// Sequence is a monad representing a sequence of elements.
type Sequence[E comparable] struct {
	seq iter.Seq[E]
}

// AsSequence wraps an iter.Seq to provide a possibility to pipe several method calls.
func AsSequence[E comparable](seq iter.Seq[E]) Sequence[E] {
	return Sequence[E]{seq}
}

// ConcatSequences concatenates multiple sequences into a single sequence.
func ConcatSequences[E comparable](sequences ...Sequence[E]) Sequence[E] {
	return AsSequence(func(yield func(E) bool) {
		for _, s := range sequences {
			for v := range s.seq {
				if !yield(v) {
					break
				}
			}
		}
	})
}

// Union returns a new sequence that contains all distinct elements from both input sequences.
func (s Sequence[E]) Union(other Sequence[E]) Sequence[E] {
	return AsSequence(seq.Union(s.seq, other.seq))
}

// UnionAll returns a new sequence that contains all elements from both input sequences.
func (s Sequence[E]) UnionAll(other Sequence[E]) Sequence[E] {
	return AsSequence(seq.UnionAll(s.seq, other.seq))
}

// Append appends elements to the end of a sequence.
func (s Sequence[E]) Append(elems ...E) Sequence[E] {
	return AsSequence(seq.Append(s.seq, elems...))
}

// Prepend prepends elements to the beginning of a sequence.
func (s Sequence[E]) Prepend(elems ...E) Sequence[E] {
	return AsSequence(seq.Prepend(s.seq, elems...))
}

// Filter returns a new sequence with elements that satisfy the predicate.
func (s Sequence[E]) Filter(predicate seq.Predicate[E]) Sequence[E] {
	return AsSequence(seq.Filter(s.seq, predicate))
}

// Where returns a new sequence with elements that satisfy the predicate.
// SQL-like alias for Filter
func (s Sequence[E]) Where(predicate seq.Predicate[E]) Sequence[E] {
	return AsSequence(seq.Where(s.seq, predicate))
}

// Skip returns a new sequence that skips the first n elements of the given sequence.
func (s Sequence[E]) Skip(n int) Sequence[E] {
	return AsSequence(seq.Skip(s.seq, n))
}

// Offset returns a new sequence that skips the first n elements of the given sequence.
// SQL-like alias for Skip
func (s Sequence[E]) Offset(n int) Sequence[E] {
	return AsSequence(seq.Offset(s.seq, n))
}

// Take returns a new sequence that contains only the first n elements of the given sequence.
func (s Sequence[E]) Take(n int) Sequence[E] {
	return AsSequence(seq.Take(s.seq, n))
}

// Limit returns a new sequence that contains only the first n elements of the given sequence.
// SQL-like alias for Take
func (s Sequence[E]) Limit(n int) Sequence[E] {
	return AsSequence(seq.Limit(s.seq, n))
}

// Tap returns a new sequence that calls the consumer for each element of the sequence.
func (s Sequence[E]) Tap(consumer func(E)) Sequence[E] {
	return AsSequence(seq.Tap(s.seq, consumer))
}

// Each returns a new sequence that calls the consumer for each element of the sequence.
func (s Sequence[E]) Each(consumer seq.Consumer[E]) Sequence[E] {
	return AsSequence(seq.Each(s.seq, consumer))
}

// ForEach calls the consumer for each element of the sequence.
func (s Sequence[E]) ForEach(consumer seq.Consumer[E]) {
	seq.ForEach(s.seq, consumer)
}

// Flush consumes all elements of the input sequence.
func (s Sequence[E]) Flush() {
	seq.Flush(s.seq)
}

// Collect collects the elements of the sequence into a slice.
func (s Sequence[E]) Collect() []E {
	return seq.Collect(s.seq)
}

// Count returns the number of elements in the sequence.
func (s Sequence[E]) Count() int {
	return seq.Count(s.seq)
}

// ToSlice collects the elements of the sequence into a given slice.
func (s Sequence[E]) ToSlice(slice []E) []E {
	return seq.ToSlice(s.seq, slice)
}

// Find returns the first element that satisfies the predicate.
func (s Sequence[E]) Find(predicate seq.Predicate[E]) optional.Value[E] {
	return seq.Find(s.seq, predicate)
}

// FindLast returns the last element that satisfies the predicate.
func (s Sequence[E]) FindLast(predicate seq.Predicate[E]) optional.Value[E] {
	return seq.FindLast(s.seq, predicate)
}

// FindAll returns all elements that satisfy the predicate.
func (s Sequence[E]) FindAll(predicate seq.Predicate[E]) Sequence[E] {
	return AsSequence(seq.FindAll(s.seq, predicate))
}

// Contains returns true if the element is in the sequence.
func (s Sequence[E]) Contains(elem E) bool {
	return seq.Contains(s.seq, elem)
}

// NotContains returns true if the element is not in the sequence.
func (s Sequence[E]) NotContains(elem E) bool {
	return seq.NotContains(s.seq, elem)
}

// ContainsAll returns true if all elements are in the sequence.
func (s Sequence[E]) ContainsAll(elements ...E) bool {
	return seq.ContainsAll(s.seq, elements...)
}

// Exists returns true if there is at least one element that satisfies the predicate.
func (s Sequence[E]) Exists(predicate seq.Predicate[E]) bool {
	return seq.Exists(s.seq, predicate)
}

// Every returns true if all elements satisfy the predicate.
func (s Sequence[E]) Every(predicate seq.Predicate[E]) bool {
	return seq.Every(s.seq, predicate)
}

// None returns true if no element satisfies the predicate.
func (s Sequence[E]) None(predicate seq.Predicate[E]) bool {
	return seq.None(s.seq, predicate)
}

// IsNotEmpty returns true if the sequence is not empty.
func (s Sequence[E]) IsNotEmpty() bool {
	return seq.IsNotEmpty(s.seq)
}

// IsEmpty returns true if the sequence is empty.
func (s Sequence[E]) IsEmpty() bool {
	return seq.IsEmpty(s.seq)
}

// Partition splits the sequence into chunks of the given size.
func (s Sequence[E]) Partition(size int) iter.Seq[iter.Seq[E]] {
	return seq.Partition(s.seq, size)
}

// Uniq returns a new sequence with only unique elements.
func (s Sequence[E]) Uniq() Sequence[E] {
	return AsSequence(seq.Uniq(s.seq))
}

// Distinct returns a new sequence with only unique elements.
// SQL-like alias for Uniq
func (s Sequence[E]) Distinct() Sequence[E] {
	return AsSequence(seq.Distinct(s.seq))
}

// Reverse returns a new sequence with elements in reverse order.
func (s Sequence[E]) Reverse() Sequence[E] {
	return AsSequence(seq.Reverse(s.seq))
}

// Fold applies a function against an accumulator and each element in the sequence (from left to right) to reduce it to a single value.
func (s Sequence[E]) Fold(accumulator func(agg E, item E) E) optional.Value[E] {
	return seq.Fold(s.seq, accumulator)
}

// FoldRight applies a function against an accumulator and each element in the sequence (from right to left) to reduce it to a single value.
func (s Sequence[E]) FoldRight(accumulator func(agg E, item E) E) optional.Value[E] {
	return seq.FoldRight(s.seq, accumulator)
}

package seq

import (
	"iter"

	"github.com/go-softwarelab/common/pkg/optional"
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
	return AsSequence(Union(s.seq, other.seq))
}

// UnionAll returns a new sequence that contains all elements from both input sequences.
func (s Sequence[E]) UnionAll(other Sequence[E]) Sequence[E] {
	return AsSequence(UnionAll(s.seq, other.seq))
}

// Append appends elements to the end of a sequence.
func (s Sequence[E]) Append(elems ...E) Sequence[E] {
	return AsSequence(Append(s.seq, elems...))
}

// Prepend prepends elements to the beginning of a sequence.
func (s Sequence[E]) Prepend(elems ...E) Sequence[E] {
	return AsSequence(Prepend(s.seq, elems...))
}

// Filter returns a new sequence with elements that satisfy the predicate.
func (s Sequence[E]) Filter(predicate Predicate[E]) Sequence[E] {
	return AsSequence(Filter(s.seq, predicate))
}

// Where returns a new sequence with elements that satisfy the predicate.
// SQL-like alias for Filter
func (s Sequence[E]) Where(predicate Predicate[E]) Sequence[E] {
	return AsSequence(Where(s.seq, predicate))
}

// Skip returns a new sequence that skips the first n elements of the given sequence.
func (s Sequence[E]) Skip(n int) Sequence[E] {
	return AsSequence(Skip(s.seq, n))
}

// Offset returns a new sequence that skips the first n elements of the given sequence.
// SQL-like alias for Skip
func (s Sequence[E]) Offset(n int) Sequence[E] {
	return AsSequence(Offset(s.seq, n))
}

// Take returns a new sequence that contains only the first n elements of the given sequence.
func (s Sequence[E]) Take(n int) Sequence[E] {
	return AsSequence(Take(s.seq, n))
}

// Limit returns a new sequence that contains only the first n elements of the given sequence.
// SQL-like alias for Take
func (s Sequence[E]) Limit(n int) Sequence[E] {
	return AsSequence(Limit(s.seq, n))
}

// Tap returns a new sequence that calls the consumer for each element of the sequence.
func (s Sequence[E]) Tap(consumer func(E)) Sequence[E] {
	return AsSequence(Tap(s.seq, consumer))
}

// Each returns a new sequence that calls the consumer for each element of the sequence.
func (s Sequence[E]) Each(consumer Consumer[E]) Sequence[E] {
	return AsSequence(Each(s.seq, consumer))
}

// ForEach calls the consumer for each element of the sequence.
func (s Sequence[E]) ForEach(consumer Consumer[E]) {
	ForEach(s.seq, consumer)
}

// Flush consumes all elements of the input sequence.
func (s Sequence[E]) Flush() {
	Flush(s.seq)
}

// Collect collects the elements of the sequence into a slice.
func (s Sequence[E]) Collect() []E {
	return Collect(s.seq)
}

// Count returns the number of elements in the sequence.
func (s Sequence[E]) Count() int {
	return Count(s.seq)
}

// ToSlice collects the elements of the sequence into a given slice.
func (s Sequence[E]) ToSlice(slice []E) []E {
	return ToSlice(s.seq, slice)
}

// Find returns the first element that satisfies the predicate.
func (s Sequence[E]) Find(predicate Predicate[E]) optional.Elem[E] {
	return Find(s.seq, predicate)
}

// FindLast returns the last element that satisfies the predicate.
func (s Sequence[E]) FindLast(predicate Predicate[E]) optional.Elem[E] {
	return FindLast(s.seq, predicate)
}

// FindAll returns all elements that satisfy the predicate.
func (s Sequence[E]) FindAll(predicate Predicate[E]) Sequence[E] {
	return AsSequence(FindAll(s.seq, predicate))
}

// Contains returns true if the element is in the sequence.
func (s Sequence[E]) Contains(elem E) bool {
	return Contains(s.seq, elem)
}

// NotContains returns true if the element is not in the sequence.
func (s Sequence[E]) NotContains(elem E) bool {
	return NotContains(s.seq, elem)
}

// ContainsAll returns true if all elements are in the sequence.
func (s Sequence[E]) ContainsAll(elements ...E) bool {
	return ContainsAll(s.seq, elements...)
}

// Exists returns true if there is at least one element that satisfies the predicate.
func (s Sequence[E]) Exists(predicate Predicate[E]) bool {
	return Exists(s.seq, predicate)
}

// Every returns true if all elements satisfy the predicate.
func (s Sequence[E]) Every(predicate Predicate[E]) bool {
	return Every(s.seq, predicate)
}

// None returns true if no element satisfies the predicate.
func (s Sequence[E]) None(predicate Predicate[E]) bool {
	return None(s.seq, predicate)
}

// IsNotEmpty returns true if the sequence is not empty.
func (s Sequence[E]) IsNotEmpty() bool {
	return IsNotEmpty(s.seq)
}

// IsEmpty returns true if the sequence is empty.
func (s Sequence[E]) IsEmpty() bool {
	return IsEmpty(s.seq)
}

// Partition splits the sequence into chunks of the given size.
func (s Sequence[E]) Partition(size int) iter.Seq[iter.Seq[E]] {
	return Partition(s.seq, size)
}

// Uniq returns a new sequence with only unique elements.
func (s Sequence[E]) Uniq() Sequence[E] {
	return AsSequence(Uniq(s.seq))
}

// Distinct returns a new sequence with only unique elements.
// SQL-like alias for Uniq
func (s Sequence[E]) Distinct() Sequence[E] {
	return AsSequence(Distinct(s.seq))
}

// Reverse returns a new sequence with elements in reverse order.
func (s Sequence[E]) Reverse() Sequence[E] {
	return AsSequence(Reverse(s.seq))
}

// Repeat repeats the sequence count times.
func (s Sequence[E]) Repeat(count int) Sequence[E] {
	return AsSequence(Cycle(s.seq, count))
}

// Fold applies a function against an accumulator and each element in the sequence (from left to right) to reduce it to a single value.
func (s Sequence[E]) Fold(accumulator func(agg E, item E) E) optional.Elem[E] {
	return Fold(s.seq, accumulator)
}

// FoldRight applies a function against an accumulator and each element in the sequence (from right to left) to reduce it to a single value.
func (s Sequence[E]) FoldRight(accumulator func(agg E, item E) E) optional.Elem[E] {
	return FoldRight(s.seq, accumulator)
}

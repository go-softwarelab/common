package seq2

import (
	"iter"

	"github.com/go-softwarelab/common/optional"
)

// Sequence is a wrapper around iter.Seq2 that provides fluent API for seq2 operations.
type Sequence[K comparable, V comparable] struct {
	seq iter.Seq2[K, V]
}

// AsSequence wraps an iter.Seq2 to provide a fluent API.
func AsSequence[K comparable, V comparable](seq iter.Seq2[K, V]) Sequence[K, V] {
	return Sequence[K, V]{seq}
}

// ConcatSequences concatenates multiple sequences into a single sequence.
func ConcatSequences[K comparable, V comparable](sequences ...Sequence[K, V]) Sequence[K, V] {
	var seq = Empty[K, V]()
	for _, s := range sequences {
		seq = UnionAll(seq, s.seq)
	}
	return AsSequence(seq)
}

// Union returns a new sequence that contains all distinct elements from both input sequences.
func (s Sequence[K, V]) Union(other Sequence[K, V]) Sequence[K, V] {
	return AsSequence(Union(s.seq, other.seq))
}

// UnionAll returns a new sequence that contains all elements from both input sequences.
func (s Sequence[K, V]) UnionAll(other Sequence[K, V]) Sequence[K, V] {
	return AsSequence(UnionAll(s.seq, other.seq))
}

// Append appends a key-value pair to the end of a sequence.
func (s Sequence[K, V]) Append(key K, value V) Sequence[K, V] {
	return AsSequence(Append(s.seq, key, value))
}

// Prepend prepends a key-value pair to the beginning of a sequence.
func (s Sequence[K, V]) Prepend(key K, value V) Sequence[K, V] {
	return AsSequence(Prepend(s.seq, key, value))
}

// Filter returns a new sequence with elements that satisfy the predicate.
func (s Sequence[K, V]) Filter(predicate Predicate[K, V]) Sequence[K, V] {
	return AsSequence(Filter(s.seq, predicate))
}

// Where returns a new sequence with elements that satisfy the predicate.
func (s Sequence[K, V]) Where(predicate Predicate[K, V]) Sequence[K, V] {
	return AsSequence(Where(s.seq, predicate))
}

// FilterByKey returns a new sequence with elements whose keys satisfy the predicate.
func (s Sequence[K, V]) FilterByKey(predicate KeyPredicate[K]) Sequence[K, V] {
	return AsSequence(FilterByKey(s.seq, predicate))
}

// FilterByValue returns a new sequence with elements whose values satisfy the predicate.
func (s Sequence[K, V]) FilterByValue(predicate ValuePredicate[V]) Sequence[K, V] {
	return AsSequence(FilterByValue(s.seq, predicate))
}

// Skip returns a new sequence that skips the first n elements.
func (s Sequence[K, V]) Skip(n int) Sequence[K, V] {
	return AsSequence(Skip(s.seq, n))
}

// Offset is an alias for Skip.
func (s Sequence[K, V]) Offset(n int) Sequence[K, V] {
	return s.Skip(n)
}

// Take returns a new sequence with only the first n elements.
func (s Sequence[K, V]) Take(n int) Sequence[K, V] {
	return AsSequence(Take(s.seq, n))
}

// Limit is an alias for Take.
func (s Sequence[K, V]) Limit(n int) Sequence[K, V] {
	return s.Take(n)
}

// Tap calls the consumer for each element without changing the sequence.
func (s Sequence[K, V]) Tap(consumer Consumer[K, V]) Sequence[K, V] {
	return AsSequence(Tap(s.seq, consumer))
}

// Each is an alias for Tap.
func (s Sequence[K, V]) Each(consumer Consumer[K, V]) Sequence[K, V] {
	return s.Tap(consumer)
}

// ForEach calls the consumer for each element.
func (s Sequence[K, V]) ForEach(consumer Consumer[K, V]) {
	ForEach(s.seq, consumer)
}

// Flush consumes all elements of the sequence.
func (s Sequence[K, V]) Flush() {
	Flush(s.seq)
}

// ToMap converts the sequence to a map.
func (s Sequence[K, V]) ToMap(m map[K]V) {
	ToMap(s.seq, m)
}

// Collect collects the elements into a map.
func (s Sequence[K, V]) Collect() map[K]V {
	return Collect(s.seq)
}

// Count returns the number of elements in the sequence.
func (s Sequence[K, V]) Count() int {
	return Count(s.seq)
}

// Keys returns a sequence of keys.
func (s Sequence[K, V]) Keys() iter.Seq[K] {
	return Keys(s.seq)
}

// Values returns a sequence of values.
func (s Sequence[K, V]) Values() iter.Seq[V] {
	return Values(s.seq)
}

// IsEmpty returns true if the sequence is empty.
func (s Sequence[K, V]) IsEmpty() bool {
	return IsEmpty(s.seq)
}

// IsNotEmpty returns true if the sequence is not empty.
func (s Sequence[K, V]) IsNotEmpty() bool {
	return IsNotEmpty(s.seq)
}

// Find returns the first element that satisfies the predicate.
func (s Sequence[K, V]) Find(predicate Predicate[K, V]) (optional.Elem[K], optional.Elem[V]) {
	return Find(s.seq, predicate)
}

// FindLast returns the last element that satisfies the predicate.
func (s Sequence[K, V]) FindLast(predicate Predicate[K, V]) (optional.Elem[K], optional.Elem[V]) {
	return FindLast(s.seq, predicate)
}

// FindAll returns all elements that satisfy the predicate.
func (s Sequence[K, V]) FindAll(predicate Predicate[K, V]) Sequence[K, V] {
	return AsSequence(FindAll(s.seq, predicate))
}

// FoldValues folds values with the first value as initial.
func (s Sequence[K, V]) FoldValues(accumulator func(agg V, key K, value V) V) optional.Elem[V] {
	return FoldValues(s.seq, accumulator)
}

// FoldValuesRight folds values from right to left.
func (s Sequence[K, V]) FoldValuesRight(accumulator func(agg V, key K, value V) V) optional.Elem[V] {
	return FoldValuesRight(s.seq, accumulator)
}

// SortComparingKeys sorts the sequence by keys using a custom comparator.
func (s Sequence[K, V]) SortComparingKeys(compare func(K, K) int) Sequence[K, V] {
	return AsSequence(SortComparingKeys(s.seq, compare))
}

// SortComparingValues sorts the sequence by values using a custom comparator.
func (s Sequence[K, V]) SortComparingValues(compare func(V, V) int) Sequence[K, V] {
	return AsSequence(SortComparingValues(s.seq, compare))
}

// Reverse returns a new sequence with elements in reverse order.
func (s Sequence[K, V]) Reverse() Sequence[K, V] {
	return AsSequence(Reverse(s.seq))
}

// Cycle repeats the sequence the specified number of times.
func (s Sequence[K, V]) Cycle(count int) Sequence[K, V] {
	return AsSequence(Cycle(s.seq, count))
}

// Repeat is an alias for Cycle.
func (s Sequence[K, V]) Repeat(count int) Sequence[K, V] {
	return s.Cycle(count)
}

// Contains returns true if the sequence contains the key.
func (s Sequence[K, V]) Contains(key K) bool {
	return Contains(s.seq, key)
}

// NotContains returns true if the sequence does not contain the key.
func (s Sequence[K, V]) NotContains(key K) bool {
	return NotContains(s.seq, key)
}

// ContainsValue returns true if the sequence contains the value.
func (s Sequence[K, V]) ContainsValue(value V) bool {
	return ContainsValue(s.seq, value)
}

// NotContainsValue returns true if the sequence does not contain the value.
func (s Sequence[K, V]) NotContainsValue(value V) bool {
	return NotContainsValue(s.seq, value)
}

// ContainsPair returns true if the sequence contains the key-value pair.
func (s Sequence[K, V]) ContainsPair(key K, value V) bool {
	return ContainsPair(s.seq, key, value)
}

// NotContainsPair returns true if the sequence does not contain the key-value pair.
func (s Sequence[K, V]) NotContainsPair(key K, value V) bool {
	return NotContainsPair(s.seq, key, value)
}

// Exists returns true if any element satisfies the predicate.
func (s Sequence[K, V]) Exists(predicate Predicate[K, V]) bool {
	return Exists(s.seq, predicate)
}

// Every returns true if all elements satisfy the predicate.
func (s Sequence[K, V]) Every(predicate Predicate[K, V]) bool {
	return Every(s.seq, predicate)
}

// None returns true if no element satisfies the predicate.
func (s Sequence[K, V]) None(predicate Predicate[K, V]) bool {
	return None(s.seq, predicate)
}

// Uniq returns a sequence with unique key-value pairs.
func (s Sequence[K, V]) Uniq() Sequence[K, V] {
	return AsSequence(Uniq(s.seq))
}

// UniqKeys returns a sequence with unique keys.
func (s Sequence[K, V]) UniqKeys() Sequence[K, V] {
	return AsSequence(UniqKeys(s.seq))
}

// UniqValues returns a sequence with unique values.
func (s Sequence[K, V]) UniqValues() Sequence[K, V] {
	return AsSequence(UniqValues(s.seq))
}

// Distinct is an alias for Uniq.
func (s Sequence[K, V]) Distinct() Sequence[K, V] {
	return s.Uniq()
}

// Get returns the value associated with the key.
func (s Sequence[K, V]) Get(key K) optional.Elem[V] {
	return Get(s.seq, key)
}

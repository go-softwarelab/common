package seq

import (
	"iter"

	"github.com/go-softwarelab/common/pkg/optional"
)

// Find returns the first element that satisfies the predicate.
func Find[E any](seq iter.Seq[E], predicate Predicate[E]) optional.Elem[E] {
	for v := range seq {
		if predicate(v) {
			return optional.Of(v)
		}
	}
	return optional.Empty[E]()
}

// FindLast returns the last element that satisfies the predicate.
func FindLast[E any](seq iter.Seq[E], predicate Predicate[E]) optional.Elem[E] {
	result := optional.Empty[E]()
	for v := range seq {
		if predicate(v) {
			result = optional.Of(v)
		}
	}
	return result
}

// FindAll returns all elements that satisfy the predicate.
func FindAll[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E] {
	return Filter(seq, predicate)
}

// Contains returns true if the element is in the sequence.
func Contains[E comparable](seq iter.Seq[E], elem E) bool {
	for v := range seq {
		if v == elem {
			return true
		}
	}
	return false
}

// NotContains returns true if the element is not in the sequence.
func NotContains[E comparable](seq iter.Seq[E], elem E) bool {
	return !Contains(seq, elem)
}

// ContainsAll returns true if all elements are in the sequence.
func ContainsAll[E comparable](seq iter.Seq[E], elements ...E) bool {
	for _, elem := range elements {
		if !Contains(seq, elem) {
			return false
		}
	}
	return true
}

// Exists returns true if there is at least one element that satisfies the predicate.
func Exists[E any](seq iter.Seq[E], predicate Predicate[E]) bool {
	for v := range seq {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Every returns true if all elements satisfy the predicate.
func Every[E any](seq iter.Seq[E], predicate Predicate[E]) bool {
	for v := range seq {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// None returns true if no element satisfies the predicate.
func None[E any](seq iter.Seq[E], predicate Predicate[E]) bool {
	return !Exists(seq, predicate)
}

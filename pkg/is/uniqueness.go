package is

import "iter"

// Unique returns true if all elements in the sequence are unique,
// and false if there are any duplicates.
// See: UniqueBy if it doesn't have comparable elements.
func Unique[E comparable, Collection ~[]E | iter.Seq[E]](collection Collection) bool {
	switch any(collection).(type) {
	case iter.Seq[E]:
		return UniqueSeq(any(collection).(iter.Seq[E]))
	default:
		// if a collection is not an iterator, then, according to generics, it must be a slice.
		return UniqueSlice(any(collection).([]E))
	}
}

// UniqueSeq checks if all elements in the given sequence are unique.
// It returns true if all elements are distinct, otherwise false.
// The sequence must provide elements of a comparable type.
// See: UniqueSeqBy if it doesn't have comparable elements.§
func UniqueSeq[E comparable](seq iter.Seq[E]) bool {
	seen := make(map[E]struct{})
	for v := range seq {
		if _, ok := seen[v]; ok {
			return false
		}
		seen[v] = struct{}{}
	}
	return true
}

// UniqueSlice checks if all elements in the provided slice are unique. Returns true if unique, otherwise false.
// See: UniqueSliceBy if it doesn't have comparable elements.
func UniqueSlice[E comparable](slice []E) bool {
	seen := make(map[E]struct{})
	for _, v := range slice {
		if _, ok := seen[v]; ok {
			return false
		}
		seen[v] = struct{}{}
	}
	return true
}

// UniqueBy returns true if all elements in the sequence are unique
// according to the given key function, false if any duplicate key exists.
func UniqueBy[E any, Collection ~[]E | iter.Seq[E], K comparable](collection Collection, key func(E) K) bool {
	switch any(collection).(type) {
	case iter.Seq[E]:
		return UniqueSeqBy(any(collection).(iter.Seq[E]), key)
	default:
		// if a collection is not an iterator, then, according to generics, it must be a slice.
		return UniqueSliceBy(any(collection).([]E), key)
	}
}

// UniqueSliceBy checks if all elements in the given slice are unique.
// It returns true if all elements are distinct, otherwise false.
func UniqueSliceBy[E any, K comparable](slice []E, key func(E) K) bool {
	seen := make(map[K]struct{})
	for _, v := range slice {
		k := key(v)
		if _, ok := seen[k]; ok {
			return false
		}
		seen[k] = struct{}{}
	}
	return true
}

// UniqueSeqBy checks if all elements in the given iter.Seq are unique.
// It returns true if all elements are distinct, otherwise false.
func UniqueSeqBy[E any, K comparable](seq iter.Seq[E], key func(E) K) bool {
	seen := make(map[K]struct{})
	for v := range seq {
		k := key(v)
		if _, ok := seen[k]; ok {
			return false
		}
		seen[k] = struct{}{}
	}
	return true
}

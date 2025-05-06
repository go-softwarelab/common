package seq2

import (
	"iter"

	"github.com/go-softwarelab/common/pkg/optional"
)

// FindAll returns all elements that satisfy the predicate.
func FindAll[K, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V] {
	return Filter(seq, predicate)
}

// findByKey returns the first element that satisfies the predicate.
func findByKey[K, V any](seq iter.Seq2[K, V], predicate KeyPredicate[K]) (optional.Value[K], optional.Value[V]) {
	for k, v := range seq {
		if predicate(k) {
			return optional.Of(k), optional.Of(v)
		}
	}
	return optional.Empty[K](), optional.Empty[V]()
}

// Get returns the element at the specified key.
func Get[K comparable, V any](seq iter.Seq2[K, V], key K) optional.Value[V] {
	_, v := findByKey(seq, func(k K) bool {
		return k == key
	})
	return v
}

// Contains returns true if the key is in the sequence.
func Contains[K comparable, V any](seq iter.Seq2[K, V], key K) bool {
	for k := range seq {
		if k == key {
			return true
		}
	}
	return false
}

// NotContains returns true if the key is not in the sequence.
func NotContains[K comparable, V any](seq iter.Seq2[K, V], key K) bool {
	return !Contains(seq, key)
}

// ContainsValue returns true if the value is in the sequence.
func ContainsValue[K any, V comparable](seq iter.Seq2[K, V], value V) bool {
	for _, v := range seq {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsAll returns true if all keys are in the sequence.
func ContainsAll[K comparable, V any](seq iter.Seq2[K, V], keys ...K) bool {
	for _, key := range keys {
		if !Contains(seq, key) {
			return false
		}
	}
	return true
}

// NotContainsValue returns true if the value is not in the sequence.
func NotContainsValue[K any, V comparable](seq iter.Seq2[K, V], value V) bool {
	return !ContainsValue(seq, value)
}

// ContainsAllValues returns true if all values are in the sequence.
func ContainsAllValues[K any, V comparable](seq iter.Seq2[K, V], values ...V) bool {
	for _, value := range values {
		if !ContainsValue(seq, value) {
			return false
		}
	}
	return true
}

// ContainsPair returns true if the key-value pair is in the sequence.
func ContainsPair[K comparable, V comparable](seq iter.Seq2[K, V], key K, value V) bool {
	for k, v := range seq {
		if k == key && v == value {
			return true
		}
	}
	return false
}

// NotContainsPair returns true if the key-value pair is not in the sequence.
func NotContainsPair[K comparable, V comparable](seq iter.Seq2[K, V], key K, value V) bool {
	return !ContainsPair(seq, key, value)
}

// Exists returns true if there is at least one element that satisfies the predicate.
func Exists[K, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) bool {
	for k, v := range seq {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

// Every returns true if all elements satisfy the predicate.
func Every[K, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) bool {
	for k, v := range seq {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

// None returns true if no element satisfies the predicate.
func None[K, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) bool {
	return !Exists(seq, predicate)
}

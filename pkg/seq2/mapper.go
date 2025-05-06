package seq2

import "iter"

// Mapper is a function that takes an element and returns a new element.
type Mapper[K, V, R any] = func(K, V) R

// DoubleMapper is a function that takes an element and returns a new sequence element.
type DoubleMapper[K, V, RK, RV any] = func(K, V) (RK, RV)

// KeyMapper is a function that takes Key a new Key.
type KeyMapper[K, R any] = func(K) R

// ValueMapper is a function that takes Value a new Value.
type ValueMapper[V, R any] = func(V) R

// Map applies a mapper function to each element of the sequence.
func Map[K, V, RK, RV any](seq iter.Seq2[K, V], mapper DoubleMapper[K, V, RK, RV]) iter.Seq2[RK, RV] {
	return func(yield func(RK, RV) bool) {
		for k, v := range seq {
			rk, rv := mapper(k, v)
			if !yield(rk, rv) {
				break
			}
		}
	}
}

// MapKeys applies a mapper function to each key of the sequence.
func MapKeys[K, V, RK any](seq iter.Seq2[K, V], mapper KeyMapper[K, RK]) iter.Seq2[RK, V] {
	return Map[K, V, RK, V](seq, func(k K, v V) (RK, V) {
		return mapper(k), v
	})
}

// MapValues applies a mapper function to each value of the sequence.
func MapValues[K, V, RV any](seq iter.Seq2[K, V], mapper ValueMapper[V, RV]) iter.Seq2[K, RV] {
	return Map[K, V, K, RV](seq, func(k K, v V) (K, RV) {
		return k, mapper(v)
	})
}

// MapTo applies a mapper function to each element of the sequence and returns a sequence of mapper results.
func MapTo[K, V, RV any](seq iter.Seq2[K, V], mapper Mapper[K, V, RV]) iter.Seq[RV] {
	return func(yield func(RV) bool) {
		for k, v := range seq {
			rv := mapper(k, v)
			if !yield(rv) {
				break
			}
		}
	}
}

// Cycle repeats the sequence indefinitely.
func Cycle[K, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// CycleTimes repeats the sequence count times.
func CycleTimes[K, V any](seq iter.Seq2[K, V], count int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := 0; i < count; i++ {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

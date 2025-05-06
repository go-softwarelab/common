package seq2

import "iter"

// Predicate is a function that takes an element and returns a boolean.
type Predicate[K any, V any] = Mapper[K, V, bool]

// KeyPredicate is a function that is used to filter by key.
type KeyPredicate[E any] = KeyMapper[E, bool]

// ValuePredicate is a function that is used to filter by value.
type ValuePredicate[E any] = ValueMapper[E, bool]

// Filter returns a new sequence that contains only the elements that satisfy the predicate.
func Filter[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if predicate(k, v) {
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// Where returns a new sequence that contains only the elements that satisfy the predicate.
// SQL-like alias for Filter
func Where[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V] {
	return Filter(seq, predicate)
}

// FilterByKey returns a new sequence that contains only the elements that satisfy the predicate.
func FilterByKey[K any, V any](seq iter.Seq2[K, V], predicate KeyPredicate[K]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if predicate(k) {
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// FilterByValue returns a new sequence that contains only the elements that satisfy the predicate.
func FilterByValue[K any, V any](seq iter.Seq2[K, V], predicate ValuePredicate[V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if predicate(v) {
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// Skip returns a new sequence that skips the first n elements of the given sequence.
func Skip[K any, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		i := 1
		for k, v := range seq {
			if i > n {
				if !yield(k, v) {
					break
				}
			}
			i++
		}
	}
}

// SkipWhile returns a new sequence that skips elements from the given sequence while the predicate is satisfied.
func SkipWhile[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		skip := true
		for k, v := range seq {
			if skip && predicate(k, v) {
				continue
			}
			skip = false
			if !yield(k, v) {
				break
			}
		}
	}
}

// SkipUntil returns a new sequence that skips elements from the given sequence until the predicate is satisfied.
func SkipUntil[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		skip := true
		for k, v := range seq {
			if skip && !predicate(k, v) {
				continue
			}
			skip = false
			if !yield(k, v) {
				break
			}
		}
	}
}

// Offset returns a new sequence that skips the first n elements of the given sequence.
// SQL-like alias for Skip.
func Offset[K any, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return Skip(seq, n)
}

// Take returns a new sequence that contains only the first n elements of the given sequence.
func Take[K any, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		i := 1
		for k, v := range seq {
			if i > n {
				break
			}
			if !yield(k, v) {
				break
			}
			i++
		}
	}
}

// TakeWhile returns a new sequence that takes elements from the given sequence while the predicate is satisfied.
func TakeWhile[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !predicate(k, v) || !yield(k, v) {
				break
			}
		}
	}
}

// TakeUntil returns a new sequence that takes elements from the given sequence until the predicate is satisfied.
func TakeUntil[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if predicate(k, v) || !yield(k, v) {
				break
			}
		}
	}
}

// Limit returns a new sequence that contains only the first n elements of the given sequence.
// SQL-like alias for Take.
func Limit[K any, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return Take(seq, n)
}

// Uniq returns a new sequence that contains only the unique elements of the given sequence.
// It compares both key and value.
// In case of pointers, pointers are compared, not the values they point to.
func Uniq[K comparable, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	type entry struct {
		k K
		v V
	}
	return func(yield func(K, V) bool) {
		seen := make(map[entry]struct{})
		for k, v := range seq {
			e := entry{k, v}
			if _, ok := seen[e]; !ok {
				seen[e] = struct{}{}
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// UniqKeys returns a new sequence that contains only the elements with unique keys from the given sequence.
func UniqKeys[K comparable, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		keys := make(map[K]struct{})
		for k, v := range seq {
			if _, ok := keys[k]; !ok {
				keys[k] = struct{}{}
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// UniqValues returns a new sequence that contains only the elements with unique values from the given sequence.
func UniqValues[K any, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		values := make(map[V]struct{})
		for k, v := range seq {
			if _, ok := values[v]; !ok {
				values[v] = struct{}{}
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// UniqBy returns a new sequence that contains only the unique elements of the given sequence based on result of the mapper.
func UniqBy[K any, V any, K2 comparable](seq iter.Seq2[K, V], mapper Mapper[K, V, K2]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seen := make(map[K2]struct{})
		for k, v := range seq {
			key := mapper(k, v)
			if _, ok := seen[key]; !ok {
				seen[key] = struct{}{}
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// UniqByKeys returns a new sequence that contains only the unique elements of the given sequence based on a result of key mapper.
func UniqByKeys[K any, V any, K2 comparable](seq iter.Seq2[K, V], mapper KeyMapper[K, K2]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seen := make(map[K2]struct{})
		for k, v := range seq {
			key := mapper(k)
			if _, ok := seen[key]; !ok {
				seen[key] = struct{}{}
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// UniqByValues returns a new sequence that contains only the unique elements of the given sequence based on a result of value mapper.
func UniqByValues[K any, V any, V2 comparable](seq iter.Seq2[K, V], mapper ValueMapper[V, V2]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seen := make(map[V2]struct{})
		for k, v := range seq {
			key := mapper(v)
			if _, ok := seen[key]; !ok {
				seen[key] = struct{}{}
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// Distinct returns a new sequence that contains only the unique elements of the given sequence.
// SQL-like alias for Uniq.
func Distinct[K comparable, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return Uniq(seq)
}

// DistinctKeys returns a new sequence that contains only the unique keys of the given sequence.
func DistinctKeys[K comparable, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return UniqKeys(seq)
}

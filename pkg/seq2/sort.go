package seq2

import (
	"cmp"
	"iter"
	"slices"

	"github.com/go-softwarelab/common/types"
)

// Sort sorts the elements of a sequence by key in ascending order.
func Sort[K types.Ordered, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return sortComparingPair(seq, func(a, b *pair[K, V]) int {
		return cmp.Compare(a.k, b.k)
	})
}

// SortBy sorts the elements of a sequence by result of the mapper.
func SortBy[K any, V any, R types.Ordered](seq iter.Seq2[K, V], mapper func(K, V) R) iter.Seq2[K, V] {
	withNewKey := MapPairs(seq, func(k K, v V) (R, *pair[K, V]) {
		return mapper(k, v), &pair[K, V]{k, v}
	})

	return func(yield func(K, V) bool) {
		sorted := sortComparingPair(withNewKey, func(a, b *pair[R, *pair[K, V]]) int {
			return cmp.Compare(a.k, b.k)
		})

		for _, p := range sorted {
			yield(p.k, p.v)
		}
	}
}

// SortComparingKeys sorts the elements of a sequence by key in ascending order.
func SortComparingKeys[K any, V any](seq iter.Seq2[K, V], cmp func(K, K) int) iter.Seq2[K, V] {
	return sortComparingPair(seq, func(a, b *pair[K, V]) int {
		return cmp(a.k, b.k)
	})
}

// SortComparingValues sorts the elements of a sequence by value in ascending order.
func SortComparingValues[K any, V any](seq iter.Seq2[K, V], cmp func(V, V) int) iter.Seq2[K, V] {
	return sortComparingPair(seq, func(a, b *pair[K, V]) int {
		return cmp(a.v, b.v)
	})
}

type pair[K, V any] struct {
	k K
	v V
}

func sortComparingPair[K any, V any](seq iter.Seq2[K, V], cmp func(*pair[K, V], *pair[K, V]) int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		pairs := Narrow(seq, func(k K, v V) *pair[K, V] {
			return &pair[K, V]{k, v}
		})
		sorted := slices.SortedFunc(pairs, cmp)
		for _, p := range sorted {
			if !yield(p.k, p.v) {
				break
			}
		}
	}
}

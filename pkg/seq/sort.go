package seq

import (
	"cmp"
	"iter"
	"slices"

	"github.com/go-softwarelab/common/types"
)

// Sort sorts the elements of a sequence in ascending order.
func Sort[E types.Ordered](seq iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		s := Collect(seq)
		slices.Sort(s)
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

// SortBy sorts the elements of a sequence in ascending order by the key returned by keyFn.
func SortBy[E any, K types.Ordered](seq iter.Seq[E], keyFn Mapper[E, K]) iter.Seq[E] {
	type pair struct {
		k K
		e E
	}

	return func(yield func(E) bool) {
		withKey := Map(seq, func(e E) pair {
			return pair{keyFn(e), e}
		})

		s := slices.SortedFunc(withKey, func(a, b pair) int {
			return cmp.Compare(a.k, b.k)
		})

		result := Map(FromSlice(s), func(p pair) E {
			return p.e
		})

		result(yield)
	}
}

// SortComparing sorts the elements of a sequence in ascending order using the cmp function.
func SortComparing[E any](seq iter.Seq[E], cmp func(a, b E) int) iter.Seq[E] {
	return func(yield func(E) bool) {
		s := slices.SortedFunc(seq, cmp)
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

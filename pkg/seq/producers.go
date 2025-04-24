package seq

import (
	"iter"
	"slices"
	"time"

	"github.com/go-softwarelab/common/pkg/types"
)

// Of creates a new sequence from the given elements.
func Of[E any](elems ...E) iter.Seq[E] {
	return FromSlice(elems)
}

// FromSlice creates a new sequence from the given slice.
func FromSlice[Slice ~[]E, E any](slice Slice) iter.Seq[E] {
	return slices.Values(slice)
}

// FromSliceReversed creates a new sequence from the given slice starting from last elements to first.
// It is more efficient then first creating a seq from slice and then reversing it.
func FromSliceReversed[Slice ~[]E, E any](slice Slice) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := len(slice) - 1; i >= 0; i-- {
			if !yield(slice[i]) {
				break
			}
		}
	}
}

// PointersFromSlice creates a new sequence of pointers for the given slice of value elements.
func PointersFromSlice[Slice ~[]E, E any](slice Slice) iter.Seq[*E] {
	return func(yield func(*E) bool) {
		for i := range slice {
			if !yield(&slice[i]) {
				break
			}
		}
	}
}

// Reverse creates a new sequence that iterates over the elements of the given sequence in reverse order.
func Reverse[E any](seq iter.Seq[E]) iter.Seq[E] {
	return FromSliceReversed(Collect(seq))
}

// Repeat returns a sequence that yields the same element `count` times.
func Repeat[E any, N types.Integer](elem E, count N) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := N(0); i < count; i++ {
			if !yield(elem) {
				break
			}
		}
	}
}

// RangeWithStep returns a sequence that yields integers from `start` to `end` with `step`.
func RangeWithStep[E types.Integer](start, end, step E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := start; i < end; i += step {
			if !yield(i) {
				break
			}
		}
	}
}

// Range returns a sequence that yields integers from `start` to `end`.
func Range[E types.Integer](start, end E) iter.Seq[E] {
	return RangeWithStep(start, end, 1)
}

// RangeTo returns a sequence that yields integers from 0 to `end`.
func RangeTo[E types.Integer](end E) iter.Seq[E] {
	return RangeWithStep(0, end, 1)
}

// Tick returns a sequence that yields the current time every duration.
func Tick(d time.Duration) iter.Seq[time.Time] {
	return func(yield func(time.Time) bool) {
		ticker := time.NewTicker(d)
		defer ticker.Stop()
		//nolint:gosimple
		for {
			select {
			case t := <-ticker.C:
				if !yield(t) {
					return
				}
			}
		}
	}
}

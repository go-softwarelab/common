package seq

import (
	"iter"
	"time"

	"github.com/go-softwarelab/common/types"
)

// Repeat returns a sequence that yields the same element `count` times.
func Repeat[E any](elem E, count int) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := 0; i < count; i++ {
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

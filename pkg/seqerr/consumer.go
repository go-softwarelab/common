package seqerr

import (
	"iter"
)

// ToSlice collects the elements of the given sequence into a slice.
func ToSlice[Slice ~[]E, E any](seq iter.Seq2[E, error], slice Slice) (Slice, error) {
	for v, err := range seq {
		if err != nil {
			return slice, err
		}
		slice = append(slice, v)
	}
	return slice, nil
}

// Collect collects the elements of the given sequence into a slice.
func Collect[E any](seq iter.Seq2[E, error]) ([]E, error) {
	return ToSlice(seq, []E(nil))
}

// Count returns the number of elements in the sequence.
func Count[E any](seq iter.Seq2[E, error]) (int, error) {
	i := 0
	for _, err := range seq {
		if err != nil {
			return 0, err
		}
		i++
	}
	return i, nil
}

// Tap returns a sequence that applies the given consumer to each element of the input sequence and pass it further.
// In case if consumer returns an error, the sequence stops and pass only the error from consumer further.
func Tap[E any, C Consumer[E]](seq iter.Seq2[E, error], consumer C) iter.Seq2[E, error] {
	consume := toConsumerWithError[E](consumer)
	return func(yield func(E, error) bool) {
		for v, err := range seq {
			if err != nil {
				yield(v, err)
				break
			}

			err = consume(v)
			if err != nil {
				yield(v, err)
				break
			}
			if !yield(v, nil) {
				break
			}
		}
	}
}

// Each returns a sequence that applies the given consumer to each element of the input sequence and pass it further.
// Each is an alias for Tap.
// Comparing to ForEach, this is a lazy function and doesn't consume the input sequence.
func Each[E any, C Consumer[E]](seq iter.Seq2[E, error], consumer C) iter.Seq2[E, error] {
	return Tap(seq, consumer)
}

// ForEach applies consumer to each element of the input sequence.
// Comparing to Each, this is not a lazy function and consumes the input sequence.
func ForEach[E any, C Consumer[E]](seq iter.Seq2[E, error], consumer C) error {
	consume := toConsumerWithError[E](consumer)

	for v, err := range seq {
		if err != nil {
			return err
		}

		err = consume(v)
		if err != nil {
			return err
		}
	}
	return nil
}

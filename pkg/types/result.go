package types

import "iter"

// PairLike represents any type that is pair-like, it is used for creating a Result instance.
type PairLike[L, R any] interface {
	GetLeft() L
	GetRight() R
}

// Result is a type representing a result that could be either a value or an error.
type Result[V any] struct {
	value V
	err   error
}

// SuccessResult creates a new Result instance with the provided value.
func SuccessResult[V any](value V) *Result[V] {
	return &Result[V]{value: value}
}

// FailureResult creates a new Result instance with the provided error.
func FailureResult[V any](err error) *Result[V] {
	return &Result[V]{err: err}
}

// ResultOf creates a new Result instance with the provided value and error.
func ResultOf[V any](value V, err error) *Result[V] {
	if err != nil {
		return FailureResult[V](err)
	}
	return SuccessResult(value)
}

// ResultFrom creates a Result instance from a function that returns a value and an error.
func ResultFrom[V any](provider func() (V, error)) *Result[V] {
	value, err := provider()
	if err != nil {
		return FailureResult[V](err)
	}
	return SuccessResult(value)
}

// ResultFromPair creates a Result instance from a PairLike argument.
func ResultFromPair[V any](pair PairLike[V, error]) *Result[V] {
	if pair.GetRight() != nil {
		return &Result[V]{err: pair.GetRight()}
	}
	return &Result[V]{value: pair.GetLeft()}
}

// IsError checks if the Result instance contains an error.
func (m *Result[V]) IsError() bool {
	return m.err != nil
}

// IsNotError checks if the Result instance does not contain an error.
func (m *Result[V]) IsNotError() bool {
	return m.err == nil
}

// Get returns the value and error from the Result instance.
func (m *Result[V]) Get() (V, error) {
	return m.value, m.err
}

// MustGetValue returns the value from the Result instance, panicking if there is an error.
func (m *Result[V]) MustGetValue() V {
	if m.IsError() {
		panic(m.err)
	}
	return m.value
}

// GetError returns the error from the Result instance.
// If there is no error, it returns nil.
func (m *Result[V]) GetError() error {
	return m.err
}

// MustGetError returns the error from the Result instance, panicking if there is no error.
func (m *Result[V]) MustGetError() error {
	if m.IsNotError() {
		panic("no error present")
	}
	return m.err
}

// OrElse returns the value if there is no error, otherwise it returns the provided default value.
func (m *Result[V]) OrElse(defaultValue V) V {
	if m.IsError() {
		return defaultValue
	}
	return m.value
}

// OrElseGet returns the value if there is no error, otherwise it returns the result of the provided function.
func (m *Result[V]) OrElseGet(defaultValue func() V) V {
	if m.IsError() {
		return defaultValue()
	}
	return m.value
}

// Or returns this Result if there is no error, otherwise it returns the provided alternative Result instance.
func (m *Result[V]) Or(alternative *Result[V]) *Result[V] {
	if m.IsError() {
		return alternative
	}
	return m
}

// Seq returns an iter.Seq with this Result.
//
// This is useful for reusing functions provided by package seq.
func (m *Result[V]) Seq() iter.Seq[Result[V]] {
	return func(yield func(Result[V]) bool) {
		yield(*m)
	}
}

// Seq2 returns an iter.Seq2 with value and error.
//
// This is useful for reusing functions provided by package seq2 or seqerr.
func (m *Result[V]) Seq2() iter.Seq2[V, error] {
	return func(yield func(V, error) bool) {
		yield(m.value, m.err)
	}
}

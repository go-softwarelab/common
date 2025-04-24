package types

import "iter"

// PairLike represents any type that is pair-like, it is used for creating a Maybe instance.
type PairLike[L, R any] interface {
	GetLeft() L
	GetRight() R
}

// Maybe is a type representing a result that could be either a value or an error.
type Maybe[V any] struct {
	value V
	err   error
}

// NewMaybe creates a new Maybe instance with the provided value and error.
func NewMaybe[V any](value V, err error) *Maybe[V] {
	if err != nil {
		return &Maybe[V]{err: err}
	}
	return &Maybe[V]{value: value}
}

// MaybeOf creates a Maybe instance from a PairLike argument.
func MaybeOf[V any](pair PairLike[V, error]) *Maybe[V] {
	if pair.GetRight() != nil {
		return &Maybe[V]{err: pair.GetRight()}
	}
	return &Maybe[V]{value: pair.GetLeft()}
}

// IsError checks if the Maybe instance contains an error.
func (m *Maybe[V]) IsError() bool {
	return m.err != nil
}

// IsNotError checks if the Maybe instance does not contain an error.
func (m *Maybe[V]) IsNotError() bool {
	return m.err == nil
}

// ShouldGetValue returns the value and error from the Maybe instance.
func (m *Maybe[V]) ShouldGetValue() (V, error) {
	return m.value, m.err
}

// MustGetValue returns the value from the Maybe instance, panicking if there is an error.
func (m *Maybe[V]) MustGetValue() V {
	if m.IsError() {
		panic(m.err)
	}
	return m.value
}

// OrElse returns the value if there is no error, otherwise it returns the provided default value.
func (m *Maybe[V]) OrElse(defaultValue V) V {
	if m.IsError() {
		return defaultValue
	}
	return m.value
}

// OrElseGet returns the value if there is no error, otherwise it returns the result of the provided function.
func (m *Maybe[V]) OrElseGet(defaultValue func() V) V {
	if m.IsError() {
		return defaultValue()
	}
	return m.value
}

// Or returns this Maybe if there is no error, otherwise it returns the provided alternative Maybe instance.
func (m *Maybe[V]) Or(alternative *Maybe[V]) *Maybe[V] {
	if m.IsError() {
		return alternative
	}
	return m
}

// Seq returns an iter.Seq with this Maybe.
//
// This is useful for reusing functions provided by package seq.
func (m *Maybe[V]) Seq() iter.Seq[Maybe[V]] {
	return func(yield func(Maybe[V]) bool) {
		yield(*m)
	}
}

// Seq2 returns an iter.Seq2 with value and error.
//
// This is useful for reusing functions provided by package seq2 or seqerr.
func (m *Maybe[V]) Seq2() iter.Seq2[V, error] {
	return func(yield func(V, error) bool) {
		yield(m.value, m.err)
	}
}

package optional

import (
	"errors"
	"fmt"
	"iter"
	"reflect"

	"github.com/go-softwarelab/common/pkg/is"
	"github.com/go-softwarelab/common/pkg/to"
)

const valueNotPresentErrorMessage = "value is not present"

// ValueNotPresent is the error returned or passed to iter.Seq2 when the value is not present.
var ValueNotPresent = errors.New(valueNotPresentErrorMessage)

// Value represents an optional value.
type Value[V any] struct {
	value *V
}

// Empty returns an empty optional value.
func Empty[V any]() Value[V] {
	return Value[V]{}
}

// None returns an empty optional value.
// alias: Empty
func None[V any]() Value[V] {
	return Empty[V]()
}

// Some returns an optional with the given value.
// It doesn't make any checks on value - it was caller decision to understand this value as present.
func Some[V any](v V) Value[V] {
	return Value[V]{value: &v}
}

// Of returns an optional with the given value.
// If the value is a pointer, and it's nil, it returns an empty optional.
// Otherwise, it returns non-empty optional with the given value.
func Of[E any](v E) Value[E] {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return Value[E]{}
	}

	return Value[E]{value: &v}
}

// OfPtr returns an optional with the value from pointer.
// If the pointer is nil, it returns an empty optional.
// Otherwise, it returns non-empty optional with the value pointed to by the pointer.
func OfPtr[E any](v *E) Value[E] {
	if v == nil {
		return Value[E]{}
	}

	return Value[E]{value: v}
}

// OfValue returns an optional for the given value.
// If value is zero value, it returns an empty optional.
// Otherwise, it returns non-empty optional with the given value.
//
// If zero value is valid existing value for you, for example when the value is int, then prefer Of() instead.
func OfValue[E comparable](v E) Value[E] {
	if is.Zero[E](v) {
		return Value[E]{}
	}

	return Value[E]{value: &v}
}

// Or returns this optional if present, otherwise returns the other optional.
func (o Value[V]) Or(other Value[V]) Value[V] {
	if o.IsPresent() {
		return o
	}

	return other
}

// ShouldGet returns the value if present, otherwise returns the error ValueNotPresent.
func (o Value[V]) ShouldGet() (V, error) {
	if o.IsEmpty() {
		return to.ZeroValue[V](), ValueNotPresent
	}

	return *o.value, nil
}

// MustGet returns the value if present, otherwise panics.
func (o Value[V]) MustGet() V {
	return o.MustGetf(valueNotPresentErrorMessage)
}

// MustGetf returns the value if present, otherwise panics with a custom message.
func (o Value[V]) MustGetf(msg string, args ...any) V {
	if o.IsEmpty() {
		panic(fmt.Sprintf(msg, args...))
	}

	return *o.value
}

// OrZeroValue returns the value if present, otherwise returns the zero value of the type.
func (o Value[V]) OrZeroValue() V {
	if o.IsEmpty() {
		return to.ZeroValue[V]()
	}

	return *o.value
}

// OrElse returns the value if present, otherwise returns the default value.
func (o Value[V]) OrElse(defaultValue V) V {
	if o.IsEmpty() {
		return defaultValue
	}

	return *o.value
}

// OrElseGet returns the value if present, otherwise returns the default value from the function.
func (o Value[V]) OrElseGet(defaultValue func() V) V {
	if o.IsEmpty() {
		return defaultValue()
	}

	return *o.value
}

// OrError returns the value if present, otherwise returns the error.
func (o Value[V]) OrError(err error) (V, error) {
	if o.IsEmpty() {
		return to.ZeroValue[V](), err
	}

	return *o.value, nil
}

// OrErrorGet returns the value if present, otherwise returns the error from the function.
func (o Value[V]) OrErrorGet(err func() error) (V, error) {
	if o.IsEmpty() {
		return to.ZeroValue[V](), err()
	}

	return *o.value, nil
}

// IfPresent executes the function if the value is present.
func (o Value[V]) IfPresent(fn func(V)) {
	if o.IsPresent() {
		fn(*o.value)
	}
}

// IfNotPresent executes the function if the value is not present.
func (o Value[V]) IfNotPresent(fn func()) {
	if o.IsEmpty() {
		fn()
	}
}

// IsEmpty returns true if the value is not present.
func (o Value[V]) IsEmpty() bool {
	return o.value == nil
}

// IsPresent returns true if the value is present.
func (o Value[V]) IsPresent() bool {
	return o.value != nil
}

// IsNotEmpty returns true if the value is present.
func (o Value[V]) IsNotEmpty() bool {
	return o.value != nil
}

// Seq returns the sequence with yelded value if present, otherwise returns an empty sequence.
func (o Value[V]) Seq() iter.Seq[V] {
	return func(yield func(V) bool) {
		if o.IsPresent() {
			yield(*o.value)
		}
	}
}

// Seq2 returns the iter.Seq2[V, error] with yelded value if present, otherwise yields an error.
// Useful with usage of seqerr package.
func (o Value[V]) Seq2() iter.Seq2[V, error] {
	return func(yield func(V, error) bool) {
		if o.IsPresent() {
			yield(*o.value, nil)
		} else {
			yield(to.ZeroValue[V](), ValueNotPresent)
		}
	}
}

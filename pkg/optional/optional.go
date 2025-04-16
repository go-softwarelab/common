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

// Elem represents an optional value.
type Elem[E any] struct {
	value *E
}

// Empty returns an empty optional value.
func Empty[E any]() Elem[E] {
	return Elem[E]{}
}

// Of returns an optional with the given value.
// If the value is a pointer, and it's nil, it returns an empty optional.
// Otherwise, it returns non-empty optional with the given value.
func Of[E any](v E) Elem[E] {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return Elem[E]{}
	}

	return Elem[E]{value: &v}
}

// OfPtr returns an optional with the value from pointer.
// If the pointer is nil, it returns an empty optional.
// Otherwise, it returns non-empty optional with the value pointed to by the pointer.
func OfPtr[E any](v *E) Elem[E] {
	if v == nil {
		return Elem[E]{}
	}

	return Elem[E]{value: v}
}

// OfValue returns an optional for the given value.
// If value is zero value, it returns an empty optional.
// Otherwise, it returns non-empty optional with the given value.
//
// If zero value is valid existing value for you, for example when the value is int, then prefer Of() instead.
func OfValue[E comparable](v E) Elem[E] {
	if is.Zero[E](v) {
		return Elem[E]{}
	}

	return Elem[E]{value: &v}
}

// Or returns this optional if present, otherwise returns the other optional.
func (o Elem[E]) Or(other Elem[E]) Elem[E] {
	if o.IsPresent() {
		return o
	}

	return other
}

// ShouldGet returns the value if present, otherwise returns the error ValueNotPresent.
func (o Elem[E]) ShouldGet() (E, error) {
	if o.IsEmpty() {
		return to.ZeroValue[E](), ValueNotPresent
	}

	return *o.value, nil
}

// MustGet returns the value if present, otherwise panics.
func (o Elem[E]) MustGet() E {
	return o.MustGetf(valueNotPresentErrorMessage)
}

// MustGetf returns the value if present, otherwise panics with a custom message.
func (o Elem[E]) MustGetf(msg string, args ...any) E {
	if o.IsEmpty() {
		panic(fmt.Sprintf(msg, args...))
	}

	return *o.value
}

// OrZeroValue returns the value if present, otherwise returns the zero value of the type.
func (o Elem[E]) OrZeroValue() E {
	if o.IsEmpty() {
		return to.ZeroValue[E]()
	}

	return *o.value
}

// OrElse returns the value if present, otherwise returns the default value.
func (o Elem[E]) OrElse(defaultValue E) E {
	if o.IsEmpty() {
		return defaultValue
	}

	return *o.value
}

// OrElseGet returns the value if present, otherwise returns the default value from the function.
func (o Elem[E]) OrElseGet(defaultValue func() E) E {
	if o.IsEmpty() {
		return defaultValue()
	}

	return *o.value
}

// OrError returns the value if present, otherwise returns the error.
func (o Elem[E]) OrError(err error) (E, error) {
	if o.IsEmpty() {
		return to.ZeroValue[E](), err
	}

	return *o.value, nil
}

// OrErrorGet returns the value if present, otherwise returns the error from the function.
func (o Elem[E]) OrErrorGet(err func() error) (E, error) {
	if o.IsEmpty() {
		return to.ZeroValue[E](), err()
	}

	return *o.value, nil
}

// IfPresent executes the function if the value is present.
func (o Elem[E]) IfPresent(fn func(E)) {
	if o.IsPresent() {
		fn(*o.value)
	}
}

// IfNotPresent executes the function if the value is not present.
func (o Elem[E]) IfNotPresent(fn func()) {
	if o.IsEmpty() {
		fn()
	}
}

// IsEmpty returns true if the value is not present.
func (o Elem[E]) IsEmpty() bool {
	return o.value == nil
}

// IsPresent returns true if the value is present.
func (o Elem[E]) IsPresent() bool {
	return o.value != nil
}

// IsNotEmpty returns true if the value is present.
func (o Elem[E]) IsNotEmpty() bool {
	return o.value != nil
}

// ToSeq returns the sequence with yelded value if present, otherwise returns an empty sequence.
func (o Elem[E]) ToSeq() iter.Seq[E] {
	return func(yield func(E) bool) {
		if o.IsPresent() {
			yield(*o.value)
		}
	}
}

// ToSeq2 returns the iter.Seq2[E, error] with yelded value if present, otherwise yields an error.
// Useful with usage of seqerr package.
func (o Elem[E]) ToSeq2() iter.Seq2[E, error] {
	return func(yield func(E, error) bool) {
		if o.IsPresent() {
			yield(*o.value, nil)
		} else {
			yield(to.ZeroValue[E](), ValueNotPresent)
		}
	}
}

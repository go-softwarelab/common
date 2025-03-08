package optional

import (
	"fmt"
	"reflect"
)

// Elem represents an optional value.
type Elem[E any] struct {
	value E
	exist bool
}

// Empty returns an empty optional value.
func Empty[E any]() Elem[E] {
	return Elem[E]{}
}

// Of returns an optional value with the given value.
func Of[E any](v E) Elem[E] {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return Elem[E]{}
	}

	return Elem[E]{value: v, exist: true}
}

// MustGet returns the value if present, otherwise panics.
func (o Elem[E]) MustGet() E {
	return o.MustGetf("value is not present")
}

// MustGetf returns the value if present, otherwise panics with a custom message.
func (o Elem[E]) MustGetf(msg string, args ...any) E {
	if o.IsEmpty() {
		panic(fmt.Sprintf(msg, args...))
	}

	return o.value
}

// IsEmpty returns true if the value is not present.
func (o Elem[E]) IsEmpty() bool {
	return !o.exist
}

// IsPresent returns true if the value is present.
func (o Elem[E]) IsPresent() bool {
	return o.exist
}

// IsNotEmpty returns true if the value is present.
func (o Elem[E]) IsNotEmpty() bool {
	return o.exist
}

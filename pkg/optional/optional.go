package optional

import (
	"fmt"
	"reflect"
)

type Elem[E any] struct {
	value E
	exist bool
}

func (o Elem[E]) MustGet() E {
	return o.MustGetf("value is not present")
}

func (o Elem[E]) MustGetf(msg string, args ...any) E {
	if o.IsEmpty() {
		panic(fmt.Sprintf(msg, args...))
	}

	return o.value
}

func (o Elem[E]) IsEmpty() bool {
	return !o.exist
}

func (o Elem[E]) IsPresent() bool {
	return o.exist
}

func (o Elem[E]) IsNotEmpty() bool {
	return o.exist
}

func Empty[E any]() Elem[E] {
	return Elem[E]{}
}

func Of[E any](v E) Elem[E] {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return Elem[E]{}
	}

	return Elem[E]{value: v, exist: true}
}

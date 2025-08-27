package is

import "reflect"

// Nil checks if a value is nil.
func Nil[T any](value T) bool {
	return isNil(value)
}

func isNil(value any) bool {
	defer func() { recover() }() //nolint:errcheck
	return value == nil || reflect.ValueOf(value).IsNil()
}

// NotNil checks if a value is not nil.
func NotNil[T any](value T) bool {
	return !isNil(value)
}

// Empty checks if a value is zero value.
func Empty[T comparable](value T) bool {
	var zero T
	return zero == value
}

// NotEmpty checks if a value is not zero value.
func NotEmpty[T comparable](value T) bool {
	return !Empty(value)
}

// Type checks if a value is of a specific type.
func Type[T any](value any) bool {
	_, ok := value.(T)
	return ok
}

// String checks if a value is a string.
var String = Type[string]

// Int checks if a value is an int.
var Int = Type[int]

// Bool checks if a value is a bool.
var Bool = Type[bool]

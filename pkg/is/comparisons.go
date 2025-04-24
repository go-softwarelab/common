package is

import (
	"cmp"

	"github.com/go-softwarelab/common/pkg/types"
)

// Zero checks if a value is zero.
func Zero[T comparable](value T) bool {
	return Empty(value)
}

// Equal checks if two values are equal.
func Equal[T comparable](a, b T) bool {
	return a == b
}

// NotEqual checks if two values are not equal.
func NotEqual[T comparable](a, b T) bool {
	return a != b
}

// EqualTo returns the function that checks if a value is equal to another.
func EqualTo[T comparable](expected T) func(T) bool {
	return func(v T) bool {
		return v == expected
	}
}

// Greater checks if a value is greater than another.
func Greater[T types.Ordered](a, b T) bool {
	return cmp.Compare(a, b) > 0
}

// GreaterThen returns the function that checks if a value is greater than another.
func GreaterThen[T types.Ordered](expected T) func(T) bool {
	return func(v T) bool {
		return cmp.Compare(v, expected) > 0
	}
}

// GreaterOrEqual checks if a value is greater than or equal to another.
func GreaterOrEqual[T types.Ordered](a, b T) bool {
	return cmp.Compare(a, b) >= 0
}

// GreaterOrEqualTo returns the function that checks if a value is greater than or equal to another.
func GreaterOrEqualTo[T types.Ordered](expected T) func(T) bool {
	return func(v T) bool {
		return cmp.Compare(v, expected) >= 0
	}
}

// Less checks if a value is less than another.
func Less[T types.Ordered](a, b T) bool {
	return cmp.Compare(a, b) < 0
}

// LessThen returns the function that checks if a value is less than another.
func LessThen[T types.Ordered](expected T) func(T) bool {
	return func(v T) bool {
		return cmp.Compare(v, expected) < 0
	}
}

// LessOrEqual checks if a value is less than or equal to another.
func LessOrEqual[T types.Ordered](a, b T) bool {
	return cmp.Compare(a, b) <= 0
}

// LessOrEqualTo returns the function that checks if a value is less than or equal to another.
func LessOrEqualTo[T types.Ordered](expected T) func(T) bool {
	return func(v T) bool {
		return cmp.Compare(v, expected) <= 0
	}
}

// Between checks if a value is between two others.
func Between[T types.Ordered](value, a, b T) bool {
	return GreaterOrEqual(value, a) && LessOrEqual(value, b)
}

// BetweenThe checks returns the function that checks if a value is between two others.
func BetweenThe[T types.Ordered](a, b T) func(T) bool {
	return func(value T) bool {
		return Between(value, a, b)
	}
}

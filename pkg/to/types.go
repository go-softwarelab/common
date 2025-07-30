package to

// Ptr returns a pointer copy of value.
func Ptr[T any](x T) *T {
	return &x
}

// Nil returns a nil pointer of type.
func Nil[T any]() *T {
	return nil
}

// NilOfType returns a nil pointer of type.
func NilOfType[T any](_ *T) *T {
	return nil
}

// EmptyValue returns the zero value of type.
func EmptyValue[T any]() T {
	var zero T
	return zero
}

// ZeroValue returns the zero value of type.
// alias: EmptyValue
func ZeroValue[T any]() T {
	var zero T
	return zero
}

// Value returns the pointer value or zero value.
func Value[T any](x *T) T {
	if x == nil {
		return EmptyValue[T]()
	}

	return *x
}

// ValueOr returns the pointer value or the fallback value.
func ValueOr[T any](x *T, fallback T) T {
	if x == nil {
		return fallback
	}

	return *x
}

// ValueOrGet returns the pointer value or the fallback value from result of fallback function call.
func ValueOrGet[T any](x *T, fallback func() T) T {
	if x == nil {
		return fallback()
	}

	return *x
}

// SliceOfPtr returns a slice of pointer copy of value.
func SliceOfPtr[T any](collection []T) []*T {
	ptrs := make([]*T, len(collection))
	for i, v := range collection {
		ptrs[i] = &v
	}
	return ptrs
}

// SliceOfValue returns a slice with the pointer values.
func SliceOfValue[T any](collection []*T) []T {
	values := make([]T, len(collection))
	for i, v := range collection {
		values[i] = Value(v)
	}
	return values
}

// Any casts the value to an any type.
func Any[T any](value T) any {
	return value
}

// SliceOfAny casts the slice to a slice of any type.
func SliceOfAny[T any](collection []T) []any {
	anyCollection := make([]any, len(collection))
	for i, v := range collection {
		anyCollection[i] = v
	}
	return anyCollection
}

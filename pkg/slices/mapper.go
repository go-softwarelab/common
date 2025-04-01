package slices

// Mapper is a function that maps a value of type T to a value of type R.
type Mapper[T any, R any] = func(T) R

// Map returns new slice where each element is a result of applying mapper to each element of the original slice and flattening the result.
func Map[E any, R any](collection []E, mapper Mapper[E, R]) []R {
	result := make([]R, len(collection))

	for i := range collection {
		result[i] = mapper(collection[i])
	}

	return result
}

// FlatMap returns new slice where each element is a result of applying mapper to each element of the original slice and flattening the result.
func FlatMap[E any, R any](collection []E, mapper Mapper[E, []R]) []R {
	result := make([]R, 0, len(collection))

	for i := range collection {
		result = append(result, mapper(collection[i])...)
	}

	return result
}

// Flatten flattens a slice of slices.
func Flatten[E any, Slice ~[]E](collection []Slice) Slice {
	totalLen := 0
	for i := range collection {
		totalLen += len(collection[i])
	}

	result := make(Slice, 0, totalLen)
	for i := range collection {
		result = append(result, collection[i]...)
	}

	return result
}

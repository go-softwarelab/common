package slices

// Mapper is a function that maps a value of type T to a value of type R.
type Mapper[T any, R any] = func(T) R

// MapperWithError is a function that maps a value of type T to a value of type R or returns an error.
type MapperWithError[T any, R any] = func(T) (R, error)

// Map returns new slice where each element is a result of applying mapper to each element of the original slice.
func Map[E any, R any](collection []E, mapper Mapper[E, R]) []R {
	result := make([]R, len(collection))

	for i := range collection {
		result[i] = mapper(collection[i])
	}

	return result
}

// MapOrError returns new slice where each element is a result of applying mapper to each element of the original slice.
// If any of the mappers return an error, the function returns an error.
func MapOrError[E any, R any](collection []E, mapper MapperWithError[E, R]) ([]R, error) {
	result := make([]R, len(collection))

	for i, v := range collection {
		mapped, err := mapper(v)
		if err != nil {
			return nil, err
		}
		result[i] = mapped
	}

	return result, nil
}

// FlatMap returns new slice where each element is a result of applying mapper to each element of the original slice and flattening the result.
func FlatMap[E any, R any](collection []E, mapper Mapper[E, []R]) []R {
	result := make([]R, 0, len(collection))

	for i := range collection {
		result = append(result, mapper(collection[i])...)
	}

	return result
}

// FlatMapOrError returns new slice where each element is a result of applying mapper to each element of the original slice and flattening the result.
// If any of the mappers return an error, the function returns an error.
func FlatMapOrError[E any, R any](collection []E, mapper MapperWithError[E, []R]) ([]R, error) {
	result := make([]R, 0, len(collection))

	for _, v := range collection {
		mapped, err := mapper(v)
		if err != nil {
			return nil, err
		}
		result = append(result, mapped...)
	}

	return result, nil
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

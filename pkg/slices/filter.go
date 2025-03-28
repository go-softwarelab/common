package slices

// Filter returns a new collection that contains only the elements that satisfy the predicate.
func Filter[E any, Slice ~[]E](collection Slice, predicate func(item E) bool) Slice {
	result := make(Slice, 0, len(collection))

	for _, e := range collection {
		if predicate(e) {
			result = append(result, e)
		}
	}

	return result
}

// Uniq returns a collection with only unique elements.
// The order of result values is determined by the order they occur in the array, only the first occurrence of each element is kept.
func Uniq[E comparable, Slice ~[]E](collection Slice) Slice {
	result := make(Slice, 0, len(collection))
	seen := make(map[E]struct{}, len(collection))

	for _, e := range collection {
		if _, ok := seen[e]; ok {
			continue
		}

		seen[e] = struct{}{}
		result = append(result, e)
	}

	return result
}

// UniqBy returns a collection with only unique elements based on a key.
// The order of result values is determined by the order they occur in the array, only the first occurrence of each element is kept.
func UniqBy[E any, R comparable, Slice ~[]E](collection Slice, mapper func(E) R) Slice {
	result := make(Slice, 0, len(collection))
	seen := make(map[R]struct{}, len(collection))

	for _, v := range collection {
		key := mapper(v)

		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, v)
	}

	return result
}

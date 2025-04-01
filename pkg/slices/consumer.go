package slices

// Consumer is a function that consumes an element of a sequence.
type Consumer[E any] = func(E)

// ForEach applies consumer to each element of the collection.
func ForEach[E any](collection []E, consumer Consumer[E]) {
	for _, e := range collection {
		consumer(e)
	}
}

// Count returns the number of elements in the collection.
func Count[E any](collection []E) int {
	return len(collection)
}

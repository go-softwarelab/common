package slices_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func ExampleMap() {
	collection := []int{1, 2, 3}

	squared := slices.Map(collection, func(v int) int {
		return v * v
	})
	fmt.Println(squared)

	// Output:
	// [1 4 9]
}

func ExampleFlatMap() {
	collection := []int{1, 2, 3}

	duplicated := slices.FlatMap(collection, func(v int) []int {
		return []int{v, v}
	})
	fmt.Println(duplicated)

	// Output:
	// [1 1 2 2 3 3]
}

func ExampleFlatten() {
	collection := [][]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}

	flattened := slices.Flatten(collection)
	fmt.Println(flattened)

	// Output:
	// [a b c d e f]
}

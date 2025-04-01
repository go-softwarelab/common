package slices_test

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/slices"
)

func ExampleFilter() {
	collection := []int{1, 2, 3, 4, 5}

	even := slices.Filter(collection, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(even)

	// Output:
	// [2 4]
}

func ExampleUniq() {
	collection := []string{"apple", "banana", "apple", "cherry", "banana"}

	unique := slices.Uniq(collection)
	fmt.Println(unique)

	// Output:
	// [apple banana cherry]
}

func ExampleUniqBy() {
	collection := []string{"Apple", "BANANA", "apple", "Cherry", "banana"}

	unique := slices.UniqBy(collection, func(s string) string {
		return strings.ToLower(s)
	})
	fmt.Println(unique)

	// Output:
	// [Apple BANANA Cherry]
}

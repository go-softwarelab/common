package slices_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func ExampleForEach() {
	collection := []int{1, 2, 3}

	slices.ForEach(collection, func(v int) {
		fmt.Println(v)
	})

	// Output:
	// 1
	// 2
	// 3
}

func ExampleCount() {
	collection := []string{"apple", "banana", "cherry"}

	count := slices.Count(collection)
	fmt.Println(count)

	emptyCollection := []int{}
	emptyCount := slices.Count(emptyCollection)
	fmt.Println(emptyCount)

	// Output:
	// 3
	// 0
}

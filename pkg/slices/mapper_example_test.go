package slices_test

import (
	"fmt"
	"strconv"

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

func ExampleMapOrError() {
	collection := []string{"1", "2", "3"}

	// Parse strings to ints, which could return an error
	parsed, err := slices.MapOrError(collection, func(v string) (int, error) {
		return strconv.Atoi(v)
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(parsed)

	// Now with an error case
	collection = []string{"1", "invalid", "3"}
	parsed, err = slices.MapOrError(collection, func(v string) (int, error) {
		return strconv.Atoi(v)
	})

	if err != nil {
		fmt.Println("Error occurred")
	}

	// Output:
	// [1 2 3]
	// Error occurred
}

func ExampleFlatMapOrError() {
	collection := []int{1, 2, 3}

	// Create pairs for each number, which could return an error
	result, err := slices.FlatMapOrError(collection, func(v int) ([]int, error) {
		if v <= 0 {
			return nil, fmt.Errorf("value must be positive: %d", v)
		}
		return []int{v, v * 10}, nil
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// Now with an error case
	collection = []int{1, 0, 3}
	result, err = slices.FlatMapOrError(collection, func(v int) ([]int, error) {
		if v <= 0 {
			return nil, fmt.Errorf("value must be positive: %d", v)
		}
		return []int{v, v * 10}, nil
	})

	if err != nil {
		fmt.Println("Error occurred")
	}

	// Output:
	// [1 10 2 20 3 30]
	// Error occurred
}

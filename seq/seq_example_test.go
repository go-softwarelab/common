package seq_test

import (
	"fmt"

	"github.com/go-softwarelab/common/seq"
)

func ExampleOf() {
	sequence := seq.Of(1, 2, 3)

	result := seq.Collect(sequence)

	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleFromSlice() {
	slice := []int{1, 2, 3}

	sequence := seq.FromSlice(slice)

	result := seq.Collect(sequence)

	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleToSlice() {
	sequence := seq.Of(1, 2, 3)

	slice := make([]int, 0, 3)
	result := seq.ToSlice(sequence, slice)

	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleCollect() {
	sequence := seq.Of(1, 2, 3)

	result := seq.Collect(sequence)

	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleReverse() {
	sequence := seq.Of(1, 2, 3)

	reversed := seq.Reverse(sequence)

	result := seq.Collect(reversed)
	fmt.Println(result)
	// Output:
	// [3 2 1]
}

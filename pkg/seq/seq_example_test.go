package seq_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
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

func ExampleFromSliceReversed() {
	slice := []int{1, 2, 3}

	sequence := seq.FromSliceReversed(slice)

	result := seq.Collect(sequence)

	fmt.Println(result)
	// Output:
	// [3 2 1]
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

func ExamplePointersFromSlice() {
	slice := []int{1, 2, 3}

	pointersSequence := seq.PointersFromSlice(slice)

	backToValues := seq.Map(pointersSequence, func(p *int) int {
		// NOTE: p is a pointer so no copy is made here
		return *p
	})

	result := seq.Collect(backToValues)
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

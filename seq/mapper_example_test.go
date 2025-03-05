package seq_test

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/seq"
)

func ExampleMap() {
	input := seq.Of(1, 2, 3)

	mapped := seq.Map(input, func(v int) string {
		return fmt.Sprintf("Number_%d", v)
	})

	result := seq.Collect(mapped)

	fmt.Println(result)
	// Output:
	// [Number_1 Number_2 Number_3]
}

func ExampleSelect() {
	input := seq.Of(1, 2, 3)

	mapped := seq.Select(input, func(v int) string {
		return fmt.Sprintf("Number_%d", v)
	})

	result := seq.Collect(mapped)

	fmt.Println(result)
	// Output:
	// [Number_1 Number_2 Number_3]
}

func ExampleFlatMap() {
	input := seq.Of(0, 3)

	flatMapped := seq.FlatMap(input, func(it int) iter.Seq[int] {
		return seq.Of[int](1+it, 2+it, 3+it)
	})

	result := seq.Collect(flatMapped)

	fmt.Println(result)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleFlatten() {
	input := seq.Of(seq.Of(1, 2), seq.Of(3, 4))

	flattened := seq.Flatten(input)

	result := seq.Collect(flattened)

	fmt.Println(result)
	// Output:
	// [1 2 3 4]
}

func ExampleCycle() {
	input := seq.Of(1, 2, 3)

	cycled := seq.Cycle(input, 2)

	result := seq.Collect(cycled)

	fmt.Println(result)
	// Output:
	// [1 2 3 1 2 3]
}

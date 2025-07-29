package seq_test

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
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

func ExampleMapOrErr() {
	input := seq.Of(1, 2, 3)

	// Example mapper function that returns an error for values > 2
	mapper := func(v int) (string, error) {
		if v > 2 {
			return "", fmt.Errorf("value %d is too large", v)
		}
		return fmt.Sprintf("Number_%d", v), nil
	}

	results := seq.MapOrErr(input, mapper)

	for val, err := range results {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Mapped value: %s\n", val)
		}
	}

	// Output:
	// Mapped value: Number_1
	// Mapped value: Number_2
	// Error: value 3 is too large
}

func ExampleMapTo() {
	input := seq.Of(1, 2, 3)

	mapped := seq.MapTo(input, func(v int) (int, string) {
		return v, fmt.Sprintf("Number_%d", v)
	})

	result := seq2.CollectToMap(mapped)

	fmt.Println(result)

	// Output:
	// map[1:Number_1 2:Number_2 3:Number_3]
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

func ExampleFlatMapOrErr() {
	input := seq.Of(1, 2, 3)

	mapper := func(v int) (iter.Seq[string], error) {
		if v > 2 {
			return nil, fmt.Errorf("value %d is too large", v)
		}
		return seq.Of(fmt.Sprintf("Number_%d_1", v), fmt.Sprintf("Number_%d_2", v)), nil
	}

	results := seq.FlatMapOrErr(input, mapper)

	for val, err := range results {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Mapped value: %s\n", val)
		}
	}

	// Output:
	// Mapped value: Number_1_1
	// Mapped value: Number_1_2
	// Mapped value: Number_2_1
	// Mapped value: Number_2_2
	// Error: value 3 is too large
}

func ExampleFlatMapSlices() {
	input := seq.Of(1, 2)

	mapped := seq.FlatMapSlices(input, func(v int) []string {
		return []string{
			fmt.Sprintf("Number_%d_1", v),
			fmt.Sprintf("Number_%d_2", v),
		}
	})

	result := seq.Collect(mapped)

	fmt.Println(result)
	// Output:
	// [Number_1_1 Number_1_2 Number_2_1 Number_2_2]
}

func ExampleFlatMapSlicesOrErr() {
	input := seq.Of(1, 2, 3)

	mapper := func(v int) ([]string, error) {
		if v > 2 {
			return nil, fmt.Errorf("value %d is too large", v)
		}
		return []string{
			fmt.Sprintf("Number_%d_1", v),
			fmt.Sprintf("Number_%d_2", v),
		}, nil
	}

	results := seq.FlatMapSlicesOrErr(input, mapper)

	for val, err := range results {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Mapped value: %s\n", val)
		}
	}

	// Output:
	// Mapped value: Number_1_1
	// Mapped value: Number_1_2
	// Mapped value: Number_2_1
	// Mapped value: Number_2_2
	// Error: value 3 is too large
}

func ExampleFlattenSlices() {
	// Create a sequence of slices
	sequence := seq.Of(1, 2, 3)

	seqOfSlices := seq.Map(sequence, func(n int) []int {
		return []int{n, n + 1}
	})

	// Flatten the sequence of slices
	flattened := seq.FlattenSlices(seqOfSlices)

	// Collect results
	result := seq.Collect(flattened)

	fmt.Println(result)

	// Output:
	// [1 2 2 3 3 4]
}

func ExampleCycle() {
	input := seq.Of(1, 2, 3)

	cycled := seq.Cycle(input)

	cycled = seq.Take(cycled, 9) // Limit to 9 elements for demonstration

	result := seq.Collect(cycled)

	fmt.Println(result)
	// Output:
	// [1 2 3 1 2 3 1 2 3]
}

func ExampleCycleTimes() {
	input := seq.Of(1, 2, 3)

	cycled := seq.CycleTimes(input, 2)

	cycled = seq.Take(cycled, 9) // Limit to 9 elements for demonstration difference between Cycle and CycleTimes

	result := seq.Collect(cycled)

	fmt.Println(result)
	// Output:
	// [1 2 3 1 2 3]
}

package seq_test

import (
	"fmt"

	"github.com/go-softwarelab/common/seq"
)

func ExampleFilter() {
	input := seq.Of(1, 2, 3, 4, 5)

	filtered := seq.Filter(input, func(v int) bool {
		return v%2 == 0
	})

	result := seq.Collect(filtered)

	fmt.Printf("%v\n", result)
	// Output:
	// [2 4]
}

func ExampleWhere() {
	input := seq.Of(1, 2, 3, 4, 5)

	filtered := seq.Where(input, func(v int) bool {
		return v%2 == 0
	})

	result := seq.Collect(filtered)

	fmt.Printf("%v\n", result)
	// Output:
	// [2 4]
}

func ExampleSkip() {
	input := seq.Of(1, 2, 3, 4, 5)

	skipped := seq.Skip(input, 2)

	result := seq.Collect(skipped)

	fmt.Printf("%v\n", result)
	// Output:
	// [3 4 5]
}

func ExampleOffset() {
	input := seq.Of(1, 2, 3, 4, 5)

	skipped := seq.Offset(input, 2)

	result := seq.Collect(skipped)

	fmt.Printf("%v\n", result)
	// Output:
	// [3 4 5]
}

func ExampleTake() {
	input := seq.Of(1, 2, 3, 4, 5)

	taken := seq.Take(input, 3)

	result := seq.Collect(taken)

	fmt.Printf("%v\n", result)
	// Output:
	// [1 2 3]
}

func ExampleTakeWhile() {
	input := seq.Of(1, 2, 3, 2, 1)

	taken := seq.TakeWhile(input, func(v int) bool {
		return v < 3
	})

	result := seq.Collect(taken)

	fmt.Printf("%v\n", result)
	// Output:
	// [3 2 1]
}

func ExampleLimit() {
	input := seq.Of(1, 2, 3, 4, 5)

	limited := seq.Limit(input, 2)

	result := seq.Collect(limited)

	fmt.Printf("%v\n", result)
	// Output:
	// [1 2]
}

func ExampleUniq() {
	input := seq.Of(1, 2, 2, 3, 3, 3)

	unique := seq.Uniq(input)

	result := seq.Collect(unique)

	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleDistinct() {
	input := seq.Of(1, 2, 2, 3, 3, 3)

	distinct := seq.Distinct(input)

	result := seq.Collect(distinct)

	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleUniqBy() {
	input := seq.Of("apple", "banana", "apricot", "blueberry")

	uniqueBy := seq.UniqBy(input, func(v string) string {
		return v[:1] // unique by first letter
	})

	result := seq.Collect(uniqueBy)

	fmt.Println(result)
	// Output:
	// [apple banana]
}

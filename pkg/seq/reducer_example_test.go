package seq_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func ExampleReduce() {
	input := seq.Of("a", "b", "c")

	concat := seq.Reduce(input, func(agg, item string) string {
		return agg + item
	}, "")

	fmt.Println(concat)
	// Output:
	// abc
}

func ExampleReduceRight() {
	input := seq.Of("a", "b", "c")

	concat := seq.ReduceRight(input, func(agg, item string) string {
		return agg + item
	}, "")

	fmt.Println(concat)
	// Output:
	// cba
}

func ExampleFold() {
	input := seq.Of("a", "b", "c")

	sum := seq.Fold(input, func(agg, item string) string {
		return agg + item
	})

	fmt.Println(sum.MustGet())
	// Output:
	// abc
}

func ExampleFoldRight() {
	input := seq.Of("a", "b", "c")

	sum := seq.FoldRight(input, func(agg, item string) string {
		return agg + item
	})

	fmt.Println(sum.MustGet())
	// Output:
	// cba
}

func ExampleMax() {
	input := seq.Of(2, 3, 1, 5, 4)

	maxVal := seq.Max(input)

	fmt.Println(maxVal.MustGet())
	// Output:
	// 5
}

func ExampleMin() {
	input := seq.Of(2, 3, 1, 5, 4)

	maxVal := seq.Min(input)

	fmt.Println(maxVal.MustGet())
	// Output:
	// 1
}

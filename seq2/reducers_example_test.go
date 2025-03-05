package seq2_test

import (
	"fmt"

	"github.com/go-softwarelab/common/seq2"
)

func ExampleReduce() {
	type Reduced struct {
		Key   string
		Value int
	}

	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// Reduce to calculate the sum of all values
	reduced := seq2.Reduce(input, func(agg Reduced, key string, value int) Reduced {
		return Reduced{
			Key:   agg.Key + key,
			Value: agg.Value + value,
		}
	}, Reduced{})

	fmt.Println(reduced)
	// Output:
	// {abc 6}
}

func ExampleReduceRight() {
	type Reduced struct {
		Key   string
		Value int
	}

	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// Reduce to calculate the sum of all values
	reduced := seq2.ReduceRight(input, func(agg Reduced, key string, value int) Reduced {
		return Reduced{
			Key:   agg.Key + key,
			Value: agg.Value + value,
		}
	}, Reduced{})

	fmt.Println(reduced)
	// Output:
	// {cba 6}
}

func ExampleFoldValues() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// FoldValues to calculate the product of all values
	product := seq2.FoldValues(input, func(agg int, key string, value int) int {
		return agg * value
	})

	fmt.Println(product.MustGet())
	// Output:
	// 6
}

func ExampleFoldValuesRight() {
	input := seq2.FromMap(map[string]int{"a": 3, "b": 5, "c": 10})
	input = seq2.Sort(input)

	// FoldValuesRight performs right-to-left division
	result := seq2.FoldValuesRight(input, func(agg int, key string, value int) int {
		return agg - value
	})

	fmt.Println(result.MustGet())
	// Output:
	// 2
}

func ExampleFoldKeys() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// FoldKeys concatenates all keys starting with the first key
	result := seq2.FoldKeys(input, func(agg string, key string, value int) string {
		return agg + key
	})

	fmt.Println(result.MustGet())
	// Output:
	// abc
}

func ExampleFoldKeysRight() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// FoldKeysRight concatenates all keys starting from the right
	result := seq2.FoldKeysRight(input, func(agg string, key string, value int) string {
		return agg + key
	})

	fmt.Println(result.MustGet())
	// Output:
	// cba
}

func ExampleMaxValue() {
	input := seq2.FromMap(map[string]int{"a": 5, "b": 10, "c": 3})
	input = seq2.Sort(input)

	// Find the maximum value
	maxV := seq2.MaxValue(input)

	fmt.Println(maxV.MustGet())
	// Output:
	// 10
}

func ExampleMinValue() {
	input := seq2.FromMap(map[string]int{"a": 5, "b": 10, "c": 3})
	input = seq2.Sort(input)

	// Find the minimum value
	minV := seq2.MinValue(input)

	fmt.Println(minV.MustGet())
	// Output:
	// 3
}

func ExampleMaxKey() {
	input := seq2.FromMap(map[string]int{"a": 5, "z": 10, "c": 3})
	input = seq2.Sort(input)

	// Find the maximum key (lexicographically)
	maxK := seq2.MaxKey(input)

	fmt.Println(maxK.MustGet())
	// Output:
	// z
}

func ExampleMinKey() {
	input := seq2.FromMap(map[string]int{"a": 5, "z": 10, "c": 3})
	input = seq2.Sort(input)

	// Find the minimum key (lexicographically)
	minK := seq2.MinKey(input)

	fmt.Println(minK.MustGet())
	// Output:
	// a
}

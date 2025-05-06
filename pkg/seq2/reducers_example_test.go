package seq2_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func ExampleReduce() {
	type Reduced struct {
		Key   string
		Value int
	}

	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

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
	input = seq2.SortByKeys(input)

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

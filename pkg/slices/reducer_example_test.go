package slices_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func ExampleReduce() {
	collection := []int{1, 2, 3, 4}

	sum := slices.Reduce(collection, func(agg int, item int) int {
		return agg + item
	}, 0)
	fmt.Println(sum)

	// Output:
	// 10
}

func ExampleReduceRight() {
	collection := []string{"a", "b", "c"}

	concatenated := slices.ReduceRight(collection, func(agg string, item string) string {
		return agg + item
	}, "")
	fmt.Println(concatenated)

	// Output:
	// cba
}

func ExampleReduce_customType() {
	type Product struct {
		Name  string
		Price int
	}

	collection := []Product{
		{Name: "Apple", Price: 1},
		{Name: "Banana", Price: 2},
		{Name: "Cherry", Price: 3},
	}

	totalPrice := slices.Reduce(collection, func(agg int, item Product) int {
		return agg + item.Price
	}, 0)
	fmt.Println(totalPrice)

	// Output:
	// 6
}

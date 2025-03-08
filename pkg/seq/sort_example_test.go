package seq_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func ExampleSort() {
	input := seq.Of(3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5)

	sorted := seq.Sort(input)

	result := seq.Collect(sorted)
	fmt.Println(result)
	// Output:
	// [1 1 2 3 3 4 5 5 5 6 9]
}

func ExampleSortBy() {
	type Person struct {
		Name string
		Age  int
	}
	input := seq.Of(
		Person{"Alice", 30},
		Person{"Bob", 25},
		Person{"Charlie", 35},
	)

	sorted := seq.SortBy(input, func(p Person) int {
		return p.Age
	})

	for p := range sorted {
		fmt.Printf("%s (%d)\n", p.Name, p.Age)
	}
	// Output:
	// Bob (25)
	// Alice (30)
	// Charlie (35)
}

func ExampleSortComparing() {
	type Person struct {
		Name string
		Age  int
	}
	input := seq.Of(
		Person{"Alice", 30},
		Person{"Bob", 25},
		Person{"Charlie", 35},
	)

	sorted := seq.SortComparing(input, func(a, b Person) int {
		return a.Age - b.Age
	})

	for p := range sorted {
		fmt.Printf("%s (%d)\n", p.Name, p.Age)
	}
	// Output:
	// Bob (25)
	// Alice (30)
	// Charlie (35)
}

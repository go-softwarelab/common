package seq_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func ExampleFind() {
	input := seq.Of(1, 2, 3, 4, 5)

	found := seq.Find(input, func(v int) bool { return v > 3 })

	fmt.Println(found.MustGet())
	// Output:
	// 4
}

func ExampleFindLast() {
	input := seq.Of(1, 2, 3, 4, 5)

	found := seq.FindLast(input, func(v int) bool { return v > 3 })

	fmt.Println(found.MustGet())
	// Output:
	// 5
}

func ExampleFindAll() {
	input := seq.Of(1, 2, 3, 4, 5)

	found := seq.FindAll(input, func(v int) bool { return v > 3 })

	result := seq.Collect(found)

	fmt.Println(result)
	// Output:
	// [4 5]
}

func ExampleContains() {
	input := seq.Of(1, 2, 3, 4, 5)

	contains := seq.Contains(input, 3)

	fmt.Println(contains)
	// Output:
	// true
}

func ExampleNotContains() {
	input := seq.Of(1, 2, 3, 4, 5)

	contains := seq.NotContains(input, 3)

	fmt.Println(contains)
	// Output:
	// false
}

func ExampleContainsAll() {
	input := seq.Of(1, 2, 3, 4, 5)

	containsAll := seq.ContainsAll(input, 2, 3)

	fmt.Println(containsAll)
	// Output:
	// true
}

func ExampleExists() {
	input := seq.Of(1, 2, 3, 4, 5)

	exists := seq.Exists(input, func(v int) bool { return v > 4 })

	fmt.Println(exists)
	// Output:
	// true
}

func ExampleEvery() {
	input := seq.Of(2, 4, 6, 8)

	every := seq.Every(input, func(v int) bool { return v%2 == 0 })

	fmt.Println(every)
	// Output:
	// true
}

func ExampleNone() {
	input := seq.Of(1, 2, 3, 4, 5)

	none := seq.None(input, func(v int) bool { return v > 5 })

	fmt.Println(none)
	// Output:
	// true
}

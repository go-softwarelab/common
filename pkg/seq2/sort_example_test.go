package seq2_test

import (
	"cmp"
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func ExampleSort() {
	input := seq2.Pair("a", 1)
	input = seq2.Append(input, "c", 3)
	input = seq2.Append(input, "b", 2)

	// Sort by keys (alphabetically)
	sorted := seq2.Sort(input)

	seq2.ForEach(sorted, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
	// Output:
	// a : 1
	// b : 2
	// c : 3
}

func ExampleSortBy() {
	input := seq2.Pair("a", 1)
	input = seq2.Append(input, "c", 3)
	input = seq2.Append(input, "b", 2)

	// Sort by values
	sorted := seq2.SortBy(input, func(k string, v int) int {
		return v
	})

	seq2.ForEach(sorted, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
	// Output:
	// a : 1
	// b : 2
	// c : 3
}

func ExampleSortComparingKeys() {
	input := seq2.Pair("a", 1)
	input = seq2.Append(input, "c", 3)
	input = seq2.Append(input, "b", 2)

	// Sort by keys in reverse order
	sorted := seq2.SortComparingKeys(input, func(a, b string) int {
		return -cmp.Compare(a, b)
	})

	seq2.ForEach(sorted, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
	// Output:
	// c : 3
	// b : 2
	// a : 1
}

func ExampleSortComparingValues() {
	// Unordered map with string keys
	input := seq2.Pair("a", 1)
	input = seq2.Append(input, "c", 3)
	input = seq2.Append(input, "b", 2)

	// Sort by values in descending order
	sorted := seq2.SortComparingValues(input, func(a, b int) int {
		return -cmp.Compare(a, b)
	})

	seq2.ForEach(sorted, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
	// Output:
	// c : 3
	// b : 2
	// a : 1
}

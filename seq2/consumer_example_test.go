package seq2_test

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/seq2"
)

func ExampleTap() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	// Ensure to have consistent output
	input = seq2.Sort(input)

	// Use Tap to print key-value pairs while passing them through
	tapped := seq2.Tap(input, func(k string, v int) {
		fmt.Printf("Processing: %s => %d\n", k, v)
	})

	seq2.Flush(tapped)

	// Output:
	// Processing: a => 1
	// Processing: b => 2
	// Processing: c => 3
}

func ExampleEach() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	tapped := seq2.Each(input, func(k string, v int) {
		fmt.Printf("Each: %s -> %d\n", k, v)
	})

	seq2.Flush(tapped)

	// Output:
	// Each: a -> 1
	// Each: b -> 2
	// Each: c -> 3
}

func ExampleForEach() {
	input := seq2.Of("a", "b", "c")

	// ForEach consumes the sequence and applies the function
	seq2.ForEach(input, func(k int, v string) {
		fmt.Printf("%d: %s\n", k, v)
	})

	// Output:
	// 0: a
	// 1: b
	// 2: c
}

func ExampleFlush() {
	// Create a sequence that has side effects when consumed
	sideEffects := iter.Seq2[string, int](func(yield func(string, int) bool) {
		fmt.Println("First element consumed")
		if !yield("a", 1) {
			return
		}
		fmt.Println("Second element consumed")
		if !yield("b", 2) {
			return
		}
		fmt.Println("Third element consumed")
		if !yield("c", 3) {
			return
		}
	})

	// Flush consumes all elements without doing anything with them
	seq2.Flush(sideEffects)

	// Output:
	// First element consumed
	// Second element consumed
	// Third element consumed
}

func ExampleToMap() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	result := make(map[string]int, 3)
	seq2.ToMap(input, result)

	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:3]
}

func ExampleCollect() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	result := seq2.Collect(input)

	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:3]
}

func ExampleCount() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// Count returns the number of elements in the sequence
	count := seq2.Count(input)

	fmt.Println(count)
	// Output:
	// 3
}

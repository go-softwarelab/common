package seq2_test

import (
	"fmt"

	"github.com/go-softwarelab/common/seq"
	"github.com/go-softwarelab/common/seq2"
)

func ExampleKeys() {
	input := seq2.Of("a", "b", "c")

	keys := seq2.Keys(input)

	seq.ForEach(keys, func(k int) {
		fmt.Print(k, " ")
	})
	// Output:
	// 0 1 2
}

func ExampleValues() {
	input := seq2.Of("a", "b", "c")

	keys := seq2.Values(input)

	seq.ForEach(keys, func(v string) {
		fmt.Print(v, " ")
	})
	// Output:
	// a b c
}

func ExampleIsEmpty() {
	emptySeq := seq2.Empty[any, any]()

	nonEmptySeq := seq2.Of("a")

	fmt.Printf("Empty sequence: %v\n", seq2.IsEmpty(emptySeq))
	fmt.Printf("Non-empty sequence: %v\n", seq2.IsEmpty(nonEmptySeq))
	// Output:
	// Empty sequence: true
	// Non-empty sequence: false
}

func ExampleIsNotEmpty() {
	emptySeq := seq2.Empty[any, any]()

	nonEmptySeq := seq2.Of("a")

	fmt.Printf("Empty sequence: %v\n", seq2.IsNotEmpty(emptySeq))
	fmt.Printf("Non-empty sequence: %v\n", seq2.IsNotEmpty(nonEmptySeq))
	// Output:
	// Empty sequence: false
	// Non-empty sequence: true
}

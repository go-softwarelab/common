package seq_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func ExampleConcat() {
	seq1 := seq.Of(1, 2, 3)
	seq2 := seq.Of(4, 5, 6)

	concatenated := seq.Concat(seq1, seq2, nil)

	result := seq.Collect(concatenated)

	fmt.Println(result)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleUnion() {
	seq1 := seq.Of(1, 2, 3)
	seq2 := seq.Of(3, 4, 5)

	union := seq.Union(seq1, seq2)

	result := seq.Collect(union)

	fmt.Println(result)
	// Output:
	// [1 2 3 4 5]
}

func ExampleUnionAll() {
	seq1 := seq.Of(1, 2, 3)
	seq2 := seq.Of(3, 4, 5)

	unionAll := seq.UnionAll(seq1, seq2)

	result := seq.Collect(unionAll)

	fmt.Println(result)
	// Output:
	// [1 2 3 3 4 5]
}

func ExampleAppend() {
	initial := seq.Of(1, 2, 3)

	appended := seq.Append(initial, 4, 5, 6)

	result := seq.Collect(appended)

	fmt.Println(result)
	// Output:
	// [1 2 3 4 5 6]
}

func ExamplePrepend() {
	initial := seq.Of(4, 5, 6)

	prepended := seq.Prepend(initial, 1, 2, 3)

	result := seq.Collect(prepended)

	fmt.Println(result)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleZip() {
	seq1 := seq.Of(1, 2, 3)
	seq2 := seq.Of("a", "b", "c")

	zipped := seq.Zip(seq1, seq2)

	for k, v := range zipped {
		fmt.Printf("%d: %s\n", k, v)
	}
	// Output:
	// 1: a
	// 2: b
	// 3: c
}

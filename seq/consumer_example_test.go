package seq_test

import (
	"fmt"

	"github.com/go-softwarelab/common/seq"
)

func ExampleTap() {
	sequence := seq.Tap(seq.Of(1, 2, 3), func(v int) {
		fmt.Println(v)
	})

	seq.Flush(sequence)

	// Output:
	// 1
	// 2
	// 3
}

func ExampleEach() {
	sequence := seq.Each(seq.Of(1, 2, 3), func(v int) {
		fmt.Println(v)
	})

	seq.Flush(sequence)

	// Output:
	// 1
	// 2
	// 3
}

func ExampleForEach() {
	seq.ForEach(seq.Of(1, 2, 3), func(v int) {
		fmt.Println(v)
	})

	// Output:
	// 1
	// 2
	// 3
}

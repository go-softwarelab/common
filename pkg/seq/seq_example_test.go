package seq_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func ExampleIsEmpty() {
	input := seq.Of(1, 2, 3)

	fmt.Println(seq.IsEmpty(input))
	// Output:
	// false
}

func ExampleIsNotEmpty() {
	input := seq.Of(1, 2, 3)

	isNotEmpty := seq.IsNotEmpty(input)

	fmt.Println(isNotEmpty)
	// Output:
	// true
}

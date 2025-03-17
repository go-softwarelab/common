package is_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func ExampleNot() {
	// Create a function to check if a value is positive
	isPositive := func(x int) bool {
		return x > 0
	}

	// Create the negation of that function
	isNotPositive := is.Not(isPositive)

	// Test with various values
	fmt.Printf("isNotPositive(5): %T(%v) - positive number\n", isNotPositive(5), isNotPositive(5))
	fmt.Printf("isNotPositive(0): %T(%v) - zero\n", isNotPositive(0), isNotPositive(0))
	fmt.Printf("isNotPositive(-3): %T(%v) - negative number\n", isNotPositive(-3), isNotPositive(-3))

	// Output:
	// isNotPositive(5): bool(false) - positive number
	// isNotPositive(0): bool(true) - zero
	// isNotPositive(-3): bool(true) - negative number
}

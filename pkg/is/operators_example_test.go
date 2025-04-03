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

func ExampleNotOrError() {
	// Create a function to check if a value is positive with error handling
	isPositiveWithError := func(x int) (bool, error) {
		return x > 0, nil
	}

	// Create the negation of that function
	isNotPositiveWithError := is.NotOrError(isPositiveWithError)

	// Test with various values
	result1, err1 := isNotPositiveWithError(5)
	fmt.Printf("isNotPositiveWithError(5): %T(%v), err: %v - positive number\n", result1, result1, err1)

	result2, err2 := isNotPositiveWithError(0)
	fmt.Printf("isNotPositiveWithError(0): %T(%v), err: %v - zero\n", result2, result2, err2)

	result3, err3 := isNotPositiveWithError(-3)
	fmt.Printf("isNotPositiveWithError(-3): %T(%v), err: %v - negative number\n", result3, result3, err3)

	// Output:
	// isNotPositiveWithError(5): bool(false), err: <nil> - positive number
	// isNotPositiveWithError(0): bool(true), err: <nil> - zero
	// isNotPositiveWithError(-3): bool(true), err: <nil> - negative number
}

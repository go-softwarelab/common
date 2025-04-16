package optional_test

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/optional"
)

func ExampleMap() {
	// Map a present value
	opt := optional.Of("hello")

	// Map to uppercase
	upperOpt := optional.Map(opt, strings.ToUpper)
	fmt.Println("Mapped value:", upperOpt.MustGet())

	// Map with more complex function
	lenOpt := optional.Map(opt, func(s string) int {
		return len(s)
	})
	fmt.Println("String length:", lenOpt.MustGet())

	// Output:
	// Mapped value: HELLO
	// String length: 5
}

func ExampleMap_empty() {
	// Map an empty optional
	emptyOpt := optional.Empty[string]()

	// Map function won't be called for empty optionals
	mapped := optional.Map(emptyOpt, strings.ToUpper)

	fmt.Println("Is mapped empty:", mapped.IsEmpty())
	fmt.Println("Is mapped present:", mapped.IsPresent())

	// Output:
	// Is mapped empty: true
	// Is mapped present: false
}

func ExampleMap_chain() {
	// Chaining multiple Map operations
	opt := optional.Of(42)

	msg := optional.Map(opt, func(n int) string {
		return fmt.Sprintf("Number: %d", n)
	})

	result := optional.Map(msg, func(s string) []byte {
		return []byte(s)
	})

	// Check if result is present
	if result.IsPresent() {
		fmt.Printf("Result type: %T\n", result.MustGet())
		fmt.Printf("Result length: %d\n", len(result.MustGet()))
	}

	// Output:
	// Result type: []uint8
	// Result length: 10
}

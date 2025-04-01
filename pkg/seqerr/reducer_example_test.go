package seqerr_test

import (
	"errors"
	"fmt"
	"iter"
	"strings"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func ExampleReduce() {
	// Create a sequence with numbers 1-5
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 5; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Sum all numbers
	sum, err := seqerr.Reduce(sequence, func(acc int, item int) int {
		return acc + item
	}, 0)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Sum:", sum)

	// Output:
	// Sum: 15
}

func ExampleReduce_withError() {
	// Create a sequence with an error
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 5; i++ {
			var err error
			if i == 3 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Sum all numbers
	sum, err := seqerr.Reduce(sequence, func(acc int, item int) int {
		return acc + item
	}, 0)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Sum:", sum)

	// Output:
	// Error: source error
	// Sum: 0
}

func ExampleReduceRight() {
	// Create a sequence with words
	sequence := iter.Seq2[string, error](func(yield func(string, error) bool) {
		words := []string{"hello", "beautiful", "world"}
		for _, word := range words {
			if !yield(word, nil) {
				break
			}
		}
	})

	// Join words right-to-left
	joined, err := seqerr.ReduceRight(sequence, func(acc string, item string) string {
		if acc == "" {
			return item
		}
		return item + " " + acc
	}, "")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(joined)

	// Output:
	// hello beautiful world
}

func ExampleFold() {
	// Create a sequence with numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 5; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Calculate the product of all numbers
	result, err := seqerr.Fold(sequence, func(acc int, item int) int {
		return acc * item
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if result.IsPresent() {
		fmt.Println("Product:", result.MustGet())
	} else {
		fmt.Println("Empty sequence")
	}

	// Output:
	// Product: 120
}

func ExampleFold_emptySequence() {
	// Create an empty sequence
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		// Do nothing
	})

	// Try to fold the empty sequence
	result, err := seqerr.Fold(sequence, func(acc int, item int) int {
		return acc + item
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if result.IsPresent() {
		fmt.Println("Sum:", result.MustGet())
	} else {
		fmt.Println("Empty sequence")
	}

	// Output:
	// Empty sequence
}

func ExampleFoldRight() {
	// Create a sequence with strings
	sequence := iter.Seq2[string, error](func(yield func(string, error) bool) {
		words := []string{"hello", "world"}
		for _, word := range words {
			if !yield(word, nil) {
				break
			}
		}
	})

	// Concatenate strings right-to-left
	result, err := seqerr.FoldRight(sequence, func(acc string, item string) string {
		return strings.ToUpper(acc) + "-" + item
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if result.IsPresent() {
		fmt.Println(result.MustGet())
	} else {
		fmt.Println("Empty sequence")
	}

	// Output:
	// WORLD-hello
}

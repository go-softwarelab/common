package seqerr_test

import (
	"errors"
	"fmt"
	"iter"
	"strconv"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func ExampleFilter_predicateWithoutError() {
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := range 5 {
			yield(i, nil)
		}
	})

	// Filter even numbers
	evenNumbers := seqerr.Filter(sequence, func(n int) bool {
		return n%2 == 0
	})

	// Print results
	for n, err := range evenNumbers {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 0
	// 2
	// 4
}

func ExampleFilter_predicateWithError() {
	sequence := iter.Seq2[string, error](func(yield func(string, error) bool) {
		for i := range 5 {
			if !yield(fmt.Sprintf("%d", i), nil) {
				break
			}
		}
	})

	// Filter strings that are even numbers when converted to int
	filtered := seqerr.Filter(sequence, func(s string) (bool, error) {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false, err
		}
		return i%2 == 0, nil
	})

	// Print results
	for s, err := range filtered {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(s)
	}

	// Output:
	// 0
	// 2
	// 4
}

func ExampleFilter_validator() {
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := range 5 {
			if !yield(i+1, nil) {
				break
			}
		}
	})

	// Filter using a validator (non-zero values)
	nonZero := seqerr.Filter(sequence, func(n int) error {
		if n%2 == 0 {
			return errors.New("even value not allowed")
		}
		return nil
	})

	// Collect results
	for n, err := range nonZero {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 1
	// Error: even value not allowed
}

func ExampleFilter_sourceError() {
	// Create a sequence that produces an error
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := range 5 {
			var err error
			if i == 2 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Filter even numbers
	filtered := seqerr.Filter(sequence, func(n int) bool {
		return n%2 == 0
	})

	// Collect results
	for n, err := range filtered {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 0
	// Error: source error
}

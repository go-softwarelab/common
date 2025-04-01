package seqerr_test

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func ExampleToSlice() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Collect elements into a slice
	slice, err := seqerr.ToSlice(sequence, make([]int, 0))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(slice)

	// Output:
	// [1 2 3]
}

func ExampleToSlice_withError() {
	// Create a sequence with an error
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			var err error
			if i == 2 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Collect elements into a slice
	slice, err := seqerr.ToSlice(sequence, make([]int, 0))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(slice)

	// Output:
	// Error: source error
	// [1]
}

func ExampleCollect() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Collect elements into a new slice
	result, err := seqerr.Collect(sequence)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(result)

	// Output:
	// [1 2 3]
}

func ExampleCollect_withError() {
	// Create a sequence with an error
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			var err error
			if i == 2 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Collect elements into a new slice
	result, err := seqerr.Collect(sequence)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(result)

	// Output:
	// Error: source error
	// [1]
}

func ExampleCount() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Count elements in the sequence
	count, err := seqerr.Count(sequence)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(count)

	// Output:
	// 3
}

func ExampleCount_withError() {
	// Create a sequence with an error
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			var err error
			if i == 2 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Count elements in the sequence
	count, err := seqerr.Count(sequence)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(count)

	// Output:
	// Error: source error
	// 0
}

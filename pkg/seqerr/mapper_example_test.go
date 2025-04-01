package seqerr_test

import (
	"errors"
	"fmt"
	"iter"
	"strconv"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seqerr"
)

func ExampleMap() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			yield(i, nil)
		}
	})

	// Map numbers to their squares
	squared := seqerr.Map(sequence, func(n int) int {
		return n * n
	})

	// Collect results
	for n, err := range squared {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 1
	// 4
	// 9
}

func ExampleMap_sourceError() {
	// Create a sequence with error
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

	// Map numbers to their squares
	squared := seqerr.Map(sequence, func(n int) int {
		return n * n
	})

	// Collect results
	for n, err := range squared {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 1
	// Error: source error
}

func ExampleMapOrErr() {
	// Create a sequence of strings
	sequence := iter.Seq2[string, error](func(yield func(string, error) bool) {
		for i := 1; i <= 3; i++ {
			var value string
			if i == 2 {
				value = "two"
			} else {
				value = strconv.Itoa(i)
			}
			if !yield(value, nil) {
				break
			}
		}
	})

	// Map strings to integers
	numbers := seqerr.MapOrErr(sequence, func(s string) (int, error) {
		return strconv.Atoi(s)
	})

	// Collect results
	for n, err := range numbers {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 1
	// Error: strconv.Atoi: parsing "two": invalid syntax
}

func ExampleFlatMap() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// FlatMap to create duplicates of each number
	duplicated := seqerr.FlatMap(sequence, func(n int) iter.Seq[int] {
		return seq.Repeat(n, 2)
	})

	// Collect results
	for n, err := range duplicated {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 1
	// 1
	// 2
	// 2
	// 3
	// 3
}

func ExampleFlatMapOrErr() {
	// Create a sequence of strings
	sequence := iter.Seq2[string, error](func(yield func(string, error) bool) {
		for i := 1; i <= 3; i++ {
			var value string
			if i == 2 {
				value = "two"
			} else {
				value = strconv.Itoa(i)
			}
			if !yield(value, nil) {
				break
			}
		}
	})

	// Duplicate integers or return an error if cannot be parsed
	numbers := seqerr.FlatMapOrErr(sequence, func(s string) (iter.Seq[int], error) {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		return seq.Repeat(i, 2), nil
	})

	// Collect results
	for n, err := range numbers {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 1
	// 1
	// Error: strconv.Atoi: parsing "two": invalid syntax
}

func ExampleFlatten() {
	// Create a sequence of sequences
	sequence := seq.RangeTo(3)

	seqOfSeq := seqerr.MapSeq(sequence, func(n int) (iter.Seq[int], error) {
		return seq.Repeat(n, 2), nil
	})

	// Flatten the sequence of sequences
	flattened := seqerr.Flatten(seqOfSeq)

	// Collect results
	for n, err := range flattened {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 0
	// 0
	// 1
	// 1
	// 2
	// 2
}

func ExampleFlattenSlices() {
	// Create a sequence of slices
	sequence := iter.Seq2[[]int, error](func(yield func([]int, error) bool) {
		slices := [][]int{{1, 2}, {3, 4, 5}, {6}}
		for _, slice := range slices {
			if !yield(slice, nil) {
				break
			}
		}
	})

	// Flatten the sequence of slices
	flattened := seqerr.FlattenSlices(sequence)

	// Collect results
	for n, err := range flattened {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func ExampleFlattenSlices_withError() {
	// Create a sequence of slices with an error
	sequence := iter.Seq2[[]int, error](func(yield func([]int, error) bool) {
		if !yield([]int{1, 2}, nil) {
			return
		}
		if !yield(nil, errors.New("slice error")) {
			return
		}
		// This won't be processed due to the error
		if !yield([]int{3, 4}, nil) {
			return
		}
	})

	// Flatten the sequence of slices
	flattened := seqerr.FlattenSlices(sequence)

	// Collect results
	for n, err := range flattened {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 1
	// 2
	// Error: slice error
}

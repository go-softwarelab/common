package seqerr_test

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seqerr"
)

func ExampleConcat() {
	// Create first sequence with numbers 1-2
	first := seqerr.FromSeq(seq.Range(1, 3))

	// Create second sequence with numbers 3-4
	second := seqerr.FromSeq(seq.Range(3, 5))

	// Concatenate the two sequences
	combined := seqerr.Concat(first, second)

	// Collect results
	for n, err := range combined {
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
}

func ExampleConcat_withError() {
	// Create first sequence with an error
	first := iter.Seq2[int, error](func(yield func(int, error) bool) {
		if !yield(1, nil) {
			return
		}
		if !yield(2, errors.New("first sequence error")) {
			return
		}
		// This won't be processed due to the error
		if !yield(3, nil) {
			return
		}
	})

	// Create second sequence
	second := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 4; i <= 5; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Concatenate the two sequences
	combined := seqerr.Concat(first, second)

	// Collect results
	for n, err := range combined {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 1
	// Error: first sequence error
}

func ExampleAppend() {
	// Create a sequence with numbers 1-3
	sequence := seq.Range(1, 4)

	// Append number 4 to the sequence
	appended := seqerr.Append(seqerr.FromSeq(sequence), 4)

	// Collect results
	for n, err := range appended {
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
}

func ExamplePrepend() {
	// Create a sequence with numbers 2-4
	sequence := seq.Range(2, 5)

	// Prepend number 1 to the sequence
	prepended := seqerr.Prepend(seqerr.FromSeq(sequence), 1)

	// Collect results
	for n, err := range prepended {
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
}

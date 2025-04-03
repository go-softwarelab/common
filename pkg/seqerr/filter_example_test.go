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

func ExampleTake() {
	// Create a sequence of numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := range 10 {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Take only the first 3 elements
	taken := seqerr.Take(sequence, 3)

	// Print results
	for n, err := range taken {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 0
	// 1
	// 2
}

func ExampleTakeWhile() {
	// Create a sequence of numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Take elements while they are less than 5
	taken := seqerr.TakeWhile(sequence, func(n int) bool {
		return n < 5
	})

	// Print results
	for n, err := range taken {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleTakeUntil() {
	// Create a sequence of numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Take elements until we find one that's equal to 5
	taken := seqerr.TakeUntil(sequence, func(n int) bool {
		return n == 5
	})

	// Print results
	for n, err := range taken {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleTakeWhileTrue() {
	// Create a sequence of numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Create a condition that will be true only for the first 4 elements
	count := 0
	condition := func() bool {
		count++
		return count <= 4
	}

	// Take elements while the condition remains true
	taken := seqerr.TakeWhileTrue(sequence, condition)

	// Print results
	for n, err := range taken {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
}

func ExampleTakeUntilTrue() {
	// Create a sequence of numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Create a condition that will become true after 5 elements
	count := 0
	condition := func() bool {
		count++
		return count > 5
	}

	// Take elements until the condition becomes true
	taken := seqerr.TakeUntilTrue(sequence, condition)

	// Print results
	for n, err := range taken {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

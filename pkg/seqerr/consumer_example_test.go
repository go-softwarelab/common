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

func ExampleTap() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Use Tap to print each element while passing it through
	tapped := seqerr.Tap(sequence, func(n int) {
		fmt.Printf("Processing: %d\n", n)
	})

	// Collect the elements after tapping
	result, err := seqerr.Collect(tapped)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Result:", result)

	// Output:
	// Processing: 1
	// Processing: 2
	// Processing: 3
	// Result: [1 2 3]
}

func ExampleTap_withError() {
	// Create a sequence
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Use Tap with a consumer that returns an error
	tapped := seqerr.Tap(sequence, func(n int) error {
		if n == 2 {
			return errors.New("consumer error")
		}
		fmt.Printf("Processing: %d\n", n)
		return nil
	})

	// Collect the elements after tapping
	result, err := seqerr.Collect(tapped)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Result:", result)

	// Output:
	// Processing: 1
	// Error: consumer error
	// Result: [1]
}

func ExampleEach() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Use Each (alias for Tap) to print each element while passing it through
	each := seqerr.Each(sequence, func(n int) {
		fmt.Printf("Element: %d\n", n)
	})

	// Collect the elements after processing
	result, err := seqerr.Collect(each)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Result:", result)

	// Output:
	// Element: 1
	// Element: 2
	// Element: 3
	// Result: [1 2 3]
}

func ExampleForEach() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Use ForEach to process all elements
	sum := 0
	err := seqerr.ForEach(sequence, func(n int) {
		fmt.Printf("Adding: %d\n", n)
		sum += n
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Sum:", sum)

	// Output:
	// Adding: 1
	// Adding: 2
	// Adding: 3
	// Sum: 6
}

func ExampleForEach_withError() {
	// Create a sequence with numbers 1-3
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

	// Use ForEach to process elements
	sum := 0
	err := seqerr.ForEach(sequence, func(n int) {
		fmt.Printf("Adding: %d\n", n)
		sum += n
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Sum:", sum)

	// Output:
	// Adding: 1
	// Error: source error
	// Sum: 1
}

func ExampleForEach_withConsumerError() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Use ForEach with a consumer that returns an error
	sum := 0
	err := seqerr.ForEach(sequence, func(n int) error {
		if n == 2 {
			return errors.New("consumer error")
		}
		fmt.Printf("Adding: %d\n", n)
		sum += n
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Sum:", sum)

	// Output:
	// Adding: 1
	// Error: consumer error
	// Sum: 1
}

package seqerr_test

import (
	"fmt"
	"strconv"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func ExampleProduce() {
	sequence := seqerr.Produce(func(i int) ([]string, int, error) {
		if i == 2 {
			return []string{}, i + 1, nil
		}

		num := strconv.Itoa(i)
		result := []string{"a" + num, "b" + num, "c" + num}
		return result, i + 1, nil
	})

	for item, err := range sequence {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(item)
	}

	// Output:
	// [a0 b0 c0]
	// [a1 b1 c1]
}

func ExampleProduceWithArg() {
	sequence := seqerr.ProduceWithArg(func(i int) ([]string, int, error) {
		if i == 2 {
			return []string{}, i + 1, nil
		}

		num := strconv.Itoa(i)
		result := []string{"a" + num, "b" + num, "c" + num}
		return result, i + 1, nil
	}, 1)

	for item, err := range sequence {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(item)
	}

	// Output:
	// [a1 b1 c1]
}

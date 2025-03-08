package seq2_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
)

func ExampleEmpty() {
	// Create an empty sequence
	empty := seq2.Empty[any, any]()

	seq2.ForEach(empty, func(any, any) {
		fmt.Println("Should not be called")
	})
	// Output:
}

func ExamplePair() {
	// Create a sequence with a single key-value pair
	single := seq2.Pair("key", 42)

	seq2.ForEach(single, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
	// Output:
	// key : 42
}

func ExampleOf() {
	// Create a sequence from individual elements
	indexed := seq2.Of(1, 2, 3)

	result := seq2.Collect(indexed)
	fmt.Println(result)
	// Output:
	// map[0:1 1:2 2:3]
}

func ExampleFromSlice() {
	// Create a sequence from a slice
	slice := []string{"a", "b", "c"}
	sequence := seq2.FromSlice(slice)

	result := seq2.Collect(sequence)
	fmt.Println(result)
	// Output:
	// map[0:a 1:b 2:c]
}

func ExampleWithIndex() {
	// Create a iter.Seq of values
	values := seq.Of("a", "b", "c")

	// Add indexes
	indexed := seq2.WithIndex(values)

	result := seq2.Collect(indexed)
	fmt.Println(result)
	// Output:
	// map[0:a 1:b 2:c]
}

func ExampleWithoutIndex() {
	// Create an indexed sequence
	indexed := seq2.Of("a", "b", "c")

	// Remove indexes
	values := seq2.WithoutIndex(indexed)

	result := seq.Collect(values)
	fmt.Println(result)
	// Output:
	// [a b c]
}

func ExampleRepeat() {
	// Create a sequence that repeats a key-value pair 3 times
	repeated := seq2.Repeat("key", 42, 3)

	seq2.ForEach(repeated, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
	// Output:
	// key : 42
	// key : 42
	// key : 42
}

func ExampleReverse() {
	// Create an indexed sequence
	sequence := seq2.Of("a", "b", "c")

	// Reverse it
	reversed := seq2.Reverse(sequence)

	// Collect into pairs for ordered display
	var pairs []string
	seq2.ForEach(reversed, func(k int, v string) {
		fmt.Println(k, ":", v)
	})

	fmt.Println(strings.Join(pairs, ", "))
	// Output:
	// 2 : c
	// 1 : b
	// 0 : a
}

func ExampleTick() {
	ticker := seq2.Tick(1 * time.Millisecond)

	ticker = seq2.Take(ticker, 5)

	seq2.ForEach(ticker, func(tick int, v time.Time) {
		fmt.Printf("tick %d at %s \n", tick, v.Format("15:04:05.000"))
	})

	// Example Output:
	// tick 1 at 00:00:00.000
	// tick 2 at 00:00:00.001
	// tick 3 at 00:00:00.002
	// tick 4 at 00:00:00.003
	// tick 5 at 00:00:00.004
}

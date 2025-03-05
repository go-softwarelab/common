package seq2_test

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/seq2"
)

func ExampleMapPairs() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// Map keys to uppercase and multiply values by 10
	mapped := seq2.MapPairs(input, func(k string, v int) (string, int) {
		return strings.ToUpper(k), v * 10
	})

	result := seq2.Collect(mapped)
	fmt.Println(result)
	// Output:
	// map[A:10 B:20 C:30]
}

func ExampleMap() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// Map both key and value to produce a new value (keeps original keys)
	mapped := seq2.Map(input, func(k string, v int) string {
		return fmt.Sprintf("Value of %s is %d", k, v)
	})

	result := seq2.Collect(mapped)
	fmt.Println(result)
	// Output:
	// map[a:Value of a is 1 b:Value of b is 2 c:Value of c is 3]
}

func ExampleMapKeys() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// Map keys to uppercase (keeps original values)
	mapped := seq2.MapKeys(input, func(k string) string {
		return strings.ToUpper(k)
	})

	result := seq2.Collect(mapped)
	fmt.Println(result)
	// Output:
	// map[A:1 B:2 C:3]
}

func ExampleMapValues() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// Map values to their squares (keeps original keys)
	mapped := seq2.MapValues(input, func(v int) int {
		return v * v
	})

	result := seq2.Collect(mapped)
	fmt.Println(result)
	// Output:
	// map[a:1 b:4 c:9]
}

func ExampleMapToValues() {
	input := seq2.Of("a", "b", "c")

	// Map values to strings and return as sequence of values
	mapped := seq2.MapToValues(input, func(v string) string {
		return fmt.Sprintf("Value: %s", v)
	})

	for v := range mapped {
		fmt.Println(v)
	}

	// Output:
	// Value: a
	// Value: b
	// Value: c

}

func ExampleMapTo() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// Map each key-value pair to a string
	mapped := seq2.MapTo(input, func(k string, v int) string {
		return fmt.Sprintf("%s=%d", k, v)
	})

	// Collect into a slice for display
	var result []string
	for v := range mapped {
		result = append(result, v)
	}

	// Sort for consistent output
	seq2.Sort(seq2.FromSlice(result))

	fmt.Println(result)
	// Output:
	// [a=1 b=2 c=3]
}

func ExampleNarrow() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// Narrow is an alias for MapTo
	mapped := seq2.Narrow(input, func(k string, v int) string {
		return fmt.Sprintf("%s-%d", k, v)
	})

	// Collect into a slice for display
	var result []string
	for v := range mapped {
		result = append(result, v)
	}

	// Sort for consistent output
	seq2.Sort(seq2.FromSlice(result))

	fmt.Println(result)
	// Output:
	// [a-1 b-2 c-3]
}

func ExampleCycle() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2})
	input = seq2.Sort(input)

	// Repeat the sequence 2 times
	cycled := seq2.Cycle(input, 2)

	seq2.ForEach(cycled, func(k string, v int) {
		fmt.Println(k, v)
	})

	// Output:
	// a 1
	// b 2
	// a 1
	// b 2
}

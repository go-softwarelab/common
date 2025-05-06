package seq2_test

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
)

func ExampleMap() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Map both key and value to produce a new value (keeps original keys)
	mapped := seq2.Map(input, func(k string, v int) (string, int) {
		return strings.ToUpper(k), v * 10
	})

	result := seq2.CollectToMap(mapped)
	fmt.Println(result)
	// Output:
	// map[A:10 B:20 C:30]
}

func ExampleMapKeys() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Map keys to uppercase (keeps original values)
	mapped := seq2.MapKeys(input, func(k string) string {
		return strings.ToUpper(k)
	})

	result := seq2.CollectToMap(mapped)
	fmt.Println(result)
	// Output:
	// map[A:1 B:2 C:3]
}

func ExampleMapValues() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Map values to their squares (keeps original keys)
	mapped := seq2.MapValues(input, func(v int) int {
		return v * v
	})

	result := seq2.CollectToMap(mapped)
	fmt.Println(result)
	// Output:
	// map[a:1 b:4 c:9]
}

func ExampleMapTo() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Map each key-value pair to a string
	mapped := seq2.MapTo(input, func(k string, v int) string {
		return fmt.Sprintf("%s=%d", k, v)
	})

	seq.ForEach(mapped, func(v string) {
		fmt.Println(v)
	})
	// Output:
	// a=1
	// b=2
	// c=3
}

func ExampleCycleTimes() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2})
	input = seq2.SortByKeys(input)

	// Repeat the sequence 2 times
	cycled := seq2.CycleTimes(input, 2)

	seq2.ForEach(cycled, func(k string, v int) {
		fmt.Println(k, v)
	})

	// Output:
	// a 1
	// b 2
	// a 1
	// b 2
}

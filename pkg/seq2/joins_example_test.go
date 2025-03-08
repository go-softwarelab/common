package seq2_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
)

func ExampleConcat() {
	first := seq2.FromMap(map[string]int{"a": 1, "b": 2})
	first = seq2.Sort(first)
	second := seq2.FromMap(map[string]int{"c": 3, "d": 4})
	second = seq2.Sort(second)

	// Concatenate two sequences
	combined := seq2.Concat(first, second)

	result := seq2.Collect(combined)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:3 d:4]
}

func ExampleUnion() {
	first := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	first = seq2.Sort(first)
	second := seq2.FromMap(map[string]int{"c": 3, "d": 4, "e": 5})
	second = seq2.Sort(second)

	// Union returns distinct elements from both sequences
	combined := seq2.Union(first, second)

	result := seq2.Collect(combined)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:3 d:4 e:5]
}

func ExampleUnionAll() {
	first := seq2.FromMap(map[string]int{"a": 1, "b": 2})
	first = seq2.Sort(first)
	second := seq2.FromMap(map[string]int{"c": 3, "b": 2})
	second = seq2.Sort(second)

	// UnionAll is an alias for Concat
	combined := seq2.UnionAll(first, second)

	result := seq2.Collect(combined)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:3]
}

func ExampleUnZip() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// UnZip splits a sequence into keys and values sequences
	keys, values := seq2.UnZip(input)

	keySlice := make([]string, 0)
	for k := range keys {
		keySlice = append(keySlice, k)
	}
	fmt.Printf("Keys: %v\n", keySlice)

	valueSlice := make([]int, 0)
	for v := range values {
		valueSlice = append(valueSlice, v)
	}
	fmt.Printf("Values: %v\n", valueSlice)
	// Output:
	// Keys: [a b c]
	// Values: [1 2 3]
}

func ExampleUnZip_afterZip() {
	nums := seq.Of(1, 2, 3)
	letters := seq.Of("a", "b", "c")

	zipped := seq.Zip(nums, letters)

	unzipped1, unzipped2 := seq2.UnZip(zipped)

	result1 := seq.Collect(unzipped1)
	result2 := seq.Collect(unzipped2)

	fmt.Println(result1)
	fmt.Println(result2)
	// Output:
	// [1 2 3]
	// [a b c]
}

func ExampleSplit() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.Sort(input)

	// Split is an alias for UnZip
	keys, values := seq2.Split(input)

	keySlice := make([]string, 0)
	for k := range keys {
		keySlice = append(keySlice, k)
	}
	fmt.Printf("Keys: %v\n", keySlice)

	valueSlice := make([]int, 0)
	for v := range values {
		valueSlice = append(valueSlice, v)
	}
	fmt.Printf("Values: %v\n", valueSlice)
	// Output:
	// Keys: [a b c]
	// Values: [1 2 3]
}

func ExampleAppend() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2})
	input = seq2.Sort(input)

	// Append a new key-value pair to the sequence
	appended := seq2.Append(input, "c", 3)

	result := seq2.Collect(appended)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:3]
}

func ExamplePrepend() {
	input := seq2.FromMap(map[string]int{"b": 2, "c": 3})
	input = seq2.Sort(input)

	// Prepend a new key-value pair to the sequence
	prepended := seq2.Prepend(input, "a", 1)

	result := seq2.Collect(prepended)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:3]
}

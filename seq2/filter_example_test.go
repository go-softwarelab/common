package seq2_test

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-softwarelab/common/seq2"
)

func ExampleFilter() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	// Filter elements where the value is even
	filtered := seq2.Filter(input, func(k string, v int) bool {
		return v%2 == 0
	})

	result := seq2.Collect(filtered)
	fmt.Println(result)
	// Output:
	// map[b:2 d:4]
}

func ExampleWhere() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	// Where is an alias for Filter
	filtered := seq2.Where(input, func(k string, v int) bool {
		return v > 2
	})

	result := seq2.Collect(filtered)
	fmt.Println(result)
	// Output:
	// map[c:3 d:4]
}

func ExampleFilterByKey() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	// Filter elements by key
	filtered := seq2.FilterByKey(input, func(k string) bool {
		return k > "b"
	})

	result := seq2.Collect(filtered)
	fmt.Println(result)
	// Output:
	// map[c:3 d:4]
}

func ExampleFilterByValue() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	// Filter elements by value
	filtered := seq2.FilterByValue(input, func(v int) bool {
		return v <= 2
	})

	result := seq2.Collect(filtered)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2]
}

func ExampleSkip() {
	input := seq2.Of(10, 20, 30, 40, 20)

	// Skip the first 2 elements
	skipped := seq2.Skip(input, 2)

	result := seq2.Collect(skipped)
	fmt.Println(result)
	// Output:
	// map[2:30 3:40 4:20]
}

func ExampleSkipWhile() {
	input := seq2.Of(10, 20, 30, 40, 20)

	// Skip elements until the value is less than 30
	skipped := seq2.SkipWhile(input, func(k int, v int) bool {
		return v < 30
	})

	result := seq2.Collect(skipped)
	fmt.Println(result)
	// Output:
	// map[2:30 3:40 4:20]
}

func ExampleSkipUntil() {
	input := seq2.Of(10, 20, 30, 40, 20)

	// Skip elements until value of 30 is reached
	skipped := seq2.SkipUntil(input, func(k int, v int) bool {
		return v == 30
	})

	result := seq2.Collect(skipped)
	fmt.Println(result)
	// Output:
	// map[2:30 3:40 4:20]
}

func ExampleOffset() {
	input := seq2.Of(10, 20, 30, 40, 20)

	// Skip the first 2 elements
	skipped := seq2.Offset(input, 2)

	result := seq2.Collect(skipped)
	fmt.Println(result)
	// Output:
	// map[2:30 3:40 4:20]
}

func ExampleTake() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	// Take the first 2 elements
	taken := seq2.Take(input, 2)

	result := seq2.Collect(taken)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2]
}

func ExampleTakeWhile() {
	input := seq2.Of("a", "b", "c", "d")

	// Take elements while value is less than 3
	taken := seq2.TakeWhile(input, func(k int, v string) bool {
		return v != "c"
	})

	result := seq2.Collect(taken)
	fmt.Println(result)
	// Output:
	// map[0:a 1:b]
}

func ExampleTakeUntil() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	// Take elements until value is greater than 2
	taken := seq2.TakeUntil(input, func(k string, v int) bool {
		return v > 2
	})

	result := seq2.Collect(taken)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2]
}

func ExampleLimit() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	// Limit is an alias for Take
	taken := seq2.Limit(input, 3)

	result := seq2.Collect(taken)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:3]
}

func ExampleUniq() {
	// Create a sequence with duplicate key-value pairs
	input := seq2.Concat(
		seq2.FromMap(map[string]int{"a": 1, "b": 2}),
		seq2.FromMap(map[string]int{"a": 1, "c": 3}),
	)

	// Get unique key-value pairs
	unique := seq2.Uniq(input)

	result := seq2.Collect(unique)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:3]
}

func ExampleUniqKeys() {
	// Create a sequence with duplicate keys
	input := seq2.Concat(
		seq2.FromMap(map[string]int{"a": 1, "b": 2}),
		seq2.FromMap(map[string]int{"a": 3, "c": 4}),
	)

	// Get entries with unique keys (first occurrence wins)
	unique := seq2.UniqKeys(input)

	result := seq2.Collect(unique)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:4]
}

func ExampleUniqValues() {
	// Create a sequence with duplicate values
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 1, "d": 3})
	input = seq2.Sort(input)

	// Get entries with unique values (first occurrence wins)
	unique := seq2.UniqValues(input)

	result := seq2.Collect(unique)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2 d:3]
}

func ExampleUniqBy() {
	input := seq2.FromMap(map[string]int{"apple": 1, "banana": 2, "apricot": 3, "berry": 4, "blueberry": 5})
	input = seq2.Sort(input)

	// Get unique entries based on first letter and value modulo 2
	unique := seq2.UniqBy(input, func(k string, v int) string {
		return string(k[0]) + strconv.Itoa(v%2)
	})

	result := seq2.Collect(unique)
	fmt.Println(result)
	// Output:
	// map[apple:1 banana:2 blueberry:5]
}

func ExampleUniqByKeys() {
	input := seq2.FromMap(map[string]int{"Apple": 1, "apricot": 2, "Banana": 3, "berry": 4})
	input = seq2.Sort(input)

	// Get unique entries based on lowercase first letter of key
	unique := seq2.UniqByKeys(input, func(k string) string {
		return strings.ToLower(string(k[0]))
	})

	result := seq2.Collect(unique)
	fmt.Println(result)
	// Output:
	// map[Apple:1 Banana:3]
}

func ExampleUniqByValues() {
	input := seq2.FromMap(map[string]int{"a": 10, "b": 21, "c": 30, "d": 44})
	input = seq2.Sort(input)

	// Get unique entries based on value modulo 10
	unique := seq2.UniqByValues(input, func(v int) int {
		return v % 10
	})

	result := seq2.Collect(unique)
	fmt.Println(result)
	// Output:
	// map[a:10 b:21 d:44]
}

func ExampleDistinct() {
	// Create a sequence with duplicate key-value pairs
	input := seq2.Concat(
		seq2.FromMap(map[string]int{"a": 1, "b": 2}),
		seq2.FromMap(map[string]int{"a": 1, "c": 3}),
	)

	// Distinct is an alias for Uniq
	unique := seq2.Distinct(input)

	result := seq2.Collect(unique)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:3]
}

func ExampleDistinctKeys() {
	// Create a sequence with duplicate keys
	input := seq2.Concat(
		seq2.FromMap(map[string]int{"a": 1, "b": 2}),
		seq2.FromMap(map[string]int{"a": 3, "c": 4}),
	)

	// DistinctKeys is an alias for UniqKeys
	unique := seq2.DistinctKeys(input)

	result := seq2.Collect(unique)
	fmt.Println(result)
	// Output:
	// map[a:1 b:2 c:4]
}

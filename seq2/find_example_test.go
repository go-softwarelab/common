package seq2_test

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/seq2"
)

func ExampleFind() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	k, v := seq2.Find(input, func(k string, v int) bool {
		return v > 2
	})

	fmt.Printf("Key: %s, Value: %d\n", k.MustGet(), v.MustGet())
	// Output:
	// Key: c, Value: 3
}

func ExampleFindLast() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	k, v := seq2.FindLast(input, func(k string, v int) bool {
		return v > 2
	})

	fmt.Printf("Key: %s, Value: %d\n", k.MustGet(), v.MustGet())
	// Output:
	// Key: d, Value: 4
}

func ExampleFindAll() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	found := seq2.FindAll(input, func(k string, v int) bool {
		return v > 2
	})

	result := seq2.Collect(found)
	fmt.Println(result)
	// Output:
	// map[c:3 d:4]
}

func ExampleFindByKey() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	k, v := seq2.FindByKey(input, func(k string) bool {
		return strings.ToUpper(k) == "B"
	})

	fmt.Printf("Key: %s, Value: %d\n", k.MustGet(), v.MustGet())
	// Output:
	// Key: b, Value: 2
}

func ExampleFindByValue() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	k, v := seq2.FindByValue(input, func(v int) bool {
		return v%4 == 0
	})

	fmt.Printf("Key: %s, Value: %d\n", k.MustGet(), v.MustGet())
	// Output:
	// Key: d, Value: 4
}

func ExampleGet() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	value := seq2.Get(input, "c")

	fmt.Println(value.MustGet())
	// Output:
	// 3
}

func ExampleContains() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	contains := seq2.Contains(input, "b")

	fmt.Println(contains)
	// Output:
	// true
}

func ExampleNotContains() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	notContains := seq2.NotContains(input, "x")

	fmt.Println(notContains)
	// Output:
	// true
}

func ExampleContainsValue() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	containsValue := seq2.ContainsValue(input, 3)

	fmt.Println(containsValue)
	// Output:
	// true
}

func ExampleContainsAll() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	containsAll := seq2.ContainsAll(input, "a", "c", "d")

	fmt.Println(containsAll)
	// Output:
	// true
}

func ExampleNotContainsValue() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	notContainsValue := seq2.NotContainsValue(input, 5)

	fmt.Println(notContainsValue)
	// Output:
	// true
}

func ExampleContainsAllValues() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	containsAllValues := seq2.ContainsAllValues(input, 1, 3)

	fmt.Println(containsAllValues)
	// Output:
	// true
}

func ExampleContainsPair() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	containsPair := seq2.ContainsPair(input, "b", 2)

	fmt.Println(containsPair)
	// Output:
	// true
}

func ExampleNotContainsPair() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	notContainsPair := seq2.NotContainsPair(input, "b", 3)

	fmt.Println(notContainsPair)
	// Output:
	// true
}

func ExampleExists() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	exists := seq2.Exists(input, func(k string, v int) bool {
		return k > "c" && v > 3
	})

	fmt.Println(exists)
	// Output:
	// true
}

func ExampleEvery() {
	input := seq2.FromMap(map[string]int{"a": 2, "b": 4, "c": 6, "d": 8})
	input = seq2.Sort(input)

	every := seq2.Every(input, func(k string, v int) bool {
		return v%2 == 0
	})

	fmt.Println(every)
	// Output:
	// true
}

func ExampleNone() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	none := seq2.None(input, func(k string, v int) bool {
		return v > 10
	})

	fmt.Println(none)
	// Output:
	// true
}

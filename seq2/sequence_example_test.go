package seq2_test

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/seq2"
)

func ExampleAsSequence() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	var result []string
	sequence.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 0:a, 1:b, 2:c
}

func ExampleConcatSequences() {
	sq1 := seq2.AsSequence(seq2.Of("a", "b"))
	sq2 := seq2.AsSequence(seq2.Of("c", "d"))

	concatenated := seq2.ConcatSequences(sq1, sq2)

	concatenated.ForEach(func(k int, v string) {
		fmt.Println(k, v)
	})

	// Output:
	// 0 a
	// 1 b
	// 0 c
	// 1 d
}

func ExampleSequence_Union() {
	sq1 := seq2.AsSequence(seq2.Concat(seq2.Pair("a", 1), seq2.Pair("b", 2), seq2.Pair("c", 3)))
	sq2 := seq2.AsSequence(seq2.Concat(seq2.Pair("a", 1), seq2.Pair("d", 4), seq2.Pair("e", 5)))

	union := sq1.Union(sq2)

	union.ForEach(func(k string, v int) {
		fmt.Println(k, ":", v)
	})

	// Output:
	// a : 1
	// b : 2
	// c : 3
	// d : 4
	// e : 5
}

func ExampleSequence_UnionAll() {
	seq1 := seq2.AsSequence(seq2.Of("a", "b", "c"))
	seq2Inst := seq2.AsSequence(seq2.Of("c", "d", "e"))

	unionAll := seq1.UnionAll(seq2Inst)

	var result []string
	unionAll.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 0:a, 1:b, 2:c, 0:c, 1:d, 2:e
}

func ExampleSequence_Append() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))
	appended := sequence.Append(3, "d")

	var result []string
	appended.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 0:a, 1:b, 2:c, 3:d
}

func ExampleSequence_Prepend() {
	sequence := seq2.AsSequence(seq2.Of("b", "c", "d"))
	prepended := sequence.Prepend(-1, "a")

	var result []string
	prepended.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// -1:a, 0:b, 1:c, 2:d
}

func ExampleSequence_Filter() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))
	filtered := sequence.Filter(func(k int, v string) bool {
		return k%2 == 0
	})

	var result []string
	filtered.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 0:a, 2:c, 4:e
}

func ExampleSequence_FilterByKey() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))
	filtered := sequence.FilterByKey(func(k int) bool {
		return k > 2
	})

	var result []string
	filtered.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 3:d, 4:e
}

func ExampleSequence_FilterByValue() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))
	filtered := sequence.FilterByValue(func(v string) bool {
		return v > "c"
	})

	var result []string
	filtered.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 3:d, 4:e
}

func ExampleSequence_Skip() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))
	skipped := sequence.Skip(2)

	var result []string
	skipped.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 2:c, 3:d, 4:e
}

func ExampleSequence_Take() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))
	taken := sequence.Take(3)

	var result []string
	taken.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 0:a, 1:b, 2:c
}

func ExampleSequence_ForEach() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	var result []string
	sequence.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 0:a, 1:b, 2:c
}

func ExampleSequence_Tap() {
	var result []string
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c")).
		Tap(func(k int, v string) {
			result = append(result, fmt.Sprintf("%d:%s", k, v))
		})

	sequence.Flush()
	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 0:a, 1:b, 2:c
}

func ExampleSequence_Keys() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))
	keys := sequence.Keys()

	var result []int
	for k := range keys {
		result = append(result, k)
	}

	fmt.Println(result)
	// Output:
	// [0 1 2]
}

func ExampleSequence_Values() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))
	values := sequence.Values()

	var result []string
	for v := range values {
		result = append(result, v)
	}

	fmt.Println(result)
	// Output:
	// [a b c]
}

func ExampleSequence_Find() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))

	key, value := sequence.Find(func(k int, v string) bool {
		return v == "c"
	})

	fmt.Printf("Key: %d, Value: %s\n", key.MustGet(), value.MustGet())
	// Output:
	// Key: 2, Value: c
}

func ExampleSequence_FoldValues() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	result := sequence.FoldValues(func(agg string, key int, value string) string {
		return agg + value
	})

	fmt.Println("Concatenated:", result.MustGet())
	// Output:
	// Concatenated: abc
}

func ExampleSequence_Reverse() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))
	reversed := sequence.Reverse()

	var result []string
	reversed.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 2:c, 1:b, 0:a
}

func ExampleSequence_Cycle() {
	sequence := seq2.AsSequence(seq2.Of("a", "b"))
	cycled := sequence.Cycle(2)

	var result []string
	cycled.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
	// Output:
	// 0:a, 1:b, 0:a, 1:b
}

func ExampleSequence_Contains() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	fmt.Println("Contains key 1:", sequence.Contains(1))
	fmt.Println("Contains key 5:", sequence.Contains(5))
	// Output:
	// Contains key 1: true
	// Contains key 5: false
}

func ExampleSequence_ContainsValue() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	fmt.Println("Contains value 'b':", sequence.ContainsValue("b"))
	fmt.Println("Contains value 'z':", sequence.ContainsValue("z"))
	// Output:
	// Contains value 'b': true
	// Contains value 'z': false
}

func ExampleSequence_ContainsPair() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	fmt.Println("Contains pair 1:'b':", sequence.ContainsPair(1, "b"))
	fmt.Println("Contains pair 1:'c':", sequence.ContainsPair(1, "c"))
	// Output:
	// Contains pair 1:'b': true
	// Contains pair 1:'c': false
}

func ExampleSequence_Exists() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	hasValueB := sequence.Exists(func(k int, v string) bool {
		return v == "b"
	})

	fmt.Println("Has value 'b':", hasValueB)
	// Output:
	// Has value 'b': true
}

func ExampleSequence_Every() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	allLowercase := sequence.Every(func(k int, v string) bool {
		return v >= "a" && v <= "z"
	})

	allA := sequence.Every(func(k int, v string) bool {
		return v == "a"
	})

	fmt.Println("All values lowercase:", allLowercase)
	fmt.Println("All values 'a':", allA)
	// Output:
	// All values lowercase: true
	// All values 'a': false
}

func ExampleSequence_Uniq() {
	// Using FromMap to create a sequence with duplicate values
	sequence := seq2.AsSequence(seq2.Of("a", "a", "b", "b", "c"))
	unique := sequence.UniqValues()

	unique.ForEach(func(k int, v string) {
		fmt.Println(k, ":", v)
	})

	// Output:
	// 0 : a
	// 2 : b
	// 4 : c
}

func ExampleSequence_Get() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	existing := sequence.Get(1)
	notExisting := sequence.Get(5)

	fmt.Println("Value for key 1:", existing.MustGet())
	fmt.Println("Has value for key 5:", notExisting.IsPresent())
	// Output:
	// Value for key 1: b
	// Has value for key 5: false
}

package seq_test

import (
	"fmt"

	"github.com/go-softwarelab/common/seq"
)

func ExampleAsSequence() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	result := sequence.Collect()
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleConcatSequences() {
	seq1 := seq.AsSequence(seq.Of(1, 2))
	seq2 := seq.AsSequence(seq.Of(3, 4))
	concatenated := seq.ConcatSequences(seq1, seq2)
	result := concatenated.Collect()
	fmt.Println(result)
	// Output:
	// [1 2 3 4]
}

func ExampleSequence_Union() {
	seq1 := seq.AsSequence(seq.Of(1, 2, 3))
	seq2 := seq.AsSequence(seq.Of(3, 4, 5))
	union := seq1.Union(seq2)
	result := union.Collect()
	fmt.Println(result)
	// Output:
	// [1 2 3 4 5]
}

func ExampleSequence_UnionAll() {
	seq1 := seq.AsSequence(seq.Of(1, 2, 3))
	seq2 := seq.AsSequence(seq.Of(3, 4, 5))
	unionAll := seq1.UnionAll(seq2)
	result := unionAll.Collect()
	fmt.Println(result)
	// Output:
	// [1 2 3 3 4 5]
}

func ExampleSequence_Append() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	appended := sequence.Append(4, 5)
	result := appended.Collect()
	fmt.Println(result)
	// Output:
	// [1 2 3 4 5]
}

func ExampleSequence_Prepend() {
	sequence := seq.AsSequence(seq.Of(3, 4, 5))
	prepended := sequence.Prepend(1, 2)
	result := prepended.Collect()
	fmt.Println(result)
	// Output:
	// [1 2 3 4 5]
}

func ExampleSequence_Filter() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	filtered := sequence.Filter(func(v int) bool {
		return v%2 == 0
	})
	result := filtered.Collect()
	fmt.Println(result)
	// Output:
	// [2 4]
}

func ExampleSequence_Where() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	filtered := sequence.Where(func(v int) bool {
		return v%2 == 0
	})
	result := filtered.Collect()
	fmt.Println(result)
	// Output:
	// [2 4]
}

func ExampleSequence_Skip() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	skipped := sequence.Skip(2)
	result := skipped.Collect()
	fmt.Println(result)
	// Output:
	// [3 4 5]
}

func ExampleSequence_Offset() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	offset := sequence.Offset(2)
	result := offset.Collect()
	fmt.Println(result)
	// Output:
	// [3 4 5]
}

func ExampleSequence_Take() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	taken := sequence.Take(3)
	result := taken.Collect()
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleSequence_Limit() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	limited := sequence.Limit(2)
	result := limited.Collect()
	fmt.Println(result)
	// Output:
	// [1 2]
}

func ExampleSequence_Tap() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3)).Tap(func(v int) {
		fmt.Println(v)
	})
	sequence.Flush()
	// Output:
	// 1
	// 2
	// 3
}

func ExampleSequence_Each() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3)).Each(func(v int) {
		fmt.Println(v)
	})
	sequence.Flush()
	// Output:
	// 1
	// 2
	// 3
}

func ExampleSequence_ForEach() {
	seq.AsSequence(seq.Of(1, 2, 3)).ForEach(func(v int) {
		fmt.Println(v)
	})
	// Output:
	// 1
	// 2
	// 3
}

func ExampleSequence_Flush() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	sequence.Flush()
	// No output expected
}

func ExampleSequence_Collect() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	result := sequence.Collect()
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleSequence_Count() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	count := sequence.Count()
	fmt.Println(count)
	// Output:
	// 3
}

func ExampleSequence_ToSlice() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	var slice []int
	result := sequence.ToSlice(slice)
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleSequence_Find() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	found := sequence.Find(func(v int) bool {
		return v > 3
	})

	fmt.Println(found.MustGet())
	// Output:
	// 4
}

func ExampleSequence_FindLast() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	found := sequence.FindLast(func(v int) bool {
		return v > 3
	})

	fmt.Println(found.MustGet())
	// Output:
	// 5
}

func ExampleSequence_FindAll() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	foundAll := sequence.FindAll(func(v int) bool {
		return v > 3
	})
	result := foundAll.Collect()
	fmt.Println(result)
	// Output:
	// [4 5]
}

func ExampleSequence_Contains() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	contains := sequence.Contains(3)
	fmt.Println(contains)
	// Output:
	// true
}

func ExampleSequence_NotContains() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	notContains := sequence.NotContains(6)
	fmt.Println(notContains)
	// Output:
	// true
}

func ExampleSequence_ContainsAll() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	containsAll := sequence.ContainsAll(2, 3, 4)
	fmt.Println(containsAll)
	// Output:
	// true
}

func ExampleSequence_Exists() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	exists := sequence.Exists(func(v int) bool {
		return v > 3
	})
	fmt.Println(exists)
	// Output:
	// true
}

func ExampleSequence_Every() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	every := sequence.Every(func(v int) bool {
		return v > 0
	})
	fmt.Println(every)
	// Output:
	// true
}

func ExampleSequence_None() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	none := sequence.None(func(v int) bool {
		return v > 5
	})
	fmt.Println(none)
	// Output:
	// true
}

func ExampleSequence_IsNotEmpty() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	isNotEmpty := sequence.IsNotEmpty()
	fmt.Println(isNotEmpty)
	// Output:
	// true
}

func ExampleSequence_IsEmpty() {
	sequence := seq.AsSequence(seq.Of[int]())
	isEmpty := sequence.IsEmpty()
	fmt.Println(isEmpty)
	// Output:
	// true
}

func ExampleSequence_Partition() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	partitions := sequence.Partition(2)
	for partition := range partitions {
		fmt.Println(seq.Collect(partition))
	}
	// Output:
	// [1 2]
	// [3 4]
	// [5]
}

func ExampleSequence_Uniq() {
	sequence := seq.AsSequence(seq.Of(1, 2, 2, 3, 3, 3))
	unique := sequence.Uniq()
	result := unique.Collect()
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleSequence_Distinct() {
	sequence := seq.AsSequence(seq.Of(1, 2, 2, 3, 3, 3))
	distinct := sequence.Distinct()
	result := distinct.Collect()
	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleSequence_Reverse() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	reversed := sequence.Reverse()
	result := reversed.Collect()
	fmt.Println(result)
	// Output:
	// [3 2 1]
}

func ExampleSequence_Repeat() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	repeated := sequence.Repeat(2)
	result := repeated.Collect()
	fmt.Println(result)
	// Output:
	// [1 2 3 1 2 3]
}

func ExampleSequence_Fold() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	sum := sequence.Fold(func(agg, item int) int {
		return agg + item
	})

	fmt.Println(sum.MustGet())
	// Output:
	// 15
}

func ExampleSequence_FoldRight() {
	sequence := seq.AsSequence(seq.Of("a", "b", "c"))
	concat := sequence.FoldRight(func(agg, item string) string {
		return agg + item
	})

	fmt.Println(concat.MustGet())
	// Output:
	// cba
}

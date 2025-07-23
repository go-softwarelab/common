package is_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
	"github.com/go-softwarelab/common/pkg/seq"
)

// Example for is.Unique with a slice
func ExampleUnique_slice() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 1, 2}

	fmt.Println(
		is.Unique[int](slice1),
		is.Unique[int](slice2),
	)
	// Output:
	// true false
}

// Example for is.Unique with iter.Seq
func ExampleUnique_seq() {
	seq1 := seq.Of(1, 2, 3)
	seq2 := seq.Of(1, 1, 2)
	fmt.Println(is.Unique[int](seq1), is.Unique[int](seq2))
	// Output:
	// true false
}

// Example for is.UniqueBy with a slice
func ExampleUniqueBy_slice() {
	type item struct {
		ID   int
		Name string
	}
	s1 := []item{{1, "a"}, {2, "b"}, {3, "c"}}
	s2 := []item{{1, "a"}, {2, "b"}, {1, "d"}}

	key := func(i item) int { return i.ID }

	fmt.Println(is.UniqueBy(s1, key), is.UniqueBy(s2, key))
	// Output:
	// true false
}

// Example for is.UniqueBy with iter.Seq
func ExampleUniqueBy_seq() {
	type item struct {
		ID   int
		Name string
	}
	seq1 := seq.Of(item{1, "a"}, item{2, "b"}, item{3, "c"})
	seq2 := seq.Of(item{1, "a"}, item{2, "b"}, item{1, "c"})

	key := func(i item) int { return i.ID }

	fmt.Println(is.UniqueBy(seq1, key), is.UniqueBy(seq2, key))
	// Output:
	// true false
}

// Example for is.UniqueSlice
func ExampleUniqueSlice() {
	fmt.Println(is.UniqueSlice([]string{"x", "y"}), is.UniqueSlice([]string{"x", "x"}))
	// Output:
	// true false
}

// Example for is.UniqueSeq
func ExampleUniqueSeq() {
	seq1 := seq.Of("foo", "bar")
	seq2 := seq.Of("foo", "foo")
	fmt.Println(is.UniqueSeq(seq1), is.UniqueSeq(seq2))
	// Output:
	// true false
}

// Example for is.UniqueSliceBy
func ExampleUniqueSliceBy() {
	type T struct{ X int }
	s1 := []T{{1}, {2}, {3}}
	s2 := []T{{1}, {2}, {1}}
	key := func(t T) int { return t.X }
	fmt.Println(is.UniqueSliceBy(s1, key), is.UniqueSliceBy(s2, key))
	// Output:
	// true false
}

// Example for is.UniqueSeqBy
func ExampleUniqueSeqBy() {
	type T struct{ X int }
	seq1 := seq.Of(T{1}, T{2}, T{3})
	seq2 := seq.Of(T{1}, T{2}, T{1})
	key := func(t T) int { return t.X }
	fmt.Println(is.UniqueSeqBy(seq1, key), is.UniqueSeqBy(seq2, key))
	// Output:
	// true false
}

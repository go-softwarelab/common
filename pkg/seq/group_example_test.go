package seq_test

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seq"
)

func ExamplePartition() {
	input := seq.Of(1, 2, 3, 4, 5, 6)

	partitions := seq.Partition(input, 2)

	result := seq.Collect(partitions)

	for _, partition := range result {
		fmt.Println(seq.Collect(partition))
	}

	// Output:
	// [1 2]
	// [3 4]
	// [5 6]
}

func ExampleChunk() {
	input := seq.Of(1, 2, 3, 4, 5, 6)

	chunks := seq.Chunk(input, 3)

	result := seq.Collect(chunks)

	for _, chunk := range result {
		fmt.Println(seq.Collect(chunk))
	}
	// Output:
	// [1 2 3]
	// [4 5 6]
}

func ExamplePartitionBy() {
	input := seq.Of(1, 2, 3, 4, 5, 6)

	partitions := seq.PartitionBy(input, func(v int) int {
		return v % 2
	})

	// Notice that the order of the partitions is not guaranteed
	sorted := seq.SortBy(partitions, func(p iter.Seq[int]) int {
		next, stop := iter.Pull(p)
		defer stop()
		v, _ := next()
		return v
	})

	for partition := range sorted {
		fmt.Println(seq.Collect(partition))
	}
	// Output:
	// [1 3 5]
	// [2 4 6]
}

func ExampleGroupBy() {
	input := seq.Of(1, 2, 3, 4, 5, 6)

	groups := seq.GroupBy(input, func(v int) int {
		return v % 2
	})

	for k, v := range groups {
		fmt.Printf("%d: %v\n", k, seq.Collect(v))
	}

	// Output:
	// 1: [1 3 5]
	// 0: [2 4 6]
}

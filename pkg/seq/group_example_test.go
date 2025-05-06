package seq_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
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
	input := seq.Of(1, 2, 3, 4, 1, 5, 6)

	partitions := seq.PartitionBy(input, func(v int) int {
		return (v - 1) / 3
	})

	for partition := range partitions {
		fmt.Println(seq.Collect(partition))
	}
	// Output:
	// [1 2 3]
	// [4]
	// [1]
	// [5 6]
}

func ExampleGroupBy() {
	input := seq.Of(1, 2, 3, 4, 5, 6)

	groups := seq.GroupBy(input, func(v int) int {
		return v % 2
	})

	// GroupBy does not guarantee the order of keys, so we sort them for display
	groups = seq2.SortByKeys(groups)
	for k, v := range groups {
		fmt.Printf("%d: %v\n", k, seq.Collect(v))
	}

	// Output:
	// 0: [2 4 6]
	// 1: [1 3 5]
}

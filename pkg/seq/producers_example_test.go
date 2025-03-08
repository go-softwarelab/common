package seq_test

import (
	"fmt"
	"time"

	"github.com/go-softwarelab/common/pkg/seq"
)

func ExampleRepeat() {
	repeated := seq.Repeat("hello", 3)

	result := seq.Collect(repeated)

	fmt.Println(result)
	// Output:
	// [hello hello hello]
}

func ExampleRangeWithStep() {
	ranged := seq.RangeWithStep(0, 10, 2)

	result := seq.Collect(ranged)

	fmt.Println(result)
	// Output:
	// [0 2 4 6 8]
}

func ExampleRange() {
	ranged := seq.Range(0, 5)

	result := seq.Collect(ranged)

	fmt.Println(result)
	// Output:
	// [0 1 2 3 4]
}

func ExampleRangeTo() {
	ranged := seq.RangeTo(5)

	result := seq.Collect(ranged)

	fmt.Println(result)
	// Output:
	// [0 1 2 3 4]
}

func ExampleTick() {
	ticker := seq.Tick(1 * time.Millisecond)

	ticker = seq.Take(ticker, 5)

	ticker = seq.Tap(ticker, func(v time.Time) {
		fmt.Println(v.Format("15:04:05.000"))
	})

	seq.Flush(ticker)

	// Example Output:
	// 00:00:00.000
	// 00:00:00.001
	// 00:00:00.002
	// 00:00:00.003
	// 00:00:00.004
}

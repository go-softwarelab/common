# seq

```go
import "github.com/go-softwarelab/common/pkg/seq"
```



<a name="Append"></a>
## Append

```go
func Append[E any](seq iter.Seq[E], elems ...E) iter.Seq[E]
```

Append appends elements to the end of a sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	initial := seq.Of(1, 2, 3)

	appended := seq.Append(initial, 4, 5, 6)

	result := seq.Collect(appended)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3 4 5 6]
```


</details>

<a name="Chunk"></a>
## Chunk

```go
func Chunk[E any](seq iter.Seq[E], size int) iter.Seq[iter.Seq[E]]
```

Chunk splits the sequence into chunks of the given size.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5, 6)

	chunks := seq.Chunk(input, 3)

	result := seq.Collect(chunks)

	for _, chunk := range result {
		fmt.Println(seq.Collect(chunk))
	}
}
```

**Output**

```
[1 2 3]
[4 5 6]
```


</details>

<a name="Collect"></a>
## Collect

```go
func Collect[E any](seq iter.Seq[E]) []E
```

Collect collects the elements of the given sequence into a slice.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.Of(1, 2, 3)

	result := seq.Collect(sequence)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="Concat"></a>
## Concat

```go
func Concat[E any](sequences ...iter.Seq[E]) iter.Seq[E]
```

Concat concatenates multiple sequences into a single sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	seq1 := seq.Of(1, 2, 3)
	seq2 := seq.Of(4, 5, 6)

	concatenated := seq.Concat(seq1, seq2)

	result := seq.Collect(concatenated)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3 4 5 6]
```


</details>

<a name="Contains"></a>
## Contains

```go
func Contains[E comparable](seq iter.Seq[E], elem E) bool
```

Contains returns true if the element is in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	contains := seq.Contains(input, 3)

	fmt.Println(contains)
}
```

**Output**

```
true
```


</details>

<a name="ContainsAll"></a>
## ContainsAll

```go
func ContainsAll[E comparable](seq iter.Seq[E], elements ...E) bool
```

ContainsAll returns true if all elements are in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	containsAll := seq.ContainsAll(input, 2, 3)

	fmt.Println(containsAll)
}
```

**Output**

```
true
```


</details>

<a name="Count"></a>
## Count

```go
func Count[E any](seq iter.Seq[E]) int
```

Count returns the number of elements in the sequence.

<a name="Cycle"></a>
## Cycle

```go
func Cycle[E any](seq iter.Seq[E], count int) iter.Seq[E]
```

Cycle repeats the sequence count times.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3)

	cycled := seq.Cycle(input, 2)

	result := seq.Collect(cycled)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3 1 2 3]
```


</details>

<a name="Distinct"></a>
## Distinct

```go
func Distinct[E comparable](seq iter.Seq[E]) iter.Seq[E]
```

Distinct returns a sequence with only unique elements. SQL\-like alias for Uniq

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 2, 3, 3, 3)

	distinct := seq.Distinct(input)

	result := seq.Collect(distinct)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="Each"></a>
## Each

```go
func Each[E any](seq iter.Seq[E], consumer Consumer[E]) iter.Seq[E]
```

Each is an alias for Tap.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.Each(seq.Of(1, 2, 3), func(v int) {
		fmt.Println(v)
	})

	seq.Flush(sequence)

}
```

**Output**

```
1
2
3
```


</details>

<a name="Every"></a>
## Every

```go
func Every[E any](seq iter.Seq[E], predicate Predicate[E]) bool
```

Every returns true if all elements satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(2, 4, 6, 8)

	every := seq.Every(input, func(v int) bool { return v%2 == 0 })

	fmt.Println(every)
}
```

**Output**

```
true
```


</details>

<a name="Exists"></a>
## Exists

```go
func Exists[E any](seq iter.Seq[E], predicate Predicate[E]) bool
```

Exists returns true if there is at least one element that satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	exists := seq.Exists(input, func(v int) bool { return v > 4 })

	fmt.Println(exists)
}
```

**Output**

```
true
```


</details>

<a name="Filter"></a>
## Filter

```go
func Filter[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E]
```

Filter returns a new sequence that contains only the elements that satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	filtered := seq.Filter(input, func(v int) bool {
		return v%2 == 0
	})

	result := seq.Collect(filtered)

	fmt.Printf("%v\n", result)
}
```

**Output**

```
[2 4]
```


</details>

<a name="Find"></a>
## Find

```go
func Find[E any](seq iter.Seq[E], predicate Predicate[E]) optional.Elem[E]
```

Find returns the first element that satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	found := seq.Find(input, func(v int) bool { return v > 3 })

	fmt.Println(found.MustGet())
}
```

**Output**

```
4
```


</details>

<a name="FindAll"></a>
## FindAll

```go
func FindAll[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E]
```

FindAll returns all elements that satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	found := seq.FindAll(input, func(v int) bool { return v > 3 })

	result := seq.Collect(found)

	fmt.Println(result)
}
```

**Output**

```
[4 5]
```


</details>

<a name="FindLast"></a>
## FindLast

```go
func FindLast[E any](seq iter.Seq[E], predicate Predicate[E]) optional.Elem[E]
```

FindLast returns the last element that satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	found := seq.FindLast(input, func(v int) bool { return v > 3 })

	fmt.Println(found.MustGet())
}
```

**Output**

```
5
```


</details>

<a name="FlatMap"></a>
## FlatMap

```go
func FlatMap[E any, R any](seq iter.Seq[E], mapper Mapper[E, iter.Seq[R]]) iter.Seq[R]
```

FlatMap applies a mapper function to each element of the sequence and flattens the result.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(0, 3)

	flatMapped := seq.FlatMap(input, func(it int) iter.Seq[int] {
		return seq.Of[int](1+it, 2+it, 3+it)
	})

	result := seq.Collect(flatMapped)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3 4 5 6]
```


</details>

<a name="Flatten"></a>
## Flatten

```go
func Flatten[Seq iter.Seq[iter.Seq[E]], E any](seq Seq) iter.Seq[E]
```

Flatten flattens a sequence of sequences.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(seq.Of(1, 2), seq.Of(3, 4))

	flattened := seq.Flatten(input)

	result := seq.Collect(flattened)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3 4]
```


</details>

<a name="Flush"></a>
## Flush

```go
func Flush[E any](seq iter.Seq[E])
```

Flush consumes all elements of the input sequence.

<a name="Fold"></a>
## Fold

```go
func Fold[E any](seq iter.Seq[E], accumulator func(agg E, item E) E) optional.Elem[E]
```

Fold applies a function against an accumulator and each element in the sequence \(from left to right\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of("a", "b", "c")

	sum := seq.Fold(input, func(agg, item string) string {
		return agg + item
	})

	fmt.Println(sum.MustGet())
}
```

**Output**

```
abc
```


</details>

<a name="FoldRight"></a>
## FoldRight

```go
func FoldRight[E any](seq iter.Seq[E], accumulator func(agg E, item E) E) optional.Elem[E]
```

FoldRight applies a function against an accumulator and each element in the sequence \(from right to left\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of("a", "b", "c")

	sum := seq.FoldRight(input, func(agg, item string) string {
		return agg + item
	})

	fmt.Println(sum.MustGet())
}
```

**Output**

```
cba
```


</details>

<a name="ForEach"></a>
## ForEach

```go
func ForEach[E any](seq iter.Seq[E], consumer Consumer[E])
```

ForEach applies consumer to each element of the input sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	seq.ForEach(seq.Of(1, 2, 3), func(v int) {
		fmt.Println(v)
	})

}
```

**Output**

```
1
2
3
```


</details>

<a name="FromSlice"></a>
## FromSlice

```go
func FromSlice[Slice ~[]E, E any](slice Slice) iter.Seq[E]
```

FromSlice creates a new sequence from the given slice.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	slice := []int{1, 2, 3}

	sequence := seq.FromSlice(slice)

	result := seq.Collect(sequence)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="GroupBy"></a>
## GroupBy

```go
func GroupBy[E any, K comparable](seq iter.Seq[E], by Mapper[E, K]) iter.Seq2[K, iter.Seq[E]]
```

GroupBy groups the sequence by the given key.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5, 6)

	groups := seq.GroupBy(input, func(v int) int {
		return v % 2
	})

	for k, v := range groups {
		fmt.Printf("%d: %v\n", k, seq.Collect(v))
	}

}
```

**Output**

```
1: [1 3 5]
0: [2 4 6]
```


</details>

<a name="IsEmpty"></a>
## IsEmpty

```go
func IsEmpty[E any](seq iter.Seq[E]) bool
```

IsEmpty returns true if the sequence is empty.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3)

	fmt.Println(seq.IsEmpty(input))
}
```

**Output**

```
false
```


</details>

<a name="IsNotEmpty"></a>
## IsNotEmpty

```go
func IsNotEmpty[E any](seq iter.Seq[E]) bool
```

IsNotEmpty returns true if the sequence is not empty.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3)

	isNotEmpty := seq.IsNotEmpty(input)

	fmt.Println(isNotEmpty)
}
```

**Output**

```
true
```


</details>

<a name="Limit"></a>
## Limit

```go
func Limit[E any](seq iter.Seq[E], n int) iter.Seq[E]
```

Limit returns a new sequence that contains only the first n elements of the given sequence. SQL\-like alias for Take

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	limited := seq.Limit(input, 2)

	result := seq.Collect(limited)

	fmt.Printf("%v\n", result)
}
```

**Output**

```
[1 2]
```


</details>

<a name="Map"></a>
## Map

```go
func Map[E any, R any](seq iter.Seq[E], mapper Mapper[E, R]) iter.Seq[R]
```

Map applies a mapper function to each element of the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3)

	mapped := seq.Map(input, func(v int) string {
		return fmt.Sprintf("Number_%d", v)
	})

	result := seq.Collect(mapped)

	fmt.Println(result)
}
```

**Output**

```
[Number_1 Number_2 Number_3]
```


</details>

<a name="Max"></a>
## Max

```go
func Max[E types.Ordered](seq iter.Seq[E]) optional.Elem[E]
```

Max returns the maximum element in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(2, 3, 1, 5, 4)

	maxVal := seq.Max(input)

	fmt.Println(maxVal.MustGet())
}
```

**Output**

```
5
```


</details>

<a name="Min"></a>
## Min

```go
func Min[E types.Ordered](seq iter.Seq[E]) optional.Elem[E]
```

Min returns the minimum element in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(2, 3, 1, 5, 4)

	maxVal := seq.Min(input)

	fmt.Println(maxVal.MustGet())
}
```

**Output**

```
1
```


</details>

<a name="None"></a>
## None

```go
func None[E any](seq iter.Seq[E], predicate Predicate[E]) bool
```

None returns true if no element satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	none := seq.None(input, func(v int) bool { return v > 5 })

	fmt.Println(none)
}
```

**Output**

```
true
```


</details>

<a name="NotContains"></a>
## NotContains

```go
func NotContains[E comparable](seq iter.Seq[E], elem E) bool
```

NotContains returns true if the element is not in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	contains := seq.NotContains(input, 3)

	fmt.Println(contains)
}
```

**Output**

```
false
```


</details>

<a name="Of"></a>
## Of

```go
func Of[E any](elems ...E) iter.Seq[E]
```

Of creates a new sequence from the given elements.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.Of(1, 2, 3)

	result := seq.Collect(sequence)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="Offset"></a>
## Offset

```go
func Offset[E any](seq iter.Seq[E], n int) iter.Seq[E]
```

Offset returns a new sequence that skips the first n elements of the given sequence. SQL\-like alias for Skip

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	skipped := seq.Offset(input, 2)

	result := seq.Collect(skipped)

	fmt.Printf("%v\n", result)
}
```

**Output**

```
[3 4 5]
```


</details>

<a name="Partition"></a>
## Partition

```go
func Partition[E any](seq iter.Seq[E], size int) iter.Seq[iter.Seq[E]]
```

Partition splits the sequence into chunks of the given size.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5, 6)

	partitions := seq.Partition(input, 2)

	result := seq.Collect(partitions)

	for _, partition := range result {
		fmt.Println(seq.Collect(partition))
	}

}
```

**Output**

```
[1 2]
[3 4]
[5 6]
```


</details>

<a name="PartitionBy"></a>
## PartitionBy

```go
func PartitionBy[E any, K comparable](seq iter.Seq[E], by Mapper[E, K]) iter.Seq[iter.Seq[E]]
```

PartitionBy splits the sequence into chunks based on the given key.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
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
}
```

**Output**

```
[1 3 5]
[2 4 6]
```


</details>

<a name="Prepend"></a>
## Prepend

```go
func Prepend[E any](seq iter.Seq[E], elems ...E) iter.Seq[E]
```

Prepend prepends elements to the beginning of a sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	initial := seq.Of(4, 5, 6)

	prepended := seq.Prepend(initial, 1, 2, 3)

	result := seq.Collect(prepended)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3 4 5 6]
```


</details>

<a name="Range"></a>
## Range

```go
func Range[E types.Integer](start, end E) iter.Seq[E]
```

Range returns a sequence that yields integers from \`start\` to \`end\`.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	ranged := seq.Range(0, 5)

	result := seq.Collect(ranged)

	fmt.Println(result)
}
```

**Output**

```
[0 1 2 3 4]
```


</details>

<a name="RangeTo"></a>
## RangeTo

```go
func RangeTo[E types.Integer](end E) iter.Seq[E]
```

RangeTo returns a sequence that yields integers from 0 to \`end\`.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	ranged := seq.RangeTo(5)

	result := seq.Collect(ranged)

	fmt.Println(result)
}
```

**Output**

```
[0 1 2 3 4]
```


</details>

<a name="RangeWithStep"></a>
## RangeWithStep

```go
func RangeWithStep[E types.Integer](start, end, step E) iter.Seq[E]
```

RangeWithStep returns a sequence that yields integers from \`start\` to \`end\` with \`step\`.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	ranged := seq.RangeWithStep(0, 10, 2)

	result := seq.Collect(ranged)

	fmt.Println(result)
}
```

**Output**

```
[0 2 4 6 8]
```


</details>

<a name="Reduce"></a>
## Reduce

```go
func Reduce[E any, R any](seq iter.Seq[E], accumulator func(agg R, item E) R, initial R) R
```

Reduce applies a function against an accumulator and each element in the sequence \(from left to right\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of("a", "b", "c")

	concat := seq.Reduce(input, func(agg, item string) string {
		return agg + item
	}, "")

	fmt.Println(concat)
}
```

**Output**

```
abc
```


</details>

<a name="ReduceRight"></a>
## ReduceRight

```go
func ReduceRight[E any, R any](seq iter.Seq[E], accumulator func(agg R, item E) R, initial R) R
```

ReduceRight applies a function against an accumulator and each element in the sequence \(from right to left\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of("a", "b", "c")

	concat := seq.ReduceRight(input, func(agg, item string) string {
		return agg + item
	}, "")

	fmt.Println(concat)
}
```

**Output**

```
cba
```


</details>

<a name="Repeat"></a>
## Repeat

```go
func Repeat[E any](elem E, count int) iter.Seq[E]
```

Repeat returns a sequence that yields the same element \`count\` times.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	repeated := seq.Repeat("hello", 3)

	result := seq.Collect(repeated)

	fmt.Println(result)
}
```

**Output**

```
[hello hello hello]
```


</details>

<a name="Reverse"></a>
## Reverse

```go
func Reverse[E any](seq iter.Seq[E]) iter.Seq[E]
```

Reverse creates a new sequence that iterates over the elements of the given sequence in reverse order.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.Of(1, 2, 3)

	reversed := seq.Reverse(sequence)

	result := seq.Collect(reversed)
	fmt.Println(result)
}
```

**Output**

```
[3 2 1]
```


</details>

<a name="Select"></a>
## Select

```go
func Select[E any, R any](seq iter.Seq[E], mapper Mapper[E, R]) iter.Seq[R]
```

Select applies a mapper function to each element of the sequence. SQL\-like alias for Map

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3)

	mapped := seq.Select(input, func(v int) string {
		return fmt.Sprintf("Number_%d", v)
	})

	result := seq.Collect(mapped)

	fmt.Println(result)
}
```

**Output**

```
[Number_1 Number_2 Number_3]
```


</details>

<a name="Skip"></a>
## Skip

```go
func Skip[E any](seq iter.Seq[E], n int) iter.Seq[E]
```

Skip returns a new sequence that skips the first n elements of the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	skipped := seq.Skip(input, 2)

	result := seq.Collect(skipped)

	fmt.Printf("%v\n", result)
}
```

**Output**

```
[3 4 5]
```


</details>

<a name="SkipUntil"></a>
## SkipUntil

```go
func SkipUntil[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E]
```

SkipUntil returns a new sequence that skips elements until the predicate is true.

<a name="SkipWhile"></a>
## SkipWhile

```go
func SkipWhile[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E]
```

SkipWhile returns a new sequence that skips elements while the predicate is true.

<a name="Sort"></a>
## Sort

```go
func Sort[E types.Ordered](seq iter.Seq[E]) iter.Seq[E]
```

Sort sorts the elements of a sequence in ascending order.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5)

	sorted := seq.Sort(input)

	result := seq.Collect(sorted)
	fmt.Println(result)
}
```

**Output**

```
[1 1 2 3 3 4 5 5 5 6 9]
```


</details>

<a name="SortBy"></a>
## SortBy

```go
func SortBy[E any, K types.Ordered](seq iter.Seq[E], keyFn Mapper[E, K]) iter.Seq[E]
```

SortBy sorts the elements of a sequence in ascending order by the key returned by keyFn.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	input := seq.Of(
		Person{"Alice", 30},
		Person{"Bob", 25},
		Person{"Charlie", 35},
	)

	sorted := seq.SortBy(input, func(p Person) int {
		return p.Age
	})

	for p := range sorted {
		fmt.Printf("%s (%d)\n", p.Name, p.Age)
	}
}
```

**Output**

```
Bob (25)
Alice (30)
Charlie (35)
```


</details>

<a name="SortComparing"></a>
## SortComparing

```go
func SortComparing[E any](seq iter.Seq[E], cmp func(a, b E) int) iter.Seq[E]
```

SortComparing sorts the elements of a sequence in ascending order using the cmp function.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	input := seq.Of(
		Person{"Alice", 30},
		Person{"Bob", 25},
		Person{"Charlie", 35},
	)

	sorted := seq.SortComparing(input, func(a, b Person) int {
		return a.Age - b.Age
	})

	for p := range sorted {
		fmt.Printf("%s (%d)\n", p.Name, p.Age)
	}
}
```

**Output**

```
Bob (25)
Alice (30)
Charlie (35)
```


</details>

<a name="Take"></a>
## Take

```go
func Take[E any](seq iter.Seq[E], n int) iter.Seq[E]
```

Take returns a new sequence that contains only the first n elements of the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	taken := seq.Take(input, 3)

	result := seq.Collect(taken)

	fmt.Printf("%v\n", result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="TakeUntil"></a>
## TakeUntil

```go
func TakeUntil[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E]
```

TakeUntil returns a new sequence that contains elements until the predicate is true.

<a name="TakeWhile"></a>
## TakeWhile

```go
func TakeWhile[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E]
```

TakeWhile returns a new sequence that contains elements while the predicate is true.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 2, 1)

	taken := seq.TakeWhile(input, func(v int) bool {
		return v < 3
	})

	result := seq.Collect(taken)

	fmt.Printf("%v\n", result)
}
```

**Output**

```
[1 2]
```


</details>

<a name="Tap"></a>
## Tap

```go
func Tap[E any](seq iter.Seq[E], consumer func(E)) iter.Seq[E]
```

Tap returns a sequence that applies the given consumer to each element of the input sequence and pass it further.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.Tap(seq.Of(1, 2, 3), func(v int) {
		fmt.Println(v)
	})

	seq.Flush(sequence)

}
```

**Output**

```
1
2
3
```


</details>

<a name="Tick"></a>
## Tick

```go
func Tick(d time.Duration) iter.Seq[time.Time]
```

Tick returns a sequence that yields the current time every duration.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"time"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
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
```


</details>

<a name="ToSlice"></a>
## ToSlice

```go
func ToSlice[Slice ~[]E, E any](seq iter.Seq[E], slice Slice) Slice
```

ToSlice collects the elements of the given sequence into a slice.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.Of(1, 2, 3)

	slice := make([]int, 0, 3)
	result := seq.ToSlice(sequence, slice)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="Union"></a>
## Union

```go
func Union[E types.Comparable](seq1 iter.Seq[E], seq2 iter.Seq[E]) iter.Seq[E]
```

Union returns a sequence that contains all distinct elements from both input sequences.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	seq1 := seq.Of(1, 2, 3)
	seq2 := seq.Of(3, 4, 5)

	union := seq.Union(seq1, seq2)

	result := seq.Collect(union)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3 4 5]
```


</details>

<a name="UnionAll"></a>
## UnionAll

```go
func UnionAll[E any](seq1 iter.Seq[E], seq2 iter.Seq[E]) iter.Seq[E]
```

UnionAll returns a sequence that contains all elements from both input sequences.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	seq1 := seq.Of(1, 2, 3)
	seq2 := seq.Of(3, 4, 5)

	unionAll := seq.UnionAll(seq1, seq2)

	result := seq.Collect(unionAll)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3 3 4 5]
```


</details>

<a name="Uniq"></a>
## Uniq

```go
func Uniq[E comparable](seq iter.Seq[E]) iter.Seq[E]
```

Uniq returns a sequence with only unique elements.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 2, 3, 3, 3)

	unique := seq.Uniq(input)

	result := seq.Collect(unique)

	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="UniqBy"></a>
## UniqBy

```go
func UniqBy[E any, K comparable](seq iter.Seq[E], mapper Mapper[E, K]) iter.Seq[E]
```

UniqBy returns a sequence with only unique elements based on a key.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of("apple", "banana", "apricot", "blueberry")

	uniqueBy := seq.UniqBy(input, func(v string) string {
		return v[:1] // unique by first letter
	})

	result := seq.Collect(uniqueBy)

	fmt.Println(result)
}
```

**Output**

```
[apple banana]
```


</details>

<a name="Where"></a>
## Where

```go
func Where[E any](seq iter.Seq[E], predicate Predicate[E]) iter.Seq[E]
```

Where returns a new sequence that contains only the elements that satisfy the predicate. SQL\-like alias for Filter

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	input := seq.Of(1, 2, 3, 4, 5)

	filtered := seq.Where(input, func(v int) bool {
		return v%2 == 0
	})

	result := seq.Collect(filtered)

	fmt.Printf("%v\n", result)
}
```

**Output**

```
[2 4]
```


</details>

<a name="Zip"></a>
## Zip

```go
func Zip[E any, R any](seq1 iter.Seq[E], seq2 iter.Seq[R]) iter.Seq2[E, R]
```

Zip combines two sequences into a iter.Seq2.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	seq1 := seq.Of(1, 2, 3)
	seq2 := seq.Of("a", "b", "c")

	zipped := seq.Zip(seq1, seq2)

	for k, v := range zipped {
		fmt.Printf("%d: %s\n", k, v)
	}
}
```

**Output**

```
1: a
2: b
3: c
```


</details>

<a name="Consumer"></a>
## type Consumer

Consumer is a function that consumes an element of a sequence.

```go
type Consumer[E any] = func(E)
```

<a name="Mapper"></a>
## type Mapper

Mapper is a function that maps a value of type T to a value of type R.

```go
type Mapper[T any, R any] = func(T) R
```

<a name="Predicate"></a>
## type Predicate

Predicate is a function that takes an element and returns a boolean.

```go
type Predicate[E any] = Mapper[E, bool]
```

<a name="Sequence"></a>
## type Sequence

Sequence is a monad representing a sequence of elements.

```go
type Sequence[E comparable] struct {
    // contains filtered or unexported fields
}
```

<a name="AsSequence"></a>
### AsSequence

```go
func AsSequence[E comparable](seq iter.Seq[E]) Sequence[E]
```

AsSequence wraps an iter.Seq to provide a possibility to pipe several method calls.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	result := sequence.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="ConcatSequences"></a>
### ConcatSequences

```go
func ConcatSequences[E comparable](sequences ...Sequence[E]) Sequence[E]
```

ConcatSequences concatenates multiple sequences into a single sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	seq1 := seq.AsSequence(seq.Of(1, 2))
	seq2 := seq.AsSequence(seq.Of(3, 4))
	concatenated := seq.ConcatSequences(seq1, seq2)
	result := concatenated.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2 3 4]
```


</details>

<a name="Sequence[E].Append"></a>
### Sequence[E].Append

```go
func (s Sequence[E]) Append(elems ...E) Sequence[E]
```

Append appends elements to the end of a sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	appended := sequence.Append(4, 5)
	result := appended.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2 3 4 5]
```


</details>

<a name="Sequence[E].Collect"></a>
### Sequence[E].Collect

```go
func (s Sequence[E]) Collect() []E
```

Collect collects the elements of the sequence into a slice.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	result := sequence.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="Sequence[E].Contains"></a>
### Sequence[E].Contains

```go
func (s Sequence[E]) Contains(elem E) bool
```

Contains returns true if the element is in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	contains := sequence.Contains(3)
	fmt.Println(contains)
}
```

**Output**

```
true
```


</details>

<a name="Sequence[E].ContainsAll"></a>
### Sequence[E].ContainsAll

```go
func (s Sequence[E]) ContainsAll(elements ...E) bool
```

ContainsAll returns true if all elements are in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	containsAll := sequence.ContainsAll(2, 3, 4)
	fmt.Println(containsAll)
}
```

**Output**

```
true
```


</details>

<a name="Sequence[E].Count"></a>
### Sequence[E].Count

```go
func (s Sequence[E]) Count() int
```

Count returns the number of elements in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	count := sequence.Count()
	fmt.Println(count)
}
```

**Output**

```
3
```


</details>

<a name="Sequence[E].Distinct"></a>
### Sequence[E].Distinct

```go
func (s Sequence[E]) Distinct() Sequence[E]
```

Distinct returns a new sequence with only unique elements. SQL\-like alias for Uniq

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 2, 3, 3, 3))
	distinct := sequence.Distinct()
	result := distinct.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="Sequence[E].Each"></a>
### Sequence[E].Each

```go
func (s Sequence[E]) Each(consumer Consumer[E]) Sequence[E]
```

Each returns a new sequence that calls the consumer for each element of the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3)).Each(func(v int) {
		fmt.Println(v)
	})
	sequence.Flush()
}
```

**Output**

```
1
2
3
```


</details>

<a name="Sequence[E].Every"></a>
### Sequence[E].Every

```go
func (s Sequence[E]) Every(predicate Predicate[E]) bool
```

Every returns true if all elements satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	every := sequence.Every(func(v int) bool {
		return v > 0
	})
	fmt.Println(every)
}
```

**Output**

```
true
```


</details>

<a name="Sequence[E].Exists"></a>
### Sequence[E].Exists

```go
func (s Sequence[E]) Exists(predicate Predicate[E]) bool
```

Exists returns true if there is at least one element that satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	exists := sequence.Exists(func(v int) bool {
		return v > 3
	})
	fmt.Println(exists)
}
```

**Output**

```
true
```


</details>

<a name="Sequence[E].Filter"></a>
### Sequence[E].Filter

```go
func (s Sequence[E]) Filter(predicate Predicate[E]) Sequence[E]
```

Filter returns a new sequence with elements that satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	filtered := sequence.Filter(func(v int) bool {
		return v%2 == 0
	})
	result := filtered.Collect()
	fmt.Println(result)
}
```

**Output**

```
[2 4]
```


</details>

<a name="Sequence[E].Find"></a>
### Sequence[E].Find

```go
func (s Sequence[E]) Find(predicate Predicate[E]) optional.Elem[E]
```

Find returns the first element that satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	found := sequence.Find(func(v int) bool {
		return v > 3
	})

	fmt.Println(found.MustGet())
}
```

**Output**

```
4
```


</details>

<a name="Sequence[E].FindAll"></a>
### Sequence[E].FindAll

```go
func (s Sequence[E]) FindAll(predicate Predicate[E]) Sequence[E]
```

FindAll returns all elements that satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	foundAll := sequence.FindAll(func(v int) bool {
		return v > 3
	})
	result := foundAll.Collect()
	fmt.Println(result)
}
```

**Output**

```
[4 5]
```


</details>

<a name="Sequence[E].FindLast"></a>
### Sequence[E].FindLast

```go
func (s Sequence[E]) FindLast(predicate Predicate[E]) optional.Elem[E]
```

FindLast returns the last element that satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	found := sequence.FindLast(func(v int) bool {
		return v > 3
	})

	fmt.Println(found.MustGet())
}
```

**Output**

```
5
```


</details>

<a name="Sequence[E].Flush"></a>
### Sequence[E].Flush

```go
func (s Sequence[E]) Flush()
```

Flush consumes all elements of the input sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	sequence.Flush()
	// No output expected
}
```


</details>

<a name="Sequence[E].Fold"></a>
### Sequence[E].Fold

```go
func (s Sequence[E]) Fold(accumulator func(agg E, item E) E) optional.Elem[E]
```

Fold applies a function against an accumulator and each element in the sequence \(from left to right\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	sum := sequence.Fold(func(agg, item int) int {
		return agg + item
	})

	fmt.Println(sum.MustGet())
}
```

**Output**

```
15
```


</details>

<a name="Sequence[E].FoldRight"></a>
### Sequence[E].FoldRight

```go
func (s Sequence[E]) FoldRight(accumulator func(agg E, item E) E) optional.Elem[E]
```

FoldRight applies a function against an accumulator and each element in the sequence \(from right to left\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of("a", "b", "c"))
	concat := sequence.FoldRight(func(agg, item string) string {
		return agg + item
	})

	fmt.Println(concat.MustGet())
}
```

**Output**

```
cba
```


</details>

<a name="Sequence[E].ForEach"></a>
### Sequence[E].ForEach

```go
func (s Sequence[E]) ForEach(consumer Consumer[E])
```

ForEach calls the consumer for each element of the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	seq.AsSequence(seq.Of(1, 2, 3)).ForEach(func(v int) {
		fmt.Println(v)
	})
}
```

**Output**

```
1
2
3
```


</details>

<a name="Sequence[E].IsEmpty"></a>
### Sequence[E].IsEmpty

```go
func (s Sequence[E]) IsEmpty() bool
```

IsEmpty returns true if the sequence is empty.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of[int]())
	isEmpty := sequence.IsEmpty()
	fmt.Println(isEmpty)
}
```

**Output**

```
true
```


</details>

<a name="Sequence[E].IsNotEmpty"></a>
### Sequence[E].IsNotEmpty

```go
func (s Sequence[E]) IsNotEmpty() bool
```

IsNotEmpty returns true if the sequence is not empty.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	isNotEmpty := sequence.IsNotEmpty()
	fmt.Println(isNotEmpty)
}
```

**Output**

```
true
```


</details>

<a name="Sequence[E].Limit"></a>
### Sequence[E].Limit

```go
func (s Sequence[E]) Limit(n int) Sequence[E]
```

Limit returns a new sequence that contains only the first n elements of the given sequence. SQL\-like alias for Take

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	limited := sequence.Limit(2)
	result := limited.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2]
```


</details>

<a name="Sequence[E].None"></a>
### Sequence[E].None

```go
func (s Sequence[E]) None(predicate Predicate[E]) bool
```

None returns true if no element satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	none := sequence.None(func(v int) bool {
		return v > 5
	})
	fmt.Println(none)
}
```

**Output**

```
true
```


</details>

<a name="Sequence[E].NotContains"></a>
### Sequence[E].NotContains

```go
func (s Sequence[E]) NotContains(elem E) bool
```

NotContains returns true if the element is not in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	notContains := sequence.NotContains(6)
	fmt.Println(notContains)
}
```

**Output**

```
true
```


</details>

<a name="Sequence[E].Offset"></a>
### Sequence[E].Offset

```go
func (s Sequence[E]) Offset(n int) Sequence[E]
```

Offset returns a new sequence that skips the first n elements of the given sequence. SQL\-like alias for Skip

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	offset := sequence.Offset(2)
	result := offset.Collect()
	fmt.Println(result)
}
```

**Output**

```
[3 4 5]
```


</details>

<a name="Sequence[E].Partition"></a>
### Sequence[E].Partition

```go
func (s Sequence[E]) Partition(size int) iter.Seq[iter.Seq[E]]
```

Partition splits the sequence into chunks of the given size.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	partitions := sequence.Partition(2)
	for partition := range partitions {
		fmt.Println(seq.Collect(partition))
	}
}
```

**Output**

```
[1 2]
[3 4]
[5]
```


</details>

<a name="Sequence[E].Prepend"></a>
### Sequence[E].Prepend

```go
func (s Sequence[E]) Prepend(elems ...E) Sequence[E]
```

Prepend prepends elements to the beginning of a sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(3, 4, 5))
	prepended := sequence.Prepend(1, 2)
	result := prepended.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2 3 4 5]
```


</details>

<a name="Sequence[E].Repeat"></a>
### Sequence[E].Repeat

```go
func (s Sequence[E]) Repeat(count int) Sequence[E]
```

Repeat repeats the sequence count times.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	repeated := sequence.Repeat(2)
	result := repeated.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2 3 1 2 3]
```


</details>

<a name="Sequence[E].Reverse"></a>
### Sequence[E].Reverse

```go
func (s Sequence[E]) Reverse() Sequence[E]
```

Reverse returns a new sequence with elements in reverse order.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	reversed := sequence.Reverse()
	result := reversed.Collect()
	fmt.Println(result)
}
```

**Output**

```
[3 2 1]
```


</details>

<a name="Sequence[E].Skip"></a>
### Sequence[E].Skip

```go
func (s Sequence[E]) Skip(n int) Sequence[E]
```

Skip returns a new sequence that skips the first n elements of the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	skipped := sequence.Skip(2)
	result := skipped.Collect()
	fmt.Println(result)
}
```

**Output**

```
[3 4 5]
```


</details>

<a name="Sequence[E].Take"></a>
### Sequence[E].Take

```go
func (s Sequence[E]) Take(n int) Sequence[E]
```

Take returns a new sequence that contains only the first n elements of the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	taken := sequence.Take(3)
	result := taken.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="Sequence[E].Tap"></a>
### Sequence[E].Tap

```go
func (s Sequence[E]) Tap(consumer func(E)) Sequence[E]
```

Tap returns a new sequence that calls the consumer for each element of the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3)).Tap(func(v int) {
		fmt.Println(v)
	})
	sequence.Flush()
}
```

**Output**

```
1
2
3
```


</details>

<a name="Sequence[E].ToSlice"></a>
### Sequence[E].ToSlice

```go
func (s Sequence[E]) ToSlice(slice []E) []E
```

ToSlice collects the elements of the sequence into a given slice.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3))
	var slice []int
	result := sequence.ToSlice(slice)
	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="Sequence[E].Union"></a>
### Sequence[E].Union

```go
func (s Sequence[E]) Union(other Sequence[E]) Sequence[E]
```

Union returns a new sequence that contains all distinct elements from both input sequences.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	seq1 := seq.AsSequence(seq.Of(1, 2, 3))
	seq2 := seq.AsSequence(seq.Of(3, 4, 5))
	union := seq1.Union(seq2)
	result := union.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2 3 4 5]
```


</details>

<a name="Sequence[E].UnionAll"></a>
### Sequence[E].UnionAll

```go
func (s Sequence[E]) UnionAll(other Sequence[E]) Sequence[E]
```

UnionAll returns a new sequence that contains all elements from both input sequences.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	seq1 := seq.AsSequence(seq.Of(1, 2, 3))
	seq2 := seq.AsSequence(seq.Of(3, 4, 5))
	unionAll := seq1.UnionAll(seq2)
	result := unionAll.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2 3 3 4 5]
```


</details>

<a name="Sequence[E].Uniq"></a>
### Sequence[E].Uniq

```go
func (s Sequence[E]) Uniq() Sequence[E]
```

Uniq returns a new sequence with only unique elements.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 2, 3, 3, 3))
	unique := sequence.Uniq()
	result := unique.Collect()
	fmt.Println(result)
}
```

**Output**

```
[1 2 3]
```


</details>

<a name="Sequence[E].Where"></a>
### Sequence[E].Where

```go
func (s Sequence[E]) Where(predicate Predicate[E]) Sequence[E]
```

Where returns a new sequence with elements that satisfy the predicate. SQL\-like alias for Filter

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	sequence := seq.AsSequence(seq.Of(1, 2, 3, 4, 5))
	filtered := sequence.Where(func(v int) bool {
		return v%2 == 0
	})
	result := filtered.Collect()
	fmt.Println(result)
}
```

**Output**

```
[2 4]
```


</details>
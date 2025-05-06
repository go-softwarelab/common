# seq2

```go
import "github.com/go-softwarelab/common/pkg/seq2"
```

Package seq2 provides a comprehensive set of utilities for working with key\-value sequences in Go applications.

The goal of this package is to offer a rich set of functions for creating, transforming, and consuming iter.Seq2, enabling developers to work with collections of key\-value pairs in a functional programming style. The package includes utilities for filtering, mapping, reducing, and sorting sequences, as well as combining and partitioning them.

The package is designed to reduce boilerplate code and improve readability by providing a consistent API for common sequence operations. It leverages Go's type safety and generics to ensure that operations on sequences are both flexible and safe. The Sequence struct is worth mentioning explicitly, allowing method chaining and fluent composition of sequence operations.



<a name="Append"></a>
## [Append](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/joins.go#L50>)

```go
func Append[K any, V any](seq iter.Seq2[K, V], key K, value V) iter.Seq2[K, V]
```

Append appends element to the end of a sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2})
	input = seq2.SortByKeys(input)

	// Append a new key-value pair to the sequence
	appended := seq2.Append(input, "c", 3)

	result := seq2.CollectToMap(appended)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="Collect"></a>
## [Collect](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/consumer.go#L58>)

```go
func Collect[K comparable, V any](seq iter.Seq2[K, V]) []types.Pair[K, V]
```

Collect collects the elements of the given sequence into a slice of types.Pair of K and V.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	result := seq2.Collect(input)

	fmt.Println(result)

}
```

**Output**

```
[{a 1} {b 2} {c 3}]
```


</details>

<a name="CollectToMap"></a>
## [CollectToMap](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/consumer.go#L53>)

```go
func CollectToMap[K comparable, V any](seq iter.Seq2[K, V]) map[K]V
```

CollectToMap collects the elements of the given sequence into a map.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	result := seq2.CollectToMap(input)

	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="Concat"></a>
## [Concat](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/joins.go#L6>)

```go
func Concat[K any, V any](sequences ...iter.Seq2[K, V]) iter.Seq2[K, V]
```

Concat concatenates multiple sequences into a single sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	first := seq2.FromMap(map[string]int{"a": 1, "b": 2})
	first = seq2.SortByKeys(first)
	second := seq2.FromMap(map[string]int{"c": 3, "d": 4})
	second = seq2.SortByKeys(second)

	// Concatenate two sequences
	combined := seq2.Concat(first, second)

	result := seq2.CollectToMap(combined)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3 d:4]
```


</details>

<a name="Contains"></a>
## [Contains](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L33>)

```go
func Contains[K comparable, V any](seq iter.Seq2[K, V], key K) bool
```

Contains returns true if the key is in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	contains := seq2.Contains(input, "b")

	fmt.Println(contains)
}
```

**Output**

```
true
```


</details>

<a name="ContainsAll"></a>
## [ContainsAll](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L58>)

```go
func ContainsAll[K comparable, V any](seq iter.Seq2[K, V], keys ...K) bool
```

ContainsAll returns true if all keys are in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	containsAll := seq2.ContainsAll(input, "a", "c", "d")

	fmt.Println(containsAll)
}
```

**Output**

```
true
```


</details>

<a name="ContainsAllValues"></a>
## [ContainsAllValues](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L73>)

```go
func ContainsAllValues[K any, V comparable](seq iter.Seq2[K, V], values ...V) bool
```

ContainsAllValues returns true if all values are in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	containsAllValues := seq2.ContainsAllValues(input, 1, 3)

	fmt.Println(containsAllValues)
}
```

**Output**

```
true
```


</details>

<a name="ContainsPair"></a>
## [ContainsPair](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L83>)

```go
func ContainsPair[K comparable, V comparable](seq iter.Seq2[K, V], key K, value V) bool
```

ContainsPair returns true if the key\-value pair is in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	containsPair := seq2.ContainsPair(input, "b", 2)

	fmt.Println(containsPair)
}
```

**Output**

```
true
```


</details>

<a name="ContainsValue"></a>
## [ContainsValue](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L48>)

```go
func ContainsValue[K any, V comparable](seq iter.Seq2[K, V], value V) bool
```

ContainsValue returns true if the value is in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	containsValue := seq2.ContainsValue(input, 3)

	fmt.Println(containsValue)
}
```

**Output**

```
true
```


</details>

<a name="Count"></a>
## [Count](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/consumer.go#L66>)

```go
func Count[K any, V any](seq iter.Seq2[K, V]) int
```

Count returns the number of elements in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Count returns the number of elements in the sequence
	count := seq2.Count(input)

	fmt.Println(count)
}
```

**Output**

```
3
```


</details>

<a name="Cycle"></a>
## [Cycle](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/mapper.go#L56>)

```go
func Cycle[K, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

Cycle repeats the sequence indefinitely.

<a name="CycleTimes"></a>
## [CycleTimes](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/mapper.go#L69>)

```go
func CycleTimes[K, V any](seq iter.Seq2[K, V], count int) iter.Seq2[K, V]
```

CycleTimes repeats the sequence count times.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2})
	input = seq2.SortByKeys(input)

	// Repeat the sequence 2 times
	cycled := seq2.CycleTimes(input, 2)

	seq2.ForEach(cycled, func(k string, v int) {
		fmt.Println(k, v)
	})

}
```

**Output**

```
a 1
b 2
a 1
b 2
```


</details>

<a name="Distinct"></a>
## [Distinct](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L258>)

```go
func Distinct[K comparable, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

Distinct returns a new sequence that contains only the unique elements of the given sequence. SQL\-like alias for Uniq.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create a sequence with duplicate key-value pairs
	input := seq2.Concat(
		seq2.FromMap(map[string]int{"a": 1, "b": 2}),
		seq2.FromMap(map[string]int{"a": 1, "c": 3}),
	)

	// Distinct is an alias for Uniq
	unique := seq2.Distinct(input)

	result := seq2.CollectToMap(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="DistinctKeys"></a>
## [DistinctKeys](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L263>)

```go
func DistinctKeys[K comparable, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

DistinctKeys returns a new sequence that contains only the unique keys of the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create a sequence with duplicate keys
	input := seq2.Concat(
		seq2.FromMap(map[string]int{"a": 1, "b": 2}),
		seq2.FromMap(map[string]int{"a": 3, "c": 4}),
	)

	// DistinctKeys is an alias for UniqKeys
	unique := seq2.DistinctKeys(input)

	result := seq2.CollectToMap(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:4]
```


</details>

<a name="Each"></a>
## [Each](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/consumer.go#L29>)

```go
func Each[K any, V any](seq iter.Seq2[K, V], consumer Consumer[K, V]) iter.Seq2[K, V]
```

Each returns a sequence that applies the given consumer to each element of the input sequence and pass it further. Each is an alias for Tap. Comparing to ForEach, this is a lazy function and doesn't consume the input sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	tapped := seq2.Each(input, func(k string, v int) {
		fmt.Printf("Each: %s -> %d\n", k, v)
	})

	seq2.Flush(tapped)

}
```

**Output**

```
Each: a -> 1
Each: b -> 2
Each: c -> 3
```


</details>

<a name="Empty"></a>
## [Empty](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L12>)

```go
func Empty[K, V any]() iter.Seq2[K, V]
```

Empty returns an empty sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create an empty sequence
	empty := seq2.Empty[any, any]()

	seq2.ForEach(empty, func(any, any) {
		fmt.Println("Should not be called")
	})
}
```

**Output**

```

```


</details>

<a name="Every"></a>
## [Every](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L108>)

```go
func Every[K, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) bool
```

Every returns true if all elements satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 2, "b": 4, "c": 6, "d": 8})
	input = seq2.SortByKeys(input)

	every := seq2.Every(input, func(k string, v int) bool {
		return v%2 == 0
	})

	fmt.Println(every)
}
```

**Output**

```
true
```


</details>

<a name="Exists"></a>
## [Exists](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L98>)

```go
func Exists[K, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) bool
```

Exists returns true if there is at least one element that satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	exists := seq2.Exists(input, func(k string, v int) bool {
		return k > "c" && v > 3
	})

	fmt.Println(exists)
}
```

**Output**

```
true
```


</details>

<a name="Filter"></a>
## [Filter](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L15>)

```go
func Filter[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V]
```

Filter returns a new sequence that contains only the elements that satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	// Filter elements where the value is even
	filtered := seq2.Filter(input, func(k string, v int) bool {
		return v%2 == 0
	})

	result := seq2.CollectToMap(filtered)
	fmt.Println(result)
}
```

**Output**

```
map[b:2 d:4]
```


</details>

<a name="FilterByKey"></a>
## [FilterByKey](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L34>)

```go
func FilterByKey[K any, V any](seq iter.Seq2[K, V], predicate KeyPredicate[K]) iter.Seq2[K, V]
```

FilterByKey returns a new sequence that contains only the elements that satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	// Filter elements by key
	filtered := seq2.FilterByKey(input, func(k string) bool {
		return k > "b"
	})

	result := seq2.CollectToMap(filtered)
	fmt.Println(result)
}
```

**Output**

```
map[c:3 d:4]
```


</details>

<a name="FilterByValue"></a>
## [FilterByValue](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L47>)

```go
func FilterByValue[K any, V any](seq iter.Seq2[K, V], predicate ValuePredicate[V]) iter.Seq2[K, V]
```

FilterByValue returns a new sequence that contains only the elements that satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	// Filter elements by value
	filtered := seq2.FilterByValue(input, func(v int) bool {
		return v <= 2
	})

	result := seq2.CollectToMap(filtered)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2]
```


</details>

<a name="FindAll"></a>
## [FindAll](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L10>)

```go
func FindAll[K, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V]
```

FindAll returns all elements that satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	found := seq2.FindAll(input, func(k string, v int) bool {
		return v > 2
	})

	result := seq2.CollectToMap(found)
	fmt.Println(result)
}
```

**Output**

```
map[c:3 d:4]
```


</details>

<a name="Flush"></a>
## [Flush](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/consumer.go#L42>)

```go
func Flush[K any, V any](seq iter.Seq2[K, V])
```

Flush consumes all elements of the input sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create a sequence that has side effects when consumed
	sideEffects := iter.Seq2[string, int](func(yield func(string, int) bool) {
		fmt.Println("First element consumed")
		if !yield("a", 1) {
			return
		}
		fmt.Println("Second element consumed")
		if !yield("b", 2) {
			return
		}
		fmt.Println("Third element consumed")
		if !yield("c", 3) {
			return
		}
	})

	// Flush consumes all elements without doing anything with them
	seq2.Flush(sideEffects)

}
```

**Output**

```
First element consumed
Second element consumed
Third element consumed
```


</details>

<a name="ForEach"></a>
## [ForEach](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/consumer.go#L35>)

```go
func ForEach[K any, V any](seq iter.Seq2[K, V], consumer Consumer[K, V])
```

ForEach applies consumer to each element of the input sequence. Comparing to Each, this is not a lazy function and consumes the input sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.OfIndexed("a", "b", "c")

	// ForEach consumes the sequence and applies the function
	seq2.ForEach(input, func(k int, v string) {
		fmt.Printf("%d: %s\n", k, v)
	})

}
```

**Output**

```
0: a
1: b
2: c
```


</details>

<a name="FromMap"></a>
## [FromMap](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L56>)

```go
func FromMap[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V]
```

FromMap creates a new iter.Seq2 from the given map.

<a name="FromSlice"></a>
## [FromSlice](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L51>)

```go
func FromSlice[Slice ~[]E, E any](slice Slice) iter.Seq2[int, E]
```

FromSlice creates a new sequence from the given slice with index as keys.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create a sequence from a slice
	slice := []string{"a", "b", "c"}
	sequence := seq2.FromSlice(slice)

	result := seq2.CollectToMap(sequence)
	fmt.Println(result)
}
```

**Output**

```
map[0:a 1:b 2:c]
```


</details>

<a name="Get"></a>
## [Get](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L25>)

```go
func Get[K comparable, V any](seq iter.Seq2[K, V], key K) optional.Value[V]
```

Get returns the element at the specified key.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	value := seq2.Get(input, "c")

	fmt.Println(value.MustGet())
}
```

**Output**

```
3
```


</details>

<a name="IsEmpty"></a>
## [IsEmpty](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/seq2.go#L6>)

```go
func IsEmpty[K, V any](seq iter.Seq2[K, V]) bool
```

IsEmpty returns true if the sequence is empty.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	emptySeq := seq2.Empty[any, any]()

	nonEmptySeq := seq2.OfIndexed("a")

	fmt.Printf("Empty sequence: %v\n", seq2.IsEmpty(emptySeq))
	fmt.Printf("Non-empty sequence: %v\n", seq2.IsEmpty(nonEmptySeq))
}
```

**Output**

```
Empty sequence: true
Non-empty sequence: false
```


</details>

<a name="IsNotEmpty"></a>
## [IsNotEmpty](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/seq2.go#L14>)

```go
func IsNotEmpty[K, V any](seq iter.Seq2[K, V]) bool
```

IsNotEmpty returns true if the sequence is not empty.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	emptySeq := seq2.Empty[any, any]()

	nonEmptySeq := seq2.OfIndexed("a")

	fmt.Printf("Empty sequence: %v\n", seq2.IsNotEmpty(emptySeq))
	fmt.Printf("Non-empty sequence: %v\n", seq2.IsNotEmpty(nonEmptySeq))
}
```

**Output**

```
Empty sequence: false
Non-empty sequence: true
```


</details>

<a name="Keys"></a>
## [Keys](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L17>)

```go
func Keys[K, V any](seq iter.Seq2[K, V]) iter.Seq[K]
```

Keys returns a sequence of keys from a sequence of key\-value pairs.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.OfIndexed("a", "b", "c")

	keys := seq2.Keys(input)

	seq.ForEach(keys, func(k int) {
		fmt.Print(k, " ")
	})
}
```

**Output**

```
0 1 2
```


</details>

<a name="Limit"></a>
## [Limit](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L152>)

```go
func Limit[K any, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V]
```

Limit returns a new sequence that contains only the first n elements of the given sequence. SQL\-like alias for Take.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	// Limit is an alias for Take
	taken := seq2.Limit(input, 3)

	result := seq2.CollectToMap(taken)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="Map"></a>
## [Map](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/mapper.go#L18>)

```go
func Map[K, V, RK, RV any](seq iter.Seq2[K, V], mapper DoubleMapper[K, V, RK, RV]) iter.Seq2[RK, RV]
```

Map applies a mapper function to each element of the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Map both key and value to produce a new value (keeps original keys)
	mapped := seq2.Map(input, func(k string, v int) (string, int) {
		return strings.ToUpper(k), v * 10
	})

	result := seq2.CollectToMap(mapped)
	fmt.Println(result)
}
```

**Output**

```
map[A:10 B:20 C:30]
```


</details>

<a name="MapKeys"></a>
## [MapKeys](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/mapper.go#L30>)

```go
func MapKeys[K, V, RK any](seq iter.Seq2[K, V], mapper KeyMapper[K, RK]) iter.Seq2[RK, V]
```

MapKeys applies a mapper function to each key of the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Map keys to uppercase (keeps original values)
	mapped := seq2.MapKeys(input, func(k string) string {
		return strings.ToUpper(k)
	})

	result := seq2.CollectToMap(mapped)
	fmt.Println(result)
}
```

**Output**

```
map[A:1 B:2 C:3]
```


</details>

<a name="MapTo"></a>
## [MapTo](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/mapper.go#L44>)

```go
func MapTo[K, V, RV any](seq iter.Seq2[K, V], mapper Mapper[K, V, RV]) iter.Seq[RV]
```

MapTo applies a mapper function to each element of the sequence and returns a sequence of mapper results.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Map each key-value pair to a string
	mapped := seq2.MapTo(input, func(k string, v int) string {
		return fmt.Sprintf("%s=%d", k, v)
	})

	seq.ForEach(mapped, func(v string) {
		fmt.Println(v)
	})
}
```

**Output**

```
a=1
b=2
c=3
```


</details>

<a name="MapValues"></a>
## [MapValues](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/mapper.go#L37>)

```go
func MapValues[K, V, RV any](seq iter.Seq2[K, V], mapper ValueMapper[V, RV]) iter.Seq2[K, RV]
```

MapValues applies a mapper function to each value of the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Map values to their squares (keeps original keys)
	mapped := seq2.MapValues(input, func(v int) int {
		return v * v
	})

	result := seq2.CollectToMap(mapped)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:4 c:9]
```


</details>

<a name="None"></a>
## [None](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L118>)

```go
func None[K, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) bool
```

None returns true if no element satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	none := seq2.None(input, func(k string, v int) bool {
		return v > 10
	})

	fmt.Println(none)
}
```

**Output**

```
true
```


</details>

<a name="NotContains"></a>
## [NotContains](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L43>)

```go
func NotContains[K comparable, V any](seq iter.Seq2[K, V], key K) bool
```

NotContains returns true if the key is not in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	notContains := seq2.NotContains(input, "x")

	fmt.Println(notContains)
}
```

**Output**

```
true
```


</details>

<a name="NotContainsPair"></a>
## [NotContainsPair](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L93>)

```go
func NotContainsPair[K comparable, V comparable](seq iter.Seq2[K, V], key K, value V) bool
```

NotContainsPair returns true if the key\-value pair is not in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	notContainsPair := seq2.NotContainsPair(input, "b", 3)

	fmt.Println(notContainsPair)
}
```

**Output**

```
true
```


</details>

<a name="NotContainsValue"></a>
## [NotContainsValue](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/find.go#L68>)

```go
func NotContainsValue[K any, V comparable](seq iter.Seq2[K, V], value V) bool
```

NotContainsValue returns true if the value is not in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	notContainsValue := seq2.NotContainsValue(input, 5)

	fmt.Println(notContainsValue)
}
```

**Output**

```
true
```


</details>

<a name="OfIndexed"></a>
## [OfIndexed](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L46>)

```go
func OfIndexed[E any](elems ...E) iter.Seq2[int, E]
```

OfIndexed creates a new indexed sequence from the given elements.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create a sequence from individual elements
	indexed := seq2.OfIndexed(1, 2, 3)

	result := seq2.CollectToMap(indexed)
	fmt.Println(result)
}
```

**Output**

```
map[0:1 1:2 2:3]
```


</details>

<a name="Offset"></a>
## [Offset](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L108>)

```go
func Offset[K any, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V]
```

Offset returns a new sequence that skips the first n elements of the given sequence. SQL\-like alias for Skip.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.OfIndexed(10, 20, 30, 40, 20)

	// Skip the first 2 elements
	skipped := seq2.Offset(input, 2)

	result := seq2.CollectToMap(skipped)
	fmt.Println(result)
}
```

**Output**

```
map[2:30 3:40 4:20]
```


</details>

<a name="Prepend"></a>
## [Prepend](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/joins.go#L55>)

```go
func Prepend[K any, V any](seq iter.Seq2[K, V], key K, value V) iter.Seq2[K, V]
```

Prepend prepends element to the beginning of a sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Prepend a new key-value pair to the sequence
	prepended := seq2.Prepend(input, "a", 1)

	result := seq2.CollectToMap(prepended)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="Reduce"></a>
## [Reduce](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/reducers.go#L8>)

```go
func Reduce[K any, V any, R any](seq2 iter.Seq2[K, V], accumulator func(agg R, key K, value V) R, initial R) R
```

Reduce applies a function against an accumulator and each element in the sequence \(from left to right\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	type Reduced struct {
		Key   string
		Value int
	}

	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Reduce to calculate the sum of all values
	reduced := seq2.Reduce(input, func(agg Reduced, key string, value int) Reduced {
		return Reduced{
			Key:   agg.Key + key,
			Value: agg.Value + value,
		}
	}, Reduced{})

	fmt.Println(reduced)
}
```

**Output**

```
{abc 6}
```


</details>

<a name="ReduceRight"></a>
## [ReduceRight](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/reducers.go#L17>)

```go
func ReduceRight[K any, V any, R any](seq2 iter.Seq2[K, V], accumulator func(agg R, key K, value V) R, initial R) R
```

ReduceRight applies a function against an accumulator and each element in the sequence \(from right to left\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	type Reduced struct {
		Key   string
		Value int
	}

	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	// Reduce to calculate the sum of all values
	reduced := seq2.ReduceRight(input, func(agg Reduced, key string, value int) Reduced {
		return Reduced{
			Key:   agg.Key + key,
			Value: agg.Value + value,
		}
	}, Reduced{})

	fmt.Println(reduced)
}
```

**Output**

```
{cba 6}
```


</details>

<a name="Repeat"></a>
## [Repeat](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L85>)

```go
func Repeat[K any, V any, N types.Integer](key K, value V, count N) iter.Seq2[K, V]
```

Repeat repeats the given pair \`count\` times.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create a sequence that repeats a key-value pair 3 times
	repeated := seq2.Repeat("key", 42, 3)

	seq2.ForEach(repeated, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
}
```

**Output**

```
key : 42
key : 42
key : 42
```


</details>

<a name="Reverse"></a>
## [Reverse](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L118>)

```go
func Reverse[K, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

Reverse reverses the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create an indexed sequence
	sequence := seq2.OfIndexed("a", "b", "c")

	// Reverse it
	reversed := seq2.Reverse(sequence)

	// CollectToMap into pairs for ordered display
	var pairs []string
	seq2.ForEach(reversed, func(k int, v string) {
		fmt.Println(k, ":", v)
	})

	fmt.Println(strings.Join(pairs, ", "))
}
```

**Output**

```
2 : c
1 : b
0 : a
```


</details>

<a name="Single"></a>
## [Single](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L39>)

```go
func Single[K, V any](k K, v V) iter.Seq2[K, V]
```

Single returns a sequence with given key value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create a sequence with a single key-value pair
	single := seq2.Single("key", 42)

	seq2.ForEach(single, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
}
```

**Output**

```
key : 42
```


</details>

<a name="Skip"></a>
## [Skip](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L60>)

```go
func Skip[K any, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V]
```

Skip returns a new sequence that skips the first n elements of the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.OfIndexed(10, 20, 30, 40, 20)

	// Skip the first 2 elements
	skipped := seq2.Skip(input, 2)

	result := seq2.CollectToMap(skipped)
	fmt.Println(result)
}
```

**Output**

```
map[2:30 3:40 4:20]
```


</details>

<a name="SkipUntil"></a>
## [SkipUntil](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L91>)

```go
func SkipUntil[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V]
```

SkipUntil returns a new sequence that skips elements from the given sequence until the predicate is satisfied.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.OfIndexed(10, 20, 30, 40, 20)

	// Skip elements until value of 30 is reached
	skipped := seq2.SkipUntil(input, func(k int, v int) bool {
		return v == 30
	})

	result := seq2.CollectToMap(skipped)
	fmt.Println(result)
}
```

**Output**

```
map[2:30 3:40 4:20]
```


</details>

<a name="SkipWhile"></a>
## [SkipWhile](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L75>)

```go
func SkipWhile[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V]
```

SkipWhile returns a new sequence that skips elements from the given sequence while the predicate is satisfied.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.OfIndexed(10, 20, 30, 40, 20)

	// Skip elements until the value is less than 30
	skipped := seq2.SkipWhile(input, func(k int, v int) bool {
		return v < 30
	})

	result := seq2.CollectToMap(skipped)
	fmt.Println(result)
}
```

**Output**

```
map[2:30 3:40 4:20]
```


</details>

<a name="SortBy"></a>
## [SortBy](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/sort.go#L19>)

```go
func SortBy[K any, V any, R types.Ordered](seq iter.Seq2[K, V], mapper func(K, V) R) iter.Seq2[K, V]
```

SortBy sorts the elements of a sequence by result of the mapper.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.Single("a", 1)
	input = seq2.Append(input, "c", 3)
	input = seq2.Append(input, "b", 2)

	// Sort by values
	sorted := seq2.SortBy(input, func(k string, v int) int {
		return v
	})

	seq2.ForEach(sorted, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
}
```

**Output**

```
a : 1
b : 2
c : 3
```


</details>

<a name="SortByKeys"></a>
## [SortByKeys](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/sort.go#L12>)

```go
func SortByKeys[K types.Ordered, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

SortByKeys sorts the elements of a sequence by key in ascending order.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.Single("a", 1)
	input = seq2.Append(input, "c", 3)
	input = seq2.Append(input, "b", 2)

	// SortByKeys by keys (alphabetically)
	sorted := seq2.SortByKeys(input)

	seq2.ForEach(sorted, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
}
```

**Output**

```
a : 1
b : 2
c : 3
```


</details>

<a name="SortComparingKeys"></a>
## [SortComparingKeys](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/sort.go#L38>)

```go
func SortComparingKeys[K any, V any](seq iter.Seq2[K, V], cmp func(K, K) int) iter.Seq2[K, V]
```

SortComparingKeys sorts the elements of a sequence by key in ascending order.

<details>
<summary>Example</summary>




```go
package main

import (
	"cmp"
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.Single("a", 1)
	input = seq2.Append(input, "c", 3)
	input = seq2.Append(input, "b", 2)

	// Sort by keys in reverse order
	sorted := seq2.SortComparingKeys(input, func(a, b string) int {
		return -cmp.Compare(a, b)
	})

	seq2.ForEach(sorted, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
}
```

**Output**

```
c : 3
b : 2
a : 1
```


</details>

<a name="SortComparingValues"></a>
## [SortComparingValues](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/sort.go#L45>)

```go
func SortComparingValues[K any, V any](seq iter.Seq2[K, V], cmp func(V, V) int) iter.Seq2[K, V]
```

SortComparingValues sorts the elements of a sequence by value in ascending order.

<details>
<summary>Example</summary>




```go
package main

import (
	"cmp"
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Unordered map with string keys
	input := seq2.Single("a", 1)
	input = seq2.Append(input, "c", 3)
	input = seq2.Append(input, "b", 2)

	// Sort by values in descending order
	sorted := seq2.SortComparingValues(input, func(a, b int) int {
		return -cmp.Compare(a, b)
	})

	seq2.ForEach(sorted, func(k string, v int) {
		fmt.Println(k, ":", v)
	})
}
```

**Output**

```
c : 3
b : 2
a : 1
```


</details>

<a name="Split"></a>
## [Split](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/joins.go#L45>)

```go
func Split[K any, V any](seq iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V])
```

Split splits a sequence of pairs into two sequences.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

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
}
```

**Output**

```
Keys: [a b c]
Values: [1 2 3]
```


</details>

<a name="Take"></a>
## [Take](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L113>)

```go
func Take[K any, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V]
```

Take returns a new sequence that contains only the first n elements of the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	// Take the first 2 elements
	taken := seq2.Take(input, 2)

	result := seq2.CollectToMap(taken)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2]
```


</details>

<a name="TakeUntil"></a>
## [TakeUntil](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L140>)

```go
func TakeUntil[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V]
```

TakeUntil returns a new sequence that takes elements from the given sequence until the predicate is satisfied.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	// Take elements until value is greater than 2
	taken := seq2.TakeUntil(input, func(k string, v int) bool {
		return v > 2
	})

	result := seq2.CollectToMap(taken)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2]
```


</details>

<a name="TakeWhile"></a>
## [TakeWhile](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L129>)

```go
func TakeWhile[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V]
```

TakeWhile returns a new sequence that takes elements from the given sequence while the predicate is satisfied.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.OfIndexed("a", "b", "c", "d")

	// Take elements while value is less than 3
	taken := seq2.TakeWhile(input, func(k int, v string) bool {
		return v != "c"
	})

	result := seq2.CollectToMap(taken)
	fmt.Println(result)
}
```

**Output**

```
map[0:a 1:b]
```


</details>

<a name="Tap"></a>
## [Tap](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/consumer.go#L15>)

```go
func Tap[K any, V any](seq iter.Seq2[K, V], consumer Consumer[K, V]) iter.Seq2[K, V]
```

Tap returns a sequence that applies the given consumer to each element of the input sequence and pass it further.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	// Ensure to have consistent output
	input = seq2.SortByKeys(input)

	// Use Tap to print key-value pairs while passing them through
	tapped := seq2.Tap(input, func(k string, v int) {
		fmt.Printf("Processing: %s => %d\n", k, v)
	})

	seq2.Flush(tapped)

}
```

**Output**

```
Processing: a => 1
Processing: b => 2
Processing: c => 3
```


</details>

<a name="Tick"></a>
## [Tick](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L96>)

```go
func Tick(d time.Duration) iter.Seq2[int, time.Time]
```

Tick returns a sequence that yields the tick number and the current time every duration.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"time"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	ticker := seq2.Tick(1 * time.Millisecond)

	ticker = seq2.Take(ticker, 5)

	seq2.ForEach(ticker, func(tick int, v time.Time) {
		fmt.Printf("tick %d at %s \n", tick, v.Format("15:04:05.000"))
	})

	// Example Output:
	// tick 1 at 00:00:00.000
	// tick 2 at 00:00:00.001
	// tick 3 at 00:00:00.002
	// tick 4 at 00:00:00.003
	// tick 5 at 00:00:00.004
}
```


</details>

<a name="ToMap"></a>
## [ToMap](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/consumer.go#L48>)

```go
func ToMap[Map ~map[K]V, K comparable, V any](seq iter.Seq2[K, V], m Map)
```

ToMap collects the elements of the given sequence into a map.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

	result := make(map[string]int, 3)
	seq2.ToMap(input, result)

	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="UnZip"></a>
## [UnZip](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/joins.go#L40>)

```go
func UnZip[K any, V any](seq iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V])
```

UnZip splits a sequence of pairs into two sequences.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	input = seq2.SortByKeys(input)

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
}
```

**Output**

```
Keys: [a b c]
Values: [1 2 3]
```


</details>

<details>
<summary>Example (After Zip)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	nums := seq.Of(1, 2, 3)
	letters := seq.Of("a", "b", "c")

	zipped := seq.Zip(nums, letters)

	unzipped1, unzipped2 := seq2.UnZip(zipped)

	result1 := seq.Collect(unzipped1)
	result2 := seq.Collect(unzipped2)

	fmt.Println(result1)
	fmt.Println(result2)
}
```

**Output**

```
[1 2 3]
[a b c]
```


</details>

<a name="Union"></a>
## [Union](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/joins.go#L19>)

```go
func Union[K comparable, V comparable](seq1 iter.Seq2[K, V], seq2 iter.Seq2[K, V]) iter.Seq2[K, V]
```

Union returns a sequence that contains all distinct elements from both input sequences.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	first := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	first = seq2.SortByKeys(first)
	second := seq2.FromMap(map[string]int{"c": 3, "d": 4, "e": 5})
	second = seq2.SortByKeys(second)

	// Union returns distinct elements from both sequences
	combined := seq2.Union(first, second)

	result := seq2.CollectToMap(combined)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3 d:4 e:5]
```


</details>

<a name="UnionAll"></a>
## [UnionAll](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/joins.go#L24>)

```go
func UnionAll[K any, V any](seq1 iter.Seq2[K, V], seq2 iter.Seq2[K, V]) iter.Seq2[K, V]
```

UnionAll returns a sequence that contains all elements from both input sequences.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	first := seq2.FromMap(map[string]int{"a": 1, "b": 2})
	first = seq2.SortByKeys(first)
	second := seq2.FromMap(map[string]int{"c": 3, "b": 2})
	second = seq2.SortByKeys(second)

	// UnionAll is an alias for Concat
	combined := seq2.UnionAll(first, second)

	result := seq2.CollectToMap(combined)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="Uniq"></a>
## [Uniq](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L159>)

```go
func Uniq[K comparable, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

Uniq returns a new sequence that contains only the unique elements of the given sequence. It compares both key and value. In case of pointers, pointers are compared, not the values they point to.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create a sequence with duplicate key-value pairs
	input := seq2.Concat(
		seq2.FromMap(map[string]int{"a": 1, "b": 2}),
		seq2.FromMap(map[string]int{"a": 1, "c": 3}),
	)

	// Get unique key-value pairs
	unique := seq2.Uniq(input)

	result := seq2.CollectToMap(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="UniqBy"></a>
## [UniqBy](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L209>)

```go
func UniqBy[K any, V any, K2 comparable](seq iter.Seq2[K, V], mapper Mapper[K, V, K2]) iter.Seq2[K, V]
```

UniqBy returns a new sequence that contains only the unique elements of the given sequence based on result of the mapper.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"strconv"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"apple": 1, "banana": 2, "apricot": 3, "berry": 4, "blueberry": 5})
	input = seq2.SortByKeys(input)

	// Get unique entries based on first letter and value modulo 2
	unique := seq2.UniqBy(input, func(k string, v int) string {
		return string(k[0]) + strconv.Itoa(v%2)
	})

	result := seq2.CollectToMap(unique)
	fmt.Println(result)
}
```

**Output**

```
map[apple:1 banana:2 blueberry:5]
```


</details>

<a name="UniqByKeys"></a>
## [UniqByKeys](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L225>)

```go
func UniqByKeys[K any, V any, K2 comparable](seq iter.Seq2[K, V], mapper KeyMapper[K, K2]) iter.Seq2[K, V]
```

UniqByKeys returns a new sequence that contains only the unique elements of the given sequence based on a result of key mapper.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"Apple": 1, "apricot": 2, "Banana": 3, "berry": 4})
	input = seq2.SortByKeys(input)

	// Get unique entries based on lowercase first letter of key
	unique := seq2.UniqByKeys(input, func(k string) string {
		return strings.ToLower(string(k[0]))
	})

	result := seq2.CollectToMap(unique)
	fmt.Println(result)
}
```

**Output**

```
map[Apple:1 Banana:3]
```


</details>

<a name="UniqByValues"></a>
## [UniqByValues](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L241>)

```go
func UniqByValues[K any, V any, V2 comparable](seq iter.Seq2[K, V], mapper ValueMapper[V, V2]) iter.Seq2[K, V]
```

UniqByValues returns a new sequence that contains only the unique elements of the given sequence based on a result of value mapper.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 10, "b": 21, "c": 30, "d": 44})
	input = seq2.SortByKeys(input)

	// Get unique entries based on value modulo 10
	unique := seq2.UniqByValues(input, func(v int) int {
		return v % 10
	})

	result := seq2.CollectToMap(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:10 b:21 d:44]
```


</details>

<a name="UniqKeys"></a>
## [UniqKeys](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L179>)

```go
func UniqKeys[K comparable, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

UniqKeys returns a new sequence that contains only the elements with unique keys from the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create a sequence with duplicate keys
	input := seq2.Concat(
		seq2.FromMap(map[string]int{"a": 1, "b": 2}),
		seq2.FromMap(map[string]int{"a": 3, "c": 4}),
	)

	// Get entries with unique keys (first occurrence wins)
	unique := seq2.UniqKeys(input)

	result := seq2.CollectToMap(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:4]
```


</details>

<a name="UniqValues"></a>
## [UniqValues](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L194>)

```go
func UniqValues[K any, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

UniqValues returns a new sequence that contains only the elements with unique values from the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create a sequence with duplicate values
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 1, "d": 3})
	input = seq2.SortByKeys(input)

	// Get entries with unique values (first occurrence wins)
	unique := seq2.UniqValues(input)

	result := seq2.CollectToMap(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 d:3]
```


</details>

<a name="Values"></a>
## [Values](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L28>)

```go
func Values[K, V any](seq iter.Seq2[K, V]) iter.Seq[V]
```

Values returns a sequence of values from a sequence of key\-value pairs.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.OfIndexed("a", "b", "c")

	keys := seq2.Values(input)

	seq.ForEach(keys, func(v string) {
		fmt.Print(v, " ")
	})
}
```

**Output**

```
a b c
```


</details>

<a name="Where"></a>
## [Where](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L29>)

```go
func Where[K any, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) iter.Seq2[K, V]
```

Where returns a new sequence that contains only the elements that satisfy the predicate. SQL\-like alias for Filter

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.SortByKeys(input)

	// Where is an alias for Filter
	filtered := seq2.Where(input, func(k string, v int) bool {
		return v > 2
	})

	result := seq2.CollectToMap(filtered)
	fmt.Println(result)
}
```

**Output**

```
map[c:3 d:4]
```


</details>

<a name="WithIndex"></a>
## [WithIndex](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L67>)

```go
func WithIndex[E any](seq iter.Seq[E]) iter.Seq2[int, E]
```

WithIndex creates a new indexed sequence from the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create a iter.Seq of values
	values := seq.Of("a", "b", "c")

	// Add indexes
	indexed := seq2.WithIndex(values)

	result := seq2.CollectToMap(indexed)
	fmt.Println(result)
}
```

**Output**

```
map[0:a 1:b 2:c]
```


</details>

<a name="WithoutIndex"></a>
## [WithoutIndex](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/producers.go#L80>)

```go
func WithoutIndex[E any](indexed iter.Seq2[int, E]) iter.Seq[E]
```

WithoutIndex creates a new sequence from the given indexed sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Create an indexed sequence
	indexed := seq2.OfIndexed("a", "b", "c")

	// Remove indexes
	values := seq2.WithoutIndex(indexed)

	result := seq.Collect(values)
	fmt.Println(result)
}
```

**Output**

```
[a b c]
```


</details>

<a name="Consumer"></a>
## type [Consumer](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/consumer.go#L12>)

Consumer is a function that consumes an element of an iter.Seq2.

```go
type Consumer[K any, V any] = func(K, V)
```

<a name="DoubleMapper"></a>
## type [DoubleMapper](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/mapper.go#L9>)

DoubleMapper is a function that takes an element and returns a new sequence element.

```go
type DoubleMapper[K, V, RK, RV any] = func(K, V) (RK, RV)
```

<a name="KeyMapper"></a>
## type [KeyMapper](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/mapper.go#L12>)

KeyMapper is a function that takes Key a new Key.

```go
type KeyMapper[K, R any] = func(K) R
```

<a name="KeyPredicate"></a>
## type [KeyPredicate](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L9>)

KeyPredicate is a function that is used to filter by key.

```go
type KeyPredicate[E any] = KeyMapper[E, bool]
```

<a name="Mapper"></a>
## type [Mapper](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/mapper.go#L6>)

Mapper is a function that takes an element and returns a new element.

```go
type Mapper[K, V, R any] = func(K, V) R
```

<a name="Predicate"></a>
## type [Predicate](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L6>)

Predicate is a function that takes an element and returns a boolean.

```go
type Predicate[K any, V any] = Mapper[K, V, bool]
```

<a name="ValueMapper"></a>
## type [ValueMapper](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/mapper.go#L15>)

ValueMapper is a function that takes Value a new Value.

```go
type ValueMapper[V, R any] = func(V) R
```

<a name="ValuePredicate"></a>
## type [ValuePredicate](<https://github.com/go-softwarelab/common/blob/main/pkg/seq2/filter.go#L12>)

ValuePredicate is a function that is used to filter by value.

```go
type ValuePredicate[E any] = ValueMapper[E, bool]
```
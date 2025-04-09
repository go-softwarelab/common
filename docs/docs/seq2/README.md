# seq2

```go
import "github.com/go-softwarelab/common/pkg/seq2"
```

Package seq2 provides a comprehensive set of utilities for working with key\-value sequences in Go applications.

The goal of this package is to offer a rich set of functions for creating, transforming, and consuming iter.Seq2, enabling developers to work with collections of key\-value pairs in a functional programming style. The package includes utilities for filtering, mapping, reducing, and sorting sequences, as well as combining and partitioning them.

The package is designed to reduce boilerplate code and improve readability by providing a consistent API for common sequence operations. It leverages Go's type safety and generics to ensure that operations on sequences are both flexible and safe. The Sequence struct is worth mentioning explicitly, allowing method chaining and fluent composition of sequence operations.



<a name="Append"></a>
## Append

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
	input = seq2.Sort(input)

	// Append a new key-value pair to the sequence
	appended := seq2.Append(input, "c", 3)

	result := seq2.Collect(appended)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="Collect"></a>
## Collect

```go
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V
```

Collect collects the elements of the given sequence into a map.

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
	input = seq2.Sort(input)

	result := seq2.Collect(input)

	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="Concat"></a>
## Concat

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
	first = seq2.Sort(first)
	second := seq2.FromMap(map[string]int{"c": 3, "d": 4})
	second = seq2.Sort(second)

	// Concatenate two sequences
	combined := seq2.Concat(first, second)

	result := seq2.Collect(combined)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3 d:4]
```


</details>

<a name="Contains"></a>
## Contains

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
	input = seq2.Sort(input)

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
## ContainsAll

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
	input = seq2.Sort(input)

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
## ContainsAllValues

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
	input = seq2.Sort(input)

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
## ContainsPair

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
	input = seq2.Sort(input)

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
## ContainsValue

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
	input = seq2.Sort(input)

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
## Count

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
	input = seq2.Sort(input)

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
## Cycle

```go
func Cycle[K, V any](seq iter.Seq2[K, V], count int) iter.Seq2[K, V]
```

Cycle repeats the sequence count times.

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
	input = seq2.Sort(input)

	// Repeat the sequence 2 times
	cycled := seq2.Cycle(input, 2)

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
## Distinct

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

	result := seq2.Collect(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="DistinctKeys"></a>
## DistinctKeys

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

	result := seq2.Collect(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:4]
```


</details>

<a name="Each"></a>
## Each

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
	input = seq2.Sort(input)

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
## Empty

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
## Every

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
	input = seq2.Sort(input)

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
## Exists

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
	input = seq2.Sort(input)

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
## Filter

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
	input = seq2.Sort(input)

	// Filter elements where the value is even
	filtered := seq2.Filter(input, func(k string, v int) bool {
		return v%2 == 0
	})

	result := seq2.Collect(filtered)
	fmt.Println(result)
}
```

**Output**

```
map[b:2 d:4]
```


</details>

<a name="FilterByKey"></a>
## FilterByKey

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
	input = seq2.Sort(input)

	// Filter elements by key
	filtered := seq2.FilterByKey(input, func(k string) bool {
		return k > "b"
	})

	result := seq2.Collect(filtered)
	fmt.Println(result)
}
```

**Output**

```
map[c:3 d:4]
```


</details>

<a name="FilterByValue"></a>
## FilterByValue

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
	input = seq2.Sort(input)

	// Filter elements by value
	filtered := seq2.FilterByValue(input, func(v int) bool {
		return v <= 2
	})

	result := seq2.Collect(filtered)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2]
```


</details>

<a name="Find"></a>
## Find

```go
func Find[K, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) (optional.Elem[K], optional.Elem[V])
```

Find returns the first element that satisfies the predicate.

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
	input = seq2.Sort(input)

	k, v := seq2.Find(input, func(k string, v int) bool {
		return v > 2
	})

	fmt.Printf("Key: %s, Value: %d\n", k.MustGet(), v.MustGet())
}
```

**Output**

```
Key: c, Value: 3
```


</details>

<a name="FindAll"></a>
## FindAll

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
	input = seq2.Sort(input)

	found := seq2.FindAll(input, func(k string, v int) bool {
		return v > 2
	})

	result := seq2.Collect(found)
	fmt.Println(result)
}
```

**Output**

```
map[c:3 d:4]
```


</details>

<a name="FindByKey"></a>
## FindByKey

```go
func FindByKey[K, V any](seq iter.Seq2[K, V], predicate KeyPredicate[K]) (optional.Elem[K], optional.Elem[V])
```

FindByKey returns the first element that satisfies the predicate.

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
	input := seq2.FromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})
	input = seq2.Sort(input)

	k, v := seq2.FindByKey(input, func(k string) bool {
		return strings.ToUpper(k) == "B"
	})

	fmt.Printf("Key: %s, Value: %d\n", k.MustGet(), v.MustGet())
}
```

**Output**

```
Key: b, Value: 2
```


</details>

<a name="FindByValue"></a>
## FindByValue

```go
func FindByValue[K, V any](seq iter.Seq2[K, V], predicate ValuePredicate[V]) (optional.Elem[K], optional.Elem[V])
```

FindByValue returns the first element that satisfies the predicate.

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
	input = seq2.Sort(input)

	k, v := seq2.FindByValue(input, func(v int) bool {
		return v%4 == 0
	})

	fmt.Printf("Key: %s, Value: %d\n", k.MustGet(), v.MustGet())
}
```

**Output**

```
Key: d, Value: 4
```


</details>

<a name="FindLast"></a>
## FindLast

```go
func FindLast[K, V any](seq iter.Seq2[K, V], predicate Predicate[K, V]) (optional.Elem[K], optional.Elem[V])
```

FindLast returns the last element that satisfies the predicate.

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
	input = seq2.Sort(input)

	k, v := seq2.FindLast(input, func(k string, v int) bool {
		return v > 2
	})

	fmt.Printf("Key: %s, Value: %d\n", k.MustGet(), v.MustGet())
}
```

**Output**

```
Key: d, Value: 4
```


</details>

<a name="Flush"></a>
## Flush

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

<a name="FoldKeys"></a>
## FoldKeys

```go
func FoldKeys[K any, V any](seq2 iter.Seq2[K, V], accumulator func(agg K, key K, value V) K) optional.Elem[K]
```

FoldKeys applies a function against an accumulator and each element in the sequence \(from left to right\) to reduce it to a single value. Notice: first value will be omitted and the first key will be used as initial value.

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
	input = seq2.Sort(input)

	// FoldKeys concatenates all keys starting with the first key
	result := seq2.FoldKeys(input, func(agg string, key string, value int) string {
		return agg + key
	})

	fmt.Println(result.MustGet())
}
```

**Output**

```
abc
```


</details>

<a name="FoldKeysRight"></a>
## FoldKeysRight

```go
func FoldKeysRight[K any, V any](seq2 iter.Seq2[K, V], accumulator func(agg K, key K, value V) K) optional.Elem[K]
```

FoldKeysRight applies a function against an accumulator and each element in the sequence \(from right to left\) to reduce it to a single value. Notice: last value will be omitted and the last key will be used as initial value.

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
	input = seq2.Sort(input)

	// FoldKeysRight concatenates all keys starting from the right
	result := seq2.FoldKeysRight(input, func(agg string, key string, value int) string {
		return agg + key
	})

	fmt.Println(result.MustGet())
}
```

**Output**

```
cba
```


</details>

<a name="FoldValues"></a>
## FoldValues

```go
func FoldValues[K any, V any](seq2 iter.Seq2[K, V], accumulator func(agg V, key K, value V) V) optional.Elem[V]
```

FoldValues applies a function against an accumulator and each element in the sequence \(from left to right\) to reduce it to a single value. Notice: first key will be omitted and the first value will be used as initial value.

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
	input = seq2.Sort(input)

	// FoldValues to calculate the product of all values
	product := seq2.FoldValues(input, func(agg int, key string, value int) int {
		return agg * value
	})

	fmt.Println(product.MustGet())
}
```

**Output**

```
6
```


</details>

<a name="FoldValuesRight"></a>
## FoldValuesRight

```go
func FoldValuesRight[K any, V any](seq2 iter.Seq2[K, V], accumulator func(agg V, key K, value V) V) optional.Elem[V]
```

FoldValuesRight applies a function against an accumulator and each element in the sequence \(from right to left\) to reduce it to a single value. Notice: last key will be omitted and the last value will be used as initial value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 3, "b": 5, "c": 10})
	input = seq2.Sort(input)

	// FoldValuesRight performs right-to-left division
	result := seq2.FoldValuesRight(input, func(agg int, key string, value int) int {
		return agg - value
	})

	fmt.Println(result.MustGet())
}
```

**Output**

```
2
```


</details>

<a name="ForEach"></a>
## ForEach

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
	input := seq2.Of("a", "b", "c")

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
## FromMap

```go
func FromMap[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V]
```

FromMap creates a new iter.Seq2 from the given map.

<a name="FromSlice"></a>
## FromSlice

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

	result := seq2.Collect(sequence)
	fmt.Println(result)
}
```

**Output**

```
map[0:a 1:b 2:c]
```


</details>

<a name="Get"></a>
## Get

```go
func Get[K comparable, V any](seq iter.Seq2[K, V], key K) optional.Elem[V]
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
	input = seq2.Sort(input)

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
## IsEmpty

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

	nonEmptySeq := seq2.Of("a")

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
## IsNotEmpty

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

	nonEmptySeq := seq2.Of("a")

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
## Keys

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
	input := seq2.Of("a", "b", "c")

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
## Limit

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
	input = seq2.Sort(input)

	// Limit is an alias for Take
	taken := seq2.Limit(input, 3)

	result := seq2.Collect(taken)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="Map"></a>
## Map

```go
func Map[K, V, RV any](seq iter.Seq2[K, V], mapper Mapper[K, V, RV]) iter.Seq2[K, RV]
```

Map applies a mapper function to each value of the sequence.

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
	input = seq2.Sort(input)

	// Map both key and value to produce a new value (keeps original keys)
	mapped := seq2.Map(input, func(k string, v int) string {
		return fmt.Sprintf("Value of %s is %d", k, v)
	})

	result := seq2.Collect(mapped)
	fmt.Println(result)
}
```

**Output**

```
map[a:Value of a is 1 b:Value of b is 2 c:Value of c is 3]
```


</details>

<a name="MapKeys"></a>
## MapKeys

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
	input = seq2.Sort(input)

	// Map keys to uppercase (keeps original values)
	mapped := seq2.MapKeys(input, func(k string) string {
		return strings.ToUpper(k)
	})

	result := seq2.Collect(mapped)
	fmt.Println(result)
}
```

**Output**

```
map[A:1 B:2 C:3]
```


</details>

<a name="MapPairs"></a>
## MapPairs

```go
func MapPairs[K, V, RK, RV any](seq iter.Seq2[K, V], mapper DoubleMapper[K, V, RK, RV]) iter.Seq2[RK, RV]
```

MapPairs applies a mapper function to each element of the sequence.

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
	input = seq2.Sort(input)

	// Map keys to uppercase and multiply values by 10
	mapped := seq2.MapPairs(input, func(k string, v int) (string, int) {
		return strings.ToUpper(k), v * 10
	})

	result := seq2.Collect(mapped)
	fmt.Println(result)
}
```

**Output**

```
map[A:10 B:20 C:30]
```


</details>

<a name="MapTo"></a>
## MapTo

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
	input = seq2.Sort(input)

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

<a name="MapToValues"></a>
## MapToValues

```go
func MapToValues[K, V, RV any](seq iter.Seq2[K, V], mapper ValueMapper[V, RV]) iter.Seq[RV]
```

MapToValues applies a mapper function to each value of the sequence and returns a sequence of values only.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.Of("a", "b", "c")

	// Map values to strings and return as sequence of values
	mapped := seq2.MapToValues(input, func(v string) string {
		return fmt.Sprintf("Value: %s", v)
	})

	for v := range mapped {
		fmt.Println(v)
	}

}
```

**Output**

```
Value: a
Value: b
Value: c
```


</details>

<a name="MapValues"></a>
## MapValues

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
	input = seq2.Sort(input)

	// Map values to their squares (keeps original keys)
	mapped := seq2.MapValues(input, func(v int) int {
		return v * v
	})

	result := seq2.Collect(mapped)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:4 c:9]
```


</details>

<a name="MaxKey"></a>
## MaxKey

```go
func MaxKey[K types.Ordered, V any](seq2 iter.Seq2[K, V]) optional.Elem[K]
```

MaxKey returns the maximum key in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 5, "z": 10, "c": 3})
	input = seq2.Sort(input)

	// Find the maximum key (lexicographically)
	maxK := seq2.MaxKey(input)

	fmt.Println(maxK.MustGet())
}
```

**Output**

```
z
```


</details>

<a name="MaxValue"></a>
## MaxValue

```go
func MaxValue[K any, V types.Ordered](seq2 iter.Seq2[K, V]) optional.Elem[V]
```

MaxValue returns the maximum value in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 5, "b": 10, "c": 3})
	input = seq2.Sort(input)

	// Find the maximum value
	maxV := seq2.MaxValue(input)

	fmt.Println(maxV.MustGet())
}
```

**Output**

```
10
```


</details>

<a name="MinKey"></a>
## MinKey

```go
func MinKey[K types.Ordered, V any](seq2 iter.Seq2[K, V]) optional.Elem[K]
```

MinKey returns the minimum key in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 5, "z": 10, "c": 3})
	input = seq2.Sort(input)

	// Find the minimum key (lexicographically)
	minK := seq2.MinKey(input)

	fmt.Println(minK.MustGet())
}
```

**Output**

```
a
```


</details>

<a name="MinValue"></a>
## MinValue

```go
func MinValue[K any, V types.Ordered](seq2 iter.Seq2[K, V]) optional.Elem[V]
```

MinValue returns the minimum value in the .

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.FromMap(map[string]int{"a": 5, "b": 10, "c": 3})
	input = seq2.Sort(input)

	// Find the minimum value
	minV := seq2.MinValue(input)

	fmt.Println(minV.MustGet())
}
```

**Output**

```
3
```


</details>

<a name="Narrow"></a>
## Narrow

```go
func Narrow[K, V, RV any](seq iter.Seq2[K, V], mapper Mapper[K, V, RV]) iter.Seq[RV]
```

Narrow applies a mapper function to each element of the sequence and returns a sequence of mapper results. Alias for MapTo

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
	input = seq2.Sort(input)

	// Narrow is an alias for MapTo
	mapped := seq2.Narrow(input, func(k string, v int) string {
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

<a name="None"></a>
## None

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
	input = seq2.Sort(input)

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
## NotContains

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
	input = seq2.Sort(input)

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
## NotContainsPair

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
	input = seq2.Sort(input)

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
## NotContainsValue

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
	input = seq2.Sort(input)

	notContainsValue := seq2.NotContainsValue(input, 5)

	fmt.Println(notContainsValue)
}
```

**Output**

```
true
```


</details>

<a name="Of"></a>
## Of

```go
func Of[E any](elems ...E) iter.Seq2[int, E]
```

Of creates a new indexed sequence from the given elements.

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
	indexed := seq2.Of(1, 2, 3)

	result := seq2.Collect(indexed)
	fmt.Println(result)
}
```

**Output**

```
map[0:1 1:2 2:3]
```


</details>

<a name="Offset"></a>
## Offset

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
	input := seq2.Of(10, 20, 30, 40, 20)

	// Skip the first 2 elements
	skipped := seq2.Offset(input, 2)

	result := seq2.Collect(skipped)
	fmt.Println(result)
}
```

**Output**

```
map[2:30 3:40 4:20]
```


</details>

<a name="Pair"></a>
## Pair

```go
func Pair[K, V any](k K, v V) iter.Seq2[K, V]
```

Pair returns a sequence with given key value.

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
	single := seq2.Pair("key", 42)

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

<a name="Prepend"></a>
## Prepend

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
	input = seq2.Sort(input)

	// Prepend a new key-value pair to the sequence
	prepended := seq2.Prepend(input, "a", 1)

	result := seq2.Collect(prepended)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="Reduce"></a>
## Reduce

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
	input = seq2.Sort(input)

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
## ReduceRight

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
	input = seq2.Sort(input)

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
## Repeat

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
## Reverse

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
	sequence := seq2.Of("a", "b", "c")

	// Reverse it
	reversed := seq2.Reverse(sequence)

	// Collect into pairs for ordered display
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

<a name="Skip"></a>
## Skip

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
	input := seq2.Of(10, 20, 30, 40, 20)

	// Skip the first 2 elements
	skipped := seq2.Skip(input, 2)

	result := seq2.Collect(skipped)
	fmt.Println(result)
}
```

**Output**

```
map[2:30 3:40 4:20]
```


</details>

<a name="SkipUntil"></a>
## SkipUntil

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
	input := seq2.Of(10, 20, 30, 40, 20)

	// Skip elements until value of 30 is reached
	skipped := seq2.SkipUntil(input, func(k int, v int) bool {
		return v == 30
	})

	result := seq2.Collect(skipped)
	fmt.Println(result)
}
```

**Output**

```
map[2:30 3:40 4:20]
```


</details>

<a name="SkipWhile"></a>
## SkipWhile

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
	input := seq2.Of(10, 20, 30, 40, 20)

	// Skip elements until the value is less than 30
	skipped := seq2.SkipWhile(input, func(k int, v int) bool {
		return v < 30
	})

	result := seq2.Collect(skipped)
	fmt.Println(result)
}
```

**Output**

```
map[2:30 3:40 4:20]
```


</details>

<a name="Sort"></a>
## Sort

```go
func Sort[K types.Ordered, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

Sort sorts the elements of a sequence by key in ascending order.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	input := seq2.Pair("a", 1)
	input = seq2.Append(input, "c", 3)
	input = seq2.Append(input, "b", 2)

	// Sort by keys (alphabetically)
	sorted := seq2.Sort(input)

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

<a name="SortBy"></a>
## SortBy

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
	input := seq2.Pair("a", 1)
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

<a name="SortComparingKeys"></a>
## SortComparingKeys

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
	input := seq2.Pair("a", 1)
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
## SortComparingValues

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
	input := seq2.Pair("a", 1)
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
## Split

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
	input = seq2.Sort(input)

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
## Take

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
	input = seq2.Sort(input)

	// Take the first 2 elements
	taken := seq2.Take(input, 2)

	result := seq2.Collect(taken)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2]
```


</details>

<a name="TakeUntil"></a>
## TakeUntil

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
	input = seq2.Sort(input)

	// Take elements until value is greater than 2
	taken := seq2.TakeUntil(input, func(k string, v int) bool {
		return v > 2
	})

	result := seq2.Collect(taken)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2]
```


</details>

<a name="TakeWhile"></a>
## TakeWhile

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
	input := seq2.Of("a", "b", "c", "d")

	// Take elements while value is less than 3
	taken := seq2.TakeWhile(input, func(k int, v string) bool {
		return v != "c"
	})

	result := seq2.Collect(taken)
	fmt.Println(result)
}
```

**Output**

```
map[0:a 1:b]
```


</details>

<a name="Tap"></a>
## Tap

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
	input = seq2.Sort(input)

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
## Tick

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
## ToMap

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
	input = seq2.Sort(input)

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
## UnZip

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
	input = seq2.Sort(input)

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
## Union

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
	first = seq2.Sort(first)
	second := seq2.FromMap(map[string]int{"c": 3, "d": 4, "e": 5})
	second = seq2.Sort(second)

	// Union returns distinct elements from both sequences
	combined := seq2.Union(first, second)

	result := seq2.Collect(combined)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3 d:4 e:5]
```


</details>

<a name="UnionAll"></a>
## UnionAll

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
	first = seq2.Sort(first)
	second := seq2.FromMap(map[string]int{"c": 3, "b": 2})
	second = seq2.Sort(second)

	// UnionAll is an alias for Concat
	combined := seq2.UnionAll(first, second)

	result := seq2.Collect(combined)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="Uniq"></a>
## Uniq

```go
func Uniq[K comparable, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

Uniq returns a new sequence that contains only the unique elements of the given sequence.

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

	result := seq2.Collect(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:3]
```


</details>

<a name="UniqBy"></a>
## UniqBy

```go
func UniqBy[K any, V any, K2 comparable](seq iter.Seq2[K, V], mapper Mapper[K, V, K2]) iter.Seq2[K, V]
```

UniqBy returns a new sequence that contains only the unique elements of the given sequence based on a key.

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
	input = seq2.Sort(input)

	// Get unique entries based on first letter and value modulo 2
	unique := seq2.UniqBy(input, func(k string, v int) string {
		return string(k[0]) + strconv.Itoa(v%2)
	})

	result := seq2.Collect(unique)
	fmt.Println(result)
}
```

**Output**

```
map[apple:1 banana:2 blueberry:5]
```


</details>

<a name="UniqByKeys"></a>
## UniqByKeys

```go
func UniqByKeys[K any, V any, K2 comparable](seq iter.Seq2[K, V], mapper KeyMapper[K, K2]) iter.Seq2[K, V]
```

UniqByKeys returns a new sequence that contains only the unique elements of the given sequence based on a key.

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
	input = seq2.Sort(input)

	// Get unique entries based on lowercase first letter of key
	unique := seq2.UniqByKeys(input, func(k string) string {
		return strings.ToLower(string(k[0]))
	})

	result := seq2.Collect(unique)
	fmt.Println(result)
}
```

**Output**

```
map[Apple:1 Banana:3]
```


</details>

<a name="UniqByValues"></a>
## UniqByValues

```go
func UniqByValues[K any, V any, V2 comparable](seq iter.Seq2[K, V], mapper ValueMapper[V, V2]) iter.Seq2[K, V]
```

UniqByValues returns a new sequence that contains only the unique elements of the given sequence based on a key.

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
	input = seq2.Sort(input)

	// Get unique entries based on value modulo 10
	unique := seq2.UniqByValues(input, func(v int) int {
		return v % 10
	})

	result := seq2.Collect(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:10 b:21 d:44]
```


</details>

<a name="UniqKeys"></a>
## UniqKeys

```go
func UniqKeys[K comparable, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

UniqKeys returns a new sequence that contains only the unique keys of the given sequence.

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

	result := seq2.Collect(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 c:4]
```


</details>

<a name="UniqValues"></a>
## UniqValues

```go
func UniqValues[K any, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V]
```

UniqValues returns a new sequence that contains only the unique values of the given sequence.

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
	input = seq2.Sort(input)

	// Get entries with unique values (first occurrence wins)
	unique := seq2.UniqValues(input)

	result := seq2.Collect(unique)
	fmt.Println(result)
}
```

**Output**

```
map[a:1 b:2 d:3]
```


</details>

<a name="Values"></a>
## Values

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
	input := seq2.Of("a", "b", "c")

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
## Where

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
	input = seq2.Sort(input)

	// Where is an alias for Filter
	filtered := seq2.Where(input, func(k string, v int) bool {
		return v > 2
	})

	result := seq2.Collect(filtered)
	fmt.Println(result)
}
```

**Output**

```
map[c:3 d:4]
```


</details>

<a name="WithIndex"></a>
## WithIndex

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

	result := seq2.Collect(indexed)
	fmt.Println(result)
}
```

**Output**

```
map[0:a 1:b 2:c]
```


</details>

<a name="WithoutIndex"></a>
## WithoutIndex

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
	indexed := seq2.Of("a", "b", "c")

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
## type Consumer

Consumer is a function that consumes an element of an iter.Seq2.

```go
type Consumer[K any, V any] = func(K, V)
```

<a name="DoubleMapper"></a>
## type DoubleMapper

DoubleMapper is a function that takes an element and returns a new sequence element.

```go
type DoubleMapper[K, V, RK, RV any] = func(K, V) (RK, RV)
```

<a name="KeyMapper"></a>
## type KeyMapper

KeyMapper is a function that takes Key a new Key.

```go
type KeyMapper[K, R any] = func(K) R
```

<a name="KeyPredicate"></a>
## type KeyPredicate

KeyPredicate is a function that is used to filter by key.

```go
type KeyPredicate[E any] = KeyMapper[E, bool]
```

<a name="Mapper"></a>
## type Mapper

Mapper is a function that takes an element and returns a new element.

```go
type Mapper[K, V, R any] = func(K, V) R
```

<a name="Predicate"></a>
## type Predicate

Predicate is a function that takes an element and returns a boolean.

```go
type Predicate[K any, V any] = Mapper[K, V, bool]
```

<a name="Sequence"></a>
## type Sequence

Sequence is a wrapper around iter.Seq2 that provides fluent API for seq2 operations.

```go
type Sequence[K comparable, V comparable] struct {
    // contains filtered or unexported fields
}
```

<a name="AsSequence"></a>
### AsSequence

```go
func AsSequence[K comparable, V comparable](seq iter.Seq2[K, V]) Sequence[K, V]
```

AsSequence wraps an iter.Seq2 to provide a fluent API.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	var result []string
	sequence.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
0:a, 1:b, 2:c
```


</details>

<a name="ConcatSequences"></a>
### ConcatSequences

```go
func ConcatSequences[K comparable, V comparable](sequences ...Sequence[K, V]) Sequence[K, V]
```

ConcatSequences concatenates multiple sequences into a single sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	sq1 := seq2.AsSequence(seq2.Of("a", "b"))
	sq2 := seq2.AsSequence(seq2.Of("c", "d"))

	concatenated := seq2.ConcatSequences(sq1, sq2)

	concatenated.ForEach(func(k int, v string) {
		fmt.Println(k, v)
	})

}
```

**Output**

```
0 a
1 b
0 c
1 d
```


</details>

<a name="Sequence[K, V].Append"></a>
### Sequence[K, V].Append

```go
func (s Sequence[K, V]) Append(key K, value V) Sequence[K, V]
```

Append appends a key\-value pair to the end of a sequence.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))
	appended := sequence.Append(3, "d")

	var result []string
	appended.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
0:a, 1:b, 2:c, 3:d
```


</details>

<a name="Sequence[K, V].Collect"></a>
### Sequence[K, V].Collect

```go
func (s Sequence[K, V]) Collect() map[K]V
```

Collect collects the elements into a map.

<a name="Sequence[K, V].Contains"></a>
### Sequence[K, V].Contains

```go
func (s Sequence[K, V]) Contains(key K) bool
```

Contains returns true if the sequence contains the key.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	fmt.Println("Contains key 1:", sequence.Contains(1))
	fmt.Println("Contains key 5:", sequence.Contains(5))
}
```

**Output**

```
Contains key 1: true
Contains key 5: false
```


</details>

<a name="Sequence[K, V].ContainsPair"></a>
### Sequence[K, V].ContainsPair

```go
func (s Sequence[K, V]) ContainsPair(key K, value V) bool
```

ContainsPair returns true if the sequence contains the key\-value pair.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	fmt.Println("Contains pair 1:'b':", sequence.ContainsPair(1, "b"))
	fmt.Println("Contains pair 1:'c':", sequence.ContainsPair(1, "c"))
}
```

**Output**

```
Contains pair 1:'b': true
Contains pair 1:'c': false
```


</details>

<a name="Sequence[K, V].ContainsValue"></a>
### Sequence[K, V].ContainsValue

```go
func (s Sequence[K, V]) ContainsValue(value V) bool
```

ContainsValue returns true if the sequence contains the value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	fmt.Println("Contains value 'b':", sequence.ContainsValue("b"))
	fmt.Println("Contains value 'z':", sequence.ContainsValue("z"))
}
```

**Output**

```
Contains value 'b': true
Contains value 'z': false
```


</details>

<a name="Sequence[K, V].Count"></a>
### Sequence[K, V].Count

```go
func (s Sequence[K, V]) Count() int
```

Count returns the number of elements in the sequence.

<a name="Sequence[K, V].Cycle"></a>
### Sequence[K, V].Cycle

```go
func (s Sequence[K, V]) Cycle(count int) Sequence[K, V]
```

Cycle repeats the sequence the specified number of times.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b"))
	cycled := sequence.Cycle(2)

	var result []string
	cycled.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
0:a, 1:b, 0:a, 1:b
```


</details>

<a name="Sequence[K, V].Distinct"></a>
### Sequence[K, V].Distinct

```go
func (s Sequence[K, V]) Distinct() Sequence[K, V]
```

Distinct is an alias for Uniq.

<a name="Sequence[K, V].Each"></a>
### Sequence[K, V].Each

```go
func (s Sequence[K, V]) Each(consumer Consumer[K, V]) Sequence[K, V]
```

Each is an alias for Tap.

<a name="Sequence[K, V].Every"></a>
### Sequence[K, V].Every

```go
func (s Sequence[K, V]) Every(predicate Predicate[K, V]) bool
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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	allLowercase := sequence.Every(func(k int, v string) bool {
		return v >= "a" && v <= "z"
	})

	allA := sequence.Every(func(k int, v string) bool {
		return v == "a"
	})

	fmt.Println("All values lowercase:", allLowercase)
	fmt.Println("All values 'a':", allA)
}
```

**Output**

```
All values lowercase: true
All values 'a': false
```


</details>

<a name="Sequence[K, V].Exists"></a>
### Sequence[K, V].Exists

```go
func (s Sequence[K, V]) Exists(predicate Predicate[K, V]) bool
```

Exists returns true if any element satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	hasValueB := sequence.Exists(func(k int, v string) bool {
		return v == "b"
	})

	fmt.Println("Has value 'b':", hasValueB)
}
```

**Output**

```
Has value 'b': true
```


</details>

<a name="Sequence[K, V].Filter"></a>
### Sequence[K, V].Filter

```go
func (s Sequence[K, V]) Filter(predicate Predicate[K, V]) Sequence[K, V]
```

Filter returns a new sequence with elements that satisfy the predicate.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))
	filtered := sequence.Filter(func(k int, v string) bool {
		return k%2 == 0
	})

	var result []string
	filtered.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
0:a, 2:c, 4:e
```


</details>

<a name="Sequence[K, V].FilterByKey"></a>
### Sequence[K, V].FilterByKey

```go
func (s Sequence[K, V]) FilterByKey(predicate KeyPredicate[K]) Sequence[K, V]
```

FilterByKey returns a new sequence with elements whose keys satisfy the predicate.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))
	filtered := sequence.FilterByKey(func(k int) bool {
		return k > 2
	})

	var result []string
	filtered.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
3:d, 4:e
```


</details>

<a name="Sequence[K, V].FilterByValue"></a>
### Sequence[K, V].FilterByValue

```go
func (s Sequence[K, V]) FilterByValue(predicate ValuePredicate[V]) Sequence[K, V]
```

FilterByValue returns a new sequence with elements whose values satisfy the predicate.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))
	filtered := sequence.FilterByValue(func(v string) bool {
		return v > "c"
	})

	var result []string
	filtered.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
3:d, 4:e
```


</details>

<a name="Sequence[K, V].Find"></a>
### Sequence[K, V].Find

```go
func (s Sequence[K, V]) Find(predicate Predicate[K, V]) (optional.Elem[K], optional.Elem[V])
```

Find returns the first element that satisfies the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))

	key, value := sequence.Find(func(k int, v string) bool {
		return v == "c"
	})

	fmt.Printf("Key: %d, Value: %s\n", key.MustGet(), value.MustGet())
}
```

**Output**

```
Key: 2, Value: c
```


</details>

<a name="Sequence[K, V].FindAll"></a>
### Sequence[K, V].FindAll

```go
func (s Sequence[K, V]) FindAll(predicate Predicate[K, V]) Sequence[K, V]
```

FindAll returns all elements that satisfy the predicate.

<a name="Sequence[K, V].FindLast"></a>
### Sequence[K, V].FindLast

```go
func (s Sequence[K, V]) FindLast(predicate Predicate[K, V]) (optional.Elem[K], optional.Elem[V])
```

FindLast returns the last element that satisfies the predicate.

<a name="Sequence[K, V].Flush"></a>
### Sequence[K, V].Flush

```go
func (s Sequence[K, V]) Flush()
```

Flush consumes all elements of the sequence.

<a name="Sequence[K, V].FoldValues"></a>
### Sequence[K, V].FoldValues

```go
func (s Sequence[K, V]) FoldValues(accumulator func(agg V, key K, value V) V) optional.Elem[V]
```

FoldValues folds values with the first value as initial.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	result := sequence.FoldValues(func(agg string, key int, value string) string {
		return agg + value
	})

	fmt.Println("Concatenated:", result.MustGet())
}
```

**Output**

```
Concatenated: abc
```


</details>

<a name="Sequence[K, V].FoldValuesRight"></a>
### Sequence[K, V].FoldValuesRight

```go
func (s Sequence[K, V]) FoldValuesRight(accumulator func(agg V, key K, value V) V) optional.Elem[V]
```

FoldValuesRight folds values from right to left.

<a name="Sequence[K, V].ForEach"></a>
### Sequence[K, V].ForEach

```go
func (s Sequence[K, V]) ForEach(consumer Consumer[K, V])
```

ForEach calls the consumer for each element.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	var result []string
	sequence.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
0:a, 1:b, 2:c
```


</details>

<a name="Sequence[K, V].Get"></a>
### Sequence[K, V].Get

```go
func (s Sequence[K, V]) Get(key K) optional.Elem[V]
```

Get returns the value associated with the key.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))

	existing := sequence.Get(1)
	notExisting := sequence.Get(5)

	fmt.Println("Value for key 1:", existing.MustGet())
	fmt.Println("Has value for key 5:", notExisting.IsPresent())
}
```

**Output**

```
Value for key 1: b
Has value for key 5: false
```


</details>

<a name="Sequence[K, V].IsEmpty"></a>
### Sequence[K, V].IsEmpty

```go
func (s Sequence[K, V]) IsEmpty() bool
```

IsEmpty returns true if the sequence is empty.

<a name="Sequence[K, V].IsNotEmpty"></a>
### Sequence[K, V].IsNotEmpty

```go
func (s Sequence[K, V]) IsNotEmpty() bool
```

IsNotEmpty returns true if the sequence is not empty.

<a name="Sequence[K, V].Keys"></a>
### Sequence[K, V].Keys

```go
func (s Sequence[K, V]) Keys() iter.Seq[K]
```

Keys returns a sequence of keys.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))
	keys := sequence.Keys()

	seq.ForEach(keys, func(k int) {
		fmt.Println(k)
	})
}
```

**Output**

```
0
1
2
```


</details>

<a name="Sequence[K, V].Limit"></a>
### Sequence[K, V].Limit

```go
func (s Sequence[K, V]) Limit(n int) Sequence[K, V]
```

Limit is an alias for Take.

<a name="Sequence[K, V].None"></a>
### Sequence[K, V].None

```go
func (s Sequence[K, V]) None(predicate Predicate[K, V]) bool
```

None returns true if no element satisfies the predicate.

<a name="Sequence[K, V].NotContains"></a>
### Sequence[K, V].NotContains

```go
func (s Sequence[K, V]) NotContains(key K) bool
```

NotContains returns true if the sequence does not contain the key.

<a name="Sequence[K, V].NotContainsPair"></a>
### Sequence[K, V].NotContainsPair

```go
func (s Sequence[K, V]) NotContainsPair(key K, value V) bool
```

NotContainsPair returns true if the sequence does not contain the key\-value pair.

<a name="Sequence[K, V].NotContainsValue"></a>
### Sequence[K, V].NotContainsValue

```go
func (s Sequence[K, V]) NotContainsValue(value V) bool
```

NotContainsValue returns true if the sequence does not contain the value.

<a name="Sequence[K, V].Offset"></a>
### Sequence[K, V].Offset

```go
func (s Sequence[K, V]) Offset(n int) Sequence[K, V]
```

Offset is an alias for Skip.

<a name="Sequence[K, V].Prepend"></a>
### Sequence[K, V].Prepend

```go
func (s Sequence[K, V]) Prepend(key K, value V) Sequence[K, V]
```

Prepend prepends a key\-value pair to the beginning of a sequence.

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
	sequence := seq2.AsSequence(seq2.Of("b", "c", "d"))
	prepended := sequence.Prepend(-1, "a")

	var result []string
	prepended.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
-1:a, 0:b, 1:c, 2:d
```


</details>

<a name="Sequence[K, V].Repeat"></a>
### Sequence[K, V].Repeat

```go
func (s Sequence[K, V]) Repeat(count int) Sequence[K, V]
```

Repeat is an alias for Cycle.

<a name="Sequence[K, V].Reverse"></a>
### Sequence[K, V].Reverse

```go
func (s Sequence[K, V]) Reverse() Sequence[K, V]
```

Reverse returns a new sequence with elements in reverse order.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))
	reversed := sequence.Reverse()

	var result []string
	reversed.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
2:c, 1:b, 0:a
```


</details>

<a name="Sequence[K, V].Skip"></a>
### Sequence[K, V].Skip

```go
func (s Sequence[K, V]) Skip(n int) Sequence[K, V]
```

Skip returns a new sequence that skips the first n elements.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))
	skipped := sequence.Skip(2)

	var result []string
	skipped.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
2:c, 3:d, 4:e
```


</details>

<a name="Sequence[K, V].SortComparingKeys"></a>
### Sequence[K, V].SortComparingKeys

```go
func (s Sequence[K, V]) SortComparingKeys(compare func(K, K) int) Sequence[K, V]
```

SortComparingKeys sorts the sequence by keys using a custom comparator.

<a name="Sequence[K, V].SortComparingValues"></a>
### Sequence[K, V].SortComparingValues

```go
func (s Sequence[K, V]) SortComparingValues(compare func(V, V) int) Sequence[K, V]
```

SortComparingValues sorts the sequence by values using a custom comparator.

<a name="Sequence[K, V].Take"></a>
### Sequence[K, V].Take

```go
func (s Sequence[K, V]) Take(n int) Sequence[K, V]
```

Take returns a new sequence with only the first n elements.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c", "d", "e"))
	taken := sequence.Take(3)

	var result []string
	taken.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
0:a, 1:b, 2:c
```


</details>

<a name="Sequence[K, V].Tap"></a>
### Sequence[K, V].Tap

```go
func (s Sequence[K, V]) Tap(consumer Consumer[K, V]) Sequence[K, V]
```

Tap calls the consumer for each element without changing the sequence.

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
	var result []string
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c")).
		Tap(func(k int, v string) {
			result = append(result, fmt.Sprintf("%d:%s", k, v))
		})

	sequence.Flush()
	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
0:a, 1:b, 2:c
```


</details>

<a name="Sequence[K, V].ToMap"></a>
### Sequence[K, V].ToMap

```go
func (s Sequence[K, V]) ToMap(m map[K]V)
```

ToMap converts the sequence to a map.

<a name="Sequence[K, V].Union"></a>
### Sequence[K, V].Union

```go
func (s Sequence[K, V]) Union(other Sequence[K, V]) Sequence[K, V]
```

Union returns a new sequence that contains all distinct elements from both input sequences.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	sq1 := seq2.AsSequence(seq2.Concat(seq2.Pair("a", 1), seq2.Pair("b", 2), seq2.Pair("c", 3)))
	sq2 := seq2.AsSequence(seq2.Concat(seq2.Pair("a", 1), seq2.Pair("d", 4), seq2.Pair("e", 5)))

	union := sq1.Union(sq2)

	union.ForEach(func(k string, v int) {
		fmt.Println(k, ":", v)
	})

}
```

**Output**

```
a : 1
b : 2
c : 3
d : 4
e : 5
```


</details>

<a name="Sequence[K, V].UnionAll"></a>
### Sequence[K, V].UnionAll

```go
func (s Sequence[K, V]) UnionAll(other Sequence[K, V]) Sequence[K, V]
```

UnionAll returns a new sequence that contains all elements from both input sequences.

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
	seq1 := seq2.AsSequence(seq2.Of("a", "b", "c"))
	seq2Inst := seq2.AsSequence(seq2.Of("c", "d", "e"))

	unionAll := seq1.UnionAll(seq2Inst)

	var result []string
	unionAll.ForEach(func(k int, v string) {
		result = append(result, fmt.Sprintf("%d:%s", k, v))
	})

	fmt.Println(strings.Join(result, ", "))
}
```

**Output**

```
0:a, 1:b, 2:c, 0:c, 1:d, 2:e
```


</details>

<a name="Sequence[K, V].Uniq"></a>
### Sequence[K, V].Uniq

```go
func (s Sequence[K, V]) Uniq() Sequence[K, V]
```

Uniq returns a sequence with unique key\-value pairs.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq2"
)

func main() {
	// Using FromMap to create a sequence with duplicate values
	sequence := seq2.AsSequence(seq2.Of("a", "a", "b", "b", "c"))
	unique := sequence.UniqValues()

	unique.ForEach(func(k int, v string) {
		fmt.Println(k, ":", v)
	})

}
```

**Output**

```
0 : a
2 : b
4 : c
```


</details>

<a name="Sequence[K, V].UniqKeys"></a>
### Sequence[K, V].UniqKeys

```go
func (s Sequence[K, V]) UniqKeys() Sequence[K, V]
```

UniqKeys returns a sequence with unique keys.

<a name="Sequence[K, V].UniqValues"></a>
### Sequence[K, V].UniqValues

```go
func (s Sequence[K, V]) UniqValues() Sequence[K, V]
```

UniqValues returns a sequence with unique values.

<a name="Sequence[K, V].Values"></a>
### Sequence[K, V].Values

```go
func (s Sequence[K, V]) Values() iter.Seq[V]
```

Values returns a sequence of values.

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
	sequence := seq2.AsSequence(seq2.Of("a", "b", "c"))
	values := sequence.Values()

	seq.ForEach(values, func(v string) {
		fmt.Println(v)
	})
}
```

**Output**

```
a
b
c
```


</details>

<a name="Sequence[K, V].Where"></a>
### Sequence[K, V].Where

```go
func (s Sequence[K, V]) Where(predicate Predicate[K, V]) Sequence[K, V]
```

Where returns a new sequence with elements that satisfy the predicate.

<a name="ValueMapper"></a>
## type ValueMapper

ValueMapper is a function that takes Value a new Value.

```go
type ValueMapper[V, R any] = func(V) R
```

<a name="ValuePredicate"></a>
## type ValuePredicate

ValuePredicate is a function that is used to filter by value.

```go
type ValuePredicate[E any] = ValueMapper[E, bool]
```
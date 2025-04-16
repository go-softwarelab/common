# slices

```go
import "github.com/go-softwarelab/common/pkg/slices"
```

Package slices provides a comprehensive set of utilities for working with slices in Go applications.

The goal of this package is to offer a rich set of functions for creating, transforming, and manipulating slices, enabling developers to work with collections of data in a functional programming style. The package includes utilities for filtering, mapping, reducing, and sorting slices, as well as combining and partitioning them.

The package is designed to reduce boilerplate code and improve readability by providing a consistent API for common slice operations. It leverages Go's type safety and generics to ensure that operations on slices are both flexible and safe. The functions in this package complement the standard library's slices package with additional functionality while maintaining similar patterns and conventions.



<a name="Count"></a>
## [Count](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/consumer.go#L14>)

```go
func Count[E any](collection []E) int
```

Count returns the number of elements in the collection.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func main() {
	collection := []string{"apple", "banana", "cherry"}

	count := slices.Count(collection)
	fmt.Println(count)

	emptyCollection := []int{}
	emptyCount := slices.Count(emptyCollection)
	fmt.Println(emptyCount)

}
```

**Output**

```
3
0
```


</details>

<a name="Filter"></a>
## [Filter](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/filter.go#L4>)

```go
func Filter[E any, Slice ~[]E](collection Slice, predicate func(item E) bool) Slice
```

Filter returns a new collection that contains only the elements that satisfy the predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func main() {
	collection := []int{1, 2, 3, 4, 5}

	even := slices.Filter(collection, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(even)

}
```

**Output**

```
[2 4]
```


</details>

<a name="FlatMap"></a>
## [FlatMap](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/mapper.go#L18>)

```go
func FlatMap[E any, R any](collection []E, mapper Mapper[E, []R]) []R
```

FlatMap returns new slice where each element is a result of applying mapper to each element of the original slice and flattening the result.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func main() {
	collection := []int{1, 2, 3}

	duplicated := slices.FlatMap(collection, func(v int) []int {
		return []int{v, v}
	})
	fmt.Println(duplicated)

}
```

**Output**

```
[1 1 2 2 3 3]
```


</details>

<a name="Flatten"></a>
## [Flatten](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/mapper.go#L29>)

```go
func Flatten[E any, Slice ~[]E](collection []Slice) Slice
```

Flatten flattens a slice of slices.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func main() {
	collection := [][]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}

	flattened := slices.Flatten(collection)
	fmt.Println(flattened)

}
```

**Output**

```
[a b c d e f]
```


</details>

<a name="ForEach"></a>
## [ForEach](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/consumer.go#L7>)

```go
func ForEach[E any](collection []E, consumer Consumer[E])
```

ForEach applies consumer to each element of the collection.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func main() {
	collection := []int{1, 2, 3}

	slices.ForEach(collection, func(v int) {
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

<a name="Map"></a>
## [Map](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/mapper.go#L7>)

```go
func Map[E any, R any](collection []E, mapper Mapper[E, R]) []R
```

Map returns new slice where each element is a result of applying mapper to each element of the original slice and flattening the result.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func main() {
	collection := []int{1, 2, 3}

	squared := slices.Map(collection, func(v int) int {
		return v * v
	})
	fmt.Println(squared)

}
```

**Output**

```
[1 4 9]
```


</details>

<a name="Reduce"></a>
## [Reduce](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/reducer.go#L4>)

```go
func Reduce[E any, R any](collection []E, accumulator func(agg R, item E) R, initial R) R
```

Reduce applies a function against an accumulator and each element in the slice \(from left to right\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func main() {
	collection := []int{1, 2, 3, 4}

	sum := slices.Reduce(collection, func(agg int, item int) int {
		return agg + item
	}, 0)
	fmt.Println(sum)

}
```

**Output**

```
10
```


</details>

<details>
<summary>Example (Custom Type)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func main() {
	type Product struct {
		Name  string
		Price int
	}

	collection := []Product{
		{Name: "Apple", Price: 1},
		{Name: "Banana", Price: 2},
		{Name: "Cherry", Price: 3},
	}

	totalPrice := slices.Reduce(collection, func(agg int, item Product) int {
		return agg + item.Price
	}, 0)
	fmt.Println(totalPrice)

}
```

**Output**

```
6
```


</details>

<a name="ReduceRight"></a>
## [ReduceRight](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/reducer.go#L13>)

```go
func ReduceRight[E any, R any](collection []E, accumulator func(agg R, item E) R, initial R) R
```

ReduceRight applies a function against an accumulator and each element in the slice \(from right to left\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func main() {
	collection := []string{"a", "b", "c"}

	concatenated := slices.ReduceRight(collection, func(agg string, item string) string {
		return agg + item
	}, "")
	fmt.Println(concatenated)

}
```

**Output**

```
cba
```


</details>

<a name="Uniq"></a>
## [Uniq](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/filter.go#L18>)

```go
func Uniq[E comparable, Slice ~[]E](collection Slice) Slice
```

Uniq returns a collection with only unique elements. The order of result values is determined by the order they occur in the array, only the first occurrence of each element is kept.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slices"
)

func main() {
	collection := []string{"apple", "banana", "apple", "cherry", "banana"}

	unique := slices.Uniq(collection)
	fmt.Println(unique)

}
```

**Output**

```
[apple banana cherry]
```


</details>

<a name="UniqBy"></a>
## [UniqBy](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/filter.go#L36>)

```go
func UniqBy[E any, R comparable, Slice ~[]E](collection Slice, mapper func(E) R) Slice
```

UniqBy returns a collection with only unique elements based on a key. The order of result values is determined by the order they occur in the array, only the first occurrence of each element is kept.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/slices"
)

func main() {
	collection := []string{"Apple", "BANANA", "apple", "Cherry", "banana"}

	unique := slices.UniqBy(collection, func(s string) string {
		return strings.ToLower(s)
	})
	fmt.Println(unique)

}
```

**Output**

```
[Apple BANANA Cherry]
```


</details>

<a name="Consumer"></a>
## type [Consumer](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/consumer.go#L4>)

Consumer is a function that consumes an element of a sequence.

```go
type Consumer[E any] = func(E)
```

<a name="Mapper"></a>
## type [Mapper](<https://github.com/go-softwarelab/common/blob/main/pkg/slices/mapper.go#L4>)

Mapper is a function that maps a value of type T to a value of type R.

```go
type Mapper[T any, R any] = func(T) R
```
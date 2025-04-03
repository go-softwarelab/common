# seqerr

```go
import "github.com/go-softwarelab/common/pkg/seqerr"
```

Package seqerr provides specialized utilities for handling errors when working with iter.Seq in Go applications.

The goal of this package is to simplify error handling in sequence processing pipelines by offering functions that work with iter.Seq2 where the second value represents an error. These utilities automatically break iteration when an error is encountered and propagate it through the processing chain, allowing errors to be collected and handled at the end of the pipeline.

The package includes error\-aware versions of common sequence operations such as mapping, filtering, and reducing, enabling developers to write clean and robust sequence processing code without explicitly handling errors at each step. This approach reduces boilerplate code and improves readability by separating the error handling logic from the business logic.

By integrating seamlessly with the iter.Seq ecosystem, this package provides a consistent way to manage errors across sequence operations, making it easier to build reliable data processing pipelines.



<a name="Collect"></a>
## Collect

```go
func Collect[E any](seq iter.Seq2[E, error]) ([]E, error)
```

Collect collects the elements of the given sequence into a slice.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Collect elements into a new slice
	result, err := seqerr.Collect(sequence)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(result)

}
```

**Output**

```
[1 2 3]
```


</details>

<details>
<summary>Example (With Error)</summary>




```go
package main

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with an error
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			var err error
			if i == 2 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Collect elements into a new slice
	result, err := seqerr.Collect(sequence)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(result)

}
```

**Output**

```
Error: source error
[1]
```


</details>

<a name="Count"></a>
## Count

```go
func Count[E any](seq iter.Seq2[E, error]) (int, error)
```

Count returns the number of elements in the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Count elements in the sequence
	count, err := seqerr.Count(sequence)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(count)

}
```

**Output**

```
3
```


</details>

<details>
<summary>Example (With Error)</summary>




```go
package main

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with an error
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			var err error
			if i == 2 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Count elements in the sequence
	count, err := seqerr.Count(sequence)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(count)

}
```

**Output**

```
Error: source error
0
```


</details>

<a name="Each"></a>
## Each

```go
func Each[E any, C Consumer[E]](seq iter.Seq2[E, error], consumer C) iter.Seq2[E, error]
```

Each returns a sequence that applies the given consumer to each element of the input sequence and pass it further. Each is an alias for Tap. Comparing to ForEach, this is a lazy function and doesn't consume the input sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Use Each (alias for Tap) to print each element while passing it through
	each := seqerr.Each(sequence, func(n int) {
		fmt.Printf("Element: %d\n", n)
	})

	// Collect the elements after processing
	result, err := seqerr.Collect(each)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Result:", result)

}
```

**Output**

```
Element: 1
Element: 2
Element: 3
Result: [1 2 3]
```


</details>

<a name="Filter"></a>
## Filter

```go
func Filter[E any, P Predicate[E]](seq iter.Seq2[E, error], predicate P) iter.Seq2[E, error]
```

Filter returns a new sequence that contains only the elements that satisfy the predicate.

<details>
<summary>Example (Predicate With Error)</summary>




```go
package main

import (
	"fmt"
	"iter"
	"strconv"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	sequence := iter.Seq2[string, error](func(yield func(string, error) bool) {
		for i := range 5 {
			if !yield(fmt.Sprintf("%d", i), nil) {
				break
			}
		}
	})

	// Filter strings that are even numbers when converted to int
	filtered := seqerr.Filter(sequence, func(s string) (bool, error) {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false, err
		}
		return i%2 == 0, nil
	})

	// Print results
	for s, err := range filtered {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(s)
	}

}
```

**Output**

```
0
2
4
```


</details>

<details>
<summary>Example (Predicate Without Error)</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := range 5 {
			yield(i, nil)
		}
	})

	// Filter even numbers
	evenNumbers := seqerr.Filter(sequence, func(n int) bool {
		return n%2 == 0
	})

	// Print results
	for n, err := range evenNumbers {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
0
2
4
```


</details>

<details>
<summary>Example (Source Error)</summary>




```go
package main

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence that produces an error
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := range 5 {
			var err error
			if i == 2 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Filter even numbers
	filtered := seqerr.Filter(sequence, func(n int) bool {
		return n%2 == 0
	})

	// Collect results
	for n, err := range filtered {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
0
Error: source error
```


</details>

<details>
<summary>Example (Validator)</summary>




```go
package main

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := range 5 {
			if !yield(i+1, nil) {
				break
			}
		}
	})

	// Filter using a validator (non-zero values)
	nonZero := seqerr.Filter(sequence, func(n int) error {
		if n%2 == 0 {
			return errors.New("even value not allowed")
		}
		return nil
	})

	// Collect results
	for n, err := range nonZero {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
1
Error: even value not allowed
```


</details>

<a name="FlatMap"></a>
## FlatMap

```go
func FlatMap[E any, R any](seq iter.Seq2[E, error], mapper MapperWithoutError[E, iter.Seq[R]]) iter.Seq2[R, error]
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
	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// FlatMap to create duplicates of each number
	duplicated := seqerr.FlatMap(sequence, func(n int) iter.Seq[int] {
		return seq.Repeat(n, 2)
	})

	// Collect results
	for n, err := range duplicated {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
1
1
2
2
3
3
```


</details>

<a name="FlatMapOrErr"></a>
## FlatMapOrErr

```go
func FlatMapOrErr[E any, R any](seq iter.Seq2[E, error], mapper MapperWithError[E, iter.Seq[R]]) iter.Seq2[R, error]
```

FlatMapOrErr applies a mapper function that can return error to each element of the sequence and flattens the result.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"
	"strconv"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence of strings
	sequence := iter.Seq2[string, error](func(yield func(string, error) bool) {
		for i := 1; i <= 3; i++ {
			var value string
			if i == 2 {
				value = "two"
			} else {
				value = strconv.Itoa(i)
			}
			if !yield(value, nil) {
				break
			}
		}
	})

	// Duplicate integers or return an error if cannot be parsed
	numbers := seqerr.FlatMapOrErr(sequence, func(s string) (iter.Seq[int], error) {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		return seq.Repeat(i, 2), nil
	})

	// Collect results
	for n, err := range numbers {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
1
1
Error: strconv.Atoi: parsing "two": invalid syntax
```


</details>

<a name="Flatten"></a>
## Flatten

```go
func Flatten[Seq iter.Seq2[iter.Seq[E], error], E any](seq Seq) iter.Seq2[E, error]
```

Flatten flattens a sequence of sequences.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence of sequences
	sequence := seq.RangeTo(3)

	seqOfSeq := seqerr.MapSeq(sequence, func(n int) (iter.Seq[int], error) {
		return seq.Repeat(n, 2), nil
	})

	// Flatten the sequence of sequences
	flattened := seqerr.Flatten(seqOfSeq)

	// Collect results
	for n, err := range flattened {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
0
0
1
1
2
2
```


</details>

<a name="FlattenSlices"></a>
## FlattenSlices

```go
func FlattenSlices[Seq iter.Seq2[[]E, error], E any](seq Seq) iter.Seq2[E, error]
```

FlattenSlices flattens a sequence of slices.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence of slices
	sequence := iter.Seq2[[]int, error](func(yield func([]int, error) bool) {
		slices := [][]int{{1, 2}, {3, 4, 5}, {6}}
		for _, slice := range slices {
			if !yield(slice, nil) {
				break
			}
		}
	})

	// Flatten the sequence of slices
	flattened := seqerr.FlattenSlices(sequence)

	// Collect results
	for n, err := range flattened {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
1
2
3
4
5
6
```


</details>

<details>
<summary>Example (With Error)</summary>




```go
package main

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence of slices with an error
	sequence := iter.Seq2[[]int, error](func(yield func([]int, error) bool) {
		if !yield([]int{1, 2}, nil) {
			return
		}
		if !yield(nil, errors.New("slice error")) {
			return
		}
		// This won't be processed due to the error
		if !yield([]int{3, 4}, nil) {
			return
		}
	})

	// Flatten the sequence of slices
	flattened := seqerr.FlattenSlices(sequence)

	// Collect results
	for n, err := range flattened {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
1
2
Error: slice error
```


</details>

<a name="Fold"></a>
## Fold

```go
func Fold[E any](seq iter.Seq2[E, error], accumulator func(agg E, item E) E) (optional.Elem[E], error)
```

Fold applies a function against an accumulator and each element in the sequence \(from left to right\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 5; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Calculate the product of all numbers
	result, err := seqerr.Fold(sequence, func(acc int, item int) int {
		return acc * item
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if result.IsPresent() {
		fmt.Println("Product:", result.MustGet())
	} else {
		fmt.Println("Empty sequence")
	}

}
```

**Output**

```
Product: 120
```


</details>

<details>
<summary>Example (Empty Sequence)</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create an empty sequence
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		// Do nothing
	})

	// Try to fold the empty sequence
	result, err := seqerr.Fold(sequence, func(acc int, item int) int {
		return acc + item
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if result.IsPresent() {
		fmt.Println("Sum:", result.MustGet())
	} else {
		fmt.Println("Empty sequence")
	}

}
```

**Output**

```
Empty sequence
```


</details>

<a name="FoldRight"></a>
## FoldRight

```go
func FoldRight[E any](seq iter.Seq2[E, error], accumulator func(agg E, item E) E) (optional.Elem[E], error)
```

FoldRight applies a function against an accumulator and each element in the sequence \(from right to left\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"
	"strings"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with strings
	sequence := iter.Seq2[string, error](func(yield func(string, error) bool) {
		words := []string{"hello", "world"}
		for _, word := range words {
			if !yield(word, nil) {
				break
			}
		}
	})

	// Concatenate strings right-to-left
	result, err := seqerr.FoldRight(sequence, func(acc string, item string) string {
		return strings.ToUpper(acc) + "-" + item
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if result.IsPresent() {
		fmt.Println(result.MustGet())
	} else {
		fmt.Println("Empty sequence")
	}

}
```

**Output**

```
WORLD-hello
```


</details>

<a name="ForEach"></a>
## ForEach

```go
func ForEach[E any, C Consumer[E]](seq iter.Seq2[E, error], consumer C) error
```

ForEach applies consumer to each element of the input sequence. Comparing to Each, this is not a lazy function and consumes the input sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Use ForEach to process all elements
	sum := 0
	err := seqerr.ForEach(sequence, func(n int) {
		fmt.Printf("Adding: %d\n", n)
		sum += n
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Sum:", sum)

}
```

**Output**

```
Adding: 1
Adding: 2
Adding: 3
Sum: 6
```


</details>

<details>
<summary>Example (With Consumer Error)</summary>




```go
package main

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Use ForEach with a consumer that returns an error
	sum := 0
	err := seqerr.ForEach(sequence, func(n int) error {
		if n == 2 {
			return errors.New("consumer error")
		}
		fmt.Printf("Adding: %d\n", n)
		sum += n
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Sum:", sum)

}
```

**Output**

```
Adding: 1
Error: consumer error
Sum: 1
```


</details>

<details>
<summary>Example (With Error)</summary>




```go
package main

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			var err error
			if i == 2 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Use ForEach to process elements
	sum := 0
	err := seqerr.ForEach(sequence, func(n int) {
		fmt.Printf("Adding: %d\n", n)
		sum += n
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Sum:", sum)

}
```

**Output**

```
Adding: 1
Error: source error
Sum: 1
```


</details>

<a name="Map"></a>
## Map

```go
func Map[E any, R any](seq iter.Seq2[E, error], mapper MapperWithoutError[E, R]) iter.Seq2[R, error]
```

Map applies a mapper function to each element of the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			yield(i, nil)
		}
	})

	// Map numbers to their squares
	squared := seqerr.Map(sequence, func(n int) int {
		return n * n
	})

	// Collect results
	for n, err := range squared {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
1
4
9
```


</details>

<details>
<summary>Example (Source Error)</summary>




```go
package main

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with error
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			var err error
			if i == 2 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Map numbers to their squares
	squared := seqerr.Map(sequence, func(n int) int {
		return n * n
	})

	// Collect results
	for n, err := range squared {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
1
Error: source error
```


</details>

<a name="MapOrErr"></a>
## MapOrErr

```go
func MapOrErr[E any, R any](seq iter.Seq2[E, error], mapper MapperWithError[E, R]) iter.Seq2[R, error]
```

MapOrErr applies a mapper function that can return error to each element of the sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"
	"strconv"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence of strings
	sequence := iter.Seq2[string, error](func(yield func(string, error) bool) {
		for i := 1; i <= 3; i++ {
			var value string
			if i == 2 {
				value = "two"
			} else {
				value = strconv.Itoa(i)
			}
			if !yield(value, nil) {
				break
			}
		}
	})

	// Map strings to integers
	numbers := seqerr.MapOrErr(sequence, func(s string) (int, error) {
		return strconv.Atoi(s)
	})

	// Collect results
	for n, err := range numbers {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
1
Error: strconv.Atoi: parsing "two": invalid syntax
```


</details>

<a name="MapSeq"></a>
## MapSeq

```go
func MapSeq[E any, R any](seq iter.Seq[E], mapper MapperWithError[E, R]) iter.Seq2[R, error]
```

MapSeq applies a mapper function to each element of the sequence. The mapper function can return an error.

<a name="Produce"></a>
## Produce

```go
func Produce[E, A any](next func(A) ([]E, A, error)) iter.Seq2[[]E, error]
```

Produce returns a new sequence that is filled by the results of calling the next function.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"strconv"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	sequence := seqerr.Produce(func(i int) ([]string, int, error) {
		if i == 2 {
			return []string{}, i + 1, nil
		}

		num := strconv.Itoa(i)
		result := []string{"a" + num, "b" + num, "c" + num}
		return result, i + 1, nil
	})

	for item, err := range sequence {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(item)
	}

}
```

**Output**

```
[a0 b0 c0]
[a1 b1 c1]
```


</details>

<a name="ProduceWithArg"></a>
## ProduceWithArg

```go
func ProduceWithArg[E, A any](next func(A) ([]E, A, error), arg A) iter.Seq2[[]E, error]
```

ProduceWithArg returns a new sequence that is filled by the results of calling the next function with the provided argument.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"strconv"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	sequence := seqerr.ProduceWithArg(func(i int) ([]string, int, error) {
		if i == 2 {
			return []string{}, i + 1, nil
		}

		num := strconv.Itoa(i)
		result := []string{"a" + num, "b" + num, "c" + num}
		return result, i + 1, nil
	}, 1)

	for item, err := range sequence {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(item)
	}

}
```

**Output**

```
[a1 b1 c1]
```


</details>

<a name="Reduce"></a>
## Reduce

```go
func Reduce[E any, R any](seq iter.Seq2[E, error], accumulator func(agg R, item E) R, initial R) (R, error)
```

Reduce applies a function against an accumulator and each element in the sequence \(from left to right\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers 1-5
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 5; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Sum all numbers
	sum, err := seqerr.Reduce(sequence, func(acc int, item int) int {
		return acc + item
	}, 0)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Sum:", sum)

}
```

**Output**

```
Sum: 15
```


</details>

<details>
<summary>Example (With Error)</summary>




```go
package main

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with an error
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 5; i++ {
			var err error
			if i == 3 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Sum all numbers
	sum, err := seqerr.Reduce(sequence, func(acc int, item int) int {
		return acc + item
	}, 0)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Sum:", sum)

}
```

**Output**

```
Error: source error
Sum: 0
```


</details>

<a name="ReduceRight"></a>
## ReduceRight

```go
func ReduceRight[E any, R any](seq iter.Seq2[E, error], accumulator func(agg R, item E) R, initial R) (R, error)
```

ReduceRight applies a function against an accumulator and each element in the sequence \(from right to left\) to reduce it to a single value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with words
	sequence := iter.Seq2[string, error](func(yield func(string, error) bool) {
		words := []string{"hello", "beautiful", "world"}
		for _, word := range words {
			if !yield(word, nil) {
				break
			}
		}
	})

	// Join words right-to-left
	joined, err := seqerr.ReduceRight(sequence, func(acc string, item string) string {
		if acc == "" {
			return item
		}
		return item + " " + acc
	}, "")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(joined)

}
```

**Output**

```
hello beautiful world
```


</details>

<a name="Take"></a>
## Take

```go
func Take[E any](seq iter.Seq2[E, error], n int) iter.Seq2[E, error]
```

Take returns a new sequence that contains only the first n elements of the given sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence of numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := range 10 {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Take only the first 3 elements
	taken := seqerr.Take(sequence, 3)

	// Print results
	for n, err := range taken {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
0
1
2
```


</details>

<a name="TakeUntil"></a>
## TakeUntil

```go
func TakeUntil[E any, P Predicate[E]](seq iter.Seq2[E, error], predicate P) iter.Seq2[E, error]
```

TakeUntil returns a new sequence that takes elements from the given sequence until the predicate is satisfied.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence of numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Take elements until we find one that's equal to 5
	taken := seqerr.TakeUntil(sequence, func(n int) bool {
		return n == 5
	})

	// Print results
	for n, err := range taken {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
0
1
2
3
4
```


</details>

<a name="TakeUntilTrue"></a>
## TakeUntilTrue

```go
func TakeUntilTrue[E any](seq iter.Seq2[E, error], stopCondition func() bool) iter.Seq2[E, error]
```

TakeUntilTrue returns a new sequence that takes elements from the given sequence until the stop condition is satisfied.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence of numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Create a condition that will become true after 5 elements
	count := 0
	condition := func() bool {
		count++
		return count > 5
	}

	// Take elements until the condition becomes true
	taken := seqerr.TakeUntilTrue(sequence, condition)

	// Print results
	for n, err := range taken {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
0
1
2
3
4
```


</details>

<a name="TakeWhile"></a>
## TakeWhile

```go
func TakeWhile[E any, P Predicate[E]](seq iter.Seq2[E, error], predicate P) iter.Seq2[E, error]
```

TakeWhile returns a new sequence that takes elements from the given sequence while the predicate is satisfied.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence of numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Take elements while they are less than 5
	taken := seqerr.TakeWhile(sequence, func(n int) bool {
		return n < 5
	})

	// Print results
	for n, err := range taken {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
0
1
2
3
4
```


</details>

<a name="TakeWhileTrue"></a>
## TakeWhileTrue

```go
func TakeWhileTrue[E any](seq iter.Seq2[E, error], stopCondition func() bool) iter.Seq2[E, error]
```

TakeWhileTrue returns a new sequence that takes elements from the given sequence while the stop condition is satisfied.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence of numbers
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Create a condition that will be true only for the first 4 elements
	count := 0
	condition := func() bool {
		count++
		return count <= 4
	}

	// Take elements while the condition remains true
	taken := seqerr.TakeWhileTrue(sequence, condition)

	// Print results
	for n, err := range taken {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(n)
	}

}
```

**Output**

```
0
1
2
3
```


</details>

<a name="Tap"></a>
## Tap

```go
func Tap[E any, C Consumer[E]](seq iter.Seq2[E, error], consumer C) iter.Seq2[E, error]
```

Tap returns a sequence that applies the given consumer to each element of the input sequence and pass it further. In case if consumer returns an error, the sequence stops and pass only the error from consumer further.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Use Tap to print each element while passing it through
	tapped := seqerr.Tap(sequence, func(n int) {
		fmt.Printf("Processing: %d\n", n)
	})

	// Collect the elements after tapping
	result, err := seqerr.Collect(tapped)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Result:", result)

}
```

**Output**

```
Processing: 1
Processing: 2
Processing: 3
Result: [1 2 3]
```


</details>

<details>
<summary>Example (With Error)</summary>




```go
package main

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Use Tap with a consumer that returns an error
	tapped := seqerr.Tap(sequence, func(n int) error {
		if n == 2 {
			return errors.New("consumer error")
		}
		fmt.Printf("Processing: %d\n", n)
		return nil
	})

	// Collect the elements after tapping
	result, err := seqerr.Collect(tapped)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Result:", result)

}
```

**Output**

```
Processing: 1
Error: consumer error
Result: [1]
```


</details>

<a name="ToSlice"></a>
## ToSlice

```go
func ToSlice[Slice ~[]E, E any](seq iter.Seq2[E, error], slice Slice) (Slice, error)
```

ToSlice collects the elements of the given sequence into a slice.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with numbers 1-3
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i, nil) {
				break
			}
		}
	})

	// Collect elements into a slice
	slice, err := seqerr.ToSlice(sequence, make([]int, 0))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(slice)

}
```

**Output**

```
[1 2 3]
```


</details>

<details>
<summary>Example (With Error)</summary>




```go
package main

import (
	"errors"
	"fmt"
	"iter"

	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	// Create a sequence with an error
	sequence := iter.Seq2[int, error](func(yield func(int, error) bool) {
		for i := 1; i <= 3; i++ {
			var err error
			if i == 2 {
				err = errors.New("source error")
			}
			if !yield(i, err) {
				break
			}
		}
	})

	// Collect elements into a slice
	slice, err := seqerr.ToSlice(sequence, make([]int, 0))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(slice)

}
```

**Output**

```
Error: source error
[1]
```


</details>

<a name="Consumer"></a>
## type Consumer

Consumer is a function that is consuming the sequence.

```go
type Consumer[E any] interface {
    // contains filtered or unexported methods
}
```

<a name="ConsumerWithError"></a>
## type ConsumerWithError

ConsumerWithError is a function that takes an element and returns an error.

```go
type ConsumerWithError[E any] = func(E) error
```

<a name="ConsumerWithoutError"></a>
## type ConsumerWithoutError

ConsumerWithoutError is a function that takes an element and returns nothing.

```go
type ConsumerWithoutError[E any] = func(E)
```

<a name="Mapper"></a>
## type Mapper

Mapper is a function that takes an element and returns a result. It can return an error.

```go
type Mapper[E any, R any] interface {
    // contains filtered or unexported methods
}
```

<a name="MapperWithError"></a>
## type MapperWithError

MapperWithError is a function that takes an element and returns a result and an error.

```go
type MapperWithError[E any, R any] = func(E) (R, error)
```

<a name="MapperWithoutError"></a>
## type MapperWithoutError

MapperWithoutError is a function that takes an element and returns a result.

```go
type MapperWithoutError[E any, R any] = func(E) R
```

<a name="Predicate"></a>
## type Predicate

Predicate is a function that takes an element and returns a boolean. It can return an error.

```go
type Predicate[E any] interface {
    // contains filtered or unexported methods
}
```

<a name="PredicateWithError"></a>
## type PredicateWithError

PredicateWithError is a function that takes an element and returns a boolean, it can fail with error.

```go
type PredicateWithError[E any] = func(E) (bool, error)
```

<a name="PredicateWithoutError"></a>
## type PredicateWithoutError

PredicateWithoutError is a function that takes an element and returns a boolean.

```go
type PredicateWithoutError[E any] = func(E) bool
```

<a name="Validator"></a>
## type Validator

Validator is a function that takes an element and returns an error.

```go
type Validator[E any] = func(E) error
```
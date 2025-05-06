# optional

```go
import "github.com/go-softwarelab/common/pkg/optional"
```



## Variables

<a name="ValueNotPresent"></a>ValueNotPresent is the error returned or passed to iter.Seq2 when the value is not present.

```go
var ValueNotPresent = errors.New(valueNotPresentErrorMessage)
```

<a name="Value"></a>
## type [Value](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L19-L21>)

Value represents an optional value.

```go
type Value[V any] struct {
    // contains filtered or unexported fields
}
```

<a name="Empty"></a>
### [Empty](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L24>)

```go
func Empty[V any]() Value[V]
```

Empty returns an empty optional value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Empty[string]()
	fmt.Println("Is empty:", opt.IsEmpty())
	fmt.Println("Is present:", opt.IsPresent())

}
```

**Output**

```
Is empty: true
Is present: false
```


</details>

<a name="Map"></a>
### [Map](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional_funcs.go#L4>)

```go
func Map[E, R any](o Value[E], f func(E) R) Value[R]
```

Map is a function that maps the value of optional if it is present.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	// Map a present value
	opt := optional.Of("hello")

	// Map to uppercase
	upperOpt := optional.Map(opt, strings.ToUpper)
	fmt.Println("Mapped value:", upperOpt.MustGet())

	// Map with more complex function
	lenOpt := optional.Map(opt, func(s string) int {
		return len(s)
	})
	fmt.Println("String length:", lenOpt.MustGet())

}
```

**Output**

```
Mapped value: HELLO
String length: 5
```


</details>

<details>
<summary>Example (Chain)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	// Chaining multiple Map operations
	opt := optional.Of(42)

	msg := optional.Map(opt, func(n int) string {
		return fmt.Sprintf("Number: %d", n)
	})

	result := optional.Map(msg, func(s string) []byte {
		return []byte(s)
	})

	// Check if result is present
	if result.IsPresent() {
		fmt.Printf("Result type: %T\n", result.MustGet())
		fmt.Printf("Result length: %d\n", len(result.MustGet()))
	}

}
```

**Output**

```
Result type: []uint8
Result length: 10
```


</details>

<details>
<summary>Example (Empty)</summary>




```go
package main

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	// Map an empty optional
	emptyOpt := optional.Empty[string]()

	// Map function won't be called for empty optionals
	mapped := optional.Map(emptyOpt, strings.ToUpper)

	fmt.Println("Is mapped empty:", mapped.IsEmpty())
	fmt.Println("Is mapped present:", mapped.IsPresent())

}
```

**Output**

```
Is mapped empty: true
Is mapped present: false
```


</details>

<a name="None"></a>
### [None](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L30>)

```go
func None[V any]() Value[V]
```

None returns an empty optional value. alias: Empty

<a name="Of"></a>
### [Of](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L43>)

```go
func Of[E any](v E) Value[E]
```

Of returns an optional with the given value. If the value is a pointer, and it's nil, it returns an empty optional. Otherwise, it returns non\-empty optional with the given value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	// With a value
	opt := optional.Of("hello")
	fmt.Println("Value:", opt.MustGet())
	fmt.Println("Is empty:", opt.IsEmpty())

	// With a nil pointer
	var ptr *string = nil
	optPtr := optional.Of(ptr)
	fmt.Println("Nil pointer optional is empty:", optPtr.IsEmpty())

}
```

**Output**

```
Value: hello
Is empty: false
Nil pointer optional is empty: true
```


</details>

<a name="OfPtr"></a>
### [OfPtr](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L55>)

```go
func OfPtr[E any](v *E) Value[E]
```

OfPtr returns an optional with the value from pointer. If the pointer is nil, it returns an empty optional. Otherwise, it returns non\-empty optional with the value pointed to by the pointer.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	// With a value pointer
	value := "hello"
	opt := optional.OfPtr(&value)
	fmt.Println("Value:", opt.MustGet())

	// With a nil pointer
	var nilPtr *string
	optNil := optional.OfPtr(nilPtr)
	fmt.Println("Is empty:", optNil.IsEmpty())

}
```

**Output**

```
Value: hello
Is empty: true
```


</details>

<a name="OfValue"></a>
### [OfValue](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L68>)

```go
func OfValue[E comparable](v E) Value[E]
```

OfValue returns an optional for the given value. If value is zero value, it returns an empty optional. Otherwise, it returns non\-empty optional with the given value.

If zero value is valid existing value for you, for example when the value is int, then prefer Of\(\) instead.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	// Non-zero value
	opt := optional.OfValue(42)
	fmt.Println("Value present:", opt.IsPresent())
	fmt.Println("Value:", opt.MustGet())

	// Zero value
	optZero := optional.OfValue(0)
	fmt.Println("Zero value present:", optZero.IsPresent())

}
```

**Output**

```
Value present: true
Value: 42
Zero value present: false
```


</details>

<a name="Some"></a>
### [Some](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L36>)

```go
func Some[V any](v V) Value[V]
```

Some returns an optional with the given value. It doesn't make any checks on value \- it was caller decision to understand this value as present.

<a name="Value[V].IfNotPresent"></a>
### [Value\[V\].IfNotPresent](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L161>)

```go
func (o Value[V]) IfNotPresent(fn func())
```

IfNotPresent executes the function if the value is not present.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	opt.IfNotPresent(func() {
		fmt.Println("This won't be printed")
	})

	empty.IfNotPresent(func() {
		fmt.Println("This executes when empty")
	})

}
```

**Output**

```
This executes when empty
```


</details>

<a name="Value[V].IfPresent"></a>
### [Value\[V\].IfPresent](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L154>)

```go
func (o Value[V]) IfPresent(fn func(V))
```

IfPresent executes the function if the value is present.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	opt.IfPresent(func(value string) {
		fmt.Println("Value is present:", value)
	})

	empty.IfPresent(func(value string) {
		fmt.Println("This won't be printed")
	})

}
```

**Output**

```
Value is present: hello
```


</details>

<a name="Value[V].IsEmpty"></a>
### [Value\[V\].IsEmpty](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L168>)

```go
func (o Value[V]) IsEmpty() bool
```

IsEmpty returns true if the value is not present.

<a name="Value[V].IsNotEmpty"></a>
### [Value\[V\].IsNotEmpty](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L178>)

```go
func (o Value[V]) IsNotEmpty() bool
```

IsNotEmpty returns true if the value is present.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	fmt.Println("First is not empty:", opt.IsNotEmpty())
	fmt.Println("Second is not empty:", empty.IsNotEmpty())

}
```

**Output**

```
First is not empty: true
Second is not empty: false
```


</details>

<a name="Value[V].IsPresent"></a>
### [Value\[V\].IsPresent](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L173>)

```go
func (o Value[V]) IsPresent() bool
```

IsPresent returns true if the value is present.

<a name="Value[V].MustGet"></a>
### [Value\[V\].MustGet](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L95>)

```go
func (o Value[V]) MustGet() V
```

MustGet returns the value if present, otherwise panics.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Of("hello")
	fmt.Println("Value:", opt.MustGet())

	// Note: Using MustGet on empty optional would panic
	// empty := optional.Empty[string]()
	// empty.MustGet() // would panic with "value is not present"

}
```

**Output**

```
Value: hello
```


</details>

<a name="Value[V].MustGetf"></a>
### [Value\[V\].MustGetf](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L100>)

```go
func (o Value[V]) MustGetf(msg string, args ...any) V
```

MustGetf returns the value if present, otherwise panics with a custom message.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Of("hello")
	fmt.Println("Value:", opt.MustGetf("Custom error: %s", "not found"))

	// Note: Using MustGetf on empty optional would panic with custom message
	// empty := optional.Empty[string]()
	// empty.MustGetf("Custom error: %s", "not found") // would panic with "Custom error: not found"

}
```

**Output**

```
Value: hello
```


</details>

<a name="Value[V].Or"></a>
### [Value\[V\].Or](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L77>)

```go
func (o Value[V]) Or(other Value[V]) Value[V]
```

Or returns this optional if present, otherwise returns the other optional.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt1 := optional.Of("first")
	opt2 := optional.Of("second")
	empty := optional.Empty[string]()

	// Present optional takes precedence
	fmt.Println("First or second:", opt1.Or(opt2).MustGet())
	fmt.Println("Empty or second:", empty.Or(opt2).MustGet())

}
```

**Output**

```
First or second: first
Empty or second: second
```


</details>

<a name="Value[V].OrElse"></a>
### [Value\[V\].OrElse](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L118>)

```go
func (o Value[V]) OrElse(defaultValue V) V
```

OrElse returns the value if present, otherwise returns the default value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	fmt.Println("Present value:", opt.OrElse("default"))
	fmt.Println("Empty value:", empty.OrElse("default"))

}
```

**Output**

```
Present value: hello
Empty value: default
```


</details>

<a name="Value[V].OrElseGet"></a>
### [Value\[V\].OrElseGet](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L127>)

```go
func (o Value[V]) OrElseGet(defaultValue func() V) V
```

OrElseGet returns the value if present, otherwise returns the default value from the function.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	counter := 0
	getDefault := func() string {
		counter++
		return fmt.Sprintf("default-%d", counter)
	}

	fmt.Println("Present value:", opt.OrElseGet(getDefault))
	fmt.Println("Empty value:", empty.OrElseGet(getDefault))
	fmt.Println("Empty value again:", empty.OrElseGet(getDefault))

}
```

**Output**

```
Present value: hello
Empty value: default-1
Empty value again: default-2
```


</details>

<a name="Value[V].OrError"></a>
### [Value\[V\].OrError](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L136>)

```go
func (o Value[V]) OrError(err error) (V, error)
```

OrError returns the value if present, otherwise returns the error.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Of(42)
	empty := optional.Empty[int]()

	val1, err1 := opt.OrError(fmt.Errorf("not found"))
	fmt.Println("Value:", val1)
	fmt.Println("Error:", err1)

	val2, err2 := empty.OrError(fmt.Errorf("not found"))
	fmt.Println("Empty value:", val2)
	fmt.Println("Empty error:", err2 != nil)

}
```

**Output**

```
Value: 42
Error: <nil>
Empty value: 0
Empty error: true
```


</details>

<a name="Value[V].OrErrorGet"></a>
### [Value\[V\].OrErrorGet](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L145>)

```go
func (o Value[V]) OrErrorGet(err func() error) (V, error)
```

OrErrorGet returns the value if present, otherwise returns the error from the function.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Of(42)
	empty := optional.Empty[int]()

	errCounter := 0
	getError := func() error {
		errCounter++
		return fmt.Errorf("not found-%d", errCounter)
	}

	val1, err1 := opt.OrErrorGet(getError)
	fmt.Println("Value:", val1)
	fmt.Println("Error:", err1)

	val2, err2 := empty.OrErrorGet(getError)
	fmt.Println("Empty value:", val2)
	fmt.Println("Empty error:", err2)

}
```

**Output**

```
Value: 42
Error: <nil>
Empty value: 0
Empty error: not found-1
```


</details>

<a name="Value[V].OrZeroValue"></a>
### [Value\[V\].OrZeroValue](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L109>)

```go
func (o Value[V]) OrZeroValue() V
```

OrZeroValue returns the value if present, otherwise returns the zero value of the type.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Of(42)
	empty := optional.Empty[int]()

	fmt.Println("Present value:", opt.OrZeroValue())
	fmt.Println("Empty value:", empty.OrZeroValue())

}
```

**Output**

```
Present value: 42
Empty value: 0
```


</details>

<a name="Value[V].Seq"></a>
### [Value\[V\].Seq](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L183>)

```go
func (o Value[V]) Seq() iter.Seq[V]
```

Seq returns the sequence with yelded value if present, otherwise returns an empty sequence.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	opt := optional.Of("hello")

	var values []string
	seq.ForEach(opt.Seq(), func(value string) {
		values = append(values, value)
	})

	fmt.Println("Values:", values)

	empty := optional.Empty[string]()
	var emptyValues []string
	seq.ForEach(empty.Seq(), func(value string) {
		emptyValues = append(emptyValues, value)
	})

	fmt.Println("Empty values length:", len(emptyValues))

}
```

**Output**

```
Values: [hello]
Empty values length: 0
```


</details>

<a name="Value[V].Seq2"></a>
### [Value\[V\].Seq2](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L193>)

```go
func (o Value[V]) Seq2() iter.Seq2[V, error]
```

Seq2 returns the iter.Seq2\[V, error\] with yelded value if present, otherwise yields an error. Useful with usage of seqerr package.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
	"github.com/go-softwarelab/common/pkg/seqerr"
)

func main() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	err := seqerr.ForEach(opt.Seq2(), func(value string) {
		fmt.Printf("Value: %s\n", value)
	})
	if err != nil {
		panic(err)
	}

	// With empty value
	err = seqerr.ForEach(empty.Seq2(), func(value string) {
		fmt.Printf("Unexpected value: %s\n", value)
	})
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
}
```

**Output**

```
Value: hello
Expected error: value is not present
```


</details>

<a name="Value[V].ShouldGet"></a>
### [Value\[V\].ShouldGet](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L86>)

```go
func (o Value[V]) ShouldGet() (V, error)
```

ShouldGet returns the value if present, otherwise returns the error ValueNotPresent.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
)

func main() {
	opt := optional.Of(42)
	empty := optional.Empty[int]()

	val1, err1 := opt.ShouldGet()
	fmt.Println("Value:", val1)
	fmt.Println("Error:", err1)

	val2, err2 := empty.ShouldGet()
	fmt.Println("Empty value:", val2)
	fmt.Println("Empty error:", err2)

}
```

**Output**

```
Value: 42
Error: <nil>
Empty value: 0
Empty error: value is not present
```


</details>
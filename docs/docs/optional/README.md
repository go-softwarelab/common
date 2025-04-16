# optional

```go
import "github.com/go-softwarelab/common/pkg/optional"
```



## Variables

<a name="ValueNotPresent"></a>ValueNotPresent is the error returned or passed to iter.Seq2 when the value is not present.

```go
var ValueNotPresent = errors.New(valueNotPresentErrorMessage)
```

<a name="Elem"></a>
## type [Elem](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L19-L21>)

Elem represents an optional value.

```go
type Elem[E any] struct {
    // contains filtered or unexported fields
}
```

<a name="Empty"></a>
### [Empty](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L24>)

```go
func Empty[E any]() Elem[E]
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

<a name="Of"></a>
### [Of](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L31>)

```go
func Of[E any](v E) Elem[E]
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
### [OfPtr](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L43>)

```go
func OfPtr[E any](v *E) Elem[E]
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
### [OfValue](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L56>)

```go
func OfValue[E comparable](v E) Elem[E]
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

<a name="Elem[E].IfNotPresent"></a>
### [Elem\[E\].IfNotPresent](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L140>)

```go
func (o Elem[E]) IfNotPresent(fn func())
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

<a name="Elem[E].IfPresent"></a>
### [Elem\[E\].IfPresent](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L133>)

```go
func (o Elem[E]) IfPresent(fn func(E))
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

<a name="Elem[E].IsEmpty"></a>
### [Elem\[E\].IsEmpty](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L147>)

```go
func (o Elem[E]) IsEmpty() bool
```

IsEmpty returns true if the value is not present.

<a name="Elem[E].IsNotEmpty"></a>
### [Elem\[E\].IsNotEmpty](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L157>)

```go
func (o Elem[E]) IsNotEmpty() bool
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

<a name="Elem[E].IsPresent"></a>
### [Elem\[E\].IsPresent](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L152>)

```go
func (o Elem[E]) IsPresent() bool
```

IsPresent returns true if the value is present.

<a name="Elem[E].MustGet"></a>
### [Elem\[E\].MustGet](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L83>)

```go
func (o Elem[E]) MustGet() E
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

<a name="Elem[E].MustGetf"></a>
### [Elem\[E\].MustGetf](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L88>)

```go
func (o Elem[E]) MustGetf(msg string, args ...any) E
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

<a name="Elem[E].Or"></a>
### [Elem\[E\].Or](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L65>)

```go
func (o Elem[E]) Or(other Elem[E]) Elem[E]
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

<a name="Elem[E].OrElse"></a>
### [Elem\[E\].OrElse](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L97>)

```go
func (o Elem[E]) OrElse(defaultValue E) E
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

<a name="Elem[E].OrElseGet"></a>
### [Elem\[E\].OrElseGet](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L106>)

```go
func (o Elem[E]) OrElseGet(defaultValue func() E) E
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

<a name="Elem[E].OrError"></a>
### [Elem\[E\].OrError](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L115>)

```go
func (o Elem[E]) OrError(err error) (E, error)
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

<a name="Elem[E].OrErrorGet"></a>
### [Elem\[E\].OrErrorGet](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L124>)

```go
func (o Elem[E]) OrErrorGet(err func() error) (E, error)
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

<a name="Elem[E].ShouldGet"></a>
### [Elem\[E\].ShouldGet](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L74>)

```go
func (o Elem[E]) ShouldGet() (E, error)
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

<a name="Elem[E].ToSeq"></a>
### [Elem\[E\].ToSeq](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L162>)

```go
func (o Elem[E]) ToSeq() iter.Seq[E]
```

ToSeq returns the sequence with yelded value if present, otherwise returns an empty sequence.

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
	seq.ForEach(opt.ToSeq(), func(value string) {
		values = append(values, value)
	})

	fmt.Println("Values:", values)

	empty := optional.Empty[string]()
	var emptyValues []string
	seq.ForEach(empty.ToSeq(), func(value string) {
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

<a name="Elem[E].ToSeq2"></a>
### [Elem\[E\].ToSeq2](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L172>)

```go
func (o Elem[E]) ToSeq2() iter.Seq2[E, error]
```

ToSeq2 returns the iter.Seq2\[E, error\] with yelded value if present, otherwise yields an error. Useful with usage of seqerr package.

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

	err := seqerr.ForEach(opt.ToSeq2(), func(value string) {
		fmt.Printf("Value: %s\n", value)
	})
	if err != nil {
		panic(err)
	}

	// With empty value
	err = seqerr.ForEach(empty.ToSeq2(), func(value string) {
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
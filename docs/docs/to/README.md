# to

```go
import "github.com/go-softwarelab/common/pkg/to"
```

Package to provides a comprehensive set of utility functions for type conversion and transformation in Go applications.

The goal of this package is to offer useful, common converters and mappers that can be easily integrated into functional programming patterns, such as being used as mapper functions with the seq package.

These utilities are designed to reduce boilerplate code and provide a consistent API for common transformation operations while maintaining type safety through generics.



## Variables

<a name="ErrInvalidStringSyntax"></a>ErrInvalidStringSyntax is returned when the string has invalid syntax for conversion to target type.

```go
var ErrInvalidStringSyntax = strconv.ErrSyntax
```

<a name="ErrValueOutOfRange"></a>ErrValueOutOfRange is returned when the value is out of range of the target type.

```go
var ErrValueOutOfRange = fmt.Errorf("%w to convert", strconv.ErrRange)
```

<a name="Any"></a>
## [Any](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L68>)

```go
func Any[T any](value T) any
```

Any casts the value to an any type.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Convert a typed value to any
	val := to.Any(42)
	fmt.Printf("Type: %T\n", val)

	// Convert a string to any
	strVal := to.Any("hello")
	fmt.Printf("Type: %T\n", strVal)

}
```

**Output**

```
Type: int
Type: string
```


</details>

<a name="AtLeast"></a>
## [AtLeast](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L341>)

```go
func AtLeast[T types.Ordered](min T) func(value T) T
```

AtLeast will return a function that will clamp the value to be at least the min value. It's wrapped around NoLessThan function, to make it usable in Map functions.

See Also: NoLessThan

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Creating a function that ensures values are at least 18
	ensureAdult := to.AtLeast(18)

	fmt.Printf("ensureAdult(15) = %d\n", ensureAdult(15))
	fmt.Printf("ensureAdult(21) = %d\n", ensureAdult(21))

}
```

**Output**

```
ensureAdult(15) = 18
ensureAdult(21) = 21
```


</details>

<a name="AtMost"></a>
## [AtMost](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L359>)

```go
func AtMost[T types.Ordered](max T) func(value T) T
```

AtMost will return a function that will clamp the value to be at most the max value. It's wrapped around NoMoreThan function, to make it usable in Map functions.

See Also: NoMoreThan

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Creating a function that ensures values are at most 100
	ensurePercentage := to.AtMost(100)

	fmt.Printf("ensurePercentage(50) = %d\n", ensurePercentage(50))
	fmt.Printf("ensurePercentage(150) = %d\n", ensurePercentage(150))

}
```

**Output**

```
ensurePercentage(50) = 50
ensurePercentage(150) = 100
```


</details>

<a name="BoolFromNumber"></a>
## [BoolFromNumber](<https://github.com/go-softwarelab/common/blob/main/pkg/to/booleans.go#L21>)

```go
func BoolFromNumber[V types.Number](value V) bool
```

BoolFromNumber will convert any number to bool 0 is false, any other number is true.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting various numeric values to bool
	fmt.Println("0 ->", to.BoolFromNumber(0))
	fmt.Println("1 ->", to.BoolFromNumber(1))
	fmt.Println("-1 ->", to.BoolFromNumber(-1))
	fmt.Println("42 ->", to.BoolFromNumber(42))
	fmt.Println("0.0 ->", to.BoolFromNumber(0.0))
	fmt.Println("0.1 ->", to.BoolFromNumber(0.1))

}
```

**Output**

```
0 -> false
1 -> true
-1 -> true
42 -> true
0.0 -> false
0.1 -> true
```


</details>

<a name="BoolFromString"></a>
## [BoolFromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/booleans.go#L11>)

```go
func BoolFromString(value string) (bool, error)
```

BoolFromString will convert any string to bool

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting various string representations to bool
	trueValue, err := to.BoolFromString("true")
	fmt.Println("true ->", trueValue, err)

	falseValue, err := to.BoolFromString("false")
	fmt.Println("false ->", falseValue, err)

	oneValue, err := to.BoolFromString("1")
	fmt.Println("1 ->", oneValue, err)

	zeroValue, err := to.BoolFromString("0")
	fmt.Println("0 ->", zeroValue, err)

	invalidValue, err := to.BoolFromString("not-a-bool")
	fmt.Println("not-a-bool ->", invalidValue, err != nil)

}
```

**Output**

```
true -> true <nil>
false -> false <nil>
1 -> true <nil>
0 -> false <nil>
not-a-bool -> false true
```


</details>

<a name="CamelCase"></a>
## [CamelCase](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L90>)

```go
func CamelCase(str string) string
```

CamelCase converts string to camel case.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Convert to camelCase
	val := to.CamelCase("hello world")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.CamelCase("some-kebab-case")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.CamelCase("PascalCaseString")
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("helloWorld")
string("someKebabCase")
string("pascalCaseString")
```


</details>

<a name="Capitalized"></a>
## [Capitalized](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L121>)

```go
func Capitalized(str string) string
```

Capitalized converts the first character of string to upper case and the remaining to lower case.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Capitalize the first letter of each sentence
	val := to.Capitalized("hello")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.Capitalized("HELLO WORLD")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.Capitalized("multiple. sentences. here.")
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("Hello")
string("Hello world")
string("Multiple. Sentences. Here.")
```


</details>

<a name="Ellipsis"></a>
## [Ellipsis](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L139>)

```go
func Ellipsis(str string, length int) string
```

Ellipsis trims and truncates a string to a specified length and appends an ellipsis if truncated.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Truncate string with ellipsis
	val := to.Ellipsis("This is a short string", 10)
	fmt.Printf("%T(%q)\n", val, val)

	val = to.Ellipsis("Short", 10)
	fmt.Printf("%T(%q)\n", val, val)

	val = to.Ellipsis("Little bit longer one, but with small length", 3)
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("This is...")
string("Short")
string("...")
```


</details>

<a name="EllipsisWith"></a>
## [EllipsisWith](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L155>)

```go
func EllipsisWith(length int) func(str string) string
```

EllipsisWith returns a function that trims and truncates a string to a specified length and appends an ellipsis if truncated. It's wrapped around Ellipsis function, to make it usable in Map functions. \#mapper

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Create a function that truncates strings to 10 characters
	truncate := to.EllipsisWith(10)

	val := truncate("This is a long string that should be truncated")
	fmt.Printf("%T(%q)\n", val, val)

	val = truncate("Short")
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("This is...")
string("Short")
```


</details>

<a name="EmptyValue"></a>
## [EmptyValue](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L19>)

```go
func EmptyValue[T any]() T
```

EmptyValue returns the zero value of type.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Get zero value of int
	val := to.EmptyValue[int]()
	fmt.Printf("%T(%d)\n", val, val)

	// Get zero value of string
	strVal := to.EmptyValue[string]()
	fmt.Printf("%T(%q)\n", strVal, strVal)

	// Get zero value of bool
	boolVal := to.EmptyValue[bool]()
	fmt.Printf("%T(%v)\n", boolVal, boolVal)

}
```

**Output**

```
int(0)
string("")
bool(false)
```


</details>

<a name="Float32"></a>
## [Float32](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L148>)

```go
func Float32[V types.SignedNumber](value V) (float32, error)
```

Float32 will convert any number to float

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Float32(3)
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

}
```

**Output**

```
float32(3), Error: <nil>
```


</details>

<a name="Float32FromString"></a>
## [Float32FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L304>)

```go
func Float32FromString(value string) (float32, error)
```

Float32FromString will convert any string to float32, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.Float32FromString("3.14159")
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

}
```

**Output**

```
float32(3.14159), Error: <nil>
```


</details>

<a name="Float32FromUnsigned"></a>
## [Float32FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L161>)

```go
func Float32FromUnsigned[V types.Unsigned](value V) (float32, error)
```

Float32FromUnsigned will convert any unassigned number to float

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Float32FromUnsigned(uint(42))
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

}
```

**Output**

```
float32(42), Error: <nil>
```


</details>

<a name="Float64"></a>
## [Float64](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L156>)

```go
func Float64[V types.SignedNumber](value V) (float64, error)
```

Float64 will convert any number to float

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Float64(3)
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

}
```

**Output**

```
float64(3), Error: <nil>
```


</details>

<a name="Float64FromString"></a>
## [Float64FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L317>)

```go
func Float64FromString(value string) (float64, error)
```

Float64FromString will convert any string to float64, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.Float64FromString("3.14159")
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

}
```

**Output**

```
float64(3.14159), Error: <nil>
```


</details>

<a name="Float64FromUnsigned"></a>
## [Float64FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L169>)

```go
func Float64FromUnsigned[V types.Unsigned](value V) (float64, error)
```

Float64FromUnsigned will convert any unassigned number to float

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Float64FromUnsigned(uint64(42))
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

}
```

**Output**

```
float64(42), Error: <nil>
```


</details>

<a name="Int"></a>
## [Int](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L27>)

```go
func Int[V types.SignedNumber](value V) (int, error)
```

Int will convert any integer to int, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Int(int16(42))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
int(42), Error: <nil>
```


</details>

<a name="Int16"></a>
## [Int16](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L61>)

```go
func Int16[V types.SignedNumber](value V) (int16, error)
```

Int16 will convert any integer to int16, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Int16(1000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int16(40000)
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

}
```

**Output**

```
int16(1000), Error: <nil>
int16(0), Error: true
```


</details>

<a name="Int16FromString"></a>
## [Int16FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L200>)

```go
func Int16FromString(value string) (int16, error)
```

Int16FromString will convert any string to int16, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.Int16FromString("12345")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Invalid syntax
	valSyntax, errSyntax := to.Int16FromString("abc")
	fmt.Printf("%T(%d), Error: %v\n", valSyntax, valSyntax, errors.Is(errSyntax, to.ErrInvalidStringSyntax))

}
```

**Output**

```
int16(12345), Error: <nil>
int16(0), Error: true
```


</details>

<a name="Int16FromUnsigned"></a>
## [Int16FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L70>)

```go
func Int16FromUnsigned[V types.Unsigned](value V) (int16, error)
```

Int16FromUnsigned will convert any unsigned integer to int16, with range checks.

<details>
<summary>Example</summary>




```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Int16FromUnsigned(uint(1000))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int16FromUnsigned(uint(40000))
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

}
```

**Output**

```
int16(1000), Error: <nil>
int16(0), Error: true
```


</details>

<a name="Int32"></a>
## [Int32](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L78>)

```go
func Int32[V types.SignedNumber](value V) (int32, error)
```

Int32 will convert any integer to int32, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Int32(1000000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
int32(1000000), Error: <nil>
```


</details>

<a name="Int32FromString"></a>
## [Int32FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L213>)

```go
func Int32FromString(value string) (int32, error)
```

Int32FromString will convert any string to int32, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.Int32FromString("1234567")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
int32(1234567), Error: <nil>
```


</details>

<a name="Int32FromUnsigned"></a>
## [Int32FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L87>)

```go
func Int32FromUnsigned[V types.Unsigned](value V) (int32, error)
```

Int32FromUnsigned will convert any unsigned integer to int32, with range checks.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Int32FromUnsigned(uint(1000000))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
int32(1000000), Error: <nil>
```


</details>

<a name="Int64"></a>
## [Int64](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L95>)

```go
func Int64[V types.SignedNumber](value V) (int64, error)
```

Int64 will convert any integer to int64, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Int64(9223372036854775807)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
int64(9223372036854775807), Error: <nil>
```


</details>

<a name="Int64FromString"></a>
## [Int64FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L226>)

```go
func Int64FromString(value string) (int64, error)
```

Int64FromString will convert any string to int64, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.Int64FromString("9223372036854775807")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
int64(9223372036854775807), Error: <nil>
```


</details>

<a name="Int64FromUnsigned"></a>
## [Int64FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L100>)

```go
func Int64FromUnsigned[V types.Unsigned](value V) (int64, error)
```

Int64FromUnsigned will convert any unsigned integer to int64, with range checks.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Int64FromUnsigned(uint64(9223372036854775807))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
int64(9223372036854775807), Error: <nil>
```


</details>

<a name="Int8"></a>
## [Int8](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L44>)

```go
func Int8[V types.SignedNumber](value V) (int8, error)
```

Int8 will convert any integer to int8, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Int8(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int8(1000)
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

}
```

**Output**

```
int8(42), Error: <nil>
int8(0), Error: true
```


</details>

<a name="Int8FromString"></a>
## [Int8FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L187>)

```go
func Int8FromString(value string) (int8, error)
```

Int8FromString will convert any string to int8, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.Int8FromString("100")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Out of range
	valOOR, errOOR := to.Int8FromString("200")
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

}
```

**Output**

```
int8(100), Error: <nil>
int8(0), Error: true
```


</details>

<a name="Int8FromUnsigned"></a>
## [Int8FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L53>)

```go
func Int8FromUnsigned[V types.Unsigned](value V) (int8, error)
```

Int8FromUnsigned will convert any unsigned integer to int8, with range checks.

<details>
<summary>Example</summary>




```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.Int8FromUnsigned(uint(42))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int8FromUnsigned(uint(200))
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

}
```

**Output**

```
int8(42), Error: <nil>
int8(0), Error: true
```


</details>

<a name="IntFromString"></a>
## [IntFromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L174>)

```go
func IntFromString(value string) (int, error)
```

IntFromString will convert any string to int, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.IntFromString("12345")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Invalid syntax
	valSyntax, errSyntax := to.IntFromString("abc")
	fmt.Printf("%T(%d), Error: %v\n", valSyntax, valSyntax, errors.Is(errSyntax, to.ErrInvalidStringSyntax))

}
```

**Output**

```
int(12345), Error: <nil>
int(0), Error: true
```


</details>

<a name="IntFromUnsigned"></a>
## [IntFromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L36>)

```go
func IntFromUnsigned[V types.Unsigned](value V) (int, error)
```

IntFromUnsigned will convert any unsigned integer to int, with range checks.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.IntFromUnsigned(uint16(42))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
int(42), Error: <nil>
```


</details>

<a name="KebabCase"></a>
## [KebabCase](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L103>)

```go
func KebabCase(str string) string
```

KebabCase converts string to kebab case.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Convert to kebab-case
	val := to.KebabCase("hello world")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.KebabCase("camelCaseString")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.KebabCase("PascalCaseString")
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("hello-world")
string("camel-case-string")
string("pascal-case-string")
```


</details>

<a name="Nil"></a>
## [Nil](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L9>)

```go
func Nil[T any]() *T
```

Nil returns a nil pointer of type.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Create a nil pointer of int type
	val := to.Nil[int]()
	fmt.Printf("%T, IsNil: %v\n", val, val == nil)

	// Create a nil pointer of a struct type
	type Person struct {
		Name string
		Age  int
	}
	person := to.Nil[Person]()
	fmt.Printf("%T, IsNil: %v\n", person, person == nil)

}
```

**Output**

```
*int, IsNil: true
*to_test.Person, IsNil: true
```


</details>

<a name="NilOfType"></a>
## [NilOfType](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L14>)

```go
func NilOfType[T any](_ *T) *T
```

NilOfType returns a nil pointer of type.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Create a nil pointer based on existing pointer type
	var existingPtr *string
	nilPtr := to.NilOfType(existingPtr)
	fmt.Printf("%T, IsNil: %v\n", nilPtr, nilPtr == nil)

}
```

**Output**

```
*string, IsNil: true
```


</details>

<a name="NoLessThan"></a>
## [NoLessThan](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L330>)

```go
func NoLessThan[T types.Ordered](value, min T) T
```

NoLessThan will return the value if it's not less than the min value or the min value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Already above min
	val := to.NoLessThan(5, 0)
	fmt.Printf("NoLessThan(5, 0) = %d\n", val)

	// Below min
	val = to.NoLessThan(-5, 0)
	fmt.Printf("NoLessThan(-5, 0) = %d\n", val)

	// Works with float values too
	floatVal := to.NoLessThan(3.14, 4.0)
	fmt.Printf("NoLessThan(3.14, 4.0) = %.2f\n", floatVal)

}
```

**Output**

```
NoLessThan(5, 0) = 5
NoLessThan(-5, 0) = 0
NoLessThan(3.14, 4.0) = 4.00
```


</details>

<a name="NoMoreThan"></a>
## [NoMoreThan](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L348>)

```go
func NoMoreThan[T types.Ordered](value, max T) T
```

NoMoreThan will return the value if it's not more than the max value or the max value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Below max
	val := to.NoMoreThan(5, 10)
	fmt.Printf("NoMoreThan(5, 10) = %d\n", val)

	// Above max
	val = to.NoMoreThan(15, 10)
	fmt.Printf("NoMoreThan(15, 10) = %d\n", val)

}
```

**Output**

```
NoMoreThan(5, 10) = 5
NoMoreThan(15, 10) = 10
```


</details>

<a name="Options"></a>
## [Options](<https://github.com/go-softwarelab/common/blob/main/pkg/to/options.go#L4>)

```go
func Options[T any](opts ...func(*T)) T
```

Options helper function to handle options pattern.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Prepare definition of config
	type Config struct {
		Host  string
		Port  int
		Debug bool
	}

	// Define reusable option functions
	withHost := func(host string) func(*Config) {
		return func(c *Config) {
			c.Host = host
		}
	}

	withPort := func(port int) func(*Config) {
		return func(c *Config) {
			c.Port = port
		}
	}

	withDebug := func(c *Config) {
		c.Debug = true
	}

	// Handle opts
	config := to.Options[Config](
		withHost("example.com"),
		withPort(443),
		withDebug,
	)

	fmt.Printf("Host: %s\n", config.Host)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)

}
```

**Output**

```
Host: example.com
Port: 443
Debug: true
```


</details>

<a name="OptionsWithDefault"></a>
## [OptionsWithDefault](<https://github.com/go-softwarelab/common/blob/main/pkg/to/options.go#L13>)

```go
func OptionsWithDefault[T any](defaultOptions T, opts ...func(*T)) T
```

OptionsWithDefault helper function to handle options pattern with default value for options.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	type Config struct {
		Host  string
		Port  int
		Debug bool
	}

	// Default configuration
	defaultConfig := Config{
		Host:  "127.0.0.1",
		Port:  3000,
		Debug: false,
	}

	// Override some default values using OptionsWithDefault
	config := to.OptionsWithDefault(defaultConfig,
		func(c *Config) {
			c.Port = 8080
		},
		func(c *Config) {
			c.Debug = true
		},
	)

	fmt.Printf("Host: %s\n", config.Host)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)

}
```

**Output**

```
Host: 127.0.0.1
Port: 8080
Debug: true
```


</details>

<a name="PascalCase"></a>
## [PascalCase](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L81>)

```go
func PascalCase(str string) string
```

PascalCase converts string to pascal case.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Convert to PascalCase
	val := to.PascalCase("hello world")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.PascalCase("some-kebab-case")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.PascalCase("snake_case_string")
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("HelloWorld")
string("SomeKebabCase")
string("SnakeCaseString")
```


</details>

<a name="Ptr"></a>
## [Ptr](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L4>)

```go
func Ptr[T any](x T) *T
```

Ptr returns a pointer copy of value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Create a pointer to a value
	val := to.Ptr(42)
	fmt.Printf("%T(%d)\n", val, *val)

	// Works with strings too
	strVal := to.Ptr("hello")
	fmt.Printf("%T(%q)\n", strVal, *strVal)

}
```

**Output**

```
*int(42)
*string("hello")
```


</details>

<a name="Sentences"></a>
## [Sentences](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L62>)

```go
func Sentences(str string) []string
```

Sentences splits string into an array of its sentences. The sentences are trimmed from leading and trailing spaces.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Splitting into sentences
	text := "Hello, world! This is a test. Is this working?"
	sentences := to.Sentences(text)
	fmt.Printf("%#v\n", sentences)

}
```

**Output**

```
[]string{"Hello, world!", "This is a test.", "Is this working?"}
```


</details>

<a name="SliceOfAny"></a>
## [SliceOfAny](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L73>)

```go
func SliceOfAny[T any](collection []T) []any
```

SliceOfAny casts the slice to a slice of any type.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Convert a slice of ints to a slice of any
	values := []int{1, 2, 3}
	anySlice := to.SliceOfAny(values)

	fmt.Printf("Type: %T, Length: %d\n", anySlice, len(anySlice))
	fmt.Printf("First element type: %T\n", anySlice[0])

}
```

**Output**

```
Type: []interface {}, Length: 3
First element type: int
```


</details>

<a name="SliceOfPtr"></a>
## [SliceOfPtr](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L50>)

```go
func SliceOfPtr[T any](collection []T) []*T
```

SliceOfPtr returns a slice of pointer copy of value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Convert a slice of values to a slice of pointers
	values := []int{1, 2, 3}
	ptrs := to.SliceOfPtr(values)

	// Print first element to demonstrate it's a pointer
	fmt.Printf("%T(%d)\n", ptrs[0], *ptrs[0])

	// Print length to show all elements were converted
	fmt.Printf("Length: %d\n", len(ptrs))

}
```

**Output**

```
*int(1)
Length: 3
```


</details>

<a name="SliceOfValue"></a>
## [SliceOfValue](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L59>)

```go
func SliceOfValue[T any](collection []*T) []T
```

SliceOfValue returns a slice with the pointer values.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Convert a slice of pointers to a slice of values
	p1, p2 := to.Ptr(10), to.Ptr(20)
	ptrs := []*int{p1, p2, nil}
	values := to.SliceOfValue(ptrs)

	fmt.Printf("%T(%d)\n", values[0], values[0])
	fmt.Printf("%T(%d)\n", values[1], values[1])
	fmt.Printf("%T(%d) (zero value from nil)\n", values[2], values[2])

}
```

**Output**

```
int(10)
int(20)
int(0) (zero value from nil)
```


</details>

<a name="SnakeCase"></a>
## [SnakeCase](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L112>)

```go
func SnakeCase(str string) string
```

SnakeCase converts string to snake case.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Convert to snake_case
	val := to.SnakeCase("hello world")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.SnakeCase("camelCaseString")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.SnakeCase("PascalCaseString")
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("hello_world")
string("camel_case_string")
string("pascal_case_string")
```


</details>

<a name="StringFromBool"></a>
## [StringFromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L30>)

```go
func StringFromBool(value bool) string
```

StringFromBool will convert any bool to string

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Boolean to string conversion
	val := to.StringFromBool(true)
	fmt.Printf("%T(%q)\n", val, val)

	val = to.StringFromBool(false)
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("true")
string("false")
```


</details>

<a name="StringFromBytes"></a>
## [StringFromBytes](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L35>)

```go
func StringFromBytes(value []byte) string
```

StringFromBytes will convert any byte slice to string

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Byte slice to string conversion
	val := to.StringFromBytes([]byte("hello"))
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("hello")
```


</details>

<a name="StringFromFloat"></a>
## [StringFromFloat](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L25>)

```go
func StringFromFloat[V types.Float](value V) string
```

StringFromFloat will convert any float to string

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Float to string conversion
	val := to.StringFromFloat(3.14159)
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("3.141590")
```


</details>

<a name="StringFromInteger"></a>
## [StringFromInteger](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L20>)

```go
func StringFromInteger[V types.Integer](value V) string
```

StringFromInteger will convert any integer to string

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Integer to string conversion
	val := to.StringFromInteger(42)
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("42")
```


</details>

<a name="StringFromRune"></a>
## [StringFromRune](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L40>)

```go
func StringFromRune(value rune) string
```

StringFromRune will convert any rune to string

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Rune to string conversion
	val := to.StringFromRune('世')
	fmt.Printf("%T(%q)\n", val, val)

}
```

**Output**

```
string("世")
```


</details>

<a name="UInt"></a>
## [UInt](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L108>)

```go
func UInt[V types.Number](value V) (uint, error)
```

UInt will convert any integer to uint, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting positive value
	val, err := to.UInt(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting negative value
	valNeg, errNeg := to.UInt(-5)
	fmt.Printf("%T(%d), Error: %v\n", valNeg, valNeg, errors.Is(errNeg, to.ErrValueOutOfRange))

}
```

**Output**

```
uint(42), Error: <nil>
uint(0), Error: true
```


</details>

<a name="UInt16"></a>
## [UInt16](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L124>)

```go
func UInt16[V types.Number](value V) (uint16, error)
```

UInt16 will convert any integer to uint16, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.UInt16(65000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
uint16(65000), Error: <nil>
```


</details>

<a name="UInt16FromString"></a>
## [UInt16FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L265>)

```go
func UInt16FromString(value string) (uint16, error)
```

UInt16FromString will convert any string to uint16, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.UInt16FromString("65000")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
uint16(65000), Error: <nil>
```


</details>

<a name="UInt32"></a>
## [UInt32](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L132>)

```go
func UInt32[V types.Number](value V) (uint32, error)
```

UInt32 will convert any integer to uint32, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.UInt32(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Negative number
	valNeg, errNeg := to.UInt32(-5)
	fmt.Printf("%T(%d), Error: %v\n", valNeg, valNeg, errors.Is(errNeg, to.ErrValueOutOfRange))

}
```

**Output**

```
uint32(42), Error: <nil>
uint32(0), Error: true
```


</details>

<a name="UInt32FromString"></a>
## [UInt32FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L278>)

```go
func UInt32FromString(value string) (uint32, error)
```

UInt32FromString will convert any string to uint32, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.UInt32FromString("4294967295")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
uint32(4294967295), Error: <nil>
```


</details>

<a name="UInt64"></a>
## [UInt64](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L140>)

```go
func UInt64[V types.Number](value V) (uint64, error)
```

UInt64 will convert any integer to uint64, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.UInt64(uint(18446744073709551000))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
uint64(18446744073709551000), Error: <nil>
```


</details>

<a name="UInt64FromString"></a>
## [UInt64FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L291>)

```go
func UInt64FromString(value string) (uint64, error)
```

UInt64FromString will convert any string to uint64, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.UInt64FromString("18446744073709551615")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
uint64(18446744073709551615), Error: <nil>
```


</details>

<a name="UInt8"></a>
## [UInt8](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L116>)

```go
func UInt8[V types.Number](value V) (uint8, error)
```

UInt8 will convert any integer to uint8, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting within range
	val, err := to.UInt8(200)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.UInt8(300)
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

}
```

**Output**

```
uint8(200), Error: <nil>
uint8(0), Error: true
```


</details>

<a name="UInt8FromString"></a>
## [UInt8FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L252>)

```go
func UInt8FromString(value string) (uint8, error)
```

UInt8FromString will convert any string to uint8, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.UInt8FromString("200")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
uint8(200), Error: <nil>
```


</details>

<a name="UIntFromString"></a>
## [UIntFromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L239>)

```go
func UIntFromString(value string) (uint, error)
```

UIntFromString will convert any string to uint, with range checks

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Valid conversion
	val, err := to.UIntFromString("42")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
uint(42), Error: <nil>
```


</details>

<a name="Value"></a>
## [Value](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L32>)

```go
func Value[T any](x *T) T
```

Value returns the pointer value or zero value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Get value from a pointer
	ptr := to.Ptr(42)
	val := to.Value(ptr)
	fmt.Printf("%T(%d)\n", val, val)

	// Get value from a nil pointer (returns zero value)
	var nilPtr *string
	strVal := to.Value(nilPtr)
	fmt.Printf("%T(%q)\n", strVal, strVal)

}
```

**Output**

```
int(42)
string("")
```


</details>

<a name="ValueBetween"></a>
## [ValueBetween](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L371>)

```go
func ValueBetween[T types.Ordered](value, min, max T) T
```

ValueBetween will clamp the value between the min and max values. In other words it ensures that result is min \<= value \<= max. For value that is less than min, it will return min. For value that is greater than max, it will return max.

See Also: ValueBetweenThe to use it in Map functions.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Within range
	val := to.ValueBetween(5, 0, 10)
	fmt.Printf("ValueBetween(5, 0, 10) = %d\n", val)

	// Lower than min
	val = to.ValueBetween(-5, 0, 10)
	fmt.Printf("ValueBetween(-5, 0, 10) = %d\n", val)

	// Higher than max
	val = to.ValueBetween(15, 0, 10)
	fmt.Printf("ValueBetween(15, 0, 10) = %d\n", val)

}
```

**Output**

```
ValueBetween(5, 0, 10) = 5
ValueBetween(-5, 0, 10) = 0
ValueBetween(15, 0, 10) = 10
```


</details>

<a name="ValueBetweenThe"></a>
## [ValueBetweenThe](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L387>)

```go
func ValueBetweenThe[T types.Ordered](min, max T) func(value T) T
```

ValueBetweenThe returns a function that clamps the value between the min and max values. In other words it ensures that result is min \<= value \<= max. For value that is less than min, it will return min. For value that is greater than max, it will return max. It's wrapped around ValueBetween function, to make it usable in Map functions.

See Also: ValueBetween

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Creating a percentage validator (0-100)
	validatePercentage := to.ValueBetweenThe(0, 100)

	fmt.Printf("validatePercentage(-20) = %d\n", validatePercentage(-20))
	fmt.Printf("validatePercentage(50) = %d\n", validatePercentage(50))
	fmt.Printf("validatePercentage(150) = %d\n", validatePercentage(150))

}
```

**Output**

```
validatePercentage(-20) = 0
validatePercentage(50) = 50
validatePercentage(150) = 100
```


</details>

<a name="ValueOr"></a>
## [ValueOr](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L41>)

```go
func ValueOr[T any](x *T, fallback T) T
```

ValueOr returns the pointer value or the fallback value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Get value from a pointer
	ptr := to.Ptr(42)
	val := to.ValueOr(ptr, 0)
	fmt.Printf("%T(%d)\n", val, val)

	// Get fallback value from a nil pointer
	var nilPtr *string
	strVal := to.ValueOr(nilPtr, "fallback")
	fmt.Printf("%T(%q)\n", strVal, strVal)

}
```

**Output**

```
int(42)
string("fallback")
```


</details>

<a name="Words"></a>
## [Words](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L45>)

```go
func Words(str string) []string
```

Words splits string into an array of its words.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Splitting into words
	camelCase := to.Words("camelCaseString")
	fmt.Printf("%#v\n", camelCase)

	pascalCase := to.Words("PascalCaseString")
	fmt.Printf("%#v\n", pascalCase)

	withNumbers := to.Words("Int8Value")
	fmt.Printf("%#v\n", withNumbers)

	withSpaces := to.Words("  hello  world  ")
	fmt.Printf("%#v\n", withSpaces)

}
```

**Output**

```
[]string{"camel", "Case", "String"}
[]string{"Pascal", "Case", "String"}
[]string{"Int", "8", "Value"}
[]string{"hello", "world"}
```


</details>

<a name="ZeroValue"></a>
## [ZeroValue](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L26>)

```go
func ZeroValue[T any]() T
```

ZeroValue returns the zero value of type. alias: EmptyValue

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Get zero value of int
	val := to.ZeroValue[int]()
	fmt.Printf("%T(%d)\n", val, val)

	// Get zero value of string
	strVal := to.ZeroValue[string]()
	fmt.Printf("%T(%q)\n", strVal, strVal)

}
```

**Output**

```
int(0)
string("")
```


</details>

<a name="IfElseCondition"></a>
## type [IfElseCondition](<https://github.com/go-softwarelab/common/blob/main/pkg/to/conditional.go#L4-L7>)

IfElseCondition is a struct that provides a fluent API for conditional value mapping.

```go
type IfElseCondition[T any] struct {
    // contains filtered or unexported fields
}
```

<a name="If"></a>
### [If](<https://github.com/go-softwarelab/common/blob/main/pkg/to/conditional.go#L10>)

```go
func If[T any](condition bool, resultProvider func() T) *IfElseCondition[T]
```

If creates a new IfElseCondition with the given condition and result provider.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Using If with function providers
	result := to.If(true, func() string {
		return "condition is true"
	}).Else(func() string {
		return "condition is false"
	})
	fmt.Println(result)

	result = to.If(false, func() string {
		return "condition is true"
	}).Else(func() string {
		return "condition is false"
	})
	fmt.Println(result)

}
```

**Output**

```
condition is true
condition is false
```


</details>

<a name="IfThen"></a>
### [IfThen](<https://github.com/go-softwarelab/common/blob/main/pkg/to/conditional.go#L18>)

```go
func IfThen[T any](condition bool, resultWhenTrue T) *IfElseCondition[T]
```

IfThen creates a new IfElseCondition with the given condition and result.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Using IfThen with direct values
	result := to.IfThen(true, "condition is true").ElseThen("condition is false")
	fmt.Println(result)

	result = to.IfThen(false, "condition is true").ElseThen("condition is false")
	fmt.Println(result)

}
```

**Output**

```
condition is true
condition is false
```


</details>

<a name="IfElseCondition[T].Else"></a>
### [\*IfElseCondition\[T\].Else](<https://github.com/go-softwarelab/common/blob/main/pkg/to/conditional.go#L52>)

```go
func (c *IfElseCondition[T]) Else(resultProvider func() T) T
```

Else accepts the default result provider and returns the result of the condition evaluation.

<a name="IfElseCondition[T].ElseIf"></a>
### [\*IfElseCondition\[T\].ElseIf](<https://github.com/go-softwarelab/common/blob/main/pkg/to/conditional.go#L28>)

```go
func (c *IfElseCondition[T]) ElseIf(condition bool, resultProvider func() T) *IfElseCondition[T]
```

ElseIf adds an else if condition.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Using ElseIf for multiple conditions
	age := 25
	result := to.If(age < 18, func() string {
		return "minor"
	}).ElseIf(age < 65, func() string {
		return "adult"
	}).Else(func() string {
		return "senior"
	})
	fmt.Println(result)

	// Testing with different age
	age = 70
	result = to.If(age < 18, func() string {
		return "minor"
	}).ElseIf(age < 65, func() string {
		return "adult"
	}).Else(func() string {
		return "senior"
	})
	fmt.Println(result)

}
```

**Output**

```
adult
senior
```


</details>

<a name="IfElseCondition[T].ElseIfThen"></a>
### [\*IfElseCondition\[T\].ElseIfThen](<https://github.com/go-softwarelab/common/blob/main/pkg/to/conditional.go#L39>)

```go
func (c *IfElseCondition[T]) ElseIfThen(condition bool, resultWhenTrue T) *IfElseCondition[T]
```

ElseIfThen adds an else if condition with a result.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Using ElseIfThen for multiple conditions with direct values
	score := 85
	grade := to.IfThen(score >= 90, "A").
		ElseIfThen(score >= 80, "B").
		ElseIfThen(score >= 70, "C").
		ElseIfThen(score >= 60, "D").
		ElseThen("F")
	fmt.Println(grade)

}
```

**Output**

```
B
```


</details>

<a name="IfElseCondition[T].ElseThen"></a>
### [\*IfElseCondition\[T\].ElseThen](<https://github.com/go-softwarelab/common/blob/main/pkg/to/conditional.go#L60>)

```go
func (c *IfElseCondition[T]) ElseThen(resultWhenFalse T) T
```

ElseThen accepts the default result and returns the result of the condition evaluation.

<a name="SwitchCase"></a>
## type [SwitchCase](<https://github.com/go-softwarelab/common/blob/main/pkg/to/conditional.go#L68-L75>)

SwitchCase provides a fluent API for conditional \(switch\-like\) value mapping. Represents case predicates.

```go
type SwitchCase[V comparable, R any] interface {
    // Case adds a case that compares equality to case value.
    Case(value V) SwitchThen[V, R]
    // When adds a case predicate function.
    When(func(V) bool) SwitchThen[V, R]
    // Default adds a default value.
    Default(R) R
}
```

<details>
<summary>Example (7hen)</summary>




```go
package main

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Using Switch with When for custom predicates
	text := "HELLO"

	result := to.Switch[string, string](text).
		When(func(s string) bool { return s == "" }).
		ThenValue("Empty string").
		When(func(s string) bool { return len(s) < 5 }).
		ThenValue("Short string").
		When(func(s string) bool { return strings.ToUpper(s) == s }).
		ThenValue("All uppercase").
		Default("Normal string")

	fmt.Println(result)

}
```

**Output**

```
All uppercase
```


</details>

<a name="Switch"></a>
### [Switch](<https://github.com/go-softwarelab/common/blob/main/pkg/to/conditional.go#L86>)

```go
func Switch[V comparable, R any](value V) SwitchCase[V, R]
```

Switch creates a new SwitchCase for the given value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Using Switch with Case for equality comparisons
	dayNum := 3
	dayName := to.Switch[int, string](dayNum).
		Case(1).ThenValue("Monday").
		Case(2).ThenValue("Tuesday").
		Case(3).ThenValue("Wednesday").
		Case(4).ThenValue("Thursday").
		Case(5).ThenValue("Friday").
		Case(6).ThenValue("Saturday").
		Case(7).ThenValue("Sunday").
		Default("Invalid day")

	fmt.Println(dayName)

}
```

**Output**

```
Wednesday
```


</details>

<a name="SwitchThen"></a>
## type [SwitchThen](<https://github.com/go-softwarelab/common/blob/main/pkg/to/conditional.go#L78-L83>)

SwitchThen provides a fluent API for conditional \(switch\-like\) value mapping. Represents case results.

```go
type SwitchThen[V comparable, R any] interface {
    // Then adds a result provider function for given case.
    Then(func(V) R) SwitchCase[V, R]
    // ThenValue adds a result value for given case.
    ThenValue(R) SwitchCase[V, R]
}
```

<details>
<summary>Example (4hen)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Using Switch with Then for computed results
	number := 15

	result := to.Switch[int, string](number).
		Case(0).Then(func(int) string { return "Zero" }).
		When(func(n int) bool { return n%3 == 0 && n%5 == 0 }).Then(func(n int) string {
		return fmt.Sprintf("FizzBuzz: %d", n)
	}).
		Default("Number")

	fmt.Println(result)

}
```

**Output**

```
FizzBuzz: 15
```


</details>

<details>
<summary>Example (4hen Value)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Using Switch with Then for computed results
	number := 15

	result := to.Switch[int, string](number).
		Case(0).ThenValue("Zero").
		When(func(n int) bool { return n%3 == 0 }).ThenValue("Fizz").
		When(func(n int) bool { return n%5 == 0 }).ThenValue("Buzz").
		Default("Number")

	fmt.Println(result)

}
```

**Output**

```
Fizz
```


</details>
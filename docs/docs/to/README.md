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
## [Any](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L77>)

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
## [AtLeast](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1313>)

```go
func AtLeast[T types.Ordered](min T) func(value T) T
```

AtLeast will return a function that will clamp the value to be at least the min value. It's wrapped around NoLessThan function, to make it usable in Map functions.

See Also: NoLessThan @Deprecated: In the future will replace ValueAtLeast, use AtLeastThe instead.

<a name="AtLeastThe"></a>
## [AtLeastThe](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1323>)

```go
func AtLeastThe[T types.Ordered](min T) func(value T) T
```

AtLeastThe will return a function that will clamp the value to be at least the min value. It's wrapped around ValueAtLeast function, to make it usable in Map functions.

See Also: ValueAtLeast

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
	ensureAdult := to.AtLeastThe(18)

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
## [AtMost](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1347>)

```go
func AtMost[T types.Ordered](max T) func(value T) T
```

AtMost will return a function that will clamp the value to be at most the max value. It's wrapped around ValueAtMost function, to make it usable in Map functions.

See Also: ValueAtMost @Deprecated: In the future will replace ValueAtMost, use AtMostThe instead.

<a name="AtMostThe"></a>
## [AtMostThe](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1357>)

```go
func AtMostThe[T types.Ordered](max T) func(value T) T
```

AtMostThe will return a function that will clamp the value to be at most the max value. It's wrapped around ValueAtMost function, to make it usable in Map functions.

See Also: ValueAtMost

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
	ensurePercentage := to.AtMostThe(100)

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
## [BoolFromNumber](<https://github.com/go-softwarelab/common/blob/main/pkg/to/booleans.go#L22>)

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
## [BoolFromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/booleans.go#L12>)

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
## [CamelCase](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L119>)

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
## [Capitalized](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L150>)

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
## [Ellipsis](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L168>)

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
## [EllipsisWith](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L184>)

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

<a name="Enum"></a>
## [Enum](<https://github.com/go-softwarelab/common/blob/main/pkg/to/enum.go#L11>)

```go
func Enum[V ~string, T ~string](value V, enumValues ...T) (T, error)
```

Enum converts the provided value to the available enum value in a case\-insensitive manner. In case when the provided value doesn't match any of the enums, it will return an empty string as the enum and error. This is a case\-insensitive version if you prefer to be more strict \(value matches exactly one of the enums\), use to.EnumStrict.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	type FooBarEnum string
	const (
		Foo FooBarEnum = "foo"
		Bar FooBarEnum = "bar"
	)

	// Case-insensitive match succeeds
	v, err := to.Enum("FoO", Foo, Bar)
	fmt.Printf("%q, Error: %v\n", v, err)

	// No match found -> returns empty string and an error
	v, err = to.Enum("baz", Foo, Bar)
	fmt.Printf("%q, Error: %v\n", v, err)

}
```

**Output**

```
"foo", Error: <nil>
"", Error: invalid value: baz doesn't match enum values [foo bar]
```


</details>

<a name="EnumStrict"></a>
## [EnumStrict](<https://github.com/go-softwarelab/common/blob/main/pkg/to/enum.go#L23>)

```go
func EnumStrict[V ~string, T ~string](value V, enumValues ...T) (T, error)
```

EnumStrict converts the provided value to the available enum value in a strict manner \(value matches exactly one of the enums\). In case when the provided value doesn't match any of the enums, it will return an empty string as the enum and error. This is a strict version if you prefer to be more flexible about value \(value matches one of the enums in case\-insensitive manner\), use to.Enum.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	type FooBarEnum string
	const (
		Foo FooBarEnum = "foo"
		Bar FooBarEnum = "bar"
	)

	// Case-sensitive match succeeds
	v, err := to.EnumStrict("foo", Foo, Bar)
	fmt.Printf("%q, Error: %v\n", v, err)

	// Case-sensitive mismatch -> returns empty string and an error
	v, err = to.EnumStrict("FOO", Foo, Bar)
	fmt.Printf("%q, Error: %v\n", v, err)

}
```

**Output**

```
"foo", Error: <nil>
"", Error: invalid value: FOO doesn't match enum values [foo bar]
```


</details>

<a name="Float32"></a>
## [Float32](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1088>)

```go
func Float32[V ConvertableToNumber](value V) (float32, error)
```

Float32 will convert bool, any number or string, and their subtypes to float32

NOTE: This is using reflection, if it is an issue, use dedicated functions for a given type, with suffix like

```
FromBool or FromSigned.
```

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Example for bool value
	boolVal, boolErr := to.Float32(true)
	fmt.Printf("Bool: %T(%g), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.Float32("3.14159")
	fmt.Printf("String: %T(%g), Error: %v\n", strVal, strVal, strErr)

	// Example for float32 value
	float32Val, float32Err := to.Float32(float32(3.14159))
	fmt.Printf("Float32: %T(%g), Error: %v\n", float32Val, float32Val, float32Err)

	// Example for int value
	intVal, intErr := to.Float32(3)
	fmt.Printf("Int: %T(%g), Error: %v\n", intVal, intVal, intErr)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.Float32(CustomString("3.14159"))
	fmt.Printf("CustomString: %T(%g), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int
	type CustomInt int
	customIntVal, customIntErr := to.Float32(CustomInt(3))
	fmt.Printf("CustomInt: %T(%g), Error: %v\n", customIntVal, customIntVal, customIntErr)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.Float32(CustomBool(true))
	fmt.Printf("CustomBool: %T(%g), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

}
```

**Output**

```
Bool: float32(1), Error: <nil>
String: float32(3.14159), Error: <nil>
Float32: float32(3.14159), Error: <nil>
Int: float32(3), Error: <nil>
CustomString: float32(3.14159), Error: <nil>
CustomInt: float32(3), Error: <nil>
CustomBool: float32(1), Error: <nil>
```


</details>

<a name="Float32FromBool"></a>
## [Float32FromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1156>)

```go
func Float32FromBool(value bool) float32
```

Float32FromBool converts a boolean value to its float32 representation \(true = 1.0, false = 0.0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.Float32FromBool(true)
	fmt.Printf("true: %T(%g)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Float32FromBool(false)
	fmt.Printf("false: %T(%g)\n", falseVal, falseVal)

}
```

**Output**

```
true: float32(1)
false: float32(0)
```


</details>

<a name="Float32FromSigned"></a>
## [Float32FromSigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1164>)

```go
func Float32FromSigned[V types.SignedNumber](value V) (float32, error)
```

Float32FromSigned will convert any signed number to float32, with range checks

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
	val, err := to.Float32FromSigned(3)
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

}
```

**Output**

```
float32(3), Error: <nil>
```


</details>

<a name="Float32FromString"></a>
## [Float32FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1180>)

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
## [Float32FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1172>)

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
## [Float64](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1197>)

```go
func Float64[V ConvertableToNumber](value V) (float64, error)
```

Float64 will convert bool, any number or string, and their subtypes to float64

NOTE: This is using reflection, if it is an issue, use dedicated functions for a given type, with suffix like

```
FromBool or FromSigned.
```

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Example for bool value
	boolVal, boolErr := to.Float64(true)
	fmt.Printf("Bool: %T(%g), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.Float64("3.14159265359")
	fmt.Printf("String: %T(%g), Error: %v\n", strVal, strVal, strErr)

	// Example for float64 value
	float64Val, float64Err := to.Float64(float64(3.14159265359))
	fmt.Printf("Float64: %T(%g), Error: %v\n", float64Val, float64Val, float64Err)

	// Example for int value
	intVal, intErr := to.Float64(3)
	fmt.Printf("Int: %T(%g), Error: %v\n", intVal, intVal, intErr)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.Float64(CustomString("3.14159265359"))
	fmt.Printf("CustomString: %T(%g), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int
	type CustomInt int
	customIntVal, customIntErr := to.Float64(CustomInt(3))
	fmt.Printf("CustomInt: %T(%g), Error: %v\n", customIntVal, customIntVal, customIntErr)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.Float64(CustomBool(true))
	fmt.Printf("CustomBool: %T(%g), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

}
```

**Output**

```
Bool: float64(1), Error: <nil>
String: float64(3.14159265359), Error: <nil>
Float64: float64(3.14159265359), Error: <nil>
Int: float64(3), Error: <nil>
CustomString: float64(3.14159265359), Error: <nil>
CustomInt: float64(3), Error: <nil>
CustomBool: float64(1), Error: <nil>
```


</details>

<a name="Float64FromBool"></a>
## [Float64FromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1265>)

```go
func Float64FromBool(value bool) float64
```

Float64FromBool converts a boolean value to its float64 representation \(true = 1.0, false = 0.0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.Float64FromBool(true)
	fmt.Printf("true: %T(%g)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Float64FromBool(false)
	fmt.Printf("false: %T(%g)\n", falseVal, falseVal)

}
```

**Output**

```
true: float64(1)
false: float64(0)
```


</details>

<a name="Float64FromSigned"></a>
## [Float64FromSigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1273>)

```go
func Float64FromSigned[V types.SignedNumber](value V) (float64, error)
```

Float64FromSigned will convert any signed number to float64

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
	val, err := to.Float64FromSigned(3)
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

}
```

**Output**

```
float64(3), Error: <nil>
```


</details>

<a name="Float64FromString"></a>
## [Float64FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1283>)

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
## [Float64FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1278>)

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
## [Int](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L37>)

```go
func Int[V ConvertableToNumber](value V) (int, error)
```

Int will convert bool, any number or string, and their subtypes to int

NOTE: This is using reflection, if it is an issue, use dedicated functions for a given type, with suffix like

```
FromBool or FromSigned.
```

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Example for bool value
	boolVal, boolErr := to.Int(true)
	fmt.Printf("Bool: %T(%d), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.Int("42")
	fmt.Printf("String: %T(%d), Error: %v\n", strVal, strVal, strErr)

	// Example for uint64 value
	uint64Val, uint64Err := to.Int(uint64(9223372036854775807))
	fmt.Printf("Uint64: %T(%d), Error: %v\n", uint64Val, uint64Val, uint64Err)

	// Example for int64 value
	int64Val, int64Err := to.Int(int64(9223372036854775807))
	fmt.Printf("Int64: %T(%d), Error: %v\n", int64Val, int64Val, int64Err)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.Int(CustomString("42"))
	fmt.Printf("CustomString: %T(%d), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int64
	type CustomInt64 int64
	customInt64Val, customInt64Err := to.Int(CustomInt64(42))
	fmt.Printf("CustomInt64: %T(%d), Error: %v\n", customInt64Val, customInt64Val, customInt64Err)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.Int(CustomBool(true))
	fmt.Printf("CustomBool: %T(%d), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

}
```

**Output**

```
Bool: int(1), Error: <nil>
String: int(42), Error: <nil>
Uint64: int(9223372036854775807), Error: <nil>
Int64: int(9223372036854775807), Error: <nil>
CustomString: int(42), Error: <nil>
CustomInt64: int(42), Error: <nil>
CustomBool: int(1), Error: <nil>
```


</details>

<a name="Int16"></a>
## [Int16](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L257>)

```go
func Int16[V ConvertableToNumber](value V) (int16, error)
```

Int16 will convert bool, any number or string, and their subtypes to int16

NOTE: This is using reflection, if it is an issue, use dedicated functions for a given type, with suffix like

```
FromBool or FromSigned.
```

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
	// Example for bool value
	boolVal, boolErr := to.Int16(true)
	fmt.Printf("Bool: %T(%d), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.Int16("1000")
	fmt.Printf("String: %T(%d), Error: %v\n", strVal, strVal, strErr)

	// Example for uint16 value
	uint16Val, uint16Err := to.Int16(uint16(1000))
	fmt.Printf("Uint16: %T(%d), Error: %v\n", uint16Val, uint16Val, uint16Err)

	// Example for int value
	intVal, intErr := to.Int16(1000)
	fmt.Printf("Int: %T(%d), Error: %v\n", intVal, intVal, intErr)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.Int16(CustomString("1000"))
	fmt.Printf("CustomString: %T(%d), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int
	type CustomInt int
	customIntVal, customIntErr := to.Int16(CustomInt(1000))
	fmt.Printf("CustomInt: %T(%d), Error: %v\n", customIntVal, customIntVal, customIntErr)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.Int16(CustomBool(true))
	fmt.Printf("CustomBool: %T(%d), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

	// Converting out of range
	valOOR, errOOR := to.Int16(40000)
	fmt.Printf("Out of range: %T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

}
```

**Output**

```
Bool: int16(1), Error: <nil>
String: int16(1000), Error: <nil>
Uint16: int16(1000), Error: <nil>
Int: int16(1000), Error: <nil>
CustomString: int16(1000), Error: <nil>
CustomInt: int16(1000), Error: <nil>
CustomBool: int16(1), Error: <nil>
Out of range: int16(0), Error: true
```


</details>

<a name="Int16FromBool"></a>
## [Int16FromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L325>)

```go
func Int16FromBool(value bool) int16
```

Int16FromBool converts a boolean value to its int16 representation \(true = 1, false = 0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.Int16FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Int16FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

}
```

**Output**

```
true: int16(1)
false: int16(0)
```


</details>

<a name="Int16FromSigned"></a>
## [Int16FromSigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L333>)

```go
func Int16FromSigned[V types.SignedNumber](value V) (int16, error)
```

Int16FromSigned will convert any signed number to int16, with range checks

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
	val, err := to.Int16FromSigned(1000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int16FromSigned(40000)
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
## [Int16FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L350>)

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
## [Int16FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L342>)

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
## [Int32](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L367>)

```go
func Int32[V ConvertableToNumber](value V) (int32, error)
```

Int32 will convert bool, any number or string, and their subtypes to int32

NOTE: This is using reflection, if it is an issue, use dedicated functions for a given type, with suffix like

```
FromBool or FromSigned.
```

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Example for bool value
	boolVal, boolErr := to.Int32(true)
	fmt.Printf("Bool: %T(%d), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.Int32("1000000")
	fmt.Printf("String: %T(%d), Error: %v\n", strVal, strVal, strErr)

	// Example for uint32 value
	uint32Val, uint32Err := to.Int32(uint32(1000000))
	fmt.Printf("Uint32: %T(%d), Error: %v\n", uint32Val, uint32Val, uint32Err)

	// Example for int value
	intVal, intErr := to.Int32(1000000)
	fmt.Printf("Int: %T(%d), Error: %v\n", intVal, intVal, intErr)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.Int32(CustomString("1000000"))
	fmt.Printf("CustomString: %T(%d), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int
	type CustomInt int
	customIntVal, customIntErr := to.Int32(CustomInt(1000000))
	fmt.Printf("CustomInt: %T(%d), Error: %v\n", customIntVal, customIntVal, customIntErr)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.Int32(CustomBool(true))
	fmt.Printf("CustomBool: %T(%d), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

}
```

**Output**

```
Bool: int32(1), Error: <nil>
String: int32(1000000), Error: <nil>
Uint32: int32(1000000), Error: <nil>
Int: int32(1000000), Error: <nil>
CustomString: int32(1000000), Error: <nil>
CustomInt: int32(1000000), Error: <nil>
CustomBool: int32(1), Error: <nil>
```


</details>

<a name="Int32FromBool"></a>
## [Int32FromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L435>)

```go
func Int32FromBool(value bool) int32
```

Int32FromBool converts a boolean value to its int32 representation \(true = 1, false = 0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.Int32FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Int32FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

}
```

**Output**

```
true: int32(1)
false: int32(0)
```


</details>

<a name="Int32FromSigned"></a>
## [Int32FromSigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L443>)

```go
func Int32FromSigned[V types.SignedNumber](value V) (int32, error)
```

Int32FromSigned will convert any signed number to int32, with range checks

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
	val, err := to.Int32FromSigned(1000000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
int32(1000000), Error: <nil>
```


</details>

<a name="Int32FromString"></a>
## [Int32FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L460>)

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
## [Int32FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L452>)

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
## [Int64](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L477>)

```go
func Int64[V ConvertableToNumber](value V) (int64, error)
```

Int64 will convert bool, any number or string, and their subtypes to int64

NOTE: This is using reflection,

```
if it is an issue, use dedicated functions for a given type, with suffix like FromBool or FromSigned.
```

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Example for bool value
	boolVal, boolErr := to.Int64(true)
	fmt.Printf("Bool: %T(%d), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.Int64("9223372036854775807")
	fmt.Printf("String: %T(%d), Error: %v\n", strVal, strVal, strErr)

	// Example for uint64 value
	uint64Val, uint64Err := to.Int64(uint64(9223372036854775807))
	fmt.Printf("Uint64: %T(%d), Error: %v\n", uint64Val, uint64Val, uint64Err)

	// Example for int value
	intVal, intErr := to.Int64(9223372036854775807)
	fmt.Printf("Int: %T(%d), Error: %v\n", intVal, intVal, intErr)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.Int64(CustomString("9223372036854775807"))
	fmt.Printf("CustomString: %T(%d), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int
	type CustomInt int
	customIntVal, customIntErr := to.Int64(CustomInt(9223372036854775807))
	fmt.Printf("CustomInt: %T(%d), Error: %v\n", customIntVal, customIntVal, customIntErr)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.Int64(CustomBool(true))
	fmt.Printf("CustomBool: %T(%d), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

}
```

**Output**

```
Bool: int64(1), Error: <nil>
String: int64(9223372036854775807), Error: <nil>
Uint64: int64(9223372036854775807), Error: <nil>
Int: int64(9223372036854775807), Error: <nil>
CustomString: int64(9223372036854775807), Error: <nil>
CustomInt: int64(9223372036854775807), Error: <nil>
CustomBool: int64(1), Error: <nil>
```


</details>

<a name="Int64FromBool"></a>
## [Int64FromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L545>)

```go
func Int64FromBool(value bool) int64
```

Int64FromBool converts a boolean value to its int64 representation \(true = 1, false = 0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.Int64FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Int64FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

}
```

**Output**

```
true: int64(1)
false: int64(0)
```


</details>

<a name="Int64FromSigned"></a>
## [Int64FromSigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L553>)

```go
func Int64FromSigned[V types.SignedNumber](value V) (int64, error)
```

Int64FromSigned will convert any signed number to int64, with range checks

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
	val, err := to.Int64FromSigned(9223372036854775807)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
int64(9223372036854775807), Error: <nil>
```


</details>

<a name="Int64FromString"></a>
## [Int64FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L566>)

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
## [Int64FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L558>)

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
## [Int8](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L147>)

```go
func Int8[V ConvertableToNumber](value V) (int8, error)
```

Int8 will convert bool, any number or string, and their subtypes to int8

NOTE: This is using reflection,

```
if it is an issue, use dedicated functions for a given type, with suffix like FromBool or FromSigned.
```

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
	// Example for bool value
	boolVal, boolErr := to.Int8(true)
	fmt.Printf("Bool: %T(%d), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.Int8("42")
	fmt.Printf("String: %T(%d), Error: %v\n", strVal, strVal, strErr)

	// Example for uint8 value
	uint8Val, uint8Err := to.Int8(uint8(42))
	fmt.Printf("Uint8: %T(%d), Error: %v\n", uint8Val, uint8Val, uint8Err)

	// Example for int value
	intVal, intErr := to.Int8(42)
	fmt.Printf("Int: %T(%d), Error: %v\n", intVal, intVal, intErr)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.Int8(CustomString("42"))
	fmt.Printf("CustomString: %T(%d), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int
	type CustomInt int
	customIntVal, customIntErr := to.Int8(CustomInt(42))
	fmt.Printf("CustomInt: %T(%d), Error: %v\n", customIntVal, customIntVal, customIntErr)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.Int8(CustomBool(true))
	fmt.Printf("CustomBool: %T(%d), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

	// Converting out of range
	valOOR, errOOR := to.Int8(1000)
	fmt.Printf("Out of range: %T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

}
```

**Output**

```
Bool: int8(1), Error: <nil>
String: int8(42), Error: <nil>
Uint8: int8(42), Error: <nil>
Int: int8(42), Error: <nil>
CustomString: int8(42), Error: <nil>
CustomInt: int8(42), Error: <nil>
CustomBool: int8(1), Error: <nil>
Out of range: int8(0), Error: true
```


</details>

<a name="Int8FromBool"></a>
## [Int8FromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L215>)

```go
func Int8FromBool(value bool) int8
```

Int8FromBool converts a boolean value to its int8 representation \(true = 1, false = 0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.Int8FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Int8FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

}
```

**Output**

```
true: int8(1)
false: int8(0)
```


</details>

<a name="Int8FromSigned"></a>
## [Int8FromSigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L223>)

```go
func Int8FromSigned[V types.SignedNumber](value V) (int8, error)
```

Int8FromSigned will convert any signed number to int8, with range checks

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
	val, err := to.Int8FromSigned(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int8FromSigned(1000)
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
## [Int8FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L240>)

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
## [Int8FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L232>)

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

<a name="IntFromBool"></a>
## [IntFromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L105>)

```go
func IntFromBool(value bool) int
```

IntFromBool converts a boolean value to its integer representation \(true = 1, false = 0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.IntFromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.IntFromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

}
```

**Output**

```
true: int(1)
false: int(0)
```


</details>

<a name="IntFromSigned"></a>
## [IntFromSigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L113>)

```go
func IntFromSigned[V types.SignedNumber](value V) (int, error)
```

IntFromSigned will convert any signed number to int, with range checks

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
	val, err := to.IntFromSigned(int16(42))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
int(42), Error: <nil>
```


</details>

<a name="IntFromString"></a>
## [IntFromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L130>)

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
## [IntFromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L122>)

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
## [KebabCase](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L132>)

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
## [NoLessThan](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1296>)

```go
func NoLessThan[T types.Ordered](value, min T) T
```

NoLessThan will return the value if it's not less than the min value or the min value.

<a name="NoMoreThan"></a>
## [NoMoreThan](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1330>)

```go
func NoMoreThan[T types.Ordered](value, max T) T
```

NoMoreThan will return the value if it's not more than the max value or the max value.

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
## [PascalCase](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L110>)

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
## [Sentences](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L91>)

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
## [SliceOfAny](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L82>)

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
## [SliceOfPtr](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L59>)

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
## [SliceOfValue](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L68>)

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
## [SnakeCase](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L141>)

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

<a name="String"></a>
## [String](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L28>)

```go
func String[T any](value T) string
```

String converts any value to string. It takes into account some interfaces like: encoding.TextMarshaler

```
if it can produce text without returning an error it will be the preferred result.
if it returns the error, it will fallback to other methods to convert to string.
The support for encoding.TextMarshaler is experimental and may change in the future
```

fmt.Stringer \- is preferred over default string conversion

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

type exampleStringer string

func (s exampleStringer) String() string {
	return "stringer-" + string(s)
}

type exampleTextMarshaler string

func (m exampleTextMarshaler) MarshalText() ([]byte, error) {
	if string(m) != "hello" {
		return nil, fmt.Errorf("this example allows only hello value")
	}
	return []byte("text-marshaler-" + string(m)), nil
}

func (m exampleTextMarshaler) String() string {
	return "fallback-on-fail-marshal-text"
}

func main() {
	var str string

	// String type
	str = to.String("hello")
	fmt.Printf("%q\n", str)

	// Custom string type
	type CustomString string
	str = to.String(CustomString("hello-custom"))
	fmt.Printf("%q\n", str)

	// fmt.Stringer
	str = to.String(exampleStringer("hello"))
	fmt.Printf("%q\n", str)

	// TextMarshaler
	str = to.String(exampleTextMarshaler("hello"))
	fmt.Printf("%q\n", str)

	// TextMarshaler with error when marshaling will fallback to other methods of stringifying
	// This is experimental behavior and may change in the future
	str = to.String(exampleTextMarshaler("failing-marshal-text"))
	fmt.Printf("%q\n", str)

	// Integer type
	str = to.String(42)
	fmt.Printf("%q\n", str)

	// Unsigned integer type
	str = to.String(uint(123))
	fmt.Printf("%q\n", str)

	// Float type
	str = to.String(3.14159)
	fmt.Printf("%q\n", str)

	// Boolean type
	str = to.String(true)
	fmt.Printf("%q\n", str)

}
```

**Output**

```
"hello"
"hello-custom"
"stringer-hello"
"text-marshaler-hello"
"fallback-on-fail-marshal-text"
"42"
"123"
"3.14159"
"true"
```


</details>

<a name="StringFromBool"></a>
## [StringFromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L59>)

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
## [StringFromBytes](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L64>)

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
## [StringFromFloat](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L54>)

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
## [StringFromInteger](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L49>)

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
## [StringFromRune](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L69>)

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
## [UInt](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L583>)

```go
func UInt[V ConvertableToNumber](value V) (uint, error)
```

UInt will convert bool, any number or string, and their subtypes to uint

NOTE: This is using reflection, if it is an issue, use dedicated functions for a given type, with suffix like

```
FromBool or FromNumber.
```

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
	// Example for bool value
	boolVal, boolErr := to.UInt(true)
	fmt.Printf("Bool: %T(%d), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.UInt("42")
	fmt.Printf("String: %T(%d), Error: %v\n", strVal, strVal, strErr)

	// Example for uint value
	uintVal, uintErr := to.UInt(uint(42))
	fmt.Printf("Uint: %T(%d), Error: %v\n", uintVal, uintVal, uintErr)

	// Example for int value
	intVal, intErr := to.UInt(42)
	fmt.Printf("Int: %T(%d), Error: %v\n", intVal, intVal, intErr)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.UInt(CustomString("42"))
	fmt.Printf("CustomString: %T(%d), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int
	type CustomInt int
	customIntVal, customIntErr := to.UInt(CustomInt(42))
	fmt.Printf("CustomInt: %T(%d), Error: %v\n", customIntVal, customIntVal, customIntErr)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.UInt(CustomBool(true))
	fmt.Printf("CustomBool: %T(%d), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

	// Converting negative value
	valNeg, errNeg := to.UInt(-5)
	fmt.Printf("Negative: %T(%d), Error: %v\n", valNeg, valNeg, errors.Is(errNeg, to.ErrValueOutOfRange))

}
```

**Output**

```
Bool: uint(1), Error: <nil>
String: uint(42), Error: <nil>
Uint: uint(42), Error: <nil>
Int: uint(42), Error: <nil>
CustomString: uint(42), Error: <nil>
CustomInt: uint(42), Error: <nil>
CustomBool: uint(1), Error: <nil>
Negative: uint(0), Error: true
```


</details>

<a name="UInt16"></a>
## [UInt16](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L785>)

```go
func UInt16[V ConvertableToNumber](value V) (uint16, error)
```

UInt16 will convert bool, any number or string, and their subtypes to uint16

NOTE: This is using reflection, if it is an issue, use dedicated functions for a given type, with suffix like

```
FromBool or FromNumber.
```

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Example for bool value
	boolVal, boolErr := to.UInt16(true)
	fmt.Printf("Bool: %T(%d), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.UInt16("65000")
	fmt.Printf("String: %T(%d), Error: %v\n", strVal, strVal, strErr)

	// Example for uint16 value
	uint16Val, uint16Err := to.UInt16(uint16(65000))
	fmt.Printf("Uint16: %T(%d), Error: %v\n", uint16Val, uint16Val, uint16Err)

	// Example for int value
	intVal, intErr := to.UInt16(65000)
	fmt.Printf("Int: %T(%d), Error: %v\n", intVal, intVal, intErr)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.UInt16(CustomString("65000"))
	fmt.Printf("CustomString: %T(%d), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int
	type CustomInt int
	customIntVal, customIntErr := to.UInt16(CustomInt(65000))
	fmt.Printf("CustomInt: %T(%d), Error: %v\n", customIntVal, customIntVal, customIntErr)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.UInt16(CustomBool(true))
	fmt.Printf("CustomBool: %T(%d), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

}
```

**Output**

```
Bool: uint16(1), Error: <nil>
String: uint16(65000), Error: <nil>
Uint16: uint16(65000), Error: <nil>
Int: uint16(65000), Error: <nil>
CustomString: uint16(65000), Error: <nil>
CustomInt: uint16(65000), Error: <nil>
CustomBool: uint16(1), Error: <nil>
```


</details>

<a name="UInt16FromBool"></a>
## [UInt16FromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L853>)

```go
func UInt16FromBool(value bool) uint16
```

UInt16FromBool converts a boolean value to its uint16 representation \(true = 1, false = 0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.UInt16FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.UInt16FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

}
```

**Output**

```
true: uint16(1)
false: uint16(0)
```


</details>

<a name="UInt16FromNumber"></a>
## [UInt16FromNumber](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L861>)

```go
func UInt16FromNumber[V types.Number](value V) (uint16, error)
```

UInt16FromNumber will convert any number to uint16, with range checks

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
	val, err := to.UInt16FromNumber(65000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
uint16(65000), Error: <nil>
```


</details>

<a name="UInt16FromString"></a>
## [UInt16FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L869>)

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
## [UInt32](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L886>)

```go
func UInt32[V ConvertableToNumber](value V) (uint32, error)
```

UInt32 will convert bool, any number or string, and their subtypes to uint32

NOTE: This is using reflection, if it is an issue, use dedicated functions for a given type, with suffix like

```
FromBool or FromNumber.
```

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
	// Example for bool value
	boolVal, boolErr := to.UInt32(true)
	fmt.Printf("Bool: %T(%d), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.UInt32("4294967295")
	fmt.Printf("String: %T(%d), Error: %v\n", strVal, strVal, strErr)

	// Example for uint32 value
	uint32Val, uint32Err := to.UInt32(uint32(4294967295))
	fmt.Printf("Uint32: %T(%d), Error: %v\n", uint32Val, uint32Val, uint32Err)

	// Example for int value
	intVal, intErr := to.UInt32(42)
	fmt.Printf("Int: %T(%d), Error: %v\n", intVal, intVal, intErr)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.UInt32(CustomString("42"))
	fmt.Printf("CustomString: %T(%d), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int
	type CustomInt int
	customIntVal, customIntErr := to.UInt32(CustomInt(42))
	fmt.Printf("CustomInt: %T(%d), Error: %v\n", customIntVal, customIntVal, customIntErr)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.UInt32(CustomBool(true))
	fmt.Printf("CustomBool: %T(%d), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

	// Negative number
	valNeg, errNeg := to.UInt32(-5)
	fmt.Printf("Negative: %T(%d), Error: %v\n", valNeg, valNeg, errors.Is(errNeg, to.ErrValueOutOfRange))

}
```

**Output**

```
Bool: uint32(1), Error: <nil>
String: uint32(4294967295), Error: <nil>
Uint32: uint32(4294967295), Error: <nil>
Int: uint32(42), Error: <nil>
CustomString: uint32(42), Error: <nil>
CustomInt: uint32(42), Error: <nil>
CustomBool: uint32(1), Error: <nil>
Negative: uint32(0), Error: true
```


</details>

<a name="UInt32FromBool"></a>
## [UInt32FromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L954>)

```go
func UInt32FromBool(value bool) uint32
```

UInt32FromBool converts a boolean value to its uint32 representation \(true = 1, false = 0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.UInt32FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.UInt32FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

}
```

**Output**

```
true: uint32(1)
false: uint32(0)
```


</details>

<a name="UInt32FromNumber"></a>
## [UInt32FromNumber](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L962>)

```go
func UInt32FromNumber[V types.Number](value V) (uint32, error)
```

UInt32FromNumber will convert any number to uint32, with range checks

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
	val, err := to.UInt32FromNumber(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Negative number
	valNeg, errNeg := to.UInt32FromNumber(-5)
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
## [UInt32FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L970>)

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
## [UInt64](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L987>)

```go
func UInt64[V ConvertableToNumber](value V) (uint64, error)
```

UInt64 will convert bool, any number or string, and their subtypes to uint64

NOTE: This is using reflection, if it is an issue, use dedicated functions for a given type, with suffix like

```
FromBool or FromNumber.
```

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Example for bool value
	boolVal, boolErr := to.UInt64(true)
	fmt.Printf("Bool: %T(%d), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.UInt64("18446744073709551615")
	fmt.Printf("String: %T(%d), Error: %v\n", strVal, strVal, strErr)

	// Example for uint64 value
	uint64Val, uint64Err := to.UInt64(uint64(18446744073709551615))
	fmt.Printf("Uint64: %T(%d), Error: %v\n", uint64Val, uint64Val, uint64Err)

	// Example for int value
	intVal, intErr := to.UInt64(42)
	fmt.Printf("Int: %T(%d), Error: %v\n", intVal, intVal, intErr)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.UInt64(CustomString("42"))
	fmt.Printf("CustomString: %T(%d), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int
	type CustomInt int
	customIntVal, customIntErr := to.UInt64(CustomInt(42))
	fmt.Printf("CustomInt: %T(%d), Error: %v\n", customIntVal, customIntVal, customIntErr)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.UInt64(CustomBool(true))
	fmt.Printf("CustomBool: %T(%d), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

}
```

**Output**

```
Bool: uint64(1), Error: <nil>
String: uint64(18446744073709551615), Error: <nil>
Uint64: uint64(18446744073709551615), Error: <nil>
Int: uint64(42), Error: <nil>
CustomString: uint64(42), Error: <nil>
CustomInt: uint64(42), Error: <nil>
CustomBool: uint64(1), Error: <nil>
```


</details>

<a name="UInt64FromBool"></a>
## [UInt64FromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1055>)

```go
func UInt64FromBool(value bool) uint64
```

UInt64FromBool converts a boolean value to its uint64 representation \(true = 1, false = 0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.UInt64FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.UInt64FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

}
```

**Output**

```
true: uint64(1)
false: uint64(0)
```


</details>

<a name="UInt64FromNumber"></a>
## [UInt64FromNumber](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1063>)

```go
func UInt64FromNumber[V types.Number](value V) (uint64, error)
```

UInt64FromNumber will convert any number to uint64, with range checks

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
	val, err := to.UInt64FromNumber(uint(18446744073709551000))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

}
```

**Output**

```
uint64(18446744073709551000), Error: <nil>
```


</details>

<a name="UInt64FromString"></a>
## [UInt64FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1071>)

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
## [UInt8](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L684>)

```go
func UInt8[V ConvertableToNumber](value V) (uint8, error)
```

UInt8 will convert bool, any number or string, and their subtypes to uint8

NOTE: This is using reflection, if it is an issue, use dedicated functions for a given type, with suffix like

```
FromBool or FromNumber.
```

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
	// Example for bool value
	boolVal, boolErr := to.UInt8(true)
	fmt.Printf("Bool: %T(%d), Error: %v\n", boolVal, boolVal, boolErr)

	// Example for string value
	strVal, strErr := to.UInt8("200")
	fmt.Printf("String: %T(%d), Error: %v\n", strVal, strVal, strErr)

	// Example for uint8 value
	uint8Val, uint8Err := to.UInt8(uint8(200))
	fmt.Printf("Uint8: %T(%d), Error: %v\n", uint8Val, uint8Val, uint8Err)

	// Example for int value
	intVal, intErr := to.UInt8(200)
	fmt.Printf("Int: %T(%d), Error: %v\n", intVal, intVal, intErr)

	// Example for custom type based on string
	type CustomString string
	customStrVal, customStrErr := to.UInt8(CustomString("200"))
	fmt.Printf("CustomString: %T(%d), Error: %v\n", customStrVal, customStrVal, customStrErr)

	// Example for custom type based on int
	type CustomInt int
	customIntVal, customIntErr := to.UInt8(CustomInt(200))
	fmt.Printf("CustomInt: %T(%d), Error: %v\n", customIntVal, customIntVal, customIntErr)

	// Example for custom type based on bool
	type CustomBool bool
	customBoolVal, customBoolErr := to.UInt8(CustomBool(true))
	fmt.Printf("CustomBool: %T(%d), Error: %v\n", customBoolVal, customBoolVal, customBoolErr)

	// Converting out of range
	valOOR, errOOR := to.UInt8(300)
	fmt.Printf("Out of range: %T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

}
```

**Output**

```
Bool: uint8(1), Error: <nil>
String: uint8(200), Error: <nil>
Uint8: uint8(200), Error: <nil>
Int: uint8(200), Error: <nil>
CustomString: uint8(200), Error: <nil>
CustomInt: uint8(200), Error: <nil>
CustomBool: uint8(1), Error: <nil>
Out of range: uint8(0), Error: true
```


</details>

<a name="UInt8FromBool"></a>
## [UInt8FromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L752>)

```go
func UInt8FromBool(value bool) uint8
```

UInt8FromBool converts a boolean value to its uint8 representation \(true = 1, false = 0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.UInt8FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.UInt8FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

}
```

**Output**

```
true: uint8(1)
false: uint8(0)
```


</details>

<a name="UInt8FromNumber"></a>
## [UInt8FromNumber](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L760>)

```go
func UInt8FromNumber[V types.Number](value V) (uint8, error)
```

UInt8FromNumber will convert any number to uint8, with range checks

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
	val, err := to.UInt8FromNumber(200)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.UInt8FromNumber(300)
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
## [UInt8FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L768>)

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

<a name="UIntFromBool"></a>
## [UIntFromBool](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L651>)

```go
func UIntFromBool(value bool) uint
```

UIntFromBool converts a boolean value to its uint representation \(true = 1, false = 0\).

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	// Converting true value
	trueVal := to.UIntFromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.UIntFromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

}
```

**Output**

```
true: uint(1)
false: uint(0)
```


</details>

<a name="UIntFromNumber"></a>
## [UIntFromNumber](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L659>)

```go
func UIntFromNumber[V types.Number](value V) (uint, error)
```

UIntFromNumber will convert any number to uint, with range checks

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
	val, err := to.UIntFromNumber(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting negative value
	valNeg, errNeg := to.UIntFromNumber(-5)
	fmt.Printf("%T(%d), Error: %v\n", valNeg, valNeg, errors.Is(errNeg, to.ErrValueOutOfRange))

}
```

**Output**

```
uint(42), Error: <nil>
uint(0), Error: true
```


</details>

<a name="UIntFromString"></a>
## [UIntFromString](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L667>)

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

<a name="ValueAtLeast"></a>
## [ValueAtLeast](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1301>)

```go
func ValueAtLeast[T types.Ordered](value, min T) T
```

ValueAtLeast will return the value if it's not less than the min value or the min value.

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
	val := to.ValueAtLeast(5, 0)
	fmt.Printf("ValueAtLeast(5, 0) = %d\n", val)

	// Below min
	val = to.ValueAtLeast(-5, 0)
	fmt.Printf("ValueAtLeast(-5, 0) = %d\n", val)

	// Works with float values too
	floatVal := to.ValueAtLeast(3.14, 4.0)
	fmt.Printf("ValueAtLeast(3.14, 4.0) = %.2f\n", floatVal)

}
```

**Output**

```
ValueAtLeast(5, 0) = 5
ValueAtLeast(-5, 0) = 0
ValueAtLeast(3.14, 4.0) = 4.00
```


</details>

<a name="ValueAtMost"></a>
## [ValueAtMost](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1335>)

```go
func ValueAtMost[T types.Ordered](value, max T) T
```

ValueAtMost will return the value if it's not more than the max value or the max value.

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
	val := to.ValueAtMost(5, 10)
	fmt.Printf("ValueAtMost(5, 10) = %d\n", val)

	// Above max
	val = to.ValueAtMost(15, 10)
	fmt.Printf("ValueAtMost(15, 10) = %d\n", val)

}
```

**Output**

```
ValueAtMost(5, 10) = 5
ValueAtMost(15, 10) = 10
```


</details>

<a name="ValueBetween"></a>
## [ValueBetween](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1369>)

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
## [ValueBetweenThe](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L1385>)

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

<a name="ValueOrGet"></a>
## [ValueOrGet](<https://github.com/go-softwarelab/common/blob/main/pkg/to/types.go#L50>)

```go
func ValueOrGet[T any](x *T, fallback func() T) T
```

ValueOrGet returns the pointer value or the fallback value from result of fallback function call.

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
	val := to.ValueOrGet(ptr, func() int { return 0 })
	fmt.Printf("%T(%d)\n", val, val)

	// Get fallback value from a nil pointer using fallback function
	var nilPtr *string
	strVal := to.ValueOrGet(nilPtr, func() string { return "computed fallback" })
	fmt.Printf("%T(%q)\n", strVal, strVal)

}
```

**Output**

```
int(42)
string("computed fallback")
```


</details>

<a name="Words"></a>
## [Words](<https://github.com/go-softwarelab/common/blob/main/pkg/to/strings.go#L74>)

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

<a name="YesNo"></a>
## [YesNo](<https://github.com/go-softwarelab/common/blob/main/pkg/to/booleans.go#L27>)

```go
func YesNo[T comparable](value T) string
```

YesNo returns "yes" if the input value is non\-zero and "no" if the input value is zero.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	type SomeStruct struct {
		Field string
	}

	// Converting various values to "yes"/"no" string
	fmt.Println("empty string ->", to.YesNo(""))
	fmt.Println("non-empty string ->", to.YesNo("hello"))
	fmt.Println("zero ->", to.YesNo(0))
	fmt.Println("positive number ->", to.YesNo(42))
	fmt.Println("negative number ->", to.YesNo(-42))
	fmt.Println("nil struct ->", to.YesNo((*SomeStruct)(nil)))
	fmt.Println("empty struct ->", to.YesNo(SomeStruct{}))
	fmt.Println("filled struct ->", to.YesNo(SomeStruct{Field: "hello"}))

}
```

**Output**

```
empty string -> no
non-empty string -> yes
zero -> no
positive number -> yes
negative number -> yes
nil struct -> no
empty struct -> no
filled struct -> yes
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

<a name="ConvertableToNumber"></a>
## type [ConvertableToNumber](<https://github.com/go-softwarelab/common/blob/main/pkg/to/numbers.go#L14-L16>)

ConvertableToNumber is a constraint that permits types that can be converted to number, such as strings, numbers, or booleans.

```go
type ConvertableToNumber interface {
    // contains filtered or unexported methods
}
```

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
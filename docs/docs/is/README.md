# is

```go
import "github.com/go-softwarelab/common/pkg/is"
```

Package is provides a comprehensive set of predicates and validation functions for common conditional checks in Go applications.

The goal of this package is to offer concise, reusable predicate functions that can be used for common validation scenarios, reducing boilerplate code and improving readability. The functions can be combined with other predicates using operators like Not\(\) to create more complex conditions.

These utilities are designed to be used in functional programming patterns and conditional logic while maintaining type safety through generics.



## Variables

<a name="Bool"></a>Bool checks if a value is a bool.

```go
var Bool = Type[bool]
```

<a name="Int"></a>Int checks if a value is an int.

```go
var Int = Type[int]
```

<a name="String"></a>String checks if a value is a string.

```go
var String = Type[string]
```

<a name="Between"></a>
## [Between](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L80>)

```go
func Between[T types.Ordered](value, a, b T) bool
```

Between checks if a value is between two others.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Between checks if a value falls within a range (inclusive)
	fmt.Printf("%T(%v)\n", is.Between(25, 18, 65), is.Between(25, 18, 65))
	fmt.Printf("%T(%v)\n", is.Between(18, 18, 65), is.Between(18, 18, 65))
	fmt.Printf("%T(%v)\n", is.Between(65, 18, 65), is.Between(65, 18, 65))
	fmt.Printf("%T(%v)\n", is.Between(17, 18, 65), is.Between(17, 18, 65))
	fmt.Printf("%T(%v)\n", is.Between(66, 18, 65), is.Between(66, 18, 65))

}
```

**Output**

```
bool(true)
bool(true)
bool(true)
bool(false)
bool(false)
```


</details>

<a name="BetweenThe"></a>
## [BetweenThe](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L85>)

```go
func BetweenThe[T types.Ordered](a, b T) func(T) bool
```

BetweenThe checks returns the function that checks if a value is between two others.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// BetweenThe creates a function checking if values are within a range
	isWorkingAge := is.BetweenThe(18, 65)

	fmt.Printf("%T(%v)\n", isWorkingAge(25), isWorkingAge(25))
	fmt.Printf("%T(%v)\n", isWorkingAge(18), isWorkingAge(18))
	fmt.Printf("%T(%v)\n", isWorkingAge(65), isWorkingAge(65))
	fmt.Printf("%T(%v)\n", isWorkingAge(16), isWorkingAge(16))
	fmt.Printf("%T(%v)\n", isWorkingAge(70), isWorkingAge(70))

	// Temperature range checker (comfortable room temperature 20-25°C)
	isComfortableTemp := is.BetweenThe(20.0, 25.0)
	fmt.Printf("%T(%v)\n", isComfortableTemp(22.5), isComfortableTemp(22.5))
	fmt.Printf("%T(%v)\n", isComfortableTemp(18.0), isComfortableTemp(18.0))

}
```

**Output**

```
bool(true)
bool(true)
bool(true)
bool(false)
bool(false)
bool(true)
bool(false)
```


</details>

<a name="BlankString"></a>
## [BlankString](<https://github.com/go-softwarelab/common/blob/main/pkg/is/strings.go#L17>)

```go
func BlankString[S string | *string](str S) (isBlank bool)
```

BlankString returns true if the string is empty or contains only whitespace

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// BlankString checks if a string is empty or contains only whitespace
	fmt.Println("Checking for blank strings:")
	fmt.Printf("is.BlankString(\"\"): %T(%v) - empty string\n", is.BlankString(""), is.BlankString(""))
	fmt.Printf("is.BlankString(\"   \"): %T(%v) - whitespace only\n", is.BlankString("   "), is.BlankString("   "))
	fmt.Printf("is.BlankString(\"\\t\\n\"): %T(%v) - tabs and newlines\n", is.BlankString("\t\n"), is.BlankString("\t\n"))
	fmt.Printf("is.BlankString(\"hello\"): %T(%v) - non-blank string\n", is.BlankString("hello"), is.BlankString("hello"))

}
```

**Output**

```
Checking for blank strings:
is.BlankString(""): bool(true) - empty string
is.BlankString("   "): bool(true) - whitespace only
is.BlankString("\t\n"): bool(true) - tabs and newlines
is.BlankString("hello"): bool(false) - non-blank string
```


</details>

<a name="Empty"></a>
## [Empty](<https://github.com/go-softwarelab/common/blob/main/pkg/is/types.go#L17>)

```go
func Empty[T comparable](value T) bool
```

Empty checks if a value is zero value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Testing zero/empty values of different types
	fmt.Println("Checking for empty/zero values:")

	fmt.Printf("is.Empty(0): %T(%v) - zero int\n", is.Empty(0), is.Empty(0))
	fmt.Printf("is.Empty(42): %T(%v) - non-zero int\n", is.Empty(42), is.Empty(42))

	fmt.Printf("is.Empty(\"\"): %T(%v) - empty string\n", is.Empty(""), is.Empty(""))
	fmt.Printf("is.Empty(\"hello\"): %T(%v) - non-empty string\n", is.Empty("hello"), is.Empty("hello"))

	fmt.Printf("is.Empty(false): %T(%v) - zero bool\n", is.Empty(false), is.Empty(false))
	fmt.Printf("is.Empty(true): %T(%v) - non-zero bool\n", is.Empty(true), is.Empty(true))

}
```

**Output**

```
Checking for empty/zero values:
is.Empty(0): bool(true) - zero int
is.Empty(42): bool(false) - non-zero int
is.Empty(""): bool(true) - empty string
is.Empty("hello"): bool(false) - non-empty string
is.Empty(false): bool(true) - zero bool
is.Empty(true): bool(false) - non-zero bool
```


</details>

<a name="EmptyString"></a>
## [EmptyString](<https://github.com/go-softwarelab/common/blob/main/pkg/is/strings.go#L6>)

```go
func EmptyString[S string | *string](s S) (isEmpty bool)
```

EmptyString returns true if the string is empty

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// EmptyString checks if a string is empty
	fmt.Println("Checking for empty strings:")
	fmt.Printf("is.EmptyString(\"\"): %T(%v) - empty string\n", is.EmptyString(""), is.EmptyString(""))
	fmt.Printf("is.EmptyString(\"hello\"): %T(%v) - non-empty string\n", is.EmptyString("hello"), is.EmptyString("hello"))

}
```

**Output**

```
Checking for empty strings:
is.EmptyString(""): bool(true) - empty string
is.EmptyString("hello"): bool(false) - non-empty string
```


</details>

<a name="Equal"></a>
## [Equal](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L15>)

```go
func Equal[T comparable](a, b T) bool
```

Equal checks if two values are equal.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Equal checks equality between two values
	fmt.Println("Checking equality:")
	fmt.Printf("is.Equal(42, 42): %v (same integers)\n", is.Equal(42, 42))
	fmt.Printf("is.Equal(42, 43): %v (different integers)\n", is.Equal(42, 43))
	fmt.Printf("is.Equal(\"hello\", \"hello\"): %v (same strings)\n", is.Equal("hello", "hello"))
	fmt.Printf("is.Equal(\"hello\", \"world\"): %v (different strings)\n", is.Equal("hello", "world"))

}
```

**Output**

```
Checking equality:
is.Equal(42, 42): true (same integers)
is.Equal(42, 43): false (different integers)
is.Equal("hello", "hello"): true (same strings)
is.Equal("hello", "world"): false (different strings)
```


</details>

<a name="EqualTo"></a>
## [EqualTo](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L25>)

```go
func EqualTo[T comparable](expected T) func(T) bool
```

EqualTo returns the function that checks if a value is equal to another.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// EqualTo creates a function that checks for equality with a specific value
	fmt.Println("Using equality checker functions:")

	// Create a function that checks if a value equals 42
	isFortyTwo := is.EqualTo(42)
	fmt.Printf("isFortyTwo(42): %v (matching value)\n", isFortyTwo(42))
	fmt.Printf("isFortyTwo(43): %v (non-matching value)\n", isFortyTwo(43))

	// Create a function that checks if a string equals "admin"
	isAdmin := is.EqualTo("admin")
	fmt.Printf("isAdmin(\"admin\"): %v (matching role)\n", isAdmin("admin"))
	fmt.Printf("isAdmin(\"user\"): %v (non-matching role)\n", isAdmin("user"))

}
```

**Output**

```
Using equality checker functions:
isFortyTwo(42): true (matching value)
isFortyTwo(43): false (non-matching value)
isAdmin("admin"): true (matching role)
isAdmin("user"): false (non-matching role)
```


</details>

<a name="Greater"></a>
## [Greater](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L32>)

```go
func Greater[T types.Ordered](a, b T) bool
```

Greater checks if a value is greater than another.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Greater checks if first value is greater than second
	fmt.Println("Checking greater than relation:")
	fmt.Printf("is.Greater(42, 30): %v (42 > 30)\n", is.Greater(42, 30))
	fmt.Printf("is.Greater(30, 42): %v (30 > 42)\n", is.Greater(30, 42))
	fmt.Printf("is.Greater(42, 42): %v (42 > 42)\n", is.Greater(42, 42))

	// String comparison is lexicographic
	fmt.Printf("is.Greater(\"zebra\", \"apple\"): %v (alphabetical order)\n", is.Greater("zebra", "apple"))
	fmt.Printf("is.Greater(\"apple\", \"zebra\"): %v (alphabetical order)\n", is.Greater("apple", "zebra"))

}
```

**Output**

```
Checking greater than relation:
is.Greater(42, 30): true (42 > 30)
is.Greater(30, 42): false (30 > 42)
is.Greater(42, 42): false (42 > 42)
is.Greater("zebra", "apple"): true (alphabetical order)
is.Greater("apple", "zebra"): false (alphabetical order)
```


</details>

<a name="GreaterOrEqual"></a>
## [GreaterOrEqual](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L44>)

```go
func GreaterOrEqual[T types.Ordered](a, b T) bool
```

GreaterOrEqual checks if a value is greater than or equal to another.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// GreaterOrEqual checks if first value is greater than or equal to second
	fmt.Println("Checking greater than or equal relation:")
	fmt.Printf("is.GreaterOrEqual(42, 30): %v (42 ≥ 30)\n", is.GreaterOrEqual(42, 30))
	fmt.Printf("is.GreaterOrEqual(42, 42): %v (42 ≥ 42)\n", is.GreaterOrEqual(42, 42))
	fmt.Printf("is.GreaterOrEqual(30, 42): %v (30 ≥ 42)\n", is.GreaterOrEqual(30, 42))

}
```

**Output**

```
Checking greater than or equal relation:
is.GreaterOrEqual(42, 30): true (42 ≥ 30)
is.GreaterOrEqual(42, 42): true (42 ≥ 42)
is.GreaterOrEqual(30, 42): false (30 ≥ 42)
```


</details>

<a name="GreaterOrEqualTo"></a>
## [GreaterOrEqualTo](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L49>)

```go
func GreaterOrEqualTo[T types.Ordered](expected T) func(T) bool
```

GreaterOrEqualTo returns the function that checks if a value is greater than or equal to another.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// GreaterOrEqualTo creates a function checking if values meet or exceed a threshold
	fmt.Println("Using minimum threshold functions:")

	// Check if age qualifies for voting (18 or older)
	canVote := is.GreaterOrEqualTo(18)
	fmt.Printf("canVote(21): %v (21 meets voting age)\n", canVote(21))
	fmt.Printf("canVote(18): %v (18 meets voting age)\n", canVote(18))
	fmt.Printf("canVote(16): %v (16 below voting age)\n", canVote(16))

	// Check if score is passing grade (60 or higher)
	isPassing := is.GreaterOrEqualTo(60.0)
	fmt.Printf("isPassing(75.5): %v (75.5 is passing)\n", isPassing(75.5))
	fmt.Printf("isPassing(60.0): %v (60.0 is passing)\n", isPassing(60.0))
	fmt.Printf("isPassing(59.9): %v (59.9 is failing)\n", isPassing(59.9))

}
```

**Output**

```
Using minimum threshold functions:
canVote(21): true (21 meets voting age)
canVote(18): true (18 meets voting age)
canVote(16): false (16 below voting age)
isPassing(75.5): true (75.5 is passing)
isPassing(60.0): true (60.0 is passing)
isPassing(59.9): false (59.9 is failing)
```


</details>

<a name="GreaterThan"></a>
## [GreaterThan](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L37>)

```go
func GreaterThan[T types.Ordered](expected T) func(T) bool
```

GreaterThan returns the function that checks if a value is greater than another.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// GreaterThan creates a function that checks if values exceed a threshold
	fmt.Println("Using threshold checker functions:")

	// Check if age is considered adult (over 18)
	isAdult := is.GreaterThan(18)
	fmt.Printf("isAdult(21): %v (21 exceeds adult threshold)\n", isAdult(21))
	fmt.Printf("isAdult(18): %v (18 equals adult threshold)\n", isAdult(18))
	fmt.Printf("isAdult(16): %v (16 below adult threshold)\n", isAdult(16))

	// Check if temperature is hot (over 30°C)
	isHot := is.GreaterThan(30.0)
	fmt.Printf("isHot(35.5): %v (35.5°C exceeds hot threshold)\n", isHot(35.5))
	fmt.Printf("isHot(25.0): %v (25.0°C below hot threshold)\n", isHot(25.0))

}
```

**Output**

```
Using threshold checker functions:
isAdult(21): true (21 exceeds adult threshold)
isAdult(18): false (18 equals adult threshold)
isAdult(16): false (16 below adult threshold)
isHot(35.5): true (35.5°C exceeds hot threshold)
isHot(25.0): false (25.0°C below hot threshold)
```


</details>

<a name="Less"></a>
## [Less](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L56>)

```go
func Less[T types.Ordered](a, b T) bool
```

Less checks if a value is less than another.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Less checks if first value is less than second
	fmt.Println("Checking less than relation:")
	fmt.Printf("is.Less(30, 42): %v (30 < 42)\n", is.Less(30, 42))
	fmt.Printf("is.Less(42, 30): %v (42 < 30)\n", is.Less(42, 30))
	fmt.Printf("is.Less(42, 42): %v (42 < 42)\n", is.Less(42, 42))

	// String comparison is lexicographic
	fmt.Printf("is.Less(\"apple\", \"zebra\"): %v (alphabetical order)\n", is.Less("apple", "zebra"))

}
```

**Output**

```
Checking less than relation:
is.Less(30, 42): true (30 < 42)
is.Less(42, 30): false (42 < 30)
is.Less(42, 42): false (42 < 42)
is.Less("apple", "zebra"): true (alphabetical order)
```


</details>

<a name="LessOrEqual"></a>
## [LessOrEqual](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L68>)

```go
func LessOrEqual[T types.Ordered](a, b T) bool
```

LessOrEqual checks if a value is less than or equal to another.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// LessOrEqual checks if first value is less than or equal to second
	fmt.Println("Checking less than or equal relation:")
	fmt.Printf("is.LessOrEqual(30, 42): %v (30 ≤ 42)\n", is.LessOrEqual(30, 42))
	fmt.Printf("is.LessOrEqual(42, 42): %v (42 ≤ 42)\n", is.LessOrEqual(42, 42))
	fmt.Printf("is.LessOrEqual(50, 42): %v (50 ≤ 42)\n", is.LessOrEqual(50, 42))

}
```

**Output**

```
Checking less than or equal relation:
is.LessOrEqual(30, 42): true (30 ≤ 42)
is.LessOrEqual(42, 42): true (42 ≤ 42)
is.LessOrEqual(50, 42): false (50 ≤ 42)
```


</details>

<a name="LessOrEqualTo"></a>
## [LessOrEqualTo](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L73>)

```go
func LessOrEqualTo[T types.Ordered](expected T) func(T) bool
```

LessOrEqualTo returns the function that checks if a value is less than or equal to another.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// LessOrEqualTo creates a function checking if values are at or below a threshold
	fmt.Println("Using maximum allowed functions:")

	// Check if child qualifies for child pricing (12 or under)
	isChildRate := is.LessOrEqualTo(12)
	fmt.Printf("isChildRate(10): %v (10 qualifies for child rate)\n", isChildRate(10))
	fmt.Printf("isChildRate(12): %v (12 qualifies for child rate)\n", isChildRate(12))
	fmt.Printf("isChildRate(13): %v (13 does not qualify for child rate)\n", isChildRate(13))

	// Check if luggage is within weight limit (23kg or less)
	isWithinWeightLimit := is.LessOrEqualTo(23.0)
	fmt.Printf("isWithinWeightLimit(20.5): %v (20.5kg is acceptable)\n", isWithinWeightLimit(20.5))
	fmt.Printf("isWithinWeightLimit(23.0): %v (23.0kg is acceptable)\n", isWithinWeightLimit(23.0))
	fmt.Printf("isWithinWeightLimit(23.5): %v (23.5kg exceeds limit)\n", isWithinWeightLimit(23.5))

}
```

**Output**

```
Using maximum allowed functions:
isChildRate(10): true (10 qualifies for child rate)
isChildRate(12): true (12 qualifies for child rate)
isChildRate(13): false (13 does not qualify for child rate)
isWithinWeightLimit(20.5): true (20.5kg is acceptable)
isWithinWeightLimit(23.0): true (23.0kg is acceptable)
isWithinWeightLimit(23.5): false (23.5kg exceeds limit)
```


</details>

<a name="LessThan"></a>
## [LessThan](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L61>)

```go
func LessThan[T types.Ordered](expected T) func(T) bool
```

LessThan returns the function that checks if a value is less than another.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// LessThan creates a function that checks if values are below a threshold
	fmt.Println("Using maximum threshold functions:")

	// Check if someone is a minor (under 18)
	isMinor := is.LessThan(18)
	fmt.Printf("isMinor(16): %v (16 is below adult threshold)\n", isMinor(16))
	fmt.Printf("isMinor(18): %v (18 equals adult threshold)\n", isMinor(18))
	fmt.Printf("isMinor(21): %v (21 exceeds adult threshold)\n", isMinor(21))

	// Check if temperature is freezing (below 0°C)
	isFreezing := is.LessThan(0.0)
	fmt.Printf("isFreezing(-5.0): %v (-5.0°C is freezing)\n", isFreezing(-5.0))
	fmt.Printf("isFreezing(0.0): %v (0.0°C at freezing point)\n", isFreezing(0.0))

}
```

**Output**

```
Using maximum threshold functions:
isMinor(16): true (16 is below adult threshold)
isMinor(18): false (18 equals adult threshold)
isMinor(21): false (21 exceeds adult threshold)
isFreezing(-5.0): true (-5.0°C is freezing)
isFreezing(0.0): false (0.0°C at freezing point)
```


</details>

<a name="Nil"></a>
## [Nil](<https://github.com/go-softwarelab/common/blob/main/pkg/is/types.go#L6>)

```go
func Nil(value any) bool
```

Nil checks if a value is nil.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Testing nil values
	fmt.Println("Checking for nil values:")

	var nilSlice []int
	fmt.Printf("is.Nil(nil): %T(%v) - literal nil\n", is.Nil(nil), is.Nil(nil))
	fmt.Printf("is.Nil(nilSlice): %T(%v) - nil slice\n", is.Nil(nilSlice), is.Nil(nilSlice))

	nonNilSlice := make([]int, 0)
	fmt.Printf("is.Nil(nonNilSlice): %T(%v) - empty but initialized slice\n", is.Nil(nonNilSlice), is.Nil(nonNilSlice))

	var nilMap map[string]int
	fmt.Printf("is.Nil(nilMap): %T(%v) - nil map\n", is.Nil(nilMap), is.Nil(nilMap))

}
```

**Output**

```
Checking for nil values:
is.Nil(nil): bool(true) - literal nil
is.Nil(nilSlice): bool(true) - nil slice
is.Nil(nonNilSlice): bool(false) - empty but initialized slice
is.Nil(nilMap): bool(true) - nil map
```


</details>

<a name="Not"></a>
## [Not](<https://github.com/go-softwarelab/common/blob/main/pkg/is/operators.go#L4>)

```go
func Not[T any](predicate func(T) bool) func(T) bool
```

Not returns a function that inverts the result of the given predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Create a function to check if a value is positive
	isPositive := func(x int) bool {
		return x > 0
	}

	// Create the negation of that function
	isNotPositive := is.Not(isPositive)

	// Test with various values
	fmt.Printf("isNotPositive(5): %T(%v) - positive number\n", isNotPositive(5), isNotPositive(5))
	fmt.Printf("isNotPositive(0): %T(%v) - zero\n", isNotPositive(0), isNotPositive(0))
	fmt.Printf("isNotPositive(-3): %T(%v) - negative number\n", isNotPositive(-3), isNotPositive(-3))

}
```

**Output**

```
isNotPositive(5): bool(false) - positive number
isNotPositive(0): bool(true) - zero
isNotPositive(-3): bool(true) - negative number
```


</details>

<a name="NotBlankString"></a>
## [NotBlankString](<https://github.com/go-softwarelab/common/blob/main/pkg/is/strings.go#L33>)

```go
func NotBlankString[S string | *string](s S) bool
```

NotBlankString returns true if the string is not empty or contains only whitespace

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// NotBlankString checks if a string has non-whitespace content
	fmt.Println("Checking for non-blank strings:")
	fmt.Printf("is.NotBlankString(\"hello\"): %T(%v) - string with content\n", is.NotBlankString("hello"), is.NotBlankString("hello"))
	fmt.Printf("is.NotBlankString(\" x \"): %T(%v) - string with whitespace and content\n", is.NotBlankString(" x "), is.NotBlankString(" x "))
	fmt.Printf("is.NotBlankString(\"   \"): %T(%v) - whitespace only\n", is.NotBlankString("   "), is.NotBlankString("   "))
	fmt.Printf("is.NotBlankString(\"\"): %T(%v) - empty string\n", is.NotBlankString(""), is.NotBlankString(""))

}
```

**Output**

```
Checking for non-blank strings:
is.NotBlankString("hello"): bool(true) - string with content
is.NotBlankString(" x "): bool(true) - string with whitespace and content
is.NotBlankString("   "): bool(false) - whitespace only
is.NotBlankString(""): bool(false) - empty string
```


</details>

<a name="NotEmpty"></a>
## [NotEmpty](<https://github.com/go-softwarelab/common/blob/main/pkg/is/types.go#L23>)

```go
func NotEmpty[T comparable](value T) bool
```

NotEmpty checks if a value is not zero value.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Testing non-empty values of different types
	fmt.Println("Checking for non-empty values:")

	fmt.Printf("is.NotEmpty(42): %T(%v) - non-zero int\n", is.NotEmpty(42), is.NotEmpty(42))
	fmt.Printf("is.NotEmpty(0): %T(%v) - zero int\n", is.NotEmpty(0), is.NotEmpty(0))

	fmt.Printf("is.NotEmpty(\"hello\"): %T(%v) - non-empty string\n", is.NotEmpty("hello"), is.NotEmpty("hello"))
	fmt.Printf("is.NotEmpty(\"\"): %T(%v) - empty string\n", is.NotEmpty(""), is.NotEmpty(""))

	fmt.Printf("is.NotEmpty(true): %T(%v) - non-zero bool\n", is.NotEmpty(true), is.NotEmpty(true))
	fmt.Printf("is.NotEmpty(false): %T(%v) - zero bool\n", is.NotEmpty(false), is.NotEmpty(false))

}
```

**Output**

```
Checking for non-empty values:
is.NotEmpty(42): bool(true) - non-zero int
is.NotEmpty(0): bool(false) - zero int
is.NotEmpty("hello"): bool(true) - non-empty string
is.NotEmpty(""): bool(false) - empty string
is.NotEmpty(true): bool(true) - non-zero bool
is.NotEmpty(false): bool(false) - zero bool
```


</details>

<a name="NotEmptyString"></a>
## [NotEmptyString](<https://github.com/go-softwarelab/common/blob/main/pkg/is/strings.go#L28>)

```go
func NotEmptyString[S string | *string](s S) bool
```

NotEmptyString returns true if the string is not empty

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// NotEmptyString checks if a string is not empty
	fmt.Println("Checking for non-empty strings:")
	fmt.Printf("is.NotEmptyString(\"hello\"): %T(%v) - non-empty string\n", is.NotEmptyString("hello"), is.NotEmptyString("hello"))
	fmt.Printf("is.NotEmptyString(\"   \"): %T(%v) - whitespace string\n", is.NotEmptyString("   "), is.NotEmptyString("   "))
	fmt.Printf("is.NotEmptyString(\"\"): %T(%v) - empty string\n", is.NotEmptyString(""), is.NotEmptyString(""))

}
```

**Output**

```
Checking for non-empty strings:
is.NotEmptyString("hello"): bool(true) - non-empty string
is.NotEmptyString("   "): bool(true) - whitespace string
is.NotEmptyString(""): bool(false) - empty string
```


</details>

<a name="NotEqual"></a>
## [NotEqual](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L20>)

```go
func NotEqual[T comparable](a, b T) bool
```

NotEqual checks if two values are not equal.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// NotEqual checks inequality between two values
	fmt.Println("Checking inequality:")
	fmt.Printf("is.NotEqual(42, 43): %v (different integers)\n", is.NotEqual(42, 43))
	fmt.Printf("is.NotEqual(42, 42): %v (same integers)\n", is.NotEqual(42, 42))
	fmt.Printf("is.NotEqual(\"hello\", \"world\"): %v (different strings)\n", is.NotEqual("hello", "world"))
	fmt.Printf("is.NotEqual(\"hello\", \"hello\"): %v (same strings)\n", is.NotEqual("hello", "hello"))

}
```

**Output**

```
Checking inequality:
is.NotEqual(42, 43): true (different integers)
is.NotEqual(42, 42): false (same integers)
is.NotEqual("hello", "world"): true (different strings)
is.NotEqual("hello", "hello"): false (same strings)
```


</details>

<a name="NotNil"></a>
## [NotNil](<https://github.com/go-softwarelab/common/blob/main/pkg/is/types.go#L12>)

```go
func NotNil(value any) bool
```

NotNil checks if a value is not nil.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Testing non-nil values
	fmt.Println("Checking for non-nil values:")

	var nilSlice []int
	nonNilSlice := make([]int, 0)

	fmt.Printf("is.NotNil(nonNilSlice): %T(%v) - empty but initialized slice\n", is.NotNil(nonNilSlice), is.NotNil(nonNilSlice))
	fmt.Printf("is.NotNil(nilSlice): %T(%v) - nil slice\n", is.NotNil(nilSlice), is.NotNil(nilSlice))
	fmt.Printf("is.NotNil(\"hello\"): %T(%v) - string value\n", is.NotNil("hello"), is.NotNil("hello"))
	fmt.Printf("is.NotNil(nil): %T(%v) - literal nil\n", is.NotNil(nil), is.NotNil(nil))

}
```

**Output**

```
Checking for non-nil values:
is.NotNil(nonNilSlice): bool(true) - empty but initialized slice
is.NotNil(nilSlice): bool(false) - nil slice
is.NotNil("hello"): bool(true) - string value
is.NotNil(nil): bool(false) - literal nil
```


</details>

<a name="NotOrError"></a>
## [NotOrError](<https://github.com/go-softwarelab/common/blob/main/pkg/is/operators.go#L11>)

```go
func NotOrError[T any](predicate func(T) (bool, error)) func(T) (bool, error)
```

NotOrError returns a function that inverts the result of the given predicate.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Create a function to check if a value is positive with error handling
	isPositiveWithError := func(x int) (bool, error) {
		return x > 0, nil
	}

	// Create the negation of that function
	isNotPositiveWithError := is.NotOrError(isPositiveWithError)

	// Test with various values
	result1, err1 := isNotPositiveWithError(5)
	fmt.Printf("isNotPositiveWithError(5): %T(%v), err: %v - positive number\n", result1, result1, err1)

	result2, err2 := isNotPositiveWithError(0)
	fmt.Printf("isNotPositiveWithError(0): %T(%v), err: %v - zero\n", result2, result2, err2)

	result3, err3 := isNotPositiveWithError(-3)
	fmt.Printf("isNotPositiveWithError(-3): %T(%v), err: %v - negative number\n", result3, result3, err3)

}
```

**Output**

```
isNotPositiveWithError(5): bool(false), err: <nil> - positive number
isNotPositiveWithError(0): bool(true), err: <nil> - zero
isNotPositiveWithError(-3): bool(true), err: <nil> - negative number
```


</details>

<a name="Type"></a>
## [Type](<https://github.com/go-softwarelab/common/blob/main/pkg/is/types.go#L28>)

```go
func Type[T any](value any) bool
```

Type checks if a value is of a specific type.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Testing type checking
	fmt.Println("Type checking:")

	value1 := "hello"
	value2 := 42
	value3 := true

	fmt.Printf("is.Type[string](value1): %T(%v) - string value\n", is.Type[string](value1), is.Type[string](value1))
	fmt.Printf("is.Type[int](value1): %T(%v) - string value checked as int\n", is.Type[int](value1), is.Type[int](value1))

	fmt.Printf("is.Type[int](value2): %T(%v) - int value\n", is.Type[int](value2), is.Type[int](value2))
	fmt.Printf("is.Type[string](value2): %T(%v) - int value checked as string\n", is.Type[string](value2), is.Type[string](value2))

	fmt.Printf("is.Type[bool](value3): %T(%v) - bool value\n", is.Type[bool](value3), is.Type[bool](value3))

}
```

**Output**

```
Type checking:
is.Type[string](value1): bool(true) - string value
is.Type[int](value1): bool(false) - string value checked as int
is.Type[int](value2): bool(true) - int value
is.Type[string](value2): bool(false) - int value checked as string
is.Type[bool](value3): bool(true) - bool value
```


</details>

<a name="Zero"></a>
## [Zero](<https://github.com/go-softwarelab/common/blob/main/pkg/is/comparisons.go#L10>)

```go
func Zero[T comparable](value T) bool
```

Zero checks if a value is zero.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func main() {
	// Zero detects zero values of different types
	fmt.Println("Checking zero values:")
	fmt.Printf("is.Zero(0): %v (integer zero value)\n", is.Zero(0))
	fmt.Printf("is.Zero(42): %v (non-zero integer)\n", is.Zero(42))
	fmt.Printf("is.Zero(\"\"): %v (empty string)\n", is.Zero(""))
	fmt.Printf("is.Zero(\"hello\"): %v (non-empty string)\n", is.Zero("hello"))

}
```

**Output**

```
Checking zero values:
is.Zero(0): true (integer zero value)
is.Zero(42): false (non-zero integer)
is.Zero(""): true (empty string)
is.Zero("hello"): false (non-empty string)
```


</details>
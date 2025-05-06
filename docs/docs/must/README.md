# must

```go
import "github.com/go-softwarelab/common/pkg/must"
```

Package must provides a collection of utility functions that panic in error scenarios instead of returning errors, simplifying code in specific contexts.

The goal of this package is to offer helper functions for situations where errors cannot be meaningfully handled at runtime, such as when errors would indicate programmer mistakes rather than external conditions. It's particularly useful in cases where errors are not expected because values have been pre\-validated or when handling initialization code that should fail fast.

These utilities are designed to reduce error\-checking boilerplate and improve code readability in initialization paths, configuration loading, and other contexts where failures represent exceptional conditions that should halt execution.



<a name="ConvertToFloat32"></a>
## [ConvertToFloat32](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L84>)

```go
func ConvertToFloat32[V types.SignedNumber](value V) float32
```

ConvertToFloat32 converts any signed number to float32, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToFloat32(3)
	fmt.Printf("%T(%g)\n", val, val)

	// Demonstrating panic with recovery for extreme values
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic because the value exceeds float32 range
		tooLarge := float64(3.5e+38) // Exceeds max float32 (approx 3.4e+38)
		_ = must.ConvertToFloat32(tooLarge)
	}()

}
```

**Output**

```
float32(3)
Error: 3.5e+38 value out of range to convert to float32
```


</details>

<a name="ConvertToFloat32FromString"></a>
## [ConvertToFloat32FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L154>)

```go
func ConvertToFloat32FromString(value string) float32
```

ConvertToFloat32FromString converts a string to float32, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToFloat32FromString("3.14")
	fmt.Printf("%T(%g)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to invalid syntax
		_ = must.ConvertToFloat32FromString("not-a-float")
	}()

}
```

**Output**

```
float32(3.14)
Error: invalid syntax of not-a-float to parse into number
```


</details>

<a name="ConvertToFloat32FromUnsigned"></a>
## [ConvertToFloat32FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L94>)

```go
func ConvertToFloat32FromUnsigned[V types.Unsigned](value V) float32
```

ConvertToFloat32FromUnsigned converts any unsigned number to float32, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToFloat32FromUnsigned(uint(42))
	fmt.Printf("%T(%g)\n", val, val)

}
```

**Output**

```
float32(42)
```


</details>

<a name="ConvertToFloat64"></a>
## [ConvertToFloat64](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L89>)

```go
func ConvertToFloat64[V types.SignedNumber](value V) float64
```

ConvertToFloat64 converts any signed number to float64

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToFloat64(3)
	fmt.Printf("%T(%g)\n", val, val)

}
```

**Output**

```
float64(3)
```


</details>

<a name="ConvertToFloat64FromString"></a>
## [ConvertToFloat64FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L159>)

```go
func ConvertToFloat64FromString(value string) float64
```

ConvertToFloat64FromString converts a string to float64, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToFloat64FromString("3.141592653589793")
	fmt.Printf("%T(%g)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to invalid syntax
		_ = must.ConvertToFloat64FromString("not-a-float")
	}()

}
```

**Output**

```
float64(3.141592653589793)
Error: invalid syntax of not-a-float to parse into number
```


</details>

<a name="ConvertToFloat64FromUnsigned"></a>
## [ConvertToFloat64FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L99>)

```go
func ConvertToFloat64FromUnsigned[V types.Unsigned](value V) float64
```

ConvertToFloat64FromUnsigned converts any unsigned number to float64

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToFloat64FromUnsigned(uint64(42))
	fmt.Printf("%T(%g)\n", val, val)

}
```

**Output**

```
float64(42)
```


</details>

<a name="ConvertToInt"></a>
## [ConvertToInt](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L9>)

```go
func ConvertToInt[V types.SignedNumber](value V) int
```

ConvertToInt converts any signed number to int, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt(int16(42))
	fmt.Printf("%T(%d)\n", val, val)

}
```

**Output**

```
int(42)
```


</details>

<a name="ConvertToInt16"></a>
## [ConvertToInt16](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L29>)

```go
func ConvertToInt16[V types.SignedNumber](value V) int16
```

ConvertToInt16 converts any signed number to int16, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt16(1000)
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToInt16(40000)
	}()

}
```

**Output**

```
int16(1000)
Error: 40000 value out of range to convert to int16
```


</details>

<a name="ConvertToInt16FromString"></a>
## [ConvertToInt16FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L114>)

```go
func ConvertToInt16FromString(value string) int16
```

ConvertToInt16FromString converts a string to int16, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt16FromString("1000")
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToInt16FromString("40000")
	}()

}
```

**Output**

```
int16(1000)
Error: 40000 value out of range to convert to int16
```


</details>

<a name="ConvertToInt16FromUnsigned"></a>
## [ConvertToInt16FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L34>)

```go
func ConvertToInt16FromUnsigned[V types.Unsigned](value V) int16
```

ConvertToInt16FromUnsigned converts any unsigned number to int16, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt16FromUnsigned(uint(1000))
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToInt16FromUnsigned(uint(40000))
	}()

}
```

**Output**

```
int16(1000)
Error: 40000 value out of range to convert to int16
```


</details>

<a name="ConvertToInt32"></a>
## [ConvertToInt32](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L39>)

```go
func ConvertToInt32[V types.SignedNumber](value V) int32
```

ConvertToInt32 converts any signed number to int32, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt32(1000000)
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToInt32(int64(3000000000))
	}()

}
```

**Output**

```
int32(1000000)
Error: 3000000000 value out of range to convert to int32
```


</details>

<a name="ConvertToInt32FromString"></a>
## [ConvertToInt32FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L119>)

```go
func ConvertToInt32FromString(value string) int32
```

ConvertToInt32FromString converts a string to int32, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt32FromString("1000000")
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToInt32FromString("3000000000")
	}()

}
```

**Output**

```
int32(1000000)
Error: 3000000000 value out of range to convert to int32
```


</details>

<a name="ConvertToInt32FromUnsigned"></a>
## [ConvertToInt32FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L44>)

```go
func ConvertToInt32FromUnsigned[V types.Unsigned](value V) int32
```

ConvertToInt32FromUnsigned converts any unsigned number to int32, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt32FromUnsigned(uint(1000000))
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToInt32FromUnsigned(uint64(3000000000))
	}()

}
```

**Output**

```
int32(1000000)
Error: 3000000000 value out of range to convert to int32
```


</details>

<a name="ConvertToInt64"></a>
## [ConvertToInt64](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L49>)

```go
func ConvertToInt64[V types.SignedNumber](value V) int64
```

ConvertToInt64 converts any signed number to int64, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt64(9223372036854775807)
	fmt.Printf("%T(%d)\n", val, val)

}
```

**Output**

```
int64(9223372036854775807)
```


</details>

<a name="ConvertToInt64FromString"></a>
## [ConvertToInt64FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L124>)

```go
func ConvertToInt64FromString(value string) int64
```

ConvertToInt64FromString converts a string to int64, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt64FromString("9223372036854775807")
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToInt64FromString("9223372036854775808")
	}()

}
```

**Output**

```
int64(9223372036854775807)
Error: 9223372036854775808 value out of range to convert to int64
```


</details>

<a name="ConvertToInt64FromUnsigned"></a>
## [ConvertToInt64FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L54>)

```go
func ConvertToInt64FromUnsigned[V types.Unsigned](value V) int64
```

ConvertToInt64FromUnsigned converts any unsigned number to int64, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt64FromUnsigned(uint64(9223372036854775807))
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToInt64FromUnsigned(uint64(9223372036854775808))
	}()

}
```

**Output**

```
int64(9223372036854775807)
Error: 9223372036854775808 value out of range to convert to int64
```


</details>

<a name="ConvertToInt8"></a>
## [ConvertToInt8](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L19>)

```go
func ConvertToInt8[V types.SignedNumber](value V) int8
```

ConvertToInt8 converts any signed number to int8, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt8(42)
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToInt8(1000)
	}()

}
```

**Output**

```
int8(42)
Error: 1000 value out of range to convert to int8
```


</details>

<a name="ConvertToInt8FromString"></a>
## [ConvertToInt8FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L109>)

```go
func ConvertToInt8FromString(value string) int8
```

ConvertToInt8FromString converts a string to int8, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt8FromString("42")
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToInt8FromString("200")
	}()

}
```

**Output**

```
int8(42)
Error: 200 value out of range to convert to int8
```


</details>

<a name="ConvertToInt8FromUnsigned"></a>
## [ConvertToInt8FromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L24>)

```go
func ConvertToInt8FromUnsigned[V types.Unsigned](value V) int8
```

ConvertToInt8FromUnsigned converts any unsigned number to int8, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToInt8FromUnsigned(uint(42))
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToInt8FromUnsigned(uint(200))
	}()

}
```

**Output**

```
int8(42)
Error: 200 value out of range to convert to int8
```


</details>

<a name="ConvertToIntFromString"></a>
## [ConvertToIntFromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L104>)

```go
func ConvertToIntFromString(value string) int
```

ConvertToIntFromString converts a string to int, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting valid string
	val := must.ConvertToIntFromString("42")
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to invalid syntax
		_ = must.ConvertToIntFromString("not-a-number")
	}()

}
```

**Output**

```
int(42)
Error: invalid syntax of not-a-number to parse into number
```


</details>

<a name="ConvertToIntFromUnsigned"></a>
## [ConvertToIntFromUnsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L14>)

```go
func ConvertToIntFromUnsigned[V types.Unsigned](value V) int
```

ConvertToIntFromUnsigned converts any unsigned number to int, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToIntFromUnsigned(uint16(42))
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		huge := uint64(9223372036854775808)
		_ = must.ConvertToIntFromUnsigned(huge)
	}()

}
```

**Output**

```
int(42)
Error: 9223372036854775808 value out of range to convert to int
```


</details>

<a name="ConvertToUInt"></a>
## [ConvertToUInt](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L59>)

```go
func ConvertToUInt[V types.Number](value V) uint
```

ConvertToUInt converts any number to uint, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting positive value
	val := must.ConvertToUInt(42)
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to negative value
		_ = must.ConvertToUInt(-5)
	}()

}
```

**Output**

```
uint(42)
Error: -5 value out of range to convert to uint
```


</details>

<a name="ConvertToUInt16"></a>
## [ConvertToUInt16](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L69>)

```go
func ConvertToUInt16[V types.Number](value V) uint16
```

ConvertToUInt16 converts any number to uint16, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToUInt16(65000)
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToUInt16(70000)
	}()

}
```

**Output**

```
uint16(65000)
Error: 70000 value out of range to convert to uint16
```


</details>

<a name="ConvertToUInt16FromString"></a>
## [ConvertToUInt16FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L139>)

```go
func ConvertToUInt16FromString(value string) uint16
```

ConvertToUInt16FromString converts a string to uint16, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToUInt16FromString("65000")
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToUInt16FromString("70000")
	}()

}
```

**Output**

```
uint16(65000)
Error: 70000 value out of range to convert to uint16
```


</details>

<a name="ConvertToUInt32"></a>
## [ConvertToUInt32](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L74>)

```go
func ConvertToUInt32[V types.Number](value V) uint32
```

ConvertToUInt32 converts any number to uint32, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Valid conversion
	val := must.ConvertToUInt32(42)
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to negative value
		_ = must.ConvertToUInt32(-5)
	}()

}
```

**Output**

```
uint32(42)
Error: -5 value out of range to convert to uint32
```


</details>

<a name="ConvertToUInt32FromString"></a>
## [ConvertToUInt32FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L144>)

```go
func ConvertToUInt32FromString(value string) uint32
```

ConvertToUInt32FromString converts a string to uint32, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToUInt32FromString("4000000000")
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToUInt32FromString("5000000000")
	}()

}
```

**Output**

```
uint32(4000000000)
Error: 5000000000 value out of range to convert to uint32
```


</details>

<a name="ConvertToUInt64"></a>
## [ConvertToUInt64](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L79>)

```go
func ConvertToUInt64[V types.Number](value V) uint64
```

ConvertToUInt64 converts any number to uint64, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToUInt64(uint(18446744073709551000))
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to negative value
		_ = must.ConvertToUInt64(-1)
	}()

}
```

**Output**

```
uint64(18446744073709551000)
Error: -1 value out of range to convert to uint64
```


</details>

<a name="ConvertToUInt64FromString"></a>
## [ConvertToUInt64FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L149>)

```go
func ConvertToUInt64FromString(value string) uint64
```

ConvertToUInt64FromString converts a string to uint64, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToUInt64FromString("18446744073709551615")
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to invalid syntax
		_ = must.ConvertToUInt64FromString("invalid")
	}()

}
```

**Output**

```
uint64(18446744073709551615)
Error: invalid syntax of invalid to parse into number
```


</details>

<a name="ConvertToUInt8"></a>
## [ConvertToUInt8](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L64>)

```go
func ConvertToUInt8[V types.Number](value V) uint8
```

ConvertToUInt8 converts any number to uint8, panicking on range errors

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToUInt8(200)
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToUInt8(300)
	}()

}
```

**Output**

```
uint8(200)
Error: 300 value out of range to convert to uint8
```


</details>

<a name="ConvertToUInt8FromString"></a>
## [ConvertToUInt8FromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L134>)

```go
func ConvertToUInt8FromString(value string) uint8
```

ConvertToUInt8FromString converts a string to uint8, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting within range
	val := must.ConvertToUInt8FromString("200")
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to value out of range
		_ = must.ConvertToUInt8FromString("300")
	}()

}
```

**Output**

```
uint8(200)
Error: 300 value out of range to convert to uint8
```


</details>

<a name="ConvertToUIntFromString"></a>
## [ConvertToUIntFromString](<https://github.com/go-softwarelab/common/blob/main/pkg/must/converters.go#L129>)

```go
func ConvertToUIntFromString(value string) uint
```

ConvertToUIntFromString converts a string to uint, panicking in case if the string is not a valid number.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func main() {
	// Converting positive value
	val := must.ConvertToUIntFromString("42")
	fmt.Printf("%T(%d)\n", val, val)

	// Demonstrating panic with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()
		// This will panic due to negative value
		_ = must.ConvertToUIntFromString("-5")
	}()

}
```

**Output**

```
uint(42)
Error: invalid syntax of -5 to parse into number
```


</details>
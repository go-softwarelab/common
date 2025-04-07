# must

```go
import "github.com/go-softwarelab/common/pkg/must"
```

Package must provides a collection of utility functions that panic in error scenarios instead of returning errors, simplifying code in specific contexts.

The goal of this package is to offer helper functions for situations where errors cannot be meaningfully handled at runtime, such as when errors would indicate programmer mistakes rather than external conditions. It's particularly useful in cases where errors are not expected because values have been pre\-validated or when handling initialization code that should fail fast.

These utilities are designed to reduce error\-checking boilerplate and improve code readability in initialization paths, configuration loading, and other contexts where failures represent exceptional conditions that should halt execution.



<a name="ConvertToFloat32"></a>
## ConvertToFloat32

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

<a name="ConvertToFloat32FromUnsigned"></a>
## ConvertToFloat32FromUnsigned

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
## ConvertToFloat64

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

<a name="ConvertToFloat64FromUnsigned"></a>
## ConvertToFloat64FromUnsigned

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
## ConvertToInt

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
## ConvertToInt16

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

<a name="ConvertToInt16FromUnsigned"></a>
## ConvertToInt16FromUnsigned

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
## ConvertToInt32

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

<a name="ConvertToInt32FromUnsigned"></a>
## ConvertToInt32FromUnsigned

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
## ConvertToInt64

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

<a name="ConvertToInt64FromUnsigned"></a>
## ConvertToInt64FromUnsigned

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
## ConvertToInt8

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

<a name="ConvertToInt8FromUnsigned"></a>
## ConvertToInt8FromUnsigned

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

<a name="ConvertToIntFromUnsigned"></a>
## ConvertToIntFromUnsigned

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
## ConvertToUInt

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
## ConvertToUInt16

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

<a name="ConvertToUInt32"></a>
## ConvertToUInt32

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

<a name="ConvertToUInt64"></a>
## ConvertToUInt64

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

<a name="ConvertToUInt8"></a>
## ConvertToUInt8

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
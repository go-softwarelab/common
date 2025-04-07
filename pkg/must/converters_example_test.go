package must_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/must"
)

func ExampleConvertToInt() {
	// Converting within range
	val := must.ConvertToInt(int16(42))
	fmt.Printf("%T(%d)\n", val, val)

	// Output:
	// int(42)
}

func ExampleConvertToIntFromUnsigned() {
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

	// Output:
	// int(42)
	// Error: 9223372036854775808 value out of range to convert to int
}

func ExampleConvertToInt8() {
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

	// Output:
	// int8(42)
	// Error: 1000 value out of range to convert to int8
}

func ExampleConvertToInt8FromUnsigned() {
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

	// Output:
	// int8(42)
	// Error: 200 value out of range to convert to int8
}

func ExampleConvertToInt16() {
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

	// Output:
	// int16(1000)
	// Error: 40000 value out of range to convert to int16
}

func ExampleConvertToInt16FromUnsigned() {
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

	// Output:
	// int16(1000)
	// Error: 40000 value out of range to convert to int16
}

func ExampleConvertToInt32() {
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

	// Output:
	// int32(1000000)
	// Error: 3000000000 value out of range to convert to int32
}

func ExampleConvertToInt32FromUnsigned() {
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

	// Output:
	// int32(1000000)
	// Error: 3000000000 value out of range to convert to int32
}

func ExampleConvertToInt64() {
	// Converting within range
	val := must.ConvertToInt64(9223372036854775807)
	fmt.Printf("%T(%d)\n", val, val)

	// Output:
	// int64(9223372036854775807)
}

func ExampleConvertToInt64FromUnsigned() {
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

	// Output:
	// int64(9223372036854775807)
	// Error: 9223372036854775808 value out of range to convert to int64
}

func ExampleConvertToUInt() {
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

	// Output:
	// uint(42)
	// Error: -5 value out of range to convert to uint
}

func ExampleConvertToUInt8() {
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

	// Output:
	// uint8(200)
	// Error: 300 value out of range to convert to uint8
}

func ExampleConvertToUInt16() {
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

	// Output:
	// uint16(65000)
	// Error: 70000 value out of range to convert to uint16
}

func ExampleConvertToUInt32() {
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

	// Output:
	// uint32(42)
	// Error: -5 value out of range to convert to uint32
}

func ExampleConvertToUInt64() {
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

	// Output:
	// uint64(18446744073709551000)
	// Error: -1 value out of range to convert to uint64
}

func ExampleConvertToFloat32() {
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

	// Output:
	// float32(3)
	// Error: 3.5e+38 value out of range to convert to float32
}

func ExampleConvertToFloat64() {
	// Converting within range
	val := must.ConvertToFloat64(3)
	fmt.Printf("%T(%g)\n", val, val)

	// Output:
	// float64(3)
}

func ExampleConvertToFloat32FromUnsigned() {
	// Converting within range
	val := must.ConvertToFloat32FromUnsigned(uint(42))
	fmt.Printf("%T(%g)\n", val, val)

	// Output:
	// float32(42)
}

func ExampleConvertToFloat64FromUnsigned() {
	// Converting within range
	val := must.ConvertToFloat64FromUnsigned(uint64(42))
	fmt.Printf("%T(%g)\n", val, val)

	// Output:
	// float64(42)
}

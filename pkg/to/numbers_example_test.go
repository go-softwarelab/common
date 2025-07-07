package to_test

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func ExampleInt() {
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

	// Output:
	// Bool: int(1), Error: <nil>
	// String: int(42), Error: <nil>
	// Uint64: int(9223372036854775807), Error: <nil>
	// Int64: int(9223372036854775807), Error: <nil>
	// CustomString: int(42), Error: <nil>
	// CustomInt64: int(42), Error: <nil>
	// CustomBool: int(1), Error: <nil>
}

func ExampleIntFromUnsigned() {
	// Converting within range
	val, err := to.IntFromUnsigned(uint16(42))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int(42), Error: <nil>
}

func ExampleIntFromBool() {
	// Converting true value
	trueVal := to.IntFromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.IntFromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

	// Output:
	// true: int(1)
	// false: int(0)
}

func ExampleInt8FromBool() {
	// Converting true value
	trueVal := to.Int8FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Int8FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

	// Output:
	// true: int8(1)
	// false: int8(0)
}

func ExampleInt16FromBool() {
	// Converting true value
	trueVal := to.Int16FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Int16FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

	// Output:
	// true: int16(1)
	// false: int16(0)
}

func ExampleInt32FromBool() {
	// Converting true value
	trueVal := to.Int32FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Int32FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

	// Output:
	// true: int32(1)
	// false: int32(0)
}

func ExampleInt64FromBool() {
	// Converting true value
	trueVal := to.Int64FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Int64FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

	// Output:
	// true: int64(1)
	// false: int64(0)
}

func ExampleUIntFromBool() {
	// Converting true value
	trueVal := to.UIntFromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.UIntFromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

	// Output:
	// true: uint(1)
	// false: uint(0)
}

func ExampleUInt8FromBool() {
	// Converting true value
	trueVal := to.UInt8FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.UInt8FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

	// Output:
	// true: uint8(1)
	// false: uint8(0)
}

func ExampleUInt16FromBool() {
	// Converting true value
	trueVal := to.UInt16FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.UInt16FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

	// Output:
	// true: uint16(1)
	// false: uint16(0)
}

func ExampleUInt32FromBool() {
	// Converting true value
	trueVal := to.UInt32FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.UInt32FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

	// Output:
	// true: uint32(1)
	// false: uint32(0)
}

func ExampleUInt64FromBool() {
	// Converting true value
	trueVal := to.UInt64FromBool(true)
	fmt.Printf("true: %T(%d)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.UInt64FromBool(false)
	fmt.Printf("false: %T(%d)\n", falseVal, falseVal)

	// Output:
	// true: uint64(1)
	// false: uint64(0)
}

func ExampleFloat32FromBool() {
	// Converting true value
	trueVal := to.Float32FromBool(true)
	fmt.Printf("true: %T(%g)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Float32FromBool(false)
	fmt.Printf("false: %T(%g)\n", falseVal, falseVal)

	// Output:
	// true: float32(1)
	// false: float32(0)
}

func ExampleFloat64FromBool() {
	// Converting true value
	trueVal := to.Float64FromBool(true)
	fmt.Printf("true: %T(%g)\n", trueVal, trueVal)

	// Converting false value
	falseVal := to.Float64FromBool(false)
	fmt.Printf("false: %T(%g)\n", falseVal, falseVal)

	// Output:
	// true: float64(1)
	// false: float64(0)
}

func ExampleInt8() {
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

	// Output:
	// Bool: int8(1), Error: <nil>
	// String: int8(42), Error: <nil>
	// Uint8: int8(42), Error: <nil>
	// Int: int8(42), Error: <nil>
	// CustomString: int8(42), Error: <nil>
	// CustomInt: int8(42), Error: <nil>
	// CustomBool: int8(1), Error: <nil>
	// Out of range: int8(0), Error: true
}

func ExampleInt8FromUnsigned() {
	// Converting within range
	val, err := to.Int8FromUnsigned(uint(42))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int8FromUnsigned(uint(200))
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

	// Output:
	// int8(42), Error: <nil>
	// int8(0), Error: true
}

func ExampleInt16() {
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

	// Output:
	// Bool: int16(1), Error: <nil>
	// String: int16(1000), Error: <nil>
	// Uint16: int16(1000), Error: <nil>
	// Int: int16(1000), Error: <nil>
	// CustomString: int16(1000), Error: <nil>
	// CustomInt: int16(1000), Error: <nil>
	// CustomBool: int16(1), Error: <nil>
	// Out of range: int16(0), Error: true
}

func ExampleInt16FromUnsigned() {
	// Converting within range
	val, err := to.Int16FromUnsigned(uint(1000))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int16FromUnsigned(uint(40000))
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

	// Output:
	// int16(1000), Error: <nil>
	// int16(0), Error: true
}

func ExampleInt32() {
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

	// Output:
	// Bool: int32(1), Error: <nil>
	// String: int32(1000000), Error: <nil>
	// Uint32: int32(1000000), Error: <nil>
	// Int: int32(1000000), Error: <nil>
	// CustomString: int32(1000000), Error: <nil>
	// CustomInt: int32(1000000), Error: <nil>
	// CustomBool: int32(1), Error: <nil>
}

func ExampleInt32FromUnsigned() {
	// Converting within range
	val, err := to.Int32FromUnsigned(uint(1000000))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int32(1000000), Error: <nil>
}

func ExampleInt64() {
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

	// Output:
	// Bool: int64(1), Error: <nil>
	// String: int64(9223372036854775807), Error: <nil>
	// Uint64: int64(9223372036854775807), Error: <nil>
	// Int: int64(9223372036854775807), Error: <nil>
	// CustomString: int64(9223372036854775807), Error: <nil>
	// CustomInt: int64(9223372036854775807), Error: <nil>
	// CustomBool: int64(1), Error: <nil>
}

func ExampleInt64FromUnsigned() {
	// Converting within range
	val, err := to.Int64FromUnsigned(uint64(9223372036854775807))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int64(9223372036854775807), Error: <nil>
}

func ExampleUInt() {
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

	// Output:
	// Bool: uint(1), Error: <nil>
	// String: uint(42), Error: <nil>
	// Uint: uint(42), Error: <nil>
	// Int: uint(42), Error: <nil>
	// CustomString: uint(42), Error: <nil>
	// CustomInt: uint(42), Error: <nil>
	// CustomBool: uint(1), Error: <nil>
	// Negative: uint(0), Error: true
}

func ExampleUIntFromNumber() {
	// Converting positive value
	val, err := to.UIntFromNumber(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting negative value
	valNeg, errNeg := to.UIntFromNumber(-5)
	fmt.Printf("%T(%d), Error: %v\n", valNeg, valNeg, errors.Is(errNeg, to.ErrValueOutOfRange))

	// Output:
	// uint(42), Error: <nil>
	// uint(0), Error: true
}

func ExampleUInt8() {
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

	// Output:
	// Bool: uint8(1), Error: <nil>
	// String: uint8(200), Error: <nil>
	// Uint8: uint8(200), Error: <nil>
	// Int: uint8(200), Error: <nil>
	// CustomString: uint8(200), Error: <nil>
	// CustomInt: uint8(200), Error: <nil>
	// CustomBool: uint8(1), Error: <nil>
	// Out of range: uint8(0), Error: true
}

func ExampleUInt8FromNumber() {
	// Converting within range
	val, err := to.UInt8FromNumber(200)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.UInt8FromNumber(300)
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

	// Output:
	// uint8(200), Error: <nil>
	// uint8(0), Error: true
}

func ExampleUInt16() {
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

	// Output:
	// Bool: uint16(1), Error: <nil>
	// String: uint16(65000), Error: <nil>
	// Uint16: uint16(65000), Error: <nil>
	// Int: uint16(65000), Error: <nil>
	// CustomString: uint16(65000), Error: <nil>
	// CustomInt: uint16(65000), Error: <nil>
	// CustomBool: uint16(1), Error: <nil>
}

func ExampleUInt16FromNumber() {
	// Converting within range
	val, err := to.UInt16FromNumber(65000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// uint16(65000), Error: <nil>
}

func ExampleUInt32() {
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

	// Output:
	// Bool: uint32(1), Error: <nil>
	// String: uint32(4294967295), Error: <nil>
	// Uint32: uint32(4294967295), Error: <nil>
	// Int: uint32(42), Error: <nil>
	// CustomString: uint32(42), Error: <nil>
	// CustomInt: uint32(42), Error: <nil>
	// CustomBool: uint32(1), Error: <nil>
	// Negative: uint32(0), Error: true
}

func ExampleUInt32FromNumber() {
	// Valid conversion
	val, err := to.UInt32FromNumber(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Negative number
	valNeg, errNeg := to.UInt32FromNumber(-5)
	fmt.Printf("%T(%d), Error: %v\n", valNeg, valNeg, errors.Is(errNeg, to.ErrValueOutOfRange))

	// Output:
	// uint32(42), Error: <nil>
	// uint32(0), Error: true
}

func ExampleUInt64() {
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

	// Output:
	// Bool: uint64(1), Error: <nil>
	// String: uint64(18446744073709551615), Error: <nil>
	// Uint64: uint64(18446744073709551615), Error: <nil>
	// Int: uint64(42), Error: <nil>
	// CustomString: uint64(42), Error: <nil>
	// CustomInt: uint64(42), Error: <nil>
	// CustomBool: uint64(1), Error: <nil>
}

func ExampleUInt64FromNumber() {
	// Converting within range
	val, err := to.UInt64FromNumber(uint(18446744073709551000))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// uint64(18446744073709551000), Error: <nil>
}

func ExampleFloat32() {
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

	// Output:
	// Bool: float32(1), Error: <nil>
	// String: float32(3.14159), Error: <nil>
	// Float32: float32(3.14159), Error: <nil>
	// Int: float32(3), Error: <nil>
	// CustomString: float32(3.14159), Error: <nil>
	// CustomInt: float32(3), Error: <nil>
	// CustomBool: float32(1), Error: <nil>
}

func ExampleFloat64() {
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

	// Output:
	// Bool: float64(1), Error: <nil>
	// String: float64(3.14159265359), Error: <nil>
	// Float64: float64(3.14159265359), Error: <nil>
	// Int: float64(3), Error: <nil>
	// CustomString: float64(3.14159265359), Error: <nil>
	// CustomInt: float64(3), Error: <nil>
	// CustomBool: float64(1), Error: <nil>
}

func ExampleFloat32FromUnsigned() {
	// Converting within range
	val, err := to.Float32FromUnsigned(uint(42))
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

	// Output:
	// float32(42), Error: <nil>
}

func ExampleFloat64FromUnsigned() {
	// Converting within range
	val, err := to.Float64FromUnsigned(uint64(42))
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

	// Output:
	// float64(42), Error: <nil>
}

func ExampleIntFromString() {
	// Valid conversion
	val, err := to.IntFromString("12345")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Invalid syntax
	valSyntax, errSyntax := to.IntFromString("abc")
	fmt.Printf("%T(%d), Error: %v\n", valSyntax, valSyntax, errors.Is(errSyntax, to.ErrInvalidStringSyntax))

	// Output:
	// int(12345), Error: <nil>
	// int(0), Error: true
}

func ExampleInt8FromString() {
	// Valid conversion
	val, err := to.Int8FromString("100")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Out of range
	valOOR, errOOR := to.Int8FromString("200")
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

	// Output:
	// int8(100), Error: <nil>
	// int8(0), Error: true
}

func ExampleInt16FromString() {
	// Valid conversion
	val, err := to.Int16FromString("12345")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Invalid syntax
	valSyntax, errSyntax := to.Int16FromString("abc")
	fmt.Printf("%T(%d), Error: %v\n", valSyntax, valSyntax, errors.Is(errSyntax, to.ErrInvalidStringSyntax))

	// Output:
	// int16(12345), Error: <nil>
	// int16(0), Error: true
}

func ExampleInt32FromString() {
	// Valid conversion
	val, err := to.Int32FromString("1234567")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int32(1234567), Error: <nil>
}

func ExampleInt64FromString() {
	// Valid conversion
	val, err := to.Int64FromString("9223372036854775807")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int64(9223372036854775807), Error: <nil>
}

func ExampleUIntFromString() {
	// Valid conversion
	val, err := to.UIntFromString("42")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// uint(42), Error: <nil>
}

func ExampleUInt8FromString() {
	// Valid conversion
	val, err := to.UInt8FromString("200")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// uint8(200), Error: <nil>
}

func ExampleUInt16FromString() {
	// Valid conversion
	val, err := to.UInt16FromString("65000")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// uint16(65000), Error: <nil>
}

func ExampleUInt32FromString() {
	// Valid conversion
	val, err := to.UInt32FromString("4294967295")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// uint32(4294967295), Error: <nil>
}

func ExampleUInt64FromString() {
	// Valid conversion
	val, err := to.UInt64FromString("18446744073709551615")
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// uint64(18446744073709551615), Error: <nil>
}

func ExampleFloat32FromString() {
	// Valid conversion
	val, err := to.Float32FromString("3.14159")
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

	// Output:
	// float32(3.14159), Error: <nil>
}

func ExampleFloat64FromString() {
	// Valid conversion
	val, err := to.Float64FromString("3.14159")
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

	// Output:
	// float64(3.14159), Error: <nil>
}

func ExampleValueBetween() {
	// Within range
	val := to.ValueBetween(5, 0, 10)
	fmt.Printf("ValueBetween(5, 0, 10) = %d\n", val)

	// Lower than min
	val = to.ValueBetween(-5, 0, 10)
	fmt.Printf("ValueBetween(-5, 0, 10) = %d\n", val)

	// Higher than max
	val = to.ValueBetween(15, 0, 10)
	fmt.Printf("ValueBetween(15, 0, 10) = %d\n", val)

	// Output:
	// ValueBetween(5, 0, 10) = 5
	// ValueBetween(-5, 0, 10) = 0
	// ValueBetween(15, 0, 10) = 10
}

func ExampleValueAtLeast() {
	// Already above min
	val := to.ValueAtLeast(5, 0)
	fmt.Printf("ValueAtLeast(5, 0) = %d\n", val)

	// Below min
	val = to.ValueAtLeast(-5, 0)
	fmt.Printf("ValueAtLeast(-5, 0) = %d\n", val)

	// Works with float values too
	floatVal := to.ValueAtLeast(3.14, 4.0)
	fmt.Printf("ValueAtLeast(3.14, 4.0) = %.2f\n", floatVal)

	// Output:
	// ValueAtLeast(5, 0) = 5
	// ValueAtLeast(-5, 0) = 0
	// ValueAtLeast(3.14, 4.0) = 4.00
}

func ExampleValueAtMost() {
	// Below max
	val := to.ValueAtMost(5, 10)
	fmt.Printf("ValueAtMost(5, 10) = %d\n", val)

	// Above max
	val = to.ValueAtMost(15, 10)
	fmt.Printf("ValueAtMost(15, 10) = %d\n", val)

	// Output:
	// ValueAtMost(5, 10) = 5
	// ValueAtMost(15, 10) = 10
}

func ExampleAtLeastThe() {
	// Creating a function that ensures values are at least 18
	ensureAdult := to.AtLeastThe(18)

	fmt.Printf("ensureAdult(15) = %d\n", ensureAdult(15))
	fmt.Printf("ensureAdult(21) = %d\n", ensureAdult(21))

	// Output:
	// ensureAdult(15) = 18
	// ensureAdult(21) = 21
}

func ExampleAtMostThe() {
	// Creating a function that ensures values are at most 100
	ensurePercentage := to.AtMostThe(100)

	fmt.Printf("ensurePercentage(50) = %d\n", ensurePercentage(50))
	fmt.Printf("ensurePercentage(150) = %d\n", ensurePercentage(150))

	// Output:
	// ensurePercentage(50) = 50
	// ensurePercentage(150) = 100
}

func ExampleValueBetweenThe() {
	// Creating a percentage validator (0-100)
	validatePercentage := to.ValueBetweenThe(0, 100)

	fmt.Printf("validatePercentage(-20) = %d\n", validatePercentage(-20))
	fmt.Printf("validatePercentage(50) = %d\n", validatePercentage(50))
	fmt.Printf("validatePercentage(150) = %d\n", validatePercentage(150))

	// Output:
	// validatePercentage(-20) = 0
	// validatePercentage(50) = 50
	// validatePercentage(150) = 100
}

func ExampleIntFromSigned() {
	// Converting within range
	val, err := to.IntFromSigned(int16(42))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int(42), Error: <nil>
}

func ExampleInt8FromSigned() {
	// Converting within range
	val, err := to.Int8FromSigned(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int8FromSigned(1000)
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

	// Output:
	// int8(42), Error: <nil>
	// int8(0), Error: true
}

func ExampleInt16FromSigned() {
	// Converting within range
	val, err := to.Int16FromSigned(1000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int16FromSigned(40000)
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

	// Output:
	// int16(1000), Error: <nil>
	// int16(0), Error: true
}

func ExampleInt32FromSigned() {
	// Converting within range
	val, err := to.Int32FromSigned(1000000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int32(1000000), Error: <nil>
}

func ExampleInt64FromSigned() {
	// Converting within range
	val, err := to.Int64FromSigned(9223372036854775807)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int64(9223372036854775807), Error: <nil>
}

func ExampleFloat32FromSigned() {
	// Converting within range
	val, err := to.Float32FromSigned(3)
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

	// Output:
	// float32(3), Error: <nil>
}

func ExampleFloat64FromSigned() {
	// Converting within range
	val, err := to.Float64FromSigned(3)
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

	// Output:
	// float64(3), Error: <nil>
}

package to_test

import (
	"errors"
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func ExampleInt() {
	// Converting within range
	val, err := to.Int(int16(42))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int(42), Error: <nil>
}

func ExampleIntFromUnsigned() {
	// Converting within range
	val, err := to.IntFromUnsigned(uint16(42))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int(42), Error: <nil>
}

func ExampleInt8() {
	// Converting within range
	val, err := to.Int8(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int8(1000)
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

	// Output:
	// int8(42), Error: <nil>
	// int8(0), Error: true
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
	// Converting within range
	val, err := to.Int16(1000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.Int16(40000)
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

	// Output:
	// int16(1000), Error: <nil>
	// int16(0), Error: true
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
	// Converting within range
	val, err := to.Int32(1000000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int32(1000000), Error: <nil>
}

func ExampleInt32FromUnsigned() {
	// Converting within range
	val, err := to.Int32FromUnsigned(uint(1000000))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int32(1000000), Error: <nil>
}

func ExampleInt64() {
	// Converting within range
	val, err := to.Int64(9223372036854775807)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int64(9223372036854775807), Error: <nil>
}

func ExampleInt64FromUnsigned() {
	// Converting within range
	val, err := to.Int64FromUnsigned(uint64(9223372036854775807))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// int64(9223372036854775807), Error: <nil>
}

func ExampleUInt() {
	// Converting positive value
	val, err := to.UInt(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting negative value
	valNeg, errNeg := to.UInt(-5)
	fmt.Printf("%T(%d), Error: %v\n", valNeg, valNeg, errors.Is(errNeg, to.ErrValueOutOfRange))

	// Output:
	// uint(42), Error: <nil>
	// uint(0), Error: true
}

func ExampleUInt8() {
	// Converting within range
	val, err := to.UInt8(200)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Converting out of range
	valOOR, errOOR := to.UInt8(300)
	fmt.Printf("%T(%d), Error: %v\n", valOOR, valOOR, errors.Is(errOOR, to.ErrValueOutOfRange))

	// Output:
	// uint8(200), Error: <nil>
	// uint8(0), Error: true
}

func ExampleUInt16() {
	// Converting within range
	val, err := to.UInt16(65000)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// uint16(65000), Error: <nil>
}

func ExampleUInt32() {
	// Valid conversion
	val, err := to.UInt32(42)
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Negative number
	valNeg, errNeg := to.UInt32(-5)
	fmt.Printf("%T(%d), Error: %v\n", valNeg, valNeg, errors.Is(errNeg, to.ErrValueOutOfRange))

	// Output:
	// uint32(42), Error: <nil>
	// uint32(0), Error: true
}

func ExampleUInt64() {
	// Converting within range
	val, err := to.UInt64(uint(18446744073709551000))
	fmt.Printf("%T(%d), Error: %v\n", val, val, err)

	// Output:
	// uint64(18446744073709551000), Error: <nil>
}

func ExampleFloat32() {
	// Converting within range
	val, err := to.Float32(3)
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

	// Output:
	// float32(3), Error: <nil>
}

func ExampleFloat64() {
	// Converting within range
	val, err := to.Float64(3)
	fmt.Printf("%T(%g), Error: %v\n", val, val, err)

	// Output:
	// float64(3), Error: <nil>
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

func ExampleClamped() {
	// Within range
	val := to.Clamped(5, 0, 10)
	fmt.Printf("Clamped(5, 0, 10) = %d\n", val)

	// Lower than min
	val = to.Clamped(-5, 0, 10)
	fmt.Printf("Clamped(-5, 0, 10) = %d\n", val)

	// Higher than max
	val = to.Clamped(15, 0, 10)
	fmt.Printf("Clamped(15, 0, 10) = %d\n", val)

	// Output:
	// Clamped(5, 0, 10) = 5
	// Clamped(-5, 0, 10) = 0
	// Clamped(15, 0, 10) = 10
}

func ExampleNoLessThan() {
	// Already above min
	val := to.NoLessThan(5, 0)
	fmt.Printf("NoLessThan(5, 0) = %d\n", val)

	// Below min
	val = to.NoLessThan(-5, 0)
	fmt.Printf("NoLessThan(-5, 0) = %d\n", val)

	// Works with float values too
	floatVal := to.NoLessThan(3.14, 4.0)
	fmt.Printf("NoLessThan(3.14, 4.0) = %.2f\n", floatVal)

	// Output:
	// NoLessThan(5, 0) = 5
	// NoLessThan(-5, 0) = 0
	// NoLessThan(3.14, 4.0) = 4.00
}

func ExampleNoMoreThan() {
	// Below max
	val := to.NoMoreThan(5, 10)
	fmt.Printf("NoMoreThan(5, 10) = %d\n", val)

	// Above max
	val = to.NoMoreThan(15, 10)
	fmt.Printf("NoMoreThan(15, 10) = %d\n", val)

	// Output:
	// NoMoreThan(5, 10) = 5
	// NoMoreThan(15, 10) = 10
}

func ExampleAtLeast() {
	// Creating a function that ensures values are at least 18
	ensureAdult := to.AtLeast(18)

	fmt.Printf("ensureAdult(15) = %d\n", ensureAdult(15))
	fmt.Printf("ensureAdult(21) = %d\n", ensureAdult(21))

	// Output:
	// ensureAdult(15) = 18
	// ensureAdult(21) = 21
}

func ExampleAtMost() {
	// Creating a function that ensures values are at most 100
	ensurePercentage := to.AtMost(100)

	fmt.Printf("ensurePercentage(50) = %d\n", ensurePercentage(50))
	fmt.Printf("ensurePercentage(150) = %d\n", ensurePercentage(150))

	// Output:
	// ensurePercentage(50) = 50
	// ensurePercentage(150) = 100
}

func ExampleClampedWith() {
	// Creating a percentage validator (0-100)
	validatePercentage := to.ClampedWith(0, 100)

	fmt.Printf("validatePercentage(-20) = %d\n", validatePercentage(-20))
	fmt.Printf("validatePercentage(50) = %d\n", validatePercentage(50))
	fmt.Printf("validatePercentage(150) = %d\n", validatePercentage(150))

	// Output:
	// validatePercentage(-20) = 0
	// validatePercentage(50) = 50
	// validatePercentage(150) = 100
}

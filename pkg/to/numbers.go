package to

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/go-softwarelab/common/pkg/types"
)

// ErrValueOutOfRange is returned when the value is out of range of the target type.
var ErrValueOutOfRange = fmt.Errorf("%w to convert", strconv.ErrRange)

// ErrInvalidStringSyntax is returned when the string has invalid syntax for conversion to target type.
var ErrInvalidStringSyntax = strconv.ErrSyntax

const (
	maxIntForUnsigned   = uint64(math.MaxInt)
	maxInt8ForUnsigned  = uint64(math.MaxInt8)
	maxInt16ForUnsigned = uint64(math.MaxInt16)
	maxInt32ForUnsigned = uint64(math.MaxInt32)
	maxInt64ForUnsigned = uint64(math.MaxInt64)
)

// Int will convert any integer to int, with range checks
func Int[V types.SignedNumber](value V) (int, error) {
	valToCompare := int64(value)
	if valToCompare < math.MinInt || valToCompare > math.MaxInt {
		return 0, fmt.Errorf("%v %w to int", value, ErrValueOutOfRange)
	}
	return int(value), nil
}

// IntFromUnsigned will convert any unsigned integer to int, with range checks.
func IntFromUnsigned[V types.Unsigned](value V) (int, error) {
	if uint64(value) > maxIntForUnsigned {
		return 0, fmt.Errorf("%d %w to int", value, ErrValueOutOfRange)
	}
	return int(value), nil
}

// Int8 will convert any integer to int8, with range checks
func Int8[V types.SignedNumber](value V) (int8, error) {
	valToCompare := int64(value)
	if valToCompare < math.MinInt8 || valToCompare > math.MaxInt8 {
		return 0, fmt.Errorf("%v %w to int8", value, ErrValueOutOfRange)
	}
	return int8(value), nil
}

// Int8FromUnsigned will convert any unsigned integer to int8, with range checks.
func Int8FromUnsigned[V types.Unsigned](value V) (int8, error) {
	if uint64(value) > maxInt8ForUnsigned {
		return 0, fmt.Errorf("%d %w to int8", value, ErrValueOutOfRange)
	}
	return int8(value), nil
}

// Int16 will convert any integer to int16, with range checks
func Int16[V types.SignedNumber](value V) (int16, error) {
	valToCompare := int64(value)
	if valToCompare < math.MinInt16 || valToCompare > math.MaxInt16 {
		return 0, fmt.Errorf("%v %w to int16", value, ErrValueOutOfRange)
	}
	return int16(value), nil
}

// Int16FromUnsigned will convert any unsigned integer to int16, with range checks.
func Int16FromUnsigned[V types.Unsigned](value V) (int16, error) {
	if uint64(value) > maxInt16ForUnsigned {
		return 0, fmt.Errorf("%d %w to int16", value, ErrValueOutOfRange)
	}
	return int16(value), nil
}

// Int32 will convert any integer to int32, with range checks
func Int32[V types.SignedNumber](value V) (int32, error) {
	valToCompare := int64(value)
	if valToCompare < math.MinInt32 || valToCompare > math.MaxInt32 {
		return 0, fmt.Errorf("%v %w to int32", value, ErrValueOutOfRange)
	}
	return int32(value), nil
}

// Int32FromUnsigned will convert any unsigned integer to int32, with range checks.
func Int32FromUnsigned[V types.Unsigned](value V) (int32, error) {
	if uint64(value) > maxInt32ForUnsigned {
		return 0, fmt.Errorf("%d %w to int32", value, ErrValueOutOfRange)
	}
	return int32(value), nil
}

// Int64 will convert any integer to int64, with range checks
func Int64[V types.SignedNumber](value V) (int64, error) {
	return int64(value), nil
}

// Int64FromUnsigned will convert any unsigned integer to int64, with range checks.
func Int64FromUnsigned[V types.Unsigned](value V) (int64, error) {
	if uint64(value) > maxInt64ForUnsigned {
		return 0, fmt.Errorf("%v %w to int64", value, ErrValueOutOfRange)
	}
	return int64(value), nil
}

// UInt will convert any integer to uint, with range checks
func UInt[V types.Number](value V) (uint, error) {
	if value < 0 || uint64(value) > math.MaxUint {
		return 0, fmt.Errorf("%v %w to uint", value, ErrValueOutOfRange)
	}
	return uint(value), nil
}

// UInt8 will convert any integer to uint8, with range checks
func UInt8[V types.Number](value V) (uint8, error) {
	if value < 0 || uint64(value) > math.MaxUint8 {
		return 0, fmt.Errorf("%v %w to uint8", value, ErrValueOutOfRange)
	}
	return uint8(value), nil
}

// UInt16 will convert any integer to uint16, with range checks
func UInt16[V types.Number](value V) (uint16, error) {
	if value < 0 || uint64(value) > math.MaxUint16 {
		return 0, fmt.Errorf("%v %w to uint16", value, ErrValueOutOfRange)
	}
	return uint16(value), nil
}

// UInt32 will convert any integer to uint32, with range checks
func UInt32[V types.Number](value V) (uint32, error) {
	if value < 0 || uint64(value) > math.MaxUint32 {
		return 0, fmt.Errorf("%v %w to uint32", value, ErrValueOutOfRange)
	}
	return uint32(value), nil
}

// UInt64 will convert any integer to uint64, with range checks
func UInt64[V types.Number](value V) (uint64, error) {
	if value < 0 {
		return 0, fmt.Errorf("%v %w to uint64", value, ErrValueOutOfRange)
	}
	return uint64(value), nil
}

// Float32 will convert any number to float
func Float32[V types.SignedNumber](value V) (float32, error) {
	if float64(value) < float64(math.SmallestNonzeroFloat32) || float64(value) > float64(math.MaxFloat32) {
		return 0, fmt.Errorf("%v %w to float32", value, ErrValueOutOfRange)
	}
	return float32(value), nil
}

// Float64 will convert any number to float
func Float64[V types.SignedNumber](value V) (float64, error) {
	return float64(value), nil
}

// Float32FromUnsigned will convert any unassigned number to float
func Float32FromUnsigned[V types.Unsigned](value V) (float32, error) {
	if float64(value) > float64(math.MaxFloat32) {
		return 0, fmt.Errorf("%v %w to float32", value, ErrValueOutOfRange)
	}
	return float32(value), nil
}

// Float64FromUnsigned will convert any unassigned number to float
func Float64FromUnsigned[V types.Unsigned](value V) (float64, error) {
	return float64(value), nil
}

// IntFromString will convert any string to int, with range checks
func IntFromString(value string) (int, error) {
	result, err := strconv.ParseInt(value, 10, strconv.IntSize)
	if err == nil {
		return int(result), nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to int", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// Int8FromString will convert any string to int8, with range checks
func Int8FromString(value string) (int8, error) {
	result, err := strconv.ParseInt(value, 10, 8)
	if err == nil {
		return int8(result), nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to int8", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// Int16FromString will convert any string to int16, with range checks
func Int16FromString(value string) (int16, error) {
	result, err := strconv.ParseInt(value, 10, 16)
	if err == nil {
		return int16(result), nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to int16", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// Int32FromString will convert any string to int32, with range checks
func Int32FromString(value string) (int32, error) {
	result, err := strconv.ParseInt(value, 10, 32)
	if err == nil {
		return int32(result), nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to int32", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// Int64FromString will convert any string to int64, with range checks
func Int64FromString(value string) (int64, error) {
	result, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		return result, nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to int64", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// UIntFromString will convert any string to uint, with range checks
func UIntFromString(value string) (uint, error) {
	result, err := strconv.ParseUint(value, 10, strconv.IntSize)
	if err == nil {
		return uint(result), nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to uint", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// UInt8FromString will convert any string to uint8, with range checks
func UInt8FromString(value string) (uint8, error) {
	result, err := strconv.ParseUint(value, 10, 8)
	if err == nil {
		return uint8(result), nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to uint8", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// UInt16FromString will convert any string to uint16, with range checks
func UInt16FromString(value string) (uint16, error) {
	result, err := strconv.ParseUint(value, 10, 16)
	if err == nil {
		return uint16(result), nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to uint16", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// UInt32FromString will convert any string to uint32, with range checks
func UInt32FromString(value string) (uint32, error) {
	result, err := strconv.ParseUint(value, 10, 32)
	if err == nil {
		return uint32(result), nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to uint32", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// UInt64FromString will convert any string to uint64, with range checks
func UInt64FromString(value string) (uint64, error) {
	result, err := strconv.ParseUint(value, 10, 64)
	if err == nil {
		return result, nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to uint64", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// Float32FromString will convert any string to float32, with range checks
func Float32FromString(value string) (float32, error) {
	result, err := strconv.ParseFloat(value, 32)
	if err == nil {
		return float32(result), nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to float32", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// Float64FromString will convert any string to float64, with range checks
func Float64FromString(value string) (float64, error) {
	result, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return result, nil
	} else if errors.Is(err, strconv.ErrSyntax) {
		return 0, fmt.Errorf("%w of %s to parse into number", ErrInvalidStringSyntax, value)
	} else if errors.Is(err, strconv.ErrRange) {
		return 0, fmt.Errorf("%s %w to float64", value, ErrValueOutOfRange)
	}
	return 0, fmt.Errorf("%w of %s: %w", ErrInvalidStringSyntax, value, err)
}

// NoLessThan will return the value if it's not less than the min value or the min value.
func NoLessThan[T types.Ordered](value, min T) T {
	if value < min {
		return min
	}
	return value
}

// AtLeast will return a function that will clamp the value to be at least the min value.
// It's wrapped around NoLessThan function, to make it usable in Map functions.
//
// See Also: NoLessThan
func AtLeast[T types.Ordered](min T) func(value T) T {
	return func(value T) T {
		return NoLessThan(value, min)
	}
}

// NoMoreThan will return the value if it's not more than the max value or the max value.
func NoMoreThan[T types.Ordered](value, max T) T {
	if value > max {
		return max
	}
	return value
}

// AtMost will return a function that will clamp the value to be at most the max value.
// It's wrapped around NoMoreThan function, to make it usable in Map functions.
//
// See Also: NoMoreThan
func AtMost[T types.Ordered](max T) func(value T) T {
	return func(value T) T {
		return NoMoreThan(value, max)
	}
}

// ValueBetween will clamp the value between the min and max values.
// In other words it ensures that result is min <= value <= max.
// For value that is less than min, it will return min.
// For value that is greater than max, it will return max.
//
// See Also: ValueBetweenThe to use it in Map functions.
func ValueBetween[T types.Ordered](value, min, max T) T {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

// ValueBetweenThe returns a function that clamps the value between the min and max values.
// In other words it ensures that result is min <= value <= max.
// For value that is less than min, it will return min.
// For value that is greater than max, it will return max.
// It's wrapped around ValueBetween function, to make it usable in Map functions.
//
// See Also: ValueBetween
func ValueBetweenThe[T types.Ordered](min, max T) func(value T) T {
	return func(value T) T {
		return ValueBetween(value, min, max)
	}
}

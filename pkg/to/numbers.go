package to

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"

	"github.com/go-softwarelab/common/pkg/types"
)

// ConvertableToNumber is a constraint that permits types that can be converted to number, such as strings, numbers, or booleans.
type ConvertableToNumber interface {
	~string | types.Number | ~bool
}

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

// Int will convert bool, any number or string, and their subtypes to int
//
// NOTE: This is using reflection,
//
//	if it is an issue, use dedicated functions for given type, with suffix like FromBool or FromSigned
func Int[V ConvertableToNumber](value V) (int, error) {
	switch v := any(value).(type) {
	case bool:
		return IntFromBool(v), nil
	case int:
		return IntFromSigned(v)
	case int8:
		return IntFromSigned(v)
	case int16:
		return IntFromSigned(v)
	case int32:
		return IntFromSigned(v)
	case int64:
		return IntFromSigned(v)
	case float32:
		return IntFromSigned(v)
	case float64:
		return IntFromSigned(v)
	case uint:
		return IntFromUnsigned(v)
	case uint8:
		return IntFromUnsigned(v)
	case uint16:
		return IntFromUnsigned(v)
	case uint32:
		return IntFromUnsigned(v)
	case uint64:
		return IntFromUnsigned(v)
	case string:
		return IntFromString(v)
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Bool:
		return IntFromBool(v.Bool()), nil
	case reflect.Int:
		return IntFromSigned(v.Int())
	case reflect.Int8:
		return IntFromSigned(v.Int())
	case reflect.Int16:
		return IntFromSigned(v.Int())
	case reflect.Int32:
		return IntFromSigned(v.Int())
	case reflect.Int64:
		return IntFromSigned(v.Int())
	case reflect.Float32:
		return IntFromSigned(v.Float())
	case reflect.Float64:
		return IntFromSigned(v.Float())
	case reflect.Uint:
		return IntFromUnsigned(v.Uint())
	case reflect.Uint8:
		return IntFromUnsigned(v.Uint())
	case reflect.Uint16:
		return IntFromUnsigned(v.Uint())
	case reflect.Uint32:
		return IntFromUnsigned(v.Uint())
	case reflect.Uint64:
		return IntFromUnsigned(v.Uint())
	case reflect.String:
		return IntFromString(v.String())
	default:
		panic(fmt.Sprintf("unsupported type %T", value))
	}

}

// IntFromBool converts a boolean value to its integer representation (true = 1, false = 0).
func IntFromBool(value bool) int {
	if value {
		return 1
	}
	return 0
}

// IntFromSigned will convert any signed number to int, with range checks
func IntFromSigned[V types.SignedNumber](value V) (int, error) {
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

// Int8 will convert any integer to int8, with range checks
func Int8[V types.SignedNumber](value V) (int8, error) {
	valToCompare := int64(value)
	if valToCompare < math.MinInt8 || valToCompare > math.MaxInt8 {
		return 0, fmt.Errorf("%v %w to int8", value, ErrValueOutOfRange)
	}
	return int8(value), nil
}

// Int8FromSigned will convert any signed number to int8, with range checks
func Int8FromSigned[V types.SignedNumber](value V) (int8, error) {
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

// Int16 will convert any integer to int16, with range checks
func Int16[V types.SignedNumber](value V) (int16, error) {
	valToCompare := int64(value)
	if valToCompare < math.MinInt16 || valToCompare > math.MaxInt16 {
		return 0, fmt.Errorf("%v %w to int16", value, ErrValueOutOfRange)
	}
	return int16(value), nil
}

// Int16FromSigned will convert any signed number to int16, with range checks
func Int16FromSigned[V types.SignedNumber](value V) (int16, error) {
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

// Int32 will convert any integer to int32, with range checks
func Int32[V types.SignedNumber](value V) (int32, error) {
	valToCompare := int64(value)
	if valToCompare < math.MinInt32 || valToCompare > math.MaxInt32 {
		return 0, fmt.Errorf("%v %w to int32", value, ErrValueOutOfRange)
	}
	return int32(value), nil
}

// Int32FromSigned will convert any signed number to int32, with range checks
func Int32FromSigned[V types.SignedNumber](value V) (int32, error) {
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

// Int64 will convert any integer to int64, with range checks
func Int64[V types.SignedNumber](value V) (int64, error) {
	return int64(value), nil
}

// Int64FromSigned will convert any signed number to int64, with range checks
func Int64FromSigned[V types.SignedNumber](value V) (int64, error) {
	return int64(value), nil
}

// Int64FromUnsigned will convert any unsigned integer to int64, with range checks.
func Int64FromUnsigned[V types.Unsigned](value V) (int64, error) {
	if uint64(value) > maxInt64ForUnsigned {
		return 0, fmt.Errorf("%v %w to int64", value, ErrValueOutOfRange)
	}
	return int64(value), nil
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

// UInt will convert any integer to uint, with range checks
func UInt[V types.Number](value V) (uint, error) {
	if value < 0 || uint64(value) > math.MaxUint {
		return 0, fmt.Errorf("%v %w to uint", value, ErrValueOutOfRange)
	}
	return uint(value), nil
}

// UIntFromNumber will convert any number to uint, with range checks
func UIntFromNumber[V types.Number](value V) (uint, error) {
	if value < 0 || uint64(value) > math.MaxUint {
		return 0, fmt.Errorf("%v %w to uint", value, ErrValueOutOfRange)
	}
	return uint(value), nil
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

// UInt8 will convert any integer to uint8, with range checks
func UInt8[V types.Number](value V) (uint8, error) {
	if value < 0 || uint64(value) > math.MaxUint8 {
		return 0, fmt.Errorf("%v %w to uint8", value, ErrValueOutOfRange)
	}
	return uint8(value), nil
}

// UInt8FromNumber will convert any number to uint8, with range checks
func UInt8FromNumber[V types.Number](value V) (uint8, error) {
	if value < 0 || uint64(value) > math.MaxUint8 {
		return 0, fmt.Errorf("%v %w to uint8", value, ErrValueOutOfRange)
	}
	return uint8(value), nil
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

// UInt16 will convert any integer to uint16, with range checks
func UInt16[V types.Number](value V) (uint16, error) {
	if value < 0 || uint64(value) > math.MaxUint16 {
		return 0, fmt.Errorf("%v %w to uint16", value, ErrValueOutOfRange)
	}
	return uint16(value), nil
}

// UInt16FromNumber will convert any number to uint16, with range checks
func UInt16FromNumber[V types.Number](value V) (uint16, error) {
	if value < 0 || uint64(value) > math.MaxUint16 {
		return 0, fmt.Errorf("%v %w to uint16", value, ErrValueOutOfRange)
	}
	return uint16(value), nil
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

// UInt32 will convert any integer to uint32, with range checks
func UInt32[V types.Number](value V) (uint32, error) {
	if value < 0 || uint64(value) > math.MaxUint32 {
		return 0, fmt.Errorf("%v %w to uint32", value, ErrValueOutOfRange)
	}
	return uint32(value), nil
}

// UInt32FromNumber will convert any number to uint32, with range checks
func UInt32FromNumber[V types.Number](value V) (uint32, error) {
	if value < 0 || uint64(value) > math.MaxUint32 {
		return 0, fmt.Errorf("%v %w to uint32", value, ErrValueOutOfRange)
	}
	return uint32(value), nil
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

// UInt64 will convert any integer to uint64, with range checks
func UInt64[V types.Number](value V) (uint64, error) {
	if value < 0 {
		return 0, fmt.Errorf("%v %w to uint64", value, ErrValueOutOfRange)
	}
	return uint64(value), nil
}

// UInt64FromNumber will convert any number to uint64, with range checks
func UInt64FromNumber[V types.Number](value V) (uint64, error) {
	if value < 0 {
		return 0, fmt.Errorf("%v %w to uint64", value, ErrValueOutOfRange)
	}
	return uint64(value), nil
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

// Float32 will convert any number to float
func Float32[V types.SignedNumber](value V) (float32, error) {
	if float64(value) < float64(math.SmallestNonzeroFloat32) || float64(value) > float64(math.MaxFloat32) {
		return 0, fmt.Errorf("%v %w to float32", value, ErrValueOutOfRange)
	}
	return float32(value), nil
}

// Float32FromSigned will convert any signed number to float32, with range checks
func Float32FromSigned[V types.SignedNumber](value V) (float32, error) {
	if float64(value) < float64(math.SmallestNonzeroFloat32) || float64(value) > float64(math.MaxFloat32) {
		return 0, fmt.Errorf("%v %w to float32", value, ErrValueOutOfRange)
	}
	return float32(value), nil
}

// Float32FromUnsigned will convert any unassigned number to float
func Float32FromUnsigned[V types.Unsigned](value V) (float32, error) {
	if float64(value) > float64(math.MaxFloat32) {
		return 0, fmt.Errorf("%v %w to float32", value, ErrValueOutOfRange)
	}
	return float32(value), nil
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

// Float64 will convert any number to float
func Float64[V types.SignedNumber](value V) (float64, error) {
	return float64(value), nil
}

// Float64FromSigned will convert any signed number to float64
func Float64FromSigned[V types.SignedNumber](value V) (float64, error) {
	return float64(value), nil
}

// Float64FromUnsigned will convert any unassigned number to float
func Float64FromUnsigned[V types.Unsigned](value V) (float64, error) {
	return float64(value), nil
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
	return ValueAtLeast(value, min)
}

// ValueAtLeast will return the value if it's not less than the min value or the min value.
func ValueAtLeast[T types.Ordered](value, min T) T {
	if value < min {
		return min
	}
	return value
}

// AtLeast will return a function that will clamp the value to be at least the min value.
// It's wrapped around NoLessThan function, to make it usable in Map functions.
//
// See Also: NoLessThan
// @Deprecated: In the future will replace ValueAtLeast, use AtLeastThe instead.
func AtLeast[T types.Ordered](min T) func(value T) T {
	return func(value T) T {
		return ValueAtLeast(value, min)
	}
}

// AtLeastThe will return a function that will clamp the value to be at least the min value.
// It's wrapped around ValueAtLeast function, to make it usable in Map functions.
//
// See Also: ValueAtLeast
func AtLeastThe[T types.Ordered](min T) func(value T) T {
	return func(value T) T {
		return ValueAtLeast(value, min)
	}
}

// NoMoreThan will return the value if it's not more than the max value or the max value.
func NoMoreThan[T types.Ordered](value, max T) T {
	return ValueAtMost(value, max)
}

// ValueAtMost will return the value if it's not more than the max value or the max value.
func ValueAtMost[T types.Ordered](value, max T) T {
	if value > max {
		return max
	}
	return value
}

// AtMost will return a function that will clamp the value to be at most the max value.
// It's wrapped around ValueAtMost function, to make it usable in Map functions.
//
// See Also: ValueAtMost
// @Deprecated: In the future will replace ValueAtMost, use AtMostThe instead.
func AtMost[T types.Ordered](max T) func(value T) T {
	return func(value T) T {
		return ValueAtMost(value, max)
	}
}

// AtMostThe will return a function that will clamp the value to be at most the max value.
// It's wrapped around ValueAtMost function, to make it usable in Map functions.
//
// See Also: ValueAtMost
func AtMostThe[T types.Ordered](max T) func(value T) T {
	return func(value T) T {
		return ValueAtMost(value, max)
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

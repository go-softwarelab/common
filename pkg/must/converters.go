package must

import (
	"github.com/go-softwarelab/common/pkg/to"
	"github.com/go-softwarelab/common/pkg/types"
)

// ConvertToInt converts any signed number to int, panicking on range errors
func ConvertToInt[V types.SignedNumber](value V) int {
	return getWithNoError(to.Int(value))
}

// ConvertToIntFromUnsigned converts any unsigned number to int, panicking on range errors
func ConvertToIntFromUnsigned[V types.Unsigned](value V) int {
	return getWithNoError(to.IntFromUnsigned(value))
}

// ConvertToInt8 converts any signed number to int8, panicking on range errors
func ConvertToInt8[V types.SignedNumber](value V) int8 {
	return getWithNoError(to.Int8(value))
}

// ConvertToInt8FromUnsigned converts any unsigned number to int8, panicking on range errors
func ConvertToInt8FromUnsigned[V types.Unsigned](value V) int8 {
	return getWithNoError(to.Int8FromUnsigned(value))
}

// ConvertToInt16 converts any signed number to int16, panicking on range errors
func ConvertToInt16[V types.SignedNumber](value V) int16 {
	return getWithNoError(to.Int16(value))
}

// ConvertToInt16FromUnsigned converts any unsigned number to int16, panicking on range errors
func ConvertToInt16FromUnsigned[V types.Unsigned](value V) int16 {
	return getWithNoError(to.Int16FromUnsigned(value))
}

// ConvertToInt32 converts any signed number to int32, panicking on range errors
func ConvertToInt32[V types.SignedNumber](value V) int32 {
	return getWithNoError(to.Int32(value))
}

// ConvertToInt32FromUnsigned converts any unsigned number to int32, panicking on range errors
func ConvertToInt32FromUnsigned[V types.Unsigned](value V) int32 {
	return getWithNoError(to.Int32FromUnsigned(value))
}

// ConvertToInt64 converts any signed number to int64, panicking on range errors
func ConvertToInt64[V types.SignedNumber](value V) int64 {
	return getWithNoError(to.Int64(value))
}

// ConvertToInt64FromUnsigned converts any unsigned number to int64, panicking on range errors
func ConvertToInt64FromUnsigned[V types.Unsigned](value V) int64 {
	return getWithNoError(to.Int64FromUnsigned(value))
}

// ConvertToUInt converts any number to uint, panicking on range errors
func ConvertToUInt[V types.Number](value V) uint {
	return getWithNoError(to.UInt(value))
}

// ConvertToUInt8 converts any number to uint8, panicking on range errors
func ConvertToUInt8[V types.Number](value V) uint8 {
	return getWithNoError(to.UInt8(value))
}

// ConvertToUInt16 converts any number to uint16, panicking on range errors
func ConvertToUInt16[V types.Number](value V) uint16 {
	return getWithNoError(to.UInt16(value))
}

// ConvertToUInt32 converts any number to uint32, panicking on range errors
func ConvertToUInt32[V types.Number](value V) uint32 {
	return getWithNoError(to.UInt32(value))
}

// ConvertToUInt64 converts any number to uint64, panicking on range errors
func ConvertToUInt64[V types.Number](value V) uint64 {
	return getWithNoError(to.UInt64(value))
}

// ConvertToFloat32 converts any signed number to float32, panicking on range errors
func ConvertToFloat32[V types.SignedNumber](value V) float32 {
	return getWithNoError(to.Float32(value))
}

// ConvertToFloat64 converts any signed number to float64
func ConvertToFloat64[V types.SignedNumber](value V) float64 {
	return getWithNoError(to.Float64(value))
}

// ConvertToFloat32FromUnsigned converts any unsigned number to float32, panicking on range errors
func ConvertToFloat32FromUnsigned[V types.Unsigned](value V) float32 {
	return getWithNoError(to.Float32FromUnsigned(value))
}

// ConvertToFloat64FromUnsigned converts any unsigned number to float64
func ConvertToFloat64FromUnsigned[V types.Unsigned](value V) float64 {
	return getWithNoError(to.Float64FromUnsigned(value))
}

// ConvertToIntFromString converts a string to int, panicking in case if the string is not a valid number.
func ConvertToIntFromString(value string) int {
	return getWithNoError(to.IntFromString(value))
}

// ConvertToInt8FromString converts a string to int8, panicking in case if the string is not a valid number.
func ConvertToInt8FromString(value string) int8 {
	return getWithNoError(to.Int8FromString(value))
}

// ConvertToInt16FromString converts a string to int16, panicking in case if the string is not a valid number.
func ConvertToInt16FromString(value string) int16 {
	return getWithNoError(to.Int16FromString(value))
}

// ConvertToInt32FromString converts a string to int32, panicking in case if the string is not a valid number.
func ConvertToInt32FromString(value string) int32 {
	return getWithNoError(to.Int32FromString(value))
}

// ConvertToInt64FromString converts a string to int64, panicking in case if the string is not a valid number.
func ConvertToInt64FromString(value string) int64 {
	return getWithNoError(to.Int64FromString(value))
}

// ConvertToUIntFromString converts a string to uint, panicking in case if the string is not a valid number.
func ConvertToUIntFromString(value string) uint {
	return getWithNoError(to.UIntFromString(value))
}

// ConvertToUInt8FromString converts a string to uint8, panicking in case if the string is not a valid number.
func ConvertToUInt8FromString(value string) uint8 {
	return getWithNoError(to.UInt8FromString(value))
}

// ConvertToUInt16FromString converts a string to uint16, panicking in case if the string is not a valid number.
func ConvertToUInt16FromString(value string) uint16 {
	return getWithNoError(to.UInt16FromString(value))
}

// ConvertToUInt32FromString converts a string to uint32, panicking in case if the string is not a valid number.
func ConvertToUInt32FromString(value string) uint32 {
	return getWithNoError(to.UInt32FromString(value))
}

// ConvertToUInt64FromString converts a string to uint64, panicking in case if the string is not a valid number.
func ConvertToUInt64FromString(value string) uint64 {
	return getWithNoError(to.UInt64FromString(value))
}

// ConvertToFloat32FromString converts a string to float32, panicking in case if the string is not a valid number.
func ConvertToFloat32FromString(value string) float32 {
	return getWithNoError(to.Float32FromString(value))
}

// ConvertToFloat64FromString converts a string to float64, panicking in case if the string is not a valid number.
func ConvertToFloat64FromString(value string) float64 {
	return getWithNoError(to.Float64FromString(value))
}

func getWithNoError[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

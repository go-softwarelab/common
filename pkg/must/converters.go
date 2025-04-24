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

func getWithNoError[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

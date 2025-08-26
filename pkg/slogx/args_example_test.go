package slogx_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func ExampleError() {
	err := fmt.Errorf("test error")
	fmt.Println(slogx.Error(err))
	// Output:
	// error=test error
}

func ExampleNumber_uint() {
	fmt.Println(slogx.Number("key", uint(42)))
	fmt.Println(slogx.Number("key", uint8(42)))
	fmt.Println(slogx.Number("key", uint16(42)))
	fmt.Println(slogx.Number("key", uint32(42)))
	fmt.Println(slogx.Number("key", uint64(42)))
	// Output:
	// key=42
	// key=42
	// key=42
	// key=42
	// key=42
}

func ExampleNumber_int() {
	fmt.Println(slogx.Number("key", 42))
	fmt.Println(slogx.Number("key", int8(42)))
	fmt.Println(slogx.Number("key", int16(42)))
	fmt.Println(slogx.Number("key", int32(42)))
	fmt.Println(slogx.Number("key", int64(42)))
	// Output:
	// key=42
	// key=42
	// key=42
	// key=42
	// key=42
}

func ExampleNumber_float() {
	fmt.Println(slogx.Number("key", float32(42.5)))
	fmt.Println(slogx.Number("key", float64(42.5)))
	// Output:
	// key=42.5
	// key=42.5
}

func ExampleNumber_customType() {
	type MyInt int
	type MyFloat float64
	type MyUint uint

	fmt.Println(slogx.Number("key", MyInt(42)))
	fmt.Println(slogx.Number("key", MyFloat(42.5)))
	fmt.Println(slogx.Number("key", MyUint(42)))
	// Output:
	// key=42
	// key=42.5
	// key=42
}

func ExampleOptionalNumber_nil() {
	var n *int
	fmt.Println(slogx.OptionalNumber("key", n))
	// Output:
	// key=<nil>
}

func ExampleOptionalNumber_int() {
	n := 42
	fmt.Println(slogx.OptionalNumber("key", &n))
	// Output:
	// key=42
}

func ExampleOptionalNumber_float() {
	n := 42.5
	fmt.Println(slogx.OptionalNumber("key", &n))
	// Output:
	// key=42.5
}

func ExampleOptionalNumber_customType() {
	type MyInt int
	n := MyInt(42)
	fmt.Println(slogx.OptionalNumber("key", &n))
	// Output:
	// key=42
}

func ExampleString_basic() {
	fmt.Println(slogx.String("key", "value"))
	fmt.Println(slogx.String("empty", ""))
	// Output:
	// key=value
	// empty=
}

func ExampleString_customType() {
	type MyString string
	s := MyString("custom value")
	fmt.Println(slogx.String("key", s))
	// Output:
	// key=custom value
}

func ExampleString_empty() {
	var s string
	fmt.Println(slogx.String("key", s))
	// Output:
	// key=
}

func ExampleOptionalString_nil() {
	var s *string
	fmt.Println(slogx.OptionalString("key", s))
	// Output:
	// key=<nil>
}

func ExampleOptionalString_basic() {
	s := "value"
	fmt.Println(slogx.OptionalString("key", &s))
	// Output:
	// key=value
}

func ExampleOptionalString_customType() {
	type MyString string
	s := MyString("custom value")
	fmt.Println(slogx.OptionalString("key", &s))
	// Output:
	// key=custom value
}

func ExampleUserID_int() {
	fmt.Println(slogx.UserID(42))
	// Output:
	// userId=42
}

func ExampleUserID_string() {
	fmt.Println(slogx.UserID("user123"))
	// Output:
	// userId=user123
}

func ExampleUserID_customType() {
	type UserIDInt int
	type UserIDString string

	fmt.Println(slogx.UserID(UserIDInt(42)))
	fmt.Println(slogx.UserID(UserIDString("user123")))
	// Output:
	// userId=42
	// userId=user123
}

func ExampleOptionalUserID_nil() {
	var id *string
	fmt.Println(slogx.OptionalUserID(id))
	// Output:
	// userId=<unknown>
}

func ExampleOptionalUserID_int() {
	id := 42
	fmt.Println(slogx.OptionalUserID(&id))
	// Output:
	// userId=42
}

func ExampleOptionalUserID_string() {
	id := "user123"
	fmt.Println(slogx.OptionalUserID(&id))
	// Output:
	// userId=user123
}

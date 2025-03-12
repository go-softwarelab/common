package to_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func ExamplePtr() {
	// Create a pointer to a value
	val := to.Ptr(42)
	fmt.Printf("%T(%d)\n", val, *val)

	// Works with strings too
	strVal := to.Ptr("hello")
	fmt.Printf("%T(%q)\n", strVal, *strVal)

	// Output:
	// *int(42)
	// *string("hello")
}

func ExampleNil() {
	// Create a nil pointer of int type
	val := to.Nil[int]()
	fmt.Printf("%T, IsNil: %v\n", val, val == nil)

	// Create a nil pointer of a struct type
	type Person struct {
		Name string
		Age  int
	}
	person := to.Nil[Person]()
	fmt.Printf("%T, IsNil: %v\n", person, person == nil)

	// Output:
	// *int, IsNil: true
	// *to_test.Person, IsNil: true
}

func ExampleNilOfType() {
	// Create a nil pointer based on existing pointer type
	var existingPtr *string
	nilPtr := to.NilOfType(existingPtr)
	fmt.Printf("%T, IsNil: %v\n", nilPtr, nilPtr == nil)

	// Output:
	// *string, IsNil: true
}

func ExampleEmptyValue() {
	// Get zero value of int
	val := to.EmptyValue[int]()
	fmt.Printf("%T(%d)\n", val, val)

	// Get zero value of string
	strVal := to.EmptyValue[string]()
	fmt.Printf("%T(%q)\n", strVal, strVal)

	// Get zero value of bool
	boolVal := to.EmptyValue[bool]()
	fmt.Printf("%T(%v)\n", boolVal, boolVal)

	// Output:
	// int(0)
	// string("")
	// bool(false)
}

func ExampleZeroValue() {
	// Get zero value of int
	val := to.ZeroValue[int]()
	fmt.Printf("%T(%d)\n", val, val)

	// Get zero value of string
	strVal := to.ZeroValue[string]()
	fmt.Printf("%T(%q)\n", strVal, strVal)

	// Output:
	// int(0)
	// string("")
}

func ExampleValue() {
	// Get value from a pointer
	ptr := to.Ptr(42)
	val := to.Value(ptr)
	fmt.Printf("%T(%d)\n", val, val)

	// Get value from a nil pointer (returns zero value)
	var nilPtr *string
	strVal := to.Value(nilPtr)
	fmt.Printf("%T(%q)\n", strVal, strVal)

	// Output:
	// int(42)
	// string("")
}

func ExampleValueOr() {
	// Get value from a pointer
	ptr := to.Ptr(42)
	val := to.ValueOr(ptr, 0)
	fmt.Printf("%T(%d)\n", val, val)

	// Get fallback value from a nil pointer
	var nilPtr *string
	strVal := to.ValueOr(nilPtr, "fallback")
	fmt.Printf("%T(%q)\n", strVal, strVal)

	// Output:
	// int(42)
	// string("fallback")
}

func ExampleSliceOfPtr() {
	// Convert a slice of values to a slice of pointers
	values := []int{1, 2, 3}
	ptrs := to.SliceOfPtr(values)

	// Print first element to demonstrate it's a pointer
	fmt.Printf("%T(%d)\n", ptrs[0], *ptrs[0])

	// Print length to show all elements were converted
	fmt.Printf("Length: %d\n", len(ptrs))

	// Output:
	// *int(1)
	// Length: 3
}

func ExampleSliceOfValue() {
	// Convert a slice of pointers to a slice of values
	p1, p2 := to.Ptr(10), to.Ptr(20)
	ptrs := []*int{p1, p2, nil}
	values := to.SliceOfValue(ptrs)

	fmt.Printf("%T(%d)\n", values[0], values[0])
	fmt.Printf("%T(%d)\n", values[1], values[1])
	fmt.Printf("%T(%d) (zero value from nil)\n", values[2], values[2])

	// Output:
	// int(10)
	// int(20)
	// int(0) (zero value from nil)
}

func ExampleAny() {
	// Convert a typed value to any
	val := to.Any(42)
	fmt.Printf("Type: %T\n", val)

	// Convert a string to any
	strVal := to.Any("hello")
	fmt.Printf("Type: %T\n", strVal)

	// Output:
	// Type: int
	// Type: string
}

func ExampleSliceOfAny() {
	// Convert a slice of ints to a slice of any
	values := []int{1, 2, 3}
	anySlice := to.SliceOfAny(values)

	fmt.Printf("Type: %T, Length: %d\n", anySlice, len(anySlice))
	fmt.Printf("First element type: %T\n", anySlice[0])

	// Output:
	// Type: []interface {}, Length: 3
	// First element type: int
}

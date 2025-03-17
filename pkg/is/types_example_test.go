package is_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func ExampleNil() {
	// Testing nil values
	fmt.Println("Checking for nil values:")

	var nilSlice []int
	fmt.Printf("is.Nil(nil): %T(%v) - literal nil\n", is.Nil(nil), is.Nil(nil))
	fmt.Printf("is.Nil(nilSlice): %T(%v) - nil slice\n", is.Nil(nilSlice), is.Nil(nilSlice))

	nonNilSlice := make([]int, 0)
	fmt.Printf("is.Nil(nonNilSlice): %T(%v) - empty but initialized slice\n", is.Nil(nonNilSlice), is.Nil(nonNilSlice))

	var nilMap map[string]int
	fmt.Printf("is.Nil(nilMap): %T(%v) - nil map\n", is.Nil(nilMap), is.Nil(nilMap))

	// Output:
	// Checking for nil values:
	// is.Nil(nil): bool(true) - literal nil
	// is.Nil(nilSlice): bool(true) - nil slice
	// is.Nil(nonNilSlice): bool(false) - empty but initialized slice
	// is.Nil(nilMap): bool(true) - nil map
}

func ExampleNotNil() {
	// Testing non-nil values
	fmt.Println("Checking for non-nil values:")

	var nilSlice []int
	nonNilSlice := make([]int, 0)

	fmt.Printf("is.NotNil(nonNilSlice): %T(%v) - empty but initialized slice\n", is.NotNil(nonNilSlice), is.NotNil(nonNilSlice))
	fmt.Printf("is.NotNil(nilSlice): %T(%v) - nil slice\n", is.NotNil(nilSlice), is.NotNil(nilSlice))
	fmt.Printf("is.NotNil(\"hello\"): %T(%v) - string value\n", is.NotNil("hello"), is.NotNil("hello"))
	fmt.Printf("is.NotNil(nil): %T(%v) - literal nil\n", is.NotNil(nil), is.NotNil(nil))

	// Output:
	// Checking for non-nil values:
	// is.NotNil(nonNilSlice): bool(true) - empty but initialized slice
	// is.NotNil(nilSlice): bool(false) - nil slice
	// is.NotNil("hello"): bool(true) - string value
	// is.NotNil(nil): bool(false) - literal nil
}

func ExampleEmpty() {
	// Testing zero/empty values of different types
	fmt.Println("Checking for empty/zero values:")

	fmt.Printf("is.Empty(0): %T(%v) - zero int\n", is.Empty(0), is.Empty(0))
	fmt.Printf("is.Empty(42): %T(%v) - non-zero int\n", is.Empty(42), is.Empty(42))

	fmt.Printf("is.Empty(\"\"): %T(%v) - empty string\n", is.Empty(""), is.Empty(""))
	fmt.Printf("is.Empty(\"hello\"): %T(%v) - non-empty string\n", is.Empty("hello"), is.Empty("hello"))

	fmt.Printf("is.Empty(false): %T(%v) - zero bool\n", is.Empty(false), is.Empty(false))
	fmt.Printf("is.Empty(true): %T(%v) - non-zero bool\n", is.Empty(true), is.Empty(true))

	// Output:
	// Checking for empty/zero values:
	// is.Empty(0): bool(true) - zero int
	// is.Empty(42): bool(false) - non-zero int
	// is.Empty(""): bool(true) - empty string
	// is.Empty("hello"): bool(false) - non-empty string
	// is.Empty(false): bool(true) - zero bool
	// is.Empty(true): bool(false) - non-zero bool
}

func ExampleNotEmpty() {
	// Testing non-empty values of different types
	fmt.Println("Checking for non-empty values:")

	fmt.Printf("is.NotEmpty(42): %T(%v) - non-zero int\n", is.NotEmpty(42), is.NotEmpty(42))
	fmt.Printf("is.NotEmpty(0): %T(%v) - zero int\n", is.NotEmpty(0), is.NotEmpty(0))

	fmt.Printf("is.NotEmpty(\"hello\"): %T(%v) - non-empty string\n", is.NotEmpty("hello"), is.NotEmpty("hello"))
	fmt.Printf("is.NotEmpty(\"\"): %T(%v) - empty string\n", is.NotEmpty(""), is.NotEmpty(""))

	fmt.Printf("is.NotEmpty(true): %T(%v) - non-zero bool\n", is.NotEmpty(true), is.NotEmpty(true))
	fmt.Printf("is.NotEmpty(false): %T(%v) - zero bool\n", is.NotEmpty(false), is.NotEmpty(false))

	// Output:
	// Checking for non-empty values:
	// is.NotEmpty(42): bool(true) - non-zero int
	// is.NotEmpty(0): bool(false) - zero int
	// is.NotEmpty("hello"): bool(true) - non-empty string
	// is.NotEmpty(""): bool(false) - empty string
	// is.NotEmpty(true): bool(true) - non-zero bool
	// is.NotEmpty(false): bool(false) - zero bool
}

func ExampleType() {
	// Testing type checking
	fmt.Println("Type checking:")

	value1 := "hello"
	value2 := 42
	value3 := true

	fmt.Printf("is.Type[string](value1): %T(%v) - string value\n", is.Type[string](value1), is.Type[string](value1))
	fmt.Printf("is.Type[int](value1): %T(%v) - string value checked as int\n", is.Type[int](value1), is.Type[int](value1))

	fmt.Printf("is.Type[int](value2): %T(%v) - int value\n", is.Type[int](value2), is.Type[int](value2))
	fmt.Printf("is.Type[string](value2): %T(%v) - int value checked as string\n", is.Type[string](value2), is.Type[string](value2))

	fmt.Printf("is.Type[bool](value3): %T(%v) - bool value\n", is.Type[bool](value3), is.Type[bool](value3))

	// Output:
	// Type checking:
	// is.Type[string](value1): bool(true) - string value
	// is.Type[int](value1): bool(false) - string value checked as int
	// is.Type[int](value2): bool(true) - int value
	// is.Type[string](value2): bool(false) - int value checked as string
	// is.Type[bool](value3): bool(true) - bool value
}

func ExampleString() {
	// Testing String type predicate
	fmt.Println("String type checking:")

	value1 := "hello"
	value2 := 42

	fmt.Printf("is.String(value1): %T(%v) - string value\n", is.String(value1), is.String(value1))
	fmt.Printf("is.String(value2): %T(%v) - int value\n", is.String(value2), is.String(value2))
	fmt.Printf("is.String(\"test\"): %T(%v) - string literal\n", is.String("test"), is.String("test"))

	// Output:
	// String type checking:
	// is.String(value1): bool(true) - string value
	// is.String(value2): bool(false) - int value
	// is.String("test"): bool(true) - string literal
}

func ExampleInt() {
	// Testing Int type predicate
	fmt.Println("Int type checking:")

	value1 := 42
	value2 := "hello"

	fmt.Printf("is.Int(value1): %T(%v) - int value\n", is.Int(value1), is.Int(value1))
	fmt.Printf("is.Int(value2): %T(%v) - string value\n", is.Int(value2), is.Int(value2))
	fmt.Printf("is.Int(100): %T(%v) - int literal\n", is.Int(100), is.Int(100))

	// Output:
	// Int type checking:
	// is.Int(value1): bool(true) - int value
	// is.Int(value2): bool(false) - string value
	// is.Int(100): bool(true) - int literal
}

func ExampleBool() {
	// Testing Bool type predicate
	fmt.Println("Bool type checking:")

	value1 := true
	value2 := 1
	value3 := "true"

	fmt.Printf("is.Bool(value1): %T(%v) - bool value\n", is.Bool(value1), is.Bool(value1))
	fmt.Printf("is.Bool(value2): %T(%v) - int value\n", is.Bool(value2), is.Bool(value2))
	fmt.Printf("is.Bool(value3): %T(%v) - string value\n", is.Bool(value3), is.Bool(value3))
	fmt.Printf("is.Bool(false): %T(%v) - bool literal\n", is.Bool(false), is.Bool(false))

	// Output:
	// Bool type checking:
	// is.Bool(value1): bool(true) - bool value
	// is.Bool(value2): bool(false) - int value
	// is.Bool(value3): bool(false) - string value
	// is.Bool(false): bool(true) - bool literal
}

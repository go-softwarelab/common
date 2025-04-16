package optional_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/optional"
	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seqerr"
)

func ExampleEmpty() {
	opt := optional.Empty[string]()
	fmt.Println("Is empty:", opt.IsEmpty())
	fmt.Println("Is present:", opt.IsPresent())

	// Output:
	// Is empty: true
	// Is present: false
}

func ExampleOf() {
	// With a value
	opt := optional.Of("hello")
	fmt.Println("Value:", opt.MustGet())
	fmt.Println("Is empty:", opt.IsEmpty())

	// With a nil pointer
	var ptr *string = nil
	optPtr := optional.Of(ptr)
	fmt.Println("Nil pointer optional is empty:", optPtr.IsEmpty())

	// Output:
	// Value: hello
	// Is empty: false
	// Nil pointer optional is empty: true
}

func ExampleOfPtr() {
	// With a value pointer
	value := "hello"
	opt := optional.OfPtr(&value)
	fmt.Println("Value:", opt.MustGet())

	// With a nil pointer
	var nilPtr *string
	optNil := optional.OfPtr(nilPtr)
	fmt.Println("Is empty:", optNil.IsEmpty())

	// Output:
	// Value: hello
	// Is empty: true
}

func ExampleOfValue() {
	// Non-zero value
	opt := optional.OfValue(42)
	fmt.Println("Value present:", opt.IsPresent())
	fmt.Println("Value:", opt.MustGet())

	// Zero value
	optZero := optional.OfValue(0)
	fmt.Println("Zero value present:", optZero.IsPresent())

	// Output:
	// Value present: true
	// Value: 42
	// Zero value present: false
}

func ExampleElem_Or() {
	opt1 := optional.Of("first")
	opt2 := optional.Of("second")
	empty := optional.Empty[string]()

	// Present optional takes precedence
	fmt.Println("First or second:", opt1.Or(opt2).MustGet())
	fmt.Println("Empty or second:", empty.Or(opt2).MustGet())

	// Output:
	// First or second: first
	// Empty or second: second
}

func ExampleElem_ShouldGet() {
	opt := optional.Of(42)
	empty := optional.Empty[int]()

	val1, err1 := opt.ShouldGet()
	fmt.Println("Value:", val1)
	fmt.Println("Error:", err1)

	val2, err2 := empty.ShouldGet()
	fmt.Println("Empty value:", val2)
	fmt.Println("Empty error:", err2)

	// Output:
	// Value: 42
	// Error: <nil>
	// Empty value: 0
	// Empty error: value is not present
}

func ExampleElem_MustGet() {
	opt := optional.Of("hello")
	fmt.Println("Value:", opt.MustGet())

	// Note: Using MustGet on empty optional would panic
	// empty := optional.Empty[string]()
	// empty.MustGet() // would panic with "value is not present"

	// Output:
	// Value: hello
}

func ExampleElem_MustGetf() {
	opt := optional.Of("hello")
	fmt.Println("Value:", opt.MustGetf("Custom error: %s", "not found"))

	// Note: Using MustGetf on empty optional would panic with custom message
	// empty := optional.Empty[string]()
	// empty.MustGetf("Custom error: %s", "not found") // would panic with "Custom error: not found"

	// Output:
	// Value: hello
}

func ExampleElem_OrElseGet() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	counter := 0
	getDefault := func() string {
		counter++
		return fmt.Sprintf("default-%d", counter)
	}

	fmt.Println("Present value:", opt.OrElseGet(getDefault))
	fmt.Println("Empty value:", empty.OrElseGet(getDefault))
	fmt.Println("Empty value again:", empty.OrElseGet(getDefault))

	// Output:
	// Present value: hello
	// Empty value: default-1
	// Empty value again: default-2
}

func ExampleElem_OrErrorGet() {
	opt := optional.Of(42)
	empty := optional.Empty[int]()

	errCounter := 0
	getError := func() error {
		errCounter++
		return fmt.Errorf("not found-%d", errCounter)
	}

	val1, err1 := opt.OrErrorGet(getError)
	fmt.Println("Value:", val1)
	fmt.Println("Error:", err1)

	val2, err2 := empty.OrErrorGet(getError)
	fmt.Println("Empty value:", val2)
	fmt.Println("Empty error:", err2)

	// Output:
	// Value: 42
	// Error: <nil>
	// Empty value: 0
	// Empty error: not found-1
}

func ExampleElem_OrElse() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	fmt.Println("Present value:", opt.OrElse("default"))
	fmt.Println("Empty value:", empty.OrElse("default"))

	// Output:
	// Present value: hello
	// Empty value: default
}

func ExampleElem_OrError() {
	opt := optional.Of(42)
	empty := optional.Empty[int]()

	val1, err1 := opt.OrError(fmt.Errorf("not found"))
	fmt.Println("Value:", val1)
	fmt.Println("Error:", err1)

	val2, err2 := empty.OrError(fmt.Errorf("not found"))
	fmt.Println("Empty value:", val2)
	fmt.Println("Empty error:", err2 != nil)

	// Output:
	// Value: 42
	// Error: <nil>
	// Empty value: 0
	// Empty error: true
}

func ExampleElem_IfPresent() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	opt.IfPresent(func(value string) {
		fmt.Println("Value is present:", value)
	})

	empty.IfPresent(func(value string) {
		fmt.Println("This won't be printed")
	})

	// Output:
	// Value is present: hello
}

func ExampleElem_IfNotPresent() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	opt.IfNotPresent(func() {
		fmt.Println("This won't be printed")
	})

	empty.IfNotPresent(func() {
		fmt.Println("This executes when empty")
	})

	// Output:
	// This executes when empty
}

func ExampleElem_ToSeq() {
	opt := optional.Of("hello")

	var values []string
	seq.ForEach(opt.ToSeq(), func(value string) {
		values = append(values, value)
	})

	fmt.Println("Values:", values)

	empty := optional.Empty[string]()
	var emptyValues []string
	seq.ForEach(empty.ToSeq(), func(value string) {
		emptyValues = append(emptyValues, value)
	})

	fmt.Println("Empty values length:", len(emptyValues))

	// Output:
	// Values: [hello]
	// Empty values length: 0
}

func ExampleElem_ToSeq2() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	err := seqerr.ForEach(opt.ToSeq2(), func(value string) {
		fmt.Printf("Value: %s\n", value)
	})
	if err != nil {
		panic(err)
	}

	// With empty value
	err = seqerr.ForEach(empty.ToSeq2(), func(value string) {
		fmt.Printf("Unexpected value: %s\n", value)
	})
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
	// Output:
	// Value: hello
	// Expected error: value is not present
}

func ExampleElem_IsNotEmpty() {
	opt := optional.Of("hello")
	empty := optional.Empty[string]()

	fmt.Println("First is not empty:", opt.IsNotEmpty())
	fmt.Println("Second is not empty:", empty.IsNotEmpty())

	// Output:
	// First is not empty: true
	// Second is not empty: false
}

func ExampleElem_OrZeroValue() {
	opt := optional.Of(42)
	empty := optional.Empty[int]()

	fmt.Println("Present value:", opt.OrZeroValue())
	fmt.Println("Empty value:", empty.OrZeroValue())

	// Output:
	// Present value: 42
	// Empty value: 0
}

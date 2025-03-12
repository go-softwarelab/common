package to_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func ExampleBoolFromString() {
	// Converting various string representations to bool
	trueValue, err := to.BoolFromString("true")
	fmt.Println("true ->", trueValue, err)

	falseValue, err := to.BoolFromString("false")
	fmt.Println("false ->", falseValue, err)

	oneValue, err := to.BoolFromString("1")
	fmt.Println("1 ->", oneValue, err)

	zeroValue, err := to.BoolFromString("0")
	fmt.Println("0 ->", zeroValue, err)

	invalidValue, err := to.BoolFromString("not-a-bool")
	fmt.Println("not-a-bool ->", invalidValue, err != nil)

	// Output:
	// true -> true <nil>
	// false -> false <nil>
	// 1 -> true <nil>
	// 0 -> false <nil>
	// not-a-bool -> false true
}

func ExampleBoolFromNumber() {
	// Converting various numeric values to bool
	fmt.Println("0 ->", to.BoolFromNumber(0))
	fmt.Println("1 ->", to.BoolFromNumber(1))
	fmt.Println("-1 ->", to.BoolFromNumber(-1))
	fmt.Println("42 ->", to.BoolFromNumber(42))
	fmt.Println("0.0 ->", to.BoolFromNumber(0.0))
	fmt.Println("0.1 ->", to.BoolFromNumber(0.1))

	// Output:
	// 0 -> false
	// 1 -> true
	// -1 -> true
	// 42 -> true
	// 0.0 -> false
	// 0.1 -> true
}

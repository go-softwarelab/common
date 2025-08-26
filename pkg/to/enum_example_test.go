package to_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func ExampleEnum() {
	type FooBarEnum string
	const (
		Foo FooBarEnum = "foo"
		Bar FooBarEnum = "bar"
	)

	// Case-insensitive match succeeds
	v, err := to.Enum("FoO", Foo, Bar)
	fmt.Printf("%q, Error: %v\n", v, err)

	// No match found -> returns empty string and an error
	v, err = to.Enum("baz", Foo, Bar)
	fmt.Printf("%q, Error: %v\n", v, err)

	// Output:
	// "foo", Error: <nil>
	// "", Error: invalid value: baz doesn't match enum values [foo bar]
}

func ExampleEnumStrict() {
	type FooBarEnum string
	const (
		Foo FooBarEnum = "foo"
		Bar FooBarEnum = "bar"
	)

	// Case-sensitive match succeeds
	v, err := to.EnumStrict("foo", Foo, Bar)
	fmt.Printf("%q, Error: %v\n", v, err)

	// Case-sensitive mismatch -> returns empty string and an error
	v, err = to.EnumStrict("FOO", Foo, Bar)
	fmt.Printf("%q, Error: %v\n", v, err)

	// Output:
	// "foo", Error: <nil>
	// "", Error: invalid value: FOO doesn't match enum values [foo bar]
}

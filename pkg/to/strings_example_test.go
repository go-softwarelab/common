package to_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func ExampleStringFromInteger() {
	// Integer to string conversion
	val := to.StringFromInteger(42)
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("42")
}

func ExampleStringFromFloat() {
	// Float to string conversion
	val := to.StringFromFloat(3.14159)
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("3.141590")
}

func ExampleStringFromBool() {
	// Boolean to string conversion
	val := to.StringFromBool(true)
	fmt.Printf("%T(%q)\n", val, val)

	val = to.StringFromBool(false)
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("true")
	// string("false")
}

func ExampleStringFromBytes() {
	// Byte slice to string conversion
	val := to.StringFromBytes([]byte("hello"))
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("hello")
}

func ExampleStringFromRune() {
	// Rune to string conversion
	val := to.StringFromRune('世')
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("世")
}

func ExampleWords() {
	// Splitting into words
	camelCase := to.Words("camelCaseString")
	fmt.Printf("%#v\n", camelCase)

	pascalCase := to.Words("PascalCaseString")
	fmt.Printf("%#v\n", pascalCase)

	withNumbers := to.Words("Int8Value")
	fmt.Printf("%#v\n", withNumbers)

	withSpaces := to.Words("  hello  world  ")
	fmt.Printf("%#v\n", withSpaces)

	// Output:
	// []string{"camel", "Case", "String"}
	// []string{"Pascal", "Case", "String"}
	// []string{"Int", "8", "Value"}
	// []string{"hello", "world"}
}

func ExampleSentences() {
	// Splitting into sentences
	text := "Hello, world! This is a test. Is this working?"
	sentences := to.Sentences(text)
	fmt.Printf("%#v\n", sentences)

	// Output:
	// []string{"Hello, world!", "This is a test.", "Is this working?"}
}

func ExamplePascalCase() {
	// Convert to PascalCase
	val := to.PascalCase("hello world")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.PascalCase("some-kebab-case")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.PascalCase("snake_case_string")
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("HelloWorld")
	// string("SomeKebabCase")
	// string("SnakeCaseString")
}

func ExampleCamelCase() {
	// Convert to camelCase
	val := to.CamelCase("hello world")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.CamelCase("some-kebab-case")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.CamelCase("PascalCaseString")
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("helloWorld")
	// string("someKebabCase")
	// string("pascalCaseString")
}

func ExampleKebabCase() {
	// Convert to kebab-case
	val := to.KebabCase("hello world")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.KebabCase("camelCaseString")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.KebabCase("PascalCaseString")
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("hello-world")
	// string("camel-case-string")
	// string("pascal-case-string")
}

func ExampleSnakeCase() {
	// Convert to snake_case
	val := to.SnakeCase("hello world")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.SnakeCase("camelCaseString")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.SnakeCase("PascalCaseString")
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("hello_world")
	// string("camel_case_string")
	// string("pascal_case_string")
}

func ExampleCapitalized() {
	// Capitalize the first letter of each sentence
	val := to.Capitalized("hello")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.Capitalized("HELLO WORLD")
	fmt.Printf("%T(%q)\n", val, val)

	val = to.Capitalized("multiple. sentences. here.")
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("Hello")
	// string("Hello world")
	// string("Multiple. Sentences. Here.")
}

func ExampleEllipsis() {
	// Truncate string with ellipsis
	val := to.Ellipsis("This is a short string", 10)
	fmt.Printf("%T(%q)\n", val, val)

	val = to.Ellipsis("Short", 10)
	fmt.Printf("%T(%q)\n", val, val)

	val = to.Ellipsis("Little bit longer one, but with small length", 3)
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("This is...")
	// string("Short")
	// string("...")
}

func ExampleEllipsisWith() {
	// Create a function that truncates strings to 10 characters
	truncate := to.EllipsisWith(10)

	val := truncate("This is a long string that should be truncated")
	fmt.Printf("%T(%q)\n", val, val)

	val = truncate("Short")
	fmt.Printf("%T(%q)\n", val, val)

	// Output:
	// string("This is...")
	// string("Short")
}

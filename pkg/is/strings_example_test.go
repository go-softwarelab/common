package is_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func ExampleEmptyString() {
	// EmptyString checks if a string is empty
	fmt.Println("Checking for empty strings:")
	fmt.Printf("is.EmptyString(\"\"): %T(%v) - empty string\n", is.EmptyString(""), is.EmptyString(""))
	fmt.Printf("is.EmptyString(\"hello\"): %T(%v) - non-empty string\n", is.EmptyString("hello"), is.EmptyString("hello"))

	// Output:
	// Checking for empty strings:
	// is.EmptyString(""): bool(true) - empty string
	// is.EmptyString("hello"): bool(false) - non-empty string
}

func ExampleBlankString() {
	// BlankString checks if a string is empty or contains only whitespace
	fmt.Println("Checking for blank strings:")
	fmt.Printf("is.BlankString(\"\"): %T(%v) - empty string\n", is.BlankString(""), is.BlankString(""))
	fmt.Printf("is.BlankString(\"   \"): %T(%v) - whitespace only\n", is.BlankString("   "), is.BlankString("   "))
	fmt.Printf("is.BlankString(\"\\t\\n\"): %T(%v) - tabs and newlines\n", is.BlankString("\t\n"), is.BlankString("\t\n"))
	fmt.Printf("is.BlankString(\"hello\"): %T(%v) - non-blank string\n", is.BlankString("hello"), is.BlankString("hello"))

	// Output:
	// Checking for blank strings:
	// is.BlankString(""): bool(true) - empty string
	// is.BlankString("   "): bool(true) - whitespace only
	// is.BlankString("\t\n"): bool(true) - tabs and newlines
	// is.BlankString("hello"): bool(false) - non-blank string
}

func ExampleNotEmptyString() {
	// NotEmptyString checks if a string is not empty
	fmt.Println("Checking for non-empty strings:")
	fmt.Printf("is.NotEmptyString(\"hello\"): %T(%v) - non-empty string\n", is.NotEmptyString("hello"), is.NotEmptyString("hello"))
	fmt.Printf("is.NotEmptyString(\"   \"): %T(%v) - whitespace string\n", is.NotEmptyString("   "), is.NotEmptyString("   "))
	fmt.Printf("is.NotEmptyString(\"\"): %T(%v) - empty string\n", is.NotEmptyString(""), is.NotEmptyString(""))

	// Output:
	// Checking for non-empty strings:
	// is.NotEmptyString("hello"): bool(true) - non-empty string
	// is.NotEmptyString("   "): bool(true) - whitespace string
	// is.NotEmptyString(""): bool(false) - empty string
}

func ExampleNotBlankString() {
	// NotBlankString checks if a string has non-whitespace content
	fmt.Println("Checking for non-blank strings:")
	fmt.Printf("is.NotBlankString(\"hello\"): %T(%v) - string with content\n", is.NotBlankString("hello"), is.NotBlankString("hello"))
	fmt.Printf("is.NotBlankString(\" x \"): %T(%v) - string with whitespace and content\n", is.NotBlankString(" x "), is.NotBlankString(" x "))
	fmt.Printf("is.NotBlankString(\"   \"): %T(%v) - whitespace only\n", is.NotBlankString("   "), is.NotBlankString("   "))
	fmt.Printf("is.NotBlankString(\"\"): %T(%v) - empty string\n", is.NotBlankString(""), is.NotBlankString(""))

	// Output:
	// Checking for non-blank strings:
	// is.NotBlankString("hello"): bool(true) - string with content
	// is.NotBlankString(" x "): bool(true) - string with whitespace and content
	// is.NotBlankString("   "): bool(false) - whitespace only
	// is.NotBlankString(""): bool(false) - empty string
}

package is

import "strings"

// EmptyString returns true if the string is empty
func EmptyString(s string) bool {
	return s == ""
}

// BlankString returns true if the string is empty or contains only whitespace
func BlankString(s string) bool {
	return s == "" || strings.TrimSpace(s) == ""
}

// NotEmptyString returns true if the string is not empty
func NotEmptyString(s string) bool {
	return !EmptyString(s)
}

// NotBlankString returns true if the string is not empty or contains only whitespace
func NotBlankString(s string) bool {
	return !BlankString(s)
}

package is

import "strings"

// EmptyString returns true if the string is empty
func EmptyString[S string | *string](s S) (isEmpty bool) {
	switch str := (any)(s).(type) {
	case string:
		isEmpty = str == ""
	case *string:
		isEmpty = str == nil || *str == ""
	}
	return isEmpty
}

// BlankString returns true if the string is empty or contains only whitespace
func BlankString[S string | *string](str S) (isBlank bool) {
	switch s := (any)(str).(type) {
	case string:
		isBlank = s == "" || strings.TrimSpace(s) == ""
	case *string:
		isBlank = s == nil || *s == "" || strings.TrimSpace(*s) == ""
	}
	return isBlank
}

// NotEmptyString returns true if the string is not empty
func NotEmptyString[S string | *string](s S) bool {
	return !EmptyString(s)
}

// NotBlankString returns true if the string is not empty or contains only whitespace
func NotBlankString[S string | *string](s S) bool {
	return !BlankString(s)
}

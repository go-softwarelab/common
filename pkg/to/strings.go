package to

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-softwarelab/common/pkg/types"
)

var (
	// bearer:disable go_lang_permissive_regex_validation
	splitWordReg = regexp.MustCompile(`([a-z])([A-Z0-9])|([a-zA-Z])([0-9])|([0-9])([a-zA-Z])|([A-Z])([A-Z])([a-z])`)
	// bearer:disable go_lang_permissive_regex_validation
	splitNumberLetterReg = regexp.MustCompile(`([0-9])([a-zA-Z])`)
)

// StringFromInteger will convert any integer to string
func StringFromInteger[V types.Integer](value V) string {
	return fmt.Sprintf("%d", value)
}

// StringFromFloat will convert any float to string
func StringFromFloat[V types.Float](value V) string {
	return fmt.Sprintf("%f", value)
}

// StringFromBool will convert any bool to string
func StringFromBool(value bool) string {
	return fmt.Sprintf("%v", value)
}

// StringFromBytes will convert any byte slice to string
func StringFromBytes(value []byte) string {
	return string(value)
}

// StringFromRune will convert any rune to string
func StringFromRune(value rune) string {
	return string(value)
}

// Words splits string into an array of its words.
func Words(str string) []string {
	str = splitWordReg.ReplaceAllString(str, `$1$3$5$7 $2$4$6$8$9`)
	// example: Int8Value => Int 8Value => Int 8 Value
	str = splitNumberLetterReg.ReplaceAllString(str, "$1 $2")
	var result strings.Builder
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(r)
		} else {
			result.WriteRune(' ')
		}
	}
	return strings.Fields(result.String())
}

// Sentences splits string into an array of its sentences.
// The sentences are trimmed from leading and trailing spaces.
func Sentences(str string) []string {
	result := make([]string, 0)
	builder := strings.Builder{}
	for _, r := range str {
		if r == '.' || r == '!' || r == '?' {
			builder.WriteRune(r)
			result = append(result, strings.TrimSpace(builder.String()))
			builder.Reset()
			continue
		}
		builder.WriteRune(r)
	}
	if builder.Len() > 0 {
		result = append(result, strings.TrimSpace(builder.String()))
	}
	return result
}

// PascalCase converts string to pascal case.
func PascalCase(str string) string {
	items := Words(str)
	for i := range items {
		items[i] = Capitalized(items[i])
	}
	return strings.Join(items, "")
}

// CamelCase converts string to camel case.
func CamelCase(str string) string {
	items := Words(str)
	for i, item := range items {
		item = strings.ToLower(item)
		if i > 0 {
			item = Capitalized(item)
		}
		items[i] = item
	}
	return strings.Join(items, "")
}

// KebabCase converts string to kebab case.
func KebabCase(str string) string {
	items := Words(str)
	for i := range items {
		items[i] = strings.ToLower(items[i])
	}
	return strings.Join(items, "-")
}

// SnakeCase converts string to snake case.
func SnakeCase(str string) string {
	items := Words(str)
	for i := range items {
		items[i] = strings.ToLower(items[i])
	}
	return strings.Join(items, "_")
}

// Capitalized converts the first character of string to upper case and the remaining to lower case.
func Capitalized(str string) string {
	sentences := Sentences(str)

	var builder strings.Builder
	for _, sentence := range sentences {
		if builder.Len() > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(strings.ToUpper(sentence[:1]))
		if len(sentence) > 1 {
			builder.WriteString(strings.ToLower(sentence[1:]))
		}

	}
	return builder.String()
}

// Ellipsis trims and truncates a string to a specified length and appends an ellipsis if truncated.
func Ellipsis(str string, length int) string {
	str = strings.TrimSpace(str)

	if len(str) > length {
		if len(str) < 3 || length < 3 {
			return "..."
		}
		return strings.TrimSpace(str[0:length-3]) + "..."
	}

	return str
}

// EllipsisWith returns a function that trims and truncates a string to a specified length and appends an ellipsis if truncated.
// It's wrapped around Ellipsis function, to make it usable in Map functions.
// #mapper
func EllipsisWith(length int) func(str string) string {
	return func(str string) string {
		return Ellipsis(str, length)
	}
}

package to_test

import (
	"reflect"
	"testing"

	"github.com/go-softwarelab/common/pkg/to"
)

func TestSentences(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single sentence with period",
			input:    "This is a test.",
			expected: []string{"This is a test."},
		},
		{
			name:     "single sentence with exclamation mark",
			input:    "Hello world!",
			expected: []string{"Hello world!"},
		},
		{
			name:     "single sentence with question mark",
			input:    "How are you?",
			expected: []string{"How are you?"},
		},
		{
			name:     "multiple sentences with mixed punctuation",
			input:    "Hello. How are you? I'm fine!",
			expected: []string{"Hello.", "How are you?", "I'm fine!"},
		},
		{
			name:     "sentence without ending punctuation",
			input:    "This sentence has no ending punctuation",
			expected: []string{"This sentence has no ending punctuation"},
		},
		{
			name:     "sentence with numbers and special characters",
			input:    "Version 1.0 is released. Download now!",
			expected: []string{"Version 1.", "0 is released.", "Download now!"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := to.Sentences(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Sentences() = %#v, want %#v", got, tt.expected)
			}
		})
	}
}

func TestCapitalized(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single word lowercase",
			input:    "hello",
			expected: "Hello",
		},
		{
			name:     "single word uppercase",
			input:    "HELLO",
			expected: "Hello",
		},
		{
			name:     "single word mixed case",
			input:    "hElLo",
			expected: "Hello",
		},
		{
			name:     "multiple words",
			input:    "hello world",
			expected: "Hello world",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "A",
		},
		{
			name:     "sentence with period",
			input:    "hello world.",
			expected: "Hello world.",
		},
		{
			name:     "multiple sentences",
			input:    "hello. world",
			expected: "Hello. World",
		},
		{
			name:     "sentence starting with number",
			input:    "123 test",
			expected: "123 test",
		},
		{
			name:     "mixed punctuation",
			input:    "hello! how? yes.",
			expected: "Hello! How? Yes.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := to.Capitalized(tt.input)
			if got != tt.expected {
				t.Errorf("Capitalized() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestEllipsis(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		length   int
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			length:   10,
			expected: "",
		},
		{
			name:     "string with spaces only",
			input:    "   ",
			length:   10,
			expected: "",
		},
		{
			name:     "short string",
			input:    "Hello",
			length:   10,
			expected: "Hello",
		},
		{
			name:     "string with exact length",
			input:    "Exact ten.",
			length:   10,
			expected: "Exact ten.",
		},
		{
			name:     "long string needs truncation",
			input:    "This string is too long and will be truncated.",
			length:   20,
			expected: "This string is to...",
		},
		{
			name:     "string with leading/trailing spaces",
			input:    "  Trim spaces  ",
			length:   15,
			expected: "Trim spaces",
		},
		{
			name:     "length less than 3",
			input:    "Hello world",
			length:   2,
			expected: "...",
		},
		{
			name:     "very short string with length less than 3",
			input:    "Hi",
			length:   1,
			expected: "...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := to.Ellipsis(tt.input, tt.length)
			if got != tt.expected {
				t.Errorf("Ellipsis(%q, %d) = %q, want %q",
					tt.input, tt.length, got, tt.expected)
			}
		})
	}
}

func TestPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty string", input: "", expected: ""},
		{name: "punctuation only", input: ".", expected: ""},
		{name: "sentence with punctuation", input: "Hello world!", expected: "HelloWorld"},
		{name: "uppercase single char", input: "A", expected: "A"},
		{name: "lowercase single char", input: "a", expected: "A"},
		{name: "simple word", input: "foo", expected: "Foo"},
		{name: "snake_case", input: "snake_case", expected: "SnakeCase"},
		{name: "UPPER_SNAKE_CASE", input: "SNAKE_CASE", expected: "SnakeCase"},
		{name: "kebab-case", input: "kebab-case", expected: "KebabCase"},
		{name: "PascalCase", input: "PascalCase", expected: "PascalCase"},
		{name: "camelCase", input: "camelCase", expected: "CamelCase"},
		{name: "Title Case", input: "Title Case", expected: "TitleCase"},
		{name: "dot.notation", input: "point.case", expected: "PointCase"},
		{name: "leading spaces", input: " leading space", expected: "LeadingSpace"},
		{name: "trailing spaces", input: "trailing space ", expected: "TrailingSpace"},
		{name: "uppercase breaks", input: "CASEBreak", expected: "CaseBreak"},
		{name: "initialisms", input: "HTTPStatusCode", expected: "HttpStatusCode"},
		{name: "numbers with letters", input: "http200", expected: "Http200"},
		{name: "mixed numbers", input: "Int8Value", expected: "Int8Value"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := to.PascalCase(tt.input)
			if result != tt.expected {
				t.Errorf("PascalCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty string", input: "", expected: ""},
		{name: "punctuation only", input: ".", expected: ""},
		{name: "sentence with punctuation", input: "Hello world!", expected: "helloWorld"},
		{name: "uppercase single char", input: "A", expected: "a"},
		{name: "lowercase single char", input: "a", expected: "a"},
		{name: "simple word", input: "foo", expected: "foo"},
		{name: "snake_case", input: "snake_case", expected: "snakeCase"},
		{name: "UPPER_SNAKE_CASE", input: "SNAKE_CASE", expected: "snakeCase"},
		{name: "kebab-case", input: "kebab-case", expected: "kebabCase"},
		{name: "PascalCase", input: "PascalCase", expected: "pascalCase"},
		{name: "camelCase", input: "camelCase", expected: "camelCase"},
		{name: "Title Case", input: "Title Case", expected: "titleCase"},
		{name: "dot.notation", input: "point.case", expected: "pointCase"},
		{name: "leading spaces", input: " leading space", expected: "leadingSpace"},
		{name: "trailing spaces", input: "trailing space ", expected: "trailingSpace"},
		{name: "uppercase breaks", input: "CASEBreak", expected: "caseBreak"},
		{name: "initialisms", input: "HTTPStatusCode", expected: "httpStatusCode"},
		{name: "numbers with letters", input: "http200", expected: "http200"},
		{name: "mixed numbers", input: "Int8Value", expected: "int8Value"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := to.CamelCase(tt.input)
			if result != tt.expected {
				t.Errorf("CamelCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestKebabCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty string", input: "", expected: ""},
		{name: "punctuation only", input: ".", expected: ""},
		{name: "sentence with punctuation", input: "Hello world!", expected: "hello-world"},
		{name: "uppercase single char", input: "A", expected: "a"},
		{name: "lowercase single char", input: "a", expected: "a"},
		{name: "simple word", input: "foo", expected: "foo"},
		{name: "snake_case", input: "snake_case", expected: "snake-case"},
		{name: "UPPER_SNAKE_CASE", input: "SNAKE_CASE", expected: "snake-case"},
		{name: "kebab-case", input: "kebab-case", expected: "kebab-case"},
		{name: "PascalCase", input: "PascalCase", expected: "pascal-case"},
		{name: "camelCase", input: "camelCase", expected: "camel-case"},
		{name: "Title Case", input: "Title Case", expected: "title-case"},
		{name: "dot.notation", input: "point.case", expected: "point-case"},
		{name: "leading spaces", input: " leading space", expected: "leading-space"},
		{name: "trailing spaces", input: "trailing space ", expected: "trailing-space"},
		{name: "uppercase breaks", input: "CASEBreak", expected: "case-break"},
		{name: "initialisms", input: "HTTPStatusCode", expected: "http-status-code"},
		{name: "numbers with letters", input: "http200", expected: "http-200"},
		{name: "mixed numbers", input: "Int8Value", expected: "int-8-value"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := to.KebabCase(tt.input)
			if result != tt.expected {
				t.Errorf("KebabCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty string", input: "", expected: ""},
		{name: "punctuation only", input: ".", expected: ""},
		{name: "sentence with punctuation", input: "Hello world!", expected: "hello_world"},
		{name: "uppercase single char", input: "A", expected: "a"},
		{name: "lowercase single char", input: "a", expected: "a"},
		{name: "simple word", input: "foo", expected: "foo"},
		{name: "snake_case", input: "snake_case", expected: "snake_case"},
		{name: "UPPER_SNAKE_CASE", input: "SNAKE_CASE", expected: "snake_case"},
		{name: "kebab-case", input: "kebab-case", expected: "kebab_case"},
		{name: "PascalCase", input: "PascalCase", expected: "pascal_case"},
		{name: "camelCase", input: "camelCase", expected: "camel_case"},
		{name: "Title Case", input: "Title Case", expected: "title_case"},
		{name: "dot.notation", input: "point.case", expected: "point_case"},
		{name: "leading spaces", input: " leading space", expected: "leading_space"},
		{name: "trailing spaces", input: "trailing space ", expected: "trailing_space"},
		{name: "uppercase breaks", input: "CASEBreak", expected: "case_break"},
		{name: "initialisms", input: "HTTPStatusCode", expected: "http_status_code"},
		{name: "numbers with letters", input: "http200", expected: "http_200"},
		{name: "mixed numbers", input: "Int8Value", expected: "int_8_value"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := to.SnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("SnakeCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

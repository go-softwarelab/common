package is_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/is"
)

func ExampleZero() {
	// Zero detects zero values of different types
	fmt.Println("Checking zero values:")
	fmt.Printf("is.Zero(0): %v (integer zero value)\n", is.Zero(0))
	fmt.Printf("is.Zero(42): %v (non-zero integer)\n", is.Zero(42))
	fmt.Printf("is.Zero(\"\"): %v (empty string)\n", is.Zero(""))
	fmt.Printf("is.Zero(\"hello\"): %v (non-empty string)\n", is.Zero("hello"))

	// Output:
	// Checking zero values:
	// is.Zero(0): true (integer zero value)
	// is.Zero(42): false (non-zero integer)
	// is.Zero(""): true (empty string)
	// is.Zero("hello"): false (non-empty string)
}

func ExampleEqual() {
	// Equal checks equality between two values
	fmt.Println("Checking equality:")
	fmt.Printf("is.Equal(42, 42): %v (same integers)\n", is.Equal(42, 42))
	fmt.Printf("is.Equal(42, 43): %v (different integers)\n", is.Equal(42, 43))
	fmt.Printf("is.Equal(\"hello\", \"hello\"): %v (same strings)\n", is.Equal("hello", "hello"))
	fmt.Printf("is.Equal(\"hello\", \"world\"): %v (different strings)\n", is.Equal("hello", "world"))

	// Output:
	// Checking equality:
	// is.Equal(42, 42): true (same integers)
	// is.Equal(42, 43): false (different integers)
	// is.Equal("hello", "hello"): true (same strings)
	// is.Equal("hello", "world"): false (different strings)
}

func ExampleNotEqual() {
	// NotEqual checks inequality between two values
	fmt.Println("Checking inequality:")
	fmt.Printf("is.NotEqual(42, 43): %v (different integers)\n", is.NotEqual(42, 43))
	fmt.Printf("is.NotEqual(42, 42): %v (same integers)\n", is.NotEqual(42, 42))
	fmt.Printf("is.NotEqual(\"hello\", \"world\"): %v (different strings)\n", is.NotEqual("hello", "world"))
	fmt.Printf("is.NotEqual(\"hello\", \"hello\"): %v (same strings)\n", is.NotEqual("hello", "hello"))

	// Output:
	// Checking inequality:
	// is.NotEqual(42, 43): true (different integers)
	// is.NotEqual(42, 42): false (same integers)
	// is.NotEqual("hello", "world"): true (different strings)
	// is.NotEqual("hello", "hello"): false (same strings)
}

func ExampleEqualTo() {
	// EqualTo creates a function that checks for equality with a specific value
	fmt.Println("Using equality checker functions:")

	// Create a function that checks if a value equals 42
	isFortyTwo := is.EqualTo(42)
	fmt.Printf("isFortyTwo(42): %v (matching value)\n", isFortyTwo(42))
	fmt.Printf("isFortyTwo(43): %v (non-matching value)\n", isFortyTwo(43))

	// Create a function that checks if a string equals "admin"
	isAdmin := is.EqualTo("admin")
	fmt.Printf("isAdmin(\"admin\"): %v (matching role)\n", isAdmin("admin"))
	fmt.Printf("isAdmin(\"user\"): %v (non-matching role)\n", isAdmin("user"))

	// Output:
	// Using equality checker functions:
	// isFortyTwo(42): true (matching value)
	// isFortyTwo(43): false (non-matching value)
	// isAdmin("admin"): true (matching role)
	// isAdmin("user"): false (non-matching role)
}

func ExampleGreater() {
	// Greater checks if first value is greater than second
	fmt.Println("Checking greater than relation:")
	fmt.Printf("is.Greater(42, 30): %v (42 > 30)\n", is.Greater(42, 30))
	fmt.Printf("is.Greater(30, 42): %v (30 > 42)\n", is.Greater(30, 42))
	fmt.Printf("is.Greater(42, 42): %v (42 > 42)\n", is.Greater(42, 42))

	// String comparison is lexicographic
	fmt.Printf("is.Greater(\"zebra\", \"apple\"): %v (alphabetical order)\n", is.Greater("zebra", "apple"))
	fmt.Printf("is.Greater(\"apple\", \"zebra\"): %v (alphabetical order)\n", is.Greater("apple", "zebra"))

	// Output:
	// Checking greater than relation:
	// is.Greater(42, 30): true (42 > 30)
	// is.Greater(30, 42): false (30 > 42)
	// is.Greater(42, 42): false (42 > 42)
	// is.Greater("zebra", "apple"): true (alphabetical order)
	// is.Greater("apple", "zebra"): false (alphabetical order)
}

func ExampleGreaterThen() {
	// GreaterThen creates a function that checks if values exceed a threshold
	fmt.Println("Using threshold checker functions:")

	// Check if age is considered adult (over 18)
	isAdult := is.GreaterThen(18)
	fmt.Printf("isAdult(21): %v (21 exceeds adult threshold)\n", isAdult(21))
	fmt.Printf("isAdult(18): %v (18 equals adult threshold)\n", isAdult(18))
	fmt.Printf("isAdult(16): %v (16 below adult threshold)\n", isAdult(16))

	// Check if temperature is hot (over 30°C)
	isHot := is.GreaterThen(30.0)
	fmt.Printf("isHot(35.5): %v (35.5°C exceeds hot threshold)\n", isHot(35.5))
	fmt.Printf("isHot(25.0): %v (25.0°C below hot threshold)\n", isHot(25.0))

	// Output:
	// Using threshold checker functions:
	// isAdult(21): true (21 exceeds adult threshold)
	// isAdult(18): false (18 equals adult threshold)
	// isAdult(16): false (16 below adult threshold)
	// isHot(35.5): true (35.5°C exceeds hot threshold)
	// isHot(25.0): false (25.0°C below hot threshold)
}

func ExampleGreaterOrEqual() {
	// GreaterOrEqual checks if first value is greater than or equal to second
	fmt.Println("Checking greater than or equal relation:")
	fmt.Printf("is.GreaterOrEqual(42, 30): %v (42 ≥ 30)\n", is.GreaterOrEqual(42, 30))
	fmt.Printf("is.GreaterOrEqual(42, 42): %v (42 ≥ 42)\n", is.GreaterOrEqual(42, 42))
	fmt.Printf("is.GreaterOrEqual(30, 42): %v (30 ≥ 42)\n", is.GreaterOrEqual(30, 42))

	// Output:
	// Checking greater than or equal relation:
	// is.GreaterOrEqual(42, 30): true (42 ≥ 30)
	// is.GreaterOrEqual(42, 42): true (42 ≥ 42)
	// is.GreaterOrEqual(30, 42): false (30 ≥ 42)
}

func ExampleGreaterOrEqualTo() {
	// GreaterOrEqualTo creates a function checking if values meet or exceed a threshold
	fmt.Println("Using minimum threshold functions:")

	// Check if age qualifies for voting (18 or older)
	canVote := is.GreaterOrEqualTo(18)
	fmt.Printf("canVote(21): %v (21 meets voting age)\n", canVote(21))
	fmt.Printf("canVote(18): %v (18 meets voting age)\n", canVote(18))
	fmt.Printf("canVote(16): %v (16 below voting age)\n", canVote(16))

	// Check if score is passing grade (60 or higher)
	isPassing := is.GreaterOrEqualTo(60.0)
	fmt.Printf("isPassing(75.5): %v (75.5 is passing)\n", isPassing(75.5))
	fmt.Printf("isPassing(60.0): %v (60.0 is passing)\n", isPassing(60.0))
	fmt.Printf("isPassing(59.9): %v (59.9 is failing)\n", isPassing(59.9))

	// Output:
	// Using minimum threshold functions:
	// canVote(21): true (21 meets voting age)
	// canVote(18): true (18 meets voting age)
	// canVote(16): false (16 below voting age)
	// isPassing(75.5): true (75.5 is passing)
	// isPassing(60.0): true (60.0 is passing)
	// isPassing(59.9): false (59.9 is failing)
}

func ExampleLess() {
	// Less checks if first value is less than second
	fmt.Println("Checking less than relation:")
	fmt.Printf("is.Less(30, 42): %v (30 < 42)\n", is.Less(30, 42))
	fmt.Printf("is.Less(42, 30): %v (42 < 30)\n", is.Less(42, 30))
	fmt.Printf("is.Less(42, 42): %v (42 < 42)\n", is.Less(42, 42))

	// String comparison is lexicographic
	fmt.Printf("is.Less(\"apple\", \"zebra\"): %v (alphabetical order)\n", is.Less("apple", "zebra"))

	// Output:
	// Checking less than relation:
	// is.Less(30, 42): true (30 < 42)
	// is.Less(42, 30): false (42 < 30)
	// is.Less(42, 42): false (42 < 42)
	// is.Less("apple", "zebra"): true (alphabetical order)
}

func ExampleLessThen() {
	// LessThen creates a function that checks if values are below a threshold
	fmt.Println("Using maximum threshold functions:")

	// Check if someone is a minor (under 18)
	isMinor := is.LessThen(18)
	fmt.Printf("isMinor(16): %v (16 is below adult threshold)\n", isMinor(16))
	fmt.Printf("isMinor(18): %v (18 equals adult threshold)\n", isMinor(18))
	fmt.Printf("isMinor(21): %v (21 exceeds adult threshold)\n", isMinor(21))

	// Check if temperature is freezing (below 0°C)
	isFreezing := is.LessThen(0.0)
	fmt.Printf("isFreezing(-5.0): %v (-5.0°C is freezing)\n", isFreezing(-5.0))
	fmt.Printf("isFreezing(0.0): %v (0.0°C at freezing point)\n", isFreezing(0.0))

	// Output:
	// Using maximum threshold functions:
	// isMinor(16): true (16 is below adult threshold)
	// isMinor(18): false (18 equals adult threshold)
	// isMinor(21): false (21 exceeds adult threshold)
	// isFreezing(-5.0): true (-5.0°C is freezing)
	// isFreezing(0.0): false (0.0°C at freezing point)
}

func ExampleLessOrEqual() {
	// LessOrEqual checks if first value is less than or equal to second
	fmt.Println("Checking less than or equal relation:")
	fmt.Printf("is.LessOrEqual(30, 42): %v (30 ≤ 42)\n", is.LessOrEqual(30, 42))
	fmt.Printf("is.LessOrEqual(42, 42): %v (42 ≤ 42)\n", is.LessOrEqual(42, 42))
	fmt.Printf("is.LessOrEqual(50, 42): %v (50 ≤ 42)\n", is.LessOrEqual(50, 42))

	// Output:
	// Checking less than or equal relation:
	// is.LessOrEqual(30, 42): true (30 ≤ 42)
	// is.LessOrEqual(42, 42): true (42 ≤ 42)
	// is.LessOrEqual(50, 42): false (50 ≤ 42)
}

func ExampleLessOrEqualTo() {
	// LessOrEqualTo creates a function checking if values are at or below a threshold
	fmt.Println("Using maximum allowed functions:")

	// Check if child qualifies for child pricing (12 or under)
	isChildRate := is.LessOrEqualTo(12)
	fmt.Printf("isChildRate(10): %v (10 qualifies for child rate)\n", isChildRate(10))
	fmt.Printf("isChildRate(12): %v (12 qualifies for child rate)\n", isChildRate(12))
	fmt.Printf("isChildRate(13): %v (13 does not qualify for child rate)\n", isChildRate(13))

	// Check if luggage is within weight limit (23kg or less)
	isWithinWeightLimit := is.LessOrEqualTo(23.0)
	fmt.Printf("isWithinWeightLimit(20.5): %v (20.5kg is acceptable)\n", isWithinWeightLimit(20.5))
	fmt.Printf("isWithinWeightLimit(23.0): %v (23.0kg is acceptable)\n", isWithinWeightLimit(23.0))
	fmt.Printf("isWithinWeightLimit(23.5): %v (23.5kg exceeds limit)\n", isWithinWeightLimit(23.5))

	// Output:
	// Using maximum allowed functions:
	// isChildRate(10): true (10 qualifies for child rate)
	// isChildRate(12): true (12 qualifies for child rate)
	// isChildRate(13): false (13 does not qualify for child rate)
	// isWithinWeightLimit(20.5): true (20.5kg is acceptable)
	// isWithinWeightLimit(23.0): true (23.0kg is acceptable)
	// isWithinWeightLimit(23.5): false (23.5kg exceeds limit)
}

func ExampleBetween() {
	// Between checks if a value falls within a range (inclusive)
	fmt.Printf("%T(%v)\n", is.Between(25, 18, 65), is.Between(25, 18, 65))
	fmt.Printf("%T(%v)\n", is.Between(18, 18, 65), is.Between(18, 18, 65))
	fmt.Printf("%T(%v)\n", is.Between(65, 18, 65), is.Between(65, 18, 65))
	fmt.Printf("%T(%v)\n", is.Between(17, 18, 65), is.Between(17, 18, 65))
	fmt.Printf("%T(%v)\n", is.Between(66, 18, 65), is.Between(66, 18, 65))

	// Output:
	// bool(true)
	// bool(true)
	// bool(true)
	// bool(false)
	// bool(false)
}

func ExampleBetweenThe() {
	// BetweenThe creates a function checking if values are within a range
	isWorkingAge := is.BetweenThe(18, 65)

	fmt.Printf("%T(%v)\n", isWorkingAge(25), isWorkingAge(25))
	fmt.Printf("%T(%v)\n", isWorkingAge(18), isWorkingAge(18))
	fmt.Printf("%T(%v)\n", isWorkingAge(65), isWorkingAge(65))
	fmt.Printf("%T(%v)\n", isWorkingAge(16), isWorkingAge(16))
	fmt.Printf("%T(%v)\n", isWorkingAge(70), isWorkingAge(70))

	// Temperature range checker (comfortable room temperature 20-25°C)
	isComfortableTemp := is.BetweenThe(20.0, 25.0)
	fmt.Printf("%T(%v)\n", isComfortableTemp(22.5), isComfortableTemp(22.5))
	fmt.Printf("%T(%v)\n", isComfortableTemp(18.0), isComfortableTemp(18.0))

	// Output:
	// bool(true)
	// bool(true)
	// bool(true)
	// bool(false)
	// bool(false)
	// bool(true)
	// bool(false)
}

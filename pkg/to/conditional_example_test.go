package to_test

import (
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/to"
)

func ExampleIf() {
	// Using If with function providers
	result := to.If(true, func() string {
		return "condition is true"
	}).Else(func() string {
		return "condition is false"
	})
	fmt.Println(result)

	result = to.If(false, func() string {
		return "condition is true"
	}).Else(func() string {
		return "condition is false"
	})
	fmt.Println(result)

	// Output:
	// condition is true
	// condition is false
}

func ExampleIfThen() {
	// Using IfThen with direct values
	result := to.IfThen(true, "condition is true").ElseThen("condition is false")
	fmt.Println(result)

	result = to.IfThen(false, "condition is true").ElseThen("condition is false")
	fmt.Println(result)

	// Output:
	// condition is true
	// condition is false
}

func ExampleIfElseCondition_ElseIf() {
	// Using ElseIf for multiple conditions
	age := 25
	result := to.If(age < 18, func() string {
		return "minor"
	}).ElseIf(age < 65, func() string {
		return "adult"
	}).Else(func() string {
		return "senior"
	})
	fmt.Println(result)

	// Testing with different age
	age = 70
	result = to.If(age < 18, func() string {
		return "minor"
	}).ElseIf(age < 65, func() string {
		return "adult"
	}).Else(func() string {
		return "senior"
	})
	fmt.Println(result)

	// Output:
	// adult
	// senior
}

func ExampleIfElseCondition_ElseIfThen() {
	// Using ElseIfThen for multiple conditions with direct values
	score := 85
	grade := to.IfThen(score >= 90, "A").
		ElseIfThen(score >= 80, "B").
		ElseIfThen(score >= 70, "C").
		ElseIfThen(score >= 60, "D").
		ElseThen("F")
	fmt.Println(grade)

	// Output:
	// B
}

func ExampleSwitch() {
	// Using Switch with Case for equality comparisons
	dayNum := 3
	dayName := to.Switch[int, string](dayNum).
		Case(1).ThenValue("Monday").
		Case(2).ThenValue("Tuesday").
		Case(3).ThenValue("Wednesday").
		Case(4).ThenValue("Thursday").
		Case(5).ThenValue("Friday").
		Case(6).ThenValue("Saturday").
		Case(7).ThenValue("Sunday").
		Default("Invalid day")

	fmt.Println(dayName)

	// Output:
	// Wednesday
}

func ExampleSwitchCase_When() {
	// Using Switch with When for custom predicates
	text := "HELLO"

	result := to.Switch[string, string](text).
		When(func(s string) bool { return s == "" }).
		ThenValue("Empty string").
		When(func(s string) bool { return len(s) < 5 }).
		ThenValue("Short string").
		When(func(s string) bool { return strings.ToUpper(s) == s }).
		ThenValue("All uppercase").
		Default("Normal string")

	fmt.Println(result)

	// Output:
	// All uppercase
}

func ExampleSwitchThen_Then() {
	// Using Switch with Then for computed results
	number := 15

	result := to.Switch[int, string](number).
		Case(0).Then(func(int) string { return "Zero" }).
		When(func(n int) bool { return n%3 == 0 && n%5 == 0 }).Then(func(n int) string {
		return fmt.Sprintf("FizzBuzz: %d", n)
	}).
		Default("Number")

	fmt.Println(result)

	// Output:
	// FizzBuzz: 15
}

func ExampleSwitchThen_ThenValue() {
	// Using Switch with Then for computed results
	number := 15

	result := to.Switch[int, string](number).
		Case(0).ThenValue("Zero").
		When(func(n int) bool { return n%3 == 0 }).ThenValue("Fizz").
		When(func(n int) bool { return n%5 == 0 }).ThenValue("Buzz").
		Default("Number")

	fmt.Println(result)

	// Output:
	// Fizz
}

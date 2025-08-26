package slogx_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func ExampleIsDebug() {
	debugLogger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelDebug).
		WritingToConsole().
		WithTextFormat().
		Logger()

	infoLogger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelInfo).
		WritingToConsole().
		WithTextFormat().
		Logger()

	fmt.Println(slogx.IsDebug(debugLogger))
	fmt.Println(slogx.IsDebug(infoLogger))

	// Output:
	// true
	// false
}

func ExampleIsInfo() {
	infoLogger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelInfo).
		WritingToConsole().
		WithTextFormat().
		Logger()

	warnLogger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelWarn).
		WritingToConsole().
		WithTextFormat().
		Logger()

	fmt.Println(slogx.IsInfo(infoLogger))
	fmt.Println(slogx.IsInfo(warnLogger))

	// Output:
	// true
	// false
}

func ExampleIsWarn() {
	warnLogger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelWarn).
		WritingToConsole().
		WithTextFormat().
		Logger()

	errorLogger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelError).
		WritingToConsole().
		WithTextFormat().
		Logger()

	fmt.Println(slogx.IsWarn(warnLogger))
	fmt.Println(slogx.IsWarn(errorLogger))

	// Output:
	// true
	// false
}

func ExampleIsError() {
	errorLogger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelError).
		WritingToConsole().
		WithTextFormat().
		Logger()

	infoLogger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelInfo).
		WritingToConsole().
		WithTextFormat().
		Logger()

	noneLogger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelNone).
		WritingToConsole().
		WithTextFormat().
		Logger()

	fmt.Println(slogx.IsError(errorLogger))
	fmt.Println(slogx.IsError(infoLogger))
	fmt.Println(slogx.IsError(noneLogger))

	// Output:
	// true
	// true
	// false
}

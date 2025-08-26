package testingx

import (
	"context"
	"fmt"
)

// AssertTB is a subinterface of testing.TB that fulfills the needs of assert package from testify library, but also includes the Helper method.
// This package also provides an implementation of this interface, called testingx.E, that can be used in runnable examples.
type AssertTB interface {
	Helper()
	Errorf(format string, args ...interface{})
}

// RequireTB is a subinterface of testing.TB that fulfills the needs of require package from testify library, but also includes the FailNow method.
// This package also provides an implementation of this interface, called testingx.E, that can be used in runnable examples.
type RequireTB interface {
	AssertTB
	FailNow()
}

// PrintingTBE is a subinterface of testing.TB that provides the printing methods of testing.TB (Log, Logf, Error, Errorf) and the Helper method.
type PrintingTBE interface {
	Helper()
	Errorf(format string, args ...interface{})
	Error(args ...any)
	Log(args ...any)
	Logf(format string, args ...any)
}

// TBE is a subinterface that gives reasonable minimum set of methods from testing.TB, that can be used in tests and runnable examples.
type TBE interface {
	PrintingTBE
	FailNow()
}

// TB is a copy of testing.TB interface, that also allows for implementing it.
// This package also provides an implementation of this interface, called testingx.E, that can be used in runnable examples.
type TB interface {
	Cleanup(func())
	Error(args ...any)
	Errorf(format string, args ...any)
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Name() string
	Setenv(key, value string)
	Chdir(dir string)
	Skip(args ...any)
	SkipNow()
	Skipf(format string, args ...any)
	Skipped() bool
	TempDir() string
	Context() context.Context
}

var _ TB = (*E)(nil)

// E is emulating a testing.TB for runnable examples.
// It is useful to show usage of the library elements dedicated to tests, how they can be used.
// Unfortunately, testing.TB cannot be implemented because it has a private method on the interface.
// This is why it's preferable to use one of testingx or own declared interfaces as argument types instead.
//
// WARNING: This struct is still under development, and still some parts are not implemented completely.
type E struct {
	Verbose bool
}

// Cleanup is a testing.TB Cleanup method implementation for runnable examples.
func (e *E) Cleanup(f func()) {
	// TODO implement me
}

// Fail is a testing.TB Fail method implementation for runnable examples.
func (e *E) Fail() {
	// TODO implement me
	panic(fmt.Sprintf("Fail was called on %T", e))
}

// Failed is a testing.TB Failed method implementation for runnable examples.
func (e *E) Failed() bool {
	// TODO implement me
	panic(fmt.Sprintf("Failed was called on %T", e))
}

// Fatal is a testing.TB Fatal method implementation for runnable examples.
func (e *E) Fatal(args ...any) {
	e.Error(args...)
	panic(fmt.Sprintf("Fatal was called on %T: %s", e, fmt.Sprint(args...)))
}

// Fatalf is a testing.TB Fatalf method implementation for runnable examples.
func (e *E) Fatalf(format string, args ...any) {
	e.Errorf(format, args...)
	panic(fmt.Sprintf("Fatalf was called on %T: %s", e, fmt.Sprintf(format, args...)))
}

// Name is a testing.TB Name method implementation for runnable examples.
func (e *E) Name() string {
	// TODO implement me
	return "Example"
}

// Setenv is a testing.TB Setenv method implementation for runnable examples.
func (e *E) Setenv(key, value string) {
	// TODO implement me
	panic("Setenv not implemented yet")
}

// Chdir is a testing.TB Chdir method implementation for runnable examples.
func (e *E) Chdir(dir string) {
	panic("Chdir not implemented yet")
}

// Skip is a testing.TB Skip method implementation for runnable examples.
func (e *E) Skip(args ...any) {
	// TODO implement me
	fmt.Println("SKIP: Skip is not implemented yet.")
	fmt.Println(append([]any{"SKIP: args:"}, args...)...)
}

// SkipNow is a testing.TB SkipNow method implementation for runnable examples.
func (e *E) SkipNow() {
	// TODO implement me
	fmt.Println("SKIP: SkipNow is not implemented yet.")
}

// Skipf is a testing.TB Skipf method implementation for runnable examples.
func (e *E) Skipf(format string, args ...any) {
	// TODO implement me
	fmt.Println("SKIP: Skipf is not implemented yet.")
	fmt.Println(append([]any{"SKIP: msg:"}, fmt.Sprintf(format, args...))...)
}

// Skipped is a testing.TB Skipped method implementation for runnable examples.
func (e *E) Skipped() bool {
	// TODO implement me
	return false
}

// TempDir is a testing.TB TempDir method implementation for runnable examples.
func (e *E) TempDir() string {
	// TODO implement me
	panic("TempDir not implemented yet")
}

// Context is a testing.TB Context method implementation for runnable examples.
func (e *E) Context() context.Context {
	// TODO implement me
	return context.Background()
}

// Helper is a testing.TB Helper method implementation for runnable examples.
func (e *E) Helper() {
	// Nothing to do here in case of Examples.
}

// FailNow is a testing.TB FailNow method implementation for runnable examples.
func (e *E) FailNow() {
	panic("FailNow was called")
}

// Error is a testing.TB Error method implementation for runnable examples.
// It writes the error message with "ERROR:" prefix.
func (e *E) Error(args ...any) {
	fmt.Println(append([]any{"ERROR:"}, args...)...)
}

// Errorf is a testing.TB Errorf method implementation for runnable examples.
// It writes the error message with "ERROR:" prefix.
func (e *E) Errorf(format string, args ...interface{}) {
	fmt.Printf("ERROR:"+format+"\n", args...)
}

// Log is a testing.TB Log method implementation for runnable examples.
func (e *E) Log(args ...any) {
	if e.Verbose {
		fmt.Println(args...)
	}
}

// Logf is a testing.TB Logf method implementation for runnable examples.
func (e *E) Logf(format string, args ...any) {
	if e.Verbose {
		fmt.Printf(format, args...)
	}
}

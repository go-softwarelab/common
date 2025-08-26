# testingx

```go
import "github.com/go-softwarelab/common/pkg/testingx"
```

Package testingx provides set of utilities for tests and runnable examples.



<a name="AssertTB"></a>
## type [AssertTB](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L10-L13>)

AssertTB is a subinterface of testing.TB that fulfills the needs of assert package from testify library, but also includes the Helper method. This package also provides an implementation of this interface, called testingx.E, that can be used in runnable examples.

```go
type AssertTB interface {
    Helper()
    Errorf(format string, args ...interface{})
}
```

<a name="E"></a>
## type [E](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L70-L72>)

E is emulating a testing.TB for runnable examples. It is useful to show usage of the library elements dedicated to tests, how they can be used. Unfortunately, testing.TB cannot be implemented because it has a private method on the interface. This is why it's preferable to use one of testingx or own declared interfaces as argument types instead.

WARNING: This struct is still under development, and still some parts are not implemented completely.

```go
type E struct {
    Verbose bool
}
```

<a name="E.Chdir"></a>
### [\*E.Chdir](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L116>)

```go
func (e *E) Chdir(dir string)
```

Chdir is a testing.TB Chdir method implementation for runnable examples.

<a name="E.Cleanup"></a>
### [\*E.Cleanup](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L75>)

```go
func (e *E) Cleanup(f func())
```

Cleanup is a testing.TB Cleanup method implementation for runnable examples.

<a name="E.Context"></a>
### [\*E.Context](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L153>)

```go
func (e *E) Context() context.Context
```

Context is a testing.TB Context method implementation for runnable examples.

<a name="E.Error"></a>
### [\*E.Error](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L170>)

```go
func (e *E) Error(args ...any)
```

Error is a testing.TB Error method implementation for runnable examples. It writes the error message with "ERROR:" prefix.

<a name="E.Errorf"></a>
### [\*E.Errorf](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L176>)

```go
func (e *E) Errorf(format string, args ...interface{})
```

Errorf is a testing.TB Errorf method implementation for runnable examples. It writes the error message with "ERROR:" prefix.

<a name="E.Fail"></a>
### [\*E.Fail](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L80>)

```go
func (e *E) Fail()
```

Fail is a testing.TB Fail method implementation for runnable examples.

<a name="E.FailNow"></a>
### [\*E.FailNow](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L164>)

```go
func (e *E) FailNow()
```

FailNow is a testing.TB FailNow method implementation for runnable examples.

<a name="E.Failed"></a>
### [\*E.Failed](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L86>)

```go
func (e *E) Failed() bool
```

Failed is a testing.TB Failed method implementation for runnable examples.

<a name="E.Fatal"></a>
### [\*E.Fatal](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L92>)

```go
func (e *E) Fatal(args ...any)
```

Fatal is a testing.TB Fatal method implementation for runnable examples.

<a name="E.Fatalf"></a>
### [\*E.Fatalf](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L98>)

```go
func (e *E) Fatalf(format string, args ...any)
```

Fatalf is a testing.TB Fatalf method implementation for runnable examples.

<a name="E.Helper"></a>
### [\*E.Helper](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L159>)

```go
func (e *E) Helper()
```

Helper is a testing.TB Helper method implementation for runnable examples.

<a name="E.Log"></a>
### [\*E.Log](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L181>)

```go
func (e *E) Log(args ...any)
```

Log is a testing.TB Log method implementation for runnable examples.

<a name="E.Logf"></a>
### [\*E.Logf](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L188>)

```go
func (e *E) Logf(format string, args ...any)
```

Logf is a testing.TB Logf method implementation for runnable examples.

<a name="E.Name"></a>
### [\*E.Name](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L104>)

```go
func (e *E) Name() string
```

Name is a testing.TB Name method implementation for runnable examples.

<a name="E.Setenv"></a>
### [\*E.Setenv](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L110>)

```go
func (e *E) Setenv(key, value string)
```

Setenv is a testing.TB Setenv method implementation for runnable examples.

<a name="E.Skip"></a>
### [\*E.Skip](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L121>)

```go
func (e *E) Skip(args ...any)
```

Skip is a testing.TB Skip method implementation for runnable examples.

<a name="E.SkipNow"></a>
### [\*E.SkipNow](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L128>)

```go
func (e *E) SkipNow()
```

SkipNow is a testing.TB SkipNow method implementation for runnable examples.

<a name="E.Skipf"></a>
### [\*E.Skipf](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L134>)

```go
func (e *E) Skipf(format string, args ...any)
```

Skipf is a testing.TB Skipf method implementation for runnable examples.

<a name="E.Skipped"></a>
### [\*E.Skipped](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L141>)

```go
func (e *E) Skipped() bool
```

Skipped is a testing.TB Skipped method implementation for runnable examples.

<a name="E.TempDir"></a>
### [\*E.TempDir](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L147>)

```go
func (e *E) TempDir() string
```

TempDir is a testing.TB TempDir method implementation for runnable examples.

<a name="PrintingTBE"></a>
## type [PrintingTBE](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L23-L29>)

PrintingTBE is a subinterface of testing.TB that provides the printing methods of testing.TB \(Log, Logf, Error, Errorf\) and the Helper method.

```go
type PrintingTBE interface {
    Helper()
    Errorf(format string, args ...interface{})
    Error(args ...any)
    Log(args ...any)
    Logf(format string, args ...any)
}
```

<a name="RequireTB"></a>
## type [RequireTB](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L17-L20>)

RequireTB is a subinterface of testing.TB that fulfills the needs of require package from testify library, but also includes the FailNow method. This package also provides an implementation of this interface, called testingx.E, that can be used in runnable examples.

```go
type RequireTB interface {
    AssertTB
    FailNow()
}
```

<a name="TB"></a>
## type [TB](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L39-L60>)

TB is a copy of testing.TB interface, that also allows for implementing it. This package also provides an implementation of this interface, called testingx.E, that can be used in runnable examples.

```go
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
```

<a name="TBE"></a>
## type [TBE](<https://github.com/go-softwarelab/common/blob/main/pkg/testingx/testing_tbe.go#L32-L35>)

TBE is a subinterface that gives reasonable minimum set of methods from testing.TB, that can be used in tests and runnable examples.

```go
type TBE interface {
    PrintingTBE
    FailNow()
}
```
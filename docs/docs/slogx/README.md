# slogx

```go
import "github.com/go-softwarelab/common/pkg/slogx"
```

Package slogx provides utility functions and helpers to complement the Go slog package. It offers convenient methods for creating and configuring loggers for both production and testing environments. Additionally, it includes tools to facilitate testing of logging output.

This package also provides a set of functions for creating slog.Args in a standardized way for some of the most popular use cases, helping ensure consistency and clarity in structured logging across your codebase.

The main goal of this package is to simplify logger setup, reduce boilerplate, and enhance the ergonomics of consistent structured logging throughout application code and tests.



## Constants

<a name="ServiceKey"></a>

```go
const (

    // ServiceKey is a predefined constant used as a key for identifying services or components in structured logging.
    ServiceKey = "service"
    // ComponentKey is a predefined constant used as a key for identifying components in structured logging.
    ComponentKey = "component"
    // ErrorKey is a predefined constant used as a key for identifying errors in structured logging.
    ErrorKey = "error"
    // UserIDKey is a predefined constant used as a key for identifying user IDs in structured logging.
    UserIDKey = "userId"
)
```

<a name="LevelNone"></a>LevelNone is a special log level that disables all logging.

```go
const LevelNone slog.Level = math.MaxInt
```

<a name="Child"></a>
## [Child](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L126>)

```go
func Child(logger *slog.Logger, serviceName string) *slog.Logger
```

Child returns a new logger with the specified service name added to its attributes. If the provided logger is nil, it uses the default slog logger as the base.

<a name="ChildForComponent"></a>
## [ChildForComponent](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L135>)

```go
func ChildForComponent(logger *slog.Logger, componentName string) *slog.Logger
```

ChildForComponent returns a new logger with the specified component name added to its attributes. If the provided logger is nil, it uses the default slog logger as the base. It is strongly recommended to use slogx.Child instead of this function. However, if you need to distinguish components \(such as library tools\) from services, this function can be useful.

<a name="Component"></a>
## [Component](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/args.go#L39>)

```go
func Component(componentName string) slog.Attr
```

Component creates a slog.Attr with the predefined ComponentKey and the given componentName. This is a conventional attribute for marking loggers for components in an application. It is strongly recommended to use slogx.Service instead of this function. However, if you need to distinguish components \(such as library tools\) from services, this function can be useful.

<a name="DefaultIfNil"></a>
## [DefaultIfNil](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L117>)

```go
func DefaultIfNil(logger *slog.Logger) *slog.Logger
```

DefaultIfNil returns the default slog.Logger if the given logger is nil.

<a name="Error"></a>
## [Error](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/args.go#L24>)

```go
func Error(err error) slog.Attr
```

Error returns a slog.Attr containing the provided error message under the "error" key.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	err := fmt.Errorf("test error")
	fmt.Println(slogx.Error(err))
}
```

**Output**

```
error=test error
```


</details>

<a name="IsDebug"></a>
## [IsDebug](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/utils.go#L9>)

```go
func IsDebug(logger *slog.Logger) bool
```

IsDebug checks if the provided logger has the debug level enabled.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
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

}
```

**Output**

```
true
false
```


</details>

<a name="IsError"></a>
## [IsError](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/utils.go#L24>)

```go
func IsError(logger *slog.Logger) bool
```

IsError checks if the provided logger is enabled at the error logging level.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
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

}
```

**Output**

```
true
true
false
```


</details>

<a name="IsInfo"></a>
## [IsInfo](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/utils.go#L14>)

```go
func IsInfo(logger *slog.Logger) bool
```

IsInfo checks if the provided logger has the info level enabled.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
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

}
```

**Output**

```
true
false
```


</details>

<a name="IsWarn"></a>
## [IsWarn](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/utils.go#L19>)

```go
func IsWarn(logger *slog.Logger) bool
```

IsWarn checks if the logger is enabled for the warning level.

<details>
<summary>Example</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
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

}
```

**Output**

```
true
false
```


</details>

<a name="NewLogger"></a>
## [NewLogger](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L75>)

```go
func NewLogger(opts ...func(options *NewLoggerOptions)) *slog.Logger
```

NewLogger creates a new slog.Logger with customizable options such as log level, writer, and format.

<details>
<summary>Example</summary>




```go
package main

import (
	"github.com/go-softwarelab/common/pkg/slogx"
)

var skipTimeInLogOutputForExamplePurposes = slogx.WithFormat(slogx.TextWithoutTimeFormat)

func main() {
	logger := slogx.NewLogger(skipTimeInLogOutputForExamplePurposes)
	logger.Info("test")

}
```

**Output**

```
level=INFO msg=test
```


</details>

<details>
<summary>Example (With Decorator And Managed Level)</summary>


Example using LogLevelManager with NewLogger and WithDecorator. The example sets a level for a pattern before logger creation, logs, then adds another pattern and logs again to show the change.

```go
package main

import (
	"log/slog"

	"github.com/go-softwarelab/common/pkg/slogx"
)

var skipTimeInLogOutputForExamplePurposes = slogx.WithFormat(slogx.TextWithoutTimeFormat)

func main() {
	// Prepare manager and set a pattern BEFORE creating the logger.
	levelManager := slogx.NewLogLevelManager(slogx.LogLevelWarn)
	_ = levelManager.SetLevelForServicePattern("OrderService", slog.LevelInfo)

	// Build the logger with the manager and a stable text format (without time).
	logger := slogx.NewLogger(skipTimeInLogOutputForExamplePurposes, slogx.WithDecorator(levelManager))

	// First, create a child logger for the service.
	orderSvc := slogx.Child(logger, "OrderService")

	// At INFO level for OrderService, debug is filtered, info is printed.
	orderSvc.Debug("dbg-1")
	orderSvc.Info("info-1")

	// Now add a more specific pattern AFTER logger creation.
	_ = levelManager.SetLevelForServicePattern("OrderService.Payment", slog.LevelDebug)

	// Create a more specific child to include the Payment sub-service.
	paymentSvc := slogx.Child(orderSvc, "Payment")

	// At DEBUG level for OrderService.Payment, both debug and info are printed.
	paymentSvc.Debug("dbg-2")
	paymentSvc.Info("info-2")

}
```

**Output**

```
level=INFO msg=info-1 service=OrderService
level=DEBUG msg=dbg-2 service=OrderService service=Payment
level=INFO msg=info-2 service=OrderService service=Payment
```


</details>

<details>
<summary>Example (With Level)</summary>


Example showing how to set a different log level using WithLevel.

```go
package main

import (
	"log/slog"

	"github.com/go-softwarelab/common/pkg/slogx"
)

var skipTimeInLogOutputForExamplePurposes = slogx.WithFormat(slogx.TextWithoutTimeFormat)

func main() {
	logger := slogx.NewLogger(skipTimeInLogOutputForExamplePurposes, slogx.WithLevel(slog.LevelWarn))

	logger.Info("info") // filtered out at WARN level
	logger.Warn("warn") // printed

}
```

**Output**

```
level=WARN msg=warn
```


</details>

<a name="NewTestLogger"></a>
## [NewTestLogger](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L100>)

```go
func NewTestLogger(t TestingTBOutput, opts ...func(options *NewLoggerOptions)) *slog.Logger
```

NewTestLogger creates a new logger instance configured for usage in tests. It writes log output through the provided testing.TB interface with debug level enabled.

<details>
<summary>Example</summary>




```go
package main

import (
	"github.com/go-softwarelab/common/pkg/slogx"
	"github.com/go-softwarelab/common/pkg/testingx"
)

var skipTimeInLogOutputForExamplePurposes = slogx.WithFormat(slogx.TextWithoutTimeFormat)

func main() {
	// We're simulating a test here in example.
	t := &testingx.E{Verbose: true}

	logger := slogx.NewTestLogger(t, skipTimeInLogOutputForExamplePurposes)
	logger.Info("test")

}
```

**Output**

```
level=INFO msg=test
```


</details>

<a name="Number"></a>
## [Number](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/args.go#L44>)

```go
func Number[T types.Number](key string, value T) slog.Attr
```

Number makes easier creation slog.Attr based on any number \(int, float, uint\) or the custom type over the number type.

<details>
<summary>Example (Custom Type)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	type MyInt int
	type MyFloat float64
	type MyUint uint

	fmt.Println(slogx.Number("key", MyInt(42)))
	fmt.Println(slogx.Number("key", MyFloat(42.5)))
	fmt.Println(slogx.Number("key", MyUint(42)))
}
```

**Output**

```
key=42
key=42.5
key=42
```


</details>

<details>
<summary>Example (Float)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	fmt.Println(slogx.Number("key", float32(42.5)))
	fmt.Println(slogx.Number("key", float64(42.5)))
}
```

**Output**

```
key=42.5
key=42.5
```


</details>

<details>
<summary>Example (Int)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	fmt.Println(slogx.Number("key", 42))
	fmt.Println(slogx.Number("key", int8(42)))
	fmt.Println(slogx.Number("key", int16(42)))
	fmt.Println(slogx.Number("key", int32(42)))
	fmt.Println(slogx.Number("key", int64(42)))
}
```

**Output**

```
key=42
key=42
key=42
key=42
key=42
```


</details>

<details>
<summary>Example (Uint)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	fmt.Println(slogx.Number("key", uint(42)))
	fmt.Println(slogx.Number("key", uint8(42)))
	fmt.Println(slogx.Number("key", uint16(42)))
	fmt.Println(slogx.Number("key", uint32(42)))
	fmt.Println(slogx.Number("key", uint64(42)))
}
```

**Output**

```
key=42
key=42
key=42
key=42
key=42
```


</details>

<a name="OptionalNumber"></a>
## [OptionalNumber](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/args.go#L54>)

```go
func OptionalNumber[T types.Number](key string, value *T) slog.Attr
```

OptionalNumber makes easier creation slog.Attr based on pointer to any number \(int, float, uint\) or the custom type over the number type.

<details>
<summary>Example (Custom Type)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	type MyInt int
	n := MyInt(42)
	fmt.Println(slogx.OptionalNumber("key", &n))
}
```

**Output**

```
key=42
```


</details>

<details>
<summary>Example (Float)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	n := 42.5
	fmt.Println(slogx.OptionalNumber("key", &n))
}
```

**Output**

```
key=42.5
```


</details>

<details>
<summary>Example (Int)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	n := 42
	fmt.Println(slogx.OptionalNumber("key", &n))
}
```

**Output**

```
key=42
```


</details>

<details>
<summary>Example (Nil)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	var n *int
	fmt.Println(slogx.OptionalNumber("key", n))
}
```

**Output**

```
key=<nil>
```


</details>

<a name="OptionalString"></a>
## [OptionalString](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/args.go#L69>)

```go
func OptionalString[S ~string](key string, value *S) slog.Attr
```

OptionalString returns a slog.Attr for a string pointer. If the pointer is nil, it sets the value to "\<nil\>".

<details>
<summary>Example (Basic)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	s := "value"
	fmt.Println(slogx.OptionalString("key", &s))
}
```

**Output**

```
key=value
```


</details>

<details>
<summary>Example (Custom Type)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	type MyString string
	s := MyString("custom value")
	fmt.Println(slogx.OptionalString("key", &s))
}
```

**Output**

```
key=custom value
```


</details>

<details>
<summary>Example (Nil)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	var s *string
	fmt.Println(slogx.OptionalString("key", s))
}
```

**Output**

```
key=<nil>
```


</details>

<a name="OptionalUserID"></a>
## [OptionalUserID](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/args.go#L87>)

```go
func OptionalUserID[ID types.Number | ~string](userID *ID) slog.Attr
```

OptionalUserID returns a slog.Attr representing the "userId" or a default value of "\<unknown\>" if userID is nil.

<details>
<summary>Example (Int)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	id := 42
	fmt.Println(slogx.OptionalUserID(&id))
}
```

**Output**

```
userId=42
```


</details>

<details>
<summary>Example (Nil)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	var id *string
	fmt.Println(slogx.OptionalUserID(id))
}
```

**Output**

```
userId=<unknown>
```


</details>

<details>
<summary>Example (String)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	id := "user123"
	fmt.Println(slogx.OptionalUserID(&id))
}
```

**Output**

```
userId=user123
```


</details>

<a name="Service"></a>
## [Service](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/args.go#L30>)

```go
func Service(serviceName string) slog.Attr
```

Service creates a slog.Attr with the predefined ServiceKey and the given serviceName. This is a conventional attr for marking loggers for services/components in application.

<a name="SilentLogger"></a>
## [SilentLogger](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L14>)

```go
func SilentLogger() *slog.Logger
```

SilentLogger returns a new logger instance configured with a handler that discards all log output, effectively creating a logger that won't log any messages.

<details>
<summary>Example</summary>




```go
package main

import (
	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	logger := slogx.SilentLogger()
	logger.Info("this message will be discarded")
	logger.Error("this error will also be discarded")

}
```

**Output**

```

```


</details>

<a name="String"></a>
## [String](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/args.go#L64>)

```go
func String[S ~string](key string, value S) slog.Attr
```

String returns a slog.Attr with the provided key and string value. It's almost the same as a slog.String function, but allows for any custom type based on string to be passed as value.

<details>
<summary>Example (Basic)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	fmt.Println(slogx.String("key", "value"))
	fmt.Println(slogx.String("empty", ""))
}
```

**Output**

```
key=value
empty=
```


</details>

<details>
<summary>Example (Custom Type)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	type MyString string
	s := MyString("custom value")
	fmt.Println(slogx.String("key", s))
}
```

**Output**

```
key=custom value
```


</details>

<details>
<summary>Example (Empty)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	var s string
	fmt.Println(slogx.String("key", s))
}
```

**Output**

```
key=
```


</details>

<a name="UserID"></a>
## [UserID](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/args.go#L77>)

```go
func UserID[ID types.Number | ~string](userID ID) slog.Attr
```

UserID creates a slog.Attr representing the "userId".

<details>
<summary>Example (Custom Type)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	type UserIDInt int
	type UserIDString string

	fmt.Println(slogx.UserID(UserIDInt(42)))
	fmt.Println(slogx.UserID(UserIDString("user123")))
}
```

**Output**

```
userId=42
userId=user123
```


</details>

<details>
<summary>Example (Int)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	fmt.Println(slogx.UserID(42))
}
```

**Output**

```
userId=42
```


</details>

<details>
<summary>Example (String)</summary>




```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/slogx"
)

func main() {
	fmt.Println(slogx.UserID("user123"))
}
```

**Output**

```
userId=user123
```


</details>

<a name="WithAdditionalDotPatternAttrKeys"></a>
## [WithAdditionalDotPatternAttrKeys](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/log_level_manager.go#L29>)

```go
func WithAdditionalDotPatternAttrKeys(keys ...string) func(*DynamicLogLevelOptions)
```

WithAdditionalDotPatternAttrKeys adds additional keys to the list of keys used for dynamic log level matching. By default, the keys are "service" and "component". The keys are matched against the logger attributes keys and the values are matched against the pattern.

<a name="WithDecorator"></a>
## [WithDecorator](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L68>)

```go
func WithDecorator(decorator ...HandlerDecorator) func(*NewLoggerOptions)
```

WithDecorator adds a handler decorator to the logger. The decorator is applied to the handler after it is created. The order of decorators applies from the first to the last.

<a name="WithFormat"></a>
## [WithFormat](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L52>)

```go
func WithFormat(format LogFormat) func(*NewLoggerOptions)
```

WithFormat sets the log format for the logger. To setup logger for tests \(NewTestLogger\) use WithTestLoggerFormat instead

<a name="WithLevel"></a>
## [WithLevel](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L34>)

```go
func WithLevel[L LogLevel | slog.Level](level L) func(*NewLoggerOptions)
```

WithLevel sets the logging level for the logger. To setup logger for tests \(NewTestLogger\) use WithTestLoggerLevel instead

<a name="WithWriter"></a>
## [WithWriter](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L59>)

```go
func WithWriter(writer io.Writer) func(*NewLoggerOptions)
```

WithWriter returns a functional option to set a custom io.Writer for logger to output to.

<a name="CollectingLogsWriter"></a>
## type [CollectingLogsWriter](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/writers.go#L40-L42>)

CollectingLogsWriter is a simple io.Writer implementation that writes to a string builder. It is useful for testing purposes \- to check what was written by the logger.

```go
type CollectingLogsWriter struct {
    strings.Builder
}
```

<a name="NewCollectingLogsWriter"></a>
### [NewCollectingLogsWriter](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/writers.go#L45>)

```go
func NewCollectingLogsWriter() *CollectingLogsWriter
```

NewCollectingLogsWriter creates a new CollectingLogsWriter.

<a name="CollectingLogsWriter.Clear"></a>
### [\*CollectingLogsWriter.Clear](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/writers.go#L55>)

```go
func (w *CollectingLogsWriter) Clear()
```

Clear the stored log output.

<a name="CollectingLogsWriter.Lines"></a>
### [\*CollectingLogsWriter.Lines](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/writers.go#L50>)

```go
func (w *CollectingLogsWriter) Lines() []string
```

Lines returns a slice of lines from the log output.

<a name="DecoratorOptions"></a>
## type [DecoratorOptions](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/handler_decorator.go#L6-L8>)

DecoratorOptions is the options for decorating a slog.Handler.

```go
type DecoratorOptions struct {
    Additional map[string]any
}
```

<a name="DynamicLogLevelOptions"></a>
## type [DynamicLogLevelOptions](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/log_level_manager.go#L22-L24>)



```go
type DynamicLogLevelOptions struct {
    // contains filtered or unexported fields
}
```

<a name="HandlerDecorator"></a>
## type [HandlerDecorator](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/handler_decorator.go#L11-L14>)

HandlerDecorator is an interface for decorating a slog.Handler.

```go
type HandlerDecorator interface {
    // DecorateHandler decorates the provided handler and returns the decorated handler.
    DecorateHandler(handler slog.Handler, options *DecoratorOptions) slog.Handler
}
```

<a name="HandlerDecoratorFunc"></a>
## type [HandlerDecoratorFunc](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/handler_decorator.go#L17>)

HandlerDecoratorFunc is a function that implements HandlerDecorator interface.

```go
type HandlerDecoratorFunc func(handler slog.Handler) slog.Handler
```

<a name="HandlerDecoratorFunc.DecorateHandler"></a>
### [HandlerDecoratorFunc.DecorateHandler](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/handler_decorator.go#L20>)

```go
func (f HandlerDecoratorFunc) DecorateHandler(handler slog.Handler, _ *DecoratorOptions) slog.Handler
```

DecorateHandler implements HandlerDecorator interface.

<a name="HandlerDecoratorFuncWithOptions"></a>
## type [HandlerDecoratorFuncWithOptions](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/handler_decorator.go#L25>)

HandlerDecoratorFuncWithOptions is a function that implements HandlerDecorator interface.

```go
type HandlerDecoratorFuncWithOptions func(handler slog.Handler, options *DecoratorOptions) slog.Handler
```

<a name="HandlerDecoratorFuncWithOptions.DecorateHandler"></a>
### [HandlerDecoratorFuncWithOptions.DecorateHandler](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/handler_decorator.go#L28>)

```go
func (f HandlerDecoratorFuncWithOptions) DecorateHandler(handler slog.Handler, options *DecoratorOptions) slog.Handler
```

DecorateHandler implements HandlerDecorator interface.

<a name="LogFormat"></a>
## type [LogFormat](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/configuration.go#L69>)

LogFormat represents different log handler types which can be configured.

```go
type LogFormat string
```

<a name="JSONFormat"></a>Supported handler types \(based on slog\).

```go
const (
    // JSONFormat is a standard slog json format.
    JSONFormat LogFormat = "json"
    // TextFormat is a standard slog text format.
    TextFormat LogFormat = "text"
    // TextWithoutTimeFormat is a text format without showing the time, it is mostly useful for testing or examples.
    TextWithoutTimeFormat LogFormat = "text-no-time"
)
```

<a name="ParseLogFormat"></a>
### [ParseLogFormat](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/configuration.go#L82>)

```go
func ParseLogFormat[H ~string](handlerType H) (LogFormat, error)
```

ParseLogFormat parses a string into a LogFormat enum \(case\-insensitive\).

<a name="LogLevel"></a>
## type [LogLevel](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/configuration.go#L16>)

LogLevel string representation of different log levels, which can be configured.

```go
type LogLevel string
```

<a name="LogLevelDebug"></a>Supported log levels \(based on slog\).

```go
const (
    LogLevelDebug LogLevel = "debug"
    LogLevelInfo  LogLevel = "info"
    LogLevelWarn  LogLevel = "warn"
    LogLevelError LogLevel = "error"
    LogLevelNone  LogLevel = "none"
)
```

<a name="ParseLogLevel"></a>
### [ParseLogLevel](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/configuration.go#L28>)

```go
func ParseLogLevel[L ~string](level L) (LogLevel, error)
```

ParseLogLevel parses a string into a LogLevel enum \(case\-insensitive\).

<a name="LogLevel.GetSlogLevel"></a>
### [LogLevel.GetSlogLevel](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/configuration.go#L51>)

```go
func (l LogLevel) GetSlogLevel() (slog.Level, error)
```

GetSlogLevel returns the slog.Level representation of the LogLevel.

<a name="LogLevel.MustGetSlogLevel"></a>
### [LogLevel.MustGetSlogLevel](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/configuration.go#L42>)

```go
func (l LogLevel) MustGetSlogLevel() slog.Level
```

MustGetSlogLevel returns the slog.Level representation of the LogLevel. Panics if the LogLevel is invalid.

<a name="LogLevelManager"></a>
## type [LogLevelManager](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/log_level_manager.go#L36-L41>)

LogLevelManager is a logger decorator that allows to configure log levels dynamically based on logger attributes.

```go
type LogLevelManager struct {
    // contains filtered or unexported fields
}
```

<a name="NewLogLevelManager"></a>
### [NewLogLevelManager](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/log_level_manager.go#L44>)

```go
func NewLogLevelManager[L slog.Level | LogLevel](level L, opts ...func(*DynamicLogLevelOptions)) *LogLevelManager
```

NewLogLevelManager creates a new LogLevelManager.

<a name="NewLoggerWithManagedLevel"></a>
### [NewLoggerWithManagedLevel](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L93>)

```go
func NewLoggerWithManagedLevel(level LogLevel) (*slog.Logger, *LogLevelManager)
```

NewLoggerWithManagedLevel creates a new slog.Logger and LogLevelManager with the specified log level. LogLevelManager is applied to the resulting logger. Use the LogLevelManager to change the log level for particular child loggers at runtime.

<details>
<summary>Example</summary>


Example using LogLevelManager obtained from NewLoggerWithManagedLevel. We rebuild a logger with the same manager to control the output format in this example.

```go
package main

import (
	"log/slog"

	"github.com/go-softwarelab/common/pkg/slogx"
)

var skipTimeInLogOutputForExamplePurposes = slogx.WithFormat(slogx.TextWithoutTimeFormat)

func main() {
	// Create a manager via helper and then build a logger with a stable format for the example output.
	_, levelManager := slogx.NewLoggerWithManagedLevel(slogx.LogLevelWarn)
	logger := slogx.NewLogger(skipTimeInLogOutputForExamplePurposes, slogx.WithDecorator(levelManager))

	// Create a service-scoped child logger.
	svc := slogx.Child(logger, "OrderService")

	// Before adding a pattern, debug is filtered out (default WARN).
	svc.Debug("debug before change")

	// Enable DEBUG for the specific service using a dot pattern.
	_ = levelManager.SetLevelForServicePattern("OrderService", slog.LevelDebug)

	// Now DEBUG is enabled for OrderService.
	svc.Debug("debug after change")

}
```

**Output**

```
level=DEBUG msg="debug after change" service=OrderService
```


</details>

<a name="LogLevelManager.Decorate"></a>
### [\*LogLevelManager.Decorate](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/log_level_manager.go#L146>)

```go
func (m *LogLevelManager) Decorate(handler slog.Handler) slog.Handler
```

Decorate decorates the given handler with the LogLevelManager.

<a name="LogLevelManager.DecorateHandler"></a>
### [\*LogLevelManager.DecorateHandler](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/log_level_manager.go#L151>)

```go
func (m *LogLevelManager) DecorateHandler(handler slog.Handler, _ *DecoratorOptions) slog.Handler
```

DecorateHandler decorates the given handler with the LogLevelManager.

<a name="LogLevelManager.SetLevel"></a>
### [\*LogLevelManager.SetLevel](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/log_level_manager.go#L71>)

```go
func (m *LogLevelManager) SetLevel(level slog.Level)
```

SetLevel sets the default logging level to the specified slog.Level.

<a name="LogLevelManager.SetLevelForServicePattern"></a>
### [\*LogLevelManager.SetLevelForServicePattern](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/log_level_manager.go#L88>)

```go
func (m *LogLevelManager) SetLevelForServicePattern(pattern string, level slog.Level) error
```

SetLevelForServicePattern associates a logging level with a given simple dot\-separated pattern for dynamic log level matching. Returns an error if the pattern cannot be parsed or the level cannot be set.

Dot\-Service\-Pattern is a string containing dot\-separated services \(or components\) names, such as "Service.Component". The pattern's dot\-separated parts are matched against the logger attributes values whose key is "service" \(or "component"\). The specified level is applied to all attributes that match the pattern. The most specific matching pattern determines the log level. If multiple patterns of equal specificity match, one is chosen arbitrarily.

For example: Given the patterns: "Service1" and "Service1.Service2", and logger attributes: service="Service1", service="Service2", user=1 the level set for the pattern "Service1.Service2" will be used.

<a name="LogLevelManager.SetLevels"></a>
### [\*LogLevelManager.SetLevels](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/log_level_manager.go#L105>)

```go
func (m *LogLevelManager) SetLevels(patterns map[string]any) error
```

SetLevels updates logging levels using a map of patterns and their corresponding levels; returns an error if invalid input. Currently only a dot\-service\-pattern is supported, see SetLevelForServicePattern for details. Value can be:

- string \- string value parseable to slogx.LogLevel
- int: any integer value \- although it's recommended to use slog.Level values
- slogx.LogLevel
- slog.Level

<a name="LogLevelManager.SetLogLevel"></a>
### [\*LogLevelManager.SetLogLevel](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/log_level_manager.go#L66>)

```go
func (m *LogLevelManager) SetLogLevel(level LogLevel)
```

SetLogLevel sets the default logging level to the specified slogx.LogLevel.

<a name="LoggerBuilderWithHandler"></a>
## type [LoggerBuilderWithHandler](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/builder.go#L44-L51>)



```go
type LoggerBuilderWithHandler interface {
    // WithHandlerDecorator adds a handler decorator to the logger.
    // The decorator is applied to the handler after it is created.
    // The order of decorators applies from the first to the last.
    WithHandlerDecorator(decorators ...HandlerDecorator) LoggerBuilderWithHandler

    LoggerFactory
}
```

<a name="LoggerFactory"></a>
## type [LoggerFactory](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/builder.go#L54-L57>)



```go
type LoggerFactory interface {
    // Logger creates and returns a configured *slog.Logger instance based on the current logger configuration.
    Logger() *slog.Logger
}
```

<a name="LoggerHandlerBuilder"></a>
## type [LoggerHandlerBuilder](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/builder.go#L32-L41>)



```go
type LoggerHandlerBuilder interface {
    // WithTextFormat configures the logger to use a text-based log format.
    WithTextFormat() LoggerBuilderWithHandler

    // WithJSONFormat configures the logger to use a JSON-based log format.
    WithJSONFormat() LoggerBuilderWithHandler

    // WithFormat sets the handler of a given type for the logger.
    WithFormat(handlerType LogFormat) LoggerBuilderWithHandler
}
```

<a name="LoggerLevelBuilder"></a>
## type [LoggerLevelBuilder](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/builder.go#L10-L17>)

LoggerLevelBuilder is the main interface for configuring a logger.

```go
type LoggerLevelBuilder interface {
    // WithLevel sets the log level for the logger.
    WithLevel(level LogLevel) LoggerOutputBuilder
    // WithSlogLevel sets the log slog.Level for the logger.
    WithSlogLevel(level slog.Level) LoggerOutputBuilder
    // Silent sets the special logger handler to discard all logs.
    Silent() LoggerFactory
}
```

<a name="NewBuilder"></a>
### [NewBuilder](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/builder.go#L67>)

```go
func NewBuilder() LoggerLevelBuilder
```

NewBuilder creates a new Builder for configuring a logger.

<a name="LoggerOpt"></a>
## type [LoggerOpt](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L27>)

LoggerOpt is an alias for functional option for slogx.NewLogger and slogx.NewTestLogger.

```go
type LoggerOpt = func(*NewLoggerOptions)
```

<a name="LoggerOpts"></a>
## type [LoggerOpts](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L30>)

LoggerOpts is an alias for a slice of functional options for slogx.NewLogger and slogx.NewTestLogger.

```go
type LoggerOpts = []LoggerOpt
```

<a name="LoggerOutputBuilder"></a>
## type [LoggerOutputBuilder](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/builder.go#L20-L29>)



```go
type LoggerOutputBuilder interface {
    // WritingToConsole configures a logger to write logs to the stdout
    WritingToConsole() LoggerHandlerBuilder

    // WritingTo configures the logger to write logs to the provided io.Writer
    WritingTo(writer io.Writer) LoggerHandlerBuilder

    // WithCustomHandler sets a custom slog.Handler for the logger.
    WithCustomHandler(handler slog.Handler) LoggerFactory
}
```

<a name="NewLoggerOptions"></a>
## type [NewLoggerOptions](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/loggers.go#L19-L24>)

NewLoggerOptions is a set of options for slogx.NewLogger and slogx.NewTestLogger.

```go
type NewLoggerOptions struct {
    // contains filtered or unexported fields
}
```

<a name="RestyLogger"></a>
## type [RestyLogger](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/resty_adapter.go#L9-L13>)

RestyLogger represents an interface of logger required by github.com/go\-resty/resty

```go
type RestyLogger interface {
    Errorf(format string, v ...interface{})
    Warnf(format string, v ...interface{})
    Debugf(format string, v ...interface{})
}
```

<a name="RestyAdapter"></a>
### [RestyAdapter](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/resty_adapter.go#L18>)

```go
func RestyAdapter(logger *slog.Logger) RestyLogger
```

RestyAdapter create an adapter on slog.Logger that will allow it to be use with github.com/go\-resty/resty library

<a name="TestingTBOutput"></a>
## type [TestingTBOutput](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/writers.go#L10-L13>)

TestingTBOutput is an interface that represents a testing.TB instance required methods to be used by NewTestLogger and TestingTBWriter.

```go
type TestingTBOutput interface {
    Helper()
    Log(args ...any)
}
```

<a name="TestingTBWriter"></a>
## type [TestingTBWriter](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/writers.go#L17-L19>)

TestingTBWriter is a utility type that implements the io.Writer interface by writing log output to testing.TB. It is commonly used in test scenarios to redirect logs to the test's output via the provided testing.TB instance.

```go
type TestingTBWriter struct {
    // contains filtered or unexported fields
}
```

<a name="NewTestingTBWriter"></a>
### [NewTestingTBWriter](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/writers.go#L22>)

```go
func NewTestingTBWriter(t TestingTBOutput) *TestingTBWriter
```

NewTestingTBWriter creates a new TestingTBWriter that writes output to the provided testing.TB instance.

<a name="TestingTBWriter.Write"></a>
### [\*TestingTBWriter.Write](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/writers.go#L27>)

```go
func (w *TestingTBWriter) Write(p []byte) (n int, err error)
```

Write writes the provided byte slice to the testing.TB instance.

<a name="TestingTBWriter.WriteString"></a>
### [\*TestingTBWriter.WriteString](<https://github.com/go-softwarelab/common/blob/main/pkg/slogx/writers.go#L32>)

```go
func (w *TestingTBWriter) WriteString(s string) (n int, err error)
```

WriteString writes the provided string to the testing.TB instance.
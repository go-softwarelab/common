package slogx

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/go-softwarelab/common/pkg/to"
)

// SilentLogger returns a new logger instance configured with a handler that discards all log output,
// effectively creating a logger that won't log any messages.
func SilentLogger() *slog.Logger {
	return NewBuilder().Silent().Logger()
}

// NewLoggerOptions is a set of options for slogx.NewLogger and slogx.NewTestLogger.
type NewLoggerOptions struct {
	level      slog.Level
	writer     io.Writer
	format     LogFormat
	decorators []HandlerDecorator
}

// LoggerOpt is an alias for functional option for slogx.NewLogger and slogx.NewTestLogger.
type LoggerOpt = func(*NewLoggerOptions)

// LoggerOpts is an alias for a slice of functional options for slogx.NewLogger and slogx.NewTestLogger.
type LoggerOpts = []LoggerOpt

// WithLevel sets the logging level for the logger.
// To setup logger for tests (NewTestLogger) use WithTestLoggerLevel instead
func WithLevel[L LogLevel | slog.Level](level L) func(*NewLoggerOptions) {
	var logLevel slog.Level
	switch l := any(level).(type) {
	case LogLevel:
		logLevel = l.MustGetSlogLevel()
	case slog.Level:
		logLevel = l
	default:
		panic(fmt.Errorf("unexpected type (%T) of level passed without compiler error", level))
	}

	return func(options *NewLoggerOptions) {
		options.level = logLevel
	}
}

// WithFormat sets the log format for the logger.
// To setup logger for tests (NewTestLogger) use WithTestLoggerFormat instead
func WithFormat(format LogFormat) func(*NewLoggerOptions) {
	return func(options *NewLoggerOptions) {
		options.format = format
	}
}

// WithWriter returns a functional option to set a custom io.Writer for logger to output to.
func WithWriter(writer io.Writer) func(*NewLoggerOptions) {
	return func(options *NewLoggerOptions) {
		options.writer = writer
	}
}

// WithDecorator adds a handler decorator to the logger.
// The decorator is applied to the handler after it is created.
// The order of decorators applies from the first to the last.
func WithDecorator(decorator ...HandlerDecorator) func(*NewLoggerOptions) {
	return func(options *NewLoggerOptions) {
		options.decorators = append(options.decorators, decorator...)
	}
}

// NewLogger creates a new slog.Logger with customizable options such as log level, writer, and format.
func NewLogger(opts ...func(options *NewLoggerOptions)) *slog.Logger {
	options := to.OptionsWithDefault(NewLoggerOptions{
		level:  slog.LevelInfo,
		writer: os.Stdout,
		format: TextFormat,
	}, opts...)

	return NewBuilder().
		WithSlogLevel(options.level).
		WritingTo(options.writer).
		WithFormat(options.format).
		WithHandlerDecorator(options.decorators...).
		Logger()
}

// NewLoggerWithManagedLevel creates a new slog.Logger and LogLevelManager with the specified log level.
// LogLevelManager is applied to the resulting logger.
// Use the LogLevelManager to change the log level for particular child loggers at runtime.
func NewLoggerWithManagedLevel(level LogLevel) (*slog.Logger, *LogLevelManager) {
	levelManager := NewLogLevelManager(level)
	return NewLogger(WithLevel(slog.LevelDebug), WithDecorator(levelManager)), levelManager
}

// NewTestLogger creates a new logger instance configured for usage in tests.
// It writes log output through the provided testing.TB interface with debug level enabled.
func NewTestLogger(t TestingTBOutput, opts ...func(options *NewLoggerOptions)) *slog.Logger {
	opts = append(opts, withTBWritter(t))
	return NewLogger(opts...)
}

func withTBWritter(t TestingTBOutput) func(*NewLoggerOptions) {
	return func(options *NewLoggerOptions) {
		testWriter := NewTestingTBWriter(t)
		if options.writer != os.Stdout {
			options.writer = io.MultiWriter(options.writer, testWriter)
		} else {
			options.writer = testWriter
		}
	}
}

// DefaultIfNil returns the default slog.Logger if the given logger is nil.
func DefaultIfNil(logger *slog.Logger) *slog.Logger {
	if logger == nil {
		return slog.Default()
	}
	return logger
}

// Child returns a new logger with the specified service name added to its attributes.
// If the provided logger is nil, it uses the default slog logger as the base.
func Child(logger *slog.Logger, serviceName string) *slog.Logger {
	return DefaultIfNil(logger).With(Service(serviceName))
}

// ChildForComponent returns a new logger with the specified component name added to its attributes.
// If the provided logger is nil, it uses the default slog logger as the base.
// It is strongly recommended to use slogx.Child instead of this function.
// However, if you need to distinguish components (such as library tools) from services,
// this function can be useful.
func ChildForComponent(logger *slog.Logger, componentName string) *slog.Logger {
	return DefaultIfNil(logger).With(Component(componentName))
}

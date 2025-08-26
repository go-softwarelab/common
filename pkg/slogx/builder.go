package slogx

import (
	"io"
	"log/slog"
	"os"
)

// LoggerLevelBuilder is the main interface for configuring a logger.
type LoggerLevelBuilder interface {
	// WithLevel sets the log level for the logger.
	WithLevel(level LogLevel) LoggerOutputBuilder
	// WithSlogLevel sets the log slog.Level for the logger.
	WithSlogLevel(level slog.Level) LoggerOutputBuilder
	// Silent sets the special logger handler to discard all logs.
	Silent() LoggerFactory
}

//nolint:revive
type LoggerOutputBuilder interface {
	// WritingToConsole configures a logger to write logs to the stdout
	WritingToConsole() LoggerHandlerBuilder

	// WritingTo configures the logger to write logs to the provided io.Writer
	WritingTo(writer io.Writer) LoggerHandlerBuilder

	// WithCustomHandler sets a custom slog.Handler for the logger.
	WithCustomHandler(handler slog.Handler) LoggerFactory
}

//nolint:revive
type LoggerHandlerBuilder interface {
	// WithTextFormat configures the logger to use a text-based log format.
	WithTextFormat() LoggerBuilderWithHandler

	// WithJSONFormat configures the logger to use a JSON-based log format.
	WithJSONFormat() LoggerBuilderWithHandler

	// WithFormat sets the handler of a given type for the logger.
	WithFormat(handlerType LogFormat) LoggerBuilderWithHandler
}

//nolint:revive
type LoggerBuilderWithHandler interface {
	// WithHandlerDecorator adds a handler decorator to the logger.
	// The decorator is applied to the handler after it is created.
	// The order of decorators applies from the first to the last.
	WithHandlerDecorator(decorators ...HandlerDecorator) LoggerBuilderWithHandler

	LoggerFactory
}

//nolint:revive
type LoggerFactory interface {
	// Logger creates and returns a configured *slog.Logger instance based on the current logger configuration.
	Logger() *slog.Logger
}

type builder struct {
	handler    slog.Handler
	level      *slog.LevelVar
	writer     io.Writer
	decorators []HandlerDecorator
}

// NewBuilder creates a new Builder for configuring a logger.
func NewBuilder() LoggerLevelBuilder {
	return &builder{
		level: new(slog.LevelVar),
	}
}

// Silent sets the special logger handler to discard all logs.
func (c *builder) Silent() LoggerFactory {
	c.handler = slog.DiscardHandler
	return c
}

// WithLevel sets the log level for the logger.
func (c *builder) WithLevel(level LogLevel) LoggerOutputBuilder {
	return c.WithSlogLevel(level.MustGetSlogLevel())
}

// WithSlogLevel sets the log slog.Level for the logger.
func (c *builder) WithSlogLevel(level slog.Level) LoggerOutputBuilder {
	c.level.Set(level)
	return c
}

// WritingToConsole configures a logger to write logs to the stdout
func (c *builder) WritingToConsole() LoggerHandlerBuilder {
	return c.WritingTo(os.Stdout)
}

// WritingTo configures the logger to write logs to the provided io.Writer
func (c *builder) WritingTo(writer io.Writer) LoggerHandlerBuilder {
	if writer == nil {
		panic("writer cannot be nil")
	}
	c.writer = writer
	return c
}

// WithCustomHandler sets a custom handler for the logger.
func (c *builder) WithCustomHandler(handler slog.Handler) LoggerFactory {
	c.handler = handler
	return c
}

// WithHandlerDecorator adds a handler decorator to the logger.
// The decorator is applied to the handler after it is created.
// The order of decorators applies from the first to the last.
func (c *builder) WithHandlerDecorator(decorators ...HandlerDecorator) LoggerBuilderWithHandler {
	c.decorators = append(c.decorators, decorators...)
	return c
}

// WithTextFormat sets the log format to "text" for the logger and returns a builder with the handler configured.
func (c *builder) WithTextFormat() LoggerBuilderWithHandler {
	return c.WithFormat(TextFormat)
}

// WithJSONFormat sets the log format to "json" for the logger and returns a builder with the handler configured.
func (c *builder) WithJSONFormat() LoggerBuilderWithHandler {
	return c.WithFormat(JSONFormat)
}

// WithFormat sets a format of logs for the logger.
func (c *builder) WithFormat(handlerType LogFormat) LoggerBuilderWithHandler {
	opts := &slog.HandlerOptions{Level: c.level}

	switch handlerType {
	case JSONFormat:
		c.handler = slog.NewJSONHandler(c.writer, opts)
	case TextFormat:
		c.handler = slog.NewTextHandler(c.writer, opts)
	case TextWithoutTimeFormat:
		opts.ReplaceAttr = func(_ []string, attr slog.Attr) slog.Attr {
			if attr.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return attr
		}
		c.handler = slog.NewTextHandler(c.writer, opts)
	default:
		panic("unsupported handler type")
	}

	return c
}

// Logger creates a new logger from the configuration.
func (c *builder) Logger() *slog.Logger {
	var handler = c.handler
	options := &DecoratorOptions{}

	for _, decorator := range c.decorators {
		handler = decorator.DecorateHandler(handler, options)
	}
	return slog.New(handler)
}

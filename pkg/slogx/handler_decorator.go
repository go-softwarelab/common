package slogx

import "log/slog"

// DecoratorOptions is the options for decorating a slog.Handler.
type DecoratorOptions struct {
	Additional map[string]any
}

// HandlerDecorator is an interface for decorating a slog.Handler.
type HandlerDecorator interface {
	// DecorateHandler decorates the provided handler and returns the decorated handler.
	DecorateHandler(handler slog.Handler, options *DecoratorOptions) slog.Handler
}

// HandlerDecoratorFunc is a function that implements HandlerDecorator interface.
type HandlerDecoratorFunc func(handler slog.Handler) slog.Handler

// DecorateHandler implements HandlerDecorator interface.
func (f HandlerDecoratorFunc) DecorateHandler(handler slog.Handler, _ *DecoratorOptions) slog.Handler {
	return f(handler)
}

// HandlerDecoratorFuncWithOptions is a function that implements HandlerDecorator interface.
type HandlerDecoratorFuncWithOptions func(handler slog.Handler, options *DecoratorOptions) slog.Handler

// DecorateHandler implements HandlerDecorator interface.
func (f HandlerDecoratorFuncWithOptions) DecorateHandler(handler slog.Handler, options *DecoratorOptions) slog.Handler {
	return f(handler, options)
}

package slogx

import (
	"fmt"
	"log/slog"
)

// RestyLogger represents an interface of logger required by github.com/go-resty/resty
type RestyLogger interface {
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

type restyAdapter slog.Logger

// RestyAdapter create an adapter on slog.Logger that will allow it to be use with github.com/go-resty/resty library
func RestyAdapter(logger *slog.Logger) RestyLogger {
	logger = Child(logger, "resty")

	return (*restyAdapter)(logger)
}

// Errorf logs an error message formatted according to a format specifier and optional arguments.
func (l *restyAdapter) Errorf(message string, v ...interface{}) {
	if len(v) > 0 {
		message = fmt.Sprintf(message, v...)
	}
	(*slog.Logger)(l).Error(message)
}

// Warnf logs a warning message formatted according to a format specifier and optional arguments.
func (l *restyAdapter) Warnf(message string, v ...interface{}) {
	if len(v) > 0 {
		message = fmt.Sprintf(message, v...)
	}
	(*slog.Logger)(l).Warn(message)
}

// Debugf logs a debug message formatted according to a format specifier and optional arguments.
func (l *restyAdapter) Debugf(message string, v ...interface{}) {
	if len(v) > 0 {
		message = fmt.Sprintf(message, v...)
	}
	(*slog.Logger)(l).Debug(message)
}

package slogx

import (
	"context"
	"log/slog"
)

// IsDebug checks if the provided logger has the debug level enabled.
func IsDebug(logger *slog.Logger) bool {
	return logger.Enabled(context.Background(), slog.LevelDebug)
}

// IsInfo checks if the provided logger has the info level enabled.
func IsInfo(logger *slog.Logger) bool {
	return logger.Enabled(context.Background(), slog.LevelInfo)
}

// IsWarn checks if the logger is enabled for the warning level.
func IsWarn(logger *slog.Logger) bool {
	return logger.Enabled(context.Background(), slog.LevelWarn)
}

// IsError checks if the provided logger is enabled at the error logging level.
func IsError(logger *slog.Logger) bool {
	return logger.Enabled(context.Background(), slog.LevelError)
}

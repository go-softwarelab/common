package slogx

import (
	"fmt"
	"log/slog"
	"math"
	"strings"

	"github.com/go-softwarelab/common/pkg/to"
)

// LevelNone is a special log level that disables all logging.
const LevelNone slog.Level = math.MaxInt

// LogLevel string representation of different log levels, which can be configured.
type LogLevel string

// Supported log levels (based on slog).
const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
	LogLevelNone  LogLevel = "none"
)

// ParseLogLevel parses a string into a LogLevel enum (case-insensitive).
func ParseLogLevel[L ~string](level L) (LogLevel, error) {
	if LogLevel(level) == LogLevelNone {
		return LogLevelNone, nil
	}
	var slevel slog.Level
	err := slevel.UnmarshalText([]byte(level))
	if err != nil {
		return "", fmt.Errorf("failed to parse log level: %w", err)
	}
	return LogLevel(strings.ToLower(slevel.String())), nil
}

// MustGetSlogLevel returns the slog.Level representation of the LogLevel.
// Panics if the LogLevel is invalid.
func (l LogLevel) MustGetSlogLevel() slog.Level {
	level, err := l.GetSlogLevel()
	if err != nil {
		panic(err)
	}
	return level
}

// GetSlogLevel returns the slog.Level representation of the LogLevel.
func (l LogLevel) GetSlogLevel() (slog.Level, error) {
	switch l {
	case LogLevelDebug:
		return slog.LevelDebug, nil
	case LogLevelInfo:
		return slog.LevelInfo, nil
	case LogLevelWarn:
		return slog.LevelWarn, nil
	case LogLevelError:
		return slog.LevelError, nil
	case LogLevelNone:
		return LevelNone, nil
	default:
		return 0, fmt.Errorf("unexpected value of log level %s", l)
	}
}

// LogFormat represents different log handler types which can be configured.
type LogFormat string

// Supported handler types (based on slog).
const (
	// JSONFormat is a standard slog json format.
	JSONFormat LogFormat = "json"
	// TextFormat is a standard slog text format.
	TextFormat LogFormat = "text"
	// TextWithoutTimeFormat is a text format without showing the time, it is mostly useful for testing or examples.
	TextWithoutTimeFormat LogFormat = "text-no-time"
)

// ParseLogFormat parses a string into a LogFormat enum (case-insensitive).
func ParseLogFormat[H ~string](handlerType H) (LogFormat, error) {
	return to.Enum(handlerType, JSONFormat, TextFormat, TextWithoutTimeFormat) //nolint:wrapcheck
}

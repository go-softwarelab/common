package slogx_test

import (
	"testing"

	"github.com/go-softwarelab/common/pkg/slogx"
	"github.com/stretchr/testify/require"
)

func TestChildLogger(t *testing.T) {
	// given:
	output := slogx.NewCollectingLogsWriter()
	logger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelDebug).
		WritingTo(output).
		WithTextFormat().
		Logger()

	// when:
	childLogger := slogx.Child(logger, "child")

	// and:
	childLogger.Debug("debug message")

	// then:
	msg := output.String()
	require.Contains(t, msg, "service=child")
	require.Contains(t, msg, `msg="debug message"`)
}

func TestNewLogger(t *testing.T) {
	t.Run("logger with debug level", func(t *testing.T) {
		// given:
		output := slogx.NewCollectingLogsWriter()

		// when:
		logger := slogx.NewLogger(
			slogx.WithLevel(slogx.LogLevelDebug),
			slogx.WithWriter(output),
		)

		// and:
		logger.Debug("debug message")

		// then:
		msg := output.String()
		require.Contains(t, msg, "level=DEBUG")
		require.Contains(t, msg, `msg="debug message"`)
	})

	t.Run("logger with error level", func(t *testing.T) {
		// given:
		output := slogx.NewCollectingLogsWriter()

		// when:
		logger := slogx.NewLogger(
			slogx.WithLevel(slogx.LogLevelError),
			slogx.WithWriter(output),
		)

		// and:
		logger.Debug("debug message")

		// then:
		msg := output.String()
		require.NotContains(t, msg, "level=DEBUG")
	})

	t.Run("logger with json format", func(t *testing.T) {
		// given:
		output := slogx.NewCollectingLogsWriter()

		// when:
		logger := slogx.NewLogger(
			slogx.WithFormat(slogx.JSONFormat),
			slogx.WithWriter(output),
		)

		// and:
		logger.Info("info message")

		// then:
		msg := output.String()
		require.Contains(t, msg, `"level":"INFO"`)
		require.Contains(t, msg, `"msg":"info message"`)
	})
}

func TestDefaultIfNil(t *testing.T) {
	// when:
	logger := slogx.DefaultIfNil(nil)

	// then:
	require.NotNil(t, logger)
}

func TestNewTestLogger(t *testing.T) {
	t.Run("default test logger will not panic", func(t *testing.T) {
		// when:
		logger := slogx.NewTestLogger(t)

		// then:
		require.NotNil(t, logger)

		require.NotPanics(t, func() {
			logger.Debug("debug message")
			logger.Info("info message")
			logger.Error("error message")
		})
	})

	t.Run("test logger with configured level will not panic", func(t *testing.T) {
		// when:
		logger := slogx.NewTestLogger(t, slogx.WithLevel(slogx.LogLevelError))

		// then:
		require.NotNil(t, logger)

		require.NotPanics(t, func() {
			logger.Debug("debug message")
			logger.Info("info message")
			logger.Error("error message")
		})
	})

	t.Run("test logger with configured format will not panic", func(t *testing.T) {
		// when:
		logger := slogx.NewTestLogger(t, slogx.WithFormat(slogx.JSONFormat))

		// then:
		require.NotNil(t, logger)

		require.NotPanics(t, func() {
			logger.Debug("debug message")
			logger.Info("info message")
			logger.Error("error message")
		})
	})
}

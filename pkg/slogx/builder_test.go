package slogx_test

import (
	"testing"

	"github.com/go-softwarelab/common/pkg/slogx"
	"github.com/stretchr/testify/require"
)

func TestTextLogger(t *testing.T) {
	// given:
	output := slogx.NewCollectingLogsWriter()
	logger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelDebug).
		WritingTo(output).
		WithFormat(slogx.TextFormat).
		Logger()

	// when:
	logger.Debug("debug message")

	// then:
	msg := output.String()
	require.Contains(t, msg, "time=")
	require.Contains(t, msg, "level=DEBUG")
	require.Contains(t, msg, `msg="debug message"`)
}

func TestJSONLogger(t *testing.T) {
	// given:
	output := slogx.NewCollectingLogsWriter()
	logger := slogx.NewBuilder().
		WithLevel(slogx.LogLevelDebug).
		WritingTo(output).
		WithFormat(slogx.JSONFormat).
		Logger()

	// when:
	logger.Debug("debug message")

	// then:
	msg := output.String()
	require.Contains(t, msg, `"time"`)
	require.Contains(t, msg, `"level":"DEBUG"`)
	require.Contains(t, msg, `"msg":"debug message"`)
}

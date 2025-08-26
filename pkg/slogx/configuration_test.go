package slogx_test

import (
	"log/slog"
	"math"
	"testing"

	"github.com/go-softwarelab/common/pkg/slogx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLogLevel(t *testing.T) {
	cases := []struct {
		in      string
		want    slogx.LogLevel
		wantErr bool
		desc    string
	}{
		{in: "debug", want: slogx.LogLevelDebug, desc: "lowercase debug"},
		{in: "DEBUG", want: slogx.LogLevelDebug, desc: "uppercase debug"},
		{in: "Debug", want: slogx.LogLevelDebug, desc: "mixedcase debug"},
		{in: "info", want: slogx.LogLevelInfo, desc: "info"},
		{in: "INFO", want: slogx.LogLevelInfo, desc: "INFO"},
		{in: "warn", want: slogx.LogLevelWarn, desc: "warn"},
		{in: "WARN", want: slogx.LogLevelWarn, desc: "WARN"},
		{in: "error", want: slogx.LogLevelError, desc: "error"},
		{in: "ERROR", want: slogx.LogLevelError, desc: "ERROR"},
		{in: "none", want: slogx.LogLevelNone, desc: "special none"},
		{in: "NoNe", wantErr: true, desc: "none is case-sensitive special-case"},
		{in: "verbose", wantErr: true, desc: "invalid"},
	}

	for _, tc := range cases {
		got, err := slogx.ParseLogLevel(tc.in)
		if tc.wantErr {
			assert.Error(t, err, tc.desc)
			continue
		}
		require.NoError(t, err, tc.desc)
		assert.Equal(t, tc.want, got, tc.desc)
	}
}

func TestGetSlogLevel(t *testing.T) {
	cases := []struct {
		in      slogx.LogLevel
		want    slog.Level
		wantErr bool
		desc    string
	}{
		{in: slogx.LogLevelDebug, want: slog.LevelDebug, desc: "debug"},
		{in: slogx.LogLevelInfo, want: slog.LevelInfo, desc: "info"},
		{in: slogx.LogLevelWarn, want: slog.LevelWarn, desc: "warn"},
		{in: slogx.LogLevelError, want: slog.LevelError, desc: "error"},
		{in: slogx.LogLevelNone, want: slogx.LevelNone, desc: "none -> max int"},
		{in: slogx.LogLevel("invalid"), wantErr: true, desc: "invalid"},
	}

	for _, tc := range cases {
		lvl, err := tc.in.GetSlogLevel()
		if tc.wantErr {
			assert.Error(t, err, tc.desc)
			continue
		}
		require.NoError(t, err, tc.desc)
		assert.Equal(t, tc.want, lvl, tc.desc)
	}
}

func TestMustGetSlogLevel(t *testing.T) {
	// valid values should not panic and should return correct slog.Level
	valid := []struct {
		in   slogx.LogLevel
		want slog.Level
	}{
		{slogx.LogLevelDebug, slog.LevelDebug},
		{slogx.LogLevelInfo, slog.LevelInfo},
		{slogx.LogLevelWarn, slog.LevelWarn},
		{slogx.LogLevelError, slog.LevelError},
		{slogx.LogLevelNone, slog.Level(math.MaxInt)},
	}

	for _, tc := range valid {
		// ensure it does not panic
		assert.NotPanics(t, func() { _ = tc.in.MustGetSlogLevel() })
		got := tc.in.MustGetSlogLevel()
		assert.Equal(t, tc.want, got)
	}

	// invalid should panic
	invalid := slogx.LogLevel("invalid")
	require.Panics(t, func() { _ = invalid.MustGetSlogLevel() })
}

func TestParseLogFormat(t *testing.T) {
	cases := []struct {
		in      string
		want    slogx.LogFormat
		wantErr bool
		desc    string
	}{
		{in: "json", want: slogx.JSONFormat, desc: "json lowercase"},
		{in: "JSON", want: slogx.JSONFormat, desc: "json uppercase"},
		{in: "JsOn", want: slogx.JSONFormat, desc: "json mixedcase"},
		{in: "text", want: slogx.TextFormat, desc: "text lowercase"},
		{in: "TEXT", want: slogx.TextFormat, desc: "text uppercase"},
		{in: "TeXt", want: slogx.TextFormat, desc: "text mixedcase"},
		{in: "xml", wantErr: true, desc: "invalid"},
	}

	for _, tc := range cases {
		got, err := slogx.ParseLogFormat(tc.in)
		if tc.wantErr {
			assert.Error(t, err, tc.desc)
			continue
		}
		require.NoError(t, err, tc.desc)
		assert.Equal(t, tc.want, got, tc.desc)
	}
}

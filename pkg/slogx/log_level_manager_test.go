package slogx_test

import (
	"log/slog"
	"testing"

	"github.com/go-softwarelab/common/pkg/slogx"
	"github.com/go-softwarelab/common/pkg/to"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDynamicLogLevelHandlerGeneralLevel(t *testing.T) {
	t.Run("level by default is info", func(t *testing.T) {
		// given:
		loggerLevel := &slogx.LogLevelManager{}

		handler := loggerLevel.Decorate(slog.DiscardHandler)

		// then:
		assert.False(t, handler.Enabled(t.Context(), slog.LevelDebug), "Default value of level is info, so debug should be disabled")

		assert.True(t, handler.Enabled(t.Context(), slog.LevelInfo), "Default value of level is info, so info should be enabled")
	})

	t.Run("Handler reflects levelCalculator level on NewLogLevelManager if not Info", func(t *testing.T) {
		// given:
		loggerLevel := slogx.NewLogLevelManager(slog.LevelWarn)

		// and:
		handler := loggerLevel.Decorate(slog.DiscardHandler)

		// then:
		assert.False(t, handler.Enabled(t.Context(), slog.LevelInfo), "Info should be disabled at Warn level")
		assert.True(t, handler.Enabled(t.Context(), slog.LevelWarn), "Warn should be enabled at Warn level")
	})

	t.Run("Handler sees updated level in levelCalculator after instantiation", func(t *testing.T) {
		// given
		loggerLevel := slogx.NewLogLevelManager(slog.LevelInfo)

		// and:
		handler := loggerLevel.Decorate(slog.DiscardHandler)

		// when:
		loggerLevel.SetLevel(slog.LevelError)

		// then:
		assert.False(t, handler.Enabled(t.Context(), slog.LevelWarn), "Warn should be disabled at Error level after change")
		assert.True(t, handler.Enabled(t.Context(), slog.LevelError), "Error should be enabled at Error level after change")
	})

	t.Run("all levels are disabled on level none", func(t *testing.T) {
		// given:
		loggerLevel := slogx.NewLogLevelManager(slogx.LogLevelNone)

		handler := loggerLevel.Decorate(slog.DiscardHandler)

		// then:
		assert.False(t, handler.Enabled(t.Context(), slog.LevelDebug), "Debug is disabled on level none")

		assert.False(t, handler.Enabled(t.Context(), slog.LevelError), "Error is disabled on level none")
	})

	t.Run("all levels are disabled when levelCalculator changed to level none", func(t *testing.T) {
		// given:
		loggerLevel := slogx.NewLogLevelManager(slogx.LogLevelDebug)

		handler := loggerLevel.Decorate(slog.DiscardHandler)

		// when:
		loggerLevel.SetLogLevel(slogx.LogLevelNone)

		// then:
		assert.False(t, handler.Enabled(t.Context(), slog.LevelDebug), "Debug is disabled on level none")

		assert.False(t, handler.Enabled(t.Context(), slog.LevelError), "Error is disabled on level none")
	})

	t.Run("Handler sees updated level after calling Enabled", func(t *testing.T) {
		// given:
		loggerLevel := slogx.NewLogLevelManager(slog.LevelInfo)
		handler := loggerLevel.Decorate(slog.DiscardHandler)

		// when:
		assert.True(t, handler.Enabled(t.Context(), slog.LevelInfo))

		// and:
		loggerLevel.SetLevel(slog.LevelDebug)

		// then:
		assert.True(t, handler.Enabled(t.Context(), slog.LevelDebug), "Debug should be enabled after lowering level")
	})

	t.Run("handler create with WithAttrs should have the same level as parent handler", func(t *testing.T) {
		// given:
		loggerLevel := slogx.NewLogLevelManager(slog.LevelWarn)

		// and:
		base := loggerLevel.Decorate(slog.DiscardHandler)

		// when
		handlerWithAttrs := base.WithAttrs([]slog.Attr{slog.String(slogx.ServiceKey, "anything")})

		// then:
		assert.False(t, handlerWithAttrs.Enabled(t.Context(), slog.LevelInfo), "Info should be disabled at Warn level")
		assert.True(t, handlerWithAttrs.Enabled(t.Context(), slog.LevelWarn), "Warn should be enabled at Warn level")
	})

	t.Run("handler create with WithAttrs should reflect changes in level in levelCalculator", func(t *testing.T) {
		// given:
		loggerLevel := slogx.NewLogLevelManager(slog.LevelWarn)

		// and:
		base := loggerLevel.Decorate(slog.DiscardHandler)

		// when
		handlerWithAttrs := base.WithAttrs([]slog.Attr{slog.String(slogx.ServiceKey, "anything")})

		// and:
		loggerLevel.SetLevel(slog.LevelDebug)

		// then:
		assert.True(t, handlerWithAttrs.Enabled(t.Context(), slog.LevelDebug), "Debug should be enable after change of level")
	})

	t.Run("handler create with WithGroup should have the same level as parent handler", func(t *testing.T) {
		// given:
		loggerLevel := slogx.NewLogLevelManager(slog.LevelWarn)

		// and:
		base := loggerLevel.Decorate(slog.DiscardHandler)

		// when
		handlerWithGroup := base.WithGroup(slogx.ServiceKey)

		// then:
		assert.False(t, handlerWithGroup.Enabled(t.Context(), slog.LevelInfo), "Info should be disabled at Warn level")
		assert.True(t, handlerWithGroup.Enabled(t.Context(), slog.LevelWarn), "Warn should be enabled at Warn level")
	})

	t.Run("handler create with WithGroup should reflect changes in level in levelCalculator", func(t *testing.T) {
		// given:
		loggerLevel := slogx.NewLogLevelManager(slog.LevelWarn)

		// and:
		base := loggerLevel.Decorate(slog.DiscardHandler)

		// when
		handlerWithGroup := base.WithGroup(slogx.ServiceKey)

		// and:
		loggerLevel.SetLevel(slog.LevelDebug)

		// then:
		assert.True(t, handlerWithGroup.Enabled(t.Context(), slog.LevelDebug), "Debug should be enable after change of level")
	})

}

func TestDynamicLogLevelHandlerLevelByPattern(t *testing.T) {
	t.Run("pattern doesn't match handler without args", func(t *testing.T) {
		// given:
		loggerLevel := slogx.NewLogLevelManager(slogx.LogLevelNone)

		err := loggerLevel.SetLevelForServicePattern("SomeService", slog.LevelDebug)
		require.NoError(t, err, "invalid test setup: invalid pattern")

		// when:
		handler := loggerLevel.Decorate(slog.DiscardHandler)

		// then:
		assert.False(t, handler.Enabled(t.Context(), slog.LevelDebug), "main handler doesn't have match the pattern so is in level none")

	})

	patternMatchingTestCases := map[string]struct {
		pattern          string
		attrs            []slog.Attr
		levelShouldApply bool
	}{
		"match single service attr pattern when there is exactly one equal service attribute": {
			pattern:          "SomeService",
			attrs:            []slog.Attr{slogx.Service("SomeService")},
			levelShouldApply: true,
		},
		"match case sensitive service attr pattern when there is exactly one equal service attribute": {
			pattern:          "SomeService",
			attrs:            []slog.Attr{slogx.Service("SomeService")},
			levelShouldApply: true,
		},
		"match single service attr pattern when one of service attribute is equal": {
			pattern:          "SomeService",
			attrs:            []slog.Attr{slogx.Service("TotallyDifferent"), slogx.Service("SomeService"), slogx.Service("OtherService")},
			levelShouldApply: true,
		},
		"do not match service attr pattern when service attribute is in different casing": {
			pattern:          "SomeService",
			attrs:            []slog.Attr{slogx.Service("someService")},
			levelShouldApply: false,
		},
		"do not match single service attr pattern when service is totally different": {
			pattern:          "SomeService",
			attrs:            []slog.Attr{slogx.Service("TotallyDifferent")},
			levelShouldApply: false,
		},
		"do not match single service attr pattern when service only contains the name from pattern": {
			pattern:          "SomeService",
			attrs:            []slog.Attr{slogx.Service("SomeService1")},
			levelShouldApply: false,
		},
		"match multiple service attr pattern": {
			pattern:          "SomeService.SomeChildService",
			attrs:            []slog.Attr{slogx.Service("SomeService"), slogx.Service("SomeChildService")},
			levelShouldApply: true,
		},
		"match multiple service attr pattern when there are more service attributes": {
			pattern:          "SomeService.SomeChildService",
			attrs:            []slog.Attr{slogx.Service("TotallyDifferent"), slogx.Service("SomeService"), slogx.Service("SomeChildService"), slogx.Service("OtherService")},
			levelShouldApply: true,
		},
		"do not match multiple service attr pattern when there are less service attributes then required services": {
			pattern:          "SomeService.SomeChildService",
			attrs:            []slog.Attr{slogx.Service("SomeChildService")},
			levelShouldApply: false,
		},
		"do not match multiple service attr pattern when only first service is equal": {
			pattern:          "SomeService.SomeChildService",
			attrs:            []slog.Attr{slogx.Service("TotallyDifferent"), slogx.Service("SomeService"), slogx.Service("OtherService")},
			levelShouldApply: false,
		},
		"do not match multiple service attr pattern when one of required service is missing": {
			pattern:          "SomeService.SomeChildService.SomeGrandChildService",
			attrs:            []slog.Attr{slogx.Service("SomeService"), slogx.Service("SomeChildService")},
			levelShouldApply: false,
		},
	}
	for name, test := range patternMatchingTestCases {
		t.Run(name, func(t *testing.T) {
			// given:
			loggerLevel := slogx.NewLogLevelManager(slogx.LogLevelNone)

			err := loggerLevel.SetLevelForServicePattern(test.pattern, slog.LevelDebug)
			require.NoError(t, err, "invalid test setup: invalid pattern")

			// and:
			handler := loggerLevel.Decorate(slog.DiscardHandler)

			// when:
			handlerForService := handler.WithAttrs(test.attrs)

			// and:
			applied := handlerForService.Enabled(t.Context(), slog.LevelDebug)

			// then:
			assert.Equalf(t, test.levelShouldApply, applied, to.IfThen(test.levelShouldApply, "level should be applied").ElseThen("level should not be applied"))
		})
	}

	t.Run("matching pattern added after handler creation should impact the handler", func(t *testing.T) {
		// given:
		loggerLevel := slogx.NewLogLevelManager(slogx.LogLevelNone)

		// and:
		handler := loggerLevel.Decorate(slog.DiscardHandler)

		// and:
		handlerForService := handler.WithAttrs([]slog.Attr{slogx.Service("SomeService")})

		// when:
		err := loggerLevel.SetLevelForServicePattern("SomeService", slog.LevelDebug)
		require.NoError(t, err, "invalid test setup: invalid pattern")

		// then:
		assert.True(t, handlerForService.Enabled(t.Context(), slog.LevelDebug), "level should be applied")
	})

	patternsPrioritizationTests := map[string]struct {
		patterns      []string
		expectedLevel slog.Level
	}{
		"the most specific pattern at first place": {
			patterns: []string{
				"SomeService.ChildService",
				"SomeService",
			},
			expectedLevel: slog.LevelDebug,
		},
		"the most specific pattern at last place": {
			patterns: []string{
				"SomeService",
				"SomeService.ChildService",
				"SomeService.ChildService.GrandChildService",
			},
			expectedLevel: slog.LevelWarn,
		},
		"the most specific pattern set in the middle place": {
			patterns: []string{
				"SomeService",
				"SomeService.ChildService.GrandChildService",
				"SomeService.ChildService",
			},
			expectedLevel: slog.LevelInfo,
		},
	}
	var availableLevels = []slog.Level{slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError}
	for name, test := range patternsPrioritizationTests {
		t.Run(name, func(t *testing.T) {
			// given:
			loggerLevel := slogx.NewLogLevelManager(slogx.LogLevelNone)

			for i, pattern := range test.patterns {
				err := loggerLevel.SetLevelForServicePattern(pattern, availableLevels[i])
				require.NoError(t, err, "invalid test setup: invalid pattern")
			}

			// and:
			handler := loggerLevel.Decorate(slog.DiscardHandler)

			// when:
			handlerForService := handler.WithAttrs([]slog.Attr{
				slogx.Service("SomeService"),
				slogx.Service("ChildService"),
				slogx.Service("GrandChildService"),
			})

			// and:
			assertLevelEnabled(t, handlerForService, test.expectedLevel)

		})
	}

	patternsFromMapTests := map[string]struct {
		patterns       map[string]any
		expectedLevels []slog.Level
	}{
		"single pattern with slog.Level in map": {
			patterns: map[string]any{
				"SomeService": slog.LevelDebug,
			},
			expectedLevels: []slog.Level{slog.LevelDebug, slog.LevelDebug, slog.LevelDebug},
		},
		"single pattern with LogLevel in map": {
			patterns: map[string]any{
				"SomeService": slogx.LogLevelDebug,
			},
			expectedLevels: []slog.Level{slog.LevelDebug, slog.LevelDebug, slog.LevelDebug},
		},
		"single pattern with string log level in map": {
			patterns: map[string]any{
				"SomeService": "debug",
			},
			expectedLevels: []slog.Level{slog.LevelDebug, slog.LevelDebug, slog.LevelDebug},
		},
		"single pattern with int log level in map": {
			patterns: map[string]any{
				"SomeService": -4,
			},
			expectedLevels: []slog.Level{slog.LevelDebug, slog.LevelDebug, slog.LevelDebug},
		},
		"multiple patterns in map": {
			patterns: map[string]any{
				"SomeService":                                slog.LevelDebug,
				"SomeService.ChildService":                   slog.LevelInfo,
				"SomeService.ChildService.GrandChildService": slog.LevelWarn,
			},
			expectedLevels: []slog.Level{slog.LevelDebug, slog.LevelInfo, slog.LevelWarn},
		},
	}
	for name, test := range patternsFromMapTests {
		t.Run(name, func(t *testing.T) {
			// given:
			loggerLevel := slogx.NewLogLevelManager(slogx.LogLevelNone)

			// and:
			handler := loggerLevel.Decorate(slog.DiscardHandler)

			// and:
			childHandlers := []slog.Handler{
				handler.WithAttrs([]slog.Attr{slogx.Service("SomeService")}),
				handler.WithAttrs([]slog.Attr{slogx.Service("SomeService"), slogx.Service("ChildService")}),
				handler.WithAttrs([]slog.Attr{slogx.Service("SomeService"), slogx.Service("ChildService"), slogx.Service("GrandChildService")}),
			}

			// when:
			err := loggerLevel.SetLevels(test.patterns)

			// then:
			require.NoError(t, err)

			for i, level := range test.expectedLevels {
				assertLevelEnabled(t, childHandlers[i], level)
			}

		})
	}

	componentAttrSupportTests := map[string]struct {
		pattern string
		attrs   []slog.Attr
	}{
		"match single component attr": {
			pattern: "SomeComponent",
			attrs:   []slog.Attr{slogx.Component("SomeComponent")},
		},
		"match multiple component attr": {
			pattern: "SomeComponent.SomeChildComponent",
			attrs:   []slog.Attr{slogx.Component("SomeComponent"), slogx.Component("SomeChildComponent")},
		},
		"match single service and component attrs": {
			pattern: "SomeService.SomeComponent",
			attrs:   []slog.Attr{slogx.Service("SomeService"), slogx.Component("SomeComponent")},
		},
		"match multiple service and component attrs": {
			pattern: "SomeService.SomeComponent.SomeChildComponent",
			attrs:   []slog.Attr{slogx.Service("SomeService"), slogx.Component("SomeComponent"), slogx.Component("SomeChildComponent")},
		},
		"match multiple service and component attrs with different order": {
			pattern: "SomeService.SomeComponent.SomeChildComponent",
			attrs:   []slog.Attr{slogx.Component("SomeComponent"), slogx.Service("SomeService"), slogx.Component("SomeChildComponent")},
		},
	}
	for name, test := range componentAttrSupportTests {
		t.Run(name, func(t *testing.T) {
			// given:
			loggerLevel := slogx.NewLogLevelManager(slogx.LogLevelNone)

			err := loggerLevel.SetLevelForServicePattern(test.pattern, slog.LevelDebug)
			require.NoError(t, err, "invalid test setup: invalid pattern")

			// and:
			handler := loggerLevel.Decorate(slog.DiscardHandler)

			// when:
			handlerForService := handler.WithAttrs(test.attrs)

			// then:
			assertLevelEnabled(t, handlerForService, slog.LevelDebug)
		})
	}

	t.Run("support additional keys for attrs for dot pattern", func(t *testing.T) {
		// given:
		loggerLevel := slogx.NewLogLevelManager(slogx.LogLevelNone, slogx.WithAdditionalDotPatternAttrKeys("foo"))

		err := loggerLevel.SetLevelForServicePattern("SomeService.Bar", slog.LevelDebug)
		require.NoError(t, err, "invalid test setup: invalid pattern")

		// and:
		handler := loggerLevel.Decorate(slog.DiscardHandler)

		// when:
		handlerForService := handler.WithAttrs([]slog.Attr{slogx.Service("SomeService"), slog.String("foo", "Bar")})

		// then:
		assertLevelEnabled(t, handlerForService, slog.LevelDebug)
	})
}

func assertLevelEnabled(t *testing.T, handler slog.Handler, level slog.Level) {
	t.Helper()
	var availableLevels = []slog.Level{slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError}
	for _, l := range availableLevels {
		if l >= level {
			assert.True(t, handler.Enabled(t.Context(), l), "level %s should be enabled", l)
		} else {
			assert.False(t, handler.Enabled(t.Context(), l), "level %s should be disabled", l)
		}
	}
}

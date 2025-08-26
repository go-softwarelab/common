package slogx

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strings"
	"sync/atomic"

	"github.com/go-softwarelab/common/pkg/is"
	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
	slicesx "github.com/go-softwarelab/common/pkg/slices"
	"github.com/go-softwarelab/common/pkg/to"
)

var _ slog.Handler = (*managedLogLevelHandler)(nil)

//nolint:revive
type DynamicLogLevelOptions struct {
	serviceAttrKeys []string
}

// WithAdditionalDotPatternAttrKeys adds additional keys to the list of keys used for dynamic log level matching.
// By default, the keys are "service" and "component".
// The keys are matched against the logger attributes keys and the values are matched against the pattern.
func WithAdditionalDotPatternAttrKeys(keys ...string) func(*DynamicLogLevelOptions) {
	return func(options *DynamicLogLevelOptions) {
		options.serviceAttrKeys = append(options.serviceAttrKeys, keys...)
	}
}

// LogLevelManager is a logger decorator that allows to configure log levels dynamically based on logger attributes.
type LogLevelManager struct {
	version      atomic.Uint64
	defaultLevel slog.LevelVar
	dotMatchers  []*dotPatternMatcher
	options      DynamicLogLevelOptions
}

// NewLogLevelManager creates a new LogLevelManager.
func NewLogLevelManager[L slog.Level | LogLevel](level L, opts ...func(*DynamicLogLevelOptions)) *LogLevelManager {
	options := to.OptionsWithDefault(DynamicLogLevelOptions{
		serviceAttrKeys: []string{ServiceKey, ComponentKey},
	}, opts...)

	logLevel := &LogLevelManager{
		options: options,
	}

	switch l := any(level).(type) {
	case LogLevel:
		logLevel.SetLogLevel(l)
	case slog.Level:
		logLevel.SetLevel(l)
	default:
		panic(fmt.Errorf("unexpected type (%T) of level passed without compiler error", level))
	}

	return logLevel
}

// SetLogLevel sets the default logging level to the specified slogx.LogLevel.
func (m *LogLevelManager) SetLogLevel(level LogLevel) {
	m.defaultLevel.Set(level.MustGetSlogLevel())
}

// SetLevel sets the default logging level to the specified slog.Level.
func (m *LogLevelManager) SetLevel(level slog.Level) {
	m.defaultLevel.Set(level)
}

// SetLevelForServicePattern associates a logging level with a given simple dot-separated pattern for dynamic log level matching.
// Returns an error if the pattern cannot be parsed or the level cannot be set.
//
// Dot-Service-Pattern is a string containing dot-separated services (or components) names, such as "Service.Component".
// The pattern's dot-separated parts are matched against the logger attributes values whose key is "service" (or "component").
// The specified level is applied to all attributes that match the pattern.
// The most specific matching pattern determines the log level.
// If multiple patterns of equal specificity match, one is chosen arbitrarily.
//
// For example:
// Given the patterns: "Service1" and "Service1.Service2",
// and logger attributes: service="Service1", service="Service2", user=1
// the level set for the pattern "Service1.Service2" will be used.
func (m *LogLevelManager) SetLevelForServicePattern(pattern string, level slog.Level) error {
	matcher, err := dotPatternMatcherFromString(pattern, level)
	if err != nil {
		return fmt.Errorf("failed to set level for pattern: %w", err)
	}

	m.addMatcher(matcher)
	return nil
}

// SetLevels updates logging levels using a map of patterns and their corresponding levels; returns an error if invalid input.
// Currently only a dot-service-pattern is supported, see SetLevelForServicePattern for details.
// Value can be:
//   - string - string value parseable to slogx.LogLevel
//   - int: any integer value - although it's recommended to use slog.Level values
//   - slogx.LogLevel
//   - slog.Level
func (m *LogLevelManager) SetLevels(patterns map[string]any) error {
	dotPatterns := seq2.Map(seq2.FromMap(patterns), func(pattern string, levelForPattern any) (*dotPatternMatcher, error) {
		var level slog.Level
		switch l := levelForPattern.(type) {
		case string:
			logLevel, err := ParseLogLevel(l)
			if err != nil {
				return nil, fmt.Errorf("invalid string level %q for pattern %s: %w", l, pattern, err)
			}
			level = logLevel.MustGetSlogLevel()
		case int:
			level = slog.Level(l)
		case LogLevel:
			level = l.MustGetSlogLevel()
		case slog.Level:
			level = l
		default:
			return nil, fmt.Errorf("unexpected type (%T) of level passed for pattern %s", levelForPattern, pattern)
		}
		return dotPatternMatcherFromString(pattern, level)
	})

	var err error
	var matchers []*dotPatternMatcher
	for matcher, errCreateMatcher := range dotPatterns {
		if errCreateMatcher != nil {
			err = errors.Join(err, errCreateMatcher)
			continue
		}
		if matcher != nil {
			matchers = append(matchers, matcher)
		}
	}
	if err != nil {
		return err
	}
	m.addMatcher(matchers...)
	return nil
}

// Decorate decorates the given handler with the LogLevelManager.
func (m *LogLevelManager) Decorate(handler slog.Handler) slog.Handler {
	return newManagedLogLevelHandler(handler, m)
}

// DecorateHandler decorates the given handler with the LogLevelManager.
func (m *LogLevelManager) DecorateHandler(handler slog.Handler, _ *DecoratorOptions) slog.Handler {
	return newManagedLogLevelHandler(handler, m)
}

func (m *LogLevelManager) addMatcher(matchers ...*dotPatternMatcher) {
	m.dotMatchers = append(m.dotMatchers, matchers...)
	slices.SortFunc(m.dotMatchers, func(i, j *dotPatternMatcher) int {
		return j.Priority() - i.Priority()
	})
	m.version.Add(1)
}

func (m *LogLevelManager) getVersion() uint64 {
	return m.version.Load()
}

func (m *LogLevelManager) calculateLevel(attrs []slog.Attr) *slog.LevelVar {
	// optimization for dot matchers
	serviceOnlyAttrs := slicesx.Filter(attrs, func(attr slog.Attr) bool {
		return slices.Contains(m.options.serviceAttrKeys, attr.Key)
	})

	for _, matcher := range m.dotMatchers {
		if matcher.Match(serviceOnlyAttrs) {
			return matcher.level
		}
	}

	return &m.defaultLevel
}

type dotPatternMatcher struct {
	level    *slog.LevelVar
	services []string
}

func dotPatternMatcherFromString(pattern string, level slog.Level) (*dotPatternMatcher, error) {
	matcher, err := parseDotPattern(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pattern %q: %w", pattern, err)
	}
	matcher.level.Set(level)
	return matcher, nil
}

// Match returns true if the matcher matches the given attributes.
func (m *dotPatternMatcher) Match(attrs []slog.Attr) bool {
	filtered := seq.Filter(seq.FromSlice(attrs), func(attr slog.Attr) bool {
		// WARNING: there is prefiltering applied in calculateLevel,
		// 	so currently we will skip filtering by attr key here
		//  but if we would like to publish this struct, we should consider adding it,
		//  or add type switch in calculateLevel before prefiltering

		return attr.Value.Kind() == slog.KindString
	})

	values := seq.Map(filtered, func(attr slog.Attr) string {
		return attr.Value.String()
	})

	return seq.ContainsAll(values, m.services...)
}

// Priority returns the matcher's priority.'
func (m *dotPatternMatcher) Priority() int {
	return len(m.services)
}

func parseDotPattern(pattern string) (*dotPatternMatcher, error) {
	if is.BlankString(pattern) {
		return nil, fmt.Errorf("pattern must not be empty")
	}

	services := strings.Split(pattern, ".")

	return &dotPatternMatcher{
		services: services,
		level:    &slog.LevelVar{},
	}, nil
}

type managedLogLevelHandler struct {
	levelCalculator *LogLevelManager
	levelVersion    uint64
	h               slog.Handler
	level           *slog.LevelVar
	attrs           []slog.Attr
}

func newManagedLogLevelHandler(wrappedHandler slog.Handler, level *LogLevelManager) *managedLogLevelHandler {
	return &managedLogLevelHandler{
		h:               wrappedHandler,
		levelCalculator: level,
		level:           &level.defaultLevel,
	}
}

// Enabled returns true if the log level is enabled.
func (h *managedLogLevelHandler) Enabled(_ context.Context, l slog.Level) bool {
	if h.levelVersion < h.levelCalculator.getVersion() {
		h.levelVersion = h.levelCalculator.getVersion()
		h.level = h.levelCalculator.calculateLevel(h.attrs)
	}
	return l >= h.level.Level()
}

// WithAttrs returns a new handler with the given attributes.
func (h *managedLogLevelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	wh := h.h.WithAttrs(attrs)
	attrs = append(h.attrs[:], attrs...)
	levelVersion := h.levelCalculator.getVersion()
	level := h.levelCalculator.calculateLevel(attrs)

	return &managedLogLevelHandler{
		h:               wh,
		attrs:           attrs,
		levelCalculator: h.levelCalculator,
		level:           level,
		levelVersion:    levelVersion,
	}
}

// WithGroup returns a new handler with the given group name.
func (h *managedLogLevelHandler) WithGroup(name string) slog.Handler {
	return newManagedLogLevelHandler(h.h.WithGroup(name), h.levelCalculator)
}

// Handle handles the log record.
func (h *managedLogLevelHandler) Handle(ctx context.Context, record slog.Record) error {
	return h.h.Handle(ctx, record) //nolint:wrapcheck
}

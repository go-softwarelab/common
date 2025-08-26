package slogx

import (
	"strings"

	"github.com/go-softwarelab/common/pkg/seq"
)

// TestingTBOutput is an interface that represents a testing.TB instance required methods to be used by NewTestLogger and TestingTBWriter.
type TestingTBOutput interface {
	Helper()
	Log(args ...any)
}

// TestingTBWriter is a utility type that implements the io.Writer interface by writing log output to testing.TB.
// It is commonly used in test scenarios to redirect logs to the test's output via the provided testing.TB instance.
type TestingTBWriter struct {
	t TestingTBOutput
}

// NewTestingTBWriter creates a new TestingTBWriter that writes output to the provided testing.TB instance.
func NewTestingTBWriter(t TestingTBOutput) *TestingTBWriter {
	return &TestingTBWriter{t: t}
}

// Write writes the provided byte slice to the testing.TB instance.
func (w *TestingTBWriter) Write(p []byte) (n int, err error) {
	return w.WriteString(string(p))
}

// WriteString writes the provided string to the testing.TB instance.
func (w *TestingTBWriter) WriteString(s string) (n int, err error) {
	w.t.Helper()
	w.t.Log(s)
	return len(s), nil
}

// CollectingLogsWriter is a simple io.Writer implementation that writes to a string builder.
// It is useful for testing purposes - to check what was written by the logger.
type CollectingLogsWriter struct {
	strings.Builder
}

// NewCollectingLogsWriter creates a new CollectingLogsWriter.
func NewCollectingLogsWriter() *CollectingLogsWriter {
	return &CollectingLogsWriter{}
}

// Lines returns a slice of lines from the log output.
func (w *CollectingLogsWriter) Lines() []string {
	return seq.Collect(strings.Lines(w.String()))
}

// Clear the stored log output.
func (w *CollectingLogsWriter) Clear() {
	w.Builder.Reset()
}

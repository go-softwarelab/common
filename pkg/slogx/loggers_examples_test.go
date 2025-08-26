package slogx_test

import (
	"log/slog"

	"github.com/go-softwarelab/common/pkg/slogx"
	"github.com/go-softwarelab/common/pkg/testingx"
)

var skipTimeInLogOutputForExamplePurposes = slogx.WithFormat(slogx.TextWithoutTimeFormat)

func ExampleNewLogger() {
	logger := slogx.NewLogger(skipTimeInLogOutputForExamplePurposes)
	logger.Info("test")

	// Output:
	// level=INFO msg=test
}

func ExampleNewTestLogger() {
	// We're simulating a test here in example.
	t := &testingx.E{Verbose: true}

	logger := slogx.NewTestLogger(t, skipTimeInLogOutputForExamplePurposes)
	logger.Info("test")

	// Output:
	// level=INFO msg=test
}

// Example showing how to set a different log level using WithLevel.
func ExampleNewLogger_withLevel() {
	logger := slogx.NewLogger(skipTimeInLogOutputForExamplePurposes, slogx.WithLevel(slog.LevelWarn))

	logger.Info("info") // filtered out at WARN level
	logger.Warn("warn") // printed

	// Output:
	// level=WARN msg=warn
}

// Example using LogLevelManager obtained from NewLoggerWithManagedLevel.
// We rebuild a logger with the same manager to control the output format in this example.
func ExampleNewLoggerWithManagedLevel() {
	// Create a manager via helper and then build a logger with a stable format for the example output.
	_, levelManager := slogx.NewLoggerWithManagedLevel(slogx.LogLevelWarn)
	logger := slogx.NewLogger(skipTimeInLogOutputForExamplePurposes, slogx.WithDecorator(levelManager))

	// Create a service-scoped child logger.
	svc := slogx.Child(logger, "OrderService")

	// Before adding a pattern, debug is filtered out (default WARN).
	svc.Debug("debug before change")

	// Enable DEBUG for the specific service using a dot pattern.
	_ = levelManager.SetLevelForServicePattern("OrderService", slog.LevelDebug)

	// Now DEBUG is enabled for OrderService.
	svc.Debug("debug after change")

	// Output:
	// level=DEBUG msg="debug after change" service=OrderService
}

// Example using LogLevelManager with NewLogger and WithDecorator.
// The example sets a level for a pattern before logger creation, logs, then adds another pattern and logs again to show the change.
func ExampleNewLogger_withDecoratorAndManagedLevel() {
	// Prepare manager and set a pattern BEFORE creating the logger.
	levelManager := slogx.NewLogLevelManager(slogx.LogLevelWarn)
	_ = levelManager.SetLevelForServicePattern("OrderService", slog.LevelInfo)

	// Build the logger with the manager and a stable text format (without time).
	logger := slogx.NewLogger(skipTimeInLogOutputForExamplePurposes, slogx.WithDecorator(levelManager))

	// First, create a child logger for the service.
	orderSvc := slogx.Child(logger, "OrderService")

	// At INFO level for OrderService, debug is filtered, info is printed.
	orderSvc.Debug("dbg-1")
	orderSvc.Info("info-1")

	// Now add a more specific pattern AFTER logger creation.
	_ = levelManager.SetLevelForServicePattern("OrderService.Payment", slog.LevelDebug)

	// Create a more specific child to include the Payment sub-service.
	paymentSvc := slogx.Child(orderSvc, "Payment")

	// At DEBUG level for OrderService.Payment, both debug and info are printed.
	paymentSvc.Debug("dbg-2")
	paymentSvc.Info("info-2")

	// Output:
	// level=INFO msg=info-1 service=OrderService
	// level=DEBUG msg=dbg-2 service=OrderService service=Payment
	// level=INFO msg=info-2 service=OrderService service=Payment
}

func ExampleSilentLogger() {
	logger := slogx.SilentLogger()
	logger.Info("this message will be discarded")
	logger.Error("this error will also be discarded")

	// Output:
}

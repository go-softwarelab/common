// Package must provides a collection of utility functions that panic in error scenarios
// instead of returning errors, simplifying code in specific contexts.
//
// The goal of this package is to offer helper functions for situations where errors
// cannot be meaningfully handled at runtime, such as when errors would indicate
// programmer mistakes rather than external conditions. It's particularly useful
// in cases where errors are not expected because values have been pre-validated
// or when handling initialization code that should fail fast.
//
// These utilities are designed to reduce error-checking boilerplate and improve
// code readability in initialization paths, configuration loading, and other
// contexts where failures represent exceptional conditions that should halt execution.
package must

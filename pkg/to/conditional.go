package to

// IfElseCondition is a struct that provides a fluent API for conditional value mapping.
type IfElseCondition[T any] struct {
	isTrue   bool
	whenTrue func() T
}

// If creates a new IfElseCondition with the given condition and result provider.
func If[T any](condition bool, resultProvider func() T) *IfElseCondition[T] {
	return &IfElseCondition[T]{
		isTrue:   condition,
		whenTrue: resultProvider,
	}
}

// IfThen creates a new IfElseCondition with the given condition and result.
func IfThen[T any](condition bool, resultWhenTrue T) *IfElseCondition[T] {
	return &IfElseCondition[T]{
		isTrue: condition,
		whenTrue: func() T {
			return resultWhenTrue
		},
	}
}

// ElseIf adds an else if condition.
func (c *IfElseCondition[T]) ElseIf(condition bool, resultProvider func() T) *IfElseCondition[T] {
	if c.isTrue {
		return c
	}

	c.isTrue = condition
	c.whenTrue = resultProvider
	return c
}

// ElseIfThen adds an else if condition with a result.
func (c *IfElseCondition[T]) ElseIfThen(condition bool, resultWhenTrue T) *IfElseCondition[T] {
	if c.isTrue {
		return c
	}

	c.isTrue = condition
	c.whenTrue = func() T {
		return resultWhenTrue
	}
	return c
}

// Else accepts the default result provider and returns the result of the condition evaluation.
func (c *IfElseCondition[T]) Else(resultProvider func() T) T {
	if c.isTrue {
		return c.whenTrue()
	}
	return resultProvider()
}

// ElseThen accepts the default result and returns the result of the condition evaluation.
func (c *IfElseCondition[T]) ElseThen(resultWhenFalse T) T {
	if c.isTrue {
		return c.whenTrue()
	}
	return resultWhenFalse
}

// SwitchCase provides a fluent API for conditional (switch-like) value mapping. Represents case predicates.
type SwitchCase[V comparable, R any] interface {
	// Case adds a case that compares equality to case value.
	Case(value V) SwitchThen[V, R]
	// When adds a case predicate function.
	When(func(V) bool) SwitchThen[V, R]
	// Default adds a default value.
	Default(R) R
}

// SwitchThen provides a fluent API for conditional (switch-like) value mapping. Represents case results.
type SwitchThen[V comparable, R any] interface {
	// Then adds a result provider function for given case.
	Then(func(V) R) SwitchCase[V, R]
	// ThenValue adds a result value for given case.
	ThenValue(R) SwitchCase[V, R]
}

// Switch creates a new SwitchCase for the given value.
func Switch[V comparable, R any](value V) SwitchCase[V, R] {
	return &switchCase[V, R]{
		value: value,
	}
}

type switchCase[V comparable, R any] struct {
	value V
}

func (s *switchCase[V, R]) Case(predicate V) SwitchThen[V, R] {
	if s.value == predicate {
		return &matchedCase[V, R]{value: s.value}
	}
	// no match return self
	return s
}

func (s *switchCase[V, R]) When(predicate func(V) bool) SwitchThen[V, R] {
	if predicate(s.value) {
		return &matchedCase[V, R]{value: s.value}
	}
	// no match return self
	return s
}

func (s *switchCase[V, R]) ThenValue(_ R) SwitchCase[V, R] {
	// this wasn't the case, so we ignore it
	return s
}

func (s *switchCase[V, R]) Then(_ func(V) R) SwitchCase[V, R] {
	// this wasn't the case, so we ignore it
	return s
}

func (s *switchCase[V, R]) Default(defaultResult R) R {
	// no case matched, return default
	return defaultResult
}

type matchedCase[V comparable, R any] struct {
	value V
}

func (c *matchedCase[V, R]) Then(provider func(V) R) SwitchCase[V, R] {
	result := provider(c.value)
	return &resultOfSwitchCase[V, R]{
		result: result,
	}
}

func (m *matchedCase[V, R]) ThenValue(provider R) SwitchCase[V, R] {
	return &resultOfSwitchCase[V, R]{
		result: provider,
	}
}

type resultOfSwitchCase[V comparable, R any] struct {
	result R
}

func (r *resultOfSwitchCase[V, R]) Then(_ func(V) R) SwitchCase[V, R] {
	return r
}

func (r *resultOfSwitchCase[V, R]) ThenValue(_ R) SwitchCase[V, R] {
	return r
}

func (r *resultOfSwitchCase[V, R]) Case(_ V) SwitchThen[V, R] {
	return r
}

func (r *resultOfSwitchCase[V, R]) When(_ func(V) bool) SwitchThen[V, R] {
	return r
}

func (r *resultOfSwitchCase[V, R]) Default(_ R) R {
	return r.result
}

package types

// Pair is a generic type that represents a pair of values.
type Pair[L, R any] struct {
	Left  L
	Right R
}

// NewPair creates a new Pair with the given left and right values.
func NewPair[L, R any](left L, right R) *Pair[L, R] {
	return &Pair[L, R]{Left: left, Right: right}
}

// GetLeft returns the left value of the pair.
func (p *Pair[L, R]) GetLeft() L {
	return p.Left
}

// GetRight returns the right value of the pair.
func (p *Pair[L, R]) GetRight() R {
	return p.Right
}

// Unpack returns the left and right values of the pair.
func (p *Pair[L, R]) Unpack() (L, R) {
	return p.Left, p.Right
}

package types

import "iter"

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

// ToTuple converts the Pair to a Tuple.
func (p *Pair[L, R]) ToTuple() Tuple2[L, R] {
	return Tuple2[L, R]{A: p.Left, B: p.Right}
}

// Seq returns an iter.Seq with this Pair.
//
// This is useful for reusing functions provided by package seq.
func (p *Pair[L, R]) Seq() iter.Seq[Pair[L, R]] {
	return func(yield func(Pair[L, R]) bool) {
		yield(*p)
	}
}

// Seq2 returns an iter.Seq2 with left and right value.
//
// This is useful for reusing functions provided by package seq2.
func (p *Pair[L, R]) Seq2() iter.Seq2[L, R] {
	return func(yield func(L, R) bool) {
		yield(p.Left, p.Right)
	}
}

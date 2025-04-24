package types

// Tuple is a group of 2 elements
type Tuple[A, B any] = Tuple2[A, B]

// NewTuple creates a Tuple from given values.
func NewTuple[A, B any](a A, b B) Tuple[A, B] {
	return Tuple[A, B]{A: a, B: b}
}

// Tuple2 is a group of 2 elements.
type Tuple2[A, B any] struct {
	A A
	B B
}

// NewTuple2 creates a Tuple2 from given values.
func NewTuple2[A, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{A: a, B: b}
}

// GetLeft returns the left value of the tuple.
func (t Tuple2[A, B]) GetLeft() A {
	return t.A
}

// GetRight returns the right value of the tuple.
func (t Tuple2[A, B]) GetRight() B {
	return t.B
}

// Tuple3 is a group of 3 elements.
type Tuple3[A, B, C any] struct {
	A A
	B B
	C C
}

// NewTuple3 creates a Tuple3 from given values.
func NewTuple3[A, B, C any](a A, b B, c C) Tuple3[A, B, C] {
	return Tuple3[A, B, C]{A: a, B: b, C: c}
}

// Tuple4 is a group of 4 elements.
type Tuple4[A, B, C, D any] struct {
	A A
	B B
	C C
	D D
}

// NewTuple4 creates a Tuple4 from given values.
func NewTuple4[A, B, C, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return Tuple4[A, B, C, D]{A: a, B: b, C: c, D: d}
}

// Tuple5 is a group of 5 elements.
type Tuple5[A, B, C, D, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

// NewTuple5 creates a Tuple5 from given values.
func NewTuple5[A, B, C, D, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return Tuple5[A, B, C, D, E]{A: a, B: b, C: c, D: d, E: e}
}

// Tuple6 is a group of 6 elements.
type Tuple6[A, B, C, D, E, F any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
}

// NewTuple6 creates a Tuple6 from given values.
func NewTuple6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F] {
	return Tuple6[A, B, C, D, E, F]{A: a, B: b, C: c, D: d, E: e, F: f}
}

// Tuple7 is a group of 7 elements.
type Tuple7[A, B, C, D, E, F, G any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
}

// NewTuple7 creates a Tuple7 from given values.
func NewTuple7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G] {
	return Tuple7[A, B, C, D, E, F, G]{A: a, B: b, C: c, D: d, E: e, F: f, G: g}
}

// Tuple8 is a group of 8 elements.
type Tuple8[A, B, C, D, E, F, G, H any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
}

// NewTuple8 creates a Tuple8 from given values.
func NewTuple8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H] {
	return Tuple8[A, B, C, D, E, F, G, H]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h}
}

// Tuple9 is a group of 9 elements.
type Tuple9[A, B, C, D, E, F, G, H, I any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
}

// NewTuple9 creates a Tuple9 from given values.
func NewTuple9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I] {
	return Tuple9[A, B, C, D, E, F, G, H, I]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h, I: i}
}

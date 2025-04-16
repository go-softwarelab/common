# collection

```go
import "github.com/go-softwarelab/common/pkg/collection"
```



<a name="Tuple2"></a>
## type [Tuple2](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L4-L7>)

Tuple2 is a group of 2 elements \(pair\).

```go
type Tuple2[A, B any] struct {
    A   A
    B   B
}
```

<a name="NewTuple2"></a>
### [NewTuple2](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L10>)

```go
func NewTuple2[A, B any](a A, b B) Tuple2[A, B]
```

NewTuple2 creates a Tuple2 from given values.

<a name="Pair"></a>
### [Pair](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples.go#L4>)

```go
func Pair[A, B any](a A, b B) Tuple2[A, B]
```

Pair creates a pair tuple from given values.

<a name="Tuple3"></a>
## type [Tuple3](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L15-L19>)

Tuple3 is a group of 3 elements.

```go
type Tuple3[A, B, C any] struct {
    A   A
    B   B
    C   C
}
```

<a name="NewTuple3"></a>
### [NewTuple3](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L22>)

```go
func NewTuple3[A, B, C any](a A, b B, c C) Tuple3[A, B, C]
```

NewTuple3 creates a Tuple3 from given values.

<a name="Tuple4"></a>
## type [Tuple4](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L27-L32>)

Tuple4 is a group of 4 elements.

```go
type Tuple4[A, B, C, D any] struct {
    A   A
    B   B
    C   C
    D   D
}
```

<a name="NewTuple4"></a>
### [NewTuple4](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L35>)

```go
func NewTuple4[A, B, C, D any](a A, b B, c C, d D) Tuple4[A, B, C, D]
```

NewTuple4 creates a Tuple4 from given values.

<a name="Tuple5"></a>
## type [Tuple5](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L40-L46>)

Tuple5 is a group of 5 elements.

```go
type Tuple5[A, B, C, D, E any] struct {
    A   A
    B   B
    C   C
    D   D
    E   E
}
```

<a name="NewTuple5"></a>
### [NewTuple5](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L49>)

```go
func NewTuple5[A, B, C, D, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E]
```

NewTuple5 creates a Tuple5 from given values.

<a name="Tuple6"></a>
## type [Tuple6](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L54-L61>)

Tuple6 is a group of 6 elements.

```go
type Tuple6[A, B, C, D, E, F any] struct {
    A   A
    B   B
    C   C
    D   D
    E   E
    F   F
}
```

<a name="NewTuple6"></a>
### [NewTuple6](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L64>)

```go
func NewTuple6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F]
```

NewTuple6 creates a Tuple6 from given values.

<a name="Tuple7"></a>
## type [Tuple7](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L69-L77>)

Tuple7 is a group of 7 elements.

```go
type Tuple7[A, B, C, D, E, F, G any] struct {
    A   A
    B   B
    C   C
    D   D
    E   E
    F   F
    G   G
}
```

<a name="NewTuple7"></a>
### [NewTuple7](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L80>)

```go
func NewTuple7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G]
```

NewTuple7 creates a Tuple7 from given values.

<a name="Tuple8"></a>
## type [Tuple8](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L85-L94>)

Tuple8 is a group of 8 elements.

```go
type Tuple8[A, B, C, D, E, F, G, H any] struct {
    A   A
    B   B
    C   C
    D   D
    E   E
    F   F
    G   G
    H   H
}
```

<a name="NewTuple8"></a>
### [NewTuple8](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L97>)

```go
func NewTuple8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H]
```

NewTuple8 creates a Tuple8 from given values.

<a name="Tuple9"></a>
## type [Tuple9](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L102-L112>)

Tuple9 is a group of 9 elements.

```go
type Tuple9[A, B, C, D, E, F, G, H, I any] struct {
    A   A
    B   B
    C   C
    D   D
    E   E
    F   F
    G   G
    H   H
    I   I
}
```

<a name="NewTuple9"></a>
### [NewTuple9](<https://github.com/go-softwarelab/common/blob/main/pkg/collection/tuples_types.go#L115>)

```go
func NewTuple9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I]
```

NewTuple9 creates a Tuple9 from given values.
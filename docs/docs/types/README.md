# types

```go
import "github.com/go-softwarelab/common/pkg/types"
```

Package types defines a set of useful constraints to be used with type parameters.

Package types provides a collection of useful generic types and constraints for Go applications, enhancing type safety and enabling more expressive generic programming.

This package includes generic constraint types that define sets of types usable with type parameters, such as numeric types, ordered types, or comparable types.

It also includes utility types and structs like tuples or pairs, which simplify working with grouped data.

The \`types\` package is designed to complement Go's type parameter features, making it easier to write reusable and type\-safe code.



<a name="Comparable"></a>
## type [Comparable](<https://github.com/go-softwarelab/common/blob/main/pkg/types/constraints.go#L76>)

Comparable is an interface that is implemented by all comparable types \(booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types\). The comparable interface may only be used as a type parameter constraint, not as the type of a variable.

```go
type Comparable = comparable
```

<a name="Complex"></a>
## type [Complex](<https://github.com/go-softwarelab/common/blob/main/pkg/types/constraints.go#L56-L58>)

Complex is a constraint that permits any complex numeric type. If future releases of Go add new predeclared complex numeric types, this constraint will be modified to include them.

```go
type Complex interface {
    // contains filtered or unexported methods
}
```

<a name="Float"></a>
## type [Float](<https://github.com/go-softwarelab/common/blob/main/pkg/types/constraints.go#L35-L37>)

Float is a constraint that permits any floating\-point type. If future releases of Go add new predeclared floating\-point types, this constraint will be modified to include them.

```go
type Float interface {
    // contains filtered or unexported methods
}
```

<a name="Integer"></a>
## type [Integer](<https://github.com/go-softwarelab/common/blob/main/pkg/types/constraints.go#L28-L30>)

Integer is a constraint that permits any integer type. If future releases of Go add new predeclared integer types, this constraint will be modified to include them.

```go
type Integer interface {
    // contains filtered or unexported methods
}
```

<a name="Number"></a>
## type [Number](<https://github.com/go-softwarelab/common/blob/main/pkg/types/constraints.go#L42-L44>)

Number is a constraint that permits any numeric type. If future releases of Go add new predeclared numeric types, this constraint will be modified to include them.

```go
type Number interface {
    // contains filtered or unexported methods
}
```

<a name="Ordered"></a>
## type [Ordered](<https://github.com/go-softwarelab/common/blob/main/pkg/types/constraints.go#L69>)

Ordered is a constraint that permits any ordered type: any type that supports the operators \< \<= \>= \>. If future releases of Go add new ordered types, this constraint will be modified to include them.

Note that floating\-point types may contain NaN \("not\-a\-number"\) values. An operator such as == or \< will always report false when comparing a NaN value with any other value, NaN or not. See the \[Compare\] function for a consistent way to compare NaN values.

```go
type Ordered = cmp.Ordered
```

<a name="Pair"></a>
## type [Pair](<https://github.com/go-softwarelab/common/blob/main/pkg/types/pair.go#L6-L9>)

Pair is a generic type that represents a pair of values.

```go
type Pair[L, R any] struct {
    Left  L
    Right R
}
```

<a name="NewPair"></a>
### [NewPair](<https://github.com/go-softwarelab/common/blob/main/pkg/types/pair.go#L12>)

```go
func NewPair[L, R any](left L, right R) *Pair[L, R]
```

NewPair creates a new Pair with the given left and right values.

<a name="Pair[L, R].GetLeft"></a>
### [\*Pair\[L, R\].GetLeft](<https://github.com/go-softwarelab/common/blob/main/pkg/types/pair.go#L17>)

```go
func (p *Pair[L, R]) GetLeft() L
```

GetLeft returns the left value of the pair.

<a name="Pair[L, R].GetRight"></a>
### [\*Pair\[L, R\].GetRight](<https://github.com/go-softwarelab/common/blob/main/pkg/types/pair.go#L22>)

```go
func (p *Pair[L, R]) GetRight() R
```

GetRight returns the right value of the pair.

<a name="Pair[L, R].Seq"></a>
### [\*Pair\[L, R\].Seq](<https://github.com/go-softwarelab/common/blob/main/pkg/types/pair.go#L39>)

```go
func (p *Pair[L, R]) Seq() iter.Seq[Pair[L, R]]
```

Seq returns an iter.Seq with this Pair.

This is useful for reusing functions provided by package seq.

<a name="Pair[L, R].Seq2"></a>
### [\*Pair\[L, R\].Seq2](<https://github.com/go-softwarelab/common/blob/main/pkg/types/pair.go#L48>)

```go
func (p *Pair[L, R]) Seq2() iter.Seq2[L, R]
```

Seq2 returns an iter.Seq2 with left and right value.

This is useful for reusing functions provided by package seq2.

<a name="Pair[L, R].ToTuple"></a>
### [\*Pair\[L, R\].ToTuple](<https://github.com/go-softwarelab/common/blob/main/pkg/types/pair.go#L32>)

```go
func (p *Pair[L, R]) ToTuple() Tuple2[L, R]
```

ToTuple converts the Pair to a Tuple.

<a name="Pair[L, R].Unpack"></a>
### [\*Pair\[L, R\].Unpack](<https://github.com/go-softwarelab/common/blob/main/pkg/types/pair.go#L27>)

```go
func (p *Pair[L, R]) Unpack() (L, R)
```

Unpack returns the left and right values of the pair.

<a name="PairLike"></a>
## type [PairLike](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L6-L9>)

PairLike represents any type that is pair\-like, it is used for creating a Result instance.

```go
type PairLike[L, R any] interface {
    GetLeft() L
    GetRight() R
}
```

<a name="Result"></a>
## type [Result](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L12-L15>)

Result is a type representing a result that could be either a value or an error.

```go
type Result[V any] struct {
    // contains filtered or unexported fields
}
```

<a name="FailureResult"></a>
### [FailureResult](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L23>)

```go
func FailureResult[V any](err error) *Result[V]
```

FailureResult creates a new Result instance with the provided error.

<a name="ResultFrom"></a>
### [ResultFrom](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L36>)

```go
func ResultFrom[V any](provider func() (V, error)) *Result[V]
```

ResultFrom creates a Result instance from a function that returns a value and an error.

<a name="ResultFromPair"></a>
### [ResultFromPair](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L45>)

```go
func ResultFromPair[V any](pair PairLike[V, error]) *Result[V]
```

ResultFromPair creates a Result instance from a PairLike argument.

<a name="ResultOf"></a>
### [ResultOf](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L28>)

```go
func ResultOf[V any](value V, err error) *Result[V]
```

ResultOf creates a new Result instance with the provided value and error.

<a name="SuccessResult"></a>
### [SuccessResult](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L18>)

```go
func SuccessResult[V any](value V) *Result[V]
```

SuccessResult creates a new Result instance with the provided value.

<a name="Result[V].Get"></a>
### [\*Result\[V\].Get](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L63>)

```go
func (m *Result[V]) Get() (V, error)
```

Get returns the value and error from the Result instance.

<a name="Result[V].GetError"></a>
### [\*Result\[V\].GetError](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L77>)

```go
func (m *Result[V]) GetError() error
```

GetError returns the error from the Result instance. If there is no error, it returns nil.

<a name="Result[V].IsError"></a>
### [\*Result\[V\].IsError](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L53>)

```go
func (m *Result[V]) IsError() bool
```

IsError checks if the Result instance contains an error.

<a name="Result[V].IsNotError"></a>
### [\*Result\[V\].IsNotError](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L58>)

```go
func (m *Result[V]) IsNotError() bool
```

IsNotError checks if the Result instance does not contain an error.

<a name="Result[V].MustGetError"></a>
### [\*Result\[V\].MustGetError](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L82>)

```go
func (m *Result[V]) MustGetError() error
```

MustGetError returns the error from the Result instance, panicking if there is no error.

<a name="Result[V].MustGetValue"></a>
### [\*Result\[V\].MustGetValue](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L68>)

```go
func (m *Result[V]) MustGetValue() V
```

MustGetValue returns the value from the Result instance, panicking if there is an error.

<a name="Result[V].Or"></a>
### [\*Result\[V\].Or](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L106>)

```go
func (m *Result[V]) Or(alternative *Result[V]) *Result[V]
```

Or returns this Result if there is no error, otherwise it returns the provided alternative Result instance.

<a name="Result[V].OrElse"></a>
### [\*Result\[V\].OrElse](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L90>)

```go
func (m *Result[V]) OrElse(defaultValue V) V
```

OrElse returns the value if there is no error, otherwise it returns the provided default value.

<a name="Result[V].OrElseGet"></a>
### [\*Result\[V\].OrElseGet](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L98>)

```go
func (m *Result[V]) OrElseGet(defaultValue func() V) V
```

OrElseGet returns the value if there is no error, otherwise it returns the result of the provided function.

<a name="Result[V].Seq"></a>
### [\*Result\[V\].Seq](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L116>)

```go
func (m *Result[V]) Seq() iter.Seq[Result[V]]
```

Seq returns an iter.Seq with this Result.

This is useful for reusing functions provided by package seq.

<a name="Result[V].Seq2"></a>
### [\*Result\[V\].Seq2](<https://github.com/go-softwarelab/common/blob/main/pkg/types/result.go#L125>)

```go
func (m *Result[V]) Seq2() iter.Seq2[V, error]
```

Seq2 returns an iter.Seq2 with value and error.

This is useful for reusing functions provided by package seq2 or seqerr.

<a name="Signed"></a>
## type [Signed](<https://github.com/go-softwarelab/common/blob/main/pkg/types/constraints.go#L14-L16>)

Signed is a constraint that permits any signed integer type. If future releases of Go add new predeclared signed integer types, this constraint will be modified to include them.

```go
type Signed interface {
    // contains filtered or unexported methods
}
```

<a name="SignedNumber"></a>
## type [SignedNumber](<https://github.com/go-softwarelab/common/blob/main/pkg/types/constraints.go#L49-L51>)

SignedNumber is a constraint that permits any signed numeric type. If future releases of Go add new predeclared signed numeric types, this constraint will be modified to include them.

```go
type SignedNumber interface {
    // contains filtered or unexported methods
}
```

<a name="Tuple"></a>
## type [Tuple](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L4>)

Tuple is a group of 2 elements

```go
type Tuple[A, B any] = Tuple2[A, B]
```

<a name="NewTuple"></a>
### [NewTuple](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L7>)

```go
func NewTuple[A, B any](a A, b B) Tuple[A, B]
```

NewTuple creates a Tuple from given values.

<a name="Tuple2"></a>
## type [Tuple2](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L12-L15>)

Tuple2 is a group of 2 elements.

```go
type Tuple2[A, B any] struct {
    A   A
    B   B
}
```

<a name="NewTuple2"></a>
### [NewTuple2](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L18>)

```go
func NewTuple2[A, B any](a A, b B) Tuple2[A, B]
```

NewTuple2 creates a Tuple2 from given values.

<a name="Tuple2[A, B].GetLeft"></a>
### [Tuple2\[A, B\].GetLeft](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L23>)

```go
func (t Tuple2[A, B]) GetLeft() A
```

GetLeft returns the left value of the tuple.

<a name="Tuple2[A, B].GetRight"></a>
### [Tuple2\[A, B\].GetRight](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L28>)

```go
func (t Tuple2[A, B]) GetRight() B
```

GetRight returns the right value of the tuple.

<a name="Tuple3"></a>
## type [Tuple3](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L33-L37>)

Tuple3 is a group of 3 elements.

```go
type Tuple3[A, B, C any] struct {
    A   A
    B   B
    C   C
}
```

<a name="NewTuple3"></a>
### [NewTuple3](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L40>)

```go
func NewTuple3[A, B, C any](a A, b B, c C) Tuple3[A, B, C]
```

NewTuple3 creates a Tuple3 from given values.

<a name="Tuple4"></a>
## type [Tuple4](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L45-L50>)

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
### [NewTuple4](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L53>)

```go
func NewTuple4[A, B, C, D any](a A, b B, c C, d D) Tuple4[A, B, C, D]
```

NewTuple4 creates a Tuple4 from given values.

<a name="Tuple5"></a>
## type [Tuple5](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L58-L64>)

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
### [NewTuple5](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L67>)

```go
func NewTuple5[A, B, C, D, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E]
```

NewTuple5 creates a Tuple5 from given values.

<a name="Tuple6"></a>
## type [Tuple6](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L72-L79>)

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
### [NewTuple6](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L82>)

```go
func NewTuple6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F]
```

NewTuple6 creates a Tuple6 from given values.

<a name="Tuple7"></a>
## type [Tuple7](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L87-L95>)

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
### [NewTuple7](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L98>)

```go
func NewTuple7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G]
```

NewTuple7 creates a Tuple7 from given values.

<a name="Tuple8"></a>
## type [Tuple8](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L103-L112>)

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
### [NewTuple8](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L115>)

```go
func NewTuple8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H]
```

NewTuple8 creates a Tuple8 from given values.

<a name="Tuple9"></a>
## type [Tuple9](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L120-L130>)

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
### [NewTuple9](<https://github.com/go-softwarelab/common/blob/main/pkg/types/tuples.go#L133>)

```go
func NewTuple9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I]
```

NewTuple9 creates a Tuple9 from given values.

<a name="Unsigned"></a>
## type [Unsigned](<https://github.com/go-softwarelab/common/blob/main/pkg/types/constraints.go#L21-L23>)

Unsigned is a constraint that permits any unsigned integer type. If future releases of Go add new predeclared unsigned integer types, this constraint will be modified to include them.

```go
type Unsigned interface {
    // contains filtered or unexported methods
}
```
# optional

```go
import "github.com/go-softwarelab/common/pkg/optional"
```



<a name="Elem"></a>
## type Elem

Elem represents an optional value.

```go
type Elem[E any] struct {
    // contains filtered or unexported fields
}
```

<a name="Empty"></a>
### Empty

```go
func Empty[E any]() Elem[E]
```

Empty returns an empty optional value.

<a name="Of"></a>
### Of

```go
func Of[E any](v E) Elem[E]
```

Of returns an optional value with the given value.

<a name="Elem[E].IsEmpty"></a>
### Elem[E].IsEmpty

```go
func (o Elem[E]) IsEmpty() bool
```

IsEmpty returns true if the value is not present.

<a name="Elem[E].IsNotEmpty"></a>
### Elem[E].IsNotEmpty

```go
func (o Elem[E]) IsNotEmpty() bool
```

IsNotEmpty returns true if the value is present.

<a name="Elem[E].IsPresent"></a>
### Elem[E].IsPresent

```go
func (o Elem[E]) IsPresent() bool
```

IsPresent returns true if the value is present.

<a name="Elem[E].MustGet"></a>
### Elem[E].MustGet

```go
func (o Elem[E]) MustGet() E
```

MustGet returns the value if present, otherwise panics.

<a name="Elem[E].MustGetf"></a>
### Elem[E].MustGetf

```go
func (o Elem[E]) MustGetf(msg string, args ...any) E
```

MustGetf returns the value if present, otherwise panics with a custom message.
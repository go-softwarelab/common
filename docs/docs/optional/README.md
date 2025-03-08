# optional

```go
import "github.com/go-softwarelab/common/pkg/optional"
```



<a name="Elem"></a>
## type [Elem](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L8-L11>)



```go
type Elem[E any] struct {
    // contains filtered or unexported fields
}
```

<a name="Empty"></a>
### [Empty](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L37>)

```go
func Empty[E any]() Elem[E]
```



<a name="Of"></a>
### [Of](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L41>)

```go
func Of[E any](v E) Elem[E]
```



<a name="Elem[E].IsEmpty"></a>
### [Elem\[E\].IsEmpty](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L25>)

```go
func (o Elem[E]) IsEmpty() bool
```



<a name="Elem[E].IsNotEmpty"></a>
### [Elem\[E\].IsNotEmpty](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L33>)

```go
func (o Elem[E]) IsNotEmpty() bool
```



<a name="Elem[E].IsPresent"></a>
### [Elem\[E\].IsPresent](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L29>)

```go
func (o Elem[E]) IsPresent() bool
```



<a name="Elem[E].MustGet"></a>
### [Elem\[E\].MustGet](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L13>)

```go
func (o Elem[E]) MustGet() E
```



<a name="Elem[E].MustGetf"></a>
### [Elem\[E\].MustGetf](<https://github.com/go-softwarelab/common/blob/main/pkg/optional/optional.go#L17>)

```go
func (o Elem[E]) MustGetf(msg string, args ...any) E
```


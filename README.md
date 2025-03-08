## Common Utilities for Go

Welcome to the `common` library by `go-softwarelab`! This library is designed to simplify your Go development experience by providing a set of common tools and utilities. Our goal is to turn complexity into clarity, making your everyday coding tasks easier and more efficient.

One of the standout features of this library is its powerful set of functions that enhance the usability of Go iterators (`iter.Seq` and `iter.Seq2`). These functions are designed to make working with sequences in Go as seamless and pleasant as possible.

## Installation

To install the `common` library, simply run:

```sh
go get github.com/go-softwarelab/common
```

## Basic Example

Here's a basic example demonstrating how to use some of the key features of the `common` library:

```go
package main

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/seq"
)

func main() {
	// Create a sequence of integers
	numbers := seq.Of(1, 2, 3, 4, 5, 6)

	// Filter the sequence to include only even numbers
	evenNumbers := seq.Filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	// Map the filtered sequence to their squares
	squaredEvenNumbers := seq.Map(evenNumbers, func(n int) int {
		return n * n
	})

	// Print the result
	seq.ForEach(squaredEvenNumbers, func(i int) {
        fmt.Println(i)
	})
}
```

## Documentation

For more detailed documentation and examples, please visit [Go Common Utils](https://go-softwarelab.github.io/common) page.

We hope you find the `common` library useful and that it helps you write Go code that's as clear as a summer sky!

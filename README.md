test helpers
====

# Packages

This module contains test helpers for one-parameter-input-two-parameter-return style functions (e.g. `json.Marshal`, `os.Open`, etc.).

Package `must` uses `t.Fatal` to report any errors; package `should` uses `t.Error` to report any errors.

# Example

```go
package test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/iwittkau/th/must"
	"github.com/iwittkau/th/should"
)

func TestExamples(t *testing.T) {
	data := must.Return[any](t, []string{"test"}, json.Marshal) // type of input must be passed; return type []byte automatically inferred
	t.Log(string(data))

	f := should.Return(t, "must.go", os.Open) // returns an *os.File (automatically inferred)
	defer f.Close()
}
```

# Installation

This module is based on Go's generics support which was introduced in Go 1.18, therefor this module requires Go 1.18.

```
go get github.com/iwittkau/th
```

# Inspiration

The simple API design and the idea to separate the helpers into two packages was inspired by [testify's](https://github.com/stretchr/testify) `assert` and `require` packages.

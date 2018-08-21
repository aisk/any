# any.Type

[![GoDoc](https://godoc.org/github.com/aisk/any?status.svg)](https://godoc.org/github.com/aisk/any) [![Build Status](https://travis-ci.org/aisk/any.svg?branch=master)](https://travis-ci.org/aisk/any) [![Codecov](https://img.shields.io/codecov/c/github/aisk/any.svg)](https://codecov.io/gh/aisk/any)

The missing `any.Type` for go.

![any](http://d.justpo.st/media/images/2017/12/30/she-cant-love-you-for-your-money-if-you-dont-have-any-meme-1514663282.jpg)

---

`any.Type` is just a type alias for `interface{}`, you can use it as any value's type.

Example:

```go
package main

import (
    "github.com/aisk/any"
)

func main() {
    var a any.Type
    a = 42
    a = true
    a = "don't panic"
    println(a)
}
```

There is also an `any.Map` and `any.Slice`, behave like it looks like.

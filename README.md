[![Go Report Card](https://goreportcard.com/badge/github.com/dkharms/sliceutils)](https://goreportcard.com/report/github.com/dkharms/sliceutils)
[![CI](https://github.com/dkharms/sliceutils/actions/workflows/main-workflow.yml/badge.svg?branch=main)](https://github.com/dkharms/sliceutils/actions/workflows/main-workflow.yml)

## sliceutils

Some generic stuff for slices.

Here you can find frequently used methods for slices.

### Available methods

```go
func Unique[T comparable](s ...T) []T
```
Unique returns unique elements from slice.


```go
func Filter[T any](f func(T) bool, s ...T) []T
```
Filter returns values from slice which satisfies given predicate.

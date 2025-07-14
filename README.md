# 🛠️Ptr Utils

[![Go](https://github.com/manuelarte/ptrutils/actions/workflows/go.yml/badge.svg)](https://github.com/manuelarte/ptrutils/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/manuelarte/ptrutils)](https://goreportcard.com/report/github.com/manuelarte/ptrutils)
![version](https://img.shields.io/github/v/release/manuelarte/ptrutils)

Some utility functions to use with pointers in Go.

## ⬇️  Getting Started

To use it run:

```bash
go get github.com/manuelarte/ptrutils@latest
```

## 🚀 Features

### Ptr

Simple function that converts a value to a pointer. Handy if you want to inline pointers' initialization.

```go
annoyingWay := "myAnnoyingWay"
myStruct := MyStruct {
  AnnoyingWay: annoyingWay,
  LibWay: ptrutils.Ptr("myLibWay")
}
```

### DefOr

Simple function that either dereferences a pointer or, if `nil`, returns a default value.

```go
var p *string
ptrutils.DefOr(p, "default")
```

### DefOrL

Simple function that either dereferences a pointer or, if `nil`, it runs the function that returns a default value.

```go
var p *string
ptrutils.DefOrL(p, ptrutils.LazyIdentity("default"))
```

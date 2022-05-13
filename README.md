# go-js

[![push](https://github.com/frantjc/go-js/actions/workflows/ci.yml/badge.svg?branch=main&event=push)](https://github.com/frantjc/go-js/actions)

[Golang](https://go.dev) module for functions that mimic useful JavaScript functions using [Go 1.18's Generics](https://go.dev/blog/intro-generics)

## Install

```sh
go get github.com/frantjc/go-js
```

## Usage

See [examples](examples/).

```go
	array := []int{1,2,3,4}
	mappable := js.MappableArray[int, string](array)
	some := mappable.Map(func(a, _ int, _ []int) string {
		return fmt.Sprint(a)
	}).Some(func (b string, _ int, _ []string) bool {
		return b == "1"
	})

	fmt.Println(some)
	// true
```

package main

import (
	"fmt"

	"github.com/frantjc/go-js"
)

func main() {
	array := []int{1, 2, 3, 4}
	mappable := js.MappableArray[int, string](array)
	some := mappable.Map(func(a, _ int, _ []int) string {
		return fmt.Sprint(a)
	}).Some(func(b string, _ int, _ []string) bool {
		return b == "1"
	})

	fmt.Println(some)
	// true
}

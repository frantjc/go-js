package js_test

import (
	"testing"

	"github.com/frantjc/go-js"
	"github.com/stretchr/testify/assert"
)

func TestReducibleReduceRight(t *testing.T) {
	var (
		a = js.ReducibleArray[int, int]([]int{1, 2, 3, 4})
		f = func(acc, a, i int, s []int) int {
			return acc + a
		}
		expected = 10
		actual   = a.ReduceRight(f, 0)
	)
	assert.Equal(t, expected, actual)
}

func TestReducibleReduce(t *testing.T) {
	var (
		a = js.ReducibleArray[int, int]([]int{1, 2, 3, 4})
		f = func(acc, a, i int, s []int) int {
			return acc + a
		}
		expected = 10
		actual   = a.Reduce(f, 0)
	)
	assert.Equal(t, expected, actual)
}

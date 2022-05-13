package js_test

import (
	"testing"

	"github.com/frantjc/go-js"
	"github.com/stretchr/testify/assert"
)

func TestEveryTrue(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4}
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = true
		actual   = js.Every(a, f)
	)
	assert.Equal(t, expected, actual)
}

func TestEveryFalse(t *testing.T) {
	var (
		a = []int{0, 1, 2, 3}
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = false
		actual   = js.Every(a, f)
	)
	assert.Equal(t, expected, actual)
}

func TestFilter(t *testing.T) {
	var (
		a = []int{0, 1, 2, 3}
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = []int{1, 2, 3}
		actual   = js.Filter(a, f)
	)
	assert.Equal(t, expected, actual)
}

func TestFindIndex(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4}
		f = func(a, _ int, _ []int) bool {
			return a == 2
		}
		expected = 1
		actual   = js.FindIndex(a, f)
	)
	assert.Equal(t, expected, actual)
}

func TestFind(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4}
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = 1
		actual   = js.Find(a, f)
	)
	assert.Equal(t, expected, actual)
}

func TestIncludesTrue(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4}
		expected = true
		actual   = js.Includes(a, 1)
	)
	assert.Equal(t, expected, actual)
}

func TestIncludesFalse(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4}
		expected = false
		actual   = js.Includes(a, 0)
	)
	assert.Equal(t, expected, actual)
}

func TestIndexOf(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4}
		expected = 0
		actual   = js.IndexOf(a, 1)
	)
	assert.Equal(t, expected, actual)
}

func TestLastIndexOf(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4, 3, 2}
		expected = 4
		actual   = js.LastIndexOf(a, 3)
	)
	assert.Equal(t, expected, actual)
}

func TestMap(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4, 3, 2}
		f = func(a, _ int, _ []int) int {
			return a + 1
		}
		expected = []int{2, 3, 4, 5, 4, 3}
		actual   = js.Map(a, f)
	)
	assert.Equal(t, expected, actual)
}

func TestReduceRight(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4, 3, 2}
		f = func(acc, a, _ int, _ []int) int {
			return acc + a
		}
		expected = 15
		actual   = js.ReduceRight(a, f, 0)
	)
	assert.Equal(t, expected, actual)
}

func TestReduce(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4, 3, 2}
		f = func(acc, a, _ int, _ []int) int {
			return acc + a
		}
		expected = 15
		actual   = js.Reduce(a, f, 0)
	)
	assert.Equal(t, expected, actual)
}

func TestReverse(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4, 3, 2}
		expected = []int{2, 3, 4, 3, 2, 1}
		actual   = js.Reverse(a)
	)
	assert.Equal(t, expected, actual)
}

func TestSliceFromStart(t *testing.T) {
	var (
		a        = []int{0, 1, 2, 3}
		expected = []int{0, 1}
		actual   = js.Slice(a, 0, 2)
	)
	assert.Equal(t, expected, actual)
}

func TestSliceFromEnd(t *testing.T) {
	var (
		a        = []int{0, 1, 2, 3}
		expected = []int{0, 1, 2}
		actual   = js.Slice(a, 0, -1)
	)
	assert.Equal(t, expected, actual)
}

func TestSomeTrue(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4}
		f = func(a, _ int, _ []int) bool {
			return a == 4
		}
		expected = true
		actual   = js.Some(a, f)
	)
	assert.Equal(t, expected, actual)
}

func TestSomeFalse(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4}
		f = func(a, _ int, _ []int) bool {
			return a == 5
		}
		expected = false
		actual   = js.Some(a, f)
	)
	assert.Equal(t, expected, actual)
}

package js_test

import (
	"testing"

	"github.com/frantjc/go-js"
	"github.com/stretchr/testify/assert"
)

func TestAnyEveryTrue(t *testing.T) {
	var (
		a = js.NewAnyArray([]int{1, 2, 3, 4})
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = true
		actual   = a.Every(f)
	)
	assert.Equal(t, expected, actual)
}

func TestAnyEveryFalse(t *testing.T) {
	var (
		a = js.NewAnyArray([]int{0, 1, 2, 3})
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = false
		actual   = a.Every(f)
	)
	assert.Equal(t, expected, actual)
}

func TestAnyFilter(t *testing.T) {
	var (
		a = js.NewAnyArray([]int{0, 1, 2, 3})
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = js.NewAnyArray([]int{1, 2, 3})
		actual   = a.Filter(f)
	)
	assert.Equal(t, expected, actual)
}

func TestAnyFindIndex(t *testing.T) {
	var (
		a = js.NewAnyArray([]int{1, 2, 3, 4})
		f = func(a, _ int, _ []int) bool {
			return a == 2
		}
		expected = 1
		actual   = a.FindIndex(f)
	)
	assert.Equal(t, expected, actual)
}

func TestAnyFind(t *testing.T) {
	var (
		a = js.NewAnyArray([]int{1, 2, 3, 4})
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = 1
		actual   = a.Find(f)
	)
	assert.Equal(t, expected, actual)
}

func TestAnyMap(t *testing.T) {
	var (
		a = js.NewAnyArray([]int{1, 2, 3, 4, 3, 2})
		f = func(a, _ int, _ []int) any {
			return a + 1
		}
		expected = js.NewAnyArray([]interface{}{2, 3, 4, 5, 4, 3})
		actual   = a.Map(f)
	)
	assert.Equal(t, expected, actual)
}

// func TestAnyReduceRight(t *testing.T) {
// 	var (
// 		a = js.NewAnyArray([]int{1, 2, 3, 4, 3, 2})
// 		f = func(acc, a, _ int, _ []int) int {
// 			return acc + a
// 		}
// 		expected = 15
// 		actual   = a.ReduceRight(f, 0)
// 	)
// 	assert.Equal(t, expected, actual)
// }

// func TestAnyReduce(t *testing.T) {
// 	var (
// 		a = js.NewAnyArray([]int{1, 2, 3, 4, 3, 2})
// 		f = func(acc, a, _ int, _ []int) int {
// 			return acc + a
// 		}
// 		expected = 15
// 		actual   = a.Reduce(f, 0)
// 	)
// 	assert.Equal(t, expected, actual)
// }

func TestAnyReverse(t *testing.T) {
	var (
		a        = js.NewAnyArray([]int{1, 2, 3, 4, 3, 2})
		expected = js.NewAnyArray([]int{2, 3, 4, 3, 2, 1})
		actual   = a.Reverse()
	)
	assert.Equal(t, expected, actual)
}

func TestAnySliceFromStart(t *testing.T) {
	var (
		a        = js.NewAnyArray([]int{0, 1, 2, 3})
		expected = js.NewAnyArray([]int{0, 1})
		actual   = a.Slice(0, 2)
	)
	assert.Equal(t, expected, actual)
}

func TestAnySliceFromEnd(t *testing.T) {
	var (
		a        = js.NewAnyArray([]int{0, 1, 2, 3})
		expected = js.NewAnyArray([]int{2})
		actual   = a.Slice(-2, -1)
	)
	assert.Equal(t, expected, actual)
}

func TestAnySomeTrue(t *testing.T) {
	var (
		a = js.NewAnyArray([]int{1, 2, 3, 4})
		f = func(a, _ int, _ []int) bool {
			return a == 4
		}
		expected = true
		actual   = a.Some(f)
	)
	assert.Equal(t, expected, actual)
}

func TestAnySomeFalse(t *testing.T) {
	var (
		a = js.NewAnyArray([]int{1, 2, 3, 4})
		f = func(a, _ int, _ []int) bool {
			return a == 5
		}
		expected = false
		actual   = a.Some(f)
	)
	assert.Equal(t, expected, actual)
}

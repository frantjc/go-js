package js_test

import (
	"testing"

	"github.com/frantjc/go-js"
	"github.com/stretchr/testify/assert"
)

func TestComparableEveryTrue(t *testing.T) {
	var (
		a = js.NewComparableArray([]int{1, 2, 3, 4})
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = true
		actual   = a.Every(f)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableEveryFalse(t *testing.T) {
	var (
		a = js.NewComparableArray([]int{0, 1, 2, 3})
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = false
		actual   = a.Every(f)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableFilter(t *testing.T) {
	var (
		a = js.NewComparableArray([]int{0, 1, 2, 3})
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = js.NewComparableArray([]int{1, 2, 3})
		actual   = a.Filter(f)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableFindIndex(t *testing.T) {
	var (
		a = js.NewComparableArray([]int{1, 2, 3, 4})
		f = func(a, _ int, _ []int) bool {
			return a == 2
		}
		expected = 1
		actual   = a.FindIndex(f)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableFind(t *testing.T) {
	var (
		a = js.NewComparableArray([]int{1, 2, 3, 4})
		f = func(a, _ int, _ []int) bool {
			return a > 0
		}
		expected = 1
		actual   = a.Find(f)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableIncludesTrue(t *testing.T) {
	var (
		a        = js.NewComparableArray([]int{1, 2, 3, 4})
		expected = true
		actual   = a.Includes(1)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableIncludesFalse(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4}
		expected = false
		actual   = js.Includes(a, 0)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableIndexOf(t *testing.T) {
	var (
		a        = js.NewComparableArray([]int{1, 2, 3, 4})
		expected = 0
		actual   = a.IndexOf(1)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableLastIndexOf(t *testing.T) {
	var (
		a        = js.NewComparableArray([]int{1, 2, 3, 4, 3, 2})
		expected = 4
		actual   = a.LastIndexOf(3, 0)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableLastIndexOfFrom(t *testing.T) {
	var (
		a        = js.NewComparableArray([]int{1, 2, 3, 4, 3, 2})
		expected = 2
		actual   = a.LastIndexOf(3, 3)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableMap(t *testing.T) {
	var (
		a = js.NewComparableArray([]int{1, 2, 3, 4, 3, 2})
		f = func(a, _ int, _ []int) any {
			return a + 1
		}
		expected = js.NewAnyArray([]interface{}{2, 3, 4, 5, 4, 3})
		actual   = a.Map(f)
	)
	assert.Equal(t, expected, actual)
}

// func TestComparableReduceRight(t *testing.T) {
// 	var (
// 		a = js.NewComparableArray([]int{1, 2, 3, 4, 3, 2})
// 		f = func(acc, a, _ int, _ []int) int {
// 			return acc + a
// 		}
// 		expected = 15
// 		actual   = a.ReduceRight(f, 0)
// 	)
// 	assert.Equal(t, expected, actual)
// }

// func TestComparableReduce(t *testing.T) {
// 	var (
// 		a = js.NewComparableArray([]int{1, 2, 3, 4, 3, 2})
// 		f = func(acc, a, _ int, _ []int) int {
// 			return acc + a
// 		}
// 		expected = 15
// 		actual   = a.Reduce(f, 0)
// 	)
// 	assert.Equal(t, expected, actual)
// }

func TestComparableReverse(t *testing.T) {
	var (
		a        = js.NewComparableArray([]int{1, 2, 3, 4, 3, 2})
		expected = js.NewComparableArray([]int{2, 3, 4, 3, 2, 1})
		actual   = a.Reverse()
	)
	assert.Equal(t, expected, actual)
}

func TestComparableSliceFromStart(t *testing.T) {
	var (
		a        = js.NewComparableArray([]int{0, 1, 2, 3})
		expected = js.NewComparableArray([]int{0, 1})
		actual   = a.Slice(0, 2)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableSliceFromEnd(t *testing.T) {
	var (
		a        = js.NewComparableArray([]int{0, 1, 2, 3})
		expected = js.NewComparableArray([]int{2})
		actual   = a.Slice(-2, -1)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableSomeTrue(t *testing.T) {
	var (
		a = js.NewComparableArray([]int{1, 2, 3, 4})
		f = func(a, _ int, _ []int) bool {
			return a == 4
		}
		expected = true
		actual   = a.Some(f)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableSomeFalse(t *testing.T) {
	var (
		a = js.NewComparableArray([]int{1, 2, 3, 4})
		f = func(a, _ int, _ []int) bool {
			return a == 5
		}
		expected = false
		actual   = a.Some(f)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableUnique(t *testing.T) {
	var (
		a        = js.NewComparableArray([]int{1, 2, 3, 4, 3, 2, 1})
		expected = js.NewComparableArray([]int{1, 2, 3, 4})
		actual   = a.Unique()
	)
	assert.Equal(t, expected, actual)
}

func TestComparableSortWorst(t *testing.T) {
	var (
		a        = js.NewComparableArray([]int{4, 3, 2, 1})
		expected = js.NewComparableArray([]int{1, 2, 3, 4})
		actual   = a.Sort(func(a, b int) int {
			return a - b
		})
	)
	assert.Equal(t, expected, actual)
}

func TestComparableSortBest(t *testing.T) {
	var (
		a        = js.NewComparableArray([]int{1, 2, 3, 4})
		expected = js.NewComparableArray([]int{1, 2, 3, 4})
		actual   = a.Sort(func(a, b int) int {
			return a - b
		})
	)
	assert.Equal(t, expected, actual)
}

package js_test

import (
	"testing"

	"github.com/frantjc/go-js"
	"github.com/stretchr/testify/assert"
)

func TestComparableIncludesTrue(t *testing.T) {
	var (
		a        = js.ComparableArray[int]([]int{1, 2, 3, 4})
		expected = true
		actual   = a.Includes(1)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableIncludesFalse(t *testing.T) {
	var (
		a        = js.ComparableArray[int]([]int{1, 2, 3, 4})
		expected = false
		actual   = a.Includes(0)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableIndexOf(t *testing.T) {
	var (
		a        = js.ComparableArray[int]([]int{1, 2, 3, 4})
		expected = 0
		actual   = a.IndexOf(1)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableLastIndexOf(t *testing.T) {
	var (
		a        = js.ComparableArray[int]([]int{1, 2, 3, 4, 3, 2})
		expected = 4
		actual   = a.LastIndexOf(3, 0)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableLastIndexOfFrom(t *testing.T) {
	var (
		a        = js.ComparableArray[int]([]int{1, 2, 3, 4, 3, 2})
		expected = 2
		actual   = a.LastIndexOf(3, 3)
	)
	assert.Equal(t, expected, actual)
}

func TestComparableUnique(t *testing.T) {
	var (
		a        = js.ComparableArray[int]([]int{1, 2, 3, 4, 3, 2, 1})
		expected = js.ComparableArray[int]([]int{1, 2, 3, 4})
		actual   = a.Unique()
	)
	assert.Equal(t, expected, actual)
}

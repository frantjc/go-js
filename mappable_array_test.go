package js_test

import (
	"testing"

	"github.com/frantjc/go-js"
	"github.com/stretchr/testify/assert"
)

func TestMappableMap(t *testing.T) {
	var (
		a = js.MappableArray[int, int]([]int{1, 2, 3, 4})
		f = func(a, i int, s []int) int {
			return a + 1
		}
		expected = js.AnyArray[int]([]int{2, 3, 4, 5})
		actual   = a.Map(f)
	)
	assert.Equal(t, expected, actual)
}

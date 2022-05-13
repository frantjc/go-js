package js_test

import (
	"testing"

	"github.com/frantjc/go-js"
	"github.com/stretchr/testify/assert"
)

func TestTernaryTrue(t *testing.T) {
	var (
		expected = 1
		actual   = js.Ternary(true, 1, 0)
	)
	assert.Equal(t, expected, actual)
}

func TestTernaryFalse(t *testing.T) {
	var (
		expected = 0
		actual   = js.Ternary(false, 1, 0)
	)
	assert.Equal(t, expected, actual)
}

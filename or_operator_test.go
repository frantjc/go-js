package js_test

import (
	"testing"

	"github.com/frantjc/go-js"
	"github.com/stretchr/testify/assert"
)

func TestCoalesceTrue(t *testing.T) {
	var (
		expected = "default"
		actual   = js.Coalesce("", "", "", expected, "")
	)
	assert.Equal(t, expected, actual)
}

func TestCoalesceFalse(t *testing.T) {
	var (
		expected = ""
		actual   = js.Coalesce("", "", "")
	)
	assert.Equal(t, expected, actual)
}

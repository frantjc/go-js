package js_test

import (
	"net/url"
	"testing"

	"github.com/frantjc/go-js"
	"github.com/stretchr/testify/assert"
)

var (
	localhost8080 *url.URL
	localhost3000 *url.URL
)

func init() {
	localhost8080, _ = url.Parse("http://localhost:8080/")
	localhost3000, _ = url.Parse("http://localhost:3000/")
}

func TestStringerJoin(t *testing.T) {
	var (
		a         = js.NewStringerArray([]*url.URL{localhost8080, localhost3000})
		separator = ","
		expected  = localhost8080.String() + separator + localhost3000.String()
		actual    = a.Join(separator)
	)
	assert.Equal(t, expected, actual)
}

package util

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSHA1(t *testing.T) {
	var cases = []struct {
		in  string
		out string
	}{
		{in: "tom", out: "96835dd8bfa718bd6447ccc87af89ae1675daeca"},
	}

	for _, c := range cases {
		got := SHA1(c.in)
		assert.Equal(t, c.out, got)
	}
}

package types

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestTokenString(t *testing.T) {
	var cases = []struct {
		token Token
		out   string
	}{
		{"", ""},
		{"se", "***"},
		{"secret", "***"},
	}

	for _, c := range cases {
		got := c.token.String()
		assert.Equal(t, c.out, got)
	}
}

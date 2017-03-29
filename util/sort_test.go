package util

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSortByDict(t *testing.T) {
	var cases = []struct {
		in     []string
		expect []string
	}{
		{[]string{"b", "a", "c"}, []string{"a", "b", "c"}},
		{[]string{"tom21", "tom12", "bob"}, []string{"bob", "tom12", "tom21"}},
	}

	for _, c := range cases {
		got := SortByDict(c.in)
		assert.Equal(t, got, c.expect)
	}
}

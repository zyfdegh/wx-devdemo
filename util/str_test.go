package util

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestAppendToken(t *testing.T) {
	var cases = []struct {
		baseURL string
		token   string
		expect  string
	}{
		{
			"http://some.site.com/things?limit=10",
			"abc123",
			"http://some.site.com/things?access_token=abc123&limit=10",
		},
		{
			"https://api.weixin.qq.com/menu/create",
			"abc123",
			"https://api.weixin.qq.com/menu/create?access_token=abc123",
		},
	}

	for _, c := range cases {
		got, _ := AppendToken(c.baseURL, c.token)
		assert.Equal(t, got, c.expect)
	}
}

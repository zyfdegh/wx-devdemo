package util

import (
	"crypto/sha1"
	"fmt"
	"io"
)

// SHA1 hash
func SHA1(content string) string {
	h := sha1.New()
	io.WriteString(h, content)
	return fmt.Sprintf("%x", h.Sum(nil))
}

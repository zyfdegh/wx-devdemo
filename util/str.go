package util

import (
	"log"
	"net/url"
)

// AppendToken appends set query param 'access_token' to token
func AppendToken(baseURL, token string) (fullURL string, err error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		log.Printf("parse %s to url error: %v\n", baseURL, err)
		return
	}

	q := u.Query()
	q.Set("access_token", token)
	u.RawQuery = q.Encode()

	fullURL = u.String()
	return
}

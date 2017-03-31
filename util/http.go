package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// PostJSON post json to url
func PostJSON(url string, obj interface{}) (resp *http.Response, err error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(jsonData)
	return http.Post(url, "application/json", body)
}

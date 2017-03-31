package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// DecodeJSONFile reads a JSON file and unmarshal it to struct
func DecodeJSONFile(file string, v interface{}) error {
	// read file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("read file error: %v\n", err)
		return err
	}
	return json.Unmarshal(data, v)
}

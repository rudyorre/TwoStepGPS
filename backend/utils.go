package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// fetchAndUnmarshal fetches data from the specified URL and unmarshals it into the provided value.
// It returns an error if there was a problem fetching the data or unmarshaling it.
func fetchAndUnmarshal(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ParseBody reads and decodes the JSON body of an HTTP request into the provided struct.
func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing request body:", err)
		}
	}(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		return err
	}

	return nil
}

package httpparser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody[T any](r *http.Request) (T, error) {
	var payload T
	fmt.Println("Parsing body")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Println("Error while parsing", err)
		return payload, err
	}
	return payload, nil
}

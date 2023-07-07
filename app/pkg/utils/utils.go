package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	fmt.Println("Parsing body")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		fmt.Println("Error while parsing", err)
		return
	}
	return
}

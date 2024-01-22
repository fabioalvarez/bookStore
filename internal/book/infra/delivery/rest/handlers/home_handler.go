package handlers

import "net/http"

func (h BookHandler) Home(w http.ResponseWriter, _ *http.Request) {
	// Response message
	res := []byte("Welcome to the Book Store API")

	// Return response
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

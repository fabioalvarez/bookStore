package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	schema "github.com/fabioalvarez/bookStore/internal/book/infra/delivery/rest/handlers/schemas"
)

func (h BookHandler) GetBooks(w http.ResponseWriter, _ *http.Request) {

	books, err := h.BookAServ.GetBooks(context.Background())
	if err != nil {
		return
	}
	// Return response
	res, _ := json.Marshal(schema.ToGetSchemas(books))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	schema "github.com/fabioalvarez/bookStore/internal/book/infra/delivery/rest/handlers/schemas"
	"github.com/fabioalvarez/bookStore/pkg/httpparser"
	"github.com/go-playground/validator/v10"
)

func (h BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	// Parse request to the created empty book
	newBook, err := httpparser.ParseBody[schema.BookSchema](r)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}

	validate := validator.New()

	err = validate.Struct(newBook)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}

	// Create new book in the database
	b, err := h.BookAServ.CreateBook(context.Background(), schema.ToEntity(newBook))
	if err != nil {
		return
	}

	// Transform []Book to a JSON
	res, _ := json.Marshal(b)

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

package rest

import (
	"github.com/fabioalvarez/bookStore/internal/book/infra/delivery/rest/handlers"
	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(routes *mux.Router, bookHandlers handlers.BookHandler) {
	routes.HandleFunc("/", bookHandlers.Home)
	routes.HandleFunc("/books/", bookHandlers.CreateBook).Methods("POST")
	routes.HandleFunc("/books/", bookHandlers.GetBooks)
}

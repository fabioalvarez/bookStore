package handlers

import (
	"github.com/fabioalvarez/bookStore/internal/book/domain/services"
)

type BookHandler struct {
	BookAServ services.BookServicer
}

func NewBookHandler(service services.BookServicer) *BookHandler {
	return &BookHandler{
		BookAServ: service,
	}
}

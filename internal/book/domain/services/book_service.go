package services

import (
	"context"

	"github.com/fabioalvarez/bookStore/internal/book/domain"
)

type BookServicer interface {
	CreateBook(ctx context.Context, book domain.Book) (domain.Book, error)
	GetBooks(ctx context.Context) ([]domain.Book, error)
}

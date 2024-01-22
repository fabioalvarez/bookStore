package repository

import (
	"context"

	"github.com/fabioalvarez/bookStore/internal/book/domain"
)

type Booker interface {
	CreateBook(ctx context.Context, book domain.Book) (domain.Book, error)
	GetAllBooks(ctx context.Context) ([]domain.Book, error)
}

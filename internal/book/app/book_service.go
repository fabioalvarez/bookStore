package app

import (
	"context"

	"github.com/fabioalvarez/bookStore/internal/book/domain"
	"github.com/fabioalvarez/bookStore/internal/book/domain/repository"
	"github.com/fabioalvarez/bookStore/internal/search/app"
)

type BookService struct {
	BookRepository repository.Booker
}

func NewBookService(bookRepo repository.Booker) *BookService {
	return &BookService{BookRepository: bookRepo}
}

func (bs BookService) CreateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	app.Get()
	// all aplication logic
	return bs.BookRepository.CreateBook(ctx, book)
}

func (bs BookService) GetBooks(ctx context.Context) ([]domain.Book, error) {
	app.Get()
	// all aplication logic
	return bs.BookRepository.GetAllBooks(ctx)
}

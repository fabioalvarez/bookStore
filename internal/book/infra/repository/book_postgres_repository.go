package repository

import (
	"context"
	"fmt"

	"github.com/fabioalvarez/bookStore/internal/book/domain"
	"gorm.io/gorm"
)

type BookPostgresRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookPostgresRepository {
	return &BookPostgresRepository{
		db: db,
	}
}

func (br *BookPostgresRepository) CreateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	b := EntityToModel(book)
	result := br.db.WithContext(ctx).Create(&b)
	fmt.Println("CreateBook - Rows affected", result.RowsAffected)
	fmt.Println("CreateBook - Error", result.Error)
	return ModelToEntity(b), nil
}

func (br *BookPostgresRepository) GetAllBooks(ctx context.Context) ([]domain.Book, error) {
	var books []BookSchema
	br.db.WithContext(ctx).Find(&books)
	return ModelsToEntities(books), nil
}

package repository

import (
	"github.com/fabioalvarez/bookStore/internal/book/domain"
	"gorm.io/gorm"
)

type BookSchemaer interface {
}

type BookSchema struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func EntityToModel(entity domain.Book) BookSchema {
	return BookSchema{
		Name:        entity.Name,
		Author:      entity.Author,
		Publication: entity.Publication,
	}
}

func ModelToEntity(model BookSchema) domain.Book {
	return domain.Book{
		ID:          model.ID,
		Name:        model.Name,
		Author:      model.Author,
		Publication: model.Publication,
	}
}

func ModelsToEntities(models []BookSchema) []domain.Book {
	books := make([]domain.Book, 0, len(models))
	for _, book := range models {
		b := domain.Book{
			ID:          book.ID,
			Name:        book.Name,
			Author:      book.Author,
			Publication: book.Publication,
		}
		books = append(books, b)
	}

	return books
}

package schema

import "github.com/fabioalvarez/bookStore/internal/book/domain"

type BookSchema struct {
	Name        string `json:"name" validate:"required,min=5,max=100"`
	Author      string `json:"author" validate:"required,min=5,max=150"`
	Publication string `json:"publication"`
}

func ToEntity(body BookSchema) domain.Book {
	return domain.Book{
		Name:        body.Name,
		Author:      body.Author,
		Publication: body.Publication,
	}
}

func ToSchema(entity domain.Book) BookSchema {
	return BookSchema{
		Name:        entity.Name,
		Author:      entity.Author,
		Publication: entity.Publication,
	}
}

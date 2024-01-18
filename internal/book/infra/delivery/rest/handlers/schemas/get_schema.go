package schema

import "github.com/fabioalvarez/bookStore/internal/book/domain"

type GetBookSchema struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" validate:"required,min=5,max=100"`
	Author      string `json:"author" validate:"required,min=5,max=150"`
	Publication string `json:"publication"`
}

func GetSchemaToEntity(body GetBookSchema) domain.Book {
	return domain.Book{
		Name:        body.Name,
		Author:      body.Author,
		Publication: body.Publication,
	}
}

func ToGetSchema(entity domain.Book) GetBookSchema {
	return GetBookSchema{
		Name:        entity.Name,
		Author:      entity.Author,
		Publication: entity.Publication,
	}
}

func ToGetSchemas(books []domain.Book) []GetBookSchema {
	resp := make([]GetBookSchema, 0, len(books))

	for _, book := range books {
		resp = append(resp, GetBookSchema{
			ID:          book.ID,
			Name:        book.Name,
			Author:      book.Author,
			Publication: book.Publication,
		})

	}
	return resp
}

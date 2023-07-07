package models

import (
	"fmt"
	"github.com/fabioalvarez/bookStore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		return
	}
}

func (b *Book) CreateBook() *Book {
	result := db.Create(&b)
	fmt.Println("CreateBook - Row afected", result.RowsAffected)
	fmt.Println("CreateBook - Error", result.Error)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
}

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/fabioalvarez/bookStore/pkg/models"
	"github.com/fabioalvarez/bookStore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	res := []byte("Welcome to the Book Store API")

	// Return response
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func GetBook(w http.ResponseWriter, _ *http.Request) {
	// Get all books
	newBooks := models.GetAllBooks()

	// Transform []Book to a JSON
	res, _ := json.Marshal(newBooks)

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// Extract from Path Parameter
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Transform string to int
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	// Get book by Id
	bookDetails, _ := models.GetBookById(ID)

	// Transform []Book to a JSON
	res, _ := json.Marshal(bookDetails)

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Create empty book
	CreateBook := &models.Book{}

	// Parse request to the created empty book
	utils.ParseBody(r, CreateBook)

	// Create new book in the database
	b := CreateBook.CreateBook()

	// Transform []Book to a JSON
	res, _ := json.Marshal(b)

	// Return response
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Extract from Path Parameter
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Transform string to int
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	// Delete the book from the database
	models.DeleteBook(ID)

	// Return the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("The book was deleted"))
	if err != nil {
		return
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Create empty book
	var updateBook = &models.Book{}

	// Parse request to the created empty book
	utils.ParseBody(r, updateBook)

	// Extract from Path Parameter
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Transform string to int
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	// Extract the required book from the database
	bookDetails, db := models.GetBookById(ID)

	// Conditions to change values
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// Save changes in the database
	db.Save(&bookDetails)

	// Transform []Book to a JSON
	res, _ := json.Marshal(bookDetails)

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

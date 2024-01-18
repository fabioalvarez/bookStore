package main

import (
	"log"
	"net/http"

	"github.com/fabioalvarez/bookStore/internal/book/app"
	"github.com/fabioalvarez/bookStore/internal/book/infra/delivery/rest"
	"github.com/fabioalvarez/bookStore/internal/book/infra/delivery/rest/handlers"
	"github.com/fabioalvarez/bookStore/internal/book/infra/repository"
	"github.com/fabioalvarez/bookStore/internal/common/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := mux.NewRouter()

	db.Connect()
	conn := db.DB()
	log.Println("DB connected")
	book_repo := repository.NewBookRepository(conn)
	book_service := app.NewBookService(book_repo)
	h := handlers.NewBookHandler(*book_service)

	log.Println("routes")
	rest.RegisterBookStoreRoutes(r, *h)
	log.Fatal(http.ListenAndServe(":4000", r))
}

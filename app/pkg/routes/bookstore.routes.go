package routes

import (
	"github.com/fabioalvarez/bookStore/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(routes *mux.Router) {
	routes.HandleFunc("/", controllers.Home)
	routes.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	routes.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	routes.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	routes.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	routes.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

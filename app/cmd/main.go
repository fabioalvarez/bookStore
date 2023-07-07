package main

import (
	"github.com/fabioalvarez/bookStore/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe(":4000", r))
}

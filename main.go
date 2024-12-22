package main

import (
	"Go-BooksStore/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/postgres"
	"log"
	"net/http"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterBookStoreRoutes(route)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe("localhost:8080", route))
}

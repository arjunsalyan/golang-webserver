package routes

import (
	"learning_project/handlers"

	"github.com/gorilla/mux"
)

func APIRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/book/get/{id}", handlers.GetBook)
	r.HandleFunc("/books/get", handlers.GetBooks)
	return r
}

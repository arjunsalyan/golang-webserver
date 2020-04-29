package routes

import (
	"learning_project/handlers"

	"github.com/gorilla/mux"
)

func WebRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("404", handlers.NotFoundHandler)
	r.HandleFunc("/read-file/", handlers.FileReadHandler)
	r.HandleFunc("/books/", handlers.BooksViewHandler)
	r.HandleFunc("/get-params/", handlers.GetParamsViewHandler)
	r.HandleFunc("/elements/", handlers.ElementsViewHandler)
	r.HandleFunc("/", handlers.HomeViewHandler)
	return r
}

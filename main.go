package main

import (
	"learning_project/handlers"
	"log"
	"net/http"
)

func main() {
	// Create the database connection string
	http.HandleFunc("404", handlers.NotFoundHandler)
	http.HandleFunc("/read-file/", handlers.FileReadHandler)
	http.HandleFunc("/get-params/", handlers.GetParamsViewHandler)
	http.HandleFunc("/elements/", handlers.ElementsViewHandler)
	http.HandleFunc("/", handlers.HomeViewHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

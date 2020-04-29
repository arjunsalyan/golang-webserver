package main

import (
	"learning_project/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r = routes.APIRoutes(r)
	r = routes.WebRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

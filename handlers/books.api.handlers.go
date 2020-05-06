package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"learning_project/auth"
	"learning_project/postgres"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	tokenAuth, err := auth.ExtractTokenMetadata(r)
	if err != nil {
		json.NewEncoder(w).Encode(ResponseError{http.StatusUnauthorized, "unauthorized"})
		return
	}
	_, err = auth.FetchAuth(tokenAuth)
	if err != nil {
		json.NewEncoder(w).Encode(ResponseError{http.StatusUnauthorized, "unauthorized"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		APIError(w, 400, "Bad Request. Invalid ID")
		return
	}
	book, err := postgres.GetOneBook(id)
	if err != nil {
		APIError(w, 404, err.Error())
		return
	}
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := postgres.GetAllBooks()
	if err != nil {
		APIError(w, 200, err.Error())
	}
	json.NewEncoder(w).Encode(books)
}

func APIError(w http.ResponseWriter, status int, m string) {
	e := ResponseError{
		Status:  status,
		Message: m,
	}
	json.NewEncoder(w).Encode(e)
}

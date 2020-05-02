package handlers

import (
	"learning_project/postgres"
	"learning_project/redis"
	"learning_project/templates"
	"net/http"
	"strconv"
)

func BooksViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			templates.BasicHTTPRender(w, "Failed to parse POST data.")
			return
		}
		name := r.FormValue("name")
		author := r.FormValue("author")
		pageStr := r.FormValue("pages")

		pages, err := strconv.Atoi(pageStr)
		if err != nil {
			templates.BasicHTTPRender(w, "Pages must be a valid integer.")
		}
		bookObj := postgres.Book{redis.RedisBook{
			Name:   name,
			Author: author,
			Pages:  pages,
		},
		}

		ok := postgres.Insert(bookObj)
		if !ok {
			templates.BasicHTTPRender(w, "Failed to add the book to the database")
		}
	}
	books, err := postgres.GetAllBooks()
	if err != nil {
		templates.BasicHTTPRender(w, "Error reading from database")
		return
	}
	templates.RenderTemplate(w, "templates/books", &templates.Page{Title: "Books", Books: books})
}

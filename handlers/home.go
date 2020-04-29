package handlers

import (
	"learning_project/templates"
	"net/http"
)

func HomeViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}
	p := &templates.Page{Title: "Home"}
	templates.RenderTemplate(w, "templates/home", p)
}

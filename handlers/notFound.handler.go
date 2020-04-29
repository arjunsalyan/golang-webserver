package handlers

import (
	"learning_project/templates"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	p := &templates.Page{Title: "Not Found"}
	templates.RenderTemplate(w, "templates/404", p)
}

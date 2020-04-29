package handlers

import (
	"learning_project/templates"
	"net/http"
	"strings"
)

func GetParamsViewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := r.URL.Query()["title"]
	if !err || len(title[0]) < 1 {
		templates.BasicHTTPRender(w, "Error: Title is not supplied")
		return
	}
	body, err := r.URL.Query()["body"]
	if !err || len(body[0]) < 1 {
		templates.BasicHTTPRender(w, "Error: Body is not supplied")
		return
	}

	p := &templates.Page{Title: strings.Join(title, " "), Body: strings.Join(body, " ")}
	templates.RenderTemplate(w, "templates/dynamic", p)
}

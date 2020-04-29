package handlers

import (
	"io/ioutil"
	"learning_project/templates"
	"net/http"
)

func FileReadHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/read-file/"):]
	p, err := loadPageFromFile(title)
	if err != nil {
		templates.BasicHTTPRender(w, "Error opening that file. Expected path: /read-file/{{filename}}  [please do not use file extension]")
		return
	}
	templates.RenderTemplate(w, "templates/read-file", p)
}

func loadPageFromFile(title string) (*templates.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &templates.Page{Title: title, Body: string(body)}, nil
}

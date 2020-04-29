package handlers

import (
	"html/template"
	"learning_project/postgres"
	"learning_project/templates"
	"net/http"
)

func ElementsViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			templates.BasicHTTPRender(w, "Failed to add the element")
			return
		}
		ele := postgres.Element{
			Text: r.FormValue("list_object"),
		}
		isAdded := postgres.Insert(ele)
		if !isAdded {
			templates.BasicHTTPRender(w, "Failed")
			return
		}
	}

	elements, err := postgres.GetAllElements()
	if err != nil {
		templates.BasicHTTPRender(w, "Error reading from database")
		return
	}
	t, err := template.ParseFiles("templates/list.html")
	if err != nil {
		templates.BasicHTTPRender(w, "Failed to parse the template.")
	}
	p := &templates.Page{Title: "List View", Elements: elements}
	err = t.Execute(w, p)
	if err != nil {
		templates.BasicHTTPRender(w, "Failed to parse the template.")
	}
}

package handlers

import (
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
	templates.RenderTemplate(w, "templates/elements", &templates.Page{Title: "Elements View", Elements: elements})
}

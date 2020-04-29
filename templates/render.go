package templates

import (
	"fmt"
	"html/template"
	"learning_project/postgres"
	"net/http"
)

// Page is structure that defines a page
type Page struct {
	Title    string
	Body     string
	Elements []postgres.Element
}

// BasicHTTPRender writes a simple text based http response
func BasicHTTPRender(w http.ResponseWriter, message string) {
	fmt.Fprintf(w, "%s", message)
}

// RenderTemplate is a helper function to render the template by taking in
// Page and name of the template as parameters.
func RenderTemplate(w http.ResponseWriter, v string, p *Page) {
	t, _ := template.ParseFiles(v + ".html")
	t.Execute(w, p)
}

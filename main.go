package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Page is structure that defines a page
type Page struct {
	Title string
	Body  string
}

func basicHTTPRender(w http.ResponseWriter, message string) {
	fmt.Fprintf(w, "%s", message)
}

// renderTemplate is a helper function to render the template by taking in
// Page and name of the template as parameters.
func renderTemplate(w http.ResponseWriter, v string, p *Page) {
	t, _ := template.ParseFiles(v + ".html")
	t.Execute(w, p)
}

// loadPageFromFile is a helper function to open a file and read its content.
// It then returns a Page object
func loadPageFromFile(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(body)}, nil
}

func main() {
	http.HandleFunc("404", notFoundHandler)
	http.HandleFunc("/read-file/", fileReadHandler)
	http.HandleFunc("/view/", dynamicViewHandler)
	http.HandleFunc("/", homeViewHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

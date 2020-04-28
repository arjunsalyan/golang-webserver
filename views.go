package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func homeViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
	}
	p := &Page{Title: "Home"}
	renderTemplate(w, "templates/home", p)
}

func fileReadHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/read-file/"):]
	p, err := loadPageFromFile(title)
	if err != nil {
		basicHTTPRender(w, "Error opening that file.")
		return
	}
	renderTemplate(w, "templates/read-file", p)
}

func dynamicViewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := r.URL.Query()["title"]
	if !err || len(title[0]) < 1 {
		basicHTTPRender(w, "Error: Title is not supplied")
		return
	}
	body, err := r.URL.Query()["body"]
	if !err || len(body[0]) < 1 {
		basicHTTPRender(w, "Error: Body is not supplied")
		return
	}

	p := &Page{Title: strings.Join(title, " "), Body: strings.Join(body, " ")}
	renderTemplate(w, "templates/dynamic", p)
}

func listViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Here")
		err := r.ParseForm()
		if err != nil {
			basicHTTPRender(w, "Failed to add the element")
			return
		}
		ele := r.FormValue("list_object")
		isAdded := inserElement(ele)
		if !isAdded {
			basicHTTPRender(w, "Failed")
			return
		}
	}

	elements, err := getAllElements()
	if err != nil {
		basicHTTPRender(w, "Error reading from database")
		return
	}
	t, err := template.ParseFiles("templates/list.html")
	if err != nil {
		basicHTTPRender(w, "Failed to parse the template.")
	}
	p := &Page{Title: "List View", Elements: elements}
	err = t.Execute(w, p)
	if err != nil {
		basicHTTPRender(w, "Failed to parse the template.")
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Not Found"}
	renderTemplate(w, "templates/404", p)
}

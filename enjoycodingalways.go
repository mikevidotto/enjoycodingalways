package main

import (
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
)

func main() {
	// directory := http.Dir("./static/")
	// fs := http.FileServer(directory)
	// http.Handle("/", fs)

	date1 := time.Now()
	title1 := "Hello WOrld"
	body1 := "body body body body body body body body body body body body "
	http.Handle("/", templ.Handler(index(date1, title1, body1)))
	http.Handle("/projects.html", templ.Handler(projects(date1, title1, body1)))
	http.Handle("/documentation.html", templ.Handler(documentation(date1, title1, body1)))
	http.Handle("/about.html", templ.Handler(about(date1, title1, body1)))
	http.Handle("/404", templ.Handler(notFoundComponent(), templ.WithStatus(http.StatusNotFound)))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error with ListenAndServe:", err)
	}
}

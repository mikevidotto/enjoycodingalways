package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	// directory := http.Dir("./static/")
	// fs := http.FileServer(directory)
	// http.Handle("/", fs)

	// date1 := time.Now()
	// title1 := "Hello WOrld"
	// body1 := "body body body body body body body body body body body body "
	// http.Handle("/projects.html", templ.Handler(projects(date1, title1, body1)))
	// http.Handle("/documentation.html", templ.Handler(documentation(date1, title1, body1)))
	// http.Handle("/about.html", templ.Handler(about(date1, title1, body1)))
	// http.Handle("/404", templ.Handler(notFoundComponent(), templ.WithStatus(http.StatusNotFound)))

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error with ListenAndServe:", nil)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	date := time.Now()
	title := "Hello WOrld"
	body := "body body body body body body body body body body body body "

	path := r.URL.Path

	switch path {
	case "/", "/index.html":
		index(date, title, body).Render(context.Background(), w)
	case "/projects.html":
		projects(date, title, body).Render(context.Background(), w)
	case "/documentation.html":
		documentation(date, title, body).Render(context.Background(), w)
	case "/about.html":
		about(date, title, body).Render(context.Background(), w)
	}

}

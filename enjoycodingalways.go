package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/documentation", documentationHandler)
	http.HandleFunc("/about", aboutHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error with ListenAndServe:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		webpage().Render(context.Background(), w)
	case "POST":
		fmt.Println("POST: rendering home.")
		HomeMain().Render(context.Background(), w)
	}
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println("POST: rendering projects.")
		ProjectsMain().Render(context.Background(), w)
	}
}

func documentationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println("POST: rendering documentation.")
		DocumentationMain().Render(context.Background(), w)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println("POST: rendering about.")

		AboutMain().Render(context.Background(), w)
	}
}

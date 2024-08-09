package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Article struct {
	date  time.Time
	title string
	body  string
}

var articles []Article

func main() {

	articles = append(articles, Article{date: time.Now(), title: "Default article", body: "this article was created upon entering the website."})
	articles = append(articles, Article{date: time.Now(), title: "Article 2", body: "two two two two two two two two two two "})
	articles = append(articles, Article{date: time.Now(), title: "Third article", body: "three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. "})

	directory := http.Dir(".")
	fs := http.FileServer(directory)
	http.Handle("/static/", fs)

	http.HandleFunc("/", handler)
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/documentation", documentationHandler)
	http.HandleFunc("/about", aboutHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error with ListenAndServe:", err)
	}
}

// <div id="recent-project">
//			<div style="display:flex">
//             <h6>Currently working on...</h6>
//             <h1 class="project-title">Enjoy Coding Always</h1>
//             <img src="/static/images/projectimage.png" alt="project image" width="700" height="500" class="project-image"/>
//             <h6 class="project-description"></h6>
//         </div>

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		webpage(articles).Render(context.Background(), w)
	case "POST":
		fmt.Println("POST: rendering home.")
		HomeMain(articles).Render(context.Background(), w)
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
		DocumentationMain(articles).Render(context.Background(), w)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println("POST: rendering about.")

		AboutMain().Render(context.Background(), w)
	}
}

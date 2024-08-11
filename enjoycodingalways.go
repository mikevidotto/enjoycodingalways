package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Project struct {
	title       string
	description string
}

type Article struct {
	date  time.Time
	title string
	body  string
}

var articles1 = []Article{
	{date: time.Now(), title: "title1", body: "body1body1body1body1body1body1body1"},
	{date: time.Now(), title: "title2", body: "body2body2body2body2body2body2body2"},
	{date: time.Now(), title: "title3", body: "body3body3body3body3body3body3body3"},
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
	http.HandleFunc("/viewdoc", viewDocHandler)
	http.HandleFunc("/createdoc", createDocHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/clicked", clickedHandler)

	fmt.Println("Starting server...")
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		log.Fatal("Error with ListenAndServe:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("Rendering homepage.")
		webpage(articles1).Render(context.Background(), w)
	case "POST":
		fmt.Println("POST: rendering home page.")
		HomeMain(articles1).Render(context.Background(), w)
	}
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println("POST: rendering projects page.")
		ProjectsMain().Render(context.Background(), w)
	}
}

func documentationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println("POST: rendering documentation page.")
		DocumentationMain(articles1).Render(context.Background(), w)
	}
}

func viewDocHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET: rendering view documentation page.")
		articlesList().Render(context.Background(), w)
	case "POST":
	}
}

func createDocHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET: rendering create documentation page.")
		CreateDoc().Render(context.Background(), w)
	case "POST":
		//		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		fmt.Println("POST: posted article.")
		newArticle := Article{title: r.FormValue("title"), body: r.FormValue("body"), date: time.Now()}
		articles = append(articles, newArticle)
		articlesList().Render(context.Background(), w)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		fmt.Println("POST: rendering about page.")

		AboutMain().Render(context.Background(), w)
	}
}

func clickedHandler(w http.ResponseWriter, r *http.Request) {
	//do something but not sure what yet.
	fmt.Println("something was clicked.")
}

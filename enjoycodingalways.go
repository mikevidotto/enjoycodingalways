package main

import (
	"database/sql"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	connStr = "user=postgres dbname=enjoycodingalways host=localhost port=5432 password=password sslmode=disable"
	
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

var articles []Article

func main() {
	createArticlesTable()

	articles = append(articles, Article{date: time.Now(), title: "Default article", body: "this article was created upon entering the website."})
	articles = append(articles, Article{date: time.Now(), title: "Article 2", body: "two two two two two two two two two two "})
	articles = append(articles, Article{date: time.Now(), title: "Third article", body: "three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. three. "})


	for _, article := range articles {
		fmt.Println(article.title)
	}
	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	checkError(err)
	fmt.Println("db open")

	pingDatabase(db)
	fmt.Println("ping successful")


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

	fmt.Println("Starting server at localhost:5000...")
	err = http.ListenAndServe("localhost:5000", nil)
	if err != nil {
		log.Fatal("Error with ListenAndServe:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("Rendering homepage.")
		webpage(articles).Render(context.Background(), w)
	case "POST":
		fmt.Println("POST: rendering home page.")
		HomeMain(articles).Render(context.Background(), w)
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
		DocumentationMain(articles).Render(context.Background(), w)
	}
}

func viewDocHandler(w http.ResponseWriter, r *http.Request) {
	articles = updatedList(w)
	fmt.Println("Getting articles:")
	fmt.Println(articles)
	switch r.Method { case "GET":
		fmt.Println("GET: rendering view documentation page.")
		articlesList().Render(context.Background(), w)
	case "POST":
	}
}

func createArticlesTable() {
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	checkError(err)

  query := `CREATE TABLE IF NOT EXISTS articles(id SERIAL PRIMARY KEY,date DATE,title VARCHAR(50),body VARCHAR(100))`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("ERROR CREATING TABLE:", err)
	}
}

func updatedList(w http.ResponseWriter) []Article {
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	checkError(err)

	var updatedArticles []Article

	query := `SELECT * FROM articles`
  
	rows, err := db.Query(query)
	checkError(err)

	defer rows.Close()

	for rows.Next() {
		var date time.Time 
		var title string 
		var body string 

		err = rows.Scan(&date, &title, &body)
		checkError(err)

		_, err = fmt.Fprint(w, date, title, body)
		checkError(err)
	}

	return updatedArticles 
}

func pingDatabase(db *sql.DB) {
	err := db.Ping()
	checkError(err)
	fmt.Println("PING SUCCESSFUL.")
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

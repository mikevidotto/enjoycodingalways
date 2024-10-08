package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Todo struct {
	task string
	done int
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// getarticles writes every record in the table to the http.ResponseWriter
func getArticles(w http.ResponseWriter, db *sql.DB) {
	query := `SELECT * FROM articles`
	rows, err := db.Query(query)
	checkError(err)

	defer rows.Close()

	for rows.Next() {
		var id int
		var task string
		var done int

		err := rows.Scan(&id, &task, &done)
		checkError(err)

		_, err = fmt.Fprintf(w, "{\"id\":%d, \"task\":\"%s\", \"isDone\":%d}", id, task, done)
		checkError(err)
	}
}

func insertData(a Article) (pk int) {
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	checkError(err)

	query := `INSERT INTO articles (date, title, body) VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(query, a.date, a.title, a.body).Scan(&pk)
	if err != nil {
		fmt.Println("couldn't insert into articles.")
		log.Fatal(err)
	}
	return pk
}

func newTodo(task string) Todo {
	return Todo{task: task, done: 0}
}

func deleteTodo(db *sql.DB, id int) (pk int, err error) {
	query := `DELETE FROM articles WHERE id=$1 RETURNING id`
	err = db.QueryRow(query, id).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}

	return pk, err
}


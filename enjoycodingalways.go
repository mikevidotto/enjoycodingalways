package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error with ListenAndServe:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	date := time.Now()
	title := "Hello WOrld"
	body := "body body body body body body body body body body body body "
	path := r.URL.Path

	index(date, title, body, path).Render(context.Background(), w)

}

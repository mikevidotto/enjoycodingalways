package main

import (
	"log"
	"net/http"
)

func main() {
	directory := http.Dir("./static/")
	fs := http.FileServer(directory)
	http.Handle("/", fs)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error with ListenAndServe:", err)
	}
}

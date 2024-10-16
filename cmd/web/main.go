package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /chore/create", createChore)
	mux.HandleFunc("GET /chore/view/{choreId}", viewChore)
	mux.HandleFunc("POST /chore/create", createChorePost)

	log.Print("Listening on port 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

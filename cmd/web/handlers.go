package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createChore(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a chore"))
}

func viewChore(w http.ResponseWriter, r *http.Request) {
	choreId, err := strconv.Atoi(r.PathValue("choreId"))
	if err != nil || choreId < 1 {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "Displaying a specific chore with id %d", choreId)
}

func createChorePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("creating a new chore"))
}

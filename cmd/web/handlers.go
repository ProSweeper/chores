package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	templates := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	ts, err := template.ParseFiles(templates...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewChore(w http.ResponseWriter, r *http.Request) {
	choreId, err := strconv.Atoi(r.PathValue("choreId"))
	if err != nil || choreId < 1 {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "Displaying a specific chore with id %d", choreId)
}
func createChore(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a chore"))
}

func createChorePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("creating a new chore"))
}

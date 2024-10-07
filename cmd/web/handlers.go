package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Chores"))
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

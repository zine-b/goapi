package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	dbConnect()
	r := mux.NewRouter()
	r.HandleFunc("/films", getFilms).Methods("GET")
	r.HandleFunc("/films/{id}", getFilmById).Methods("GET")
	r.HandleFunc("/films", addFilm).Methods("POST")
	http.Handle("/", r)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

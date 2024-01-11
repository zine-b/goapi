package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"net/http"
	"strconv"
	"strings"
)

type MessageDto struct {
	Message string `json:"message"`
}

func getFilms(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query(`SELECT film_id, title,description,release_year,language_id,rental_duration, 
                           rental_rate, length, replacement_cost, rating, last_update, special_features, fulltext  
                         FROM film ORDER BY film_id`)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var films []Film
	for rows.Next() {
		film := scanFilm(err, rows)

		films = append(films, film)
	}
	json.NewEncoder(w).Encode(films)
}

func getFilmById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	filmId := vars["id"]

	rows, err := db.Query(`SELECT film_id, 
       title,
       description,
       release_year,
       language_id,
       rental_duration,
       rental_rate, 
       length, 
       replacement_cost, 
       rating,
       last_update,
       special_features,
       fulltext
FROM film WHERE film_id=$1`, filmId)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if !rows.Next() {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(asMessage("Film not found"))

		return
	}
	film := scanFilm(err, rows)
	json.NewEncoder(w).Encode(film)
}

func asMessage(mm string) MessageDto {
	m := MessageDto{
		Message: mm,
	}
	return m
}

func addFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var film Film
	if err := json.NewDecoder(r.Body).Decode(&film); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := db.Exec(`INSERT INTO film (title, 
     description, 
     release_year, 
     language_id, 
     rental_duration, 
     rental_rate, 
     length, 
     replacement_cost, 
     rating,
     special_features, 
     fulltext)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10, $11)`,
		film.Title,
		film.Description,
		film.Release_year,
		film.Language_id,
		film.Rental_duration,
		film.Rental_rate,
		film.Length,
		film.Replacement_cost,
		film.Rating,
		pq.Array(film.Special_features),
		mapToString(film.Fulltext))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(asMessage("Film added successfully"))
}

func mapToString(data map[string]int) string {
	var keyValuePairs []string

	for key, value := range data {
		keyValuePairs = append(keyValuePairs, fmt.Sprintf("'%s':%d", key, value))
	}

	return strings.Join(keyValuePairs, " ")
}

func stringToMap(input string) map[string]int {
	result := make(map[string]int)
	pairs := strings.Fields(input)
	for _, pair := range pairs {
		// Split each pair into key and value using ":"
		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			return nil
		}
		key := strings.Trim(parts[0], "'")      // Remove single quotes from the key
		valueStr := strings.Trim(parts[1], "'") // Remove single quotes from the value
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil
		}
		result[key] = value
	}
	return result
}

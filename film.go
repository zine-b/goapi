package main

import (
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type Film struct {
	FilmId           int            `json:"filmId"`
	Title            string         `json:"title"`
	Description      string         `json:"description"`
	Release_year     int            `json:"releaseYear"`
	Language_id      int            `json:"languageId"`
	Rental_duration  int            `json:"rentalDuration"`
	Rental_rate      float64        `json:"rentalRate"`
	Length           int            `json:"length"`
	Replacement_cost float64        `json:"replacementCost"`
	Rating           string         `json:"rating"`
	Last_update      time.Time      `json:"lastUpdate"`
	Special_features []string       `json:"specialFeatures"`
	Fulltext         map[string]int `json:"fullText"`
}

func scanFilm(err error, rows *sql.Rows) Film {
	var fulltext string
	var film Film
	err = rows.Scan(&film.FilmId,
		&film.Title,
		&film.Description,
		&film.Release_year,
		&film.Language_id,
		&film.Rental_duration,
		&film.Rental_rate,
		&film.Length,
		&film.Replacement_cost,
		&film.Rating,
		&film.Last_update,
		pq.Array(&film.Special_features),
		&fulltext)
	if err != nil {
		panic(err)
	}
	film.Fulltext = stringToMap(fulltext)
	return film
}

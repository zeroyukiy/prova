package model

import "time"

type Film struct {
	ID              int        `db:"film_id"`
	Title           string     `db:"title"`
	Description     string     `db:"description"`
	ReleaseYear     int        `db:"release_year"`
	LanguageID      int        `db:"language_id"`
	RentalDuration  int        `db:"rental_duration"`
	RentalRate      float32    `db:"rental_rate"`
	Length          int        `db:"length"`
	ReplacementCost float32    `db:"replacement_cost"`
	Rating          string     `db:"rating"`
	SpecialFeatures string     `db:"special_features"`
	LastUpdate      *time.Time `db:"last_update"`
}

package model

import "time"

type Film struct {
	ID              int        `mysql:"film_id"`
	Title           string     `mysql:"title"`
	Description     string     `mysql:"description"`
	ReleaseYear     int        `mysql:"release_year"`
	LanguageID      int        `mysql:"language_id"`
	RentalDuration  int        `mysql:"rental_duration"`
	RentalRate      float32    `mysql:"rental_rate"`
	Length          int        `mysql:"length"`
	ReplacementCost float32    `mysql:"replacement_cost"`
	Rating          string     `mysql:"rating"`
	SpecialFeatures string     `mysql:"special_features"`
	LastUpdate      *time.Time `mysql:"last_update"`
}

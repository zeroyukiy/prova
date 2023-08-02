package model

import "time"

type Actor struct {
	ID         int        `db:"actor_id"`
	FirstName  string     `db:"first_name"`
	LastName   string     `db:"last_name"`
	LastUpdate *time.Time `db:"last_update"`
}

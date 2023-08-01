package model

import "time"

type Actor struct {
	ID         int        `mysql:"actor_id"`
	FirstName  string     `mysql:"first_name"`
	LastName   string     `mysql:"last_name"`
	LastUpdate *time.Time `mysql:"last_update"`
}

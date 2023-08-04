package model

import "time"

type Actor struct {
	ID         int        `db:"actor_id" json:",omitempty"`
	FirstName  string     `db:"first_name"`
	LastName   string     `db:"last_name"`
	LastUpdate *time.Time `db:"last_update" json:",omitempty"`
}

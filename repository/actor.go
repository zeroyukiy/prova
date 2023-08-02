package repository

import (
	"log"
	"prova/model"

	"github.com/jmoiron/sqlx"
)

type ActorRepository struct {
	DB  *sqlx.DB
	Log log.Logger
}

func (ar *ActorRepository) All() (interface{}, error) {
	query := `SELECT actor_id, first_name, last_name, last_update FROM actor LIMIT 20`
	rows, err := ar.DB.Queryx(query)
	if err != nil {
		ar.Log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var actors []model.Actor
	for rows.Next() {
		var actor model.Actor
		err := rows.StructScan(&actor)
		if err != nil {
			ar.Log.Fatal(err)
			return nil, err
		}
		actors = append(actors, actor)
	}

	return actors, nil
}

func (ar *ActorRepository) Get(id int) (interface{}, error) {
	query := `SELECT actor_id, first_name, last_name, last_update FROM actor WHERE actor_id = ?`

	var actor model.Actor
	err := ar.DB.QueryRowx(query, id).StructScan(&actor)
	if err != nil {
		ar.Log.Fatal(err)
	}

	return actor, nil
}
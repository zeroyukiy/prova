package repository

import (
	"database/sql"
	"log"
	"prova/model"
)

type ActorRepository struct {
	DB  *sql.DB
	Log log.Logger
}

func (ar *ActorRepository) All() ([]model.Actor, error) {
	query := "SELECT actor_id, first_name, last_name, last_update FROM actor LIMIT 20"
	rows, err := ar.DB.Query(query)
	if err != nil {
		ar.Log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var actors []model.Actor
	for rows.Next() {
		var actor model.Actor
		err := rows.Scan(&actor.ID, &actor.FirstName, &actor.LastName, &actor.LastUpdate)
		if err != nil {
			ar.Log.Fatal(err)
			return nil, err
		}
		actors = append(actors, actor)
	}

	return actors, nil
}

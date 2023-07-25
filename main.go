package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	Title string
	Todos []Todo
}

type Actor struct {
	ID         int        `mysql:"actor_id"`
	FirstName  string     `mysql:"first_name"`
	LastName   string     `mysql:"last_name"`
	LastUpdate *time.Time `mysql:"last_update"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sakila?parseTime=true")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	tmpl, err := template.ParseFiles("todo-list.html")
	if err != nil {
		log.Println(err)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		query := `SELECT actor_id, first_name, last_name, last_update FROM actor WHERE actor_id = ?`
		// err := db.QueryRow(query, 1).Scan(&actor.ID, &actor.FirstName, &actor.LastName, &actor.LastUpdate)
		res, err := db.Query(query, 1)
		if err != nil {
			log.Println(err)
		}
		defer res.Close()

		if res.Next() {
			var actor Actor
			err := res.Scan(&actor.ID, &actor.FirstName, &actor.LastName, &actor.LastUpdate)
			if err != nil {
				log.Println(err)
			}

			fmt.Printf("Actor: %+v", &actor)

			data, err := json.Marshal(&actor)
			if err != nil {
				log.Println(err)
			}

			w.Write(data)
		}
	})

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "index.html")
		todos := TodoPageData{
			Title: "Todos List",
			Todos: []Todo{
				{
					Title: "Fare la spesa",
					Done:  true,
				},
				{
					Title: "Correre 5 km",
					Done:  true,
				},
				{
					Title: "Andare al cinema",
					Done:  false,
				},
			},
		}
		tmpl.Execute(w, todos)
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}

package main

import (
	"database/sql"
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

type Film struct {
	ID          int    `mysql:"film_id"`
	Title       string `mysql:"title"`
	Description string `mysql:"description"`
	ReleaseYear int    `mysql:"release_year"`
}

type ViewFilm struct {
	Title string
	Films []Film
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sakila?parseTime=true")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./assets/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			log.Println(err)
		}

		query := `SELECT f.film_id, f.title, f.description, f.release_year FROM actor a INNER JOIN film_actor fa ON fa.actor_id = a.actor_id INNER JOIN film f ON f.film_id = fa.film_id WHERE a.actor_id = ?`

		res, err := db.Query(query, 1)
		if err != nil {
			log.Println(err)
		}
		defer res.Close()

		var films []Film
		for res.Next() {
			var film Film
			err := res.Scan(&film.ID, &film.Title, &film.Description, &film.ReleaseYear)
			if err != nil {
				log.Println(err)
			}
			films = append(films, film)
		}

		view := ViewFilm{
			Title: "Films",
			Films: films,
		}

		tmpl.Execute(w, view)
	})

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("todo-list.html")
		if err != nil {
			log.Println(err)
		}

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

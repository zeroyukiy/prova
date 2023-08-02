package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"prova/handler"
	"prova/mailer"
	"prova/repository"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	Title string
	Todos []Todo
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sakila?parseTime=true")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	actorHandler := handler.Handler{
		Repository: &repository.ActorRepository{
			DB:  db,
			Log: *log.Default(),
		},
		Mailer: mailer.NewMailer(),
	}

	e := echo.New()
	e.Static("/static", "assets")
	e.Renderer = t
	e.GET("/actors", actorHandler.GetAll)
	e.GET("/actors/:id", actorHandler.Get)
	e.POST("/addfilm", addFilm)

	e.GET("/todos", showTodos)

	e.Logger.Fatal(e.Start(":8000"))
}

func addFilm(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		d, err := json.Marshal(struct {
			Status int
			Error  string
		}{
			Status: http.StatusOK,
			Error:  "method is not POST",
		})
		if err != nil {
			log.Println(err)
		}

		return c.JSON(http.StatusBadRequest, d)
	} else {
		var filmForm struct {
			Title       string
			Description string
		}
		filmForm.Title = c.FormValue("title")
		filmForm.Description = c.FormValue("description")

		fmt.Println(filmForm)
		// d, err := json.Marshal(filmForm)
		// if err != nil {
		// 	log.Println(err)
		// }

		// a := fmt.Sprintf(`<h2 style="color:red;">%s</h2>
		// <p>%s</p>`, filmForm.Title, filmForm.Description)

		// b := c.Render(http.StatusOK, "index.html", filmForm)

		return c.Render(http.StatusOK, "a.html", filmForm)

		// return c.String(http.StatusOK, a)
	}
}

func showTodos(c echo.Context) error {
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
	return c.Render(http.StatusOK, "todo-list.html", todos)
}

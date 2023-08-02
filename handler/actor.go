package handler

import (
	"fmt"
	"net/http"
	"prova/mailer"
	"prova/model"
	"prova/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Repository *repository.ActorRepository
	Mailer *mailer.Mailer
}

func (h Handler) GetAll(c echo.Context) error {
	actors, err := h.Repository.All()
	if err != nil {
		h.Repository.Log.Fatal(err)
		fmt.Println(err)
	}

	var viewActors struct {
		Actors []model.Actor
	}
	viewActors.Actors = actors.([]model.Actor)

	h.Mailer.SendEmail("localhost@localhost.example")

	return c.Render(http.StatusOK, "actors.html", viewActors)
}

func (h Handler) Get(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	actor, err := h.Repository.Get(id)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, actor)
}
package handler

import (
	"fmt"
	"net/http"
	"prova/model"
	"prova/repository"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Repository *repository.ActorRepository
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
	viewActors.Actors = actors

	return c.Render(http.StatusOK, "actors.html", viewActors)
}

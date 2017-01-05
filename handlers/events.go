package handlers

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	loccasions "github.com/ruprict/loccasions-go"
	"github.com/ruprict/loccasions-go/repository"
)

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
}

type EventsHandler struct {
	Repo repository.Repository
}

func (e *EventsHandler) CreateEvent(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims)
	json := new(loccasions.Event)
	if err := c.Bind(json); err == nil {
		fmt.Println(json)
		e.Repo.CreateEventForUser(claims.ID, json)

		return c.JSON(201, json)
	} else {
		return c.JSON(500, map[string]string{"error": "error creating event"})
	}
}
func (e *EventsHandler) GetEvents(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims)
	events := e.Repo.GetEventsForUser(claims.ID)
	return c.JSON(200, events)
}

package handlers

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	loccasions "github.com/ruprict/loccasions-go"
	"github.com/ruprict/loccasions-go/repository"
)

type EventsHandler struct {
	Repo repository.Repository
}

func (e *EventsHandler) CreateEvent(c echo.Context) error {
	cc := c.(*CustomContext)
	claims := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims)
	json := new(loccasions.Event)
	if err := c.Bind(json); err == nil {
		cc.Repo.CreateEventForUser(claims.ID, json)

		return c.JSON(201, json)
	} else {
		return c.JSON(500, map[string]string{"error": "error creating event"})
	}
}
func (e *EventsHandler) GetEvents(c echo.Context) error {
	cc := c.(*CustomContext)
	claims := cc.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims)
	events := cc.Repo.GetEventsForUser(claims.ID)
	return c.JSON(200, events)
}

func (e *EventsHandler) GetEvent(c echo.Context) error {
	cc := c.(*CustomContext)
	claims := cc.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims)
	id := c.Param("id")
	event := cc.Repo.GetEventForUser(claims.ID, id)
	if event == nil {
		return cc.JSON(404, map[string]string{"error": "event not found"})
	} else {
		return cc.JSON(200, event)
	}
}

func (e *EventsHandler) PatchEvent(c echo.Context) error {
	var json loccasions.Event
	id := c.Param("id")
	cc := c.(*CustomContext)
	//claims := cc.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims)

	//TODO: Make sure user owns event
	if cc.Bind(&json) == nil {
		cc.Repo.UpdateEvent(id, &json)
		return cc.JSON(200, json)
	} else {
		return cc.JSON(500, map[string]string{"error": "error updating event"})
	}
}

func (e *EventsHandler) DeleteEvent(c echo.Context) error {
	id := c.Param("id")
	cc := c.(*CustomContext)
	//TODO: Make sure user owns event
	cc.Repo.DeleteEvent(id)
	return c.NoContent(http.StatusNoContent)
}

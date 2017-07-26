package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	loccasions "github.com/ruprict/loccasions-go"
	"github.com/ruprict/loccasions-go/repository"
)

type OccasionsHandler struct {
	Repo repository.Repository
}

func (o *OccasionsHandler) GetOccasionsForEvent(c echo.Context) error {
	id := c.Param("event_id")
	cc := c.(*CustomContext)
	occasions := cc.Repo.GetOccasionsForEvent(id)
	return cc.JSONApi(200, occasions)
}

func (o *OccasionsHandler) PostOccasionToEvent(c echo.Context) error {
	var json loccasions.Occasion
	id := c.Param("event_id")
	json.OccurredOn = time.Now()
	cc := c.(*CustomContext)
	if c.Bind(&json) == nil {
		fmt.Println("event_id=", id)
		cc.Repo.AddOccasionToEvent(id, &json)
		return cc.JSONApi(201, &json)
	} else {
		return cc.JSON(500, map[string]string{"error": "error creating occasion"})
	}
}

func (o *OccasionsHandler) DeleteOccasion(c echo.Context) error {

	id := c.Param("id")
	fmt.Println("*** Deleting Occasion ", id)
	//TODO: Make sure user owns event

	cc := c.(*CustomContext)
	cc.Repo.DeleteOccasion(id)
	return c.NoContent(http.StatusNoContent)
}

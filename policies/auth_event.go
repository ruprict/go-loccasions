package policies

import (
	"net/http"

	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	loccasions "github.com/ruprict/loccasions-go"
	"github.com/ruprict/loccasions-go/handlers"
)

type AuthEvent struct{}

func Event(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*handlers.CustomContext)
		claims := cc.Get("user").(*jwt.Token).Claims.(*handlers.JwtCustomClaims)

		id_param := c.Param("id")
		if id_param == "" {
			id_param = c.Param("event_id")
		}

		if id_param == "" {
			log.Println("** No ID param")
			return next(c)
		}

		event := cc.Repo.GetEventForUser(claims.ID, id_param)
		if event == nil {
			log.Println("*** Event ", id_param, " not found for user: ", claims.ID)
			return cc.JSON(http.StatusNotFound, map[string]string{
				"error": "not found",
			})
		}

		log.Println("** returning event", event)
		cc.Events = []*loccasions.Event{event}

		return next(c)
	}

}

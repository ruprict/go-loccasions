package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	loccasions "github.com/ruprict/loccasions-go"
	"github.com/ruprict/loccasions-go/handlers"
	"github.com/ruprict/loccasions-go/repository"
)

var CurrentUser *loccasions.User
var EventsHandler *handlers.EventsHandler

func init() {
	EventsHandler = &handlers.EventsHandler{&repository.Postgres{}}
}

func main() {
	e := echo.New()
	e.Use(NewCors())
	e.Use(middleware.Logger())
	userHandler := createUserHandler()
	loginHandler := createLoginHandler()
	// Login route
	e.POST("/login", loginHandler.Login)
	e.POST("/register", userHandler.CreateUser)
	//e.GET("/profile", handlers.GetUser)
	r := e.Group("/events")
	config := middleware.JWTConfig{
		Claims:     &handlers.JwtCustomClaims{},
		SigningKey: []byte("sk00kum"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.POST("/:event_id/occasions", PostOccasionToEventEndpoint)
	r.GET("", EventsHandler.GetEvents)
	r.GET("/", EventsHandler.GetEvents)
	r.GET("/:id", GetEventEndpoint)
	r.POST("/", EventsHandler.CreateEvent)
	r.PATCH("/:id", PatchEventEndpoint)
	r.DELETE("/:id", DeleteEventEndpoint)
	r.GET("/:event_id/occasions", GetOccasionsForEventEndpoint)

	log.Fatal(e.Start(":8080")) // listen and serve on 0.0.0.0:8080

}

func GetEventEndpoint(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, map[string]string{"error": "bad id"})
	}
	event := repository.GetEvent(uint(id))
	if event == nil {
		return c.JSON(404, map[string]string{"error": "event not found"})
	} else {
		return c.JSON(200, event)
	}
}

func PostEventEndpoint(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*handlers.JwtCustomClaims)
	json := new(loccasions.Event)
	if err := c.Bind(json); err == nil {
		fmt.Println(json)
		repository.CreateEventForUser(claims.ID, json)

		return c.JSON(201, json)
	} else {
		return c.JSON(500, map[string]string{"error": "error creating event"})
	}
}

func PatchEventEndpoint(c echo.Context) error {
	var json loccasions.Event
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]string{"error": "bad id"})
	}
	if c.Bind(&json) == nil {
		repository.UpdateEvent(uint(id), &json)
		return c.JSON(200, json)
	} else {
		return c.JSON(500, map[string]string{"error": "error updating event"})
	}
}

func DeleteEventEndpoint(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]string{"error": "bad id"})
	}
	repository.DeleteEvent(uint(id))
	return c.NoContent(http.StatusNoContent)
}

func GetOccasionsForEventEndpoint(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("event_id"))
	if err != nil {
		return c.JSON(500, map[string]string{"error": "bad id"})
	}
	occasions := repository.GetOccasionsForEvent(uint(id))
	return c.JSON(200, occasions)
}

func PostOccasionToEventEndpoint(c echo.Context) error {
	var json loccasions.Occasion
	id, err := strconv.Atoi(c.Param("event_id"))
	json.OccurredOn = time.Now()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "bad id"})
	}
	if c.Bind(&json) == nil {
		repository.AddOccasionToEvent(uint(id), &json)
		return c.JSON(201, json)
	} else {
		return c.JSON(500, map[string]string{"error": "error creating occasion"})
	}
}

func createLoginHandler() *handlers.LoginHandler {
	return &handlers.LoginHandler{&repository.Postgres{}}
}

func createUserHandler() *handlers.UsersHandler {
	return &handlers.UsersHandler{&repository.Postgres{}}
}

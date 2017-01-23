package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	loccasions "github.com/ruprict/loccasions-go"
	"github.com/ruprict/loccasions-go/handlers"
	"github.com/ruprict/loccasions-go/policies"
	"github.com/ruprict/loccasions-go/repository"
)

var CurrentUser *loccasions.User
var EventsHandler *handlers.EventsHandler
var OccasionsHandler *handlers.OccasionsHandler

func init() {
	EventsHandler = &handlers.EventsHandler{}
	OccasionsHandler = &handlers.OccasionsHandler{}
}

func main() {
	e := echo.New()
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &handlers.CustomContext{c, &repository.Postgres{}, nil}
			return h(cc)
		}
	})
	e.Use(NewCors())
	e.Use(middleware.Logger())
	// Login route
	e.POST("/login", handlers.Login)
	e.POST("/register", handlers.CreateUser)
	//e.GET("/profile", handlers.GetUser)
	r := e.Group("/events")
	config := middleware.JWTConfig{
		Claims:     &handlers.JwtCustomClaims{},
		SigningKey: []byte("sk00kum"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.POST("/:event_id/occasions", OccasionsHandler.PostOccasionToEvent)
	r.GET("", EventsHandler.GetEvents)
	r.GET("/", EventsHandler.GetEvents)
	r.GET("/:id", policies.Event(EventsHandler.GetEvent))
	r.POST("/", EventsHandler.CreateEvent)
	r.PATCH("/:id", policies.Event(EventsHandler.PatchEvent))
	r.DELETE("/:id", policies.Event(EventsHandler.DeleteEvent))
	r.GET("/:event_id/occasions", OccasionsHandler.GetOccasionsForEvent)

	e.Logger.SetLevel(log.DEBUG)
	repository.DB.SetLogger(e.Logger)
	e.Logger.Fatal(e.Start(":8080")) // listen and serve on 0.0.0.0:8080

}

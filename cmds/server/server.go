package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	loccasions "github.com/ruprict/loccasions-go"
	"github.com/ruprict/loccasions-go/handlers"
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
			cc := &handlers.CustomContext{c, &repository.Postgres{}}
			return h(cc)
		}
	})
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
	r.POST("/:event_id/occasions", OccasionsHandler.PostOccasionToEvent)
	r.GET("", EventsHandler.GetEvents)
	r.GET("/", EventsHandler.GetEvents)
	r.GET("/:id", EventsHandler.GetEvent)
	r.POST("/", EventsHandler.CreateEvent)
	r.PATCH("/:id", EventsHandler.PatchEvent)
	r.DELETE("/:id", EventsHandler.DeleteEvent)
	r.GET("/:event_id/occasions", OccasionsHandler.GetOccasionsForEvent)

	log.Fatal(e.Start(":8080")) // listen and serve on 0.0.0.0:8080

}

func createLoginHandler() *handlers.LoginHandler {
	return &handlers.LoginHandler{&repository.Postgres{}}
}

func createUserHandler() *handlers.UsersHandler {
	return &handlers.UsersHandler{&repository.Postgres{}}
}

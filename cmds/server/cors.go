package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewCors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://ruprict.net", "http://localhost:8081", "http://0.0.0.0:3000", "http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	})
}

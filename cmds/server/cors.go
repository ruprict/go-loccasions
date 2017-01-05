package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewCors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://ruprict.net"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	})
}

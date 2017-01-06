package handlers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
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

type CustomContext struct {
	echo.Context
	Repo repository.Repository
}

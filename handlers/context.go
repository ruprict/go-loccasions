package handlers

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/jsonapi"
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

type CustomContext struct {
	echo.Context
	Repo   repository.Repository
	Events []*loccasions.Event
}

func (cc *CustomContext) JSONApi(code int, i interface{}) (err error) {

	err = jsonapi.MarshalPayload(cc.Response(), i)
	if err != nil {
		fmt.Println("** ERRor marshalling, ", err)
		return
	}
	return

}

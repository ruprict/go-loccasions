package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
	loccasions "github.com/ruprict/loccasions-go"
)

func CreateUser(c echo.Context) error {
	cc := c.(*CustomContext)
	email := cc.FormValue("email")
	password := cc.FormValue("password")
	name := cc.FormValue("name")
	user := loccasions.User{
		Email:    email,
		Password: salted(strings.TrimSpace(password)),
		Name:     name,
	}

	_, err := cc.Repo.CreateUser(&user)
	if err != nil {
		return cc.JSON(http.StatusBadRequest, &err)
	}

	return c.JSON(http.StatusCreated, &user)

}

func salted(pass string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(hashedPassword))
	return string(hashedPassword)
}

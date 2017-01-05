package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
	loccasions "github.com/ruprict/loccasions-go"
	"github.com/ruprict/loccasions-go/repository"
)

type UsersHandler struct {
	Repo repository.Repository
}

func (u *UsersHandler) CreateUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	fmt.Println("password = ", password)
	name := c.FormValue("name")
	user := loccasions.User{
		Email:    email,
		Password: salted(strings.TrimSpace(password)),
		Name:     name,
	}

	_, err := u.Repo.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &err)
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

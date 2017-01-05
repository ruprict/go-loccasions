package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/ruprict/loccasions-go/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	Repo repository.Repository
}

func (l *LoginHandler) Login(c echo.Context) error {

	email := c.FormValue("email")
	password := strings.TrimSpace(c.FormValue("password"))

	user := l.Repo.GetUserForEmail(email)
	if user == nil {
		fmt.Println("**** user not found")
		return echo.ErrNotFound
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["admin"] = false
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("sk00kum"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

	return echo.ErrUnauthorized

}

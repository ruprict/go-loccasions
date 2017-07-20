package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {

	cc := c.(*CustomContext)
	email := c.FormValue("email")
	fmt.Println(cc.FormParams())

	password := strings.TrimSpace(c.FormValue("password"))

	user := cc.Repo.GetUserForEmail(email)
	fmt.Println(user.ID)
	if user.ID == "" {
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
	return cc.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}

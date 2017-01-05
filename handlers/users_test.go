package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	uuid "github.com/nu7hatch/gouuid"
	loccasions "github.com/ruprict/loccasions-go"
	"github.com/ruprict/loccasions-go/handlers"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]*loccasions.User{
		"glenn@skookum.com": loccasions.NewUser("Glenn Goodrich", "glenn@skookum.com", "password"),
	}
)

type mockRepo struct{}

func (m mockRepo) CreateUser(user *loccasions.User) (string, error) {
	uuid, _ := uuid.NewV4()
	user.ID = uuid.String()
	return user.ID, nil
}

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("name", "Glenn Goodrich")
	f.Set("email", "glenn@skookum.com")
	f.Set("password", "password")
	req, err := http.NewRequest(echo.POST, "/register", strings.NewReader(f.Encode()))
	if assert.NoError(t, err) {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &handlers.UsersHandler{mockRepo{}}

		// Assertions
		if assert.NoError(t, h.CreateUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	}
}

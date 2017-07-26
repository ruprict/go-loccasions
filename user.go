package loccasions

import (
	"encoding/json"
	"time"
)

//User is our model object for profiles
type User struct {
	Email     string `sql:"unique" jsonapi:"attr,email"`
	Name      string `jsonapi:"attr,name"`
	Password  string
	Events    []*Event `jsonapi:"relation,events"`
	ID        string   `sql:"type:uuid;primary_key;default:uuid_generate_v4()" jsonapi:"primary,users"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// NewUser initializes a new User object
func NewUser(email string, name string, password string) *User {
	u := new(User)
	u.Email = email
	u.Name = name
	u.Password = password
	return u
}

// MarshalJSON customizes the user JSON, only including email, name, and id
func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		email string `json:"email"`
		name  string `json:"name"`
		id    string `json:"id"`
	}{
		email: u.Email,
		name:  u.Name,
		id:    u.ID,
	})
}

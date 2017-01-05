package repository

import loccasions "github.com/ruprict/loccasions-go"

type Repository interface {
	CreateUser(user *loccasions.User) (string, error)
	GetUserForEmail(email string) *loccasions.User
	CreateEventForUser(user_id string, event *loccasions.Event) (uint, error)
	GetEventsForUser(user_id string) *[]loccasions.Event
}

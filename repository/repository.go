package repository

import loccasions "github.com/ruprict/loccasions-go"

type Repository interface {
	CreateUser(user *loccasions.User) (string, error)
	GetUserForEmail(email string) *loccasions.User
	CreateEventForUser(user_id string, event *loccasions.Event) (string, error)
	GetEventsForUser(user_id string) []loccasions.Event
	GetEventForUser(user_id string, id string) *loccasions.Event
	DeleteEvent(id string) error
	UpdateEvent(id string, event *loccasions.Event) (*loccasions.Event, error)
	GetOccasionsForEvent(event_id string) []loccasions.Occasion
	AddOccasionToEvent(event_id string, occasion *loccasions.Occasion) (string, error)
}

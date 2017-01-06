package loccasions

import "time"

//TODO: Write MarshalJSON to return link to User and Occasions
// Event is the model object reprsenting our Event
type Event struct {
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description"`
	Occasions   []Occasion `json:"occasions"`
	UserID      string     `sql:"type:uuid"`
	ID          string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

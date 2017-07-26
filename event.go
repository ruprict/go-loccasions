package loccasions

import (
	"fmt"
	"time"

	"github.com/google/jsonapi"
)

//TODO: Write MarshalJSON to return link to User and Occasions
// Event is the model object reprsenting our Event
type Event struct {
	Name        string      `jsonapi:"attr,name" validate:"required"`
	Description string      `jsonapi:"attr,description"`
	Occasions   []*Occasion `jsonapi:"relation,occasions"`
	UserID      string      `sql:"type:uuid" gorm:"index:idx_user_id" jsonapi:"attr,ownerID"`
	ID          string      `sql:"type:uuid;primary_key;default:uuid_generate_v4()" jsonapi:"primary,events"`
	CreatedAt   time.Time   `jsonapi: "attr,createdAt"`
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (event Event) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": jsonapi.Link{
			Href: fmt.Sprintf("https://api.loccasions.dev/events/%s", event.ID),
		},
		"occasions": jsonapi.Link{
			Href: fmt.Sprintf("https://api.loccasions.dev/events/%s/occasions", event.ID),
			Meta: map[string]interface{}{
				"counts": map[string]uint{
					"likes": 4,
				},
			},
		},
	}
}

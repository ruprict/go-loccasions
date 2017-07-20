package loccasions

import (
	"encoding/json"
	"fmt"
	"time"
)

//TODO: Write MarshalJSON to return link to User and Occasions
// Event is the model object reprsenting our Event
type Event struct {
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description"`
	Occasions   []Occasion `json:"occasions"`
	UserID      string     `sql:"type:uuid" gorm:"index:idx_user_id" json:"userID"`
	ID          string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (u Event) MarshalJSON() ([]byte, error) {
	var occs []Occasion
	occs = u.Occasions
	if len(u.Occasions) == 0 {
		fmt.Println("shiza ***")
		occs = make([]Occasion, 0)
	}
	return json.Marshal(&struct {
		Name        string     `json:"name"`
		Description string     `json:"description"`
		ID          string     `json:"id"`
		Occasions   []Occasion `json:"occasions"`
	}{
		Name:        u.Name,
		Description: u.Description,
		ID:          u.ID,
		Occasions:   occs,
	})

}

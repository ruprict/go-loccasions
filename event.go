package loccasions

import "github.com/jinzhu/gorm"

// Event is the model object reprsenting our Event
type Event struct {
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description"`
	Occasions   []Occasion `json:"occasions"`
	UserID      string     `sql:"type:uuid"`
	gorm.Model
}

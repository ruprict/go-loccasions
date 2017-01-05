package loccasions

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/nferruzzi/gormGIS"
)

// Occasion is the model object representing an occurance of an Event
type Occasion struct {
	OccurredOn time.Time        `json:"occurredOn"`
	Note       string           `json:"note"`
	Location   gormGIS.GeoPoint `json:"location" sql:"type:geometry(Geometry,4326)"`
	EventID    uint             `json:"eventId"`
	gorm.Model
}

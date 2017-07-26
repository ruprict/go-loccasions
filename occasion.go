package loccasions

import (
	"encoding/json"
	"time"

	"github.com/nferruzzi/gormGIS"
)

//TODO: Write MarshalJSON to return link to Event
// Occasion is the model object representing an occurance of an Event
type Occasion struct {
	OccurredOn time.Time        `json:"occurredOn" jsonapi:"attr,occurredOn"`
	Note       string           `json:"note" jsonapi:"attr,note"`
	Location   gormGIS.GeoPoint `json:"location" sql:"type:geometry(Geometry,4326)" jsonapi:"attr,location"`
	EventID    string           `json:"eventId" sql:"type:uuid"`
	ID         string           `sql:"type:uuid;primary_key;default:uuid_generate_v4()" jsonapi:"primary,occasions"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type geoJSONFeature struct {
	Type       string                 `json:"type" jsonapi:"attr,type"`
	Properties map[string]interface{} `json:"properties" jsonapi:"attr,properties"`
	Geometry   map[string]interface{} `json:"geometry"`
}

func (o *Occasion) MarshalJSON() ([]byte, error) {
	geojson := geoJSONFeature{
		Type: "Feature",
		Properties: map[string]interface{}{
			"note":       o.Note,
			"created_at": o.CreatedAt,
			"id":         o.ID,
		},
		Geometry: map[string]interface{}{
			"type":        "Point",
			"coordinates": []float64{o.Location.Lng, o.Location.Lat},
		},
	}
	return json.Marshal(&geojson)

}

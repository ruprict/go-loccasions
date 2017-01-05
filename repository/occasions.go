package repository

import loccasions "github.com/ruprict/loccasions-go"

func GetOccasionsForEvent(event_id uint) *[]loccasions.Occasion {
	var occasions []loccasions.Occasion
	event := GetEvent(event_id)
	DB.Find(&event).Related(&occasions)
	return &occasions
}

func AddOccasionToEvent(event_id uint, occasion *loccasions.Occasion) (uint, error) {
	event := GetEvent(event_id)
	DB.Model(&event).Association("Occasions").Append(occasion)
	return occasion.ID, nil

}

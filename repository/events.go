package repository

import loccasions "github.com/ruprict/loccasions-go"

func GetEventsForUser(user_id string) *[]loccasions.Event {
	var events []loccasions.Event
	u := new(loccasions.User)
	u.ID = user_id
	DB.Model(&u).Association("Events").Find(&events)

	return &events
}

func GetEvent(id uint) *loccasions.Event {
	if id > 1 {
		return nil
	}
	var event loccasions.Event
	DB.Find(&event, id)
	return &event

}

func CreateEventForUser(user_id string, event *loccasions.Event) (uint, error) {
	event.UserID = user_id
	DB.Create(event)
	return event.ID, nil
}

func UpdateEvent(id uint, event *loccasions.Event) (*loccasions.Event, error) {
	event.ID = id
	DB.Save(&event)
	return event, nil
}

func DeleteEvent(id uint) error {
	event := loccasions.Event{}
	event.ID = id
	DB.Delete(&event)
	return nil
}

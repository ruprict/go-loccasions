package repository

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	loccasions "github.com/ruprict/loccasions-go"
)

type Postgres struct{}

func init() {
	db, err := gorm.Open("postgres", "postgresql://docker:docker@localhost:5433/gis?sslmode=disable")
	db.LogMode(true)
	if err == nil {
		log.Println("*** DB migrating....")
		db.AutoMigrate(&loccasions.User{}, &loccasions.Event{}, &loccasions.Occasion{})
		DB = Database{db}
	} else {
		log.Fatalln(err)
	}

}
func (p *Postgres) CreateUser(user *loccasions.User) (string, error) {
	db := DB.Create(user)
	if db.Error != nil {
		return "", db.Error
	}
	return user.ID, nil
}
func (p *Postgres) GetUserForEmail(email string) *loccasions.User {
	var user loccasions.User
	DB.Where("email = ?", email).First(&user)
	return &user
}

func (p *Postgres) GetEventsForUser(user_id string) []*loccasions.Event {
	fmt.Println("GET some mofoocking events")
	var events []*loccasions.Event
	u := new(loccasions.User)
	u.ID = user_id
	DB.Model(&u).Association("Events").Find(&events)

	return events
}

func (p *Postgres) CreateEventForUser(user_id string, event *loccasions.Event) (string, error) {
	event.UserID = user_id
	DB.Create(event)
	return event.ID, nil
}

func (p *Postgres) GetEventForUser(user_id string, id string) *loccasions.Event {
	var events []loccasions.Event

	DB.Preload("Occasions").Where("user_id=? and id=?", user_id, id).First(&events)

	if len(events) == 1 {
		return &events[0]
	}

	return nil
}
func (p *Postgres) GetEvent(id string) *loccasions.Event {
	var event loccasions.Event
	DB.Where(&loccasions.Event{ID: id}).First(&event)
	if DB.Error != nil {
		fmt.Println("*** Error finding event: ", DB.Error)
	}

	return &event
}
func (p *Postgres) UpdateEvent(id string, event *loccasions.Event) (*loccasions.Event, error) {
	ev := p.GetEvent(id)
	err := DB.Model(ev).Update(event).Error
	if err != nil {
		fmt.Println("** WHA THE FOOOK", err)
		return nil, err
	}
	return ev, nil
}
func (p *Postgres) DeleteEvent(id string) error {
	event := p.GetEvent(id)
	DB.Delete(&event)
	return nil
}

func (p *Postgres) GetOccasionsForEvent(event_id string) []loccasions.Occasion {
	var occasions []loccasions.Occasion
	event := p.GetEvent(event_id)
	DB.Model(&event).Related(&occasions)
	return occasions
}

func (p *Postgres) AddOccasionToEvent(event_id string, occasion *loccasions.Occasion) (string, error) {
	fmt.Println("***")
	fmt.Println(event_id)
	fmt.Println("***")
	event := p.GetEvent(event_id)
	fmt.Println(event.ID)
	fmt.Println("***")
	DB.Model(&event).Association("Occasions").Append(occasion)
	return occasion.ID, nil

}

func (p *Postgres) GetOccasion(id string) *loccasions.Occasion {
	var occ loccasions.Occasion
	DB.Where(&loccasions.Occasion{ID: id}).First(&occ)
	if DB.Error != nil {
		fmt.Println("*** Error finding occcasion: ", DB.Error)
	}

	fmt.Println("** Found occasion ", occ)
	return &occ
}

func (p *Postgres) DeleteOccasion(id string) error {
	occ := p.GetOccasion(id)
	fmt.Println("*** About to delete occasion ", occ.ID)
	DB.Delete(&occ)
	fmt.Println("*** deleted")
	return nil
}

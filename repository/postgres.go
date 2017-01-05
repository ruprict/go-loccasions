package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	loccasions "github.com/ruprict/loccasions-go"
)

type Postgres struct{}

func init() {
	db, err := gorm.Open("postgres", "postgresql://docker:docker@localhost/gis?sslmode=disable")
	if err == nil {
		log.Println("*** DB migrating....")
		db.AutoMigrate(&loccasions.Occasion{}, &loccasions.Event{}, &loccasions.User{})
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

func (p *Postgres) GetEventsForUser(user_id string) *[]loccasions.Event {
	var events []loccasions.Event
	u := new(loccasions.User)
	u.ID = user_id
	DB.Model(&u).Association("Events").Find(&events)

	return &events
}

func (p *Postgres) CreateEventForUser(user_id string, event *loccasions.Event) (uint, error) {
	event.UserID = user_id
	DB.Create(event)
	return event.ID, nil
}

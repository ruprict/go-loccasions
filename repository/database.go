package repository

import (
	"github.com/jinzhu/gorm"
)

type Database struct {
	*gorm.DB
}

var DB Database

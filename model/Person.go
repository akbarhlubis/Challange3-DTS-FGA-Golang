package model

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
}

type InDB struct {
	DB *gorm.DB
}

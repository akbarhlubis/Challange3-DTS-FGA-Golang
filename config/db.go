package config

import (
	"challange3_17Mar/model"
	"log"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db_challange2_go")
	if err != nil {
		log.Fatalf("Failed to connect to databases: %v", err)
	}
	db.AutoMigrate(model.Person{})
	return db
}

package model

import (
	"github.com/jinzhu/gorm"
)

// this is how the db will be modeled
type User struct {
	//email with max length of 255, set as PRIMARY KEY, must be UNIQUE
	Email string `gorm:"primary_key; size:255; unique; not null"`
	//name with max length of 255
	Name string `gorm:"size:255; not null"`
	//user level as an int
	userLevel int `gorm:"not null"`
	//password with max length of 25
	Password string `gorm:"not null; size:25"`
}

func migrateDB(db *gorm.DB) *gorm.DB {
	//format db to replicate user struct
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic("Error in migrating db")
	}

	return db
}

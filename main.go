package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

var (
	db  *gorm.DB
	err error
)

func connectAndMigrateDB() {
	//establish connection to db
	db, err = gorm.Open(sqlite.Open("usersDB"), &gorm.Config{})
	if err != nil {
		panic("Error in opening DB")
	}

	//format db to replicate user struct
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Error in migrating db")
	}
}

func main() {
	fmt.Println("Main entered")
	fmt.Println("Set up Complete")
	return
}

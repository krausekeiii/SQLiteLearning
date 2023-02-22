package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Main entered")
	//establish connection to db
	db, err := gorm.Open(sqlite.Open("usersDB"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in opening DB")
		return
	}

	//format db to replicate user struct
	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("Error in migrating db")
		return
	}
	fmt.Println("Set up Complete")
	return
}

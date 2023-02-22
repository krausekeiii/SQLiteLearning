package main

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

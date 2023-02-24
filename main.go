package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/krausekeiii/golang-gorm-sqlite/src/api/app"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	err    error
	router *mux.Router
)

func main() {
	fmt.Println("Main entered")
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
	return
}

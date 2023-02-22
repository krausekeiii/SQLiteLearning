package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	err    error
	router *mux.Router
)

func main() {
	fmt.Println("Main entered")

	return
}

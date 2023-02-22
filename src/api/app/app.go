package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	db     *gorm.DB
}

func (a *App) Initialize() {
	//establish connection to db
	db, err := gorm.Open(sqlite.Open("usersDB"), &gorm.Config{})
	if err != nil {
		panic("Error in opening DB")
	}

	a.db = model.migrateDB(db)
	a.Router = mux.NewRouter()
	//this will allow us to update API calls easily
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/employees", a.GetAllEmployees)
	a.Post("/employees", a.CreateEmployee)
	a.Get("/employees/{title}", a.GetEmployee)
	a.Put("/employees/{title}", a.UpdateEmployee)
	a.Delete("/employees/{title}", a.DeleteEmployee)
	a.Put("/employees/{title}/disable", a.DisableEmployee)
	a.Put("/employees/{title}/enable", a.EnableEmployee)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

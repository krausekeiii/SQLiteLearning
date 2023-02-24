package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/krausekeiii/golang-gorm-sqlite/src/api/app/handler"
	"github.com/krausekeiii/golang-gorm-sqlite/src/api/app/model"
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
	//calls MigrateDB in migrate which uses AutoMigrate and throws error if cannot migrate
	a.db = model.MigrateDB(db)
	//creating new router
	a.Router = mux.NewRouter()
	//this will allow us to update API calls easily
	a.setRouters()

}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/users", a.GetAllUsers)
	a.Post("/users", a.CreateUser)
	//a.Get("/users/{title}", a.GetUser)
	//a.Put("/users/{title}", a.UpdateUser)
	//a.Delete("/users/{title}", a.DeleteEmployee)
	//a.Put("/users/{title}/disable", a.DisableEmployee)
	//a.Put("/users/{title}/enable", a.EnableEmployee)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.db, w, r)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.CreateUser(a.db, w, r)
}

/*func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {

}*/

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

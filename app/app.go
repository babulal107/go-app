package app

import (
	"bitbucket.org/go-api/app/handler"
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/go-api/app/model"
	"bitbucket.org/go-api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)
 
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}
 
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}
 
// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.setUserRouters()
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
 
// Handlers to manage Employee Data
func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}
 
func (a *App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.CreateUser(a.DB, w, r)
}
 
func (a *App) GetEmployee(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(a.DB, w, r)
}
 
func (a *App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUser(a.DB, w, r)
}
 
func (a *App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(a.DB, w, r)
}
 
func (a *App) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DisableUser(a.DB, w, r)
}
 
func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.EnableUser(a.DB, w, r)
}
 
// Run the app on it's router
func (a *App) Run(host string) {
	log.Printf("Starting server at port %v", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}

package app

import (
	"bitbucket.org/go-app/app/handler"
	"net/http"
)

// Set all required routers
func (a *App) setUserRouters() {
	// Routing for handling the projects
	a.Get("/users", a.GetAllUsers)
	a.Post("/users", a.CreateUser)
	a.Get("/users/{title}", a.GetUser)
	a.Put("/users/{title}", a.UpdateUser)
	a.Delete("/users/{title}", a.DeleteUser)
	a.Put("/users/{title}/disable", a.DisableUser)
	a.Put("/users/{title}/enable", a.EnableUser)
}


// Handlers to manage user Data
func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.CreateUser(a.DB, w, r)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(a.DB, w, r)
}

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUser(a.DB, w, r)
}

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(a.DB, w, r)
}

func (a *App) DisableUser(w http.ResponseWriter, r *http.Request) {
	handler.DisableUser(a.DB, w, r)
}

func (a *App) EnableUser(w http.ResponseWriter, r *http.Request) {
	handler.EnableUser(a.DB, w, r)
}
package app

import (
	"log"
	"net/http"
)

// App struct
type App struct {
	Router *Router
}

// Init app
func (app *App) Init() {
	app.Router = &Router{}
	app.Router.Init()
}

// Run app
func (app *App) Run() {
	log.Fatal(
		http.ListenAndServe(
			":3000",
			app.Router.Mux,
		),
	)
}

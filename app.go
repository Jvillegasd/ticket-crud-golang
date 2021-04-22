package main

import (
	"log"
	"net/http"

	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}

func (app *App) initRoutes() {
	app.Router = mux.NewRouter()

}

func (app *App) initDatabase(user, password, dbname string) {
	
}

func (app *App) run() {
	log.Fatal(http.ListenAndServe(":2600", app.Router))
}
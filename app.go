package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) InitRoutes() {
	app.Router = mux.NewRouter()

}

func (app *App) InitDatabase(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	fmt.Println(connectionString)
	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":2600", app.Router))
}

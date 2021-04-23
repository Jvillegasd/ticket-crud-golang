package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) InitRoutes() {
	app.Router = mux.NewRouter()

	app.Router.HandleFunc("/", app.ping)
	app.Router.HandleFunc("/tickets", app.getAllTickets).Methods("GET")
	app.Router.HandleFunc("/tickets", app.createTicket).Methods("POST")
	app.Router.HandleFunc("/tickets/{id:[0-9]+}", app.getTicket).Methods("GET")
	app.Router.HandleFunc("/tickets/{id:[0-9]+}", app.updateTicket).Methods("PUT")
	app.Router.HandleFunc("/tickets/{id:[0-9]+}", app.deleteTicket).Methods("DELETE")
}

func (app *App) InitDatabase(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":2600", app.Router))
}

func (app *App) ping(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"result": "server's up"})
	return
}

func (app *App) getTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ticket Id")
		return
	}

	t := ticket{ID: id}
	if err := t.getTicket(app.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Ticket not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJson(w, http.StatusOK, t)
}

func (app *App) getAllTickets(w http.ResponseWriter, r *http.Request) {
	count := 10
	start, _ := strconv.Atoi(r.FormValue("start"))

	if start < 0 {
		start = 0
	}

	tickets, err := getAllTickets(app.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, tickets)
}

func (app *App) createTicket(w http.ResponseWriter, r *http.Request) {
	var t ticket
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := t.createTicket(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, t)
}

func (app *App) updateTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ticket Id")
		return
	}

	var t ticket
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	t.ID = id

	if err := t.updateTicket(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, t)
}

func (app *App) deleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ticket Id")
		return
	}

	t := ticket{ID: id}
	if err := t.deleteTicket(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

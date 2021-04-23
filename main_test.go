package main_test

import (
	"log"
	"os"
	"testing"

	"prueba_tecnica"
)

var a main.App

func TestMain(m *testing.M) {
	a.InitDatabase(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
		
	a.InitRoutes()
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM tickets")
	a.DB.Exec("ALTER SEQUENCE tickets_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS tickets
(
	id SERIAL,
	user TEXT NOT NULL,
	created_at DATETIME,
	updated_at DATETIME,
	status BOOLEAN,
	CONSTRAIN tickets_pkey PRIMARY KEY (id)
)`

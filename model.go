package main

import (
	"database/sql"
	"time"
)

type ticket struct {
	ID        int       `json:"id"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    bool      `json:"status"`
}

func (t *ticket) getTicket(db *sql.DB) error {
	return db.QueryRow("SELECT id, user, status FROM tickets WHERE id=$1",
		t.ID).Scan(&t.ID, &t.User, &t.Status)
}

func getAllTickets(db *sql.DB, start, count int) ([]ticket, error) {
	rows, err := db.Query(
		"SELECT id user, status FROM tickets LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tickets := []ticket{}
	for rows.Next() {
		var t ticket
		if err := rows.Scan(&t.ID, &t.User, &t.Status); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	return tickets, nil
}

func (t *ticket) createTicket(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO tickets(user, status) VALUES($1, $2) RETURNING id",
		t.User, t.Status).Scan(&t.ID)

	if err != nil {
		return err
	}

	return nil
}

func (t *ticket) updateTicket(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE tickets SET User user=$1 status=$2 WHERE id=$3",
		t.User, t.Status, t.ID)

	return err
}

func (t *ticket) deleteTicket(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM tickets WHERE id=$id", t.ID)

	return err
}

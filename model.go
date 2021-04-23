package main

import (
	"database/sql"
	"time"
)

type ticket struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	TicketStatus bool   `json:"ticket_status"`
}

func (t *ticket) getTicket(db *sql.DB) error {
	return db.QueryRow("SELECT id, username, ticket_status, created_at, updated_at FROM tickets WHERE id=$1",
		t.ID).Scan(&t.ID, &t.Username, &t.TicketStatus, &t.CreatedAt, &t.UpdatedAt)
}

func getAllTickets(db *sql.DB, start, count int) ([]ticket, error) {
	rows, err := db.Query(
		"SELECT id, username, ticket_status, created_at, updated_at FROM tickets LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tickets := []ticket{}
	for rows.Next() {
		var t ticket
		if err := rows.Scan(&t.ID, &t.Username, &t.TicketStatus, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	return tickets, nil
}

func (t *ticket) createTicket(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO tickets(username, ticket_status) VALUES($1, $2) RETURNING id",
		t.Username, t.TicketStatus).Scan(&t.ID)

	if err != nil {
		return err
	}

	return nil
}

func (t *ticket) updateTicket(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE tickets SET username=$1, ticket_status=$2, updated_at=$3 WHERE id=$4",
		t.Username, t.TicketStatus, time.Now(), t.ID)

	return err
}

func (t *ticket) deleteTicket(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM tickets WHERE id=$1", t.ID)

	return err
}

package main

import (
	"database/sql"
	"errors"
	"time"
)

type ticket struct {
	ID        int       `json:"id"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	status    bool      `json:"status"`
}

func (t *ticket) getTicket(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (t *ticket) getAllTickets(db *sql.DB) ([]ticket, error) {
	return nil, errors.New("Not implemented")
}

func (t *ticket) createTicket(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (t *ticket) updateTicket(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (t *ticket) editTicket(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (t *ticket) deleteTicket(db *sql.DB) error {
	return errors.New("Not implemented")
}

CREATE TABLE IF NOT EXISTS tickets (
	id SERIAL PRIMARY KEY,
	username TEXT NOT NULL,
	created_at timestamp DEFAULT now(),
	updated_at timestamp DEFAULT now(),
	ticket_status BOOLEAN
);
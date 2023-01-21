-- +goose Up
CREATE TABLE IF NOT EXISTS clients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    email TEXT
);

CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    category TEXT,
    client_id INT,
    title TEXT,
    starts_at TIMESTAMP,
    ends_at TIMESTAMP,
    CONSTRAINT fk_clients
    FOREIGN KEY (client_id)
    REFERENCES clients(id)
);
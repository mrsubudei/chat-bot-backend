-- +goose Up
CREATE TABLE IF NOT EXISTS clients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    email TEXT
);

CREATE TABLE IF NOT EXISTS events (
    id SERIAL,
    category TEXT,
    client_id INT,
    title TEXT,
    starts_at TIMESTAMP,
    ends_at TIMESTAMP,
    docktor_id INT,
    PRIMARY KEY (starts_at, ends_at, docktor_id),
    CONSTRAINT fk_clients
        FOREIGN KEY (client_id)
            REFERENCES clients(id),
        FOREIGN KEY (docktor_id)
            REFERENCES doctors(id),
);

CREATE TABLE IF NOT EXISTS doctors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    phone TEXT
);
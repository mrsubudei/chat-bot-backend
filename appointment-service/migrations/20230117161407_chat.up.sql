-- +goose Up
SET TIMEZONE='Asia/Almaty';

CREATE TABLE IF NOT EXISTS clients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    email TEXT,
    events INT[]
);

CREATE TABLE IF NOT EXISTS doctors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    phone TEXT
);

CREATE TABLE IF NOT EXISTS events (
    id SERIAL,
    client_id INT,
    doctor_id INT,
    starts_at TIMESTAMP,
    ends_at TIMESTAMP,
    PRIMARY KEY (doctor_id, starts_at, ends_at),
    CONSTRAINT fk_clients
        FOREIGN KEY (client_id)
            REFERENCES clients(id),
        FOREIGN KEY (doctor_id)
            REFERENCES doctors(id)
);
SET TIMEZONE='Asia/Almaty';

CREATE TABLE IF NOT EXISTS doctors (
    id SERIAL PRIMARY KEY UNIQUE,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    phone TEXT,
    CONSTRAINT doctor_name_surname_unique UNIQUE (name, surname)
);

CREATE TABLE IF NOT EXISTS events (
    id SERIAL UNIQUE,
    client_id INT NULL,
    doctor_id INT NOT NULL,
    starts_at TIMESTAMP,
    ends_at TIMESTAMP,
    PRIMARY KEY (doctor_id, starts_at, ends_at),
    CONSTRAINT fk_doctors
        FOREIGN KEY (doctor_id)
            REFERENCES doctors(id)
                ON DELETE CASCADE
);

CREATE INDEX "events_client_id"
ON events(client_id);
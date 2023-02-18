SET TIMEZONE='Asia/Almaty';

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY UNIQUE,
    name TEXT NOT NULL,
    phone TEXT NOT NULL UNIQUE,
    email TEXT,
    password TEXT NOT NULL,
    role TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions (
    user_id INT UNIQUE,
    session_token TEXT,
    session_ttl TIMESTAMP,
    CONSTRAINT fk_users
        FOREIGN KEY (user_id)
            REFERENCES users(id)
                ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS verifications (
    user_id INT UNIQUE,
    sms_code TEXT,
    verified BOOLEAN NOT NULL,
    CONSTRAINT fk_users
        FOREIGN KEY (user_id)
            REFERENCES users(id)
                ON DELETE CASCADE
);
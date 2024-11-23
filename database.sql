DROP TABLE IF EXISTS pin;
CREATE TABLE pin (
    id SERIAL PRIMARY KEY,
    pin int NOT NULL,
    guess int NOT NULL
);

DROP TABLE IF EXISTS pin_guess;
CREATE TABLE pin_guess (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    pin int NOT NULL,
    guess_time date
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

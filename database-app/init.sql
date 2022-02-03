DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) NOT NULL UNIQUE,
    password  VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL
);


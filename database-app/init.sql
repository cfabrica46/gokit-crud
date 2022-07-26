DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(64) NOT NULL UNIQUE,
    password  VARCHAR(128) NOT NULL,
    email VARCHAR(64) NOT NULL UNIQUE
);

INSERT INTO users(username, password,email)
    VALUES
        ('cesar',	'01234',	'cesar@gmail.com'),
        ('luis',	'12345',	'luis@gmail.com')

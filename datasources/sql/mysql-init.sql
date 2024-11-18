-- name: create-user-table
CREATE TABLE IF NOT EXISTS users (
    id int NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    role INT NOT NULL DEFAULT 0,
    PRIMARY KEY(id)
);
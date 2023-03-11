-- +goose Up
CREATE TABLE auth_user
(
    email TEXT PRIMARY KEY
);

CREATE TABLE login_pwd
(
    id         TEXT PRIMARY KEY,
    name       TEXT NOT NULL,
    metadata   TEXT NOT NULL,
    login      TEXT NOT NULL,
    password   TEXT NOT NULL,
    user_email TEXT NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_email)
        REFERENCES auth_user (email)
);

CREATE TABLE card_info
(
    id               TEXT PRIMARY KEY,
    name             TEXT NOT NULL,
    metadata         TEXT NOT NULL,
    card_number      TEXT NOT NULL,
    cvv              TEXT NOT NULL,
    expiration_month TEXT NOT NULL,
    expiration_year  TEXT NOT NULL,
    user_email       TEXT NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_email)
        REFERENCES auth_user (email)
);

CREATE TABLE text_info
(
    id         TEXT PRIMARY KEY,
    name       TEXT NOT NULL,
    metadata   TEXT NOT NULL,
    text       TEXT NOT NULL,
    user_email TEXT NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_email)
        REFERENCES auth_user (email)
);

-- +goose Down
DROP TABLE text_info;
DROP TABLE card_info;
DROP TABLE login_pwd;

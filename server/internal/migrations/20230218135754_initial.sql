-- +goose Up
CREATE TABLE auth_user
(
    email    VARCHAR(32) PRIMARY KEY,
    password VARCHAR(128) NOT NULL
);

CREATE TABLE login_pwd
(
    id         serial PRIMARY KEY,
    name       VARCHAR(64),
    metadata   TEXT         NOT NULL,
    user_email VARCHAR(32),
    login      VARCHAR(128) NOT NULL,
    password   VARCHAR(128) NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_email)
        REFERENCES auth_user (email)
);

CREATE TABLE card_info
(
    id               serial PRIMARY KEY,
    name             VARCHAR(64),
    metadata         TEXT NOT NULL,
    user_email       VARCHAR(32),
    card_number      VARCHAR(16),
    cvv              VARCHAR(4),
    expiration_month integer CHECK (expiration_month >= 1 AND expiration_month >= 12),
    expiration_year  integer CHECK (expiration_year >= 1900 AND expiration_year <= 2500),
    CONSTRAINT fk_user FOREIGN KEY (user_email)
        REFERENCES auth_user (email)
);

CREATE TABLE text_info
(
    id         serial PRIMARY KEY,
    name       VARCHAR(64),
    metadata   TEXT NOT NULL,
    user_email VARCHAR(32),
    text       TEXT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_email)
        REFERENCES auth_user (email)
);

-- +goose Down
DROP TABLE text_info;
DROP TABLE card_info;
DROP TABLE login_pwd;
DROP TABLE auth_user;

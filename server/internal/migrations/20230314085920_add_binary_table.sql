-- +goose Up
CREATE TABLE binary_info
(
    id         uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name       VARCHAR(64),
    metadata   TEXT NOT NULL,
    user_email VARCHAR(32),
    data       TEXT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_email)
        REFERENCES auth_user (email)
);

-- +goose Down
DROP TABLE binary_info;

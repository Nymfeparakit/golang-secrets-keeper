-- +goose Up
ALTER TABLE login_pwd ADD COLUMN updated_at timestamptz default current_timestamp;
ALTER TABLE card_info ADD COLUMN updated_at timestamptz default current_timestamp;
ALTER TABLE text_info ADD COLUMN updated_at timestamptz default current_timestamp;

-- +goose Down
ALTER TABLE login_pwd DROP COLUMN updated_at;
ALTER TABLE card_info DROP COLUMN updated_at;
ALTER TABLE text_info DROP COLUMN updated_at;
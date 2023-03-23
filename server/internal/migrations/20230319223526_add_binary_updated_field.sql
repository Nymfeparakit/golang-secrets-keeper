-- +goose Up
ALTER TABLE binary_info ADD COLUMN updated_at timestamptz default current_timestamp;

-- +goose Down
ALTER TABLE binary_info DROP COLUMN updated_at;

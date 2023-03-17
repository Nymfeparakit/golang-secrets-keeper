-- +goose Up
ALTER TABLE login_pwd ADD COLUMN deleted BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE card_info ADD COLUMN deleted BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE binary_info ADD COLUMN deleted BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE text_info ADD COLUMN deleted BOOLEAN NOT NULL DEFAULT false;

-- +goose Down
ALTER TABLE login_pwd DROP COLUMN deleted;
ALTER TABLE card_info DROP COLUMN deleted;
ALTER TABLE binary_info DROP COLUMN deleted;
ALTER TABLE text_info DROP COLUMN deleted;

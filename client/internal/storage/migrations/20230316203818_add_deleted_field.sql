-- +goose Up
ALTER TABLE login_pwd ADD COLUMN deleted INTEGER NOT NULL DEFAULT 0 CHECK (deleted IN (0, 1));
ALTER TABLE card_info ADD COLUMN deleted INTEGER NOT NULL DEFAULT 0 CHECK (deleted IN (0, 1));
ALTER TABLE binary_info ADD COLUMN deleted INTEGER NOT NULL DEFAULT 0 CHECK (deleted IN (0, 1));
ALTER TABLE text_info ADD COLUMN deleted INTEGER NOT NULL DEFAULT 0 CHECK (deleted IN (0, 1));

-- +goose Down
ALTER TABLE login_pwd DROP COLUMN deleted;
ALTER TABLE card_info DROP COLUMN deleted;
ALTER TABLE binary_info DROP COLUMN deleted;
ALTER TABLE text_info DROP COLUMN deleted;

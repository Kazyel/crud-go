-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN IF NOT EXISTS last_updated TIMESTAMP
WITH
  TIME ZONE DEFAULT CURRENT_TIMESTAMP;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN IF EXISTS last_updated;

-- +goose StatementEnd
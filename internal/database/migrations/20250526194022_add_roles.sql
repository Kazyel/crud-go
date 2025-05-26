-- +goose Up
-- +goose StatementBegin
CREATE TYPE role AS ENUM ('admin', 'member');

ALTER TABLE users
ADD COLUMN IF NOT EXISTS role role NOT NULL DEFAULT 'member'
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN IF EXISTS role;

DROP TYPE role;

-- +goose StatementEnd
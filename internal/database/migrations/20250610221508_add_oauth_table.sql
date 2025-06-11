-- +goose Up
CREATE TABLE
  oauth_users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    provider VARCHAR(50),
    provider_id VARCHAR(255),
    name VARCHAR(255),
    avatar_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT NOW ()
  );

-- +goose Down
DROP TABLE IF EXISTS oauth_users
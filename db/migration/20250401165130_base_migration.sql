-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email text UNIQUE,
    password text
);

-- +goose Down
DROP TABLE IF EXISTS users;


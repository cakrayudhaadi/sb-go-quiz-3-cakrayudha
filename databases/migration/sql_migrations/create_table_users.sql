-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    modified_by VARCHAR NOT NULL
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS users;

-- +migrate StatementEnd
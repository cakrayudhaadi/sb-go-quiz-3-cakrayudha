-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS categories
(
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    modified_by VARCHAR NOT NULL
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS categories;

-- +migrate StatementEnd
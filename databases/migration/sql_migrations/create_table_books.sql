-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS books
(
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    image_url VARCHAR NOT NULL,
    release_year INTEGER NOT NULL,
    price INTEGER NOT NULL,
    total_page INTEGER NOT NULL,
    thickness VARCHAR NOT NULL,
    category_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    modified_by VARCHAR NOT NULL
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS books;

-- +migrate StatementEnd
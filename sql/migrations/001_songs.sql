-- +goose Up 

CREATE TABLE IF NOT EXISTS songs(
    id UUID PRIMARY KEY,
    band VARCHAR(70) NOT NULL,
    song VARCHAR(255) NOT NULL,
    release_date TIMESTAMP NOT NULL,
    lirics TEXT NOT NULL,
    link VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down

DROP TABLE IF EXISTS songs;
-- +goose Up 

CREATE TABLE IF NOT EXISTS songs(
    id UUID PRIMARY KEY,
    band_id UUID NOT NULL REFERENCES bands(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    release_date TIMESTAMP NOT NULL,
    lirics TEXT NOT NULL,
    link VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
CREATE INDEX IF NOT EXISTS name_idx ON songs(name);

-- +goose Down

DROP TABLE IF EXISTS songs;

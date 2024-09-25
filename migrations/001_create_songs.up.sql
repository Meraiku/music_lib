CREATE TABLE IF NOT EXISTS songs(
    id UUID PRIMARY KEY,
    band TEXT NOT NULL,
    song TEXT NOT NULL,
    release_date TIMESTAMP NOT NULL,
    lirics TEXT NOT NULL,
    link TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
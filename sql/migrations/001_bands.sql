-- +goose Up
CREATE TABLE IF NOT EXISTS bands (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE
);
CREATE INDEX IF NOT EXISTS bands_name_idx ON bands(name);

-- +goose Down
DROP TABLE IF EXISTS bands;

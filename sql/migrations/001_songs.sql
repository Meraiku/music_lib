-- goose +Up 

CREATE TABLE IF NOT EXISTS songs(
    id UUID PRIMARY KEY,
    band VARVHAR(70) NOT NULL,
    song VARVHAR(255) NOT NULL,
    release_date TIMESTAMP NOT NULL,
    lirics TEXT NOT NULL,
    link VARVHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- goose +Down

DROP TABLE IF EXISTS songs;
package postgre

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/meraiku/music_lib/internal/model"
	"github.com/meraiku/music_lib/internal/repo"
	"github.com/uptrace/bun"
)

func (db *postgre) GetSongs(ctx context.Context, params *model.Parameters) ([]model.Song, error) {

	limit := 20
	offset := limit * (params.Page - 1)
	songs := make([]repo.Song, 0, limit)

	if err := db.db.NewRaw("SELECT * FROM songs ORDER BY ? LIMIT ? OFFSET ?", bun.Ident(params.Filter), limit, offset).Scan(ctx, &songs); err != nil {
		return nil, err
	}

	out := make([]model.Song, len(songs))
	for i := range songs {
		out[i] = *repo.ToSongFromRepo(&songs[i])
	}

	return out, nil
}

func (db *postgre) AddSong(ctx context.Context, song *model.Song) (*model.Song, error) {

	query := `INSERT INTO songs 
	(id, band, song, release_date, lirics, link, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	s := repo.FromSongToRepo(song)

	s.ID = uuid.NewString()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	_, err := db.db.NewRaw(query,
		s.ID,
		s.Band,
		s.Song,
		s.ReleaseDate,
		s.Lirics,
		s.Link,
		s.CreatedAt,
		s.UpdatedAt,
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return repo.ToSongFromRepo(s), nil
}

func (db *postgre) DeleteSong(ctx context.Context, song *model.Song) error {

	s := repo.FromSongToRepo(song)

	_, err := db.db.NewRaw("DELETE FROM songs WHERE id = ?", s.ID).Exec(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}

	return nil
}

func (db *postgre) UpdateSong(ctx context.Context, song *model.Song) error {

	s := repo.FromSongToRepo(song)

	s.UpdatedAt = time.Now()

	if _, err := db.db.NewRaw("UPDATE songs SET band = ?, song = ?, updated_at = ? WHERE id = ?", s.Band, s.Song, s.UpdatedAt, s.ID).Exec(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repo.ErrSongIsNotExist
		}

		return err
	}
	return nil
}

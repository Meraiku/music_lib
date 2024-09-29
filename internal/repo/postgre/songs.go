package postgre

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mdobak/go-xerrors"
	"github.com/meraiku/music_lib/internal/model"
	"github.com/meraiku/music_lib/internal/repo"
	"github.com/uptrace/bun"
)

func (db *postgre) GetSongs(ctx context.Context, params *model.Parameters) ([]model.Song, error) {

	if params.Filter == "group" {
		params.Filter = "band"
	}

	limit := 20
	offset := limit * (params.Page - 1)
	songs := make([]repo.Song, 0, limit)

	if err := db.db.NewRaw("SELECT * FROM songs ORDER BY ? LIMIT ? OFFSET ?", bun.Ident(params.Filter), limit, offset).Scan(ctx, &songs); err != nil {
		return nil, xerrors.WithStackTrace(err, 0)
	}

	out := make([]model.Song, len(songs))
	for i := range songs {
		out[i] = *repo.ToSongFromRepo(&songs[i])
	}

	return out, nil
}

func (db *postgre) GetTextByID(ctx context.Context, id string) (string, error) {

	query := `SELECT (lirics) FROM songs WHERE id = ?`

	var text string

	if err := db.db.NewRaw(query, id).Scan(ctx, &text); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", repo.ErrSongIsNotExist
		}
		return "", xerrors.WithStackTrace(err, 0)
	}

	return text, nil
}

func (db *postgre) AddSong(ctx context.Context, song *model.Song) (*model.Song, error) {

	query := `INSERT INTO songs 
	(id, band, song, release_date, lirics, link, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?) RETURNING *`

	s := repo.FromSongToRepo(song)

	s.ID = uuid.NewString()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	err := db.db.NewRaw(query,
		s.ID,
		s.Band,
		s.Song,
		s.ReleaseDate,
		s.Lirics,
		s.Link,
		s.CreatedAt,
		s.UpdatedAt,
	).Scan(ctx, s)
	if err != nil {
		return nil, xerrors.WithStackTrace(err, 0)
	}

	return repo.ToSongFromRepo(s), nil
}

func (db *postgre) DeleteSong(ctx context.Context, id string) error {

	_, err := db.db.NewRaw("DELETE FROM songs WHERE id = ?", id).Exec(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return xerrors.WithStackTrace(err, 0)
	}

	return nil
}

func (db *postgre) UpdateSong(ctx context.Context, song *model.Update) (*model.Song, error) {

	var s repo.Song

	queryStart := `UPDATE songs SET `
	queyEnd := ` WHERE id = ? RETURNING *`

	u := repo.FromUpdateToRepo(song)

	upd := u.SQLUpdates()

	query := fmt.Sprintf("%s%s%s", queryStart, upd.Assignments(), queyEnd)

	args := upd.Values()
	args = append(args, u.ID)

	fmt.Println(query)
	fmt.Println(args...)

	if err := db.db.NewRaw(query, args...).Scan(ctx, &s); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repo.ErrSongIsNotExist
		}

		return nil, xerrors.WithStackTrace(err, 0)
	}
	return repo.ToSongFromRepo(&s), nil
}

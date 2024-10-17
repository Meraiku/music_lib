package postgre

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
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

	query := fmt.Sprintf(`SELECT songs.id, bands.name AS band, songs.name, songs.release_date, songs.lirics, songs.link
		FROM songs
		JOIN bands ON songs.band_id = bands.id 
		ORDER BY ? %s LIMIT ? OFFSET ? `, params.Order)

	db.log.DebugContext(ctx,
		"Getting songs list",
		slog.String("query", query),
	)

	limit := 20
	offset := limit * (params.Page - 1)
	songs := make([]repo.Song, 0, limit)

	if err := db.db.NewRaw(query, bun.Ident(params.Filter), limit, offset).Scan(ctx, &songs); err != nil {
		return nil, xerrors.WithStackTrace(err, 0)
	}

	out := make([]model.Song, len(songs))
	for i := range songs {
		out[i] = *repo.ToSongFromRepo(&songs[i])
	}

	db.log.DebugContext(ctx,
		"Got songs list",
		slog.Int("song_number", len(songs)),
	)

	return out, nil
}

func (db *postgre) GetTextByID(ctx context.Context, id string) (string, error) {

	query := `SELECT (lirics) FROM songs WHERE id = ?`

	db.log.DebugContext(ctx,
		"Getting song lirics",
		slog.String("id", id),
		slog.String("query", query),
	)

	var text string

	if err := db.db.NewRaw(query, id).Scan(ctx, &text); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", repo.ErrSongIsNotExist
		}
		return "", xerrors.WithStackTrace(err, 0)
	}

	db.log.DebugContext(ctx,
		"Got song lirics",
		slog.String("id", id),
		slog.String("lirics", text),
	)

	return text, nil
}

func (db *postgre) AddSong(ctx context.Context, song *model.Song) (*model.Song, error) {

	tx, err := db.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, xerrors.WithStackTrace(err, 0)
	}

	getBandIDQuery := `SELECT (id) FROM bands WHERE name = ?`

	s := repo.FromSongToRepo(song)
	bandID := ""

	db.log.DebugContext(ctx,
		"Getting band id",
		slog.String("name", s.Band),
		slog.String("query", getBandIDQuery),
	)

	if err := tx.NewRaw(getBandIDQuery,
		s.Band).
		Scan(ctx, &bandID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			createQuery := `INSERT INTO bands (id, name) VALUES (?, ?) RETURNING id`

			db.log.DebugContext(ctx,
				"Band not Found. Creating new one",
				slog.String("name", s.Band))

			if err := tx.NewRaw(createQuery,
				uuid.NewString(),
				s.Band).
				Scan(ctx, &bandID); err != nil {
				tx.Rollback()
				return nil, xerrors.WithStackTrace(err, 0)
			}

		} else {
			tx.Rollback()
			return nil, xerrors.WithStackTrace(err, 0)
		}
	}

	query := `INSERT INTO songs 
	(id, band_id, name, release_date, lirics, link, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?`

	s.ID = uuid.NewString()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	db.log.DebugContext(ctx,
		"Adding new song",
		slog.String("id", song.ID),
		slog.String("query", query),
	)

	err = tx.NewRaw(query,
		s.ID,
		bandID,
		s.Name,
		s.ReleaseDate,
		s.Lirics,
		s.Link,
		s.CreatedAt,
		s.UpdatedAt,
	).Scan(ctx, &s.ID)
	if err != nil {
		tx.Rollback()
		return nil, xerrors.WithStackTrace(err, 0)
	}

	tx.Commit()

	db.log.DebugContext(ctx,
		"Song added",
		slog.String("id", song.ID),
	)

	return repo.ToSongFromRepo(s), nil
}

func (db *postgre) DeleteSong(ctx context.Context, id string) error {

	query := "DELETE FROM songs WHERE id = ?"

	db.log.DebugContext(ctx,
		"Deleting Song",
		slog.String("song_id", id),
		slog.String("query", query),
	)

	_, err := db.db.NewRaw(query, id).Exec(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return xerrors.WithStackTrace(err, 0)
	}

	db.log.DebugContext(ctx,
		"Song deleted",
		slog.String("song_id", id),
	)

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

	db.log.DebugContext(ctx,
		"Updating song information",
		slog.String("id", song.ID),
		slog.String("query", query),
	)

	if err := db.db.NewRaw(query, args...).Scan(ctx, &s); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repo.ErrSongIsNotExist
		}

		return nil, xerrors.WithStackTrace(err, 0)
	}

	db.log.DebugContext(ctx,
		"Song updated",
		slog.String("song_id", song.ID),
	)

	return repo.ToSongFromRepo(&s), nil
}

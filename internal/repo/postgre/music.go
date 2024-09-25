package postgre

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/meraiku/music_lib/internal/model"
	repoModel "github.com/meraiku/music_lib/internal/repo/postgre/model"
	"github.com/uptrace/bun"

	"github.com/meraiku/music_lib/internal/repo/postgre/converter"
)

func (db *postgre) GetSongs(ctx context.Context, params *model.Parameters) ([]model.Song, error) {

	limit := 20
	offset := limit * (params.Page - 1)
	songs := make([]repoModel.Song, 0, limit)

	//if err := db.db.NewSelect().Model(&songs).Order(params.Filter).Limit(limit).Offset(offset).Scan(ctx); err != nil {
	//	return nil, err
	//}

	if err := db.db.NewRaw("SELECT * FROM songs ORDER BY ? LIMIT ? OFFSET ?", bun.Ident(params.Filter), limit, offset).Scan(ctx, &songs); err != nil {
		return nil, err
	}

	out := make([]model.Song, len(songs))
	for i := range songs {
		out[i] = *converter.ToSongFromRepo(&songs[i])
	}

	return out, nil
}

func (db *postgre) AddSong(ctx context.Context, song *model.Song) error {

	s := converter.FromSongToRepo(song)

	s.ID = uuid.NewString()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	_, err := db.db.NewInsert().Model(s).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (db *postgre) DeleteSong(ctx context.Context, song *model.Song) error {
	return nil
}

func (db *postgre) UpdateSong(ctx context.Context, song *model.Song) error {
	return nil
}

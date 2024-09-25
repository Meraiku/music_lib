package postgre

import (
	"context"

	"github.com/meraiku/music_lib/internal/model"
)

func (db *postgre) GetSongs(ctx context.Context, params *model.Parameters) ([]model.Song, error) {
	return nil, nil
}

func (db *postgre) AddSong(ctx context.Context, song *model.Song) error {
	return nil
}

func (db *postgre) DeleteSong(ctx context.Context, song *model.Song) error {
	return nil
}

func (db *postgre) UpdateSong(ctx context.Context, song *model.Song) error {
	return nil
}

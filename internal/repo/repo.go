package repo

import (
	"context"

	"github.com/meraiku/music_lib/internal/model"
)

type MusicRepository interface {
	GetSongs(ctx context.Context, params *model.Parameters) ([]model.Song, error)
	AddSong(ctx context.Context, song *model.Song) error
	DeleteSong(ctx context.Context, song *model.Song) error
	UpdateSong(ctx context.Context, song *model.Song) error
}

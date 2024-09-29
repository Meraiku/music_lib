package repo

import (
	"context"

	"github.com/meraiku/music_lib/internal/model"
)

type MusicRepository interface {
	GetSongs(ctx context.Context, params *model.Parameters) ([]model.Song, error)
	GetTextByID(ctx context.Context, id string) (string, error)
	AddSong(ctx context.Context, song *model.Song) (*model.Song, error)
	DeleteSong(ctx context.Context, id string) error
	UpdateSong(ctx context.Context, song *model.Update) (*model.Song, error)
}

package service

import (
	"context"

	"github.com/meraiku/music_lib/internal/model"
)

type MusicService interface {
	GetSongs(ctx context.Context, params *model.Parameters) ([]model.Song, error)
	GetText(ctx context.Context, id string, couplet int) ([]model.Text, error)
	PostSong(ctx context.Context, song *model.Song) (*model.Song, error)
	UpdateSong(ctx context.Context, song *model.Update) (*model.Song, error)
	DeleteSong(ctx context.Context, id string) error
}

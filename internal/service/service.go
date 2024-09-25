package service

import (
	"context"

	"github.com/meraiku/music_lib/internal/model"
)

type MusicService interface {
	GetSongs(ctx context.Context, params *model.Parameters) ([]model.Song, error)
	PostSong(ctx context.Context, song *model.Song) error
	DeleteSong(ctx context.Context, song *model.Song) error
	UpdateSong(ctx context.Context, song *model.Song) error
	GetText(ctx context.Context, song *model.Song, params *model.Parameters) ([]model.Text, error)
}

package music

import (
	"context"

	"github.com/meraiku/music_lib/internal/model"
)

func (s *service) GetSongs(ctx context.Context, params *model.Parameters) ([]model.Song, error) {
	return nil, nil
}

func (s *service) PostSong(ctx context.Context, song *model.Song) error {
	return nil
}

func (s *service) DeleteSong(ctx context.Context, song *model.Song) error {
	return nil
}

func (s *service) UpdateSong(ctx context.Context, song *model.Song) error {
	return nil
}

func (s *service) GetText(ctx context.Context, song *model.Song, params *model.Parameters) ([]model.Text, error) {
	return nil, nil
}

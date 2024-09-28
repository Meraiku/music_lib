package music

import (
	"context"

	"github.com/meraiku/music_lib/internal/model"
)

func (s *service) GetSongs(ctx context.Context, params *model.Parameters) ([]model.Song, error) {

	s.log.DebugContext(ctx, "Get Song OK")

	if params.Filter == "" {
		params.Filter = "song"
	}
	if params.Page == 0 {
		params.Page = 1
	}

	songs, err := s.repo.GetSongs(ctx, params)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (s *service) PostSong(ctx context.Context, song *model.Song) (*model.Song, error) {

	s.log.DebugContext(ctx, "Post Song OK")

	song, err := s.repo.AddSong(ctx, song)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (s *service) DeleteSong(ctx context.Context, song *model.Song) error {

	s.log.DebugContext(ctx, "Delete Song OK")

	return s.repo.DeleteSong(ctx, song)
}

func (s *service) UpdateSong(ctx context.Context, song *model.Song) (*model.Song, error) {

	s.log.DebugContext(ctx, "Update Song OK")

	return s.repo.UpdateSong(ctx, song)
}

func (s *service) GetText(ctx context.Context, song *model.Song, params *model.Parameters) ([]model.Text, error) {

	s.log.DebugContext(ctx, "Get Text OK")

	return nil, nil
}

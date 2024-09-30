package music

import (
	"context"
	"errors"
	"log/slog"

	"github.com/mdobak/go-xerrors"
	"github.com/meraiku/music_lib/internal/lib/fetcher"
	"github.com/meraiku/music_lib/internal/model"
)

func (s *service) GetSongs(ctx context.Context, params *model.Parameters) ([]model.Song, error) {

	s.log.DebugContext(ctx, "Get Song OK")

	songs, err := s.repo.GetSongs(ctx, params)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (s *service) GetText(ctx context.Context, id string, verse int) ([]model.Text, error) {

	s.log.DebugContext(ctx, "Get Text OK")

	s.log.DebugContext(ctx, "Getting text from repo", slog.String("song_id", id))

	text, err := s.repo.GetTextByID(ctx, id)
	if err != nil {
		return nil, err
	}

	s.log.DebugContext(ctx, "Got text from repo", slog.String("text", text))

	out := paginateText(text)

	if verse > len(out) || verse < 0 {
		verse = 0
	}

	switch verse {
	case 0:
		return out, nil
	default:
		return out[verse-1 : verse : verse], nil
	}
}

func (s *service) PostSong(ctx context.Context, song *model.Song) (*model.Song, error) {

	s.log.DebugContext(ctx, "Post Song OK")

	s.log.DebugContext(ctx, "Fetching Data from Info service")

	f := fetcher.NewInfo()

	err := f.BeginWithContext(ctx, song)
	if err != nil {
		if errors.Is(err, fetcher.ErrNoData) {
			return nil, err
		}
		return nil, xerrors.WithStackTrace(err, -1)
	}

	s.log.DebugContext(ctx, "Add data to repository")

	song, err = s.repo.AddSong(ctx, song)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (s *service) DeleteSong(ctx context.Context, id string) error {

	s.log.DebugContext(ctx, "Delete Song OK")

	return s.repo.DeleteSong(ctx, id)
}

func (s *service) UpdateSong(ctx context.Context, song *model.Update) (*model.Song, error) {

	s.log.DebugContext(ctx, "Update Song OK")

	out, err := s.repo.UpdateSong(ctx, song)
	if err != nil {
		return nil, err
	}

	return out, nil
}

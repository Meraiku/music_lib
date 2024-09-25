package converter

import (
	"github.com/meraiku/music_lib/internal/model"
	repoModel "github.com/meraiku/music_lib/internal/repo/postgre/model"
)

func FromSongToRepo(s *model.Song) *repoModel.Song {
	return &repoModel.Song{
		ID:          s.ID,
		Band:        s.Group,
		Song:        s.Song,
		ReleaseDate: s.ReleaseDate,
		Lirics:      s.Text,
		Link:        s.Link,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func ToSongFromRepo(s *repoModel.Song) *model.Song {
	return &model.Song{
		ID:          s.ID,
		Group:       s.Band,
		Song:        s.Song,
		ReleaseDate: s.ReleaseDate,
		Text:        s.Lirics,
		Link:        s.Link,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

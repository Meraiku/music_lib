package repo

import (
	"github.com/meraiku/music_lib/internal/model"
)

func FromSongToRepo(s *model.Song) *Song {
	return &Song{
		ID:          s.ID,
		Band:        s.Group,
		Song:        s.Song,
		ReleaseDate: s.ReleaseDate,
		Lirics:      s.Text,
		Link:        s.Link,
	}
}

func ToSongFromRepo(s *Song) *model.Song {
	return &model.Song{
		ID:          s.ID,
		Group:       s.Band,
		Song:        s.Song,
		ReleaseDate: s.ReleaseDate,
		Text:        s.Lirics,
		Link:        s.Link,
	}
}

package converter

import (
	"time"

	"github.com/meraiku/music_lib/internal/api"
	"github.com/meraiku/music_lib/internal/model"
)

func FromAddSongRequestToModel(req *api.AddSongRequest) *model.Song {
	return &model.Song{
		Group: req.Group,
		Song:  req.Song,
	}
}

func FromPatchRequestToModel(req *api.PatchRequest, id string) *model.Update {
	var date *time.Time

	if req.ReleaseDate != nil {
		rd, err := time.Parse("02.01.2006", *req.ReleaseDate)
		if err == nil {
			date = &rd
		}
	}

	return &model.Update{
		ID:          id,
		Group:       req.Group,
		Song:        req.Song,
		ReleaseDate: date,
		Text:        req.Text,
		Link:        req.Link,
	}
}

func FromSongToApi(s *model.Song) *api.Song {
	return &api.Song{
		ID:          s.ID,
		Group:       s.Group,
		Song:        s.Song,
		ReleaseDate: s.ReleaseDate.Format("02.01.2006"),
		Text:        s.Text,
		Link:        s.Link,
	}
}

func FromSongsToApiSongs(s []model.Song) []api.Song {
	out := make([]api.Song, len(s))

	for i, song := range s {
		apiSong := FromSongToApi(&song)
		out[i] = *apiSong
	}

	return out
}

func FromTextToApi(t []model.Text) []api.Text {
	out := make([]api.Text, len(t))

	for i, text := range t {
		out[i] = api.Text{
			VerseNumber: text.VerseNumber,
			Verse:       text.Verse,
		}
	}

	return out
}

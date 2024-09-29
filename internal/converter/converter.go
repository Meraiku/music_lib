package converter

import (
	"github.com/meraiku/music_lib/internal/api/rest/request"
	"github.com/meraiku/music_lib/internal/model"
)

func FromAddSongRequestToModel(req *request.AddSongRequest) *model.Song {
	return &model.Song{
		Group: req.Group,
		Song:  req.Song,
	}
}

func FromPatchRequestToModel(req *request.PatchRequest, id string) *model.Update {
	return &model.Update{
		ID:          id,
		Group:       req.Group,
		Song:        req.Song,
		ReleaseDate: req.ReleaseDate,
		Text:        req.Text,
		Link:        req.Link,
	}
}

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

func FromModifySongRequestToModel(req *request.ModifySongRequest) *model.Song {
	return &model.Song{
		ID:    req.ID,
		Group: req.Group,
		Song:  req.Song,
	}
}

package repo

import "errors"

var (
	ErrSongIsNotExist = errors.New("song with that id is not exists")
)

package repo

import "errors"

var (
	ErrSongIsNotExist = errors.New("song with provided id is not exists")
)

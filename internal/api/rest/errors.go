package rest

import "errors"

var (
	ErrNoBody = errors.New("request body is empty")
)

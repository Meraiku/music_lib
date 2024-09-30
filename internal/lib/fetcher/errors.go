package fetcher

import "errors"

var (
	ErrNoData                 = errors.New("song data is not found")
	ErrBadServiceEndpoint     = errors.New("incorrect service endpoint")
	ErrInvalidServiceURL      = errors.New("service not found")
	ErrInfoServiceUnavailable = errors.New("info service unavailable")
)

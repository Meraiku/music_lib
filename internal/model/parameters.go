package model

import "errors"

var (
	ErrInvalidPage   = errors.New("invalid page number")
	ErrInvalidFilter = errors.New("invalid filter parameter")

	maxPage = 1000
)

type Parameters struct {
	Filter string `json:"filter"`
	Page   int    `json:"page"`
}

func (p *Parameters) Validate() error {
	if p.Page < 0 || p.Page > maxPage {
		return ErrInvalidPage
	}

	switch p.Filter {
	case "group", "song":
		return nil
	default:
		return ErrInvalidFilter
	}
}

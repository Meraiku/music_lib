package model

import (
	"errors"
	"strings"
)

var (
	ErrInvalidPage   = errors.New("invalid page number")
	ErrInvalidFilter = errors.New("invalid filter parameter")

	maxPage = 1000
)

type Parameters struct {
	Filter string `json:"filter"`
	Page   int    `json:"page"`
	Order  string `json:"order"`
}

func (p *Parameters) Validate() error {

	p.Order = strings.ToUpper(p.Order)

	switch p.Order {
	case "ASC", "DESC":
	default:
		p.Order = "ASC"
	}

	if p.Page < 0 || p.Page > maxPage {
		return ErrInvalidPage
	}

	switch p.Filter {
	case "group", "song", "created_at", "updated_at", "releaseDate":
		return nil
	default:
		return ErrInvalidFilter
	}
}

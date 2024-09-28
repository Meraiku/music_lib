package model

import (
	"time"
)

type Song struct {
	ID          string    `json:"id"`
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	ReleaseDate time.Time `json:"releaseDate"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

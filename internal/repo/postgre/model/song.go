package model

import (
	"time"
)

type Song struct {
	ID          string    `json:"id"`
	Band        string    `json:"band"`
	Song        string    `json:"song"`
	ReleaseDate time.Time `json:"releaseDate"`
	Lirics      string    `json:"lirics"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

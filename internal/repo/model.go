package repo

import "time"

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

type Parameters struct {
	Filter *string `json:"filter"`
	Page   *int    `json:"page"`
}

type Text struct {
	Text string `json:"text"`
}

package repo

import "time"

type Song struct {
	ID          string    `json:"id"`
	Band        string    `json:"bands"`
	Name        string    `json:"name"`
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

type Update struct {
	ID          string     `json:"id"`
	Song        *string    `json:"song"`
	ReleaseDate *time.Time `json:"releaseDate"`
	Lirics      *string    `json:"lirics"`
	Link        *string    `json:"link"`
}

func (u *Update) SQLUpdates() SQLUpdate {
	var s SQLUpdate
	if u.Song != nil {
		s.add("song", *u.Song)
	}
	if u.ReleaseDate != nil {
		s.add("release_date", *u.ReleaseDate)
	}
	if u.Lirics != nil {
		s.add("lirics", *u.Lirics)
	}
	if u.Link != nil {
		s.add("link", *u.Link)
	}

	s.add("updated_at", time.Now().UTC())
	return s
}

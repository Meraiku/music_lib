package main

import (
	"errors"
)

type InfoRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type InfoResponse struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func (req InfoRequest) validate() error {
	lib := []InfoRequest{
		InfoRequest{
			Group: "Muse",
			Song:  "Supermassive Black Hole",
		},
	}

	for i := range lib {
		if req == lib[i] {
			return nil
		}
	}

	return errors.New("no records")
}

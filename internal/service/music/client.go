package music

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/meraiku/music_lib/internal/model"
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

func FetchSongInfo(ctx context.Context, song *model.Song) error {
	url := os.Getenv("SERVICE_URL") + "/info"

	c := http.Client{
		Timeout: 5 * time.Second,
	}

	b := &InfoRequest{
		Group: song.Group,
		Song:  song.Song,
	}

	data, err := json.Marshal(b)
	if err != nil {
		return err
	}

	r := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, r)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return errors.New("bad status code")
	}

	var out InfoResponse

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return err
	}

	song.Text = out.Text
	song.Link = out.Link

	return nil
}

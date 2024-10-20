package info

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/meraiku/music_lib/internal/lib/fetcher"
	"github.com/meraiku/music_lib/internal/model"
)

type InfoFetcher struct {
	timeout time.Duration
	url     string
}

func New(url string) *InfoFetcher {
	return &InfoFetcher{
		timeout: 5 * time.Second,
		url:     url,
	}
}

func NewInfo() *InfoFetcher {
	return &InfoFetcher{
		timeout: 5 * time.Second,
		url:     os.Getenv("SERVICE_URL") + "/info",
	}
}

func (f *InfoFetcher) FetchCtx(ctx context.Context, song *model.Song) error {

	b := &InfoRequest{
		Group: song.Group,
		Song:  song.Song,
	}

	resp, err := getRequest(ctx, b, f)
	if err != nil {
		return fetcher.ErrInfoServiceUnavailable
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusBadRequest:
		return fetcher.ErrNoData
	case http.StatusNotFound:
		return fetcher.ErrBadServiceEndpoint
	case http.StatusInternalServerError:

		// TODO Retry Logic

		return fetcher.ErrNoData
	case http.StatusOK:

		var out InfoResponse

		if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
			return err
		}

		releaseDate, _ := time.Parse("02.01.2006", out.ReleaseDate)

		song.ReleaseDate = releaseDate
		song.Text = out.Text
		song.Link = out.Link

		return nil
	default:
		return fetcher.ErrInvalidServiceURL
	}
}

func getRequest(ctx context.Context, payload any, f *InfoFetcher) (*http.Response, error) {

	c := http.Client{
		Timeout: f.timeout,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, f.url, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

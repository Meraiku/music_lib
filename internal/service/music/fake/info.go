package fake

import (
	"context"
	"time"

	"github.com/meraiku/music_lib/internal/model"
)

var (
	ReleaseDate = time.Date(2006, 7, 16, 0, 0, 0, 0, time.UTC)
)

type InfoFake struct {
}

func (f *InfoFake) FetchCtx(ctx context.Context, song *model.Song) error {

	// Simulate service response latency
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-ticker.C:

		song.ReleaseDate = ReleaseDate
		song.Text = "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"
		song.Link = "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
	}

	return nil
}

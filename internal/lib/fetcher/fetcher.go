package fetcher

import (
	"context"

	"github.com/meraiku/music_lib/internal/model"
)

type Fetcher interface {
	FetchCtx(ctx context.Context, song *model.Song) error
}

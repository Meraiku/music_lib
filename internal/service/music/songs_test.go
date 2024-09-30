package music

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/meraiku/music_lib/internal/model"
	mock_repo "github.com/meraiku/music_lib/internal/repo/mocks"
	"github.com/meraiku/music_lib/pkg/logging"
	"github.com/stretchr/testify/require"
)

func TestGetSongs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockMusicRepository(ctrl)

	ctx := context.Background()
	params := &model.Parameters{}

	expResp := []model.Song{
		{
			Group:       "Muse",
			Song:        "Supermassive Black Hole",
			ReleaseDate: time.Now(),
			Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
			Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
		},
	}

	s := NewService(repo, logging.Init("testing"))

	repo.EXPECT().GetSongs(ctx, params).Return(expResp, nil).Times(1)
	songs, err := s.GetSongs(ctx, params)

	require.NoError(t, err)
	require.Equal(t, expResp, songs)
}

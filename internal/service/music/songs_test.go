package music

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/meraiku/music_lib/internal/model"
	mock_repo "github.com/meraiku/music_lib/internal/repo/mocks"
	"github.com/meraiku/music_lib/internal/service/music/fake"
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

	s := NewService(repo, logging.Init("testing"), nil)

	repo.EXPECT().GetSongs(ctx, params).Return(expResp, nil).Times(1)
	songs, err := s.GetSongs(ctx, params)

	require.NoError(t, err)
	require.Equal(t, expResp, songs)
}

func TestGetSongsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockMusicRepository(ctrl)

	ctx := context.Background()
	params := &model.Parameters{}

	s := NewService(repo, logging.Init("testing"), nil)

	repo.EXPECT().GetSongs(ctx, params).Return(nil, errors.New("db in unavailable")).Times(1)
	songs, err := s.GetSongs(ctx, params)

	require.Error(t, err)
	require.Nil(t, songs)
}

func TestGetText(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockMusicRepository(ctrl)
	s := NewService(repo, logging.Init("testing"), nil)

	ctx := context.Background()
	id := uuid.NewString()
	text := "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"

	tt := []struct {
		name         string
		verse        int
		textExpected []model.Text
	}{
		{
			name:  "Get text without verse number",
			verse: 0,
			textExpected: []model.Text{
				{
					VerseNumber: 1,
					Verse:       "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?",
				},
				{
					VerseNumber: 2,
					Verse:       "Ooh\nYou set my soul alight\nOoh\nYou set my soul alight",
				},
			},
		},
		{
			name:  "Get text with verse number provided",
			verse: 2,
			textExpected: []model.Text{
				{
					VerseNumber: 2,
					Verse:       "Ooh\nYou set my soul alight\nOoh\nYou set my soul alight",
				},
			},
		},
		{
			name:  "Get text with random verse number",
			verse: -241,
			textExpected: []model.Text{
				{
					VerseNumber: 1,
					Verse:       "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?",
				},
				{
					VerseNumber: 2,
					Verse:       "Ooh\nYou set my soul alight\nOoh\nYou set my soul alight",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			repo.EXPECT().GetTextByID(ctx, id).Return(text, nil).Times(1)
			texts, err := s.GetText(ctx, id, tc.verse)

			require.NoError(t, err)
			require.Equal(t, tc.textExpected, texts)
		})
	}
}

func TestGetTextError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockMusicRepository(ctrl)

	ctx := context.Background()
	id := uuid.NewString()
	verse := 0

	s := NewService(repo, logging.Init("testing"), nil)

	repo.EXPECT().GetTextByID(ctx, id).Return("", errors.New("db in unavailable")).Times(1)
	texts, err := s.GetText(ctx, id, verse)

	require.Error(t, err)
	require.Nil(t, texts)
}

func TestDeleteSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockMusicRepository(ctrl)

	ctx := context.Background()
	id := uuid.NewString()

	s := NewService(repo, logging.Init("testing"), nil)

	repo.EXPECT().DeleteSong(ctx, id).Return(nil).Times(1)
	err := s.DeleteSong(ctx, id)

	require.NoError(t, err)
}

func TestDeleteSongError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockMusicRepository(ctrl)

	ctx := context.Background()
	id := uuid.NewString()

	s := NewService(repo, logging.Init("testing"), nil)

	repo.EXPECT().DeleteSong(ctx, id).Return(errors.New("db in unavailable")).Times(1)
	err := s.DeleteSong(ctx, id)

	require.Error(t, err)
}

func TestUpdateSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockMusicRepository(ctrl)

	ctx := context.Background()

	song := "Supermassive Black Holes"
	id := uuid.NewString()

	in := &model.Update{
		ID:   id,
		Song: &song,
	}

	expResp := &model.Song{
		ID:          id,
		Group:       "Muse",
		Song:        song,
		ReleaseDate: time.Now(),
		Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	}

	s := NewService(repo, logging.Init("testing"), nil)

	repo.EXPECT().UpdateSong(ctx, in).Return(expResp, nil).Times(1)
	songs, err := s.UpdateSong(ctx, in)

	require.NoError(t, err)
	require.Equal(t, expResp.ID, in.ID)
	require.Equal(t, expResp.Song, songs.Song)
}

func TestUpdateSongError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockMusicRepository(ctrl)

	ctx := context.Background()
	id := uuid.NewString()

	in := &model.Update{
		ID: id,
	}

	s := NewService(repo, logging.Init("testing"), nil)

	repo.EXPECT().UpdateSong(ctx, in).Return(nil, errors.New("db in unavailable")).Times(1)
	song, err := s.UpdateSong(ctx, in)

	require.Error(t, err)
	require.Nil(t, song)
}

func TestPostSong(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockMusicRepository(ctrl)

	ctx := context.Background()

	id := uuid.New()

	in := model.Song{
		Song:  "Supermassive Black Holes",
		Group: "Muse",
	}

	expResp := model.Song{
		ID:          id.String(),
		Song:        "Supermassive Black Holes",
		Group:       "Muse",
		Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
		ReleaseDate: fake.ReleaseDate,
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	}

	f := &fake.InfoFake{}

	s := NewService(repo, logging.Init("testing"), f)

	repo.EXPECT().AddSong(ctx, &in).Return(&expResp, nil).Times(1)
	songs, err := s.PostSong(ctx, &in)

	require.NoError(t, err)
	require.Equal(t, &expResp, songs)
}

func TestPostSongInfoServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repo.NewMockMusicRepository(ctrl)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	in := model.Song{
		Song:  "Supermassive Black Holes",
		Group: "Muse",
	}

	f := &fake.InfoFake{}

	s := NewService(repo, logging.Init("testing"), f)

	_, err := s.PostSong(ctx, &in)

	require.Error(t, err)
}

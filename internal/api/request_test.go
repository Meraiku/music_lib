package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSongVaildate(t *testing.T) {

	req := &AddSongRequest{
		Group: "band",
		Song:  "song",
	}

	errors := req.Validate()

	require.Nil(t, errors)
}

func TestAddSongVaildateError(t *testing.T) {

	tt := []struct {
		name string
		in   AddSongRequest
		keys []string
	}{
		{
			name: "Missing Group And Song names",
			in:   AddSongRequest{},
			keys: []string{"group", "song"},
		},
		{
			name: "Missing Song name",
			in: AddSongRequest{
				Group: "name",
			},
			keys: []string{"song"},
		},
		{
			name: "Missing Group name",
			in: AddSongRequest{
				Song: "name",
			},
			keys: []string{"group"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			errors := tc.in.Validate()

			for _, key := range tc.keys {
				_, ok := errors[key]

				require.True(t, ok)
			}
		})
	}
}

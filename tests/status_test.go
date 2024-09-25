package tests

import (
	"net/http"
	"testing"

	"github.com/meraiku/music_lib/internal/config"
	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {

	c := http.Client{}

	config.Load(".env")

	cfg := config.NewRESTConfig()

	addr := "http://" + cfg.Address()
	path := "/healthz"

	url := addr + path

	resp, err := c.Get(url)

	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode)
}

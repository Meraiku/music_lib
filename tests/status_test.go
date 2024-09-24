package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {

	c := http.Client{}

	resp, err := c.Get("http://localhost:8080/healthz")

	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode)
}

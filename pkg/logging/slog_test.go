package logging

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	log := Init("dev")

	require.NotNil(t, log)
}

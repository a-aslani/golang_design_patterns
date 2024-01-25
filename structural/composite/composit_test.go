package composite

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	err := Run()
	require.NoError(t, err)
}

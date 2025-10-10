package bpx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	key    = ""
	secret = ""
)

func TestAsserts_GetAssets(t *testing.T) {
	c := NewClient(key, secret)
	res, err := c.Asserts().GetAssets()
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestAsserts_GetCollaterals(t *testing.T) {
	c := NewClient(key, secret)
	res, err := c.Asserts().GetCollaterals()
	require.NoError(t, err)
	require.NotNil(t, res)
}

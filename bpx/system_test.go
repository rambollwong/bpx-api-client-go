package bpx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSystem_Status(t *testing.T) {
	c := NewClient(key, secret)
	resp, err := c.System().Status()
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestSystem_Ping(t *testing.T) {
	c := NewClient(key, secret)
	resp, err := c.System().Ping()
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestSystem_Time(t *testing.T) {
	c := NewClient(key, secret)
	resp, err := c.System().Time()
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestSystem_GetWallets(t *testing.T) {
	c := NewClient(key, secret)
	resp, err := c.System().GetWallets()
	require.NoError(t, err)
	require.NotNil(t, resp)
}

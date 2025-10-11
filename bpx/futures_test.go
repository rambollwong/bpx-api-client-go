package bpx

import (
	"testing"

	"github.com/rambollwong/bpx-api-client-go/types"

	"github.com/stretchr/testify/require"
)

func TestFutures_GetOpenPositions(t *testing.T) {
	c := NewClient(key, secret).Futures()
	resp, err := c.GetOpenPositions(types.GetOpenPositionsReq{})
	require.NoError(t, err)
	require.NotNil(t, resp)

	symbol := "DOGE_USDC_PERP"
	resp, err = c.GetOpenPositions(types.GetOpenPositionsReq{
		Symbol: &symbol,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	symbol = "ABC"
	_, err = c.GetOpenPositions(types.GetOpenPositionsReq{
		Symbol: &symbol,
	})
	require.Error(t, err)
}

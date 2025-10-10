package bpx

import (
	"testing"

	"bpx-api-client-go/types"

	"github.com/stretchr/testify/require"
)

func TestTrades_GetRecentTrades(t *testing.T) {
	c := NewClient(key, secret).Trades()
	trades, err := c.GetRecentTrades(types.GetRecentTradesReq{
		Symbol: "ETH_USDC_PERP",
		Limit:  100,
	})
	require.NoError(t, err)
	require.NotNil(t, trades)

	_, err = c.GetRecentTrades(types.GetRecentTradesReq{
		Symbol: "ETH_USDC_PERP",
		Limit:  99,
	})
	require.Error(t, err)

	_, err = c.GetRecentTrades(types.GetRecentTradesReq{
		Symbol: "ETH_USDC_PERP",
		Limit:  1001,
	})
	require.Error(t, err)
}

func TestTrades_GetHistoricalTrades(t *testing.T) {
	c := NewClient(key, secret).Trades()
	trades, err := c.GetHistoricalTrades(types.GetHistoricalTradesReq{
		Symbol: "ETH_USDC_PERP",
		Limit:  100,
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotNil(t, trades)

	trades, err = c.GetHistoricalTrades(types.GetHistoricalTradesReq{
		Symbol: "ETH_USDC_PERP",
		Limit:  100,
		Offset: 1,
	})
	require.NoError(t, err)
	require.NotNil(t, trades)

	_, err = c.GetHistoricalTrades(types.GetHistoricalTradesReq{
		Symbol: "ETH_USDC_PERP",
		Limit:  99,
		Offset: 0,
	})
	require.Error(t, err)

	_, err = c.GetHistoricalTrades(types.GetHistoricalTradesReq{
		Symbol: "ETH_USDC_PERP",
		Limit:  1001,
		Offset: 0,
	})
	require.Error(t, err)
}

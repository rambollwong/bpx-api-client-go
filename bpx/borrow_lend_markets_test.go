package bpx

import (
	"testing"

	"github.com/rambollwong/bpx-api-client-go/types"

	"github.com/stretchr/testify/require"
)

func TestBorrowLendMarkets_GetBorrowLendMarkets(t *testing.T) {
	c := NewClient(key, secret)
	res, err := c.BorrowLendMarkets().GetBorrowLendMarkets()
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestBorrowLendMarkets_GetBorrowLendMarketsHistory(t *testing.T) {
	c := NewClient(key, secret)
	_, err := c.BorrowLendMarkets().GetBorrowLendMarketsHistory(types.BorrowLendMarketHistoryReq{
		Interval: "",
		Symbol:   "",
	})
	require.Error(t, err)

	res, err := c.BorrowLendMarkets().GetBorrowLendMarketsHistory(types.BorrowLendMarketHistoryReq{
		Interval: types.BorrowLendMarketHistoryInterval1D,
		Symbol:   "",
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Greater(t, len(res), 1)

	res, err = c.BorrowLendMarkets().GetBorrowLendMarketsHistory(types.BorrowLendMarketHistoryReq{
		Interval: types.BorrowLendMarketHistoryInterval1D,
		Symbol:   "ETH",
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Greater(t, len(res), 1)
}

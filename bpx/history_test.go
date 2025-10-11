package bpx

import (
	"testing"

	"github.com/rambollwong/bpx-api-client-go/types"

	"github.com/stretchr/testify/require"
)

func TestHistory_GetBorrowHistory(t *testing.T) {
	c := NewClient(key, secret).History()
	res, rh, err := c.GetBorrowHistory(types.GetBorrowHistoryReq{})
	require.NoError(t, err)
	require.NotNil(t, rh)
	require.NotNil(t, res)
}

func TestHistory_GetInterestHistory(t *testing.T) {
	c := NewClient(key, secret).History()
	res, rh, err := c.GetInterestHistory(types.GetInterestHistoryReq{})
	require.NoError(t, err)
	require.NotNil(t, rh)
	require.NotNil(t, res)
}

func TestHistory_GetBorrowPositionHistory(t *testing.T) {
	c := NewClient(key, secret).History()
	res, rh, err := c.GetBorrowPositionHistory(types.GetBorrowPositionHistoryReq{})
	require.NoError(t, err)
	require.NotNil(t, rh)
	require.NotNil(t, res)
}

func TestHistory_GetDustConversionHistory(t *testing.T) {
	c := NewClient(key, secret).History()
	res, rh, err := c.GetDustConversionHistory(types.GetDustConversionHistoryReq{})
	require.NoError(t, err)
	require.NotNil(t, rh)
	require.NotNil(t, res)
}

func TestHistory_GetBorrowHistoryWithContext(t *testing.T) {
	c := NewClient(key, secret).History()
	res, rh, err := c.GetFillHistory(types.GetFillHistoryReq{})
	require.NoError(t, err)
	require.NotNil(t, rh)
	require.NotNil(t, res)
}

func TestHistory_GetFundingPayments(t *testing.T) {
	c := NewClient(key, secret).History()
	res, rh, err := c.GetFundingPayments(types.GetFundingPaymentsReq{})
	require.NoError(t, err)
	require.NotNil(t, rh)
	require.NotNil(t, res)
}

func TestHistory_GetOrderHistory(t *testing.T) {
	c := NewClient(key, secret).History()
	res, rh, err := c.GetOrderHistory(types.GetOrderHistoryReq{})
	require.NoError(t, err)
	require.NotNil(t, rh)
	require.NotNil(t, res)
}

func TestHistory_GetRfqHistory(t *testing.T) {
	c := NewClient(key, secret).History()
	res, rh, err := c.GetRfqHistory(types.GetRfqHistoryReq{})
	require.NoError(t, err)
	require.NotNil(t, rh)
	require.NotNil(t, res)
}

func TestHistory_GetQuoteHistory(t *testing.T) {
	c := NewClient(key, secret).History()
	res, rh, err := c.GetQuoteHistory(types.GetQuoteHistoryReq{})
	require.NoError(t, err)
	require.NotNil(t, rh)
	require.NotNil(t, res)
}

func TestHistory_GetSettlementHistory(t *testing.T) {
	c := NewClient(key, secret).History()
	res, rh, err := c.GetSettlementHistory(types.GetSettlementHistoryReq{})
	require.NoError(t, err)
	require.NotNil(t, rh)
	require.NotNil(t, res)
}

func TestHistory_GetStrategyHistory(t *testing.T) {
	c := NewClient(key, secret).History()
	res, rh, err := c.GetStrategyHistory(types.GetStrategyHistoryReq{})
	require.NoError(t, err)
	require.NotNil(t, rh)
	require.NotNil(t, res)
}

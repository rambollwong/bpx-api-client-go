package bpx

import (
	"testing"
	"time"

	"github.com/rambollwong/bpx-api-client-go/types"

	"github.com/stretchr/testify/require"
)

func TestCapital_GetBalances(t *testing.T) {
	c := NewClient(key, secret).Capital()
	resp, err := c.GetBalances()
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestCapital_GetCollateral(t *testing.T) {
	c := NewClient(key, secret).Capital()
	resp, err := c.GetCollateral(types.GetCollateralReq{})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestCapital_GetDeposits(t *testing.T) {
	c := NewClient(key, secret).Capital()
	resp, err := c.GetDeposits(types.GetDepositsReq{})
	require.NoError(t, err)
	require.NotNil(t, resp)

	var from, to = time.Now().Add(-time.Hour * 24 * 365).UnixMilli(), time.Now().UnixMilli()
	var limit, offset uint64 = 2, 0
	resp, err = c.GetDeposits(types.GetDepositsReq{
		From:   &from,
		To:     &to,
		Limit:  &limit,
		Offset: &offset,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestCapital_GetDepositAddress(t *testing.T) {
	c := NewClient(key, secret).Capital()
	resp, err := c.GetDepositAddress(types.GetDepositAddressReq{
		Blockchain: "Solana",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	_, err = c.GetDepositAddress(types.GetDepositAddressReq{
		Blockchain: "",
	})
	require.Error(t, err)
}

func TestCapital_GetWithdrawals(t *testing.T) {
	c := NewClient(key, secret).Capital()
	resp, err := c.GetWithdrawals(types.GetWithdrawalsReq{})
	require.NoError(t, err)
	require.NotNil(t, resp)

	var from, to = time.Now().Add(-time.Hour * 24 * 365).UnixMilli(), time.Now().UnixMilli()
	var limit, offset uint64 = 1, 1
	resp, err = c.GetWithdrawals(types.GetWithdrawalsReq{
		From:   &from,
		To:     &to,
		Limit:  &limit,
		Offset: &offset,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

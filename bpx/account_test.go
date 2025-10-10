package bpx

import (
	"testing"

	"bpx-api-client-go/types"

	"github.com/stretchr/testify/require"
)

func TestAccount_GetAccount(t *testing.T) {
	c := NewClient(key, secret).Account()
	resp, err := c.GetAccount(types.GetAccountReq{})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestAccount_UpdateAccount(t *testing.T) {
	c := NewClient(key, secret).Account()
	err := c.UpdateAccount(types.UpdateAccountReq{
		AutoBorrowSettlements: true,
		AutoLend:              false,
		AutoRepayBorrows:      false,
		LeverageLimit:         "15",
	})
	require.NoError(t, err)
}

func TestAccount_ConvertDust(t *testing.T) {
	c := NewClient(key, secret).Account()
	err := c.ConvertDust(types.ConvertDustReq{
		Symbol: "USDT",
	})
	require.NoError(t, err)
}

func TestAccount_MaxBorrowQuantity(t *testing.T) {
	c := NewClient(key, secret).Account()
	resp, err := c.GetMaxBorrowQuantity(types.MaxBorrowQuantityReq{
		Symbol: "USDC",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestAccount_GetMaxOrderQuantity(t *testing.T) {
	c := NewClient(key, secret).Account()
	resp, err := c.GetMaxOrderQuantity(types.MaxOrderQuantityReq{
		Symbol:          "ETH_USDC_PERP",
		Side:            "Bid",
		Price:           "",
		ReduceOnly:      false,
		AutoBorrow:      false,
		AutoBorrowRepay: false,
		AutoLendRedeem:  false,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = c.GetMaxOrderQuantity(types.MaxOrderQuantityReq{
		Symbol:          "ETH_USDC_PERP",
		Side:            "Ask",
		Price:           "4321.12",
		ReduceOnly:      true,
		AutoBorrow:      true,
		AutoBorrowRepay: true,
		AutoLendRedeem:  true,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestAccount_GetMaxWithdrawalQuantity(t *testing.T) {
	c := NewClient(key, secret).Account()
	resp, err := c.GetMaxWithdrawalQuantity(types.MaxWithdrawalQuantityReq{
		Symbol:         "USDC",
		AutoBorrow:     nil,
		AutoLendRedeem: nil,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	b, r := false, false
	resp, err = c.GetMaxWithdrawalQuantity(types.MaxWithdrawalQuantityReq{
		Symbol:         "USDC",
		AutoBorrow:     &b,
		AutoLendRedeem: &r,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	b, r = true, true
	resp, err = c.GetMaxWithdrawalQuantity(types.MaxWithdrawalQuantityReq{
		Symbol:         "USDC",
		AutoBorrow:     &b,
		AutoLendRedeem: &r,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	_, err = c.GetMaxWithdrawalQuantity(types.MaxWithdrawalQuantityReq{
		Symbol:         "ABC",
		AutoBorrow:     &b,
		AutoLendRedeem: &r,
	})
	require.Error(t, err)
}

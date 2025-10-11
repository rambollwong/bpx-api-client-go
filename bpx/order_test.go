package bpx

import (
	"testing"

	"github.com/rambollwong/bpx-api-client-go/types"

	"github.com/stretchr/testify/require"
)

func TestOrder_GetOpenOrder(t *testing.T) {
	c := NewClient(key, secret).Order()
	var orderId = "12169809136"
	resp, err := c.GetOpenOrder(types.GetOpenOrderReq{
		OrderId: &orderId,
		Symbol:  "DOGE_USDC_PERP",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestOrder_ExecuteOrder(t *testing.T) {
	c := NewClient(key, secret).Order()
	quantity := "5"
	postOnly := true
	price := "0.25070"
	resp, err := c.ExecuteOrder(types.ExecuteOrderReq{
		OrderType: "Limit",
		Symbol:    "DOGE_USDC_PERP",
		Side:      types.SideBid,
		Quantity:  &quantity,
		PostOnly:  &postOnly,
		Price:     &price,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestOrder_CancelOrder(t *testing.T) {
	c := NewClient(key, secret).Order()
	var orderId = "12744882364"
	resp, err := c.CancelOrder(types.CancelOpenOrderReq{
		OrderId: &orderId,
		Symbol:  "DOGE_USDC_PERP",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestOrder_ExecuteOrders(t *testing.T) {
	c := NewClient(key, secret).Order()
	quantity := "5"
	postOnly := true
	bidPrice := "0.23"
	askPrice := "0.28"
	resp, err := c.ExecuteOrders(types.ExecuteOrdersReq{{
		OrderType: "Limit",
		Symbol:    "DOGE_USDC_PERP",
		Side:      types.SideBid,
		Quantity:  &quantity,
		PostOnly:  &postOnly,
		Price:     &bidPrice,
	}, {
		OrderType: "Limit",
		Symbol:    "DOGE_USDC_PERP",
		Side:      types.SideAsk,
		Quantity:  &quantity,
		PostOnly:  &postOnly,
		Price:     &askPrice,
	},
	},
	)
	require.NoError(t, err)
	require.NotNil(t, resp)

	// 12746420652
	// 12746420653
}

func TestOrder_GetOpenOrders(t *testing.T) {
	c := NewClient(key, secret).Order()
	resp, err := c.GetOpenOrders(types.GetOpenOrdersReq{})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestOrder_CancelOpenOrders(t *testing.T) {
	c := NewClient(key, secret).Order()
	resp, err := c.CancelOpenOrders(types.CancelOpenOrdersReq{
		Symbol: "DOGE_USDC_PERP",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

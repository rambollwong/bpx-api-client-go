package bpx

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWsClient_Connect(t *testing.T) {
	ws := NewWsClient(key, secret)
	var triggerOnConnect bool
	ws.OnConnect(func() {
		triggerOnConnect = true
	})
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	require.True(t, triggerOnConnect)
}

func TestWsClient_Disconnect(t *testing.T) {
	ws := NewWsClient(key, secret)
	var e error
	ws.OnDisconnect(func(err error) {
		e = err
	})
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	ws.Disconnect(errors.New("test"))

	require.Error(t, e)
}

func TestWsClient_SubscribeDepth(t *testing.T) {
	ws := NewWsClient(key, secret)
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	var depthC = make(chan WsDepthData)
	err = ws.SubscribeDepth("ETH_USDC_PERP", "", func(data WsDepthData) {
		depthC <- data
		ws.Disconnect(nil)
	})
	require.NoError(t, err)

	select {
	case <-time.After(time.Second * 10):
		require.Fail(t, "timeout")
	case d := <-depthC:
		require.NotEmpty(t, d)
	case err := <-ws.ErrChan():
		require.NoError(t, err)
	}
}

func TestWsClient_SubscribeBookTicker(t *testing.T) {
	ws := NewWsClient(key, secret)
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	var c = make(chan WsBookTickerData)
	err = ws.SubscribeBookTicker("ETH_USDC", func(data WsBookTickerData) {
		c <- data
		ws.Disconnect(nil)
	})
	require.NoError(t, err)

	select {
	case <-time.After(time.Second * 30):
		require.Fail(t, "timeout")
	case v := <-c:
		require.NotEmpty(t, v)
	case err := <-ws.ErrChan():
		require.NoError(t, err)
	}
}

func TestWsClient_SubscribeKline(t *testing.T) {
	ws := NewWsClient(key, secret)
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	var c = make(chan WsKlineData)
	err = ws.SubscribeKline("ETH_USDC_PERP", "1m", func(data WsKlineData) {
		c <- data
		ws.Disconnect(nil)
	})
	require.NoError(t, err)

	select {
	case <-time.After(time.Second * 30):
		require.Fail(t, "timeout")
	case v := <-c:
		require.NotEmpty(t, v)
	case err := <-ws.ErrChan():
		require.NoError(t, err)
	}
}

func TestWsClient_SubscribeLiquidation(t *testing.T) {
	ws := NewWsClient(key, secret)
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	var c = make(chan WsLiquidationData)
	err = ws.SubscribeLiquidation("ETH_USDC_PERP", func(data WsLiquidationData) {
		c <- data
		ws.Disconnect(nil)
	})
	require.NoError(t, err)

	select {
	case <-time.After(time.Second * 120):
		require.Fail(t, "timeout")
	case v := <-c:
		require.NotEmpty(t, v)
	case err := <-ws.ErrChan():
		require.NoError(t, err)
	}
}

func TestWsClient_SubscribeMarkPrice(t *testing.T) {
	ws := NewWsClient(key, secret)
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	var c = make(chan WsMarkPriceData)
	err = ws.SubscribeMarkPrice("ETH_USDC_PERP", func(data WsMarkPriceData) {
		c <- data
		ws.Disconnect(nil)
	})
	require.NoError(t, err)

	select {
	case <-time.After(time.Second * 30):
		require.Fail(t, "timeout")
	case v := <-c:
		require.NotEmpty(t, v)
	case err := <-ws.ErrChan():
		require.NoError(t, err)
	}
}

func TestWsClient_SubscribeTicker(t *testing.T) {
	ws := NewWsClient(key, secret)
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	var c = make(chan WsTickerData)
	err = ws.SubscribeTicker("ETH_USDC_PERP", func(data WsTickerData) {
		c <- data
		ws.Disconnect(nil)
	})
	require.NoError(t, err)

	select {
	case <-time.After(time.Second * 30):
		require.Fail(t, "timeout")
	case v := <-c:
		require.NotEmpty(t, v)
	case err := <-ws.ErrChan():
		require.NoError(t, err)
	}
}

func TestWsClient_SubscribeOpenInterest(t *testing.T) {
	ws := NewWsClient(key, secret)
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	var c = make(chan WsOpenInterestData)
	err = ws.SubscribeOpenInterest("ETH_USDC_PERP", func(data WsOpenInterestData) {
		c <- data
		ws.Disconnect(nil)
	})
	require.NoError(t, err)

	select {
	case <-time.After(time.Second * 30):
		require.Fail(t, "timeout")
	case v := <-c:
		require.NotEmpty(t, v)
	case err := <-ws.ErrChan():
		require.NoError(t, err)
	}
}

func TestWsClient_SubscribeTrade(t *testing.T) {
	ws := NewWsClient(key, secret)
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	var c = make(chan WsTradeData)
	err = ws.SubscribeTrade("ETH_USDC_PERP", func(data WsTradeData) {
		c <- data
		ws.Disconnect(nil)
	})
	require.NoError(t, err)

	select {
	case <-time.After(time.Second * 30):
		require.Fail(t, "timeout")
	case v := <-c:
		require.NotEmpty(t, v)
	case err := <-ws.ErrChan():
		require.NoError(t, err)
	}
}

func TestWsClient_SubscribeOrderUpdate(t *testing.T) {
	ws := NewWsClient(key, secret)
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	var c = make(chan WsOrderData)
	err = ws.SubscribeOrderUpdate("", func(data WsOrderData) {
		c <- data
		ws.Disconnect(nil)
	})
	require.NoError(t, err)

	select {
	case <-time.After(time.Second * 120):
		require.Fail(t, "timeout")
	case v := <-c:
		require.NotEmpty(t, v)
	case err := <-ws.ErrChan():
		require.NoError(t, err)
	}
}

func TestWsClient_SubscribePositionUpdate(t *testing.T) {
	ws := NewWsClient(key, secret)
	err := ws.Connect(context.Background())
	require.NoError(t, err)
	var c = make(chan WsPositionData)
	err = ws.SubscribePositionUpdate("", func(data WsPositionData) {
		c <- data
		ws.Disconnect(nil)
	})
	require.NoError(t, err)

	select {
	case <-time.After(time.Second * 120):
		require.Fail(t, "timeout")
	case v := <-c:
		require.NotEmpty(t, v)
	case err := <-ws.ErrChan():
		require.NoError(t, err)
	}
}

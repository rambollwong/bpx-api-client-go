package bpx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rambollwong/bpx-api-client-go/types"
)

const (
	ApiGetRecentTrades     = "/api/v1/trades"
	ApiGetHistoricalTrades = "/api/v1/trades/history"
)

type Trades struct {
	c *Client
}

func (c Trades) GetRecentTradesWithContext(ctx context.Context, req types.GetRecentTradesReq) (resp []types.Trade, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetRecentTrades)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetRecentTradesReq, []types.Trade](ctx, http.MethodGet, u, req, c.c)
	return
}

func (c Trades) GetRecentTrades(req types.GetRecentTradesReq) ([]types.Trade, error) {
	return c.GetRecentTradesWithContext(context.Background(), req)
}

func (c Trades) GetHistoricalTradesWithContext(ctx context.Context, req types.GetHistoricalTradesReq) (resp []types.Trade, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetHistoricalTrades)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetHistoricalTradesReq, []types.Trade](ctx, http.MethodGet, u, req, c.c)
	return
}
func (c Trades) GetHistoricalTrades(req types.GetHistoricalTradesReq) ([]types.Trade, error) {
	return c.GetHistoricalTradesWithContext(context.Background(), req)
}

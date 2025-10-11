package bpx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rambollwong/bpx-api-client-go/types"
)

const (
	ApiOrder  = "/api/v1/order"
	ApiOrders = "/api/v1/orders"
)

type Order struct {
	*Client
}

func (c Order) GetOpenOrderWithContext(ctx context.Context, req types.GetOpenOrderReq) (resp *types.Order, err error) {
	u, err := url.JoinPath(EndpointApi, ApiOrder)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetOpenOrderReq, *types.Order](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c Order) GetOpenOrder(req types.GetOpenOrderReq) (*types.Order, error) {
	return c.GetOpenOrderWithContext(c.ctx, req)
}

func (c Order) ExecuteOrderWithContext(ctx context.Context, req types.ExecuteOrderReq) (resp *types.Order, err error) {
	u, err := url.JoinPath(EndpointApi, ApiOrder)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.ExecuteOrderReq, *types.Order](ctx, http.MethodPost, u, req, c.Client)
	return
}

func (c Order) ExecuteOrder(req types.ExecuteOrderReq) (*types.Order, error) {
	return c.ExecuteOrderWithContext(c.ctx, req)
}

func (c Order) CancelOrderWithContext(ctx context.Context, req types.CancelOpenOrderReq) (resp *types.Order, err error) {
	u, err := url.JoinPath(EndpointApi, ApiOrder)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.CancelOpenOrderReq, *types.Order](ctx, http.MethodDelete, u, req, c.Client)
	return
}

func (c Order) CancelOrder(req types.CancelOpenOrderReq) (*types.Order, error) {
	return c.CancelOrderWithContext(c.ctx, req)
}

func (c Order) ExecuteOrdersWithContext(ctx context.Context, req types.ExecuteOrdersReq) (resp []types.Order, err error) {
	u, err := url.JoinPath(EndpointApi, ApiOrders)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.ExecuteOrdersReq, []types.Order](ctx, http.MethodPost, u, req, c.Client)
	return
}

func (c Order) ExecuteOrders(req types.ExecuteOrdersReq) ([]types.Order, error) {
	return c.ExecuteOrdersWithContext(c.ctx, req)
}

func (c Order) GetOpenOrdersWithContext(ctx context.Context, req types.GetOpenOrdersReq) (resp []types.Order, err error) {
	u, err := url.JoinPath(EndpointApi, ApiOrders)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetOpenOrdersReq, []types.Order](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c Order) GetOpenOrders(req types.GetOpenOrdersReq) ([]types.Order, error) {
	return c.GetOpenOrdersWithContext(c.ctx, req)
}

func (c Order) CancelOpenOrdersWithContext(ctx context.Context, req types.CancelOpenOrdersReq) (resp []types.Order, err error) {
	u, err := url.JoinPath(EndpointApi, ApiOrders)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.CancelOpenOrdersReq, []types.Order](ctx, http.MethodDelete, u, req, c.Client)
	return
}

func (c Order) CancelOpenOrders(req types.CancelOpenOrdersReq) ([]types.Order, error) {
	return c.CancelOpenOrdersWithContext(c.ctx, req)
}

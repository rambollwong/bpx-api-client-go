package bpx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rambollwong/bpx-api-client-go/types"
)

const (
	ApiSystemStatus        = "/api/v1/status"
	ApiSystemPing          = "/api/v1/ping"
	ApiSystemGetSystemTime = "/api/v1/time"
	ApiSystemGetWallets    = "/api/v1/wallets"
)

type System struct {
	*Client
}

func (c System) StatusWithContext(ctx context.Context) (resp *types.Status, err error) {
	u, err := url.JoinPath(EndpointApi, ApiSystemStatus)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[any, *types.Status](ctx, http.MethodGet, u, nil, c.Client)
	return
}

func (c System) Status() (*types.Status, error) {
	return c.StatusWithContext(c.ctx)
}

func (c System) PingWithContext(ctx context.Context) (resp string, err error) {
	u, err := url.JoinPath(EndpointApi, ApiSystemPing)
	if err != nil {
		return "", err
	}
	resp, _, err = DoRequest[any, string](ctx, http.MethodGet, u, nil, c.Client)
	return
}

func (c System) Ping() (string, error) {
	return c.PingWithContext(c.ctx)
}

func (c System) TimeWithContext(ctx context.Context) (resp string, err error) {
	u, err := url.JoinPath(EndpointApi, ApiSystemGetSystemTime)
	if err != nil {
		return "", err
	}
	resp, _, err = DoRequest[any, string](ctx, http.MethodGet, u, nil, c.Client)
	return
}

func (c System) Time() (string, error) {
	return c.TimeWithContext(c.ctx)
}

func (c System) GetWalletsWithContext(ctx context.Context) (resp []types.Wallet, err error) {
	u, err := url.JoinPath(EndpointApi, ApiSystemGetWallets)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[any, []types.Wallet](ctx, http.MethodGet, u, nil, c.Client)
	return
}

func (c System) GetWallets() ([]types.Wallet, error) {
	return c.GetWalletsWithContext(c.ctx)
}

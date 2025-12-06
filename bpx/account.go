package bpx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rambollwong/bpx-api-client-go/types"
)

const (
	ApiAccount              = "/api/v1/account"
	ApiAccountConvertDust   = "/api/v1/account/convertDust"
	ApiAccountLimitBorrow   = "/api/v1/account/limits/borrow"
	ApiAccountLimitOrder    = "/api/v1/account/limits/order"
	ApiAccountLimitWithdraw = "/api/v1/account/limits/withdrawal"
)

type Account struct {
	c *Client
}

func (c Account) GetAccountWithContext(ctx context.Context, req types.GetAccountReq) (resp *types.Account, err error) {
	u, err := url.JoinPath(EndpointApi, ApiAccount)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetAccountReq, *types.Account](ctx, http.MethodGet, u, req, c.c)
	return
}

func (c Account) GetAccount(req types.GetAccountReq) (resp *types.Account, err error) {
	return c.GetAccountWithContext(context.Background(), req)
}

func (c Account) UpdateAccountWithContext(ctx context.Context, req types.UpdateAccountReq) (err error) {
	u, err := url.JoinPath(EndpointApi, ApiAccount)
	if err != nil {
		return err
	}
	_, _, err = DoRequest[types.UpdateAccountReq, struct{}](ctx, http.MethodPatch, u, req, c.c)
	return
}

func (c Account) UpdateAccount(req types.UpdateAccountReq) (err error) {
	return c.UpdateAccountWithContext(context.Background(), req)
}

func (c Account) ConvertDustWithContext(ctx context.Context, req types.ConvertDustReq) (err error) {
	u, err := url.JoinPath(EndpointApi, ApiAccountConvertDust)
	if err != nil {
		return err
	}
	_, _, err = DoRequest[types.ConvertDustReq, struct{}](ctx, http.MethodPost, u, req, c.c)
	return
}

func (c Account) ConvertDust(req types.ConvertDustReq) (err error) {
	return c.ConvertDustWithContext(context.Background(), req)
}

func (c Account) GetMaxBorrowQuantityWithContext(ctx context.Context, req types.MaxBorrowQuantityReq) (resp *types.MaxBorrowQuantityResp, err error) {
	u, err := url.JoinPath(EndpointApi, ApiAccountLimitBorrow)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.MaxBorrowQuantityReq, *types.MaxBorrowQuantityResp](ctx, http.MethodGet, u, req, c.c)
	return
}

func (c Account) GetMaxBorrowQuantity(req types.MaxBorrowQuantityReq) (resp *types.MaxBorrowQuantityResp, err error) {
	return c.GetMaxBorrowQuantityWithContext(context.Background(), req)
}

func (c Account) GetMaxOrderQuantityWithContext(ctx context.Context, req types.MaxOrderQuantityReq) (resp *types.MaxOrderQuantityResp, err error) {
	u, err := url.JoinPath(EndpointApi, ApiAccountLimitOrder)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.MaxOrderQuantityReq, *types.MaxOrderQuantityResp](ctx, http.MethodGet, u, req, c.c)
	return
}

func (c Account) GetMaxOrderQuantity(req types.MaxOrderQuantityReq) (resp *types.MaxOrderQuantityResp, err error) {
	return c.GetMaxOrderQuantityWithContext(context.Background(), req)
}

func (c Account) GetMaxWithdrawalQuantityWithContext(ctx context.Context, req types.MaxWithdrawalQuantityReq) (resp *types.MaxWithdrawalQuantityResp, err error) {
	u, err := url.JoinPath(EndpointApi, ApiAccountLimitWithdraw)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.MaxWithdrawalQuantityReq, *types.MaxWithdrawalQuantityResp](ctx, http.MethodGet, u, req, c.c)
	return
}

func (c Account) GetMaxWithdrawalQuantity(req types.MaxWithdrawalQuantityReq) (resp *types.MaxWithdrawalQuantityResp, err error) {
	return c.GetMaxWithdrawalQuantityWithContext(context.Background(), req)
}

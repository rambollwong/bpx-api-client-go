package bpx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rambollwong/bpx-api-client-go/types"
)

const (
	ApiGetBalances       = "/api/v1/capital"
	ApiGetCollateral     = "/api/v1/capital/collateral"
	ApiGetDeposits       = "/wapi/v1/capital/deposits"
	ApiGetDepositAddress = "/wapi/v1/capital/deposit/address"
	ApiGetWithdrawals    = "/wapi/v1/capital/withdrawals"
)

type Capital struct {
	c *Client
}

func (c Capital) GetBalancesWithContext(ctx context.Context, req types.GetBalancesReq) (resp types.GetBalancesResp, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetBalances)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetBalancesReq, types.GetBalancesResp](ctx, http.MethodGet, u, req, c.c)
	return
}

func (c Capital) GetBalances() (resp types.GetBalancesResp, err error) {
	return c.GetBalancesWithContext(context.Background(), types.GetBalancesReq{})
}

func (c Capital) GetCollateralWithContext(ctx context.Context, req types.GetCollateralReq) (resp *types.GetCollateralResp, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetCollateral)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetCollateralReq, *types.GetCollateralResp](ctx, http.MethodGet, u, req, c.c)
	return
}

func (c Capital) GetCollateral(req types.GetCollateralReq) (resp *types.GetCollateralResp, err error) {
	return c.GetCollateralWithContext(context.Background(), req)
}

func (c Capital) GetDepositsWithContext(ctx context.Context, req types.GetDepositsReq) (resp []types.Deposit, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetDeposits)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetDepositsReq, []types.Deposit](ctx, http.MethodGet, u, req, c.c)
	return
}

func (c Capital) GetDeposits(req types.GetDepositsReq) ([]types.Deposit, error) {
	return c.GetDepositsWithContext(context.Background(), req)
}

func (c Capital) GetDepositAddressWithContext(ctx context.Context, req types.GetDepositAddressReq) (resp *types.GetDepositAddressResp, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetDepositAddress)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetDepositAddressReq, *types.GetDepositAddressResp](ctx, http.MethodGet, u, req, c.c)
	return
}

func (c Capital) GetDepositAddress(req types.GetDepositAddressReq) (resp *types.GetDepositAddressResp, err error) {
	return c.GetDepositAddressWithContext(context.Background(), req)
}

func (c Capital) GetWithdrawalsWithContext(ctx context.Context, req types.GetWithdrawalsReq) (resp []types.Withdrawal, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetWithdrawals)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetWithdrawalsReq, []types.Withdrawal](ctx, http.MethodGet, u, req, c.c)
	return
}

func (c Capital) GetWithdrawals(req types.GetWithdrawalsReq) ([]types.Withdrawal, error) {
	return c.GetWithdrawalsWithContext(context.Background(), req)
}

func (c Capital) RequestWithdrawalWithContext(ctx context.Context, req types.RequestWithdrawalReq) (resp *types.Withdrawal, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetWithdrawals)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.RequestWithdrawalReq, *types.Withdrawal](ctx, http.MethodPost, u, req, c.c)
	return
}

func (c Capital) RequestWithdrawal(req types.RequestWithdrawalReq) (resp *types.Withdrawal, err error) {
	return c.RequestWithdrawalWithContext(context.Background(), req)
}

package bpx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rambollwong/bpx-api-client-go/types"
)

const (
	ApiAssets     = "/api/v1/assets"
	ApiCollateral = "/api/v1/collateral"
)

type Assets struct {
	*Client
}

func (c Assets) GetAssetsWithContext(ctx context.Context) (resp []types.AssetResp, err error) {
	u, err := url.JoinPath(EndpointApi, ApiAssets)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[any, []types.AssetResp](ctx, http.MethodGet, u, nil, c.Client)
	return
}

func (c Assets) GetAssets() ([]types.AssetResp, error) {
	return c.GetAssetsWithContext(context.Background())
}

func (c Assets) GetCollateralsWithContext(ctx context.Context) (resp []types.CollateralResp, err error) {
	u, err := url.JoinPath(EndpointApi, ApiCollateral)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[any, []types.CollateralResp](ctx, http.MethodGet, u, nil, c.Client)
	return

}

func (c Assets) GetCollaterals() ([]types.CollateralResp, error) {
	return c.GetCollateralsWithContext(context.Background())
}

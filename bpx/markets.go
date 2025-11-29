package bpx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rambollwong/bpx-api-client-go/types"
)

const (
	ApiGetMarkets              = "/api/v1/markets"
	ApiGetMarket               = "/api/v1/market"
	ApiGetTicker               = "/api/v1/ticker"
	ApiGetTickers              = "/api/v1/tickers"
	ApiGetDepth                = "/api/v1/depth"
	ApiGetKlines               = "/api/v1/klines"
	ApiGetAllMarkPrices        = "/api/v1/markPrices"
	ApiGetOpenInterest         = "/api/v1/openInterest"
	ApiGetFundingIntervalRates = "/api/v1/fundingRates"
)

type Markets struct {
	*Client
}

func (c Markets) GetMarketsWithContext(ctx context.Context) (resp []types.Market, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetMarkets)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[any, []types.Market](ctx, http.MethodGet, u, nil, c.Client)
	return
}

func (c Markets) GetMarkets() ([]types.Market, error) {
	return c.GetMarketsWithContext(context.Background())
}

func (c Markets) GetMarketWithContext(ctx context.Context, req types.GetMarketReq) (resp *types.Market, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetMarket)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetMarketReq, *types.Market](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c Markets) GetMarket(req types.GetMarketReq) (*types.Market, error) {
	return c.GetMarketWithContext(context.Background(), req)
}

func (c Markets) GetTickerWithContext(ctx context.Context, req types.GetTickerReq) (resp *types.Ticker, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetTicker)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetTickerReq, *types.Ticker](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c Markets) GetTicker(req types.GetTickerReq) (*types.Ticker, error) {
	return c.GetTickerWithContext(context.Background(), req)
}

func (c Markets) GetTickersWithContext(ctx context.Context, req types.GetTickersReq) (resp []types.Ticker, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetTickers)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetTickersReq, []types.Ticker](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c Markets) GetTickers(req types.GetTickersReq) ([]types.Ticker, error) {
	return c.GetTickersWithContext(context.Background(), req)
}

func (c Markets) GetDepthWithContext(ctx context.Context, req types.GetDepthReq) (resp *types.Depth, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetDepth)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetDepthReq, *types.Depth](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c Markets) GetDepth(req types.GetDepthReq) (*types.Depth, error) {
	return c.GetDepthWithContext(context.Background(), req)
}

func (c Markets) GetKlinesWithContext(ctx context.Context, req types.GetKlinesReq) (resp []types.Kline, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetKlines)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetKlinesReq, []types.Kline](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c Markets) GetKlines(req types.GetKlinesReq) ([]types.Kline, error) {
	return c.GetKlinesWithContext(context.Background(), req)
}

func (c Markets) GetAllMarkPricesWithContext(ctx context.Context, req types.GetAllMarkPricesReq) (resp []types.MarkPrice, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetAllMarkPrices)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetAllMarkPricesReq, []types.MarkPrice](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c Markets) GetAllMarketPrices(req types.GetAllMarkPricesReq) ([]types.MarkPrice, error) {
	return c.GetAllMarkPricesWithContext(context.Background(), req)
}

func (c Markets) GetOpenInterestWithContext(ctx context.Context, req types.GetOpenInterestReq) (resp []types.OpenInterest, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetOpenInterest)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetOpenInterestReq, []types.OpenInterest](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c Markets) GetOpenInterest(req types.GetOpenInterestReq) ([]types.OpenInterest, error) {
	return c.GetOpenInterestWithContext(context.Background(), req)
}

func (c Markets) GetFundingIntervalRatesWithContext(ctx context.Context, req types.GetFundingIntervalRatesReq) (resp types.GetFundingIntervalRatesResp, respHeader *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetFundingIntervalRates)
	if err != nil {
		return nil, nil, err
	}
	return DoRequest[types.GetFundingIntervalRatesReq, types.GetFundingIntervalRatesResp](ctx, http.MethodGet, u, req, c.Client)
}

func (c Markets) GetFundingIntervalRates(req types.GetFundingIntervalRatesReq) (types.GetFundingIntervalRatesResp, *types.ResponseHeaders, error) {
	return c.GetFundingIntervalRatesWithContext(context.Background(), req)
}

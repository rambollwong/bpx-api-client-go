package bpx

import (
	"context"
	"net/http"
	"net/url"

	"bpx-api-client-go/types"
)

const (
	ApiBorrowLendMarkets        = "/api/v1/borrowLend/markets"
	ApiBorrowLendMarketsHistory = "/api/v1/borrowLend/markets/history"
)

type BorrowLendMarkets struct {
	*Client
}

func (c BorrowLendMarkets) GetBorrowLendMarketsWithContext(ctx context.Context) (resp []types.BorrowLendMarketResp, err error) {
	u, err := url.JoinPath(EndpointApi, ApiBorrowLendMarkets)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[any, []types.BorrowLendMarketResp](ctx, http.MethodGet, u, nil, c.Client)
	return
}

func (c BorrowLendMarkets) GetBorrowLendMarkets() ([]types.BorrowLendMarketResp, error) {
	return c.GetBorrowLendMarketsWithContext(c.ctx)
}

func (c BorrowLendMarkets) GetBorrowLendMarketsHistoryWithContext(ctx context.Context, req types.BorrowLendMarketHistoryReq) (resp []types.BorrowLendMarketHistoryResp, err error) {
	u, err := url.JoinPath(EndpointApi, ApiBorrowLendMarketsHistory)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.BorrowLendMarketHistoryReq, []types.BorrowLendMarketHistoryResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c BorrowLendMarkets) GetBorrowLendMarketsHistory(req types.BorrowLendMarketHistoryReq) ([]types.BorrowLendMarketHistoryResp, error) {
	return c.GetBorrowLendMarketsHistoryWithContext(c.ctx, req)
}

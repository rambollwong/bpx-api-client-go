package bpx

import (
	"context"
	"net/http"
	"net/url"

	"bpx-api-client-go/types"
)

const (
	ApiGetBorrowHistory         = "/wapi/v1/history/borrowLend"
	ApiGetInterestHistory       = "/wapi/v1/history/interest"
	ApiGetBorrowPositionHistory = "/wapi/v1/history/borrowLend/positions"
	ApiGetDustConversionHistory = "/wapi/v1/history/dust"
	ApiGetFillHistory           = "/wapi/v1/history/fills"
	ApiGetFundingPayments       = "/wapi/v1/history/funding"
	ApiGetOrderHistory          = "/wapi/v1/history/orders"
	ApiGetRfqHistory            = "/wapi/v1/history/rfq"
	ApiGetQuoteHistory          = "/wapi/v1/history/quote"
	ApiGetSettlementHistory     = "/wapi/v1/history/settlement"
	ApiGetStrategyHistory       = "/wapi/v1/history/strategies"
)

type History struct {
	*Client
}

func (c History) GetBorrowHistoryWithContext(ctx context.Context, req types.GetBorrowHistoryReq) (resp types.GetBorrowHistoryResp, rh *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetBorrowHistory)
	if err != nil {
		return nil, nil, err
	}
	resp, rh, err = DoRequest[types.GetBorrowHistoryReq, types.GetBorrowHistoryResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c History) GetBorrowHistory(req types.GetBorrowHistoryReq) (resp types.GetBorrowHistoryResp, rh *types.ResponseHeaders, err error) {
	return c.GetBorrowHistoryWithContext(c.ctx, req)
}

func (c History) GetInterestHistoryWithContext(ctx context.Context, req types.GetInterestHistoryReq) (resp types.GetInterestHistoryResp, rh *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetInterestHistory)
	if err != nil {
		return nil, nil, err
	}
	resp, rh, err = DoRequest[types.GetInterestHistoryReq, types.GetInterestHistoryResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c History) GetInterestHistory(req types.GetInterestHistoryReq) (resp types.GetInterestHistoryResp, rh *types.ResponseHeaders, err error) {
	return c.GetInterestHistoryWithContext(c.ctx, req)
}

func (c History) GetBorrowPositionHistoryWithContext(ctx context.Context, req types.GetBorrowPositionHistoryReq) (resp types.GetBorrowPositionHistoryResp, rh *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetBorrowPositionHistory)
	if err != nil {
		return nil, nil, err
	}
	resp, rh, err = DoRequest[types.GetBorrowPositionHistoryReq, types.GetBorrowPositionHistoryResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c History) GetBorrowPositionHistory(req types.GetBorrowPositionHistoryReq) (resp types.GetBorrowPositionHistoryResp, rh *types.ResponseHeaders, err error) {
	return c.GetBorrowPositionHistoryWithContext(c.ctx, req)
}

func (c History) GetDustConversionHistoryWithContext(ctx context.Context, req types.GetDustConversionHistoryReq) (resp types.GetDustConversionHistoryResp, rh *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetDustConversionHistory)
	if err != nil {
		return nil, nil, err
	}
	resp, rh, err = DoRequest[types.GetDustConversionHistoryReq, types.GetDustConversionHistoryResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c History) GetDustConversionHistory(req types.GetDustConversionHistoryReq) (resp types.GetDustConversionHistoryResp, rh *types.ResponseHeaders, err error) {
	return c.GetDustConversionHistoryWithContext(c.ctx, req)
}

func (c History) GetFillHistoryWithContext(ctx context.Context, req types.GetFillHistoryReq) (resp types.GetFillHistoryResp, rh *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetFillHistory)
	if err != nil {
		return nil, nil, err
	}
	resp, rh, err = DoRequest[types.GetFillHistoryReq, types.GetFillHistoryResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c History) GetFillHistory(req types.GetFillHistoryReq) (resp types.GetFillHistoryResp, rh *types.ResponseHeaders, err error) {
	return c.GetFillHistoryWithContext(c.ctx, req)
}

func (c History) GetFundingPaymentsWithContext(ctx context.Context, req types.GetFundingPaymentsReq) (resp types.GetFundingPaymentsResp, rh *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetFundingPayments)
	if err != nil {
		return nil, nil, err
	}
	resp, rh, err = DoRequest[types.GetFundingPaymentsReq, types.GetFundingPaymentsResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c History) GetFundingPayments(req types.GetFundingPaymentsReq) (resp types.GetFundingPaymentsResp, rh *types.ResponseHeaders, err error) {
	return c.GetFundingPaymentsWithContext(c.ctx, req)
}

func (c History) GetOrderHistoryWithContext(ctx context.Context, req types.GetOrderHistoryReq) (resp types.GetOrderHistoryResp, rh *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetOrderHistory)
	if err != nil {
		return nil, nil, err
	}
	resp, rh, err = DoRequest[types.GetOrderHistoryReq, types.GetOrderHistoryResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c History) GetOrderHistory(req types.GetOrderHistoryReq) (resp types.GetOrderHistoryResp, rh *types.ResponseHeaders, err error) {
	return c.GetOrderHistoryWithContext(c.ctx, req)
}

func (c History) GetRfqHistoryWithContext(ctx context.Context, req types.GetRfqHistoryReq) (resp types.GetRfqHistoryResp, rh *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetRfqHistory)
	if err != nil {
		return nil, nil, err
	}
	resp, rh, err = DoRequest[types.GetRfqHistoryReq, types.GetRfqHistoryResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c History) GetRfqHistory(req types.GetRfqHistoryReq) (resp types.GetRfqHistoryResp, rh *types.ResponseHeaders, err error) {
	return c.GetRfqHistoryWithContext(c.ctx, req)
}

func (c History) GetQuoteHistoryWithContext(ctx context.Context, req types.GetQuoteHistoryReq) (resp types.GetQuoteHistoryResp, rh *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetQuoteHistory)
	if err != nil {
		return nil, nil, err
	}
	resp, rh, err = DoRequest[types.GetQuoteHistoryReq, types.GetQuoteHistoryResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c History) GetQuoteHistory(req types.GetQuoteHistoryReq) (resp types.GetQuoteHistoryResp, rh *types.ResponseHeaders, err error) {
	return c.GetQuoteHistoryWithContext(c.ctx, req)
}

func (c History) GetSettlementHistoryWithContext(ctx context.Context, req types.GetSettlementHistoryReq) (resp types.GetSettlementHistoryResp, rh *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetSettlementHistory)
	if err != nil {
		return nil, nil, err
	}
	resp, rh, err = DoRequest[types.GetSettlementHistoryReq, types.GetSettlementHistoryResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c History) GetSettlementHistory(req types.GetSettlementHistoryReq) (resp types.GetSettlementHistoryResp, rh *types.ResponseHeaders, err error) {
	return c.GetSettlementHistoryWithContext(c.ctx, req)
}

func (c History) GetStrategyHistoryWithContext(ctx context.Context, req types.GetStrategyHistoryReq) (resp types.GetStrategyHistoryResp, rh *types.ResponseHeaders, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetStrategyHistory)
	if err != nil {
		return nil, nil, err
	}
	resp, rh, err = DoRequest[types.GetStrategyHistoryReq, types.GetStrategyHistoryResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c History) GetStrategyHistory(req types.GetStrategyHistoryReq) (resp types.GetStrategyHistoryResp, rh *types.ResponseHeaders, err error) {
	return c.GetStrategyHistoryWithContext(c.ctx, req)
}

package bpx

import (
	"context"
	"net/http"
	"net/url"

	"bpx-api-client-go/types"
)

const (
	ApiBorrowLendPositions                                          = "/api/v1/borrowLend/positions"
	ApiExecuteBorrowLend                                            = "/api/v1/borrowLend"
	ApiGetAnEstimatedLiquidationPriceForPotentialBorrowLendPosition = "/api/v1/borrowLend/position/liquidationPrice"
)

type BorrowLend struct {
	*Client
}

func (c BorrowLend) GetBorrowLendPositionsWithContext(ctx context.Context, req types.GetBorrowLendPositionReq) (resp []types.BorrowLendPosition, err error) {
	u, err := url.JoinPath(EndpointApi, ApiBorrowLendPositions)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetBorrowLendPositionReq, []types.BorrowLendPosition](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c BorrowLend) GetBorrowLendPositions(req types.GetBorrowLendPositionReq) (resp []types.BorrowLendPosition, err error) {
	return c.GetBorrowLendPositionsWithContext(c.ctx, req)
}

func (c BorrowLend) ExecuteBorrowLendWithContext(ctx context.Context, req types.ExecuteBorrowLendReq) (err error) {
	u, err := url.JoinPath(EndpointApi, ApiExecuteBorrowLend)
	if err != nil {
		return err
	}
	_, _, err = DoRequest[types.ExecuteBorrowLendReq, struct{}](ctx, http.MethodPost, u, req, c.Client)
	return
}

func (c BorrowLend) ExecuteBorrowLend(req types.ExecuteBorrowLendReq) (err error) {
	return c.ExecuteBorrowLendWithContext(c.ctx, req)
}

func (c BorrowLend) GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionWithContext(
	ctx context.Context,
	req types.GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionReq,
) (resp types.GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionResp, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetAnEstimatedLiquidationPriceForPotentialBorrowLendPosition)
	if err != nil {
		return
	}
	resp, _, err = DoRequest[
		types.GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionReq,
		types.GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionResp](ctx, http.MethodGet, u, req, c.Client)
	return
}

func (c BorrowLend) GetAnEstimatedLiquidationPriceForPotentialBorrowLendPosition(
	req types.GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionReq,
) (resp types.GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionResp, err error) {
	return c.GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionWithContext(c.ctx, req)
}

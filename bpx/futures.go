package bpx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rambollwong/bpx-api-client-go/types"
)

const (
	ApiGetOpenPositions = "/api/v1/position"
)

type Futures struct {
	c *Client
}

func (c Futures) GetOpenPositionsWithContext(ctx context.Context, req types.GetOpenPositionsReq) (resp []types.Position, err error) {
	u, err := url.JoinPath(EndpointApi, ApiGetOpenPositions)
	if err != nil {
		return nil, err
	}
	resp, _, err = DoRequest[types.GetOpenPositionsReq, []types.Position](ctx, http.MethodGet, u, req, c.c)
	return

}

func (c Futures) GetOpenPositions(req types.GetOpenPositionsReq) (resp []types.Position, err error) {
	return c.GetOpenPositionsWithContext(context.Background(), req)
}

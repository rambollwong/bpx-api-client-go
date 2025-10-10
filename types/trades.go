package types

import (
	"errors"
	"fmt"
	"strings"
)

type GetRecentTradesReq struct {
	Symbol string
	Limit  uint64
}

func (req GetRecentTradesReq) Validate() error {
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	if req.Limit > 0 {
		if req.Limit < 100 {
			return errors.New("limit must be greater than or equal to 100")
		}
		if req.Limit > 1000 {
			return errors.New("limit must be less than or equal to 1000")
		}
	}
	return nil
}

func (req GetRecentTradesReq) BuildQueryParams() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("symbol=")
	strBuilder.WriteString(req.Symbol)
	if req.Limit > 0 {
		strBuilder.WriteString("&limit=")
		strBuilder.WriteString(fmt.Sprintf("%d", req.Limit))
	}
	return strBuilder.String()
}

type Trade struct {
	Id            int64  `json:"id,omitempty"`
	Price         string `json:"price,omitempty"`
	Quantity      string `json:"quantity,omitempty"`
	QuoteQuantity string `json:"quoteQuantity,omitempty"`
	Timestamp     int64  `json:"timestamp,omitempty"`
	IsBuyerMaker  bool   `json:"isBuyerMaker,omitempty"`
}

type GetHistoricalTradesReq struct {
	Symbol string
	Limit  uint64
	Offset uint64
}

func (req GetHistoricalTradesReq) Validate() error {
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	if req.Limit > 0 {
		if req.Limit < 100 {
			return errors.New("limit must be greater than or equal to 100")
		}
		if req.Limit > 1000 {
			return errors.New("limit must be less than or equal to 1000")
		}
	}
	return nil
}

func (req GetHistoricalTradesReq) BuildQueryParams() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("symbol=")
	strBuilder.WriteString(req.Symbol)
	if req.Limit > 0 {
		strBuilder.WriteString("&limit=")
		strBuilder.WriteString(fmt.Sprintf("%d", req.Limit))
	}
	if req.Offset > 0 {
		strBuilder.WriteString("&offset=")
		strBuilder.WriteString(fmt.Sprintf("%d", req.Offset))
	}
	return strBuilder.String()
}

package types

import (
	"errors"
	"strings"
)

type BorrowLendMarketResp struct {
	State                        string `json:"state"`
	AssetMarkPrice               string `json:"assetMarkPrice"`
	BorrowInterestRate           string `json:"borrowInterestRate"`
	BorrowedQuantity             string `json:"borrowedQuantity"`
	Fee                          string `json:"fee"`
	LendInterestRate             string `json:"lendInterestRate"`
	LentQuantity                 string `json:"lentQuantity"`
	MaxUtilization               string `json:"maxUtilization"`
	OpenBorrowLendLimit          string `json:"openBorrowLendLimit"`
	OptimalUtilization           string `json:"optimalUtilization"`
	Symbol                       string `json:"symbol"`
	Timestamp                    string `json:"timestamp"`
	ThrottleUtilizationThreshold string `json:"throttleUtilizationThreshold"`
	ThrottleUtilizationBound     string `json:"throttleUtilizationBound"`
	ThrottleUpdateFraction       string `json:"throttleUpdateFraction"`
	Utilization                  string `json:"utilization"`
	StepSize                     string `json:"stepSize"`
}

const (
	BorrowLendMarketHistoryInterval1D     = "1d"
	BorrowLendMarketHistoryInterval1W     = "1w"
	BorrowLendMarketHistoryInterval1Month = "1month"
	BorrowLendMarketHistoryInterval1Year  = "1year"
)

type BorrowLendMarketHistoryReq struct {
	Interval string `json:"interval"`
	Symbol   string `json:"symbol"`
}

func (req BorrowLendMarketHistoryReq) Validate() error {
	if req.Interval == "" {
		return errors.New("symbol is required")
	}
	switch req.Interval {
	case BorrowLendMarketHistoryInterval1D,
		BorrowLendMarketHistoryInterval1W,
		BorrowLendMarketHistoryInterval1Month,
		BorrowLendMarketHistoryInterval1Year: // ignore
	default:
		return errors.New("invalid interval")
	}
	return nil
}

func (req BorrowLendMarketHistoryReq) BuildQueryParams() string {
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("interval=")
	stringBuilder.WriteString(req.Interval)
	if req.Symbol != "" {
		stringBuilder.WriteString("&symbol=")
		stringBuilder.WriteString(req.Symbol)
	}
	return stringBuilder.String()
}

type BorrowLendMarketHistoryResp struct {
	BorrowInterestRate string `json:"borrowInterestRate"`
	BorrowedQuantity   string `json:"borrowedQuantity"`
	LendInterestRate   string `json:"lendInterestRate"`
	LentQuantity       string `json:"lentQuantity"`
	Timestamp          string `json:"timestamp"`
	Utilization        string `json:"utilization"`
}

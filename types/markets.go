package types

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Market struct {
	Symbol                string            `json:"symbol,omitempty"`
	BaseSymbol            string            `json:"baseSymbol,omitempty"`
	QuoteSymbol           string            `json:"quoteSymbol,omitempty"`
	MarketType            string            `json:"marketType,omitempty"`
	Filters               *OrderBookFilters `json:"filters,omitempty"`
	ImfFunction           *Function         `json:"imfFunction,omitempty"`
	MmfFunction           *Function         `json:"mmfFunction,omitempty"`
	FundingInterval       uint64            `json:"fundingInterval,omitempty"`
	FundingRateUpperBound string            `json:"fundingRateUpperBound,omitempty"`
	FundingRateLowerBound string            `json:"fundingRateLowerBound,omitempty"`
	OpenInterestLimit     string            `json:"openInterestLimit,omitempty"`
	OrderBookState        string            `json:"orderBookState,omitempty"`
	CreatedAt             string            `json:"createdAt,omitempty"`
	Visible               bool              `json:"visible,omitempty"`
}

type GetMarketReq struct {
	Symbol string
}

func (req GetMarketReq) Validate() error {
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	return nil
}

func (req GetMarketReq) BuildQueryParams() string {
	return "symbol=" + req.Symbol
}

const (
	TickerInterval1D = "1d"
	TickerInterval1W = "1w"
)

type GetTickerReq struct {
	Symbol   string
	Interval string
}

func (req GetTickerReq) Validate() error {
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	return nil
}

func (req GetTickerReq) BuildQueryParams() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("symbol=")
	strBuilder.WriteString(req.Symbol)
	if req.Interval != "" {
		strBuilder.WriteString("&interval=")
		strBuilder.WriteString(req.Interval)
	}
	return strBuilder.String()
}

type GetTickersReq struct {
	Interval string
}

func (req GetTickersReq) Validate() error {
	return nil
}

func (req GetTickersReq) BuildQueryParams() string {
	if req.Interval != "" {
		return "interval=" + req.Interval
	}
	return ""
}

type Ticker struct {
	Symbol             string `json:"symbol,omitempty"`
	FirstPrice         string `json:"firstPrice,omitempty"`
	LastPrice          string `json:"lastPrice,omitempty"`
	PriceChange        string `json:"priceChange,omitempty"`
	PriceChangePercent string `json:"priceChangePercent,omitempty"`
	High               string `json:"high,omitempty"`
	Low                string `json:"low,omitempty"`
	Volume             string `json:"volume,omitempty"`
	QuoteVolume        string `json:"quoteVolume,omitempty"`
	Trades             string `json:"trades,omitempty"`
}

type GetDepthReq struct {
	Symbol string
}

func (req GetDepthReq) Validate() error {
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	return nil
}

func (req GetDepthReq) BuildQueryParams() string {
	return "symbol=" + req.Symbol
}

type Depth struct {
	Asks         [][]string `json:"asks,omitempty"`
	Bids         [][]string `json:"bids,omitempty"`
	LastUpdateId string     `json:"lastUpdateId,omitempty"`
	Timestamp    int64      `json:"timestamp,omitempty"`
}

const (
	KlineInterval1M     = "1m"
	KlineInterval3M     = "3m"
	KlineInterval5M     = "5m"
	KlineInterval15M    = "15m"
	KlineInterval30M    = "30m"
	KlineInterval1H     = "1h"
	KlineInterval2H     = "2h"
	KlineInterval4H     = "4h"
	KlineInterval6H     = "6h"
	KlineInterval8H     = "8h"
	KlineInterval12H    = "12h"
	KlineInterval1D     = "1d"
	KlineInterval3D     = "3d"
	KlineInterval1W     = "1w"
	KlineInterval1Month = "1month"

	KlinePriceTypeLast  = "Last"
	KlinePriceTypeIndex = "Index"
	KlinePriceTypeMark  = "Mark"
)

type GetKlinesReq struct {
	Symbol    string
	Interval  string
	StartTime int64
	EndTime   int64
	PriceType string
}

func (req GetKlinesReq) Validate() error {
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	if req.Interval == "" {
		return errors.New("interval is required")
	}
	if req.StartTime <= 0 {
		return errors.New("startTime is required")
	}
	return nil
}

func (req GetKlinesReq) BuildQueryParams() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("symbol=")
	strBuilder.WriteString(req.Symbol)
	strBuilder.WriteString("&interval=")
	strBuilder.WriteString(req.Interval)
	strBuilder.WriteString("&startTime=")
	strBuilder.WriteString(fmt.Sprintf("%d", req.StartTime))
	if req.EndTime > 0 {
		strBuilder.WriteString("&endTime=")
		strBuilder.WriteString(fmt.Sprintf("%d", req.EndTime))
	}
	if req.PriceType != "" {
		strBuilder.WriteString("&priceType=")
		strBuilder.WriteString(req.PriceType)
	}
	return strBuilder.String()
}

type Kline struct {
	Start       string `json:"start,omitempty"`
	End         string `json:"end,omitempty"`
	Open        string `json:"open,omitempty"`
	High        string `json:"high,omitempty"`
	Low         string `json:"low,omitempty"`
	Close       string `json:"close,omitempty"`
	Volume      string `json:"volume,omitempty"`
	QuoteVolume string `json:"quoteVolume,omitempty"`
	Trades      string `json:"trades,omitempty"`
}
type GetAllMarkPricesReq struct {
	Symbol string
}

func (req GetAllMarkPricesReq) Validate() error {
	return nil
}

func (req GetAllMarkPricesReq) BuildQueryParams() string {
	if req.Symbol != "" {
		return "symbol=" + req.Symbol
	}
	return ""
}

type MarkPrice struct {
	FundingRate          string `json:"fundingRate,omitempty"`
	IndexPrice           string `json:"indexPrice,omitempty"`
	MarkPrice            string `json:"markPrice,omitempty"`
	NextFundingTimestamp int64  `json:"nextFundingTimestamp,omitempty"`
	Symbol               string `json:"symbol,omitempty"`
}

type GetOpenInterestReq struct {
	Symbol string
}

func (req GetOpenInterestReq) Validate() error {
	return nil
}

func (req GetOpenInterestReq) BuildQueryParams() string {
	if req.Symbol != "" {
		return "symbol=" + req.Symbol
	}
	return ""
}

type OpenInterest struct {
	Symbol       string `json:"symbol,omitempty"`
	OpenInterest string `json:"openInterest,omitempty"`
	Timestamp    int64  `json:"timestamp,omitempty"`
}

type GetFundingIntervalRatesReq struct {
	Symbol string
	Limit  uint64
	Offset uint64
}

func (req GetFundingIntervalRatesReq) Validate() error {
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

func (req GetFundingIntervalRatesReq) BuildQueryParams() string {
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

type FundingIntervalRates struct {
	Symbol               string `json:"symbol,omitempty"`
	IntervalEndTimestamp string `json:"intervalEndTimestamp,omitempty"`
	FundingRate          string `json:"fundingRate,omitempty"`
}

type GetFundingIntervalRatesResp []FundingIntervalRates

func (GetFundingIntervalRatesResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

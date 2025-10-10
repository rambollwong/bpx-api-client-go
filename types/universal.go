package types

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	SideAsk = "Ask"
	SideBid = "Bid"

	SortDirectionAsc  = "Asc"
	SortDirectionDesc = "Desc"
)

type ResponseHeaders struct {
	AccessControlExposeHeaders string `json:"ACCESS-CONTROL-EXPOSE-HEADERS,omitempty"`
	XPageCount                 uint64 `json:"X-PAGE-COUNT,omitempty"`
	XCurrentPage               uint64 `json:"X-CURRENT-PAGE,omitempty"`
	XPageSize                  uint64 `json:"X-PAGE-SIZE,omitempty"`
	XTotal                     uint64 `json:"X-TOTAL,omitempty"`
}

type CodeMessage struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Function struct {
	Type   string `json:"type,omitempty"`
	Base   string `json:"base,omitempty"`
	Factor string `json:"factor,omitempty"`
}

type HaircutFunction struct {
	Weight string `json:"weight,omitempty"`
	Kind   struct {
		Type string `json:"type,omitempty"`
	} `json:"kind,omitempty"`
}

type PriceBandMarkPrice struct {
	MaxMultiplier string `json:"maxMultiplier,omitempty"`
	MinMultiplier string `json:"minMultiplier,omitempty"`
}

type PriceBandMeanPremium struct {
	TolerancePct string `json:"tolerancePct,omitempty"`
}

type PriceFilter struct {
	MinPrice                    string                `json:"minPrice,omitempty"`
	MaxPrice                    string                `json:"maxPrice,omitempty"`
	TickSize                    string                `json:"tickSize,omitempty"`
	MaxMultiplier               string                `json:"maxMultiplier,omitempty"`
	MinMultiplier               string                `json:"minMultiplier,omitempty"`
	MaxImpactMultiplier         string                `json:"maxImpactMultiplier,omitempty"`
	MinImpactMultiplier         string                `json:"minImpactMultiplier,omitempty"`
	MeanMarkPriceBand           *PriceBandMarkPrice   `json:"meanMarkPriceBand,omitempty"`
	MeanPremiumBand             *PriceBandMeanPremium `json:"meanPremiumBand,omitempty"`
	BorrowEntryFeeMaxMultiplier string                `json:"borrowEntryFeeMaxMultiplier,omitempty"`
	BorrowEntryFeeMinMultiplier string                `json:"borrowEntryFeeMinMultiplier,omitempty"`
}

type QuantityFilter struct {
	MinQuantity string `json:"minQuantity,omitempty"`
	MaxQuantity string `json:"maxQuantity,omitempty"`
	StepSize    string `json:"stepSize,omitempty"`
}

type OrderBookFilters struct {
	Price          *PriceFilter    `json:"price,omitempty"`
	QuantityFilter *QuantityFilter `json:"quantityFilter,omitempty"`
}

type Balance struct {
	Available string `json:"available,omitempty"`
	Locked    string `json:"locked,omitempty"`
	Staked    string `json:"staked,omitempty"`
}

type Collateral struct {
	Symbol            string `json:"symbol,omitempty"`
	AssetMarkPrice    string `json:"assetMarkPrice,omitempty"`
	TotalQuantity     string `json:"totalQuantity,omitempty"`
	BalanceNotional   string `json:"balanceNotional,omitempty"`
	CollateralWeight  string `json:"collateralWeight,omitempty"`
	CollateralValue   string `json:"collateralValue,omitempty"`
	OpenOrderQuantity string `json:"openOrderQuantity,omitempty"`
	LendQuantity      string `json:"lendQuantity,omitempty"`
	AvailableQuantity string `json:"availableQuantity,omitempty"`
}

func WrapCodeMessageError(cm CodeMessage) error {
	return fmt.Errorf("api error: code: %s, message: %s", cm.Code, cm.Message)
}

type QueryParamsBuilder interface {
	BuildQueryParams() string
	Validate() error
}

type AuthenticatedRequest interface {
	QueryParamsBuilder
	Instruction() string
}

type ResponseHeadersReader interface {
	ReadResponseHeaders(header http.Header) (*ResponseHeaders, error)
}

func ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	rh = new(ResponseHeaders)
	rh.AccessControlExposeHeaders = header.Get("ACCESS-CONTROL-EXPOSE-HEADERS")
	rh.XPageCount, err = strconv.ParseUint(header.Get("X-PAGE-COUNT"), 10, 64)
	if err != nil {
		return rh, err
	}
	rh.XCurrentPage, err = strconv.ParseUint(header.Get("X-CURRENT-PAGE"), 10, 64)
	if err != nil {
		return rh, err
	}
	rh.XPageSize, err = strconv.ParseUint(header.Get("X-PAGE-SIZE"), 10, 64)
	if err != nil {
		return rh, err
	}
	rh.XTotal, err = strconv.ParseUint(header.Get("X-TOTAL"), 10, 64)
	return rh, err
}

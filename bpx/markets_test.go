package bpx

import (
	"testing"

	"bpx-api-client-go/types"

	"github.com/stretchr/testify/require"
)

func TestMarkets_GetMarkets(t *testing.T) {
	c := NewClient(key, secret)
	res, err := c.Markets().GetMarkets()
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestMarkets_GetMarket(t *testing.T) {
	c := NewClient(key, secret)
	res, err := c.Markets().GetMarket(types.GetMarketReq{Symbol: "ETH_USDC_PERP"})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestMarkets_GetTicker(t *testing.T) {
	c := NewClient(key, secret)
	res, err := c.Markets().GetTicker(types.GetTickerReq{Symbol: "ETH_USDC_PERP"})
	require.NoError(t, err)
	require.NotNil(t, res)

	_, err = c.Markets().GetTicker(types.GetTickerReq{Symbol: ""})
	require.Error(t, err)

	_, err = c.Markets().GetTicker(types.GetTickerReq{Symbol: "ABC"})
	require.Error(t, err)
}

func TestMarkets_GetTickers(t *testing.T) {
	c := NewClient(key, secret)
	res, err := c.Markets().GetTickers(types.GetTickersReq{Interval: types.TickerInterval1D})
	require.NoError(t, err)
	require.NotNil(t, res)

	res, err = c.Markets().GetTickers(types.GetTickersReq{Interval: ""})
	require.NoError(t, err)
	require.NotNil(t, res)

	_, err = c.Markets().GetTickers(types.GetTickersReq{Interval: "ABC"})
	require.Error(t, err)
}

func TestMarkets_GetDepth(t *testing.T) {
	c := NewClient(key, secret)
	res, err := c.Markets().GetDepth(types.GetDepthReq{Symbol: "ETH_USDC_PERP"})
	require.NoError(t, err)
	require.NotNil(t, res)

	_, err = c.Markets().GetDepth(types.GetDepthReq{Symbol: ""})
	require.Error(t, err)
}

func TestMarkets_GetKlines(t *testing.T) {
	c := NewClient(key, secret)
	res, err := c.Markets().GetKlines(types.GetKlinesReq{Symbol: "ETH_USDC_PERP", Interval: types.KlineInterval1D, StartTime: 1672531200})
	require.NoError(t, err)
	require.NotNil(t, res)

	_, err = c.Markets().GetKlines(types.GetKlinesReq{Symbol: "", Interval: types.KlineInterval1D, StartTime: 1672531200})
	require.Error(t, err)

	_, err = c.Markets().GetKlines(types.GetKlinesReq{Symbol: "ETH_USDC_PERP", Interval: "ABC", StartTime: 1672531200})
	require.Error(t, err)

	_, err = c.Markets().GetKlines(types.GetKlinesReq{Symbol: "ETH_USDC_PERP", Interval: types.KlineInterval1D, StartTime: 1234567890})
	require.Error(t, err)
}

func TestMarkets_GetAllMarketPrices(t *testing.T) {
	c := NewClient(key, secret)
	res, err := c.Markets().GetAllMarketPrices(types.GetAllMarkPricesReq{Symbol: "ETH_USDC_PERP"})
	require.NoError(t, err)
	require.NotNil(t, res)

	res, err = c.Markets().GetAllMarketPrices(types.GetAllMarkPricesReq{Symbol: ""})
	require.NoError(t, err)
	require.NotNil(t, res)

	_, err = c.Markets().GetAllMarketPrices(types.GetAllMarkPricesReq{Symbol: "ABC"})
	require.Error(t, err)
}

func TestMarkets_GetOpenInterest(t *testing.T) {
	c := NewClient(key, secret)
	res, err := c.Markets().GetOpenInterest(types.GetOpenInterestReq{Symbol: "ETH_USDC_PERP"})
	require.NoError(t, err)
	require.NotNil(t, res)

	res, err = c.Markets().GetOpenInterest(types.GetOpenInterestReq{Symbol: ""})
	require.NoError(t, err)
	require.NotNil(t, res)

	_, err = c.Markets().GetOpenInterest(types.GetOpenInterestReq{Symbol: "ABC"})
	require.Error(t, err)
}

func TestMarkets_GetFundingIntervalRates(t *testing.T) {
	c := NewClient(key, secret)
	res, h, err := c.Markets().GetFundingIntervalRates(types.GetFundingIntervalRatesReq{Symbol: "ETH_USDC_PERP"})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, h)

	res, h, err = c.Markets().GetFundingIntervalRates(types.GetFundingIntervalRatesReq{Symbol: "ETH_USDC_PERP", Limit: 100})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, h)

	res, h, err = c.Markets().GetFundingIntervalRates(types.GetFundingIntervalRatesReq{Symbol: "ETH_USDC_PERP", Limit: 100, Offset: 1})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, h)

	res, h, err = c.Markets().GetFundingIntervalRates(types.GetFundingIntervalRatesReq{Symbol: "ABC"})
	require.Error(t, err)

	res, h, err = c.Markets().GetFundingIntervalRates(types.GetFundingIntervalRatesReq{Symbol: ""})
	require.Error(t, err)

	res, h, err = c.Markets().GetFundingIntervalRates(types.GetFundingIntervalRatesReq{Symbol: "ETH_USDC_PERP", Limit: 99})
	require.Error(t, err)

	res, h, err = c.Markets().GetFundingIntervalRates(types.GetFundingIntervalRatesReq{Symbol: "ETH_USDC_PERP", Limit: 1001})
	require.Error(t, err)
}

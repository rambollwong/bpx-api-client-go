package types

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

const (
	InstructionBorrowHistoryQueryAll         = "borrowHistoryQueryAll"
	InstructionInterestHistoryQueryAll       = "interestHistoryQueryAll"
	InstructionBorrowPositionHistoryQueryAll = "borrowPositionHistoryQueryAll"
	InstructionDustHistoryQueryAll           = "dustHistoryQueryAll"
	InstructionFillHistoryQueryAll           = "fillHistoryQueryAll"
	InstructionFundingHistoryQueryAll        = "fundingHistoryQueryAll"
	InstructionOrderHistoryQueryAll          = "orderHistoryQueryAll"
	InstructionRfqHistoryQueryAll            = "rfqHistoryQueryAll"
	InstructionQuoteHistoryQueryAll          = "quoteHistoryQueryAll"
	InstructionSettlementHistoryQueryAll     = "settlementHistoryQueryAll"
	InstructionStrategyHistoryQueryAll       = "strategyHistoryQueryAll"

	BorrowLendEventTypeBorrow      = "Borrow"
	BorrowLendEventTypeLend        = "Lend"
	BorrowLendEventTypeBorrowRepay = "BorrowRepay"
	BorrowLendEventTypeLendRedeem  = "LendRedeem"
)

type GetBorrowHistoryReq struct {
	Type          *string
	Sources       *string
	PositionId    *string
	Symbol        *string
	Limit         *uint64
	Offset        *uint64
	SortDirection *string
}

func (req GetBorrowHistoryReq) Validate() error {
	return nil
}

func (req GetBorrowHistoryReq) BuildQueryParams() string {
	params := make([]string, 0)

	if req.Limit != nil {
		params = append(params, fmt.Sprintf("limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *req.Offset))
	}
	if req.PositionId != nil {
		params = append(params, fmt.Sprintf("positionId=%s", *req.PositionId))
	}
	if req.SortDirection != nil {
		params = append(params, fmt.Sprintf("sortDirection=%s", *req.SortDirection))
	}
	if req.Sources != nil {
		params = append(params, fmt.Sprintf("sources=%s", *req.Sources))
	}
	if req.Symbol != nil {
		params = append(params, fmt.Sprintf("symbol=%s", *req.Symbol))
	}
	if req.Type != nil {
		params = append(params, fmt.Sprintf("type=%s", *req.Type))
	}

	if len(params) == 0 {
		return ""
	}

	// Sort params in dictionary order
	sort.Strings(params)

	return strings.Join(params, "&")
}

func (req GetBorrowHistoryReq) Instruction() string {
	return InstructionBorrowHistoryQueryAll
}

type BorrowHistory struct {
	EventType         string  `json:"eventType,omitempty"`
	PositionId        string  `json:"positionId,omitempty"`
	PositionQuantity  *string `json:"positionQuantity,omitempty"`
	Quantity          string  `json:"quantity,omitempty"`
	Source            string  `json:"source,omitempty"`
	Symbol            string  `json:"symbol,omitempty"`
	Timestamp         string  `json:"timestamp,omitempty"`
	SpotMarginOrderId *string `json:"spotMarginOrderId,omitempty"`
}

type GetBorrowHistoryResp []BorrowHistory

func (GetBorrowHistoryResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

type GetInterestHistoryReq struct {
	Asset         *string
	Limit         *uint64
	Offset        *uint64
	PositionId    *string
	SortDirection *string
	Sources       *string
	Symbol        *string
}

func (req GetInterestHistoryReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetInterestHistoryReq) BuildQueryParams() string {
	params := make([]string, 0)

	if req.Asset != nil {
		params = append(params, fmt.Sprintf("asset=%s", *req.Asset))
	}
	if req.Limit != nil {
		params = append(params, fmt.Sprintf("limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *req.Offset))
	}
	if req.PositionId != nil {
		params = append(params, fmt.Sprintf("positionId=%s", *req.PositionId))
	}
	if req.SortDirection != nil {
		params = append(params, fmt.Sprintf("sortDirection=%s", *req.SortDirection))
	}
	if req.Sources != nil {
		params = append(params, fmt.Sprintf("sources=%s", *req.Sources))
	}
	if req.Symbol != nil {
		params = append(params, fmt.Sprintf("symbol=%s", *req.Symbol))
	}

	if len(params) == 0 {
		return ""
	}

	// Sort params in dictionary order
	sort.Strings(params)

	return strings.Join(params, "&")
}

func (req GetInterestHistoryReq) Instruction() string {
	return InstructionInterestHistoryQueryAll
}

type InterestHistory struct {
	PaymentType  string `json:"paymentType,omitempty"`
	InterestRate string `json:"interestRate,omitempty"`
	Interval     uint64 `json:"interval,omitempty"`
	MarketSymbol string `json:"marketSymbol,omitempty"`
	PositionId   string `json:"positionId,omitempty"`
	Quantity     string `json:"quantity,omitempty"`
	Symbol       string `json:"symbol,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
}

type GetInterestHistoryResp []InterestHistory

func (GetInterestHistoryResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

type GetBorrowPositionHistoryReq struct {
	Symbol        *string
	Side          *string
	State         *string
	Limit         *uint64
	Offset        *uint64
	SortDirection *string
}

func (req GetBorrowPositionHistoryReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetBorrowPositionHistoryReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.Limit != nil {
		params = append(params, fmt.Sprintf("limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *req.Offset))
	}
	if req.SortDirection != nil {
		params = append(params, fmt.Sprintf("sortDirection=%s", *req.SortDirection))
	}
	if req.Side != nil {
		params = append(params, fmt.Sprintf("side=%s", *req.Side))
	}
	if req.State != nil {
		params = append(params, fmt.Sprintf("state=%s", *req.State))
	}
	if req.Symbol != nil {
		params = append(params, fmt.Sprintf("symbol=%s", *req.Symbol))
	}
	if len(params) == 0 {
		return ""
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}

func (req GetBorrowPositionHistoryReq) Instruction() string {
	return InstructionBorrowPositionHistoryQueryAll
}

type BorrowPositionHistory struct {
	PositionId         string `json:"positionId,omitempty"`
	Quantity           string `json:"quantity,omitempty"`
	Symbol             string `json:"symbol,omitempty"`
	Source             string `json:"source,omitempty"`
	CumulativeInterest string `json:"cumulativeInterest,omitempty"`
	AvgInterestRate    string `json:"avgInterestRate,omitempty"`
	Side               string `json:"side,omitempty"`
	CreatedAt          string `json:"createdAt,omitempty"`
}

type GetBorrowPositionHistoryResp []BorrowPositionHistory

func (GetBorrowPositionHistoryResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

type GetDustConversionHistoryReq struct {
	Id            *int64
	Symbol        *string
	Limit         *uint64
	Offset        *uint64
	SortDirection *string
}

func (req GetDustConversionHistoryReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetDustConversionHistoryReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.Id != nil {
		params = append(params, fmt.Sprintf("id=%d", *req.Id))
	}
	if req.Symbol != nil {
		params = append(params, fmt.Sprintf("symbol=%s", *req.Symbol))
	}
	if req.Limit != nil {
		params = append(params, fmt.Sprintf("limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *req.Offset))
	}
	if req.SortDirection != nil {
		params = append(params, fmt.Sprintf("sortDirection=%s", *req.SortDirection))
	}
	if len(params) == 0 {
		return ""
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}

func (req GetDustConversionHistoryReq) Instruction() string {
	return InstructionDustHistoryQueryAll
}

type DustConversionHistory struct {
	Id           string `json:"id,omitempty"`
	Quantity     string `json:"quantity,omitempty"`
	Symbol       string `json:"symbol,omitempty"`
	UsdcReceived string `json:"usdcReceived,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
}

type GetDustConversionHistoryResp []DustConversionHistory

func (GetDustConversionHistoryResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

type GetFillHistoryReq struct {
	OrderId       *string
	StrategyId    *string
	From          *int64
	To            *int64
	Symbol        *string
	Limit         *uint64
	Offset        *uint64
	FillType      *string
	MarketType    *string
	SortDirection *string
}

func (req GetFillHistoryReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetFillHistoryReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.OrderId != nil {
		params = append(params, fmt.Sprintf("orderId=%s", *req.OrderId))
	}
	if req.StrategyId != nil {
		params = append(params, fmt.Sprintf("strategyId=%s", *req.StrategyId))
	}
	if req.From != nil {
		params = append(params, fmt.Sprintf("from=%d", *req.From))
	}
	if req.To != nil {
		params = append(params, fmt.Sprintf("to=%d", *req.To))
	}
	if req.Symbol != nil {
		params = append(params, fmt.Sprintf("symbol=%s", *req.Symbol))
	}
	if req.Limit != nil {
		params = append(params, fmt.Sprintf("limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *req.Offset))
	}
	if req.FillType != nil {
		params = append(params, fmt.Sprintf("fillType=%s", *req.FillType))
	}
	if req.MarketType != nil {
		params = append(params, fmt.Sprintf("marketType=%s", *req.MarketType))
	}
	if req.SortDirection != nil {
		params = append(params, fmt.Sprintf("sortDirection=%s", *req.SortDirection))
	}
	if len(params) == 0 {
		return ""
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}

func (req GetFillHistoryReq) Instruction() string {
	return InstructionFillHistoryQueryAll
}

type FillHistory struct {
	ClientId        *string `json:"clientId,omitempty"`
	Fee             string  `json:"fee,omitempty"`
	FeeSymbol       string  `json:"feeSymbol,omitempty"`
	IsMaker         bool    `json:"isMaker,omitempty"`
	OrderId         string  `json:"orderId,omitempty"`
	Price           string  `json:"price,omitempty"`
	Quantity        string  `json:"quantity,omitempty"`
	Side            string  `json:"side,omitempty"`
	Symbol          string  `json:"symbol,omitempty"`
	SystemOrderType *string `json:"systemOrderType,omitempty"`
	Timestamp       string  `json:"timestamp,omitempty"`
	TradeId         *int64  `json:"tradeId,omitempty"`
}

type GetFillHistoryResp []FillHistory

func (GetFillHistoryResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

type GetFundingPaymentsReq struct {
	SubaccountId  *uint16
	Symbol        *string
	Limit         *uint64
	Offset        *uint64
	SortDirection *string
}

func (req GetFundingPaymentsReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetFundingPaymentsReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.SubaccountId != nil {
		params = append(params, fmt.Sprintf("subaccountId=%d", *req.SubaccountId))
	}
	if req.Symbol != nil {
		params = append(params, fmt.Sprintf("symbol=%s", *req.Symbol))
	}
	if req.Limit != nil {
		params = append(params, fmt.Sprintf("limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *req.Offset))
	}
	if req.SortDirection != nil {
		params = append(params, fmt.Sprintf("sortDirection=%s", *req.SortDirection))
	}
	if len(params) == 0 {
		return ""
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}

func (req GetFundingPaymentsReq) Instruction() string {
	return InstructionFundingHistoryQueryAll
}

type FundingPayment struct {
	UserId               int32   `json:"userId,omitempty"`
	SubaccountId         *uint16 `json:"subaccountId,omitempty"`
	Symbol               string  `json:"symbol,omitempty"`
	Quantity             string  `json:"quantity,omitempty"`
	IntervalEndTimestamp string  `json:"intervalEndTimestamp,omitempty"`
	FundingRate          string  `json:"fundingRate,omitempty"`
}

type GetFundingPaymentsResp []FundingPayment

func (GetFundingPaymentsResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

type GetOrderHistoryReq struct {
	OrderId    *string
	StrategyId *string
	Symbol     *string
	Limit      *uint64
	Offset     *uint64
	MarketType []string
}

func (req GetOrderHistoryReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetOrderHistoryReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.OrderId != nil {
		params = append(params, fmt.Sprintf("orderId=%s", *req.OrderId))
	}
	if req.StrategyId != nil {
		params = append(params, fmt.Sprintf("strategyId=%s", *req.StrategyId))
	}
	if req.Symbol != nil {
		params = append(params, fmt.Sprintf("symbol=%s", *req.Symbol))
	}
	if req.Limit != nil {
		params = append(params, fmt.Sprintf("limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *req.Offset))
	}
	if req.MarketType != nil {
		for _, s := range req.MarketType {
			params = append(params, fmt.Sprintf("marketType=%s", s))
		}
	}
	if len(params) == 0 {
		return ""
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}

func (req GetOrderHistoryReq) Instruction() string {
	return InstructionOrderHistoryQueryAll
}

type OrderHistory struct {
	Id                     string  `json:"id,omitempty"`
	CreatedAt              string  `json:"createdAt,omitempty"`
	ExecutedQuantity       *string `json:"executedQuantity,omitempty"`
	ExecutedQuoteQuantity  *string `json:"executedQuoteQuantity,omitempty"`
	ExpiryReason           *string `json:"expiryReason,omitempty"`
	OrderType              string  `json:"orderType,omitempty"`
	PostOnly               *bool   `json:"postOnly,omitempty"`
	Price                  *string `json:"price,omitempty"`
	Quantity               *string `json:"quantity,omitempty"`
	QuoteQuantity          *string `json:"quoteQuantity,omitempty"`
	SelfTradePrevention    string  `json:"selfTradePrevention,omitempty"`
	Status                 string  `json:"status,omitempty"`
	Side                   string  `json:"side,omitempty"`
	StopLossTriggerPrice   *string `json:"stopLossTriggerPrice,omitempty"`
	StopLossLimitPrice     *string `json:"stopLossLimitPrice,omitempty"`
	StopLossTriggerBy      *string `json:"stopLossTriggerBy,omitempty"`
	Symbol                 string  `json:"symbol,omitempty"`
	TakeProfitTriggerPrice *string `json:"takeProfitTriggerPrice,omitempty"`
	TakeProfitLimitPrice   *string `json:"takeProfitLimitPrice,omitempty"`
	TakeProfitTriggerBy    *string `json:"takeProfitTriggerBy,omitempty"`
	TimeInForce            string  `json:"timeInForce,omitempty"`
	TriggerBy              *string `json:"triggerBy,omitempty"`
	TriggerPrice           *string `json:"triggerPrice,omitempty"`
	TriggerQuantity        *string `json:"triggerQuantity,omitempty"`
	ClientId               *uint32 `json:"clientId,omitempty"`
	SystemOrderType        *string `json:"systemOrderType,omitempty"`
	StrategyId             *string `json:"strategyId,omitempty"`
	SlippageTolerance      *string `json:"slippageTolerance,omitempty"`
	SlippageToleranceType  *string `json:"slippageToleranceType,omitempty"`
}

type GetOrderHistoryResp []OrderHistory

func (GetOrderHistoryResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

type GetRfqHistoryReq struct {
	RfqId         *string
	Symbol        *string
	Status        *string
	Side          *string
	Limit         *uint64
	Offset        *uint64
	SortDirection *string
}

func (req GetRfqHistoryReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetRfqHistoryReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.RfqId != nil {
		params = append(params, fmt.Sprintf("rfqId=%s", *req.RfqId))
	}
	if req.Symbol != nil {
		params = append(params, fmt.Sprintf("symbol=%s", *req.Symbol))
	}
	if req.Status != nil {
		params = append(params, fmt.Sprintf("status=%s", *req.Status))
	}
	if req.Side != nil {
		params = append(params, fmt.Sprintf("side=%s", *req.Side))
	}
	if req.Limit != nil {
		params = append(params, fmt.Sprintf("limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *req.Offset))
	}
	if req.SortDirection != nil {
		params = append(params, fmt.Sprintf("sortDirection=%s", *req.SortDirection))
	}
	if len(params) == 0 {
		return ""
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}

func (req GetRfqHistoryReq) Instruction() string {
	return InstructionRfqHistoryQueryAll
}

type RfqHistory struct {
	UserId         int32   `json:"userId,omitempty"`
	SubaccountId   *int32  `json:"subaccountId,omitempty"`
	RfqId          string  `json:"rfqId,omitempty"`
	ClientId       *uint32 `json:"clientId,omitempty"`
	Symbol         string  `json:"symbol,omitempty"`
	Side           string  `json:"side,omitempty"`
	Price          *string `json:"price,omitempty"`
	Quantity       *string `json:"quantity,omitempty"`
	QuoteQuantity  *string `json:"quoteQuantity,omitempty"`
	SubmissionType string  `json:"submissionType,omitempty"`
	ExpiryTime     string  `json:"expiryTime,omitempty"`
}

type GetRfqHistoryResp []RfqHistory

func (GetRfqHistoryResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

type GetQuoteHistoryReq struct {
	QuoteId       *string
	Symbol        *string
	Status        *string
	Limit         *uint64
	Offset        *uint64
	SortDirection *string
}

func (req GetQuoteHistoryReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetQuoteHistoryReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.QuoteId != nil {
		params = append(params, fmt.Sprintf("quoteId=%s", *req.QuoteId))
	}
	if req.Symbol != nil {
		params = append(params, fmt.Sprintf("symbol=%s", *req.Symbol))
	}
	if req.Status != nil {
		params = append(params, fmt.Sprintf("status=%s", *req.Status))
	}
	if req.Limit != nil {
		params = append(params, fmt.Sprintf("limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *req.Offset))
	}
	if req.SortDirection != nil {
		params = append(params, fmt.Sprintf("sortDirection=%s", *req.SortDirection))
	}
	if len(params) == 0 {
		return ""
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}

func (req GetQuoteHistoryReq) Instruction() string {
	return InstructionQuoteHistoryQueryAll
}

type QuoteHistory struct {
	UserId       int32   `json:"userId,omitempty"`
	SubaccountId *int32  `json:"subaccountId,omitempty"`
	RfqId        string  `json:"rfqId,omitempty"`
	QuoteId      string  `json:"quoteId,omitempty"`
	ClientId     *uint32 `json:"clientId,omitempty"`
	BidPrice     string  `json:"bidPrice,omitempty"`
	AskPrice     string  `json:"askPrice,omitempty"`
	Status       string  `json:"status,omitempty"`
	CreatedAt    string  `json:"createdAt,omitempty"`
}

type GetQuoteHistoryResp []QuoteHistory

func (GetQuoteHistoryResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

type GetSettlementHistoryReq struct {
	Limit         *uint64
	Offset        *uint64
	Source        *string
	SortDirection *string
}

func (req GetSettlementHistoryReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetSettlementHistoryReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.Limit != nil {
		params = append(params, fmt.Sprintf("limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *req.Offset))
	}
	if req.Source != nil {
		params = append(params, fmt.Sprintf("source=%s", *req.Source))
	}
	if req.SortDirection != nil {
		params = append(params, fmt.Sprintf("sortDirection=%s", *req.SortDirection))
	}
	if len(params) == 0 {
		return ""
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}

func (req GetSettlementHistoryReq) Instruction() string {
	return InstructionSettlementHistoryQueryAll
}

type SettlementHistory struct {
	Quantity     string `json:"quantity,omitempty"`
	Source       string `json:"source,omitempty"`
	SubaccountId *int32 `json:"subaccountId,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
	UserId       int32  `json:"userId,omitempty"`
}

type GetSettlementHistoryResp []SettlementHistory

func (GetSettlementHistoryResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

type GetStrategyHistoryReq struct {
	StrategyId    *string
	Symbol        *string
	Limit         *uint64
	Offset        *uint64
	MarketType    []string
	SortDirection *string
}

func (req GetStrategyHistoryReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetStrategyHistoryReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.StrategyId != nil {
		params = append(params, fmt.Sprintf("strategyId=%s", *req.StrategyId))
	}
	if req.Symbol != nil {
		params = append(params, fmt.Sprintf("symbol=%s", *req.Symbol))
	}
	if req.Limit != nil {
		params = append(params, fmt.Sprintf("limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *req.Offset))
	}
	if len(req.MarketType) > 0 {
		for _, s := range req.MarketType {
			params = append(params, fmt.Sprintf("marketType=%s", s))
		}
	}
	if req.SortDirection != nil {
		params = append(params, fmt.Sprintf("sortDirection=%s", *req.SortDirection))
	}
	if len(params) == 0 {
		return ""
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}

func (req GetStrategyHistoryReq) Instruction() string {
	return InstructionStrategyHistoryQueryAll
}

type StrategyHistory struct {
	Id                         string  `json:"id,omitempty"`
	CreatedAt                  string  `json:"createdAt,omitempty"`
	ExecutedQuantity           *string `json:"executedQuantity,omitempty"`
	ExecutedQuoteQuantity      *string `json:"executedQuoteQuantity,omitempty"`
	CancelReason               *string `json:"cancelReason,omitempty"`
	StrategyType               string  `json:"strategyType,omitempty"`
	Quantity                   *string `json:"quantity,omitempty"`
	SelfTradePrevention        string  `json:"selfTradePrevention,omitempty"`
	Status                     string  `json:"status,omitempty"`
	Side                       string  `json:"side,omitempty"`
	Symbol                     string  `json:"symbol,omitempty"`
	TimeInForce                string  `json:"timeInForce,omitempty"`
	ClientStrategyId           *uint32 `json:"clientStrategyId,omitempty"`
	Duration                   uint64  `json:"duration,omitempty"`
	Interval                   uint64  `json:"interval,omitempty"`
	RandomizedIntervalQuantity bool    `json:"randomizedIntervalQuantity,omitempty"`
	SlippageTolerance          *string `json:"slippageTolerance,omitempty"`
	SlippageToleranceType      *string `json:"slippageToleranceType,omitempty"`
}

type GetStrategyHistoryResp []StrategyHistory

func (GetStrategyHistoryResp) ReadResponseHeaders(header http.Header) (rh *ResponseHeaders, err error) {
	return ReadResponseHeaders(header)
}

package types

import (
	"fmt"
	"sort"
	"strings"
)

const (
	InstructionOrderQuery     = "orderQuery"
	InstructionOrderExecute   = "orderExecute"
	InstructionOrderCancel    = "orderCancel"
	InstructionOrderQueryAll  = "orderQueryAll"
	InstructionOrderCancelAll = "orderCancelAll"
)

type GetOpenOrderReq struct {
	ClientId *uint32
	OrderId  *string
	Symbol   string
}

func (req GetOpenOrderReq) Validate() error {
	return nil
}

func (req GetOpenOrderReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.ClientId != nil {
		params = append(params, fmt.Sprintf("clientId=%d", *req.ClientId))
	}
	if req.OrderId != nil {
		params = append(params, fmt.Sprintf("orderId=%s", *req.OrderId))
	}
	params = append(params, fmt.Sprintf("symbol=%s", req.Symbol))
	sort.Strings(params)
	return strings.Join(params, "&")
}

func (req GetOpenOrderReq) Instruction() string {
	return InstructionOrderQuery
}

type Order struct {
	OrderType              string  `json:"orderType,omitempty"`
	Id                     string  `json:"id,omitempty"`
	ClientId               *uint32 `json:"clientId,omitempty"`
	CreatedAt              int64   `json:"createdAt,omitempty"`
	ExecutedQuantity       string  `json:"executedQuantity,omitempty"`
	ExecutedQuoteQuantity  string  `json:"executedQuoteQuantity,omitempty"`
	Price                  *string `json:"price,omitempty"`
	Quantity               *string `json:"quantity,omitempty"`
	QuoteQuantity          *string `json:"quoteQuantity,omitempty"`
	ReduceOnly             *bool   `json:"reduceOnly,omitempty"`
	TimeInForce            string  `json:"timeInForce,omitempty"`
	SelfTradePrevention    string  `json:"selfTradePrevention,omitempty"`
	Side                   string  `json:"side,omitempty"`
	Status                 string  `json:"status,omitempty"`
	StopLossTriggerPrice   *string `json:"stopLossTriggerPrice,omitempty"`
	StopLossLimitPrice     *string `json:"stopLossLimitPrice,omitempty"`
	StopLossTriggerBy      *string `json:"stopLossTriggerBy,omitempty"`
	Symbol                 string  `json:"symbol,omitempty"`
	TakeProfitTriggerPrice *string `json:"takeProfitTriggerPrice,omitempty"`
	TakeProfitLimitPrice   *string `json:"takeProfitLimitPrice,omitempty"`
	TakeProfitTriggerBy    *string `json:"takeProfitTriggerBy,omitempty"`
	TriggerBy              *string `json:"triggerBy,omitempty"`
	TriggerPrice           *string `json:"triggerPrice,omitempty"`
	TriggerQuantity        *string `json:"triggerQuantity,omitempty"`
	TriggeredAt            *int64  `json:"triggeredAt,omitempty"`
	RelatedOrderId         *string `json:"relatedOrderId,omitempty"`
	StrategyId             *string `json:"strategyId,omitempty"`
	SlippageTolerance      *string `json:"slippageTolerance,omitempty"`
	SlippageToleranceType  *string `json:"slippageToleranceType,omitempty"`
}

type ExecuteOrderReq struct {
	AutoLend               *bool   `json:"autoLend,omitempty"`
	AutoLendRedeem         *bool   `json:"autoLendRedeem,omitempty"`
	AutoBorrow             *bool   `json:"autoBorrow,omitempty"`
	AutoBorrowRepay        *bool   `json:"autoBorrowRepay,omitempty"`
	BrokerId               *uint16 `json:"brokerId,omitempty"`
	ClientId               *uint32 `json:"clientId,omitempty"`
	OrderType              string  `json:"orderType"`
	PostOnly               *bool   `json:"postOnly,omitempty"`
	Price                  *string `json:"price,omitempty"`
	Quantity               *string `json:"quantity,omitempty"`
	QuoteQuantity          *string `json:"quoteQuantity,omitempty"`
	ReduceOnly             *bool   `json:"reduceOnly,omitempty"`
	SelfTradePrevention    *string `json:"selfTradePrevention,omitempty"`
	Side                   string  `json:"side"`
	StopLossLimitPrice     *string `json:"stopLossLimitPrice,omitempty"`
	StopLossTriggerBy      *string `json:"stopLossTriggerBy,omitempty"`
	StopLossTriggerPrice   *string `json:"stopLossTriggerPrice,omitempty"`
	Symbol                 string  `json:"symbol"`
	TakeProfitLimitPrice   *string `json:"takeProfitLimitPrice,omitempty"`
	TakeProfitTriggerBy    *string `json:"takeProfitTriggerBy,omitempty"`
	TakeProfitTriggerPrice *string `json:"takeProfitTriggerPrice,omitempty"`
	TimeInForce            *string `json:"timeInForce,omitempty"`
	TriggerBy              *string `json:"triggerBy,omitempty"`
	TriggerPrice           *string `json:"triggerPrice,omitempty"`
	TriggerQuantity        *string `json:"triggerQuantity,omitempty"`
	SlippageTolerance      *string `json:"slippageTolerance,omitempty"`
	SlippageToleranceType  *string `json:"slippageToleranceType,omitempty"`
}

func (req ExecuteOrderReq) Validate() error {
	return nil
}

func (req ExecuteOrderReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.AutoLend != nil {
		params = append(params, fmt.Sprintf("autoLend=%t", *req.AutoLend))
	}
	if req.AutoLendRedeem != nil {
		params = append(params, fmt.Sprintf("autoLendRedeem=%t", *req.AutoLendRedeem))
	}
	if req.AutoBorrow != nil {
		params = append(params, fmt.Sprintf("autoBorrow=%t", *req.AutoBorrow))
	}
	if req.AutoBorrowRepay != nil {
		params = append(params, fmt.Sprintf("autoBorrowRepay=%t", *req.AutoBorrowRepay))
	}
	if req.BrokerId != nil {
		params = append(params, fmt.Sprintf("brokerId=%d", *req.BrokerId))
	}
	if req.ClientId != nil {
		params = append(params, fmt.Sprintf("clientId=%d", *req.ClientId))
	}
	params = append(params, fmt.Sprintf("orderType=%s", req.OrderType))
	if req.PostOnly != nil {
		params = append(params, fmt.Sprintf("postOnly=%t", *req.PostOnly))
	}
	if req.Price != nil {
		params = append(params, fmt.Sprintf("price=%s", *req.Price))
	}
	if req.Quantity != nil {
		params = append(params, fmt.Sprintf("quantity=%s", *req.Quantity))
	}
	if req.QuoteQuantity != nil {
		params = append(params, fmt.Sprintf("quoteQuantity=%s", *req.QuoteQuantity))
	}
	if req.ReduceOnly != nil {
		params = append(params, fmt.Sprintf("reduceOnly=%t", *req.ReduceOnly))
	}
	if req.SelfTradePrevention != nil {
		params = append(params, fmt.Sprintf("selfTradePrevention=%s", *req.SelfTradePrevention))
	}
	params = append(params, fmt.Sprintf("side=%s", req.Side))
	if req.StopLossLimitPrice != nil {
		params = append(params, fmt.Sprintf("stopLossLimitPrice=%s", *req.StopLossLimitPrice))
	}
	if req.StopLossTriggerBy != nil {
		params = append(params, fmt.Sprintf("stopLossTriggerBy=%s", *req.StopLossTriggerBy))
	}
	if req.StopLossTriggerPrice != nil {
		params = append(params, fmt.Sprintf("stopLossTriggerPrice=%s", *req.StopLossTriggerPrice))
	}
	params = append(params, fmt.Sprintf("symbol=%s", req.Symbol))
	if req.TakeProfitLimitPrice != nil {
		params = append(params, fmt.Sprintf("takeProfitLimitPrice=%s", *req.TakeProfitLimitPrice))
	}
	if req.TakeProfitTriggerBy != nil {
		params = append(params, fmt.Sprintf("takeProfitTriggerBy=%s", *req.TakeProfitTriggerBy))
	}
	if req.TakeProfitTriggerPrice != nil {
		params = append(params, fmt.Sprintf("takeProfitTriggerPrice=%s", *req.TakeProfitTriggerPrice))
	}
	if req.TimeInForce != nil {
		params = append(params, fmt.Sprintf("timeInForce=%s", *req.TimeInForce))
	}
	if req.TriggerBy != nil {
		params = append(params, fmt.Sprintf("triggerBy=%s", *req.TriggerBy))
	}
	if req.TriggerPrice != nil {
		params = append(params, fmt.Sprintf("triggerPrice=%s", *req.TriggerPrice))
	}
	if req.TriggerQuantity != nil {
		params = append(params, fmt.Sprintf("triggerQuantity=%s", *req.TriggerQuantity))
	}
	if req.SlippageTolerance != nil {
		params = append(params, fmt.Sprintf("slippageTolerance=%s", *req.SlippageTolerance))
	}
	if req.SlippageToleranceType != nil {
		params = append(params, fmt.Sprintf("slippageToleranceType=%s", *req.SlippageToleranceType))
	}
	return strings.Join(params, "&")
}

func (req ExecuteOrderReq) Instruction() string {
	return InstructionOrderExecute
}

type CancelOpenOrderReq struct {
	ClientId *uint32 `json:"clientId,omitempty"`
	OrderId  *string `json:"orderId,omitempty"`
	Symbol   string  `json:"symbol"`
}

func (req CancelOpenOrderReq) Validate() error {
	return nil
}

func (req CancelOpenOrderReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.ClientId != nil {
		params = append(params, fmt.Sprintf("clientId=%d", *req.ClientId))
	}
	if req.OrderId != nil {
		params = append(params, fmt.Sprintf("orderId=%s", *req.OrderId))
	}
	params = append(params, fmt.Sprintf("symbol=%s", req.Symbol))
	return strings.Join(params, "&")
}

func (req CancelOpenOrderReq) Instruction() string {
	return InstructionOrderCancel
}

type ExecuteOrdersReq []ExecuteOrderReq

func (req ExecuteOrdersReq) Validate() error {
	for _, r := range req {
		if err := r.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (req ExecuteOrdersReq) BuildQueryParams() string {
	params := make([]string, 0)
	for _, r := range req {
		params = append(params, r.BuildQueryParams())
	}
	return strings.Join(params, "&instruction="+InstructionOrderExecute+"&")
}

func (req ExecuteOrdersReq) Instruction() string {
	return InstructionOrderExecute
}

type GetOpenOrdersReq struct {
	MarketType *string `json:"marketType,omitempty"`
	Symbol     *string `json:"symbol,omitempty"`
}

func (req GetOpenOrdersReq) Validate() error {
	return nil
}
func (req GetOpenOrdersReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.MarketType != nil {
		params = append(params, fmt.Sprintf("marketType=%s", *req.MarketType))
	}
	if req.Symbol != nil {
		params = append(params, fmt.Sprintf("symbol=%s", *req.Symbol))
	}
	if len(params) == 0 {
		return ""
	}
	return strings.Join(params, "&")
}

func (req GetOpenOrdersReq) Instruction() string {
	return InstructionOrderQueryAll
}

type CancelOpenOrdersReq struct {
	OrderType *string `json:"orderType,omitempty"`
	Symbol    string  `json:"symbol,omitempty"`
}

func (req CancelOpenOrdersReq) Validate() error {
	return nil
}

func (req CancelOpenOrdersReq) BuildQueryParams() string {
	params := make([]string, 0)
	if req.OrderType != nil {
		params = append(params, fmt.Sprintf("orderType=%s", *req.OrderType))
	}
	params = append(params, fmt.Sprintf("symbol=%s", req.Symbol))
	return strings.Join(params, "&")
}

func (req CancelOpenOrdersReq) Instruction() string {
	return InstructionOrderCancelAll
}

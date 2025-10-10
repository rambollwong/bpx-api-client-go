package types

import (
	"errors"
	"fmt"
	"strings"
)

const (
	InstructionAccountQuery          = "accountQuery"
	InstructionAccountUpdate         = "accountUpdate"
	InstructionConvertDust           = "convertDust"
	InstructionMaxBorrowQuantity     = "maxBorrowQuantity"
	InstructionMaxOrderQuantity      = "maxOrderQuantity"
	InstructionMaxWithdrawalQuantity = "maxWithdrawalQuantity"
)

type GetAccountReq struct {
}

func (req GetAccountReq) Validate() error {
	return nil
}

func (req GetAccountReq) BuildQueryParams() string {
	return ""
}

func (req GetAccountReq) Instruction() string {
	return InstructionAccountQuery
}

type UpdateAccountReq struct {
	AutoBorrowSettlements bool   `json:"autoBorrowSettlements"`
	AutoLend              bool   `json:"autoLend"`
	AutoRepayBorrows      bool   `json:"autoRepayBorrows"`
	LeverageLimit         string `json:"leverageLimit"`
}

func (req UpdateAccountReq) Validate() error {
	return nil
}

func (req UpdateAccountReq) BuildQueryParams() string {
	return fmt.Sprintf("autoBorrowSettlements=%t&autoLend=%t&autoRepayBorrows=%t&leverageLimit=%s",
		req.AutoBorrowSettlements, req.AutoLend, req.AutoRepayBorrows, req.LeverageLimit)
}

func (req UpdateAccountReq) Instruction() string {
	return InstructionAccountUpdate
}

type Account struct {
	AutoBorrowSettlements bool   `json:"autoBorrowSettlements,omitempty"`
	AutoLend              bool   `json:"autoLend,omitempty"`
	AutoRealizePnl        bool   `json:"autoRealizePnl,omitempty"`
	AutoRepayBorrows      bool   `json:"autoRepayBorrows,omitempty"`
	BorrowLimit           string `json:"borrowLimit,omitempty"`
	FuturesMakerFee       string `json:"futuresMakerFee,omitempty"`
	FuturesTakerFee       string `json:"futuresTakerFee,omitempty"`
	LeverageLimit         string `json:"leverageLimit,omitempty"`
	LimitOrders           uint64 `json:"limitOrders,omitempty"`
	Liquidating           bool   `json:"liquidating,omitempty"`
	PositionLimit         string `json:"positionLimit,omitempty"`
	SpotMakerFee          string `json:"spotMakerFee,omitempty"`
	SpotTakerFee          string `json:"spotTakerFee,omitempty"`
	TriggerOrders         uint64 `json:"triggerOrders,omitempty"`
}

type ConvertDustReq struct {
	Symbol string `json:"symbol"`
}

func (req ConvertDustReq) Validate() error {
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	return nil
}

func (req ConvertDustReq) BuildQueryParams() string {
	return "symbol=" + req.Symbol
}

func (req ConvertDustReq) Instruction() string {
	return InstructionConvertDust
}

type MaxBorrowQuantityReq struct {
	Symbol string `json:"symbol"`
}

func (req MaxBorrowQuantityReq) Validate() error {
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	return nil
}

func (req MaxBorrowQuantityReq) BuildQueryParams() string {
	return "symbol=" + req.Symbol
}

func (req MaxBorrowQuantityReq) Instruction() string {
	return InstructionMaxBorrowQuantity
}

type MaxBorrowQuantityResp struct {
	MaxBorrowQuantity string `json:"maxBorrowQuantity,omitempty"`
	Symbol            string `json:"symbol,omitempty"`
}

type MaxOrderQuantityReq struct {
	Symbol          string `json:"symbol,omitempty"`
	Side            string `json:"side,omitempty"`
	Price           string `json:"price,omitempty"`
	ReduceOnly      bool   `json:"reduceOnly,omitempty"`
	AutoBorrow      bool   `json:"autoBorrow,omitempty"`
	AutoBorrowRepay bool   `json:"autoBorrowRepay,omitempty"`
	AutoLendRedeem  bool   `json:"autoLendRedeem,omitempty"`
}

func (req MaxOrderQuantityReq) Validate() error {
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	if req.Side == "" {
		return errors.New("side is required")
	}
	return nil
}

func (req MaxOrderQuantityReq) BuildQueryParams() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(fmt.Sprintf("autoBorrow=%t&autoBorrowRepay=%t&autoLendRedeem=%t",
		req.AutoBorrow, req.AutoBorrowRepay, req.AutoLendRedeem))
	if req.Price != "" {
		strBuilder.WriteString("&price=")
		strBuilder.WriteString(req.Price)
	}
	strBuilder.WriteString(fmt.Sprintf("&reduceOnly=%t&side=%s&symbol=%s", req.ReduceOnly, req.Side, req.Symbol))
	return strBuilder.String()
}

func (req MaxOrderQuantityReq) Instruction() string {
	return InstructionMaxOrderQuantity
}

type MaxOrderQuantityResp struct {
	AutoBorrow       bool   `json:"autoBorrow,omitempty"`
	AutoBorrowRepay  bool   `json:"autoBorrowRepay,omitempty"`
	AutoLendRedeem   bool   `json:"autoLendRedeem,omitempty"`
	MaxOrderQuantity string `json:"maxOrderQuantity,omitempty"`
	Price            string `json:"price,omitempty"`
	Side             string `json:"side,omitempty"`
	Symbol           string `json:"symbol,omitempty"`
	ReduceOnly       bool   `json:"reduceOnly,omitempty"`
}

type MaxWithdrawalQuantityReq struct {
	AutoBorrow     *bool  `json:"autoBorrow,omitempty"`
	AutoLendRedeem *bool  `json:"autoLendRedeem,omitempty"`
	Symbol         string `json:"symbol,omitempty"`
}

func (req MaxWithdrawalQuantityReq) Validate() error {
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	return nil
}

func (req MaxWithdrawalQuantityReq) BuildQueryParams() string {
	strBuilder := strings.Builder{}
	if req.AutoBorrow != nil {
		strBuilder.WriteString(fmt.Sprintf("&autoBorrow=%t", *req.AutoBorrow))
	}
	if req.AutoLendRedeem != nil {
		strBuilder.WriteString(fmt.Sprintf("&autoLendRedeem=%t", *req.AutoLendRedeem))
	}
	strBuilder.WriteString("&symbol=")
	strBuilder.WriteString(req.Symbol)
	return strBuilder.String()[1:]
}

func (req MaxWithdrawalQuantityReq) Instruction() string {
	return InstructionMaxWithdrawalQuantity
}

type MaxWithdrawalQuantityResp struct {
	AutoBorrow            *bool  `json:"autoBorrow,omitempty"`
	AutoLendRedeem        *bool  `json:"autoLendRedeem,omitempty"`
	MaxWithdrawalQuantity string `json:"maxWithdrawalQuantity,omitempty"`
	Symbol                string `json:"symbol,omitempty"`
}

package types

import (
	"errors"
	"fmt"
	"strings"
)

const (
	InstructionBorrowLendPositionQuery = "borrowLendPositionQuery"
	InstructionBorrowLendExecute       = "borrowLendExecute"

	BorrowLendSideBorrow = "Borrow"
	BorrowLendSideLend   = "Lend"
)

type GetBorrowLendPositionReq struct {
}

func (req GetBorrowLendPositionReq) Validate() error {

	return nil
}

func (req GetBorrowLendPositionReq) BuildQueryParams() string {

	return ""
}

func (req GetBorrowLendPositionReq) Instruction() string {
	return InstructionBorrowLendPositionQuery
}

type BorrowLendPosition struct {
	CumulativeInterest  string   `json:"cumulativeInterest,omitempty"`
	Id                  string   `json:"id,omitempty"`
	Imf                 string   `json:"imf,omitempty"`
	ImfFunction         Function `json:"imfFunction,omitempty"`
	NetQuantity         string   `json:"netQuantity,omitempty"`
	MarkPrice           string   `json:"markPrice,omitempty"`
	Mmf                 string   `json:"mmf,omitempty"`
	MmfFunction         Function `json:"mmfFunction,omitempty"`
	NetExposureQuantity string   `json:"netExposureQuantity,omitempty"`
	NetExposureNotional string   `json:"netExposureNotional,omitempty"`
	Symbol              string   `json:"symbol,omitempty"`
}

type ExecuteBorrowLendReq struct {
	Quantity string `json:"quantity,omitempty"`
	Side     string `json:"side,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
}

func (req ExecuteBorrowLendReq) Validate() error {
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	if req.Side == "" {
		return errors.New("side is required")
	}
	if req.Quantity == "" {
		return errors.New("quantity is required")
	}
	return nil
}

func (req ExecuteBorrowLendReq) BuildQueryParams() string {
	return fmt.Sprintf("quantity=%s&side=%s&symbol=%s", req.Quantity, req.Side, req.Symbol)
}

func (req ExecuteBorrowLendReq) Instruction() string {
	return InstructionBorrowLendExecute
}

type GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionReq struct {
	Borrow       string
	SubaccountId *uint16
}

func (req GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionReq) Validate() error {
	if req.Borrow == "" {
		return errors.New("borrow is required")
	}
	return nil
}

func (req GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionReq) BuildQueryParams() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(fmt.Sprintf("borrow=%s", req.Borrow))
	if req.SubaccountId != nil {
		strBuilder.WriteString(fmt.Sprintf("&subaccountId=%d", *req.SubaccountId))
	}
	return strBuilder.String()
}

type GetAnEstimatedLiquidationPriceForPotentialBorrowLendPositionResp struct {
	LiquidationPrice string `json:"liquidationPrice,omitempty"`
	MarkPrice        string `json:"markPrice,omitempty"`
}

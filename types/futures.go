package types

import "fmt"

const (
	InstructionPositionQuery = "positionQuery"
)

type GetOpenPositionsReq struct {
	Symbol *string
}

func (req GetOpenPositionsReq) Validate() error {
	return nil
}

func (req GetOpenPositionsReq) BuildQueryParams() string {
	if req.Symbol != nil {
		return fmt.Sprintf("symbol=%s", *req.Symbol)
	}
	return ""
}

func (req GetOpenPositionsReq) Instruction() string {
	return InstructionPositionQuery
}

type Position struct {
	BreakEvenPrice           string   `json:"breakEvenPrice,omitempty"`
	EntryPrice               string   `json:"entryPrice,omitempty"`
	EstLiquidationPrice      string   `json:"estLiquidationPrice,omitempty"`
	Imf                      string   `json:"imf,omitempty"`
	ImfFunction              Function `json:"imfFunction,omitempty"`
	MarkPrice                string   `json:"markPrice,omitempty"`
	Mmf                      string   `json:"mmf,omitempty"`
	MmfFunction              Function `json:"mmfFunction,omitempty"`
	NetCost                  string   `json:"netCost,omitempty"`
	NetQuantity              string   `json:"netQuantity,omitempty"`
	NetExposureQuantity      string   `json:"netExposureQuantity,omitempty"`
	NetExposureNotional      string   `json:"netExposureNotional,omitempty"`
	PnlRealized              string   `json:"pnlRealized,omitempty"`
	PnlUnrealized            string   `json:"pnlUnrealized,omitempty"`
	CumulativeFundingPayment string   `json:"cumulativeFundingPayment,omitempty"`
	SubaccountId             *uint16  `json:"subaccountId,omitempty"`
	Symbol                   string   `json:"symbol,omitempty"`
	UserId                   int32    `json:"userId,omitempty"`
	PositionId               string   `json:"positionId,omitempty"`
	CumulativeInterest       string   `json:"cumulativeInterest,omitempty"`
}

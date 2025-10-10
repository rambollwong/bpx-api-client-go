package types

import (
	"errors"
	"fmt"
	"strings"
)

const (
	InstructionBalanceQuery        = "balanceQuery"
	InstructionCollateralQuery     = "collateralQuery"
	InstructionDepositQueryAll     = "depositQueryAll"
	InstructionDepositAddressQuery = "depositAddressQuery"
	InstructionWithdrawalQueryAll  = "withdrawalQueryAll"
	InstructionWithdraw            = "withdraw"
)

type GetBalancesReq struct {
}

func (req GetBalancesReq) Validate() error {
	return nil
}

func (req GetBalancesReq) BuildQueryParams() string {
	return ""
}

func (req GetBalancesReq) Instruction() string {
	return InstructionBalanceQuery
}

type GetBalancesResp map[string]Balance

type GetCollateralReq struct {
	SubaccountId *uint16
}

func (req GetCollateralReq) Validate() error {
	return nil
}

func (req GetCollateralReq) BuildQueryParams() string {
	if req.SubaccountId != nil {
		return fmt.Sprintf("subaccountId=%d", *req.SubaccountId)
	}
	return ""
}

func (req GetCollateralReq) Instruction() string {
	return InstructionCollateralQuery
}

type GetCollateralResp struct {
	AssetsValue        string       `json:"assetsValue,omitempty"`
	BorrowLiability    string       `json:"borrowLiability,omitempty"`
	Collateral         []Collateral `json:"collateral,omitempty"`
	Imf                string       `json:"imf,omitempty"`
	UnsettledEquity    string       `json:"unsettledEquity,omitempty"`
	LiabilitiesValue   string       `json:"liabilitiesValue,omitempty"`
	MarginFraction     *string      `json:"marginFraction,omitempty"`
	Mmf                string       `json:"mmf,omitempty"`
	NetEquity          string       `json:"netEquity,omitempty"`
	NetEquityAvailable string       `json:"netEquityAvailable,omitempty"`
	NetEquityLocked    string       `json:"netEquityLocked,omitempty"`
	NetExposureFutures string       `json:"netExposureFutures,omitempty"`
	PnlUnrealized      string       `json:"pnlUnrealized,omitempty"`
}

type GetDepositsReq struct {
	From   *int64
	To     *int64
	Limit  *uint64
	Offset *uint64
}

func (req GetDepositsReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetDepositsReq) BuildQueryParams() string {
	strBuilder := strings.Builder{}
	if req.From != nil {
		strBuilder.WriteString(fmt.Sprintf("&from=%d", *req.From))
	}
	if req.Limit != nil {
		strBuilder.WriteString(fmt.Sprintf("&limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		strBuilder.WriteString(fmt.Sprintf("&offset=%d", *req.Offset))
	}
	if req.To != nil {
		strBuilder.WriteString(fmt.Sprintf("&to=%d", *req.To))
	}
	if strBuilder.Len() > 0 {
		return strBuilder.String()[1:]
	}
	return ""
}

func (req GetDepositsReq) Instruction() string {
	return InstructionDepositQueryAll
}

type Deposit struct {
	Id              int32    `json:"id,omitempty"`
	ToAddress       *string  `json:"toAddress,omitempty"`
	FromAddress     *string  `json:"fromAddress,omitempty"`
	Source          string   `json:"source,omitempty"`
	Status          string   `json:"status,omitempty"`
	TransactionHash *string  `json:"transactionHash,omitempty"`
	Symbol          string   `json:"symbol,omitempty"`
	Quantity        string   `json:"quantity,omitempty"`
	CreatedAt       string   `json:"createdAt,omitempty"`
	FiatAmount      *float64 `json:"fiatAmount,omitempty"`
	FiatCurrency    *string  `json:"fiatCurrency,omitempty"`
	InstitutionBic  *string  `json:"institutionBic,omitempty"`
	PlatformMemo    *string  `json:"platformMemo,omitempty"`
}

type GetDepositAddressReq struct {
	Blockchain string
}

func (req GetDepositAddressReq) Validate() error {
	if req.Blockchain == "" {
		return errors.New("blockchain is required")
	}
	return nil
}

func (req GetDepositAddressReq) BuildQueryParams() string {
	return fmt.Sprintf("blockchain=%s", req.Blockchain)
}

func (req GetDepositAddressReq) Instruction() string {
	return InstructionDepositAddressQuery
}

type GetDepositAddressResp struct {
	Address string `json:"address,omitempty"`
}

type GetWithdrawalsReq GetDepositsReq

func (req GetWithdrawalsReq) Validate() error {
	if req.Limit != nil && *req.Limit > 1000 {
		return errors.New("limit must be less than 1000")
	}
	return nil
}

func (req GetWithdrawalsReq) BuildQueryParams() string {
	strBuilder := strings.Builder{}
	if req.From != nil {
		strBuilder.WriteString(fmt.Sprintf("&from=%d", *req.From))
	}
	if req.Limit != nil {
		strBuilder.WriteString(fmt.Sprintf("&limit=%d", *req.Limit))
	}
	if req.Offset != nil {
		strBuilder.WriteString(fmt.Sprintf("&offset=%d", *req.Offset))
	}
	if req.To != nil {
		strBuilder.WriteString(fmt.Sprintf("&to=%d", *req.To))
	}
	if strBuilder.Len() > 0 {
		return strBuilder.String()[1:]
	}
	return ""
}
func (req GetWithdrawalsReq) Instruction() string {
	return InstructionWithdrawalQueryAll
}

type Withdrawal struct {
	Id                int32   `json:"id,omitempty"`
	Blockchain        string  `json:"blockchain,omitempty"`
	ClientId          *string `json:"clientId,omitempty"`
	Identifier        *string `json:"identifier,omitempty"`
	Quantity          string  `json:"quantity,omitempty"`
	Fee               string  `json:"fee,omitempty"`
	FiatFee           *string `json:"fiatFee,omitempty"`
	FiatState         *string `json:"fiatState,omitempty"`
	FiatSymbol        *string `json:"fiatSymbol,omitempty"`
	ProviderId        *string `json:"providerId,omitempty"`
	Symbol            string  `json:"symbol,omitempty"`
	Status            string  `json:"status,omitempty"`
	SubaccountId      *uint16 `json:"subaccountId,omitempty"`
	ToAddress         string  `json:"toAddress,omitempty"`
	TransactionHash   *string `json:"transactionHash,omitempty"`
	CreatedAt         string  `json:"createdAt,omitempty"`
	IsInternal        bool    `json:"isInternal,omitempty"`
	BankName          *string `json:"bankName,omitempty"`
	BankIdentifier    *string `json:"bankIdentifier,omitempty"`
	AccountIdentifier *string `json:"accountIdentifier,omitempty"`
	TriggerAt         *string `json:"triggerAt,omitempty"`
}

type RequestWithdrawalReq struct {
	Address        string  `json:"address,omitempty"`
	AutoBorrow     *bool   `json:"autoBorrow,omitempty"`
	AutoLendRedeem *bool   `json:"autoLendRedeem,omitempty"`
	Blockchain     string  `json:"blockchain,omitempty"`
	ClientId       *string `json:"clientId,omitempty"`
	Quantity       string  `json:"quantity,omitempty"`
	Symbol         string  `json:"symbol,omitempty"`
	TwoFactorToken *string `json:"twoFactorToken,omitempty"`
}

func (req RequestWithdrawalReq) Validate() error {
	if req.Address == "" {
		return errors.New("address is required")
	}
	if req.Blockchain == "" {
		return errors.New("blockchain is required")
	}
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	if req.Quantity == "" {
		return errors.New("quantity is required")
	}
	return nil
}

func (req RequestWithdrawalReq) BuildQueryParams() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(fmt.Sprintf("address=%s", req.Address))
	if req.AutoBorrow != nil {
		strBuilder.WriteString(fmt.Sprintf("&autoBorrow=%t", *req.AutoBorrow))
	}
	if req.AutoLendRedeem != nil {
		strBuilder.WriteString(fmt.Sprintf("&autoLendRedeem=%t", *req.AutoLendRedeem))
	}
	strBuilder.WriteString(fmt.Sprintf("&blockchain=%s", req.Blockchain))
	if req.ClientId != nil {
		strBuilder.WriteString(fmt.Sprintf("&clientId=%s", *req.ClientId))
	}
	strBuilder.WriteString(fmt.Sprintf("&quantity=%s", req.Quantity))
	strBuilder.WriteString(fmt.Sprintf("&symbol=%s", req.Symbol))
	if req.TwoFactorToken != nil {
		strBuilder.WriteString(fmt.Sprintf("&twoFactorToken=%s", *req.TwoFactorToken))
	}
	return strBuilder.String()
}

func (req RequestWithdrawalReq) Instruction() string {
	return InstructionWithdraw
}

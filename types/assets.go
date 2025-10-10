package types

type Token struct {
	DisplayName       string `json:"displayName"`
	Blockchain        string `json:"blockchain"`
	ContractAddress   string `json:"contractAddress"`
	DepositEnabled    bool   `json:"depositEnabled"`
	MinimumDeposit    string `json:"minimumDeposit"`
	WithdrawEnabled   bool   `json:"withdrawEnabled"`
	MinimumWithdrawal string `json:"minimumWithdrawal"`
	MaximumWithdrawal string `json:"maximumWithdrawal"`
	WithdrawalFee     string `json:"withdrawalFee"`
}
type AssetResp struct {
	Symbol      string  `json:"symbol"`
	DisplayName string  `json:"displayName"`
	CoingeckoId string  `json:"coingeckoId"`
	Tokens      []Token `json:"tokens"`
}

type CollateralResp struct {
	Symbol          string          `json:"symbol"`
	ImfFunction     Function        `json:"imfFunction"`
	MmfFunction     Function        `json:"mmfFunction"`
	HaircutFunction HaircutFunction `json:"haircutFunction"`
}

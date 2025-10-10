package types

const (
	SystemStatusOk          = "Ok"
	SystemStatusMaintenance = "Maintenance"
)

type Status struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Wallet struct {
	Blockchain string `json:"blockchain"`
	Address    string `json:"address"`
}

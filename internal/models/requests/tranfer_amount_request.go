package requests

type TransferAmountRequest struct {
	FromAccNo string  `json:"from"`
	ToAccNo   string  `json:"to"`
	Amount    float64 `json:"amount"`
}

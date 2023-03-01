package db

import (
	"bank/internal/models/response"
)

type BankExecutor interface {
	GetDetailByAccNumber(string) (*response.AccountDetail, error)
	TransferMoney(string, string, float64) (*response.TransferAmountResponse, error)
}

package response

import "time"

type AccountDetail struct {
	AccNumber    string        `json:"acc_number"`
	Balance      float64       `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	TransactionId     string    `json:"acc_number"`
	FromAccNo         string    `json:"from_acc_no"`
	ToAccNo           string    `json:"to_acc_no"`
	TransferredAmount float64   `json:"transferred_amount"`
	CreatedDatetime   time.Time `json:"created_datetime"`
}

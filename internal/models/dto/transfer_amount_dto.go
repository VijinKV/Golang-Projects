package dto

import (
	"time"
)

type TransferAmountDto struct {
	Id              string           `json:"id"`
	From            AccountDetailDto `json:"from"`
	To              AccountDetailDto `json:"to"`
	Transferred     float64          `json:"transferred"`
	CreatedDatetime time.Time        `json:"created_datetime"`
}

package response

import (
	"time"
)

type TransferAmountResponse struct {
	Id              string                `json:"id"`
	From            AccountDetailResponse `json:"from"`
	To              AccountDetailResponse `json:"to"`
	Transferred     float64               `json:"transferred"`
	CreatedDatetime time.Time             `json:"created_datetime"`
}

//func ConvertDTOToResponseModel(transferAmountDto dto.TransferAmountDto) (*TransferAmountResponse, error) {
//	// Convert the DTO to JSON
//	dtoJson, err := json.Marshal(transferAmountDto)
//	if err != nil {
//		return nil, err
//	}
//
//	// Unmarshal the JSON to the model struct
//	var response *TransferAmountResponse
//	err = json.Unmarshal(dtoJson, response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
